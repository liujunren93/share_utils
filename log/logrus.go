package log

import (
	"fmt"
	"path"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

var Logger *logrus.Logger
var shaLogConfig string

func init() {
	Logger = logrus.New()
}
func Upgrade(conf *Config) {
	Logger.Debug("upgrade.log", conf)
	Init(conf)
}

// func GetLogger(group string) *logrus.Logger {
// 	fmt.Println("GetLogger", group)
// 	Logger.SetFormatter(&JSONFormatter{group, jsonFormatter})
// 	return Logger
// }

func Init(conf *Config) {

	Logger.SetReportCaller(conf.SetReportCaller)
	Logger.SetLevel(levelMap[strings.ToLower(conf.Level)])
	Logger.AddHook(new(TestHook))
	Logger.SetFormatter(&JSONFormatter{"", &logrus.JSONFormatter{}})
	if !conf.Debug && conf.Rotate != nil {
		rotatelog, err := rotatelogs.New(
			conf.Rotate.LogFile+".%Y%m%d%H%M",
			rotatelogs.WithLocation(time.Local),
			rotatelogs.WithLinkName(conf.Rotate.LogFile),
			rotatelogs.WithMaxAge(conf.Rotate.MaxAge*time.Second),
			rotatelogs.WithRotationTime(conf.Rotate.RotationTime*time.Second),
		)
		if err != nil {
			panic("init logrus rotatelogs err" + err.Error())
		}
		Logger.SetOutput(rotatelog)

	}

}

type TestHook struct {
	Fired bool
}

func (hook *TestHook) Fire(entry *logrus.Entry) error {

	hook.Fired = true
	return nil
}

func (hook *TestHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
	}
}

type shareFormatter struct {
	prifix string
}

func (s shareFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	fileName := path.Base(entry.Caller.File)

	sprintf := fmt.Sprintf("[%s][%s|%s:%d]:%v\n\r", entry.Level, entry.Time.Format("2006-01-02 15:04:05"), fileName, entry.Caller.Line, entry.Message)
	return []byte(sprintf), nil
}
