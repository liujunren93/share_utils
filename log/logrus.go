package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
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
		Logger.SetFormatter(shareFormatter{})
	}
}

type shareFormatter struct {
	TimestampFormat string
}

func (s shareFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if s.TimestampFormat == "" {
		s.TimestampFormat="2006-01-02 15:04:05"
	}
	fileName := path.Base( entry.Caller.File)
	sprintf := fmt.Sprintf("[%s][time:%s,file:%s,func:%s,line:%d]:%v\n\r",entry.Level, entry.Time.Format(s.TimestampFormat),fileName, entry.Caller.Function, entry.Caller.Line, entry.Message)
	return []byte(sprintf),nil
}
