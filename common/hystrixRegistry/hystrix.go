package hystrixregistry

import "github.com/afex/hystrix-go/hystrix"

func Registry(menthod string) {
	if menthod == "" {
		return
	}
	cf := hystrix.CommandConfig{
		Timeout:                0,
		MaxConcurrentRequests:  0,
		RequestVolumeThreshold: 0,
		SleepWindow:            0,
		ErrorPercentThreshold:  0,
	}
	hystrix.ConfigureCommand(menthod, cf)

}
