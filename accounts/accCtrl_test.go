package accounts_test

import (
	"context"
	"encoding/json"
	"flag"
	"os"
	"pismo/accounts"
	"pismo/errorsapi"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var update = flag.Bool("update", false, "update result file")
var fileTest = flag.String("case", "", "specifies which test case to update")

func TestCreateAccount(t *testing.T) {
	casesAcc := []struct {
		Name    string
		Input   accounts.AccountInput
		Result  string
		CaseAcc string
		ErrExp  error
	}{
		{
			Name:    "Case acc 01",
			CaseAcc: "01",
			Input: accounts.AccountInput{
				Document: "123456789",
			},
			Result: "./testdata/accs/1/result.json",
			ErrExp: nil,
		},
		{
			Name:    "Case acc 02",
			CaseAcc: "02",
			Input: accounts.AccountInput{
				Document: "01669923354",
			},
			Result: "./testdata/accs/2/result.json",
			ErrExp: nil,
		},
		{
			Name:    "Case acc 03",
			CaseAcc: "03",
			Input: accounts.AccountInput{
				Document: "",
			},
			Result: "./testdata/accs/3/result.json",
			ErrExp: errorsapi.ErrDocNotFound,
		},
	}

	for _, tc := range casesAcc {
		t.Run(tc.Name, func(t *testing.T) {
			testCreateAccount(t, tc.Input, tc.Name, tc.Result, tc.CaseAcc, tc.ErrExp)
		})
	}
}

func testCreateAccount(t *testing.T, input accounts.AccountInput, name, expResult, caseAcc string, errExp error) {
	if len(*fileTest) > 0 {
		if !strings.Contains(name, *fileTest) {
			t.Skipf("Skipped the case %s on test unity", name)
			return
		}
	}
	w := RetrieveWriteAccountMock(t, expResult, caseAcc)
	got, err := accounts.CreateAccount(context.Background(), w, input)
	if err == errExp {
		return
	}

	if err != errExp {
		t.Error(err)
		return
	}

	if *update {
		createJSONFile(t, expResult, got, true)
		return
	}

	exp := accounts.Account{}
	compareAccount(t, name, exp, got)
}

func createJSONFile(t *testing.T, fName string, data interface{}, indent ...bool) {
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

func readJSONFile(t *testing.T, fName string, data interface{}) {
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

func compareAccount(t *testing.T, name string, exp, got accounts.Account) {
	diff := cmp.Diff(exp, got, cmpopts.IgnoreFields(accounts.Account{}, "CreatedAt"))
	if len(diff) > 0 {
		t.Errorf("%s, %s", name, diff)
	}
}
