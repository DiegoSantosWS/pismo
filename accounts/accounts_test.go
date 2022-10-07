package accounts_test

import (
	"context"
	"fmt"
	"log"
	"pismo/accounts"
	"pismo/connection"
	"pismo/helpertest"
	"testing"
)

// Global vars to be used between tests.
var accID int64

func TestAccounts(t *testing.T) {
	helpertest.CheckSkipTestType(t, helpertest.FunctionalTest)
	ctx := context.Background()

	containerRequest := helpertest.ContainerRequest{
		Request:   helpertest.MountContainersPG("accounts", "postgres:latest", "postgres", "postgres", "pismodb"),
		PortToMap: "5432", // Will be used to find out which is the random mapped available port at helpertest.
	}

	pgContainer := helpertest.CreateTestContainer(ctx, t, containerRequest)
	defer func() {
		connection.Close()
		// This defer makes sure that no docker leftover will exist after the tests are done.
		if err := pgContainer.Container.Terminate(ctx); err != nil {
			t.Fatalf("Error terminating request: [ %v ]", err)
		}
	}()

	setEnvsdb(t, pgContainer)
	connection.Load(ctx)
	w := accounts.RetrieveWriteAccount()

	t.Run("CREATE ACCOUNT", func(t *testing.T) {
		acc, err := accounts.CreateAccount(ctx, w, accounts.AccountInput{
			Document: "01688797634",
		})
		if err != nil {
			t.Fatal(err)
		}
		accID = acc.ID
	})

	r := accounts.RetrieveReadAccount()
	t.Run("GET ACCOUNT CREATED", func(t *testing.T) {
		log.Println("starting tests")
		acc, err := accounts.GetAccount(ctx, r, accID)
		if err != nil {
			t.Error(err)
		}
		if acc.ID != accID {
			t.Errorf("ID not match id expected [%d] id receive [%d]", accID, acc.ID)
		}

		log.Printf("%+v", acc)
	})

	t.Run("GET ACCOUNT 1", func(t *testing.T) {
		acc, err := accounts.GetAccount(ctx, r, 1)
		if err != nil {
			t.Error(err)
		}
		if acc.ID != 1 {
			t.Errorf("ID not match id expected [%d] id receive [%d]", accID, acc.ID)
		}

		log.Printf("%+v", acc)
	})
}

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
