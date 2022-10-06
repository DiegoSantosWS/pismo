package accounts

import (
	"pismo/errorsapi"
	"time"
)

// AccountInput used to create a new account
type AccountInput struct {
	ID        int64     `db:"id" json:"id"`
	Document  string    `db:"doc_number" json:"doc_number"`
	CreatedAt time.Time `db:"created_at" json:"-"`
}

func (acc *AccountInput) setDate() (err error) {
	if acc == nil {
		return errorsapi.ErrInputEmpty
	}
	acc.CreatedAt = time.Now()
	return

}

func (acc *AccountInput) validateDocument() (err error) {
	if acc == nil {
		return errorsapi.ErrInputEmpty
	}

	if len(acc.Document) == 0 {
		err = errorsapi.ErrDocNotFound
	}

	return
}

// Account used to create a new account
type Account struct {
	ID        int64     `db:"id" json:"id"`
	Document  string    `db:"doc_number" json:"doc_number"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
