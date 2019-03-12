package invites

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openwrks/openwrks/client"
)

type (
	InviteManager struct {
		client *client.Client
	}
)

func NewInviteManager(c *client.Client) *InviteManager {
	return &InviteManager{
		client: c,
	}
}

func (m *InviteManager) CreateInvite(ctx context.Context, req *CreateInviteRequest) (*CreateInviteResponse, error) {
	jb, _ := json.Marshal(req)
	res, err := m.client.Do(ctx, "POST", "/v1/invite", bytes.NewReader(jb))

	if err != nil {
		return nil, err
	}

	out := &CreateInviteResponse{}
	return out, client.TransformResponse(res, out)
}
