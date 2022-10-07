package helpertest

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// CreateJSONFile create new json file
func CreateJSONFile(t *testing.T, fName string, data interface{}, indent ...bool) {
	var dataByte []byte
	var err error
	if len(indent) > 0 && indent[0] {
		dataByte, err = json.MarshalIndent(data, "", "\t")
	} else {
		dataByte, err = json.Marshal(data)
	}

	if err != nil {
		t.Fatalf("couldnt marshal data to json. error [%s]", err)
		return
	}
	err = os.WriteFile(fName, dataByte, 0600)
	if err != nil {
		t.Fatalf("couldnt create file [%s] error [%s]", fName, err)
	}
}

// ReadJSONFile read the json file
func ReadJSONFile(t *testing.T, fName string, data interface{}) {
	dataByte, err := os.ReadFile(fName) // #nosec
	if err != nil {
		t.Fatalf("couldnt read file. error [%s]", err)
		return
	}
	err = json.Unmarshal(dataByte, data)
	if err != nil {
		t.Fatalf("couldnt unmarshal data to json. error [%s]", err)
	}
}

// TestType is used to declare which kind of test is supposed to be run
type TestType string

const (
	// UnitTest command unit test
	UnitTest = "unit test"
	// IntegrationTest command to integration test
	IntegrationTest = "integration test"
	// FunctionalTest command to functional test
	FunctionalTest = "functional test"
)

func (t TestType) toBool() bool {
	if t == FunctionalTest {
		return false
	}
	return true
}

// CheckSkipTestType skips the test if testing.Short flag doesn't match the
// desired test type.
func CheckSkipTestType(t *testing.T, testType TestType) {
	if testing.Short() != testType.toBool() {
		t.Skip()
	}
}

// Env represent a env to system
type Env struct {
	Key, Value string
}

// SetupEnvs is used to populate Envs that are to be used on functional tests. */
func SetupEnvs(t *testing.T, envs []Env) {
	for _, env := range envs {
		if err := os.Setenv(env.Key, env.Value); err != nil {
			t.Fatalf("[ helper ] failed to set env: [ %v ]", err)
		}
	}
}

// MountContainersPG mount container
func MountContainersPG(pkgName, imageName, usr, pass, dbn string) (requestPG testcontainers.ContainerRequest) {
	workingDir, _ := os.Getwd()
	rootDir := strings.Replace(workingDir, pkgName, "", 1)
	mountFrom := fmt.Sprintf("%s/db/init.sql", rootDir)
	mountTo := "/docker-entrypoint-initdb.d/create_tables.sql"

	requestPG = testcontainers.ContainerRequest{
		Name:  fmt.Sprintf("%s_test_db", pkgName),
		Image: imageName, // ANY docker image works here, including dockerized services!
		ExposedPorts: []string{
			//When you use ExposedPorts you have to imagine yourself using docker run -p <port>. When you do so, dockerd
			//maps the selected <port> from inside the container to a random one available on your host.
			"5432/tcp",
		},
		Mounts: testcontainers.Mounts(testcontainers.BindMount(mountFrom, testcontainers.ContainerMountTarget(mountTo))),
		Env: map[string]string{
			"POSTGRES_USER":     usr,
			"POSTGRES_PASSWORD": pass,
			"POSTGRES_DB":       dbn,
		},
		//WaitingFor is a field you can use to validate when a container is ready. It is important to get this set
		//because it helps to know when the container is ready to receive any traffic. In this, case we check for the
		//logs we know come from Neo4j, telling us that it is ready to accept requests.
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	return
}
