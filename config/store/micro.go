package store

import (
	"github.com/micro/go-micro/v3/config"
	"github.com/micro/go-micro/v3/config/source"
	config2 "github.com/liujunren93/share_utils/config"
)

type microConf struct {
	conf config.Config
}

func NewMicroStore(source source.Source) (config2.ConfI, error) {
	newConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	newConfig.Load(source)
	return &microConf{
		conf: newConfig,
	}, nil
}

func (e *microConf) PublishConfig(options *config2.DataOptions) (bool, error) {
	panic("implement me")
}

func (e *microConf) GetConfig(options *config2.DataOptions) (interface{}, error) {
	get := e.conf.Get()
	return get.Bytes(), nil
}

func (e *microConf) ListenConfig(options *config2.DataOptions,f func(interface{})) {
	watch, _ := e.conf.Watch(options.Path)
	for {
		next, err := watch.Next()
		if err == nil {
			f(next.Bytes())
		}
	}

}

func (e *microConf) DeleteConfig(options *config2.DataOptions) (bool, error) {
	panic("implement me")
}
