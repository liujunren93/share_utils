package config

import (
	"errors"
	"github.com/ghodss/yaml"
	"reflect"
)

var (
	TypeErr = errors.New("type miss match")
)

func GetConfig(confInterface Configer, resData interface{}, options *DataOptions) error {

	of := reflect.TypeOf(resData)
	if of.Kind() != reflect.Ptr {
		return errors.New("resData is not ptr")
	}
	config, err := confInterface.GetConfig(options)
	if err != nil {
		return err
	}
	var NewConf []byte
	switch config.(type) {
	case []byte:
		NewConf = config.([]byte)
	case string:
		NewConf = []byte(config.(string))
	case map[string]interface{}:
		NewConf, err = yaml.Marshal(config)
		if err != nil {
			return err
		}

	default:
		return TypeErr
	}
	return yaml.Unmarshal(NewConf, resData)

}

func ListenConfig(confInterface Confter, f func(interface{}), options *DataOptions) {

	confInterface.ListenConfig(options, f)

}

func DeleteConfig(confInterface Confter, options *DataOptions) (bool, error) {
	return confInterface.DeleteConfig(options)
}

func PublishConfig(confInterface Confter, options *DataOptions) (bool, error) {
	return confInterface.PublishConfig(options)
}
