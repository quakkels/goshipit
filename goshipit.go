package main

import (
	"github.com/aphistic/gomol"
	"github.com/quakkels/goshipit/images"
)

func main() {
	startUp()
	defer cleanUp()

	images, err := images.NewImages("config.json")
	if err != nil {
		gomol.Fatal(err.Error())
		panic(err)
	}

	gomol.Infof("images: %v", images)
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
