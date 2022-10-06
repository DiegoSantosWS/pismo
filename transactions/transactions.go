package transactions

import (
	"context"
	"log"
	"pismo/errorsapi"
	"pismo/utils"
)

// CreateTransaction create a new transaction
func CreateTransaction(ctx context.Context, opT OpTypesGetter, w TransactionWriter, in TransactionInput) (trans Transaction, err error) {
	if ok, errOp := opT.CheckOperation(ctx, in.OperationType); !ok || err != nil {
		// LOG ERROR
		log.Println("", errOp)
		err = errOp
		return
	}

	if err = checkValue(in.OperationType, in.Amount); err != nil {
		return
	}

	trans, err = w.CreateTransactions(ctx, in)

	return
}

func checkValue(typeOp int64, value float64) (err error) {
	switch typeOp {
	case utils.OpWithdraw, utils.OpParceling, utils.OpAtSight:
		if value >= 0 {
			err = errorsapi.ErrTransactionAmountIsNegative
		}
	default: // pagamento
		if value <= 0 {
			err = errorsapi.ErrTransactionAmountIsPositive
		}
	}
	return
}
