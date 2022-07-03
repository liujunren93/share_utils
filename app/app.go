package app

import (
	"context"
	"flag"
	"time"

	"github.com/gin-gonic/gin"
	re "github.com/go-redis/redis/v8"
	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/server"
	"github.com/liujunren93/share_utils/app/config"
	"github.com/liujunren93/share_utils/app/config/entity"
	"github.com/liujunren93/share_utils/client/grpc"
	"github.com/liujunren93/share_utils/common/config/store"
	"github.com/liujunren93/share_utils/databases/redis"
	"github.com/liujunren93/share_utils/log"
	"github.com/liujunren93/share_utils/middleware"
	utilsServer "github.com/liujunren93/share_utils/server"
	"github.com/mitchellh/mapstructure"
)

type App struct {
	ctx context.Context
	config.AppConfigOption
	shareGrpcClient *client.Client
	monitorsCh      chan func()
}

func NewApp(ctx context.Context) *App {
	return &App{
		ctx: ctx,
		AppConfigOption: config.AppConfigOption{
			Conf: &entity.DefaultConfig,
		},
		monitorsCh: config.InitRegistryMonitor(),
	}
}

func (a *App) GetConfig() *entity.Config {
	return a.Conf

}

func (a *App) RunGw(f func(*gin.Engine) error) error {
	eng := gin.Default()
	eng.Use(middleware.Cors)
	if a.LocalConf.RunMode == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	err := f(eng)
	if err != nil {
		return err
	}
	return eng.Run(a.LocalConf.HttpHost)
}

func (a *App) RunRpc(registryAddr []string, f func(ser *server.GrpcServer) error) error {
	s := utilsServer.Server{Address: a.LocalConf.HttpHost, Mode: a.LocalConf.RunMode, ServerName: a.LocalConf.AppName}
	s.RegistryAddr = registryAddr
	gs, err := s.NewServer()
	if err != nil {
		return err
	}
	return f(gs)
}

func (a *App) RegistryConfigMonitor(fs ...func()) {
	for _, f := range fs {
		a.monitorsCh <- f
	}
}

//
func (a *App) InitConfig() {
	var configPath string
	flag.StringVar(&configPath, "c", "./conf/config.yaml", "local config path")
	v := store.NewViper(configPath)

	err := v.GetConfig(context.Background(), "", "", func(content interface{}) error {
		val := content.(map[string]interface{})
		return mapstructure.Decode(val, &a.LocalConf)
	})
	if err != nil {
		panic(err)
	}
	if a.LocalConf.ConfCenter.Enable {
		a.initConfCenter()
	}
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	err = a.Configer.GetConfig(ctx, a.LocalConf.ConfCenter.ConfName, a.LocalConf.ConfCenter.Group, config.DescConfig(a.Conf))
	if err != nil {
		panic("get Config err" + err.Error())
	}
	a.initLogger()
	go func() {
		err := a.Configer.ListenConfig(a.ctx, a.LocalConf.ConfCenter.ConfName, a.LocalConf.ConfCenter.Group, config.DescConfigAndCallbacks(a.Conf))
		if err != nil {
			log.Logger.Error(err)
		}
	}()

}

func (a *App) initConfCenter() {
	switch a.LocalConf.ConfCenter.Type {
	case "redis":
		var conf re.Options
		a.LocalConf.ConfCenter.ToConfig(&conf)
		client, err := redis.NewRedis(&conf)
		if err != nil {
			panic(err)
		}
		a.Configer = store.NewRedis(client, a.LocalConf.NameSpace)
	}
}

func (a *App) initLogger() {
	if a.Conf.Log != nil {
		log.Init(a.Conf.Log)
	}

}

func (a *App) GetGrpcClient(targetUrl string) (*client.Client, error) {
	if a.shareGrpcClient == nil {
		var utilsGrpcClient *grpc.Client
		if a.Conf.Registry != nil || a.Conf.Registry.Enable {
			utilsGrpcClient = grpc.NewClient(grpc.WithEtcdAddr(a.Conf.Registry.Etcd.Endpoints...))
		} else {
			utilsGrpcClient = grpc.NewClient(grpc.WithBuildTargetFunc(func(args ...string) string { return targetUrl }))
		}
		shareClient, err := utilsGrpcClient.GetShareClient()
		if err != nil {
			return nil, err
		}
		a.shareGrpcClient = shareClient

	}
	return a.shareGrpcClient, nil

}
