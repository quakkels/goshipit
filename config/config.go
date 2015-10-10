package config

import (
	"encoding/json"
	"github.com/aphistic/gomol"
)

func Test() {
	gomol.Info("testing logger in config.go")
	what, _ := json.Marshal(false)
	gomol.Info(string(what))
}
