package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"syscall"

	"github.com/nicholaskh/golib/locking"
	"github.com/nicholaskh/golib/server"
	"github.com/nicholaskh/golib/signal"
	log "github.com/nicholaskh/log4go"
	"github.com/nicholaskh/sentinel/config"
	"github.com/nicholaskh/sentinel/engine"
)

var (
	SentinelConf *config.SentinelConfig
)

func init() {
	parseFlags()

	if options.showVersion {
		server.ShowVersionAndExit()
	}

	if options.kill {
		if err := server.KillProcess(options.lockFile); err != nil {
			fmt.Fprintf(os.Stderr, "stop failed: %s\n", err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	server.SetupLogging(options.logFile, options.logLevel, options.crashLogFile)

	if options.lockFile != "" {
		if locking.InstanceLocked(options.lockFile) {
			fmt.Fprintf(os.Stderr, "Another actor is running, exit...\n")
			os.Exit(1)
		}

		locking.LockInstance(options.lockFile)
	}

	signal.RegisterSignalHandler(syscall.SIGINT, func(sig os.Signal) {
		shutdown()
	})

	conf := server.LoadConfig(options.configFile)
	SentinelConf = new(config.SentinelConfig)
	SentinelConf.LoadConfig(conf)
}

func main() {
	defer func() {
		cleanup()

		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	sentinel := engine.NewSentinel(SentinelConf)
	sentinel.RunForever()
}

func shutdown() {
	cleanup()
	log.Info("Terminated")
	os.Exit(0)
}

func cleanup() {
	if options.lockFile != "" {
		locking.UnlockInstance(options.lockFile)
		log.Debug("Cleanup lock %s", options.lockFile)
	}
}
