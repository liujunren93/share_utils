package config

import (
	"context"
	"encoding/json"
)

type Configer interface {
	PublishConfig(context.Context, *DataOptions) (bool, error)
	GetConfig(context.Context, *DataOptions) (interface{}, error)
	ListenConfig(context.Context, *DataOptions, func(interface{}))
	DeleteConfig(context.Context, *DataOptions) (bool, error)
}

type DataOptions struct {
	ConfigName string `json:"config_name"`
	Group      string `json:"group"` //debug product
	Content    string `json:"content"`
	Path       string `json:"path"`
	FileType   string `json:"file_type"`
}

func (opt DataOptions) String() string {
	marshal, _ := json.Marshal(&opt)
	return string(marshal)
}
