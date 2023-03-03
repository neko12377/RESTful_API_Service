package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	placeholders := mux.Vars(r)
	localSystemFilePath := placeholders["localSystemFilePath"]

	queryParams := r.URL.Query()
	orderBy := queryParams.Get("orderBy")
	orderByDirection := queryParams.Get("orderByDirection")
	filterByName := queryParams.Get("filterByName")

	println(localSystemFilePath)
	println(orderBy)
	println(orderByDirection)
	println(filterByName)
	getFile(localSystemFilePath)
}

func getFile(localSystemFilePath string) {
	println("%s %s", localSystemFilePath, 123456)
	files, err := ioutil.ReadDir("/" + localSystemFilePath)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		fmt.Printf("%s (%d bytes)\n", file.Name(), file.Size())
	}
}
