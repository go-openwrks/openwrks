package transactions

type (
	// ListTransactionsOptions provides a way of filtering down Transactions.
	ListTransactionsOptions struct {
		Page  int64
		Limit int64
	}

	// ListTransactionsResponseMetadata ...
	ListTransactionsResponseMetadata struct {
		TotalNumberOfRecords int64 `json:"totalNumberOfRecords"`
		TotalNumberOfPages   int64 `json:"totalNumberOfPages"`
		CurrentPageNumber    int64 `json:"pageNumber"`
		PageSize             int64 `json:"pageSize"`
	}

	// ListTransactionsResponseLinks ...
	ListTransactionsResponseLinks []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	}

	// ListTransactionsResponse ...
	ListTransactionsResponse struct {
		Data     []Transaction                    `json:"data"`
		Metadata ListTransactionsResponseMetadata `json:"_meta"`
		Links    ListTransactionsResponseLinks    `json:"_links"`
	}

	// Transaction ...
	Transaction struct {
		TransactionID        string  `json:"transactionId"`        // A unique identifier for the transaction (does not persist beyond refreshes)
		Amount               float64 `json:"amount"`               // The transaction amount (debits will be negative, credits will be positive)
		Description          string  `json:"description"`          // (optional) The transaction description provided by the bank
		Currency             string  `json:"currency"`             // A 3 character code that represents the currency
		MerchantName         string  `json:"merchantName"`         // (optional) The merchant name provided by the bank
		MerchantCategoryCode string  `json:"merchantCategoryCode"` // (optional) The merchant category code provided by the bank
		Timestamp            string  `json:"timestamp"`            // The date and time that the transaction occurred. Where no time is provided by the banks, the datetime defaults to midnight of the provided timezone
		Type                 string  `json:"type"`                 // An indication of whether the transaction was a debit or credit
		Category             string  `json:"category"`             // The category of the transaction, determined by OpenWrks
		RunningBalance       float64 `json:"runningBalance"`       // The account balance at the time that this transaction occurred
		Status               string  `json:"status"`               // Whether the transaction is 'Booked' or 'Pending'
	}
)
