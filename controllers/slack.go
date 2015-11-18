package controllers

import (
	"bytes"
	"github.com/aphistic/gomol"
	"github.com/quakkels/goshipit/images"
	"github.com/quakkels/goshipit/slack"
	"net/http"
	"strconv"
)

func slash(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		slashCommand := slack.NewSlashCommandModel(req)
		gomol.Infof("slash command received: %f", slashCommand)

		w.Header().Add("Content Type", "text/plain")

		if slashCommand.Text == "categories" {
			cats := images.Images.GetCategories()
			var catBuffer bytes.Buffer
			for key, count := range cats {
				catBuffer.WriteString(key +
					" has " +
					strconv.Itoa(count) +
					" shipit squirrels.\n")
			}

			catBuffer.WriteTo(w)
		} else if slashCommand.Text == "" {
			image, err := images.Images.Take()
			if err != nil {
				gomol.Err("Failed to .Take() image. " + err.Error())
			}

			incomingWebhook := slack.IncomingWebhook{}
			incomingWebhook.Username = slashCommand.UserName
			incomingWebhook.Channel = slashCommand.ChannelName
			incomingWebhook.Text = slack.GetImageMarkup(image)
			gomol.Info("image: " + image)

			slack.SendIncomingWebhook(incomingWebhook)
		} else {
			b := bytes.NewBufferString("Command not recognized.")
			b.WriteTo(w)
		}
	}
}
