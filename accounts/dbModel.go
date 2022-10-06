package accounts

import (
	"context"
	"log"
	"pismo/connection"

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

	doc := input.Document
	date := input.CreatedAt

	queryInsert := `INSERT INTO account(doc_number, created_at) VALUES ($1, $2)`
	if _, err = db.ExecContext(ctx, queryInsert, doc, date); err != nil {
		log.Println("Query: ", err)
		return
	}

	if err = db.Commit(); err != nil {
		log.Println(err)
		return
	}

	acc = Account{
		Document:  input.Document,
		CreatedAt: input.CreatedAt,
	}

	return
}
