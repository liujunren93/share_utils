package config

type ConfInterface interface {
	PublishConfig(interface{}) (bool, error)
	GetConfig() (string, error)
	ListenConfig(func(string)) error
	DeleteConfig() (bool, error)
}
