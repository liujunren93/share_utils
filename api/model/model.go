package model

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/core/registry"
	"github.com/liujunren93/share/core/registry/etcd"
	"github.com/liujunren93/share/plugins/opentrace"
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

type Model struct {
	Ctx       *gin.Context
	Redis     *redis.Client
	UserStore UserStore
	EtcdAddr  []string
	OpenTrace OpenTrace
}
type OpenTrace struct {
	ServerName string
	OpenTrace  string
}

type UserStore struct {
	KeepLoginTime int64
	Secret        string
}

func (m *Model) GClient(serverName string) (*grpc.ClientConn, error) {

	getClientOnce.Do(func() {
		newJaeger, _, err := openTrace.NewJaeger(m.OpenTrace.ServerName, m.OpenTrace.OpenTrace)
		if err != nil {
			log.Logger.Error(err)
		}
		openTracer = newJaeger

		opentracing.SetGlobalTracer(newJaeger)
		r, err := etcd.NewRegistry(registry.WithAddrs(m.EtcdAddr...))
		thisClient = client.NewClient(client.WithRegistry(r))
	})
	// agent
	newUserStore := userStore.NewUserStore(m.UserStore.KeepLoginTime, m.UserStore.Secret, m.Redis)
	if load, ok := newUserStore.Load(m.Ctx); ok {
		var agent metadata2.UserAgent
		agent = load.UserAgent
		thisClient.AddOptions(client.WithCallWrappers(opentrace.ClientGrpcCallWrap(openTracer), metadata.ClientUACallWrap(&agent)))
	}
	return thisClient.Client(serverName)
}
