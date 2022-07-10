package config

import (
	"encoding/json"

	"github.com/liujunren93/share_utils/log"

	"github.com/liujunren93/share_utils/common/config"
	"github.com/mitchellh/mapstructure"
)

var monitorMap = make(map[string][]func())

type Monitor struct {
	ConfName string
	Callback func()
}

func NewMonitor(confName, group string, callback func()) *Monitor {
	return &Monitor{confName + group, callback}
}

func InitRegistryMonitor() chan *Monitor {
	var monitorsCh = make(chan *Monitor)
	go func() {
		for mo := range monitorsCh {
			log.Logger.Debug("InitRegistryMonitor", mo.ConfName)
			monitorMap[mo.ConfName] = append(monitorMap[mo.ConfName], mo.Callback)
		}
	}()
	return monitorsCh

}

func DescConfig(desc interface{}) config.Callback {
	return func(confName, group string, content interface{}) error {
		return decode(content, desc)
	}
}

func DescConfigAndCallbacks(desc interface{}) config.Callback {
	return func(confName, group string, content interface{}) error {
		err := decode(content, desc)
		if err != nil {
			return err
		}
		for _, callback := range monitorMap[confName+group] {
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
