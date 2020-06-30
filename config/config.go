package config

func GetConfig(confInterface ConfInterface) (string, error) {
	return confInterface.GetConfig()
}

func ListenConfig(confInterface ConfInterface, f func(data string)) error {
	return confInterface.ListenConfig(f)
}

func DeleteConfig(confInterface ConfInterface) (bool, error) {
	return confInterface.DeleteConfig()
}


func PublishConfig(confInterface ConfInterface,conf interface{}) (bool, error) {
	return confInterface.PublishConfig(conf)
}