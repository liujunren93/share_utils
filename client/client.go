package client

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/wrapper/opentrace"
	"github.com/liujunren93/share_utils/auth/jwt"
	"github.com/liujunren93/share_utils/log"
	metadata2 "github.com/liujunren93/share_utils/metadata"
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
	ctx       context.Context
	redis     *redis.Client
	userStore UserStore
	etcdAddr  []string
	openTrace OpenTrace
	namespace string
	balancer  string
}
type option func(*Client)

func NewClient(opts ...option) *Client {
	cli := new(Client)
	for _, opt := range opts {
		opt(cli)
	}
	return cli

}
func WithRedis(redis *redis.Client) option {
	return func(c *Client) {
		c.redis = redis
	}
}

func WithUserStore(us UserStore) option {
	return func(c *Client) {
		c.userStore = us
	}
}

func WithEtcdAddr(addrs ...string) option {
	return func(c *Client) {
		c.etcdAddr = addrs
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

func (c *Client) WithCtx(ctx context.Context) {
	c.ctx = ctx
}
func (c *Client) GetCtx() context.Context {
	return c.ctx
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

		if len(c.etcdAddr) == 0 {
			log.Logger.Panic("register address nil")
		}
		r, err := etcd.NewRegistry(registry.WithAddrs(c.etcdAddr...))
		if err != nil {
			log.Logger.Error("registry err ", err)
		}
		// 获取share 客户端
		thisClient = client.NewClient(client.WithRegistry(r), client.WithNamespace(c.namespace))
		if openTracer != nil {
			newJaeger, _, err := openTrace.NewJaeger(c.openTrace.ClientName, c.openTrace.OpenTrace)
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
	//newUserStore := userStore.NewUserStore(c.userStore.KeepLoginTime, c.userStore.Secret, c.redis)
	//if ctx, ok := c.ctx.(*gin.Context); ok {
	//	if load, ok := newUserStore.Load(ctx); ok {
	//		var agent metadata2.UserAgent
	//		agent = load.UserAgent
	//		thisClient.AddOptions(client.WithCallWrappers(metadata.ClientUACallWrap(&agent)))
	//	}
	//}
	if ctx, ok := c.ctx.(*gin.Context); ok {
		var agent metadata2.UserAgent
		if get, exists := ctx.Get("ua"); exists {
			if data,ok := get.(*jwt.JwtClaims);ok{
				if m,ok:=data.Data.(map[string]interface{});ok{
					agent.LoginTime=int64(m["login_time"].(float64))
					agent.AppID=uint(m["app_id"].(float64))
					agent.UID=uint(m["uid"].(float64))
				}

			}

		}
		thisClient.AddOptions(client.WithCallWrappers(metadata.ClientUACallWrap(&agent)))
	}
	thisClient.AddOptions(client.WithBalancer(c.balancer))
	return thisClient.Client(serverName)
}
