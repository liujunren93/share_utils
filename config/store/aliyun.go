package store

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type AcmConf struct {
	client config_client.IConfigClient
}

func NewAcmStore(accessKey, secretKey, namespaceId, endpoint, logDir, cacheDir string) (*AcmConf, error) {
	clientConfig := constant.ClientConfig{
		Endpoint:       endpoint,
		NamespaceId:    namespaceId,
		AccessKey:      accessKey,
		SecretKey:      secretKey,
		TimeoutMs:      500 * 1000,
		ListenInterval: 60 * 1000,
	}
	if  logDir==""{
		clientConfig.LogDir=logDir
	}
	if  cacheDir==""{
		clientConfig.CacheDir=cacheDir
	}
	// Initialize client.
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})

	return &AcmConf{client: configClient}, err
}

//options  0:DataId,1:Group;2:Content
func (a *AcmConf) PublishConfig(options ...interface{}) (bool, error) {

	return a.client.PublishConfig(vo.ConfigParam{
		DataId:  options[0].(string),
		Group:   options[1].(string),
		Content: (options[2]).(string),
	})
}

//options  0:DataId,1:Group;
func (a *AcmConf) GetConfig(options ...string) (interface{}, error) {

	// Get plain content from ACM.
	return a.client.GetConfig(vo.ConfigParam{
		DataId: options[0],
		Group:  options[1],
	},
	)

}

//options  0:DataId,1:Group;
func (a *AcmConf) ListenConfig(f func(interface{}), options ...string) {
	a.client.ListenConfig(vo.ConfigParam{
		DataId: options[0],
		Group:  options[1],
		OnChange: func(namespace, group, dataId, data string) {
			f(data)
		},
	})

}

//
func (a *AcmConf) DeleteConfig(options ...string) (bool, error) {
	return a.client.DeleteConfig(vo.ConfigParam{
		DataId: options[0],
		Group:  options[1],
	})
}
