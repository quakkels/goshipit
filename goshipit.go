package main

import (
	"github.com/aphistic/gomol"
	"github.com/quakkels/goshipit/config"
)

var configs config.Images

func main() {
	startUp()
	defer cleanUp()

	configs, err := config.NewImages("config.json")
	if err != nil {
		gomol.Fatal(err.Error())
		panic(err)
	}

	gomol.Infof("Configs: %v", configs)
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
