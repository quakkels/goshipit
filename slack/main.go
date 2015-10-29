package slack

import (
	"encoding/json"
	"fmt"
	"github.com/aphistic/gomol"
	"io/ioutil"
)

var Config *SlackConfig

type SlackConfig struct {
	WebhookUrl string `json:webhookUrl`
}

func InitSlackConfig() error {
	var err error
	Config, err = newSlackConfig("slack.json")

	return err
}

func newSlackConfig(path string) (*SlackConfig, error) {
	var config *SlackConfig
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		gomol.Fatal(err.Error())
		return nil, err
	}

	if err := json.Unmarshal(configFile, &config); err != nil {
		gomol.Fatal(err.Error())
		return nil, err
	}

	return config, nil
}
