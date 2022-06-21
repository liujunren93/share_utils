package store

import (
	"context"

	"github.com/fsnotify/fsnotify"
	"github.com/liujunren93/share_utils/config"
	"github.com/spf13/viper"
)

type sViper struct {
	viper *viper.Viper
}

func NewViperStore(o *config.DataOptions) *config.Configer {

	var v sViper

	v.viper = viper.New()
	v.viper.AddConfigPath(o.Path)
	v.viper.SetConfigType(o.FileType)
	v.viper.SetConfigName(o.ConfigName)
	//v.viper.Debug()
	return &v
}

func (v *sViper) PublishConfig(ctx context.Context, options *config.DataOptions) (bool, error) {

	panic("implement me")
}

func (v *sViper) GetConfig(ctx context.Context, o *config.DataOptions) (interface{}, error) {
	if o != nil {
		v.viper.SetConfigFile(o.ConfigName)
	}
	//v.viper.AddConfigPath("config")
	err := v.viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	conf := v.viper.AllSettings()
	return conf, nil
}

func (v *sViper) ListenConfig(ctx context.Context, o *config.DataOptions, f func(interface{})) {
	if o != nil && o.ConfigName != "" {
		v.viper.SetConfigFile(o.ConfigName)
	}
	v.viper.WatchConfig()
	v.viper.OnConfigChange(func(in fsnotify.Event) {

		f(v.viper.AllSettings())
	})
}

func (v *sViper) DeleteConfig(ctx context.Context, options *config.DataOptions) (bool, error) {
	panic("implement me")
}
