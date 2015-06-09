package config

import (
	"fmt"

	conf "github.com/nicholaskh/jsconf"
)

type SentinelConfig struct {
	Services      []*ServiceConfig
	Notifications []*NotificationConfig
}

func (this *SentinelConfig) LoadConfig(cf *conf.Conf) {
	this.Notifications = make([]*NotificationConfig, 0)
	for i, _ := range cf.List("notifications", []interface{}{}) {
		section, err := cf.Section(fmt.Sprintf("notifications[%d]", i))
		if err != nil {
			panic(err)
		}

		notification := new(NotificationConfig)
		notification.LoadConfig(section)
		this.Notifications = append(this.Notifications, notification)
	}

	this.Services = make([]*ServiceConfig, 0)
	for i, _ := range cf.List("services", nil) {
		section, err := cf.Section(fmt.Sprintf("services[%d]", i))
		if err != nil {
			panic(err)
		}

		service := new(ServiceConfig)
		service.LoadConfig(section, this.Notifications)
		this.Services = append(this.Services, service)
	}
}
