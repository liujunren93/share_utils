package config

import (
	"encoding/json"
	"reflect"

	"github.com/liujunren93/share_utils/helper"
	"github.com/liujunren93/share_utils/log"

	"github.com/liujunren93/share_utils/common/config"
	"github.com/mitchellh/mapstructure"
)

var monitorMap = make(map[string][]CallbackObj)

type Monitor struct {
	ConfName string
	Callback CallbackObj
}

type CallbackObj struct {
	field    string // 字段名
	sha      string
	callback func()
}

func NewMonitor(confName, group, field string, callback func()) *Monitor {
	return &Monitor{confName + group, CallbackObj{
		field:    field,
		callback: callback,
	}}
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
		vf := reflect.ValueOf(desc).Elem()
		for k, c := range monitorMap[confName+group] {
			f := vf.FieldByName(c.field)
			sha, err := helper.Sha1Interface(f.Interface())
			if err != nil {
				log.Logger.Error("DescConfigAndCallbacks.Sha1Interface", err, vf.FieldByName(c.field).Interface())
				continue
			}
			if sha != c.sha {
				monitorMap[confName+group][k].sha = sha
				c.callback()
			}
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
