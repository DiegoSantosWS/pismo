package transactions

import (
	"pismo/errorsapi"
	"time"
)

// TransactionInput used to create a new transaction
type TransactionInput struct {
	AccountID     int64     `json:"account_id"`
	OperationType int64     `json:"operation_type_id"`
	Amount        float64   `json:"amount"`
	EventDate     time.Time `json:"-"`
}

// SetDate set date of event
func (trs *TransactionInput) SetDate() (err error) {
	if trs == nil {
		return errorsapi.ErrInputEmpty
	}
	trs.EventDate = time.Now()
	return

}

// Transaction represent one transactions
type Transaction struct {
	TransactionID int64     `db:"id" json:"transaction_id"`
	AccountID     int64     `db:"account_id" json:"account_id"`
	OperationType int64     `db:"operation_type_id" json:"operation_type_id"`
	Amount        float64   `db:"amount" json:"amount"`
	EventDate     time.Time `db:"event_date" json:"event_date"`
}

// Operation represent one operation type
type Operation struct {
	ID          int64   `json:"id"`
	Description float64 `json:"description"`
}
