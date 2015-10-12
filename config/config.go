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

var Configs []Configuration

func Load() {
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		gomol.Err(err.Error())
		panic(err)
	}

	if err := json.Unmarshal(configFile, &Configs); err != nil {
		gomol.Fatal(err.Error())
		panic(err)
	}
}
