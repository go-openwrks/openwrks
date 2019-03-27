package accounts

type (
	// ListAccountsOptions provides a way of filtering down Accounts.
	ListAccountsOptions struct {
		Page  int64
		Limit int64
	}

	// ListAccountsResponseMetadata ...
	ListAccountsResponseMetadata struct {
		TotalNumberOfRecords int64 `json:"totalNumberOfRecords"`
		TotalNumberOfPages   int64 `json:"totalNumberOfPages"`
		CurrentPageNumber    int64 `json:"pageNumber"`
		PageSize             int64 `json:"pageSize"`
	}

	// ListAccountsResponseLinks ...
	ListAccountsResponseLinks []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	}

	// ListAccountsResponse ...
	ListAccountsResponse struct {
		Data     []Account                    `json:"data"`
		Metadata ListAccountsResponseMetadata `json:"_meta"`
		Links    ListAccountsResponseLinks    `json:"_links"`
	}

	// Account ...
	Account struct {
		AccountID     string        `json:"accountId"`     // A unique identifier for the account
		AccountNumber string        `json:"accountNumber"` // (optional) The bank's account number
		SortCode      string        `json:"sortCode"`      // (optional) The bank's sort code
		IBAN          string        `json:"iban"`          // (optional) The bank's International Bank Account Number (only provided when account number and sort code are empty)
		AccountName   string        `json:"accountName"`   // (optional) The name of the account, typically set by the account owner via their bank
		Provider      string        `json:"provider"`      // The name of the bank that the account is with
		Balance       float64       `json:"balance"`       //The balance of the account
		Currency      string        `json:"currency"`      // A 3 character code that represents the currency
		CollectedDate string        `json:"collectedDate"` // The date and time that the account data was retrieved from the bank
		ConsentStatus ConsentStatus `json:"consentStatus"` // The status of the user's consent with regards to the account.
	}
)
