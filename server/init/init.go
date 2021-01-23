package init

import (
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/plugins/opentrace"
	"github.com/liujunren93/share/server"
	"github.com/liujunren93/share_utils/wrapper/openTrace"
	"github.com/opentracing/opentracing-go"
)

type Server struct {
	EtcdAddr   []string
	ServerName string
	Namespace  string
	OpenTrace  string
	Weight     int // 权重
	opts       []server.Option
}

func (s *Server) NewServer(opts ...server.Option) (*server.GrpcServer, error) {
	err := s.initOpts(opts)
	if err != nil {
		return nil, err
	}
	//注册中心
	newRegistry, err := etcd.NewRegistry(registry.WithAddrs(s.EtcdAddr...))
	if err != nil {
		return nil, err
	}

	if s.Namespace != "" {
		s.opts = append(s.opts, server.WithNamespace(s.Namespace))
	}
	grpcServer := server.NewGrpcServer(s.opts...)
	err = grpcServer.Registry(newRegistry, registry.WithWeight(s.Weight))
	if err != nil {
		return nil, err
	}
	return grpcServer, nil
}

func (s *Server) initOpts(opts []server.Option) error {

	// 链路
	jaeger, closer, err := openTrace.NewJaeger(s.ServerName, s.OpenTrace)
	defer closer.Close()
	if err != nil {
		return err
	}

	opentracing.SetGlobalTracer(jaeger)
	s.opts = append(s.opts, server.WithName(s.ServerName),
		server.WithHdlrWrappers(opentrace.ServerGrpcWrap(jaeger)))
	s.opts = append(s.opts, opts...)
	return nil
}
