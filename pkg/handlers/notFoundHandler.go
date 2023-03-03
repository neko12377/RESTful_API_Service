package handlers

import (
	"io/ioutil"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	println("Visitor use wrong url")
	html, err := ioutil.ReadFile("../../static/404.html")
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(html)
}
