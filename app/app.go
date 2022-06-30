package app

import (
	"context"
	"flag"
	"time"

	"github.com/gin-gonic/gin"
	re "github.com/go-redis/redis/v8"
	"github.com/liujunren93/share/server"
	"github.com/liujunren93/share_utils/common/config/store"
	"github.com/liujunren93/share_utils/databases/redis"
	"github.com/liujunren93/share_utils/log"
	"github.com/liujunren93/share_utils/middleware"
	utilsServer "github.com/liujunren93/share_utils/server"
	"github.com/mitchellh/mapstructure"
)

type App struct {
	appConfigOption
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

func NewApp(ctx context.Context) *App {
	return &App{
		appConfigOption: appConfigOption{ctx: ctx},
	}
}

func (a *App) AddConfigMonitor(f ...func()) {
	a.configMonitors = append(a.configMonitors, f...)
}

func (a *App) InitConfig(conf interface{}) {
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
	err = a.Configer.GetConfig(ctx, a.LocalConf.ConfCenter.ConfName, a.LocalConf.ConfCenter.Group, DescConfig(conf))
	if err != nil {
		panic(err)
	}

	go func() {
		err := a.Configer.ListenConfig(a.ctx, a.LocalConf.ConfCenter.ConfName, a.LocalConf.ConfCenter.Group, DescConfigAndCallbacks(conf, a.configMonitors))
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
