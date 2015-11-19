package slack

import (
	"bytes"
	"encoding/json"
	"github.com/aphistic/gomol"
	"io/ioutil"
	"net/http"
)

type IncomingWebhook struct {
	Username string `json:"username"`
	Text     string `json:"text"`
	Channel  string `json:"channel"`
}

func GetLinkMarkup(path string, text string) string {
	return "<" + path + "|" + text + ">"
}

func GetImageMarkup(imageUrl string) string {
	return "<" + imageUrl + ">"
}

func SendIncomingWebhook(model IncomingWebhook) ([]byte, error) {
	modelJson, err := json.Marshal(model)
	if err != nil {
		gomol.Err("Could not marshal json from IncomingWebhook model. " +
			err.Error())
		return nil, err
	}

	payload := bytes.NewBuffer(modelJson)

	// send post
	req, err := http.NewRequest("POST", Config.WebhookUrl, payload)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		gomol.Err("Could not send IncomingWebhook. " + err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // must read to close?

	if err != nil {
		gomol.Err("Could not .ReadAll() of response body when sending IncomingWebhook. " +
			err.Error())
		return nil, err
	}

	gomol.Info("Sent an IncomingWebhook")
	return body, nil
}
