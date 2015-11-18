package slack

import (
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
