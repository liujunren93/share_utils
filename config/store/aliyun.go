package store

import (
	"github.com/liujunren93/share_utils/config"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type AcmOptions struct {
	AccessKey   string `json:"access_key"`
	SecretKey   string `json:"secret_key"`
	Endpoint    string `json:"endpoint"`
	NamespaceID string `json:"namespace_id"`
	LogDir      string `json:"log_dir"`
	CacheDir    string `json:"cache_dir"`
}

type acmConf struct {
	client config_client.IConfigClient
}

func NewAcmStore(option *AcmOptions) (config.Configer, error) {
	clientConfig := constant.ClientConfig{
		Endpoint:       option.Endpoint,
		NamespaceId:    option.NamespaceID,
		AccessKey:      option.AccessKey,
		SecretKey:      option.SecretKey,
		TimeoutMs:      500 * 1000,
		ListenInterval: 60 * 1000,
		LogDir:         option.LogDir,
		CacheDir:       option.CacheDir,
	}
	var conf acmConf
	// Initialize client.
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})
	conf.client = configClient
	return &conf, err
}

//options  0:DataId,1:Group;2:Content
func (a *acmConf) PublishConfig(options *config.DataOptions) (bool, error) {

	return a.client.PublishConfig(vo.ConfigParam{
		DataId:  options.DataId,
		Group:   options.Group,
		Content: options.Content,
	})
}

//options  0:DataId,1:Group;
func (a *acmConf) GetConfig(options *config.DataOptions) (interface{}, error) {

	// Get plain content from ACM.
	return a.client.GetConfig(vo.ConfigParam{
		DataId: options.DataId,
		Group:  options.Group,
	},
	)

}

//options  0:DataId,1:Group;
func (a *acmConf) ListenConfig(options *config.DataOptions, f func(interface{})) {
	a.client.ListenConfig(vo.ConfigParam{
		DataId: options.DataId,
		Group:  options.Group,
		OnChange: func(namespace, group, dataId, data string) {
			f(data)
		},
	})

}

//
func (a *acmConf) DeleteConfig(options *config.DataOptions) (bool, error) {
	return a.client.DeleteConfig(vo.ConfigParam{
		DataId: options.DataId,
		Group:  options.Group,
	})
}
