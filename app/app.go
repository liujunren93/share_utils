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
	"github.com/liujunren93/share_utils/wrapper/recover"
	"github.com/mitchellh/mapstructure"
)

type confType int8

const (
	localConf confType = iota // local
	CloudConf                 // config center
)

var configPath string

func init() {
	// flag.StringVar(&configPath, "config")
	flag.StringVar(&configPath, "c", "./conf/config.yaml", "init config")
	// flag.Parse()
}

type appConfigOption struct {
	localConf   entity.LocalConfiger // 启动app基础配置
	cloudConfig entity.ClubConfiger  // 配置中心默认配置
	Cloud       cfg.Configer         // 配置中心
	Local       cfg.Configer         // 本地配置

}
type App struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	appConfigOption
	shareGrpcClient  *client.Client // grpc client
	monitorsCh       chan *config.Monitor
	localMonitorOnce *sync.Once
	appRouter        router                    // 自动路由
	rc               routerCenter.RouterCenter // 路由中心
	stopList         []func()
}

func NewApp(localConf entity.LocalConfiger, clubConfig entity.ClubConfiger) *App {
	ctx, cancel := context.WithCancel(context.TODO())
	app := &App{
		ctx:        ctx,
		cancelFunc: cancel,
		appConfigOption: appConfigOption{
			cloudConfig: clubConfig,
			localConf:   localConf,
		},
		monitorsCh:       config.InitRegistryMonitor(),
		localMonitorOnce: &sync.Once{},
	}
	if app.appConfigOption.cloudConfig == nil {
		app.appConfigOption.cloudConfig = entity.DefaultConfig
	}
	if app.appConfigOption.localConf == nil {

		app.appConfigOption.localConf = new(entity.LocalBase)
	}
	app.initConfig()
	return app
}

// app 停止
func (a *App) RegistryStopFunc(f func()) {
	a.stopList = append(a.stopList, f)
}

// func (a *App) RegistryTask(f func(ctx context.Context)) {
// 	a.stopList = append(a.stopList, f)
// }

func (a *App) GetAppName() string {
	return a.localConf.GetLocalBase().AppName
}

func (a *App) GetClubConfig() entity.ClubConfiger {
	if a.cloudConfig.GetVersion() == "" {
		panic("cloud config was not init")
	}
	return a.cloudConfig
}

func (a *App) GetLocalConfig() entity.LocalConfiger {
	return a.localConf
}

func (a *App) CloudConfigMonitor(confName, group, fieldName string, callbacks ...func()) {
	if confName == "" {
		confName = a.localConf.GetLocalBase().ConfCenter.ConfName
	}
	if group == "" {
		group = a.localConf.GetLocalBase().ConfCenter.Group
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

// initConfig 初始化配置
func (a *App) initConfig() {
	if !flag.Parsed() {
		flag.Parse()
	}
	var fileType, confName string
	a.Local, fileType, confName = store.NewViper(configPath)

	err := a.Local.GetConfig(context.Background(), fileType, confName, func(confName, group string, content interface{}) error {
		val := content.(map[string]interface{})
		return mapstructure.Decode(val, &a.localConf)
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("local config:%+v\n", a.localConf)

	if a.localConf.GetLocalBase().ConfCenter.Enable {
		a.initConfCenter()
	}
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	localConf := a.localConf.GetLocalBase()
	err = a.Cloud.GetConfig(ctx, localConf.ConfCenter.ConfName, localConf.ConfCenter.Group, config.DescConfig(a.cloudConfig))
	if err != nil {
		fmt.Println("get Config from cloud err:" + err.Error())
		// panic("get Config from cloud err:" + err.Error())

		return

	}
	a.initLogger()
	go func() {
		err := a.Cloud.ListenConfig(a.ctx, localConf.ConfCenter.ConfName, localConf.ConfCenter.Group, config.DescConfigAndCallbacks(a.Cloud))
		if err != nil {
			log.Logger.Error(err)
		}
	}()

}

// confType :localConf CloudConf
// if localConf group=fileType confName=fileName
func (a *App) GetConfig(ct confType, ctx context.Context, confName, group string, dest interface{}) error {
	if ct == localConf {
		return a.Local.GetConfig(ctx, group, confName, config.DescConfig(dest))
	} else {
		return a.Cloud.GetConfig(ctx, confName, group, config.DescConfig(dest))
	}
}

func (a *App) initConfCenter() {
	localConf := a.localConf.GetLocalBase()
	switch localConf.ConfCenter.Type {
	case "redis":

		client, err := redis.NewClient(localConf.ConfCenter.Config)
		if err != nil {
			panic(err)
		}
		a.Cloud = store.NewRedis(client, localConf.Namespace)
	}
}

func (a *App) initLogger() {
	if logConf, ok := a.cloudConfig.GetLogConfig(); ok {
		log.Init(logConf)
	}
	shLog.Logger = log.Logger.GetLogrus()
	localConf := a.localConf.GetLocalBase()
	a.CloudConfigMonitor(localConf.ConfCenter.ConfName, localConf.GetLocalBase().ConfCenter.Group, "Log", func() {
		if logConf, ok := a.cloudConfig.GetLogConfig(); ok {
			log.Upgrade(logConf)
		}

	})
}

func (a *App) GetGrpcClient(targetUrl string) (*client.Client, error) {
	if a.shareGrpcClient == nil {
		var utilsGrpcClient *grpc.Client
		registryConf, ok := a.cloudConfig.GetRegistryConfig()
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
	localConf := a.localConf.GetLocalBase()
	if localConf.RunMode == "debug" {
		a.shareGrpcClient.AddOptions(client.WithTimeout(time.Second * 30))
	}
	if localConf.Namespace != "" {
		a.shareGrpcClient.AddOptions(client.WithNamespace(localConf.Namespace))
	}
	return a.shareGrpcClient, nil

}

func (a *App) RunGw(f func(*gin.Engine) (shareRouter.Router, error), middlewares ...MiddlewareItem) error {
	eng := gin.Default()
	localConf := a.localConf.GetLocalBase()
	eng.Use(middleware.Cors)
	if localConf.RunMode == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	router, err := f(eng)
	if err != nil {
		return err
	}
	if a.cloudConfig.GetRouterCenter().Enable {
		for _, v := range middlewares {
			if a.appRouter.middlewares == nil {
				a.appRouter.middlewares = make(map[string]middlewareFunc)
			}
			a.appRouter.middlewares[v.Name] = v.MiddlewareFunc
		}
		a.autoRoute(router)
	}

	go func() {
		err := eng.Run(localConf.ListenAddr)
		if err != nil {
			panic(err)
		}
	}()
	a.watchSignal()
	return nil
}

func (a *App) RunRpc(f func(ser *server.GrpcServer) error) error {
	localConf := a.localConf.GetLocalBase()

	s := utilsServer.Server{ListenAddr: localConf.ListenAddr, Mode: localConf.RunMode, Namespace: localConf.Namespace, ServerName: localConf.AppName}
	if reg, ok := a.cloudConfig.GetRegistryConfig(); ok {
		if reg.Etcd != nil {
			s.RegistryAddr = reg.Etcd.Endpoints
		}

	}

	gs, err := s.NewServer(server.WithHdlrWrappers(recover.Recover()))
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	err = f(gs)
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	go func() {
		err := gs.Run()
		if err != nil {
			panic(err)
		}
	}()
	a.watchSignal()
	return nil

}

func (a *App) watchSignal() {

	ch := make(chan os.Signal, 1)
	signals := []os.Signal{
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL,
	}
	signal.Notify(ch, signals...)
	select {
	// wait on kill signal
	case sign := <-ch:
		a.cancelFunc()
		for _, stop := range a.stopList {
			stop()
		}

		time.Sleep(time.Second * 4)
		log.Logger.Info("app exit", sign)
		os.Exit(0)
	}

}
