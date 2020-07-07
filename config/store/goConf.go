package store

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"
)

type goConf struct {
	conf config.Config
}

func (g goConf) PublishConfig(conf ...interface{}) (bool, error) {

	return true, nil
}

func (g goConf) GetConfig(options ...string) (interface{}, error) {
	get := g.conf.Get(options...)
	return get.Bytes(), nil
	//return g.Conf.Map(), nil
}

func (g goConf) ListenConfig(f func(string), options ...string) {
	watch, _ := g.conf.Watch(options...)
	for {
		next, err := watch.Next()
		if err == nil {
			f(string(next.Bytes()))
		}
	}

}

func (g goConf) DeleteConfig(options ...string) (bool, error) {
	panic("implement me")
}

func NewGoConfStore(s source.Source) (*goConf, error) {

	newConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	newConfig.Load(s)
	return &goConf{
		conf: newConfig,
	}, nil
}
