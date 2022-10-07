package transactions

import (
	"context"
	"log"
	"pismo/connection"
	"pismo/errorsapi"
)

func checkOperation(ctx context.Context, opID int64) (ok bool, err error) {
	db, err := connection.GetConnection(ctx)
	if err != nil {
		log.Println(err)
		err = errorsapi.ErrDataBaseConnect
		return
	}

	op := struct {
		operation int64 `db:"id"`
	}{}
	querySelect := `SELECT id FROM operation_types WHERE id = $1`
	err = db.QueryRowContext(ctx, querySelect, opID).Scan(&op.operation)
	if err != nil {
		log.Println("Query operation: ", err)
		err = errorsapi.ErrFindTableDB
		return
	}

	if op.operation == opID {
		ok = true
	}

	return
}

func createTransactions(ctx context.Context, input TransactionInput) (trans Transaction, err error) {
	db, err := connection.GetTransaction(ctx)
	if err != nil {
		return
	}
	defer db.Rollback()

	var id int64
	queryInsert := `INSERT INTO transactions(account_id, operation_type_id, amount, event_date) VALUES ($1, $2, $3, $4) RETURNING id`
	err = db.QueryRowContext(ctx, queryInsert, input.AccountID, input.OperationType, input.Amount, input.EventDate).Scan(&id)
	if err != nil {
		log.Println("Query: ", err)
		return
	}

	if err = db.Commit(); err != nil {
		log.Println(err)
		return
	}

	if id > 0 {
		trans = Transaction{
			TransactionID: id,
			AccountID:     input.AccountID,
			OperationType: input.OperationType,
			Amount:        input.Amount,
			EventDate:     input.EventDate,
		}
	}

	return
}
