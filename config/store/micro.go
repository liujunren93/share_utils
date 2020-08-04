package store

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"
)

type MicroConf struct {
	conf config.Config
}

func NewMicroStore(source source.Source) (*MicroConf, error) {
	newConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	newConfig.Load(source)
	return &MicroConf{
		conf: newConfig,
	}, nil
}

func (e *MicroConf) PublishConfig(...interface{}) (bool, error) {
	panic("implement me")
}

func (e *MicroConf) GetConfig(options ...string) (interface{}, error) {
	get := e.conf.Get(options...)
	return get.Bytes(), nil
}

func (e *MicroConf) ListenConfig(f func(interface{}), options ...string) {
	watch, _ := e.conf.Watch(options...)
	for {
		next, err := watch.Next()
		if err == nil {
			f(next.Bytes())
		}
	}

}

func (e *MicroConf) DeleteConfig(...string) (bool, error) {
	panic("implement me")
}
