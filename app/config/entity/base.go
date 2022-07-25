package entity

import (
	"github.com/liujunren93/share_utils/databases/gorm"
	"github.com/liujunren93/share_utils/databases/redis"
	"github.com/liujunren93/share_utils/log"
	"github.com/mitchellh/mapstructure"
)

type ConfMap map[string]interface{}

type DefaultConfiger interface {
	GetVersion() string
	GetLogConfig() (*log.Config, bool)
	GetRegistryConfig() (*Registry, bool)
}

var DefaultConfig = &Config{
	Version: "v0.0.1",
	Log: &log.Config{
		Debug:           false,
		SetReportCaller: true,
		Level:           "debug",
	},
}

type LocalBase struct {
	AppName         string       `mapstructure:"app_name"`
	Namespace       string       `mapstructure:"namespace"`
	ConfCenter      ConfigCenter `mapstructure:"conf_center"` //redis etcd
	RunMode         string       `mapstructure:"run_mode"`    //debug product
	HttpHost        string       `mapstructure:"http_host"`
	PluginPath      string       `json:"plugin_path" mapstructure:"plugin_path"`
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

func (c *ConfigCenter) ToConfig(dest interface{}) error {
	return mapstructure.Decode(c.Config, &dest)
}

type Config struct {
	Version  string        `json:"version" yaml:"version"` //配置版本
	Log      *log.Config   `json:"log" yaml:"log"`
	Redis    *redis.Config `json:"redis" yaml:"redis"`
	Mysql    *gorm.Mysql   `json:"mysql" yaml:"mysql"`
	Registry *Registry     `json:"registry" yaml:"registry"`
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
