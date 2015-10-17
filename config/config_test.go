package config

import (
	"testing"
)

func newConfigGood() (imgs *Images, err error) {
	imgs, err = NewImages("testing/config_good.json")
	return
}

func TestNewImages(t *testing.T) {
	// arrange, act
	conf, err := newConfigGood()

	// assert
	if err != nil {
		t.Log("could not load config file")
		t.Fail()
	}

	if conf.imageOptions == nil {
		t.Log("config was nil. did not load.")
		t.Fail()
	}
}

func TestCategoryDoesntExist(t *testing.T) {
	// arrange
	imgs, err := newConfigGood()

	// act
	img, err := imgs.TakeFromCategory(
		"This cat does not exist and never will ever 123")

	// assert
	if err != nil {
		t.Log("could not handle a category that doesn't exist.")
		t.Fail()
	}

	if img != "" {
		t.Log("expected empty img path when category doesn't exist.")
		t.Fail()
	}
}
