package entity

type LocalBase struct {
	AppName      string `mapstructure:"app_name"`
	NameSpace    string `mapstructure:"namespace"`
	ConfigCenter string `mapstructure:"config_center"` //redis etcd
	RunMode      string `mapstructure:"run_mode"`      //debug product
	HttpHost     string `mapstructure:"http_host"`
}

type ConfigCenter struct {
	Enable bool   `mapstructure:"enable"`
	Type   string `mapstructure:"type"` // redis
	Redis  *Redis
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
	Enable   bool   `yaml:"enable"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
