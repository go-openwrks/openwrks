package transactions

type (
	// TransactionStatus ...
	TransactionStatus string
)

const (
	// TransactionStatusBooked ...
	TransactionStatusBooked TransactionStatus = "Booked"
	// TransactionStatusPending ...
	TransactionStatusPending TransactionStatus = "Pending"
)
