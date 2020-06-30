package config

import (
	"fmt"
	"github.com/shareChina/utils/config/store"
	"testing"
	"time"
)

func TestGetConfig(t *testing.T) {
	newStore, _ := store.NewStore("LTAI4G27vXUGFDc8fWoFfLg9", "G8ECCFNM8qNCKFXsW6msSQUUIybI4b", "81d9055a-03a5-46a9-8b3c-557146dc7376", "acm.aliyun.com", "go.micro.service.account", "test")
	config, err := GetConfig(newStore)
	fmt.Println(	config, err)
	//ListenConfig(newStore, func(data string) {
	//	fmt.Println(data)
	//})
	time.Sleep(time.Second*3)
}