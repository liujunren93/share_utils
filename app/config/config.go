package config

import (
	"encoding/json"

	"github.com/liujunren93/share_utils/app/config/entity"
	"github.com/liujunren93/share_utils/common/config"
	"github.com/mitchellh/mapstructure"
)

var monitors []func()

func InitRegistryMonitor() chan func() {
	var monitorsCh = make(chan func())
	go func() {
		for f := range monitorsCh {
			monitors = append(monitors, f)
		}
	}()
	return monitorsCh

}

type AppConfigOption struct {
	LocalConf *entity.LocalBase
	Configer  config.Configer
	Conf      *entity.Config
}

func DescConfig(desc interface{}) config.Callback {
	return func(content interface{}) error {
		return decode(content, desc)
	}
}

func DescConfigAndCallbacks(desc interface{}) config.Callback {
	return func(content interface{}) error {
		err := decode(content, desc)
		if err != nil {
			return err
		}
		for _, callback := range monitors {
			callback()
		}
		return nil
	}
}

func decode(content, desc interface{}) error {
	switch v := content.(type) {
	case string:
		return json.Unmarshal([]byte(v), desc)
	case map[string]interface{}:
		return mapstructure.Decode(v, desc)
	case []byte:
		return json.Unmarshal(v, desc)
	}
	return nil

}
