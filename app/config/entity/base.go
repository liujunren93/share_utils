package entity

import (
	"github.com/go-redis/redis/v8"
	"github.com/liujunren93/share_utils/databases/gorm"
	"github.com/liujunren93/share_utils/log"
	"github.com/mitchellh/mapstructure"
)

type ConfMap map[string]interface{}

var DefaultConfig = Config{
	Log: &log.Config{
		Debug:           false,
		SetReportCaller: true,
		Level:           "debug",
	},
}

type LocalBase struct {
	AppName    string       `mapstructure:"app_name"`
	NameSpace  string       `mapstructure:"namespace"`
	ConfCenter ConfigCenter `mapstructure:"conf_center"` //redis etcd
	RunMode    string       `mapstructure:"run_mode"`    //debug product
	HttpHost   string       `mapstructure:"http_host"`
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
	Log      *log.Config    `json:"log" yaml:"log"`
	Redis    *redis.Options `json:"redis" yaml:"redis"`
	Mysql    *gorm.Mysql    `json:"mysql" yaml:"mysql"`
	Registry *Registry      `json:"registry" yaml:"registry"`
}
