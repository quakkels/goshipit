package main

import (
	"github.com/aphistic/gomol"
	"github.com/quakkels/gojump/config"
)

func main() {
	gomol.AddLogger(gomol.NewConsoleLogger())
	gomol.InitLoggers()
	defer cleanUp()
	gomol.Info("Initialized console logger")

	config.Test()
}

func cleanUp() {
	gomol.Info("Shutting down loggers")
	gomol.ShutdownLoggers()
}
