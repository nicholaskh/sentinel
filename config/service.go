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

func (this *ServiceConfig) LoadConfig(cf *conf.Conf, notifications []*NotificationConfig) {
	this.Name = cf.String("name", "")

	this.Command = cf.String("cmd", "")

	this.Target = cf.String("target", "")
	this.Interval = cf.Duration("interval", time.Minute)
	this.Retry = cf.Int("retry", 5)
	this.RetryInterval = cf.Duration("retry_interval", time.Second*1)
	this.ReadTimeout = cf.Duration("read_timeout", time.Second*2)

	this.NotificationCmds = make([]*NotificationConfig, 0)
	for _, cmd := range cf.StringList("notification_cmds", []string{}) {
		found := false
		for _, not := range notifications {
			if cmd == not.Name {
				this.NotificationCmds = append(this.NotificationCmds, not)
				found = true
				break
			}
		}
		if !found {
			panic(fmt.Sprintf("Notification cmd name not valid: %s", cmd))
		}
	}

	if this.Name == "" ||
		this.Command == "" ||
		this.Target == "" ||
		len(this.NotificationCmds) == 0 {
		panic("Service config incomplete")
	}
}
