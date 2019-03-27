package users

import (
	"context"

	"github.com/go-openwrks/openwrks/client"
)

type (
	UserManager struct {
		client *client.Client
	}
)

func NewUserManager(c *client.Client) *UserManager {
	return &UserManager{
		client: c,
	}
}

// ListUsers will list all users.
func (m *UserManager) ListUsers(ctx context.Context, opts *ListUsersOptions) (*ListUsersResponse, error) {
	res, err := m.client.Do(ctx, "GET", "/v1/users", nil)

	if err != nil {
		return nil, err
	}

	out := &ListUsersResponse{}
	return out, client.TransformResponse(res, out)
}

// GetUser will get a user with the ID provided.
func (m *UserManager) GetUser(ctx context.Context, userID string) (*User, error) {
	res, err := m.client.Do(ctx, "GET", "/v1/users/"+userID, nil)

	if err != nil {
		return nil, err
	}

	out := &User{}
	return out, client.TransformResponse(res, out)
}
