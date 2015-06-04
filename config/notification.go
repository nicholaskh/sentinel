package config

import (
	conf "github.com/nicholaskh/jsconf"
)

type NotificationConfig struct {
	Cmd       string
	Notifiers []string
}

func (this *NotificationConfig) LoadConfig(cf *conf.Conf) {
	this.Cmd = cf.String("cmd", "")
	this.Notifiers = cf.StringList("notifiers", nil)

	if this.Cmd == "" ||
		this.Notifiers == nil {
		panic("Notification config incomplete")
	}
}
