package app

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	re "github.com/go-redis/redis/v8"
	"github.com/liujunren93/share/client"
	"github.com/liujunren93/share/server"
	"github.com/liujunren93/share_utils/app/config"
	"github.com/liujunren93/share_utils/app/config/entity"
	"github.com/liujunren93/share_utils/client/grpc"
	cfg "github.com/liujunren93/share_utils/common/config"
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

var configPath string

func init() {
	// flag.StringVar(&configPath, "config")
	flag.StringVar(&configPath, "c", "./conf/config.yaml", "init config")
	// flag.Parse()
}

type appConfigOption struct {
	LocalConf   *entity.LocalBase // 启动app基础配置
	defaultConf *entity.Config    // 配置中心默认配置
	Cloud       cfg.Configer      // 配置中心
	Local       cfg.Configer      // 本地配置

}
type App struct {
	ctx context.Context
	appConfigOption
	shareGrpcClient  *client.Client
	monitorsCh       chan *config.Monitor
	localMonitorOnce *sync.Once
}

func NewApp(ctx context.Context) *App {
	app := &App{
		ctx: ctx,
		appConfigOption: appConfigOption{
			defaultConf: &entity.DefaultConfig,
		},
		monitorsCh:       config.InitRegistryMonitor(),
		localMonitorOnce: &sync.Once{},
	}
	app.initConfig()
	return app
}

func (a *App) GetDefaultConfig() *entity.Config {
	if a.defaultConf.Version == 0 {
		panic("cloud config was not init")
	}
	return a.defaultConf
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
	if !flag.Parsed() {
		flag.Parse()
	}
	var fileType, confName string
	a.Local, fileType, confName = store.NewViper(configPath)

	err := a.Local.GetConfig(context.Background(), fileType, confName, func(confName, group string, content interface{}) error {
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
	err = a.Cloud.GetConfig(ctx, a.LocalConf.ConfCenter.ConfName, a.LocalConf.ConfCenter.Group, config.DescConfig(a.defaultConf))
	if err != nil {
		fmt.Println("get Config from cloud err:" + err.Error())
	}
	a.initLogger()
	go func() {
		err := a.Cloud.ListenConfig(a.ctx, a.LocalConf.ConfCenter.ConfName, a.LocalConf.ConfCenter.Group, config.DescConfigAndCallbacks(a.defaultConf))
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
	if a.defaultConf.Log != nil {
		log.Init(a.defaultConf.Log)
	}
	a.CloudConfigMonitor(a.LocalConf.ConfCenter.ConfName, a.LocalConf.ConfCenter.Group, func() {
		log.Upgrade(a.defaultConf.Log)
	})
}

func (a *App) GetGrpcClient(targetUrl string) (*client.Client, error) {
	if a.shareGrpcClient == nil {
		var utilsGrpcClient *grpc.Client
		if a.defaultConf.Registry != nil || a.defaultConf.Registry.Enable {
			utilsGrpcClient = grpc.NewClient(grpc.WithEtcdAddr(a.defaultConf.Registry.Etcd.Endpoints...))
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
