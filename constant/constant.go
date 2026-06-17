package constant

const (
	Success               = "Success"
	ErrFailedParseRequest = "Failed to parse request"
	ErrFailedBadRequest   = "Internal server error"
	ErrFailedCreateWallet = "Failed to create wallet"
)

const (
	TransactionStatusPending  = "PENDING"
	TransactionStatusSuccess  = "SUCCESS"
	TransactionStatusFailed   = "FAILED"
	TransactionStatusReversed = "REVERSED"
)

const (
	TransactionTypeTopup    = "TOPUP"
	TransactionTypePurchase = "PURCHASE"
	TransactionTypeRefund   = "REFUND"
)

var MapTransaction = map[string]bool{
	TransactionTypeTopup:    true,
	TransactionTypePurchase: true,
	TransactionTypeRefund:   true,
}
