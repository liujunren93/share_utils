package config

import (
	"errors"
	"github.com/ghodss/yaml"
	"reflect"
)

var (

	TypeErr = errors.New("Type mismatch")
)



func GetConfig(confInterface ConfI, resData interface{}, options DataOptions) error {

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
		default:
			return TypeErr
	}
	return yaml.Unmarshal(NewConf, resData)

}

func ListenConfig(confInterface ConfI, f func(interface{}), options DataOptions) {

	confInterface.ListenConfig( options,f)

}

func DeleteConfig(confInterface ConfI, options DataOptions) (bool, error) {
	return confInterface.DeleteConfig(options)
}

func PublishConfig(confInterface ConfI, options DataOptions) (bool, error) {
	return confInterface.PublishConfig(options)
}
