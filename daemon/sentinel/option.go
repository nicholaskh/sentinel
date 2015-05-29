package main

import (
	"flag"
)

var (
	options struct {
		configFile   string
		showVersion  bool
		logFile      string
		logLevel     string
		kill         bool
		lockFile     string
		crashLogFile string
	}
)

func parseFlags() {
	flag.BoolVar(&options.kill, "k", false, "kill sentinel")
	flag.StringVar(&options.lockFile, "lockfile", "sentinel.lock", "lock file")
	flag.StringVar(&options.configFile, "conf", "etc/sentinel.cf", "config file")
	flag.StringVar(&options.logFile, "log", "stdout", "log file")
	flag.StringVar(&options.logLevel, "level", "debug", "log level")
	flag.StringVar(&options.crashLogFile, "crashlog", "panic.dump", "crash log file")
	flag.BoolVar(&options.showVersion, "v", false, "show version and exit")

	flag.Parse()
}
