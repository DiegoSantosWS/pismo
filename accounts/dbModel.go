package accounts

import (
	"context"
	"log"
	"pismo/connection"
	"pismo/errorsapi"

	// Used pg drive on sqlx
	_ "github.com/lib/pq"
)

func getAccount(ctx context.Context, accID int64) (acc Account, err error) {
	db, err := connection.GetConnection(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	querySelect := `SELECT id, doc_number, created_at FROM account WHERE id = $1`
	err = db.QueryRowContext(ctx, querySelect, accID).Scan(&acc.ID, &acc.Document, &acc.CreatedAt)
	if err != nil {
		log.Println("Query: ", err)
	}

	return
}

func createAccount(ctx context.Context, input AccountInput) (acc Account, err error) {
	db, err := connection.GetTransaction(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	defer db.Rollback()

	var id int64
	queryInsert := `INSERT INTO account(doc_number, created_at) VALUES ($1, $2) RETURNING id`
	err = db.QueryRowContext(ctx, queryInsert, input.Document, input.CreatedAt).Scan(&id)
	if err != nil {
		log.Println("Query: ", err)
		return
	}

	if err = db.Commit(); err != nil {
		log.Println(err)
		return
	}

	log.Println("id", id)
	acc = Account{
		ID:        id,
		Document:  input.Document,
		CreatedAt: input.CreatedAt,
	}

	return
}

func getLimitAccount(ctx context.Context, accID int64) (limit float64, err error) {
	db, err := connection.GetConnection(ctx)
	if err != nil {
		log.Println("getLimitAccount", err)
		err = errorsapi.ErrConnectionDB
		return
	}

	query := `SELECT avaliable_credit_limit FROM account WHERE id = $1`
	err = db.QueryRowContext(ctx, query, accID).Scan(&limit)

	return
}
