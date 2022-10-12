package transactions_test

import (
	"context"
	"log"
	"pismo/errorsapi"
	"pismo/helpertest"
	"pismo/transactions"
	"pismo/utils"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func TestCheckValue(t *testing.T) {
	helpertest.CheckSkipTestType(t, helpertest.UnitTest)

	casesValue := []struct {
		Name   string
		Value  float64
		Limit  float64
		OpType int64
		Err    error
	}{
		{
			Name:   "Case withdraw",
			Value:  -500.00,
			Limit:  5000.00,
			OpType: utils.OpWithdraw,
			Err:    nil,
		},
		// {
		// 	Name:   "Case error withdraw, value positive",
		// 	Value:  450.00,
		// 	Limit:  0,
		// 	OpType: utils.OpWithdraw,
		// 	Err:    errorsapi.ErrTransactionAmountIsNegative,
		// },
		// {
		// 	Name:   "Case OpParceling",
		// 	Value:  -15.25,
		// 	Limit:  0,
		// 	OpType: utils.OpParceling,
		// 	Err:    nil,
		// },
		// {
		// 	Name:   "Case error OpParceling, value negative",
		// 	Value:  -1.00,
		// 	Limit:  0,
		// 	OpType: utils.OpParceling,
		// 	Err:    errorsapi.ErrTransactionAmountIsNegative,
		// },
		// {
		// 	Name:   "Case OpAtSight",
		// 	Value:  -11.99,
		// 	Limit:  0,
		// 	OpType: utils.OpAtSight,
		// 	Err:    nil,
		// },
		// {
		// 	Name:   "Case payment",
		// 	Value:  5000.00,
		// 	Limit:  0,
		// 	OpType: utils.OpPayment,
		// 	Err:    nil,
		// },
		// {
		// 	Name:   "Case error payment, value negative",
		// 	Value:  -5000.00,
		// 	Limit:  0,
		// 	OpType: utils.OpPayment,
		// 	Err:    errorsapi.ErrTransactionAmountIsPositive,
		// },
	}

	for _, tc := range casesValue {
		t.Run(tc.Name, func(t *testing.T) {
			log.Println(tc.Name)
			err := transactions.CheckValue(tc.OpType, tc.Limit, tc.Value)
			if err != nil {
				log.Println(err)
				assert.EqualError(t, err, tc.Err.Error())
			}
		})
	}
}

func TestCreateTransaction(t *testing.T) {
	casesAcc := []struct {
		Name   string
		Input  transactions.TransactionInput
		Result string
		ErrExp error
	}{
		{
			Name: "Case create transaction of payment 01",
			Input: transactions.TransactionInput{
				AccountID:     1,
				OperationType: 4,
				Amount:        18.9,
			},
			Result: "./testdata/trans/acc-1/1/result-payment.json",
			ErrExp: nil,
		},
		{
			Name: "Case create transaction of payment at sight 02",
			Input: transactions.TransactionInput{
				AccountID:     1,
				OperationType: 1,
				Amount:        -188.15,
			},
			Result: "./testdata/trans/acc-1/2/result-AtSight.json",
			ErrExp: nil,
		},
		{
			Name: "Case create transaction of parceling 03",
			Input: transactions.TransactionInput{
				AccountID:     1,
				OperationType: 2,
				Amount:        -1000.11,
			},
			Result: "./testdata/trans/acc-1/3/result-parceling.json",
			ErrExp: nil,
		},
		{
			Name: "Case create transaction of withdraw 04",
			Input: transactions.TransactionInput{
				AccountID:     1,
				OperationType: 3,
				Amount:        -14000.98,
			},
			Result: "./testdata/trans/acc-1/4/result-withdraw.json",
			ErrExp: nil,
		},
		// acc-2
		{
			Name: "Case create transaction of payment acc 2 - 01",
			Input: transactions.TransactionInput{
				AccountID:     2,
				OperationType: 4,
				Amount:        18.9,
			},
			Result: "./testdata/trans/acc-2/1/result-payment.json",
			ErrExp: nil,
		},
		{
			Name: "Case create transaction of payment at sight acc 2 - 02",
			Input: transactions.TransactionInput{
				AccountID:     2,
				OperationType: 1,
				Amount:        -188.15,
			},
			Result: "./testdata/trans/acc-2/2/result-AtSight.json",
			ErrExp: nil,
		},
		{
			Name: "Case create transaction of parceling acc 2 - 03",
			Input: transactions.TransactionInput{
				AccountID:     2,
				OperationType: 2,
				Amount:        -1000.11,
			},
			Result: "./testdata/trans/acc-2/3/result-parceling.json",
			ErrExp: nil,
		},
		{
			Name: "Case create transaction of withdraw acc 2 - 04",
			Input: transactions.TransactionInput{
				AccountID:     2,
				OperationType: 3,
				Amount:        14000.98,
			},
			Result: "./testdata/trans/acc-2/4/result-withdraw.json",
			ErrExp: errorsapi.ErrTransactionAmountIsNegative,
		},
	}

	for _, tc := range casesAcc {
		t.Run(tc.Name, func(t *testing.T) {
			testCreateTransaction(t, tc.Input, tc.Name, tc.Result, tc.ErrExp)
		})
	}
}

func testCreateTransaction(t *testing.T, input transactions.TransactionInput, name, expResult string, errExp error) {
	if len(*fileTest) > 0 {
		if !strings.Contains(expResult, *fileTest) {
			t.Skipf("Skipped the case %s on test unity", name)
			return
		}
	}

	opr := retrieveOperationsTypesMock()
	w := retrieveTransactionWriterMock()
	v := retrieveVerifierMock()
	got, err := transactions.CreateTransaction(context.Background(), opr, v, w, input)
	if errExp != nil {
		assert.EqualError(t, err, errExp.Error())
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	if *update {
		helpertest.CreateJSONFile(t, expResult, got, true)
		return
	}

	exp := transactions.Transaction{}
	helpertest.ReadJSONFile(t, expResult, &exp)
	compareTransaction(t, name, exp, got)
}

func compareTransaction(t *testing.T, name string, exp, got transactions.Transaction) {
	diff := cmp.Diff(exp, got, cmpopts.IgnoreFields(transactions.Transaction{}, "EventDate"))
	if len(diff) > 0 {
		t.Errorf("%s, %s", name, diff)
	}
}
