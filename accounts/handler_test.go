package accounts_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"pismo/accounts"
	"testing"
)

func TestGetterAccount(t *testing.T) {
	path := fmt.Sprintf("http://localhost:8080/account/%s", "12")
	log.Println(path)
	_, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestWriteAccount(t *testing.T) {
	path := fmt.Sprintf("http://localhost:8080/account")
	requestJSON, err := json.Marshal(accounts.AccountInput{
		ID:       15,
		Document: "147892521522",
	})
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(requestJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-type", "application/json")

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(accounts.WriteAccount)
	// handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		return
	}
}
