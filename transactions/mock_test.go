package transactions_test

import (
	"context"
	"errors"
	"log"
	"pismo/errorsapi"
	"pismo/transactions"
	"pismo/utils"
	"testing"
)

type dbMockTransaction struct {
	t      *testing.T
	path   string
	caseID string
}

type dbOpMock struct{}

func (db *dbOpMock) GetOperation(ctx context.Context, opID int64) (op transactions.Operation, err error) {
	return
}

func (db *dbOpMock) GetOperations(ctx context.Context) (op []transactions.Operation, err error) {
	return
}

func (db *dbOpMock) CheckOperation(ctx context.Context, opID int64) (ok bool, err error) {
	switch opID {
	case utils.OpAtSight, utils.OpParceling, utils.OpWithdraw, utils.OpPayment:
		ok = true
	default:
		err = errors.New("Cannot find this operation")
	}

	return
}

// RetrieveOperationsTypes access to interfacer of operations types
func retrieveOperationsTypesMock() transactions.OpTypesGetter {
	return &dbOpMock{}
}

type dbVerifierMock struct{}

func (d dbVerifierMock) GetLimitAccount(ctx context.Context, accID int64) (limit float64, err error) {

	return
}

// TransactionWriter interfaces to writer in table transaction
type TransactionWriter interface {
	CreateTransactions(ctx context.Context, v transactions.Verifier, input transactions.TransactionInput) (trans transactions.Transaction, err error)
}

type dbTransactionMock struct{}

func (d dbTransactionMock) CreateTransactions(ctx context.Context, v transactions.Verifier, input transactions.TransactionInput) (trans transactions.Transaction, err error) {
	if err = input.SetDate(); err != nil {
		log.Println(err)
		return
	}

	switch input.OperationType {
	case utils.OpAtSight, utils.OpParceling, utils.OpWithdraw:
		trans = transactions.Transaction{
			AccountID:     input.AccountID,
			OperationType: input.OperationType,
			Amount:        input.Amount,
			EventDate:     input.EventDate,
		}
	case utils.OpPayment:
		trans = transactions.Transaction{
			AccountID:     input.AccountID,
			OperationType: input.OperationType,
			Amount:        input.Amount,
			EventDate:     input.EventDate,
		}
	default:
		err = errorsapi.ErrNotFoundTableDB
	}
	return
}

// RetrieveTransactionWriter access to interfacer of operations
func retrieveTransactionWriterMock() TransactionWriter {
	return &dbTransactionMock{}
}

// RetrieveTransactionWriter access to interfacer of operations
func retrieveVerifierMock() transactions.Verifier {
	return &dbVerifierMock{}
}
