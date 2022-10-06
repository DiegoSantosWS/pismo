package accounts

import (
	"encoding/json"
	"log"
	"net/http"
	"pismo/apiansower"
	"strconv"

	"github.com/gorilla/mux"
)

// ReadAccount reading account
func ReadAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	code := mux.Vars(r)
	accID, ok := code["account_id"]
	if !ok {
		log.Println(accID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(accID)
	re := RetrieveReadAccount()
	acc, err := GetAccount(r.Context(), re, int64(id))
	apiansower.AnswerRequest(w, r, acc, err)
}

// WriteAccount ...
func WriteAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	accIn := AccountInput{}
	if err := json.NewDecoder(r.Body).Decode(&accIn); err != nil {
		apiansower.AnswerRequest(w, r, nil, err)
		return
	}

	wr := RetrieveWriteAccount()
	acc, err := CreateAccount(r.Context(), wr, accIn)
	if err != nil {
		apiansower.AnswerRequest(w, r, nil, err)
		return
	}

	apiansower.AnswerRequest(w, r, acc, err)
}
