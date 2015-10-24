package controllers

import (
	"encoding/json"
	"github.com/quakkels/goshipit/images"
	"net/http"
)

func listCategories(w http.ResponseWriter, req *http.Request) {
	categories := images.Images.GetCategories()
	w.Header().Add("Content Type", "text/json")
	json.NewEncoder(w).Encode(categories)
}
