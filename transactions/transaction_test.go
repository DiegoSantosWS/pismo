package transactions_test

import (
	"context"
	"log"
	"pismo/connection"
	"pismo/errorsapi"
	"pismo/helpertest"
	"pismo/transactions"
	"testing"
)

func TestFunctionCreateTransactions(t *testing.T) {
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
	op := transactions.RetrieveOperationsTypes()
	w := transactions.RetrieveTransactionWriter()

	t.Run("CREATE NEW TRANSACTION PAGAMENTO", func(t *testing.T) {
		op, err := transactions.CreateTransaction(ctx, op, nil, w, transactions.TransactionInput{
			AccountID:     2,
			OperationType: 4,
			Amount:        10.20,
		})
		if err != nil {
			t.Fatal(err)
		}

		log.Printf("%+v", op)
	})

	t.Run("CREATE NEW TRANSACTION WITHDRAW POSITIVE", func(t *testing.T) {
		_, err := transactions.CreateTransaction(ctx, op, nil, w, transactions.TransactionInput{
			AccountID:     2,
			OperationType: 3,
			Amount:        10.20,
		})
		if err != nil && err == errorsapi.ErrTransactionAmountIsNegative {
			t.Log(err)
		}
	})

	t.Run("CREATE NEW TRANSACTION PARCELING", func(t *testing.T) {
		_, err := transactions.CreateTransaction(ctx, op, nil, w, transactions.TransactionInput{
			AccountID:     2,
			OperationType: 2,
			Amount:        -10.20,
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("CREATE NEW TRANSACTION Sigth", func(t *testing.T) {
		_, err := transactions.CreateTransaction(ctx, op, nil, w, transactions.TransactionInput{
			AccountID:     2,
			OperationType: 1,
			Amount:        -100.20,
		})
		if err != nil {
			t.Fatal(err)
		}
	})
}
