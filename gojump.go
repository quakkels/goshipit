package main

import (
	"errors"
	"github.com/aphistic/gomol"
	"github.com/quakkels/gojump/config"
)

func main() {
	startUp()
	defer cleanUp()

	config.Load()
	gomol.Infof("Configs: %v", config.Configs)

	path, err := getImgFromCategory("something", config.Configs)
	if err != nil {
		gomol.Errf(
			"Issue getting image from category: %v",
			err.Error())
	}

	gomol.Infof("Image from category: %v", path)
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

func getImgFromCategory(category string, config []config.Configuration) (string, error) {
	if config == nil {
		msg := "Config is not initialized."
		gomol.Err(msg)
		return "", errors.New(msg)
	}

	if len(config) == 0 {
		msg := "Nothing in configuration"
		gomol.Err(msg)
		return "", errors.New(msg)
	}

	imgs := make([]string, 0)
	imgs = append(imgs, "first")

	return imgs[0], nil
}
