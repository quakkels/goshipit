package controllers

import (
	"bytes"
	"github.com/aphistic/gomol"
	"github.com/quakkels/goshipit/slack"
	"net/http"
)

func slash(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		slashCommand := slack.NewSlashCommandModel(req)
		gomol.Infof("slash command received: %f", slashCommand)

		w.Header().Add("Content Type", "text/plain")
		b := bytes.NewBufferString("Slash command received.")

		b.WriteTo(w)
	}
}
