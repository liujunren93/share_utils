package config

type ConfInterface interface {
	PublishConfig(...interface{}) (bool, error)
	GetConfig(...string) (interface{}, error)
	ListenConfig(func(string), ...string)
	DeleteConfig(...string) (bool, error)
}
