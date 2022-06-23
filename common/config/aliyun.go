package config

import (
	"context"

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

type Acm struct {
	client config_client.IConfigClient
}

func NewAcmStore(option *AcmOptions) (*Acm, error) {
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
	var conf Acm
	// Initialize client.
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})
	conf.client = configClient
	return &conf, err
}

//options  0:configName,1:Group;2:Content
func (a *Acm) PublishConfig(ctx context.Context, configName, group, content string) (bool, error) {

	return a.client.PublishConfig(vo.ConfigParam{
		DataId:  configName,
		Group:   group,
		Content: content,
	})
}

//options  0:DataId,1:Group;
func (a *Acm) GetConfig(ctx context.Context, configName, group string) (interface{}, error) {

	// Get plain content from ACM.
	return a.client.GetConfig(vo.ConfigParam{
		DataId: configName,
		Group:  group,
	},
	)

}

//options  0:DataId,1:Group;
func (a *Acm) ListenConfig(ctx context.Context, configName, group string, f func(interface{})) {
	a.client.ListenConfig(vo.ConfigParam{
		DataId: configName,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			f(data)
		},
	})

}

//
func (a *Acm) DeleteConfig(ctx context.Context, configName, group string) (bool, error) {
	return a.client.DeleteConfig(vo.ConfigParam{
		DataId: configName,
		Group:  group,
	})
}
