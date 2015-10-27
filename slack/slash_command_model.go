package slack

import (
	"net/http"
)

type SlashCommandModel struct {
	Token       string
	TeamId      string
	TeamDomain  string
	ChannelId   string
	ChannelName string
	UserId      string
	UserName    string
	Command     string
	Text        string
}

func NewSlashCommandModel(req *http.Request) *SlashCommandModel {
	myModel := &SlashCommandModel{}

	myModel.Token = req.FormValue("token")
	myModel.TeamId = req.FormValue("token_id")
	myModel.TeamDomain = req.FormValue("team_domain")
	myModel.ChannelId = req.FormValue("channel_id")
	myModel.ChannelName = req.FormValue("channel_name")
	myModel.UserId = req.FormValue("user_id")
	myModel.UserName = req.FormValue("user_name")
	myModel.Command = req.FormValue("command")
	myModel.Text = req.FormValue("text")

	return myModel
}
