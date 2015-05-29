package engine

import (
	"github.com/nicholaskh/sentinel/config"
)

var (
	availablePlugins map[string]func() Plugin
)

type Plugin interface {
	Start()
	Init(*config.ServiceConfig)
}

func RegisterPlugin(name string, factory func() Plugin) {
	availablePlugins[name] = factory
}
