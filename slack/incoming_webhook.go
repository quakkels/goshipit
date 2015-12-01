package slack

import (
	"bytes"
	"encoding/json"
	"github.com/aphistic/gomol"
	"io/ioutil"
	"net/http"
	"strings"
)

type IncomingWebhook struct {
	Username string `json:"username"`
	Text     string `json:"text"`
	Channel  string `json:"channel"`
}

func NewIncomingWebhook(channel string, text string) *IncomingWebhook {
	if !strings.HasPrefix(channel, "#") {
		channel = "#" + channel
	}
	return &IncomingWebhook{Config.BotUsername, text, channel}
}

func GetLinkMarkup(path string, text string) string {
	return "<" + path + "|" + text + ">"
}

func GetImageMarkup(imageUrl string) string {
	return "<" + imageUrl + ">"
}

func SendIncomingWebhook(model IncomingWebhook) (int, error) {

	payload, err := makePostPayload(model)
	if err != nil {
		return 0, err
	}
	req, err := makeRequest(Config.WebhookUrl, payload)
	if err != nil {
		return 0, err
	}

	code, err := sendRequest(req)
	if err != nil {
		return 0, err
	}

	gomol.Info("Sent an IncomingWebhook")
	return code, nil
}

func makePostPayload(model IncomingWebhook) (*bytes.Buffer, error) {
	modelJson, err := json.Marshal(model)
	if err != nil {
		gomol.Err("Could not marshal json from IncomingWebhook model. " +
			err.Error())
		return nil, err
	}

	payload := bytes.NewBuffer(modelJson)
	return payload, nil
}

func makeRequest(url string, payload *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		gomol.Err("Could not create create post request. " + err.Error())
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func sendRequest(req *http.Request) (int, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		gomol.Err("Could not send IncomingWebhook. " + err.Error())
		return 0, err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body) // must read to close?
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}
