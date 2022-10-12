package transactions

import (
	"context"
	"log"
	"math"
	"pismo/errorsapi"
	"pismo/utils"
)

// CreateTransaction create a new transaction
func CreateTransaction(ctx context.Context, opT OpTypesGetter, v Verifier, w TransactionWriter, in TransactionInput) (trans Transaction, err error) {
	if ok, errOp := opT.CheckOperation(ctx, in.OperationType); !ok || err != nil {
		// LOG ERROR
		log.Println("", errOp)
		err = errOp
		return
	}

	limit, err := v.GetLimitAccount(ctx, in.AccountID)
	if err != nil {
		return
	}

	if err = checkValue(in.OperationType, limit, in.Amount); err != nil {
		return
	}

	trans, err = w.CreateTransactions(ctx, v, in)

	return
}

func checkValue(typeOp int64, limit, value float64) (err error) {
	switch typeOp {
	case utils.OpWithdraw, utils.OpParceling, utils.OpAtSight:
		if !math.Signbit(value) {
			err = errorsapi.ErrTransactionAmountIsNegative
			return
		}

		ok := (limit + value)

		if !math.Signbit(ok) {
			log.Println("jkdhsdhjsd", ok)
			return
		}
		//TODO: ADD ERROR
		return

	default: // pagamento
		if math.Signbit(value) {
			err = errorsapi.ErrTransactionAmountIsPositive
		}
	}
	return
}
