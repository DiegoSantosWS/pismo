package webservice

import (
	"context"
	"log"
	"net/http"
	"pismo/accounts"
	"pismo/transactions"
	"time"

	"github.com/gorilla/mux"
)

var srv *http.Server

// NewRouter create route
func NewRouter() {
	r := mux.NewRouter()

	//accounts
	r.HandleFunc("/account", accounts.WriteAccount).Methods(http.MethodPost)
	r.HandleFunc("/account/{account_id:[0-9]+}", accounts.ReadAccount).Methods(http.MethodGet)

	//transactions
	r.HandleFunc("/transaction", transactions.WriteTransaction).Methods(http.MethodPost)

	//#nosec
	srv = &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 3,
		ReadTimeout:  time.Second * 3,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}

// Shutdown close http server
func Shutdown(ctx context.Context) {
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
}
