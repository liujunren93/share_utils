package client

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/wrapper/opentrace"
	"github.com/liujunren93/share_utils/log"
	metadata2 "github.com/liujunren93/share_utils/metadata"
	"github.com/liujunren93/share_utils/pkg/storage/userStore"
	"github.com/liujunren93/share_utils/wrapper/metadata"
	"github.com/liujunren93/share_utils/wrapper/openTrace"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"sync"
)

var (
	openTracer    opentracing.Tracer
	thisClient    *client.Client
	getClientOnce sync.Once
)

type Client struct {
	Ctx       context.Context
	Redis     *redis.Client
	UserStore UserStore
	EtcdAddr  []string
	OpenTrace OpenTrace
	Namespace string
	Balancer  string
}
type OpenTrace struct {
	ClientName string //
	OpenTrace  string
}

type UserStore struct {
	KeepLoginTime int64
	Secret        string
}

func (c *Client) GetGrpcClient(serverName string) (*grpc.ClientConn, error) {
	getClientOnce.Do(func() {

		if len(c.EtcdAddr) == 0 {
			log.Logger.Panic("register address nil")
		}
		r, err := etcd.NewRegistry(registry.WithAddrs(c.EtcdAddr...))
		if err != nil {
			log.Logger.Error("registry err ", err)
		}
		// 获取share 客户端
		thisClient = client.NewClient(client.WithRegistry(r),client.WithNamespace(c.Namespace))
		if openTracer != nil {
			newJaeger, _, err := openTrace.NewJaeger(c.OpenTrace.ClientName, c.OpenTrace.OpenTrace)
			if err != nil {
				log.Logger.Error(err)
				return
			} else {
				openTracer = newJaeger
				opentracing.SetGlobalTracer(newJaeger)
			}
			thisClient.AddOptions(client.WithCallWrappers(opentrace.NewClientWrapper(openTracer)))
		}
	})

	// agent
	newUserStore := userStore.NewUserStore(c.UserStore.KeepLoginTime, c.UserStore.Secret, c.Redis)
	if ctx, ok := c.Ctx.(*gin.Context); ok {
		if load, ok := newUserStore.Load(ctx); ok {
			var agent metadata2.UserAgent
			agent = load.UserAgent
			thisClient.AddOptions(client.WithCallWrappers(metadata.ClientUACallWrap(&agent)))
		}
	}
	thisClient.AddOptions(client.WithBalancer(c.Balancer))
	return thisClient.Client(serverName)
}
