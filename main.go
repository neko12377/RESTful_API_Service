package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
  r := mux.NewRouter()
  r.NotFoundHandler = http.HandlerFunc(defaultHandler)
  r.HandleFunc("/file/{localSystemFilePath}", fileHandler).Methods("GET")
  http.ListenAndServe(":9527", r)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  println("Try different url")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
  placeholder := mux.Vars(r)
  localSystemFilePath := placeholder["localSystemFilePath"]
  
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
    
  files, err := ioutil.ReadDir(localSystemFilePath)
  
  if err!= nil {
    fmt.Println("Error:", err)
    return
  }

  for _, file := range files {
    fmt.Printf("%s (%d bytes)\n", file.Name(), file.Size())
  }
}

func queryHandler(r *http.Request) {
}

