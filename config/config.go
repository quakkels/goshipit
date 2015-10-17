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

var configs []Configuration

func Load(path string) ([]Configuration, error) {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		gomol.Err(err.Error())
		return nil, err
	}

	if err := json.Unmarshal(configFile, &configs); err != nil {
		gomol.Err(err.Error())
		return nil, err
	}

	return configs, nil
}
