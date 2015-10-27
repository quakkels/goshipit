package slack

import (
	"net/http"
	"testing"
)

func TestNewSlashCommandModel(t *testing.T) {
	// arrange
	myReq := &http.Request{}
	// todo: set Token to "testtokenvalue"
	// todo: set post body

	// act
	myModel := NewSlashCommandModel(myReq)

	// assert
	if myModel.Token != "testtokenvalue" {
		t.Log("Wrong value in Token field.")
		t.Fail()
	}
}
