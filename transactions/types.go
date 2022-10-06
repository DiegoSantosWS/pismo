package transactions

import "time"

// TrasactionInput used to create a new transaction
type TrasactionInput struct {
	AccountID    int64     `json:"id"`
	OprationType int64     `json:"doc_number"`
	Amount       float64   `json:"client_id"`
	EventDate    time.Time `json:"-"`
}
