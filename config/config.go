package config

import (
	"encoding/json"
	"github.com/aphistic/gomol"
	"io/ioutil"
)

type Configuration struct {
	Category string
	Path     string
}

func newConfiguration() Configuration {
	c := Configuration{
		Category: "category",
		Path:     "/path/to/somewhere",
	}
	return c
}

func Test() {
	gomol.Info("Testing file io")
	configFile, err := ioutil.ReadFile("config.json")

	if err != nil {
		gomol.Err(err.Error())
		panic(err)
	}

	gomol.Info(string(configFile))
}

func Test2() {
	gomol.Info("Testing string parsing to json")
	config := []byte(`{"Category":"category","Path":"/path"}`)
	var data Configuration

	if err := json.Unmarshal(config, &data); err != nil {
		gomol.Err(err.Error())
		panic(err)
	} else {
		gomol.Info("Got the config from the []byte.")
		gomol.Infof("here's: ", data)
	}
}
