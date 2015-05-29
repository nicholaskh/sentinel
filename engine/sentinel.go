package engine

import (
	log "github.com/nicholaskh/log4go"
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
			log.Error("Unknown command[%s] found in service[%s]", service.Command, service.Name)
			continue
		}
		plugin := pluginCreator()
		plugin.Init(service)
		plugin.Start()
	}
}
