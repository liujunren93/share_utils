package log

import (
	"context"

	"github.com/sirupsen/logrus"
)

type shareLog struct {
	core *logrus.Logger
}

func New() *shareLog {
	return &shareLog{core: logrus.New()}
}

func (l *shareLog) GetLogrus() *logrus.Logger {
	return l.core
}

func (s *shareLog) log(args ...interface{}) (*logrus.Entry, []interface{}) {

	var ctx context.Context
	var ctxIndex = -1
	stringIndex := -1
	stringLen := 0
	for i, v := range args {
		if str, ok := v.(string); ok {
			if str[len(str)-1:] != ":" && stringIndex == -1 {
				stringIndex = i
			}
			if stringLen > 1 {
				break
			}
			stringLen++
		}
	}
	if stringIndex != -1 && stringLen > 1 {
		args[stringIndex] = args[stringIndex].(string) + ":"
	}

	for i, v := range args {
		if c, ok := v.(context.Context); ok {
			ctx = c
			ctxIndex = i
		}
	}
	if ctxIndex >= 0 {
		args = append(args[:ctxIndex], args[ctxIndex+1:]...)
	} else {
		ctx = context.TODO()
	}

	entry := s.core.WithContext(ctx)
	return entry, args
}

func (s *shareLog) logf(args ...interface{}) (entry *logrus.Entry, newArgs []interface{}) {
	var ctx context.Context
	var ctxIndex = -1
	for i, v := range args {
		if c, ok := v.(context.Context); ok {
			ctx = c
			ctxIndex = i
		}
	}
	if ctxIndex >= 0 {
		newArgs = args[ctxIndex : ctxIndex+1]
	} else {
		ctx = context.TODO()
	}

	entry = s.core.WithContext(ctx)
	return

}

func (s *shareLog) Trace(args ...interface{}) {

	entry, args := s.log(args...)
	entry.Trace(args...)
}

func (s *shareLog) Debug(args ...interface{}) {
	entry, args := s.log(args...)
	entry.Debug(args...)
}

func (s *shareLog) Debugf(format string, args ...interface{}) {
	entry, args := s.logf(args...)
	entry.Debugf(format, args...)
}

func (s *shareLog) Info(args ...interface{}) {
	entry, args := s.log(args...)
	entry.Info(args...)
}

func (s *shareLog) Infof(format string, args ...interface{}) {
	entry, args := s.logf(args...)
	entry.Infof(format, args...)
}

func (s *shareLog) Warn(args ...interface{}) {
	entry, args := s.log(args...)
	entry.Warn(args...)
}

func (s *shareLog) Warning(args ...interface{}) {
	entry, args := s.log(args...)
	entry.Warning(args...)
}

func (s *shareLog) Error(args ...interface{}) {
	entry, args := s.log(args...)
	entry.Error(args...)
}

func (s *shareLog) Fatal(args ...interface{}) {
	entry, args := s.log(args...)
	entry.Fatal(args...)
}

func (s *shareLog) Panic(args ...interface{}) {
	entry, args := s.log(args...)
	entry.Panic(args...)
}
