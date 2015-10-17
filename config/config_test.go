package config

import (
	"testing"
)

func TestOpenConfig(t *testing.T) {
	conf, err := Load("testing/config_good.json")
	if err != nil {
		t.Log("could not load config file")
		t.Fail()
	}

	if conf == nil {
		t.Log("config was nil. did not load.")
		t.Fail()
	}
}
