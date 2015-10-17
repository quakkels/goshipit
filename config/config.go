package config

import (
	"encoding/json"
	"github.com/aphistic/gomol"
	"io/ioutil"
)

type ImageOption struct {
	Category string `json:category`
	Path     string `json:path`
}

type Images struct {
	imageOptions []ImageOption
}

func NewImages(path string) (*Images, error) {
	images := &Images{}
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		gomol.Fatal(err.Error())
		return nil, err
	}

	if err := json.Unmarshal(configFile, &images.imageOptions); err != nil {
		gomol.Fatal(err.Error())
		return nil, err
	}

	return images, nil
}

func (i *Images) TakeFromCategory(category string) (string, error) {
	var imgs []ImageOption
	for _, img := range i.imageOptions {
		if category == img.Category {
			imgs = append(imgs, img)
		}
	}

	return "", nil
}
