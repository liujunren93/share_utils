package config

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/shareChina/utils/config/store"
	"testing"
)

func init2() {
	var endpoint = "acm.aliyun.com"
	var namespaceId = "da1af185-ef8b-4fd5-ab5a-1c15fc3fe906"
	var accessKey = "LTAI4G58UvgfoChctGeeiZTS"
	var secretKey = "S62Ik50uEKbjLERA3fGa0ZIKVoJWf9"

	clientConfig := constant.ClientConfig{
		Endpoint:       endpoint + ":8080",
		NamespaceId:    namespaceId,
		AccessKey:      accessKey,
		SecretKey:      secretKey,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
	}

	// Initialize client.
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	var dataId = "mysql"
	var group = "test"

	// Get plain content from ACM.
	configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})
	configClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {

		},
	})

}

type Mysql struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

//
var MysqlConf *Mysql

func TestGetConfig(t *testing.T) {
	var endpoint = "acm.aliyun.com"
	var namespaceId = "da1af185-ef8b-4fd5-ab5a-1c15fc3fe906"
	var accessKey = "LTAI4G58UvgfoChctGeeiZTS"
	var secretKey = "S62Ik50uEKbjLERA3fGa0ZIKVoJWf9"
	aliyunStore, err := store.NewAliyunStore(accessKey, secretKey, namespaceId, endpoint)
	fmt.Println(err)
	var dataId = "mysql"
	var group = "test"
	//err = GetConfig(aliyunStore, &MysqlConf, dataId, group)
	//fmt.Println(err)
	err = ListenConfig(aliyunStore, func(i interface{}) {

	}, dataId, group)
	fmt.Println(err)
	select {}

}
