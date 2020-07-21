package store

import (
	"fmt"
	"github.com/micro/go-micro/v2/config/source/etcd"
	"testing"
)

type authStr struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

func TestNewGoConf(t *testing.T) {

	source := etcd.NewSource(etcd.WithAddress())

	//var name  authStr
	store, _ := NewEtcdStore(source)

	getConfig, err2 := store.GetConfig("micro", "config", "acm", "auth")
	fmt.Println(getConfig, err2)
	//err := config.GetConfig(store, &name, "micro", "config", "acm", "auth")
	//fmt.Println(err,name)
}
