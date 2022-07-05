package log

import (
	"fmt"
	"path"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
}
func Upgrade(conf *Config) {
	Init(conf)
}

func Init(conf *Config) {
	Logger.SetReportCaller(conf.SetReportCaller)
	Logger.SetLevel(levelMap[strings.ToLower(conf.Level)])
	Logger.AddHook(new(TestHook))
	Logger.SetFormatter(&logrus.JSONFormatter{})
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
	TimestampFormat string
}

func (s shareFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if s.TimestampFormat == "" {
		s.TimestampFormat = "2006-01-02 15:04:05"
	}
	fileName := path.Base(entry.Caller.File)

	sprintf := fmt.Sprintf("[%s][%s|%s:%d]:%v\n\r", entry.Level, entry.Time.Format(s.TimestampFormat), fileName, entry.Caller.Line, entry.Message)
	return []byte(sprintf), nil
}
