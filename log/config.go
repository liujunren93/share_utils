package log

import (
	"time"

	"github.com/sirupsen/logrus"
)

var levelMap = map[string]logrus.Level{
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
	"trace": logrus.TraceLevel,
}

type Out string

const (
	OUT_FILE    Out = "file"
	OUT_CONSOLE Out = "console"
)

type Config struct {
	Out             Out     `json:"out" yaml:"out"`
	SetReportCaller bool    `json:"set_report_caller" yaml:"set_report_caller"` //是否打印文件
	Level           string  `json:"level" yaml:"level"`                         //required
	Rotate          *Rotate `json:"rotate" yaml:"rotate"`
	Remote          *Remote `yaml:"remote" json:"remote"` // 输出到远程
}

type Rotate struct {
	LogFile      string        `json:"log_file" yaml:"log_file"`
	MaxAge       time.Duration `json:"max_age" yaml:"max_age"`   // 保留时间
	RotationTime time.Duration `json:"rotation" yaml:"rotation"` //新文件间隔

}

type Remote struct {
	Enable bool   `yaml:"enable"` // 是否启用远程
	Host   string `yaml:"host"`
}

var defaultConfig = Config{
	Out:             OUT_CONSOLE,
	SetReportCaller: true,
	Level:           "debug",
	Rotate: &Rotate{
		LogFile:      "./log/log.json",
		MaxAge:       86400 * 30,
		RotationTime: 86400,
	},
}
