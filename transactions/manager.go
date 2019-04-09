package transactions

import (
	"context"
	"net/url"
	"strconv"

	"github.com/go-openwrks/openwrks/client"
)

const (
	transactionDateFormat = "2006-01-02T15:04:05-00:00"
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
	qv := url.Values{}

	if opts != nil {
		if opts.From.Unix() > 0 {
			qv["from"] = []string{opts.From.Format(transactionDateFormat)}
		}

		if opts.To.Unix() > 0 {
			qv["to"] = []string{opts.To.Format(transactionDateFormat)}
		}

		if opts.Page != 0 {
			qv["page"] = []string{strconv.Itoa(int(opts.Page))}
		}

		if opts.Limit != 0 {
			qv["limit"] = []string{strconv.Itoa(int(opts.Limit))}
		} else {
			qv["limit"] = []string{"100"}
		}
	}

	res, err := m.client.Do(ctx, "GET", "/v1/users/"+userID+"/accounts/"+accountID+"/transactions?"+qv.Encode(), nil)

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
