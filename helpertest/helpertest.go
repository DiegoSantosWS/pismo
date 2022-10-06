package helpertest

import (
	"encoding/json"
	"os"
	"testing"
)

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
	UnitTest        = "unit test"
	IntegrationTest = "integration test"
	FunctionalTest  = "functional test"
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
