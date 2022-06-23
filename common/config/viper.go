package config

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Viper struct {
	viper *viper.Viper
}

func NewViper(filePath string) *Viper {
	path, fileType, configName := configType(filePath)
	var v Viper
	v.viper = viper.New()
	v.viper.AddConfigPath(path)
	v.viper.SetConfigType(fileType)
	v.viper.SetConfigName(configName)
	return &v
}
func configType(configPath string) (path, fileType, configName string) {
	fileExt := strings.ToLower(filepath.Ext(configPath))
	configName = filepath.Base(configPath)
	path = filepath.Dir(configPath)

	if fileExt == ".yml" || fileExt == ".yaml" {
		fileType = "yaml"
	} else {
		fileType = fileExt[1:]
	}
	return

}

func (v *Viper) PublishConfig(ctx context.Context) (bool, error) {
	panic("implement me")
}

func (v *Viper) GetConfig(ctx context.Context, dest interface{}) error {
	//v.viper.AddConfigPath("config")
	err := v.viper.ReadInConfig()
	if err != nil {
		return err
	}
	return v.viper.Unmarshal(&dest)
}

func (v *Viper) ListenConfig(ctx context.Context, f func(interface{})) {
	v.viper.WatchConfig()
	v.viper.OnConfigChange(func(in fsnotify.Event) {
		f(v.viper.AllSettings())
	})
}

func (v *Viper) DeleteConfig(ctx context.Context) (bool, error) {
	panic("implement me")
}
