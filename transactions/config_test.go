package transactions_test

import (
	"flag"
	"fmt"
	"pismo/helpertest"
	"testing"
)

var update = flag.Bool("update", false, "update result file")
var fileTest = flag.String("case", "", "specifies which test case to update")

// Global vars to be used between tests.
var opID int64

func setEnvsdb(t *testing.T, pgc helpertest.TestContainer) {
	envs := []helpertest.Env{
		{
			Key:   "PG_HOST",
			Value: fmt.Sprint("localhost"),
		},
		{
			Key:   "PG_USER",
			Value: fmt.Sprint("postgres"),
		},
		{
			Key:   "PG_PASS",
			Value: fmt.Sprint("postgres"),
		},
		{
			Key:   "PG_DB",
			Value: fmt.Sprint("pismodb"),
		},
		{
			Key:   "PG_PORT",
			Value: fmt.Sprint(pgc.MappedPort.Int()),
		},
		{
			Key:   "ENV",
			Value: "local",
		},
	}

	helpertest.SetupEnvs(t, envs)
}
