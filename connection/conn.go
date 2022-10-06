package connection

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Load star connection
func Load(ctx context.Context) {
	db, err := GetConnection(ctx)
	if err != nil {
		log.Println("ERROR LOAD: ", err)
		return
	}

	err = db.PingContext(ctx)
	if err != nil {
		log.Println(err)
	}
}

// Close close connection case opened
func Close() {
	if db == nil {
		return
	}

	if err := db.Close(); err != nil {
		log.Println(err)
	}
}
