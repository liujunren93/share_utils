package store

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"
)

type etcdConf struct {
	conf config.Config
}

func NewEtcdStore(source source.Source) (*etcdConf, error) {
	newConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	newConfig.Load(source)
	return &etcdConf{
		conf: newConfig,
	}, nil
}

func (e *etcdConf) PublishConfig(...interface{}) (bool, error) {
	panic("implement me")
}

func (e *etcdConf) GetConfig(options ...string) (interface{}, error) {
	fmt.Println(options)
	get := e.conf.Get(options...)

	return get.Bytes(), nil
}

func (e *etcdConf) ListenConfig(f func(interface{}), options ...string) {
	watch, _ := e.conf.Watch(options...)
	for {
		next, err := watch.Next()
		if err == nil {
			f(next.Bytes())
		}
	}

}

func (e *etcdConf) DeleteConfig(...string) (bool, error) {
	panic("implement me")
}
