package connection

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"pismo/errorsapi"
	"sync"

	"github.com/jmoiron/sqlx"
	// Used pg drive on sqlx
	_ "github.com/lib/pq"
)

var (
	err  error
	pool *Pool
)

// Pool ...
type Pool struct {
	Mtx *sync.Mutex
	DB  map[string]*sqlx.DB
}

// GetConnection return connection
func GetConnection(ctx context.Context) (*sqlx.DB, error) {

	db = getDatabaseFromPool()
	if db != nil {
		return db, nil
	}

	db, err := sqlx.ConnectContext(ctx, "postgres", GetURI())
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, errorsapi.ErrConnectionDB
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(5)

	addDatabaseToPool(db)

	return db, nil
}

// GetTransaction get transaction
func GetTransaction(ctx context.Context) (tx *sql.Tx, err error) {
	db, err := GetConnection(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	tx, err = db.Begin()
	return
}

// GetURI return url of connection
func GetURI() string {
	var dbURI, dbHost, dbPass, dbUser, dbPort, dbName string
	if len(dbHost) == 0 {
		dbHost = os.Getenv("PG_HOST")
	}

	if len(dbUser) == 0 {
		dbUser = os.Getenv("PG_USER")
	}

	if len(dbPass) == 0 {
		dbPass = os.Getenv("PG_PASS")
	}

	if len(dbPort) == 0 {
		dbPort = os.Getenv("PG_PORT")
	}

	if len(dbName) == 0 {
		dbName = os.Getenv("PG_DB")
	}

	dbURI = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	return dbURI
}

// GetPool of connection
func getPool() *Pool {
	if pool == nil {
		pool = &Pool{
			Mtx: &sync.Mutex{},
			DB:  make(map[string]*sqlx.DB),
		}
	}
	return pool
}

func getDatabaseFromPool() *sqlx.DB {
	var DB *sqlx.DB
	var p *Pool

	p = getPool()

	p.Mtx.Lock()
	DB = p.DB[GetURI()]
	p.Mtx.Unlock()

	return DB
}

// AddDatabaseToPool add connection to pool
func addDatabaseToPool(db *sqlx.DB) {
	var p *Pool

	p = getPool()

	p.Mtx.Lock()
	p.DB[GetURI()] = db
	p.Mtx.Unlock()
}
