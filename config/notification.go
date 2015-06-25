package config

import (
	conf "github.com/nicholaskh/jsconf"
)

type NotificationConfig struct {
	Name      string
	Cmd       string
	Server    string
	User      string
	Pwd       string
	Notifiers []string
}

func (this *NotificationConfig) LoadConfig(cf *conf.Conf) {
	this.Name = cf.String("name", "")
	this.Cmd = cf.String("cmd", "")
	this.Server = cf.String("server", "")
	this.User = cf.String("user", "")
	this.Pwd = cf.String("pwd", "")
	this.Notifiers = cf.StringList("notifiers", nil)

	if this.Name == "" ||
		this.Cmd == "" ||
		this.Server == "" ||
		this.User == "" ||
		this.Pwd == "" ||
		this.Notifiers == nil {
		panic("Notification config incomplete")
	}
}
