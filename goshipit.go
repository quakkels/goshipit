package main

import (
	"github.com/aphistic/gomol"
	"github.com/quakkels/goshipit/controllers"
	"github.com/quakkels/goshipit/images"
	"net/http"
)

func main() {
	startUp()
	defer cleanUp()

	err := images.InitImages()
	if err != nil {
		gomol.Fatal("Could not load images.json.")
		panic(err)
	}

	controllers.Register()
	http.ListenAndServe(":8080", nil)
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
