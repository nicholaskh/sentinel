package config

import (
	conf "github.com/nicholaskh/jsconf"
)

type NotificationConfig struct {
	Name      string
	Cmd       string
	Notifiers []string
}

func (this *NotificationConfig) LoadConfig(cf *conf.Conf) {
	this.Name = cf.String("name", "")
	this.Cmd = cf.String("cmd", "")
	this.Notifiers = cf.StringList("notifiers", nil)

	if this.Name == "" ||
		this.Cmd == "" ||
		this.Notifiers == nil {
		panic("Notification config incomplete")
	}
}
