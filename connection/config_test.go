package connection_test

import (
	"fmt"
	"os"
	"pismo/connection"
	"pismo/utils"
	"testing"
)

func TestGetURI(t *testing.T) {
	utils.Load("../.env")
	exp := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASS"), os.Getenv("PG_DB"))
	gotURI := connection.GetURI()
	if gotURI != exp {
		t.Error(gotURI)
	}
}
