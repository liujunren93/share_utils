package store

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/liujunren93/share_utils/common/config"
	"github.com/spf13/viper"
)

type Viper struct {
	viper *viper.Viper
}

func NewViper(filePath string) (v *Viper, fileType, confname string) {
	path, fileType, configName := configType(filePath)
	v = new(Viper)
	v.viper = viper.New()
	v.viper.AddConfigPath(path)
	// v.viper.SetConfigType(fileType)
	// v.viper.SetConfigName(configName)
	return v, fileType, configName
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

func (v *Viper) PublishConfig(ctx context.Context, confName, group, content string) (bool, error) {
	panic("implement me")
}

func (v *Viper) GetConfig(ctx context.Context, fileType, fileName string, callback config.Callback) error {
	v.viper.SetConfigType(fileType)
	v.viper.SetConfigName(fileName)
	err := v.viper.ReadInConfig()
	if err != nil {
		return err
	}
	return callback(fileType, fileName, v.viper.AllSettings())
}

func (v *Viper) ListenConfig(ctx context.Context, fileType, fileName string, callback config.Callback) error {
	v.viper.WatchConfig()
	v.viper.OnConfigChange(func(in fsnotify.Event) {
		callback(fileType, fileName, v.viper.AllSettings())
	})
	return nil
}

func (v *Viper) DeleteConfig(ctx context.Context, confName, group string) (bool, error) {
	panic("implement me")
}
