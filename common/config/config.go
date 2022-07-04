package config

import (
	"context"
	"errors"
)

var (
	TypeErr = errors.New("type miss match")
)

type Callback func(confName, group string, configContent interface{}) error

type Configer interface {
	PublishConfig(ctx context.Context, confName, group, content string) (bool, error)
	GetConfig(ctx context.Context, confName, group string, callback Callback) error
	ListenConfig(ctx context.Context, confName, group string, callback Callback) error
	DeleteConfig(ctx context.Context, confName, group string) (bool, error)
}
