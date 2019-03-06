package wechat

import (
	"context"
	"fmt"
)

const (
	miniProgramEndpoint = "sns/jscode2session"
)

// MiniProgramAuth mini program auth.
type MiniProgramAuth struct {
	client *Client

	appid     string
	secret    string
	jscode    string
	grantType string

	params map[string]string
}

// NewMiniProgramAuthOpts return MiniProgram default opts
func NewMiniProgramAuthOpts() []ClientOptionFunc {
	return []ClientOptionFunc{
		SetBaseURI(MiniProgramBaseURI),
		SetEndpoint(miniProgramEndpoint),
	}
}

// NewMiniProgramAuth return instance of mini program auth
func NewMiniProgramAuth(client *Client) *MiniProgramAuth {
	return &MiniProgramAuth{
		client:    client,
		grantType: "authorization_code",
		params:    map[string]string{},
	}
}

// SetSecret SetSecret
func (mpa *MiniProgramAuth) SetSecret(secret string) *MiniProgramAuth {
	mpa.secret = secret
	return mpa
}

// SetJscode SetJscode
func (mpa *MiniProgramAuth) SetJscode(jscode string) *MiniProgramAuth {
	mpa.jscode = jscode
	return mpa
}

// SetAppID SetAppID
func (mpa *MiniProgramAuth) SetAppID(appid string) *MiniProgramAuth {
	mpa.appid = appid
	return mpa
}

// Validate checks if the operation is valid.
func (mpa *MiniProgramAuth) Validate() error {
	var invalid []string
	if mpa.appid == "" {
		invalid = append(invalid, "AppID")
	}
	if mpa.secret == "" {
		invalid = append(invalid, "Secret")
	}
	if mpa.jscode == "" {
		invalid = append(invalid, "Jscode")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Do Do
func (mpa *MiniProgramAuth) Do(ctx context.Context) (*MiniProgramAuthResponse, error) {
	// Check pre-conditions
	if err := mpa.Validate(); err != nil {
		return nil, err
	}
	mpa.client.PerformRequest(ctx, PerformRequestOptions{})
	return &MiniProgramAuthResponse{}, nil
}

// MiniProgramAuthRequest MiniProgramAuthRequest
type MiniProgramAuthRequest struct {
}

// MiniProgramAuthResponse MiniProgramAuthResponse
type MiniProgramAuthResponse struct {
}
