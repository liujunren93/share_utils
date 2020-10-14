package config

import "encoding/json"

type ConfI interface {
	PublishConfig(*DataOptions) (bool, error)
	GetConfig(*DataOptions) (interface{}, error)
	ListenConfig(*DataOptions, func(interface{}))
	DeleteConfig(*DataOptions) (bool, error)
}

type DataOptions struct {
	DataId     string `json:"data_id"`
	Group      string `json:"group"`
	Content    string `json:"content"`
	Path       string `json:"path"`
	FileType   string `json:"file_type"`
	ConfigName string `json:"config_name"`
} 

func (opt DataOptions) String() string {
	marshal, _ := json.Marshal(&opt)
	return string(marshal)
}
