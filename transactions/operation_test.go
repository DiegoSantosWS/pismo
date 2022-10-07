package transactions_test

import (
	"context"
	"pismo/connection"
	"pismo/helpertest"
	"pismo/transactions"
	"testing"
)

func TestFunctionGetOperation(t *testing.T) {
	helpertest.CheckSkipTestType(t, helpertest.FunctionalTest)
	ctx := context.Background()

	containerRequest := helpertest.ContainerRequest{
		Request:   helpertest.MountContainersPG("transactions", "postgres:latest", "postgres", "postgres", "pismodb"),
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

	t.Run("READ OPERATION", func(t *testing.T) {
		op, err := transactions.RetrieveOperationsTypes().CheckOperation(ctx, 4)
		if err != nil {
			t.Fatal(err)
		}
		if !op {
			t.Errorf("Operation [%d] types not exists", 4)
		}
	})
}
