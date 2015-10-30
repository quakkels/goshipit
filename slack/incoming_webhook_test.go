package slack

import (
	"testing"
)

func TestGetLink(t *testing.T) {
	result := GetLink("http://path", "Text to click")

	if "<http://path|Text to click>" != result {
		t.Log("Got wrong link.")
		t.Fail()
	}
}
