package main

import (
	"go-restful-api-service/pkg/routes"
	"net/http"
)

func main() {
	r := routes.Route()
	http.ListenAndServe(":9527", r)
}

func queryHandler(r *http.Request) {
}
