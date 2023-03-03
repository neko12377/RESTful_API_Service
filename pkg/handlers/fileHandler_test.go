package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFileHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/file/abcd", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FileHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error status: %d", status)
	}
}
