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
	// defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		log.Println(err)
	}
}

func Close() {
	if db == nil {
		return
	}

	if err := db.Close(); err != nil {
		log.Println(err)
	}
}
