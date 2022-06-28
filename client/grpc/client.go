package grpc

import (
	"sync"

	"github.com/liujunren93/share_utils/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/wrapper"
	"github.com/liujunren93/share_utils/common/opentrace/openJaeger"
	"github.com/liujunren93/share_utils/wrapper/opentrace"
	"github.com/opentracing/opentracing-go"
)

var (
	openTracer    opentracing.Tracer
	getClientOnce sync.Once
)

type Client struct {
	registryAddr    []string
	openTrace       OpenTrace
	namespace       string
	balancer        string
	buildTargetFunc client.BuildTargetFunc
	wraps           []wrapper.CallWrapper
}
type option func(*Client)

func NewClient(opts ...option) *Client {
	cli := new(Client)
	for _, opt := range opts {
		opt(cli)
	}
	return cli

}

func WithBuildTargetFunc(buildTargetFunc client.BuildTargetFunc) option {
	return func(c *Client) {
		c.buildTargetFunc = buildTargetFunc
	}
}

func WithEtcdAddr(addrs ...string) option {
	return func(c *Client) {
		c.registryAddr = addrs
	}
}
func WithOpenTrace(openTrace OpenTrace) option {
	return func(c *Client) {
		c.openTrace = openTrace
	}
}

func WithNamespace(namespace string) option {
	return func(c *Client) {
		c.namespace = namespace
	}
}
func WithBalancer(balancer string) option {
	return func(c *Client) {
		c.balancer = balancer
	}
}

type OpenTrace struct {
	ClientName string //
	OpenTrace  string
}

var shareClient *client.Client

func (c *Client) GetShareClient() (*client.Client, error) {

	var err error
	getClientOnce.Do(func() {
		// 获取share 客户端
		shareClient = client.NewClient(client.WithNamespace(c.namespace), client.WithBuildTargetFunc(c.buildTargetFunc))
		if len(c.openTrace.OpenTrace) != 0 {
			newJaeger, _, err := openJaeger.NewJaeger(c.openTrace.ClientName, c.openTrace.OpenTrace)
			if err != nil {
				log.Logger.Error(err)
				return
			} else {
				openTracer = newJaeger
				opentracing.SetGlobalTracer(newJaeger)
			}
			shareClient.AddOptions(client.WithCallWrappers(opentrace.NewClientWrapper(openTracer)))
		}
		if len(c.registryAddr) != 0 {
			r, err := etcd.NewRegistry(registry.WithAddrs(c.registryAddr...))
			if err != nil {
				return
			}
			shareClient.AddOptions(client.WithRegistry(r))

		}
	})
	if c.balancer != "" {
		shareClient.AddOptions(client.WithBalancer(c.balancer))
	}

	shareClient.AddOptions(client.WithGrpcDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())))
	return shareClient, err
}
