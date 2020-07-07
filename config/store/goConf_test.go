package store

import (
	"fmt"
	"github.com/micro/go-micro/v2/config/source/etcd"
	"testing"
)

func TestNewGoConf(t *testing.T) {
	source := etcd.NewSource()
	conf, err := NewGoConfStore(source)

	config, err := conf.GetConfig("micro", "config", "acm")
	conf.ListenConfig(func(s string) {
		fmt.Println(string(s))
	})

	fmt.Println(string(config.([]byte)), err)
}

