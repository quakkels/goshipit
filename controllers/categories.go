package controllers

import (
	"github.com/quakkels/goshipit/images"
	"net/http"
	"text/template"
)

type categoriesController struct {
	template *template.Template
}

func (this *categoriesController) list(w http.ResponseWriter, req *http.Request) {
	images.NewImages("somethings")
}
