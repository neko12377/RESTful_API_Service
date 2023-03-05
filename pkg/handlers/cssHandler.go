package handlers

import (
	"io/ioutil"
	"net/http"
)

func CssHandler(w http.ResponseWriter, r *http.Request) {
	css, err := ioutil.ReadFile("static/404.css")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	w.Write(css)
}
