package engine

import (
	"fmt"

	"github.com/nicholaskh/sentinel/config"
)

type Sentinel struct {
	config *config.SentinelConfig
}

func NewSentinel(config *config.SentinelConfig) *Sentinel {
	this := new(Sentinel)
	this.config = config

	return this
}

func (this *Sentinel) RunForever() {
	for _, service := range this.config.Services {
		pluginCreator, exists := availablePlugins[service.Command]
		if !exists {
			panic(fmt.Sprintf("Unknown command[%s] target[%s]", service.Command, service.Target))
		}
		plugin := pluginCreator()
		plugin.Init(service)
		plugin.Start()
	}
}
