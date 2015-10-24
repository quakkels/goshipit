package images

import (
	"encoding/json"
	"github.com/aphistic/gomol"
	"io/ioutil"
	"math/rand"
	"time"
)

var Images *ImageOptions

type ImageOption struct {
	Category string `json:category`
	Path     string `json:path`
}

type ImageOptions struct {
	imageOptions []ImageOption
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func InitImages() error {
	var err error
	Images, err = newImages("images.json")

	return err
}

func newImages(path string) (*ImageOptions, error) {
	imgs := &ImageOptions{}
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		gomol.Fatal(err.Error())
		return nil, err
	}

	if err := json.Unmarshal(configFile, &imgs.imageOptions); err != nil {
		gomol.Fatal(err.Error())
		return nil, err
	}

	return imgs, nil
}

func (i *ImageOptions) TakeFromCategory(category string) (string, error) {
	var imgs []ImageOption
	for _, img := range i.imageOptions {
		if category == img.Category {
			imgs = append(imgs, img)
		}
	}

	count := len(imgs)

	if count == 0 {
		return "", nil
	}

	if count == 1 {
		return imgs[0].Path, nil
	}

	index := getRandomInRange(0, uint(count-1))
	return imgs[index].Path, nil
}

func (i *ImageOptions) Take() (string, error) {
	if i.imageOptions == nil {
		return "", nil
	}

	count := len(i.imageOptions)
	if count < 1 {
		return "", nil
	}

	index := getRandomInRange(0, uint(count-1))
	return i.imageOptions[index].Path, nil

}

func getRandomInRange(bottom uint, top uint) int {
	return int(bottom) + rand.Intn(int(top)-int(bottom))
}
