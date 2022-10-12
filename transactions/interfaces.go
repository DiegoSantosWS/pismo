package transactions

import (
	"context"
	"log"
	"pismo/accounts"
)

// OpTypesGetter represent the operations types
type OpTypesGetter interface {
	GetOperations(ctx context.Context) (op []Operation, err error)
	GetOperation(ctx context.Context, opID int64) (op Operation, err error)
	CheckOperation(ctx context.Context, opID int64) (ok bool, err error)
}

type dbOp struct{}

func (db dbOp) GetOperation(ctx context.Context, opID int64) (op Operation, err error) {
	return
}

func (db dbOp) GetOperations(ctx context.Context) (op []Operation, err error) {
	return
}

func (db dbOp) CheckOperation(ctx context.Context, opID int64) (ok bool, err error) {
	return checkOperation(ctx, opID)
}

// RetrieveOperationsTypes access to interfacer of operations types
func RetrieveOperationsTypes() OpTypesGetter {
	return &dbOp{}
}

// Verifier interface verifier
type Verifier interface {
	GetLimitAccount(ctx context.Context, accID int64) (limit float64, err error)
}

type dbVerifier struct{}

func (d dbVerifier) GetLimitAccount(ctx context.Context, accID int64) (limit float64, err error) {
	limit, err = accounts.GetLimitAccount(ctx, accID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

// TransactionWriter interfaces to writer in table transaction
type TransactionWriter interface {
	CreateTransactions(ctx context.Context, v Verifier, input TransactionInput) (trans Transaction, err error)
}

type dbTransaction struct{}

func (d dbTransaction) CreateTransactions(ctx context.Context, v Verifier, input TransactionInput) (trans Transaction, err error) {

	if err = input.SetDate(); err != nil {
		log.Println(err)
		return
	}

	trans, err = createTransactions(ctx, input)

	//todo: update value of account
	return
}

// RetrieveTransactionWriter access to interfacer of operations
func RetrieveTransactionWriter() TransactionWriter {
	return &dbTransaction{}
}

// RetrieveVerifier access to interfacer of operations
func RetrieveVerifier() Verifier {
	return &dbVerifier{}
}
