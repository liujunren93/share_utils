package config

import (
	"errors"
	"github.com/ghodss/yaml"
	"reflect"
)

var TypeErr = errors.New("Type mismatch")

func GetConfig(confInterface ConfInterface, resData interface{}, options ...string) error {

	of := reflect.TypeOf(resData)
	if of.Kind() != reflect.Ptr {
		return errors.New("resData is not ptr")
	}

	config, err := confInterface.GetConfig(options...)

	if err != nil {
		return err
	}
	var NewConf []byte
	switch config.(type) {
	case []byte:
		NewConf = config.([]byte)
	case string:
		NewConf = []byte(config.(string))
	default:
		return TypeErr
	}
	return yaml.Unmarshal(NewConf, resData)

}

func ListenConfig(confInterface ConfInterface, f func(interface{}), options ...string) {

	confInterface.ListenConfig(f, options...)

}

func DeleteConfig(confInterface ConfInterface, options ...string) (bool, error) {
	return confInterface.DeleteConfig(options...)
}

func PublishConfig(confInterface ConfInterface, options ...interface{}) (bool, error) {
	return confInterface.PublishConfig(options...)
}
