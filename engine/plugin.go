package engine

import (
	"github.com/nicholaskh/sentinel/config"
)

var (
	availablePlugins map[string]func() Plugin = make(map[string]func() Plugin)
)

type Plugin interface {
	Start()
}

type CmdPlugin interface {
	Plugin
	Init(*config.ServiceConfig)
	SetNotificationPlugins([]NotificationPlugin)
}

type NotificationPlugin interface {
	Plugin
	Init(*config.NotificationConfig, *config.ServiceConfig)
	Notify()
}

func RegisterPlugin(name string, factory func() Plugin) {
	availablePlugins[name] = factory
}
