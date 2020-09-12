package store

import (
	"github.com/fsnotify/fsnotify"
	"github.com/liujunren93/share_utils/config"
	"github.com/spf13/viper"
)

type sViper struct {
	viper *viper.Viper
}

func NewViperStore(o *config.DataOptions) *sViper {

	var v sViper
	v.viper = viper.New()
	v.viper.AddConfigPath(o.Path)
	v.viper.SetConfigFile(o.FileName)
	v.viper.SetConfigType(o.FileType)
	return &v
}

func (v *sViper) PublishConfig(options *config.DataOptions) (bool, error) {

	panic("implement me")
}

func (v *sViper) GetConfig(o *config.DataOptions) (interface{}, error) {
	if o != nil {
		v.viper.SetConfigFile(o.FileName)
	}

	err := v.viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	conf := v.viper.AllSettings()
	return conf, nil
}

func (v *sViper) ListenConfig(o *config.DataOptions, f func(interface{})) {
	if o != nil && o.FileName != "" {
		v.viper.SetConfigFile(o.FileName)
	}
	v.viper.WatchConfig()
	v.viper.OnConfigChange(func(in fsnotify.Event) {

		f(v.viper.AllSettings())
	})
}

func (v *sViper) DeleteConfig(options *config.DataOptions) (bool, error) {
	panic("implement me")
}
