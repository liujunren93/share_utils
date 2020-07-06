package store

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"
)

type goConf struct {
	Conf   config.Config
	Source source.Source
}

func (g goConf) PublishConfig(conf... interface{}) (bool, error) {

	return true, nil
}

func (g goConf) GetConfig(options...string) (interface{}, error) {
	return g.Conf.Map(), nil
}

func (g goConf) ListenConfig(f func(string),options...string) error {
	watch, _ := g.Conf.Watch()
	for {
		next, err := watch.Next()
		if err == nil {
			f(string(next.Bytes()))
		}
	}
	return nil
}

func (g goConf) DeleteConfig(options...string) (bool, error) {
	panic("implement me")
}

func NewGoConf(s source.Source) (gf goConf, err error) {

	newConfig, err := config.NewConfig()
	if err != nil {
		return
	}

	newConfig.Load(s)
	gf.Conf = newConfig
	gf.Source = s
	return
}
