package routes

import (
	"go-restful-api-service/pkg/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
	r.HandleFunc("/file", handlers.FileHandler).Methods("GET")
	r.HandleFunc("/404.css", handlers.CssHandler)

	return r
}
