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
	content   string
}

var (
	client config_client.IConfigClient
)

func NewAliyunStore(accessKey, secretKey, namespaceId, endpoint string) (aliyunConf, error) {
	aliyun := aliyunConf{}
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

//options  0:DataId,1:Group;2:Content
func (a aliyunConf) PublishConfig(options ...interface{}) (bool, error) {
	return client.PublishConfig(vo.ConfigParam{
		DataId:  options[0].(string),
		Group:   options[1].(string),
		Content: (options[2]).(string),
	})
}

//
func (a aliyunConf) GetConfig(options ...string) (interface{}, error) {
	return client.GetConfig(vo.ConfigParam{
		DataId: options[0],
		Group:  options[1],
	})

}

//
func (a aliyunConf) ListenConfig(f func(string), options ...string) error {
	err := client.ListenConfig(vo.ConfigParam{
		DataId: options[0],
		Group:  options[1],
		OnChange: func(namespace, group, dataId, data string) {
			f(data)
		},
	})
	return err
}

//
func (a aliyunConf) DeleteConfig(options ...string) (bool, error) {
	return client.DeleteConfig(vo.ConfigParam{
		DataId: options[0],
		Group:  options[1],
	})
}
