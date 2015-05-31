package config

import (
	"fmt"
	"time"

	conf "github.com/nicholaskh/jsconf"
	"github.com/nicholaskh/sentinel/global"
)

type ServiceConfig struct {
	Command string

	Target        string
	Interval      time.Duration
	Retry         int
	RetryInterval time.Duration
	ReadTimeout   time.Duration

	ListenPort int
}

func (this *ServiceConfig) LoadConfig(cf *conf.Conf) {
	this.Command = cf.String("cmd", "")

	this.Target = cf.String("target", "")
	this.Interval = cf.Duration("interval", time.Minute)
	this.Retry = cf.Int("retry", 5)
	this.RetryInterval = cf.Duration("retry_interval", time.Second*1)
	this.ReadTimeout = cf.Duration("read_timeout", time.Second*2)

	this.ListenPort = cf.Int("listen_port", global.AvailableUdpPort)
	if this.ListenPort == global.AvailableUdpPort {
		global.AvailableUdpPort++
	}
	if this.Command == "" || this.Target == "" {
		panic(fmt.Sprintf("service config imcomplete"))
	}
}