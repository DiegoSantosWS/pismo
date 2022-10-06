package router

import (
	"log"
	"net/http"
	"pismo/accounts"
	"time"

	"github.com/gorilla/mux"
)

// NewRouter create route
func NewRouter() *http.Server {
	r := mux.NewRouter()

	//accounts
	r.HandleFunc("/account", accounts.WriteAccount).Methods(http.MethodPost)
	r.HandleFunc("/account/{account_id:[0-9]+}", accounts.ReadAccount).Methods(http.MethodGet)

	//transactions
	r.HandleFunc("/transaction", nil).Methods(http.MethodPost)

	srv := &http.Server{
		Addr: "127.0.0.1:8080",
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

	return srv
}
