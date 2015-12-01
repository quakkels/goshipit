package slack

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLinkMarkup(t *testing.T) {
	result := GetLinkMarkup("http://path", "Text to click")

	if "<http://path|Text to click>" != result {
		t.Log("Got wrong link.")
		t.Fail()
	}
}

func TestGetImageMarkup(t *testing.T) {
	result := GetImageMarkup("http://path/image.png")

	if "<http://path/image.png>" != result {
		t.Log("Got wrong link.")
		t.Fail()
	}
}

func TestMakePostPayload(t *testing.T) {
	// arrange
	model := IncomingWebhook{
		"username",
		"text",
		"channel",
	}
	expecting := bytes.NewBuffer(
		[]byte("{\"username\":\"username\",\"text\":\"text\",\"channel\":\"channel\"}"))

	// act
	result, err := makePostPayload(model)

	// assert
	if err != nil {
		t.Log("error making payload")
		t.Fail()
	}

	if !bytes.Equal(expecting.Bytes(), result.Bytes()) {
		t.Log("payload content did not match expected content")
		t.Fail()
	}
}

func TestMakeRequest(t *testing.T) {
	// arrange
	url := "http://localhost:8080/somewhere"
	payloadString := "{\"username\":\"username\",\"text\":\"text\",\"channel\":\"channel\"}"

	// act
	result, err := makeRequest(url, bytes.NewBuffer([]byte(payloadString)))

	// assert
	if err != nil {
		t.Log("failed to create request.")
		t.Fail()
	}

	if "POST" != result.Method {
		t.Log("request was not a post")
		t.Fail()
	}

	if url != result.URL.String() {
		t.Log("Request's url was incorrect. Expected: " +
			url +
			", received: " +
			result.URL.String())
		t.Fail()
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Log("could not read request body")
		t.Fail()
	}

	if string(body) != payloadString {
		t.Log("payload did not match request body")
		t.Log(payloadString)
		t.Log(string(body))
		t.Fail()
	}
}

func TestSendRequest(t *testing.T) {
	// arrange
	ms := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

		}))
	defer ms.Close()

	req, _ := http.NewRequest(
		"POST",
		ms.URL,
		bytes.NewBuffer([]byte("{\"one\":\"1\"}")))
	req.Header.Add("Content-Type", "application/json")

	// act
	result, err := sendRequest(req)

	// assert
	if err != nil {
		t.Log("there was an error sending a request.")
		t.Log(err.Error())
		t.Fail()
	}

	if result != 200 {
		t.Logf("result expected to be 200, but was actually: %d", result)
		t.Fail()
	}
}

func TestSendIncomingWebhook(t *testing.T) {
	// arrange
	ms := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

		}))
	defer ms.Close()
	Config = &SlackConfig{"botname", ms.URL}
	t.Log(ms.URL)
	t.Log(Config.WebhookUrl)

	model := IncomingWebhook{}
	model.Channel = "Channel"
	model.Text = "Text"
	model.Username = "Username"

	t.Logf("%v", model)
	// act
	result, err := SendIncomingWebhook(model)

	// assert
	if err != nil {
		t.Log("could not send an incoming webhook")
		t.Log(err.Error())
		t.Fail()
	}

	if result != 200 {
		t.Log("response was not a 200 code")
		t.Fail()
	}
}

func TestNewIncomingWebhookWithoutChannelHash(t *testing.T) {
	// arrange
	Config = &SlackConfig{"botname", "http://example.com"}

	// act
	result := NewIncomingWebhook("thischannel", "thistext")

	// assert
	if result == nil {
		t.Log("result was nil")
		t.Fail()
	}

	if result.Channel != "#thischannel" {
		t.Log("channel was wrong")
		t.Fail()
	}

	if result.Text != "thistext" {
		t.Log("text was wrong")
		t.Fail()
	}

	if result.Username != "botname" {
		t.Log("botusername was wrong")
		t.Fail()
	}
}

func TestNewIncomingWebhookWithChannelHash(t *testing.T) {
	// arrange
	Config = &SlackConfig{"botname", "http://example.com"}

	// act
	result := NewIncomingWebhook("#thischannel", "thistext")

	// assert
	if result == nil {
		t.Log("result was nil")
		t.Fail()
	}

	if result.Channel != "#thischannel" {
		t.Log("channel was wrong")
		t.Fail()
	}
}
