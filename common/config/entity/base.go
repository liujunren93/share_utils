package entity

type LocalBase struct {
	AppName          string `yml:"app_name"`
	NameSpace        string `yaml:"namespace"`
	ConfigCenterType string `yaml:"config_center_type"` //redis etcd
	RunMode          string `yaml:"run_mode"`           //debug product
	HttpHost         string `yaml:"http_host"`
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
