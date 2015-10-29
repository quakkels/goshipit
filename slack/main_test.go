package slack

import (
	"testing"
)

func newConfigGood() (config *SlackConfig, err error) {
	config, err = newSlackConfig("testing/slack_good.json")
	return
}

func TestNewSlackConfig(t *testing.T) {
	// arrange, act
	conf, err := newConfigGood()

	// assert
	if err != nil {
		t.Log("could not load slack config file")
		t.Log(err.Error())
		t.Fail()
	}

	if conf == nil {
		t.Log("conf is nil")
	}

	if conf.WebhookUrl == "" {
		t.Log("conf.WebhookUrl is empty")
	}

	if conf.WebhookUrl != "http://test" {
		t.Log("WebhookUrl was not the expected value: " + conf.WebhookUrl)
		t.Fail()
	}
}
