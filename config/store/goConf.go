package store

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"
)

type goConf struct {
	Conf   config.Config
	Source source.Source
}

func (g goConf) PublishConfig(conf ...interface{}) (bool, error) {

	return true, nil
}

func (g goConf) GetConfig(options ...string) (interface{}, error) {
	get := g.Conf.Get(options...)
	return get.Bytes(), nil
	//return g.Conf.Map(), nil
}

func (g goConf) ListenConfig(f func(string), options ...string) error {
	watch, _ := g.Conf.Watch()
	for {
		next, err := watch.Next()
		if err == nil {
			f(string(next.Bytes()))
		}
	}
	return nil
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
		Conf:   newConfig,
		Source: s,
	}, nil
}
