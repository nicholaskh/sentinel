package config

import (
	"fmt"
	"time"

	conf "github.com/nicholaskh/jsconf"
)

type ServiceConfig struct {
	Name string

	Command string

	Target        string
	Interval      time.Duration
	Retry         int
	RetryInterval time.Duration
	ReadTimeout   time.Duration

	NotificationCmds []*NotificationConfig
}

func (this *ServiceConfig) LoadConfig(cf *conf.Conf) {
	this.Name = cf.String("name", "")

	this.Command = cf.String("cmd", "")

	this.Target = cf.String("target", "")
	this.Interval = cf.Duration("interval", time.Minute)
	this.Retry = cf.Int("retry", 5)
	this.RetryInterval = cf.Duration("retry_interval", time.Second*1)
	this.ReadTimeout = cf.Duration("read_timeout", time.Second*2)

	this.NotificationCmds = make([]*NotificationConfig, 0)
	for i, _ := range cf.List("notification_cmds", []interface{}{}) {
		notification := new(NotificationConfig)
		section, err := cf.Section(fmt.Sprintf("notification_cmds[%d]", i))
		if err != nil {
			panic(err)
		}
		notification.LoadConfig(section)
		this.NotificationCmds = append(this.NotificationCmds, notification)
	}

	if this.Name == "" ||
		this.Command == "" ||
		this.Target == "" ||
		this.NotificationCmds == nil {
		panic("Service config incomplete")
	}
}
