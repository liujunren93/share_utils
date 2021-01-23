package init

import (
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/plugins/opentrace"
	"github.com/liujunren93/share/plugins/validator"
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
}

func (s *Server) NewServer() (*server.GrpcServer, error) {
	//注册中心
	newRegistry, err := etcd.NewRegistry(registry.WithAddrs(s.EtcdAddr...))
	if err != nil {
		return nil, err
	}
	// 链路
	jaeger, closer, err := openTrace.NewJaeger(s.ServerName, s.OpenTrace)
	defer closer.Close()
	if err != nil {
		return nil, err
	}
	opentracing.SetGlobalTracer(jaeger)
	var opts []server.Option
	opts = append(opts, server.WithName(s.ServerName),
		server.WithHdlrWrappers(opentrace.ServerGrpcWrap(jaeger), validator.NewHandlerWrapper()))
	if s.Namespace != "" {
		opts = append(opts, server.WithNamespace(s.Namespace))
	}
	grpcServer := server.NewGrpcServer(opts...)
	err = grpcServer.Registry(newRegistry, registry.WithWeight(s.Weight))
	if err != nil {
		return nil, err
	}
	return grpcServer, nil
}
