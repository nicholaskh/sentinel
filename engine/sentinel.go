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
		notificationPlugins := make([]NotificationPlugin, 0)
		for _, notificationCmd := range service.NotificationCmds {
			notificationPluginCreator, exists := availablePlugins[notificationCmd.Cmd]
			if !exists {
				panic(fmt.Sprintf("Unknown notification command[%s] in target[%s]", notificationCmd.Cmd, service.Target))
			}
			notificationPlugin, ok := notificationPluginCreator().(NotificationPlugin)
			if !ok {
				panic(fmt.Sprintf("Invalid cmd plugin: %s", notificationCmd.Cmd))
			}
			notificationPlugin.Init(notificationCmd, service)
			go notificationPlugin.Start()
			notificationPlugins = append(notificationPlugins, notificationPlugin)
		}
		pluginCreator, exists := availablePlugins[service.Command]
		if !exists {
			panic(fmt.Sprintf("Unknown command[%s] in target[%s]", service.Command, service.Target))
		}
		cmdPlugin, ok := pluginCreator().(CmdPlugin)
		if !ok {
			panic(fmt.Sprintf("Invalid cmd plugin: %s", service.Command))
		}
		cmdPlugin.Init(service)
		cmdPlugin.SetNotificationPlugins(notificationPlugins)
		go cmdPlugin.Start()
	}

	<-make(chan interface{})
}
