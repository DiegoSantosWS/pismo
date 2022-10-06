package utils

const (
	// OpAtSight represent bought at sight
	OpAtSight int64 = iota + 1
	// OpParceling represent bought at sight
	OpParceling
	// OpWithdraw represent the withdraw (saque)
	OpWithdraw
	// OpPayment represent a payment (pagamento)
	OpPayment
)
