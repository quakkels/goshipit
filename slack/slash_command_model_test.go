package slack

import (
	"net/http"
	"strings"
	"testing"
)

func TestNewSlashCommandModel(t *testing.T) {
	// arrange
	myReq, _ := http.NewRequest("POST", "Idontcare", strings.NewReader("token=testtokenvalue"))

	// act
	myModel := NewSlashCommandModel(myReq)

	// assert
	if myModel.Token != "testtokenvalue" {
		t.Log("Wrong value in Token field.")
		t.Fail()
	}
}
