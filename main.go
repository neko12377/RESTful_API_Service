package main

import (
	"go-restful-api-service/pkg/routes"
	"net/http"
)

func main() {
	r := routes.Route()
	println("Listining on localhost:9527")
	http.ListenAndServe(":9527", r)
}
