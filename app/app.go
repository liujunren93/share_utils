package app

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share/client"
	shLog "github.com/liujunren93/share/log"
	"github.com/liujunren93/share/server"
	"github.com/liujunren93/share_utils/app/config"
	"github.com/liujunren93/share_utils/app/config/entity"
	"github.com/liujunren93/share_utils/client/grpc"
	cfg "github.com/liujunren93/share_utils/common/config"
	"github.com/liujunren93/share_utils/common/config/store"
	shareRouter "github.com/liujunren93/share_utils/common/gin/router"
	"github.com/liujunren93/share_utils/db/redis"
	"github.com/liujunren93/share_utils/log"
	"github.com/liujunren93/share_utils/middleware"
	"github.com/liujunren93/share_utils/pkg/routerCenter"
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
	LocalConf  *entity.LocalBase   // 启动app基础配置
	baseConfig entity.BaseConfiger // 配置中心默认配置
	Cloud      cfg.Configer        // 配置中心
	Local      cfg.Configer        // 本地配置

}
type App struct {
	ctx context.Context
	appConfigOption
	shareGrpcClient  *client.Client
	monitorsCh       chan *config.Monitor
	localMonitorOnce *sync.Once
	appRouter        map[string]*shareRouter.Node
	rc               routerCenter.RouterCenter
	stopList         []func()
}

func NewApp(defaultConfig entity.BaseConfiger) *App {
	app := &App{
		ctx: context.TODO(),
		appConfigOption: appConfigOption{
			baseConfig: defaultConfig,
		},
		monitorsCh:       config.InitRegistryMonitor(),
		localMonitorOnce: &sync.Once{},
	}
	if app.appConfigOption.baseConfig == nil {
		app.appConfigOption.baseConfig = entity.DefaultConfig
	}
	app.initConfig()
	app.watchSignal()
	return app
}

func (a *App) RegistryStopFunc(f func()) {
	a.stopList = append(a.stopList, f)
}
func (a *App) watchSignal() {
	ctx, cancel := context.WithCancel(a.ctx)
	a.ctx = ctx
	go func() {
		ch := make(chan os.Signal, 1)
		signals := []os.Signal{
			syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL,
		}
		signal.Notify(ch, signals...)
		select {
		// wait on kill signal
		case sign := <-ch:
			cancel()
			for _, stop := range a.stopList {
				stop()
			}
			log.Logger.Info("app exit", sign)
			os.Exit(0)
		}
	}()

}

func (a *App) GetAppName() string {
	return a.LocalConf.AppName
}

func (a *App) GetBaseConfig() entity.BaseConfiger {
	if a.baseConfig.GetVersion() == "" {
		panic("cloud config was not init")
	}
	return a.baseConfig
}

func (a *App) CloudConfigMonitor(confName, group, fieldName string, callbacks ...func()) {
	if confName == "" {
		confName = a.LocalConf.ConfCenter.ConfName
	}
	if group == "" {
		group = a.LocalConf.ConfCenter.Group
	}
	for _, callback := range callbacks {
		a.monitorsCh <- config.NewMonitor(confName, group, fieldName, callback)
	}
}

func (a *App) LocalConfigMonitor(fileType, fileName, fieldName string, dest interface{}, callbacks ...func()) {
	a.localMonitorOnce.Do(func() {
		go func() {
			a.Local.ListenConfig(a.ctx, fileType, fileName, config.DescConfigAndCallbacks(dest))
		}()
	})
	for _, callback := range callbacks {
		a.monitorsCh <- config.NewMonitor(fileType, fileName, fieldName, callback)
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
	fmt.Printf("local config:%+v\n", a.LocalConf)

	if a.LocalConf.ConfCenter.Enable {
		a.initConfCenter()
	}
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	err = a.Cloud.GetConfig(ctx, a.LocalConf.ConfCenter.ConfName, a.LocalConf.ConfCenter.Group, config.DescConfig(a.baseConfig))
	if err != nil {
		fmt.Println("get Config from cloud err:" + err.Error())
		panic("get Config from cloud err:" + err.Error())

	}
	a.initLogger()
	go func() {
		err := a.Cloud.ListenConfig(a.ctx, a.LocalConf.ConfCenter.ConfName, a.LocalConf.ConfCenter.Group, config.DescConfigAndCallbacks(a.baseConfig))
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

		client, err := redis.NewClient(a.LocalConf.ConfCenter.Config)
		if err != nil {
			panic(err)
		}
		a.Cloud = store.NewRedis(client, a.LocalConf.Namespace)
	}
}

func (a *App) initLogger() {
	if logConf, ok := a.baseConfig.GetLogConfig(); ok {
		log.Init(logConf)
	}
	shLog.Logger = log.Logger
	a.CloudConfigMonitor(a.LocalConf.ConfCenter.ConfName, a.LocalConf.ConfCenter.Group, "Log", func() {
		if logConf, ok := a.baseConfig.GetLogConfig(); ok {
			log.Upgrade(logConf)
		}

	})
}

func (a *App) GetGrpcClient(targetUrl string) (*client.Client, error) {
	if a.shareGrpcClient == nil {
		var utilsGrpcClient *grpc.Client
		registryConf, ok := a.baseConfig.GetRegistryConfig()
		if ok || registryConf.Enable {
			utilsGrpcClient = grpc.NewClient(grpc.WithEtcdAddr(registryConf.Etcd.Endpoints...))
		} else {
			utilsGrpcClient = grpc.NewClient(grpc.WithBuildTargetFunc(func(args ...string) string { return targetUrl }))
		}
		shareClient, err := utilsGrpcClient.GetShareClient()
		if err != nil {
			return nil, err
		}
		a.shareGrpcClient = shareClient

	}
	if a.LocalConf.RunMode == "debug" {
		a.shareGrpcClient.AddOptions(client.WithTimeout(time.Second * 30))
	}
	return a.shareGrpcClient, nil

}

func (a *App) RunGw(f func(*gin.Engine) (shareRouter.Router, error)) error {
	eng := gin.Default()
	eng.Use(middleware.Cors)
	if a.LocalConf.RunMode == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	router, err := f(eng)
	if err != nil {
		return err
	}
	a.AutoRoute(router)
	return eng.Run(a.LocalConf.ListenAddr)
}

func (a *App) RunRpc(registryAddr []string, f func(ser *server.GrpcServer) error) error {
	s := utilsServer.Server{ListenAddr: a.LocalConf.ListenAddr, Mode: a.LocalConf.RunMode, Namespace: a.LocalConf.Namespace, ServerName: a.LocalConf.AppName}
	s.RegistryAddr = registryAddr

	gs, err := s.NewServer()
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	err = f(gs)
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	return gs.Run()

}
