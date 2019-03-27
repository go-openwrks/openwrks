package accounts

import (
	"context"

	"github.com/go-openwrks/openwrks/client"
)

type (
	// AccountsManager ...
	AccountsManager struct {
		client *client.Client
	}
)

// NewAccountsManager ...
func NewAccountsManager(c *client.Client) *AccountsManager {
	return &AccountsManager{
		client: c,
	}
}

// ListAccounts will list all accounts.
func (m *AccountsManager) ListAccounts(ctx context.Context, userID string, opts *ListAccountsOptions) (*ListAccountsResponse, error) {
	res, err := m.client.Do(ctx, "GET", "/v1/users/"+userID+"/accounts", nil)

	if err != nil {
		return nil, err
	}

	out := &ListAccountsResponse{}
	return out, client.TransformResponse(res, out)
}

// GetAccounts will get a accounts with the ID provided.
func (m *AccountsManager) GetAccounts(ctx context.Context, userID string, accountID string) (*Account, error) {
	res, err := m.client.Do(ctx, "GET", "/v1/users/"+userID+"/accounts/"+accountID, nil)

	if err != nil {
		return nil, err
	}

	out := &Account{}
	return out, client.TransformResponse(res, out)
}
