package controllers

import (
	"encoding/json"
	"net/http"
)

func listCategories(w http.ResponseWriter, req *http.Request) {
	categories := imgs.GetCategories()
	w.Header().Add("Content Type", "text/json")
	json.NewEncoder(w).Encode(categories)
}
