package main

import (
	"go-restful-api-service/pkg/routes"
	"net/http"
)

type SortOrder string

const (
	Desc SortOrder = "desc"
	Asc  SortOrder = "asc"
)

func main() {
	r := routes.Route()
	http.ListenAndServe(":9527", r)
}

func queryHandler(r *http.Request) {
}
