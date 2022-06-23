package app

import (
	"context"
	"flag"

	"github.com/liujunren93/share_utils/log"

	"github.com/liujunren93/share_utils/common/config"
	"github.com/liujunren93/share_utils/common/config/entity"
)

type App struct {
}

func (a *App) Run() {
	a.initConfig()

}

func NewApp() *App {
	return &App{}
}

func (a *App) initConfig() {
	var configPath string
	flag.StringVar(&configPath, "c", "./conf/config.yaml", "local config path")
	v := config.NewViper(configPath)
	var localconfig entity.LocalBase
	err := v.GetConfig(context.Background(), &localconfig)
	if err != nil {
		log.Logger.Panic(err)
	}

}
