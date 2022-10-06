package accounts_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"pismo/accounts"
	"pismo/connection"
	"pismo/helpertest"
	"strings"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// Global vars to be used between tests.
var accID int64

// Since calling the gRPC method directly doesn't go through main, this has to bet set up manually.

func TestAccounts(t *testing.T) {
	helpertest.CheckSkipTestType(t, helpertest.FunctionalTest)

	ctx := context.Background()
	packageName := "accounts"
	workingDir, _ := os.Getwd()
	rootDir := strings.Replace(workingDir, packageName, "", 1)
	mountFrom := fmt.Sprintf("%s/db/init.sql", rootDir)
	mountTo := "/docker-entrypoint-initdb.d/create_tables.sql"

	requestPG := testcontainers.ContainerRequest{
		Name:  "pismo_test_db",
		Image: "postgres:latest", // ANY docker image works here, including dockerized services!
		ExposedPorts: []string{
			//When you use ExposedPorts you have to imagine yourself using docker run -p <port>. When you do so, dockerd
			//maps the selected <port> from inside the container to a random one available on your host.
			"5432/tcp",
		},
		Mounts: testcontainers.Mounts(testcontainers.BindMount(mountFrom, testcontainers.ContainerMountTarget(mountTo))),
		Env: map[string]string{
			"POSTGRES_USER":     fmt.Sprintf("postgres"),
			"POSTGRES_PASSWORD": fmt.Sprint("postgres"),
			"POSTGRES_DB":       fmt.Sprint("pismodb"),
		},
		//WaitingFor is a field you can use to validate when a container is ready. It is important to get this set
		//because it helps to know when the container is ready to receive any traffic. In this, case we check for the
		//logs we know come from Neo4j, telling us that it is ready to accept requests.
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	containerRequest := helpertest.ContainerRequest{
		Request:   requestPG,
		PortToMap: "5432", // Will be used to find out which is the random mapped available port at helpertest.
	}
	pgContainer := helpertest.CreateTestContainer(t, ctx, containerRequest)
	// err := pgContainer.Container.CopyFileToContainer(ctx, "./db/init.sql", "docker-entrypoint-initdb.d/create_tables.sql", 0444)
	// log.Println("OI", err)
	defer func() {
		connection.Close()
		// This defer makes sure that no docker leftover will exist after the tests are done.
		if err := pgContainer.Container.Terminate(ctx); err != nil {
			t.Fatalf("Error terminating request: [ %v ]", err)
		}
	}()

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
			Value: fmt.Sprint(pgContainer.MappedPort.Int()),
		},
		{
			Key:   "ENV",
			Value: "local",
		},
	}

	helpertest.SetupEnvs(t, envs)
	connection.Load(ctx)
	w := accounts.RetrieveWriteAccount()

	t.Run("cREATE ACCOUNT", func(t *testing.T) {
		log.Println("starting tests")
		acc, err := accounts.CreateAccount(ctx, w, accounts.AccountInput{
			Document: "01688797634",
		})
		if err != nil {
			t.Fatal(err)
		}
		accID = acc.ID
	})

	r := accounts.RetrieveReadAccount()
	t.Run("GET ACCOUNT", func(t *testing.T) {
		log.Println("starting tests")
		acc, err := accounts.GetAccount(ctx, r, accID)
		if err != nil {
			t.Error(err)
		}
		if acc.ID != accID {
			t.Errorf("[] ID not match id expected [%d] id receive [%d]", accID, acc.ID)
		}

		log.Printf("%+v", acc)
	})
}
