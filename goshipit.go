package main

import (
	"github.com/aphistic/gomol"
	"github.com/quakkels/goshipit/context"
	"github.com/quakkels/goshipit/controllers"
	"github.com/quakkels/goshipit/images"
	"github.com/quakkels/goshipit/slack"
	"net/http"
	"os"
	"strconv"
)

func main() {
	startUp()
	defer cleanUp()

	// get context from args
	args := os.Args[1:]
	if len(args) == 3 {
		isSecure, err := strconv.ParseBool(args[2])
		if err != nil {
			panic("isSecure cli argument is not boolean.")
		}

		context.SetContext(args[0], args[1], isSecure)
	}

	err := images.InitImages()
	if err != nil {
		gomol.Fatal("Could not load images.json.")
		panic(err)
	}

	err = slack.InitSlackConfig()
	if err != nil {
		gomol.Fatal("Could not load slack.json.")
		panic(err)
	}

	controllers.Register()
	http.ListenAndServe(":"+context.Port, nil)
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
