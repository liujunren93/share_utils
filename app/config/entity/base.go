package entity

import (
	"github.com/liujunren93/share_utils/db/gorm"
	"github.com/liujunren93/share_utils/db/redis"
	"github.com/liujunren93/share_utils/log"
	"github.com/mitchellh/mapstructure"
)

type ConfMap map[string]interface{}

func (c ConfMap) ConfMap(dest interface{}) error {
	return mapstructure.Decode(c, dest)
}

type BaseConfiger interface {
	GetVersion() string
	GetLogConfig() (*log.Config, bool)
	GetRegistryConfig() (*Registry, bool)
	GetRouterCenter() *RouterCenterConf
}

var DefaultConfig = &Config{
	Version: "v0.0.1",
	Log: &log.Config{
		Debug:           false,
		SetReportCaller: true,
		Level:           "debug",
	},
	RouterCenterConf: nil,
}

type LocalBase struct {
	AppName         string       `mapstructure:"app_name"`
	Namespace       string       `mapstructure:"namespace"`
	ConfCenter      ConfigCenter `mapstructure:"conf_center"` //redis etcd
	RunMode         string       `mapstructure:"run_mode"`    //debug product
	HttpHost        string       `mapstructure:"http_host"`
	EnableAutoRoute bool         `json:"enable_auto_route" mapstructure:"enable_auto_route"` // gateway 生效
	ApiPrefix       string       `json:"api_prefix" mapstructure:"api_prefix"`
}

type ConfigCenter struct {
	Enable   bool    `mapstructure:"enable"`
	Type     string  `mapstructure:"type"`      // redis
	ConfName string  `mapstructure:"conf_name"` // 配置名
	Group    string  `mapstructure:"group"`     // debug product
	Config   ConfMap `mapstructure:"config"`
}

type Config struct {
	Version          string            `json:"version" yaml:"version"` //配置版本
	Log              *log.Config       `json:"log" yaml:"log"`
	Redis            *redis.Config     `json:"redis" yaml:"redis"`
	Mysql            *gorm.Mysql       `json:"mysql" yaml:"mysql"`
	Registry         *Registry         `json:"registry" yaml:"registry"`
	RouterCenterConf *RouterCenterConf `json:"router_center_conf" yaml:"router_center_conf"`
}

func (c *Config) GetVersion() string {
	return c.Version
}
func (c *Config) GetLogConfig() (*log.Config, bool) {
	if c.Log == nil {
		return nil, false
	}
	return c.Log, true
}
func (c *Config) GetRegistryConfig() (*Registry, bool) {
	if c.Registry == nil {
		return nil, false
	}
	return c.Registry, true
}
func (c *Config) GetRouterCenter() *RouterCenterConf {

	return c.RouterCenterConf
}

// 自动路由配置
type RouterCenterConf struct {
	Type      int8    ` json:"type" yaml:"type"`    // redis etcd
	Enable    bool    `json:"enable" yaml:"enable"` //
	RedisConf ConfMap `json:"redis" yaml:"redis"`
}
