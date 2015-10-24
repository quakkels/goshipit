package images

import (
	"testing"
)

func newConfigGood() (imgs *ImageOptions, err error) {
	imgs, err = newImages("testing/images_good.json")
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

	if len(conf.imageOptions) == 0 {
		t.Log("expected image options and found none.")
		t.Fail()
	}

	if conf.imageOptions[0].Category != "starwars" {
		t.Log("first image option has the wrong category")
		t.Fail()
	}

	if conf.imageOptions[0].Path != "pictures/shipit/01.png" {
		t.Log("first image options has wrong path")
		t.Fail()
	}
}

func TestCategoryDoesntExist(t *testing.T) {
	// arrange
	imgs, _ := newConfigGood()

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

func TestCategoryReturnsOnlyImage(t *testing.T) {
	// arrange
	imgs, _ := newConfigGood()

	// act
	img, err := imgs.TakeFromCategory("testcat")

	// assert
	if err != nil {
		t.Log("error when taking image from category that should exist")
		t.Log(err.Error())
		t.Fail()
	}

	if img == "" {
		t.Log("missing image option when image with category should exist.")
		t.Fail()
	}

	if img != "testcatpath" {
		t.Log("wrong image path returned for only image option in category.")
		t.Fail()
	}
}

func TestReturnImageWithNoCategory(t *testing.T) {
	// arrange
	imgs, _ := newConfigGood()

	// act
	img, err := imgs.Take()

	// assert
	if err != nil {
		t.Log("error when taking an image.")
		t.Log(err.Error())
		t.Fail()
	}

	if img == "" {
		t.Log("missing an image when should've gotten one")
		t.Fail()
	}
}

func TestTwoTakesReturnsDifferentImages(t *testing.T) {
	// arrange
	imgs, _ := newConfigGood()

	// act
	img1, _ := imgs.Take()
	img2, _ := imgs.Take()

	// assert
	if img1 == img2 {
		t.Log("Got duplicate images when we probably shouldn't")
		t.Fail()
	}
}

func TestRandomNumber(t *testing.T) {
	// arrange, act
	num := getRandomInRange(0, 5)
	num0 := getRandomInRange(0, 5)
	num1 := getRandomInRange(0, 1)

	// assert
	if num < 0 || num > 5 {
		t.Log("number out of range")
		t.Fail()
	}

	if num == num0 {
		t.Log("two random numbers matched when they probably shouldn't")
		t.Fail()
	}

	if num1 < 0 || num1 > 1 {
		t.Log("num2 is out of range.")
		t.Fail()
	}
}

func TestGetCategoriesFromImages(t *testing.T) {
	// arrange
	imgs, _ := newConfigGood()

	// act
	cats := imgs.GetCategories()

	// assert
	if cats["starwars"] != 1 {
		t.Log("Wrong value returned by .GetCategories()")
		t.Log(cats)
		t.Fail()
	}

	if cats["nautical"] != 1 {
		t.Log("Wrong value returned by .GetCategories()")
		t.Log(cats)
		t.Fail()
	}
	if cats["testcat"] != 2 {
		t.Log("Wrong value returned by .GetCategories()")
		t.Log(cats)
		t.Fail()
	}
}
