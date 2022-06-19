package log

import (
	"fmt"
	"path"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	if Logger == nil {
		Logger = logrus.New()
		Logger.SetReportCaller(true)
		//Logger.SetFormatter(&logrus.TextFormatter{
		//	FullTimestamp:   true,
		//	DisableColors:true,
		//	TimestampFormat:"2006-01-02 15:04:05",
		//	CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
		//		//处理文件名
		//		fileName := path.Base(frame.File)
		//		return frame.Function, fileName
		//	},
		//
		//})
		Logger.AddHook(new(TestHook))
		Logger.SetFormatter(&logrus.TextFormatter{})
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
