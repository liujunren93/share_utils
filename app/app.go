package app

import (
	"context"
	"flag"
	"sync"
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

type confType int8

const (
	LocalConf confType = iota // local
	CloudConf                 // config center
)

type App struct {
	ctx context.Context
	config.AppConfigOption
	shareGrpcClient  *client.Client
	monitorsCh       chan *config.Monitor
	localMonitorOnce *sync.Once
}

func NewApp(ctx context.Context) *App {
	app := &App{
		ctx: ctx,
		AppConfigOption: config.AppConfigOption{
			BaseConf: &entity.DefaultConfig,
		},
		monitorsCh:       config.InitRegistryMonitor(),
		localMonitorOnce: &sync.Once{},
	}
	app.initConfig()
	return app
}

func (a *App) GetDefaultConfig() *entity.Config {
	return a.BaseConf
}

func (a *App) CloudConfigMonitor(confName, group string, callbacks ...func()) {
	if confName == "" {
		confName = a.LocalConf.ConfCenter.ConfName
	}
	if group == "" {
		group = a.LocalConf.ConfCenter.Group
	}
	for _, callback := range callbacks {
		a.monitorsCh <- config.NewMonitor(confName, group, callback)
	}
}

func (a *App) LocalConfigMonitor(fileType, fileName string, dest interface{}, callbacks ...func()) {
	a.localMonitorOnce.Do(func() {
		go func() {
			a.Local.ListenConfig(a.ctx, fileType, fileName, config.DescConfigAndCallbacks(dest))
		}()
	})
	for _, callback := range callbacks {
		a.monitorsCh <- config.NewMonitor(fileType, fileName, callback)
	}
}

//
func (a *App) initConfig() {
	var configPath string
	flag.StringVar(&configPath, "c", "./conf/", "local config path")
	a.Local = store.NewViper(configPath)

	err := a.Local.GetConfig(context.Background(), "yaml", "config.yaml", func(confName, group string, content interface{}) error {
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
	err = a.Cloud.GetConfig(ctx, a.LocalConf.ConfCenter.ConfName, a.LocalConf.ConfCenter.Group, config.DescConfig(a.BaseConf))
	if err != nil {
		panic("get Config err" + err.Error())
	}
	a.initLogger()
	go func() {
		err := a.Cloud.ListenConfig(a.ctx, a.LocalConf.ConfCenter.ConfName, a.LocalConf.ConfCenter.Group, config.DescConfigAndCallbacks(a.BaseConf))
		if err != nil {
			log.Logger.Error(err)
		}
	}()

}

// confType :LocalConf CloudConf
// if LocalConf group=fileType confName=fileName
func (a *App) GetConfig(ct confType, ctx context.Context, confName, group string, dest interface{}) error {
	if ct == LocalConf {
		return a.Local.GetConfig(ctx, group, confName, config.DescConfig(dest))
	} else {
		return a.Cloud.GetConfig(ctx, confName, group, config.DescConfig(dest))
	}
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
		a.Cloud = store.NewRedis(client, a.LocalConf.NameSpace)
	}
}

func (a *App) initLogger() {
	if a.BaseConf.Log != nil {
		log.Init(a.BaseConf.Log)
	}

}

func (a *App) GetGrpcClient(targetUrl string) (*client.Client, error) {
	if a.shareGrpcClient == nil {
		var utilsGrpcClient *grpc.Client
		if a.BaseConf.Registry != nil || a.BaseConf.Registry.Enable {
			utilsGrpcClient = grpc.NewClient(grpc.WithEtcdAddr(a.BaseConf.Registry.Etcd.Endpoints...))
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
