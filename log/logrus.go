package log

import (
	"fmt"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

var Logger *shareLog
var shaLogConfig string

func init() {
	Logger = New()
}
func Upgrade(conf *Config) {
	Logger.Debug("upgrade.log", conf)
	Init(conf)
}

//	func GetLogger(group string) *logrus.Logger {
//		fmt.Println("GetLogger", group)
//		Logger.SetFormatter(&JSONFormatter{group, jsonFormatter})
//		return Logger
//	}
func SetHook(hooks ...logrus.Hook) {
	if len(hooks) != 0 {
		for _, h := range hooks {
			Logger.core.AddHook(h)
		}
	}
}
func Init(conf *Config) {

	Logger.core.SetReportCaller(conf.SetReportCaller)

	Logger.core.SetLevel(levelMap[strings.ToLower(conf.Level)])

	// Logger.core.SetFormatter(&JSONFormatter{})
	Logger.core.SetFormatter(&logrus.TextFormatter{})
	// Logger.core.SetFormatter(NewShareFormatter(conf.SetReportCaller))
	fmt.Println(conf.Out)
	if conf.Out == OUT_FILE {
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
		Logger.core.SetOutput(rotatelog)

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
	prifix          string
	setReportCaller bool
}

func NewShareFormatter(setReportCaller bool) *shareFormatter {
	return &shareFormatter{setReportCaller: setReportCaller} //
}

func (s *shareFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	msg := fmt.Sprintf("[%s,%s,%s]", entry.Level, entry.Time.Format("2006-01-02 15:04:05"), entry.Message)
	if s.setReportCaller {
		msg = fmt.Sprintf("%s,file:%s,line:%d\n\r", msg, entry.Caller.File, entry.Caller.Line)
	}

	return []byte(msg), nil
}
