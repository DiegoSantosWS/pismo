package transactions_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pismo/helpertest"
	"pismo/transactions"
	"testing"
)

func TestWriteAccount(t *testing.T) {
	helpertest.CheckSkipTestType(t, helpertest.UnitTest)

	requestJSON, err := json.Marshal(transactions.TransactionInput{AccountID: 2, OperationType: 2, Amount: -1000.11})
	if err != nil {
		t.Fatal(err)
	}

	path := fmt.Sprintf("http://localhost:8080/transaction")
	_, err = http.NewRequest(http.MethodPost, path, bytes.NewBuffer(requestJSON))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		return
	}
}
