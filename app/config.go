package app

import (
	"context"
	"encoding/json"

	"github.com/liujunren93/share_utils/common/config"
	"github.com/liujunren93/share_utils/common/config/entity"
	"github.com/mitchellh/mapstructure"
)

type appConfigOption struct {
	LocalConf      entity.LocalBase
	Configer       config.Configer
	ctx            context.Context
	configMonitors []func()
}

func DescConfig(desc interface{}) config.Callback {
	return func(content interface{}) error {
		return decode(content, desc)
	}
}

func DescConfigAndCallbacks(desc interface{}, callbacks []func()) config.Callback {
	return func(content interface{}) error {
		err := decode(content, desc)
		if err != nil {
			return err
		}
		for _, callback := range callbacks {
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
