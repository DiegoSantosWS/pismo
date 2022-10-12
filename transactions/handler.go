package transactions

import (
	"encoding/json"
	"net/http"
	"pismo/apianswer"
)

// WriteTransaction create a new account
func WriteTransaction(w http.ResponseWriter, r *http.Request) {
	transIn := TransactionInput{}
	if err := json.NewDecoder(r.Body).Decode(&transIn); err != nil {
		apianswer.AnswerRequest(w, r, nil, err)
		return
	}

	opT := RetrieveOperationsTypes()
	wr := RetrieveTransactionWriter() // return interface writer
	trans, err := CreateTransaction(r.Context(), opT, nil, wr, transIn)
	if err != nil {
		apianswer.AnswerRequest(w, r, nil, err)
		return
	}

	apianswer.AnswerRequest(w, r, trans, err)
}
