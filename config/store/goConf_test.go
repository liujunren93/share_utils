package store

import (
	"fmt"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/shareChina/utils/config"
	"testing"
	"time"
)

type app struct {
	ServerName    string `json:"service_name"`
	ServerVersion string `json:"server_version"`
	ServerAddress string `json:"server_address"`
	//RunMode         string `json:"run_mode"`
	RegistryAddress string `json:"registry_address"`
}

var AppConf *app

func TestNewGoConf(t *testing.T) {

	newSource := file.NewSource(
		file.WithPath("./init.yml"),
	)
	//AppConf.RunMode = *runMode
	microStore, err := NewMicroStore(newSource)
	config.ListenConfig(microStore, func(i interface{}) {
		fmt.Println(i)
	})

	fmt.Println(AppConf,err)
	time.Sleep(time.Hour)
}
