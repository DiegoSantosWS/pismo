package accounts

import "context"

// Writer used to writer in database
type Writer interface {
	CreateAccount(context.Context, AccountInput) (acc Account, err error)
}

// Reader used to reader in database
type Reader interface {
	GetAccount(ctx context.Context, accID int64) (acc Account, err error)
}

type dbAccount struct{}

func (db dbAccount) CreateAccount(ctx context.Context, input AccountInput) (acc Account, err error) {
	return createAccount(ctx, input)
}

func (db dbAccount) GetAccount(ctx context.Context, accID int64) (acc Account, err error) {
	return getAccount(ctx, accID)
}

// RetrieveReadAccount access the interface to read accounts
func RetrieveReadAccount() Reader {
	return &dbAccount{}
}

// RetrieveWriteAccount access the interface to write new account
func RetrieveWriteAccount() Writer {
	return &dbAccount{}
}
