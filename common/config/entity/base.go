package entity

import "github.com/mitchellh/mapstructure"

type ConfMap map[string]interface{}

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
	Log   *Log   `json:"log" yml:"log"`
	Redis *Redis `json:"redis" yml:"redis"`
	Mysql *Mysql `json:"mysql" yml:"mysql"`
}

type Log struct {
	Level  string `yaml:"level"`
	Remote struct {
		Enable bool   `yaml:"enable"` // 是否启用远程
		Host   string `yaml:"host"`
	} `yaml:"remote"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"database"`
}

type Redis struct {
	Enable       bool   `yaml:"enable" json:"enable"`
	Network      string `yaml:"network" json:"network"`
	Addr         string `yaml:"addr" json:"addr"`
	User         string `yaml:"user" json:"user"`
	Password     string `yaml:"password" json:"password"`
	DB           int    `yaml:"db" json:"db"`
	MinIdleConns int    `yaml:"min_idle_conns" json:"min_idle_conns"`
}
