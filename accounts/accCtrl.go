package accounts

import (
	"context"
	"log"
)

// CreateAccount create a new account
func CreateAccount(ctx context.Context, w Writer, input AccountInput) (acc Account, err error) {
	if err = input.setDate(); err != nil {
		log.Println(err)
		return
	}

	if err = input.validateDocument(); err != nil {
		log.Println(err)
		return
	}

	acc, err = w.CreateAccount(ctx, input)
	return
}

// GetAccount read a account
func GetAccount(ctx context.Context, r Reader, accID int64) (acc Account, err error) {
	acc, err = r.GetAccount(ctx, accID)
	return
}
