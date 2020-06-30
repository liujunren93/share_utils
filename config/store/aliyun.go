package store

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type aliyunConf struct {
	accessKey string
	secretKey string
	endpoint  string
	dataId    string
	group     string
	content   string
}

var (
	client config_client.IConfigClient
)

func NewStore(accessKey, secretKey, namespaceId, endpoint, dataId, group string) (aliyunConf, error) {
	aliyun := aliyunConf{
		dataId: dataId,
		group:  group,
	}
	err := initClient(accessKey, secretKey, namespaceId, endpoint)
	return aliyun, err
}

func initClient(accessKey, secretKey, namespaceId, endpoint string) error {
	clientConfig := constant.ClientConfig{
		Endpoint:       endpoint + ":8080",
		NamespaceId:    namespaceId,
		AccessKey:      accessKey,
		SecretKey:      secretKey,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
	}
	var err error
	client, err = clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})
	return err
}

//
func (a aliyunConf) PublishConfig(content interface{}) (bool, error) {
	return client.PublishConfig(vo.ConfigParam{
		DataId:  a.dataId,
		Group:   a.group,
		Content: content.(string),
	})
}

//
func (a aliyunConf) GetConfig() (string, error) {

	return client.GetConfig(vo.ConfigParam{
		DataId: a.dataId,
		Group:  a.group,
	})

}

//
func (a aliyunConf) ListenConfig(f func(string)) error {
	err := client.ListenConfig(vo.ConfigParam{
		DataId: a.dataId,
		Group:  a.group,
		OnChange: func(namespace, group, dataId, data string) {
			f(data)
		},
	})
	return err
}

//
func (a aliyunConf) DeleteConfig() (bool, error) {
	return client.DeleteConfig(vo.ConfigParam{
		DataId: a.dataId,
		Group:  a.group,
	})
}
