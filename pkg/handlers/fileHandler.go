package handlers

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type OrderDirection struct {
	value string
}

var (
	Desc = OrderDirection{"Descending"}
	Asc  = OrderDirection{"Ascending"}
)

type OrderConditon struct {
	value string
}

var (
	LastModified = OrderConditon{"lastModified"}
	Size         = OrderConditon{"size"}
	FileName     = OrderConditon{"fileName"}
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	localSystemFilePath := queryParams.Get("localSystemFilePath")
	orderBy := OrderConditon{queryParams.Get("orderBy")}
	orderByDirection := OrderDirection{queryParams.Get("orderByDirection")}
	filterByName := queryParams.Get("filterByName")

	files, err := getFile(localSystemFilePath)
	if err != nil {
		return
	}
	if orderBy.value != "" || orderByDirection.value != "" {
		files = sortFiles(files, orderBy, orderByDirection)
	}
	if filterByName != "" {
		files = filterFilesByName(files, filterByName)
	}
	for _, file := range files {
		fmt.Printf("%s (%d bytes)(modified %v)\n", file.Name(), file.Size(), file.ModTime())
	}
}

func getFile(localSystemFilePath string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(localSystemFilePath)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func sortFiles(files []fs.FileInfo, orderConditon OrderConditon, orderDirection OrderDirection) []fs.FileInfo {
	switch orderConditon.value {
	case FileName.value:
		sort.Slice(files, func(i, j int) bool {
			return strings.ToLower(files[i].Name()) > strings.ToLower(files[j].Name())
		})
	case Size.value:
		sort.Slice(files, func(i, j int) bool {
			return files[i].Size() > files[j].Size()
		})
	case LastModified.value:
		sort.Slice(files, func(i, j int) bool {
			return files[i].ModTime().After(files[j].ModTime())
		})
	default:
	}

	if orderDirection.value == Asc.value {
		reverseFileOrder(files)
	}

	return files
}

func reverseFileOrder(slice []fs.FileInfo) {
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func filterFilesByName(files []fs.FileInfo, name string) []fs.FileInfo {
	var filteredFiles []fs.FileInfo
	for _, file := range files {
		if strings.Contains(strings.ToLower(file.Name()), strings.ToLower(name)) {
			filteredFiles = append(filteredFiles, file)
		}
	}
	return filteredFiles
}
