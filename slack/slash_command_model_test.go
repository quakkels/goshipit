package slack

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestNewSlashCommandModel(t *testing.T) {
	// arrange
	values := url.Values{}
	values.Add("token", "testtokenvalue")
	values.Add("team_id", "testteamid")
	values.Add("team_domain", "testteamdomain")
	values.Add("channel_id", "testchannelid")
	values.Add("channel_name", "testchannelname")
	values.Add("user_id", "testuserid")
	values.Add("user_name", "testusername")
	values.Add("command", "testcommand")
	values.Add("text", "testtext")
	myReq, err := http.NewRequest(
		"POST",
		"http://example.com",
		strings.NewReader(values.Encode()))
	myReq.Header.Add("Content Type", "application/x-www-form-urlencoded")

	// act
	myModel := NewSlashCommandModel(myReq)

	// assert
	if err != nil {
		t.Log(err.Error())
	}

	if myModel.Token != "testtokenvalue" {
		t.Log("Wrong value in Token field.")
		t.Fail()
	}

	if myModel.TeamId != "testteamid" {
		t.Log("Wrong value in TeamId field.")
		t.Fail()
	}

	if myModel.TeamDomain != "testteamdomain" {
		t.Log("Wrong value in TeamDomain field.")
		t.Fail()
	}

	if myModel.ChannelId != "testchannelid" {
		t.Log("Wrong value in ChannelId field.")
		t.Fail()
	}

	if myModel.ChannelName != "testchannelname" {
		t.Log("Wrong value in ChannelName field.")
		t.Fail()
	}

	if myModel.UserId != "testuserid" {
		t.Log("Wrong value in UserId field.")
		t.Fail()
	}

	if myModel.UserName != "testusername" {
		t.Log("Wrong value in UserName field.")
		t.Fail()
	}

	if myModel.Command != "testcommand" {
		t.Log("Wrong value in Command field.")
		t.Fail()
	}

	if myModel.Text != "testtext" {
		t.Log("Wrong value in Text field.")
		t.Fail()
	}
}
