package config

type ConfInterface interface {
	PublishConfig(...interface{}) (bool, error)
	GetConfig(...string) (interface{}, error)
	ListenConfig(func(string),...string) error
	DeleteConfig(...string) (bool, error)
}

type OptionInterface interface {
	GetOption() string
}
