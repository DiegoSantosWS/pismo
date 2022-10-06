package accounts_test

import (
	"context"
	"math/rand"
	"pismo/accounts"
	"pismo/errorsapi"
	"testing"
	"time"
)

type dbMockAccount struct {
	t      *testing.T
	path   string
	caseID string
}

func (db dbMockAccount) CreateAccount(ctx context.Context, input accounts.AccountInput) (acc accounts.Account, err error) {
	var idn int
	switch db.caseID {
	case "01":
		idn = 200
	case "02":
		idn = 300
	case "03":
		idn = 400
	}

	acc = accounts.Account{
		ID:        int64(rand.Intn(idn)),
		Document:  input.Document,
		CreatedAt: input.CreatedAt,
	}
	return
}

func (db dbMockAccount) GetAccount(ctx context.Context, accID int64) (acc accounts.Account, err error) {
	switch accID {
	case 1:
		acc = accounts.Account{
			ID:        1,
			Document:  "01648808890",
			CreatedAt: time.Date(2022, time.October, 4, 12, 0, 0, 0, time.UTC),
		}
	case 2:
		acc = accounts.Account{
			ID:        2,
			Document:  "01245697845",
			CreatedAt: time.Date(2022, time.October, 5, 12, 0, 0, 0, time.UTC),
		}
	case 3:
		acc = accounts.Account{
			ID:        3,
			Document:  "01648808890",
			CreatedAt: time.Date(2022, time.October, 6, 12, 0, 0, 0, time.UTC),
		}
	default:
		err = errorsapi.ErrNotFoundTableDB
	}

	return
}

// RetrieveReadAccount access the interface to read accounts
func RetrieveReadAccountMock() accounts.Reader {
	return &dbMockAccount{}
}

// RetrieveWriteAccount access the interface to write new account
func RetrieveWriteAccountMock(t *testing.T, path, cAcc string) accounts.Writer {
	return &dbMockAccount{
		t:      t,
		path:   path,
		caseID: cAcc,
	}
}
