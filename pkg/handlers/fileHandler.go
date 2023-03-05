package handlers

import (
	"encoding/json"
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

type Directory struct {
	IsDirectory bool     `jason:"isDirectory"`
	Files       []string `jason:"files"`
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	localSystemFilePath := queryParams.Get("localSystemFilePath")
	orderBy := OrderConditon{queryParams.Get("orderBy")}
	orderByDirection := OrderDirection{queryParams.Get("orderByDirection")}
	filterByName := queryParams.Get("filterByName")

	files, err := getFile(localSystemFilePath)
	if err != nil {
		if strings.Contains(err.Error(), "not a directory") {
			content, err := readFile(localSystemFilePath)
			if err != nil {
				return
			}
			w.Write(content)
			return
		}

		http.ServeFile(w, r, "../../static/404.html")
		return
	}

	if orderBy.value != "" || orderByDirection.value != "" {
		files = sortFiles(files, orderBy, orderByDirection)
	}
	if filterByName != "" {
		files = filterFilesByName(files, filterByName)
	}

	directory := Directory{
		IsDirectory: true,
		Files:       []string{},
	}
	w.Header().Set("Content-Type", "application/json")
	for _, file := range files {
		directory.Files = append(directory.Files, formatFileInfo(file))
	}

	jsonData, err := json.Marshal(directory)
	if err != nil {
		println("Error:", err)
		return
	}
	w.Write(jsonData)
}

func getFile(localSystemFilePath string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(localSystemFilePath)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func readFile(localSystemFilePath string) ([]byte, error) {
	content, err := ioutil.ReadFile(localSystemFilePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func formatFileInfo(file fs.FileInfo) string {
	name := file.Name()
	size := file.Size()
	modTime := file.ModTime().Format("2006-01-02 15:04:05")

	return fmt.Sprintf("%s (%s) (%s)\n", name, formatSize(size), modTime)
}

func formatSize(size int64) string {
	const (
		kb = 1024
		mb = kb * 1024
		gb = mb * 1024
		tb = gb * 1024
	)

	switch {
	case size >= tb:
		return fmt.Sprintf("%.2f TB", float64(size)/tb)
	case size >= gb:
		return fmt.Sprintf("%.2f GB", float64(size)/gb)
	case size >= mb:
		return fmt.Sprintf("%.2f MB", float64(size)/mb)
	case size >= kb:
		return fmt.Sprintf("%.2f KB", float64(size)/kb)
	default:
		return fmt.Sprintf("%d bytes", size)
	}
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
