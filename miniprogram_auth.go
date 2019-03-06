package wechat

import "context"

type MiniProgramAuth struct {
	client *Client
	params map[string]string
}

func NewMiniProgramAuth(client *Client) *MiniProgramAuth {
	return &MiniProgramAuth{
		client: client,
		params: map[string]string{},
	}
}

func (oaAuth *MiniProgramAuth) Set(key, value string) *MiniProgramAuth {
	oaAuth.params[key] = value
	return oaAuth
}

type MiniProgramAuthRequest struct {
}

type MiniProgramAuthResponse struct {
}

func (oaAuth *MiniProgramAuth) Do(ctx context.Context) (*MiniProgramAuthResponse, error) {
	oaAuth.client.PerformRequest(ctx, PerformRequestOptions{})
	return &MiniProgramAuthResponse{}, nil
}
