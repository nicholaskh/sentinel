package config

import (
	"fmt"

	conf "github.com/nicholaskh/jsconf"
)

type SentinelConfig struct {
	Services []*ServiceConfig
}

func (this *SentinelConfig) LoadConfig(cf *conf.Conf) {
	this.Services = make([]*ServiceConfig, 0)
	for i, _ := range cf.List("services", nil) {
		section, err := cf.Section(fmt.Sprintf("services[%d]", i))
		if err != nil {
			panic(err)
		}

		service := new(ServiceConfig)
		service.LoadConfig(section)
		this.Services = append(this.Services, service)
	}
}
