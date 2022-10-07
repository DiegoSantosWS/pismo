package utils_test

import (
	"os"
	"pismo/utils"
	"testing"
)

func TestLoad(t *testing.T) {
	utils.Load("../.env")

	t.Run("Test read env", func(t *testing.T) {
		gotEnvH := os.Getenv("PG_HOST")
		exp := "localhost"
		if gotEnvH != exp {
			t.Errorf("the value [%s] not match with exp [%s]", gotEnvH, exp)
		}
	})
}
