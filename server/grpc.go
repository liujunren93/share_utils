package server

import (
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/server"
	"github.com/liujunren93/share_utils/common/opentrace/openJaeger"
	"github.com/liujunren93/share_utils/wrapper/opentrace"
	"github.com/opentracing/opentracing-go"
)

type Server struct {
	Mode         string
	RegistryAddr []string
	ServerName   string
	Namespace    string
	OpenTrace    string
	Weight       int // 权重
	Address      string
	opts         []server.Option
}

func (s *Server) NewServer(opts ...server.Option) (*server.GrpcServer, error) {
	err := s.initOpts(opts)
	if err != nil {
		return nil, err
	}

	if s.Namespace != "" {
		s.opts = append(s.opts, server.WithNamespace(s.Namespace))
	}
	if s.Address != "" {
		s.opts = append(s.opts, server.WithAddress(s.Address))
	}
	grpcServer := server.NewGrpcServer(s.opts...)
	//注册中心
	if len(s.RegistryAddr) > 0 {
		newRegistry, err := etcd.NewRegistry(registry.WithAddrs(s.RegistryAddr...))
		if err != nil {
			return nil, err
		}
		err = grpcServer.Registry(newRegistry, registry.WithWeight(s.Weight))
		if err != nil {
			return nil, err
		}
	}

	return grpcServer, nil
}

func (s *Server) initOpts(opts []server.Option) error {
	if s.OpenTrace != "" {
		// 链路
		jaeger, _, err := openJaeger.NewJaeger(s.ServerName, s.OpenTrace)

		if err != nil {
			return err
		}
		opentracing.SetGlobalTracer(jaeger)
		s.opts = append(s.opts, server.WithHdlrWrappers(opentrace.NewServerWrapper(jaeger)))
	}
	if s.ServerName != "" {
		s.opts = append(s.opts, server.WithName(s.ServerName))
	}

	s.opts = append(s.opts, server.WithMode(s.Mode))
	s.opts = append(s.opts, opts...)
	return nil
}
