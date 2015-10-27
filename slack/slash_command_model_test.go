package slack

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestNewSlashCommandModel(t *testing.T) {
	// arrange
	values := url.Values{}
	values.Add("token", "testtokenvalue")
	myReq, err := http.NewRequest(
		"POST",
		"http://example.com",
		strings.NewReader(values.Encode()))
	readAll, _ := ioutil.ReadAll(myReq.Body)

	// act
	myModel := NewSlashCommandModel(myReq)

	// assert
	if err != nil {
		t.Log(err.Error())
	}

	if myModel.Token != "testtokenvalue" {
		t.Log("Wrong value in Token field.")
		t.Log("myModel.Token: " + myModel.Token)
		t.Log("myReq.FormValue: " + myReq.FormValue("token"))
		t.Log("ioutil: " + string(readAll))
		t.Fail()
	}
}
