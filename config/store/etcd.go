package store

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
)

type etcdConf struct {
	conf config.Config
}

func NewEtcdStore() (*etcdConf, error) {

	if err != nil {
		return nil, err
	}
	return &etcdConf{
		conf: newConfig,
	}, nil
}

func (e *etcdConf) PublishConfig(...interface{}) (bool, error) {
	panic("implement me")
}

func (e *etcdConf) GetConfig(options ...string) (interface{}, error) {
	get := e.conf.Get(options...)
	fmt.Println(get)
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
