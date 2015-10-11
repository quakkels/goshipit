package main

import (
	"github.com/aphistic/gomol"
	"github.com/quakkels/gojump/config"
)

func main() {
	startUp()
	defer cleanUp()

	config.Load()
	gomol.Infof("Configs: %v", config.Configs)
}

func startUp() {
	loggerConfig := gomol.NewConsoleLoggerConfig()
	consoleLogger, _ := gomol.NewConsoleLogger(loggerConfig)
	gomol.AddLogger(consoleLogger)
	gomol.InitLoggers()
	gomol.Info("Initialized console logger")
}
func cleanUp() {
	gomol.Info("Shutting down loggers")
	gomol.ShutdownLoggers()
}
