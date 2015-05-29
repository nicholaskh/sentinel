package config

import (
	"fmt"
	"time"

	conf "github.com/nicholaskh/jsconf"
	"github.com/nicholaskh/sentinel/global"
)

type ServiceConfig struct {
	Name    string
	Command string

	Target     string
	Interval   time.Duration
	ListenPort int
}

func (this *ServiceConfig) LoadConfig(cf *conf.Conf) {
	this.Name = cf.String("name", "")
	this.Command = cf.String("cmd", "")

	this.Target = cf.String("target", "")
	this.Interval = cf.Duration("interval", time.Minute)
	this.ListenPort = cf.Int("listen_port", global.AvailableUdpPort)
	if this.ListenPort == global.AvailableUdpPort {
		global.AvailableUdpPort++
	}
	if this.Name == "" || this.Command == "" || this.Target == "" {
		panic(fmt.Sprintf("service config imcomplete"))
	}
}
