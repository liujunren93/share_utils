package config

type ConfInterface interface {
	PublishConfig(...interface{}) (bool, error)
	GetConfig(...string) (interface{}, error)
	ListenConfig(func(interface{}), ...string)
	DeleteConfig(...string) (bool, error)
}
