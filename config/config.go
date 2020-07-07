package config

func GetConfig(confInterface ConfInterface, options ...string) (interface{}, error) {
	return confInterface.GetConfig(options...)
}

func ListenConfig(confInterface ConfInterface, f func(data string), options ...string) {
	go confInterface.ListenConfig(f, options...)

}

func DeleteConfig(confInterface ConfInterface, options ...string) (bool, error) {
	return confInterface.DeleteConfig(options...)
}

func PublishConfig(confInterface ConfInterface, options ...interface{}) (bool, error) {
	return confInterface.PublishConfig(options...)
}
