package accounts

import (
	"encoding/json"
	"net/http"
	"pismo/apianswer"
	"pismo/errorsapi"
	"strconv"

	"github.com/gorilla/mux"
)

// ReadAccount reading account
func ReadAccount(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)
	accID, ok := code["account_id"]
	if !ok {
		apianswer.AnswerRequest(w, r, nil, errorsapi.ErrInvalidBody)
		return
	}

	id, _ := strconv.Atoi(accID)
	re := RetrieveReadAccount()
	acc, err := GetAccount(r.Context(), re, int64(id))
	if err != nil {
		apianswer.AnswerRequest(w, r, nil, err)
		return
	}

	apianswer.AnswerRequest(w, r, acc, err)
}

// WriteAccount create a new account
func WriteAccount(w http.ResponseWriter, r *http.Request) {
	accIn := AccountInput{}
	if err := json.NewDecoder(r.Body).Decode(&accIn); err != nil {
		apianswer.AnswerRequest(w, r, nil, err)
		return
	}

	wr := RetrieveWriteAccount() // return interface writer
	acc, err := CreateAccount(r.Context(), wr, accIn)
	if err != nil {
		apianswer.AnswerRequest(w, r, nil, err)
		return
	}

	apianswer.AnswerRequest(w, r, acc, err)
}
