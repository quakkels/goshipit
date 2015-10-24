package controllers

import (
	"bufio"
	"github.com/quakkels/goshipit/images"
	"net/http"
	"os"
	"strings"
)

var imgs *images.Images

func init() {
	var err error
	imgs, err = images.NewImages("images.json")
	if err != nil {
		panic(err)
	}
}

func Register() {
	http.HandleFunc("/categories", listCategories)
	http.HandleFunc("/", serveResource)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "static" + req.URL.Path
	contentType := "text/plain"

	if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg"
	} else if strings.HasSuffix(path, ".gif") {
		contentType = "image/gif"
	}

	file, err := os.Open(path)
	if err == nil {
		defer file.Close()
		w.Header().Add("Content Type", contentType)
		br := bufio.NewReader(file)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
