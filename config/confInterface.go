package config

type ConfI interface {
	PublishConfig(*DataOptions) (bool, error)
	GetConfig(*DataOptions) (interface{}, error)
	ListenConfig(*DataOptions, func(interface{}))
	DeleteConfig(*DataOptions) (bool, error)
}

type DataOptions struct {
	DataId     string
	Group      string
	Content    string
	Path       string
	FileType   string
	ConfigName string
}
