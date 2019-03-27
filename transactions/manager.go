package transactions

import (
	"context"

	"github.com/go-openwrks/openwrks/client"
)

type (
	// TransactionsManager ...
	TransactionsManager struct {
		client *client.Client
	}
)

// NewTransactionsManager ...
func NewTransactionsManager(c *client.Client) *TransactionsManager {
	return &TransactionsManager{
		client: c,
	}
}

// ListTransactions will list all transactions for the provided account.
func (m *TransactionsManager) ListTransactions(ctx context.Context, userID, accountID string, opts *ListTransactionsOptions) (*ListTransactionsResponse, error) {
	res, err := m.client.Do(ctx, "GET", "/v1/users/"+userID+"/accounts/"+accountID+"/transactions", nil)

	if err != nil {
		return nil, err
	}

	out := &ListTransactionsResponse{}
	return out, client.TransformResponse(res, out)
}

// GetTransactions will get a Transactions with the ID provided.
func (m *TransactionsManager) GetTransactions(ctx context.Context, userID, accountID, transactionID string) (*Transaction, error) {
	res, err := m.client.Do(ctx, "GET", "/v1/users/"+userID+"/accounts/"+accountID+"/transactions/"+transactionID, nil)

	if err != nil {
		return nil, err
	}

	out := &Transaction{}
	return out, client.TransformResponse(res, out)
}
