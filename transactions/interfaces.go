package transactions

import (
	"context"
	"log"
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

// TransactionWriter interfaces to writer in table transaction
type TransactionWriter interface {
	CreateTransactions(ctx context.Context, input TransactionInput) (trans Transaction, err error)
}

type dbTransaction struct{}

func (d dbTransaction) CreateTransactions(ctx context.Context, input TransactionInput) (trans Transaction, err error) {
	if err = input.SetDate(); err != nil {
		log.Println(err)
		return
	}
	return createTransactions(ctx, input)
}

// RetrieveTransactionWriter access to interfacer of operations
func RetrieveTransactionWriter() TransactionWriter {
	return &dbTransaction{}
}
