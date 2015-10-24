package controllers

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

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
