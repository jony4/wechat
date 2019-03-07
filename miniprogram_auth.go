package wechat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	// MiniProgramAuthEndpoint Endpoint
	MiniProgramAuthEndpoint = "sns/jscode2session"
)

// MiniProgramAuth mini program auth.
type MiniProgramAuth struct {
	client *Client

	appid     string
	secret    string
	jscode    string
	grantType string
}

// NewMiniProgramAuth return instance of mini program auth
func NewMiniProgramAuth(client *Client) *MiniProgramAuth {
	mpa := &MiniProgramAuth{
		client: client,
	}
	mpa.SetGrantType()
	return mpa
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

// SetGrantType SetGrantType
func (mpa *MiniProgramAuth) SetGrantType() *MiniProgramAuth {
	mpa.grantType = "authorization_code"
	return mpa
}

// Validate checks if the operation is valid.
func (mpa *MiniProgramAuth) Validate() error {
	var invalid []string
	if mpa.appid == "" {
		invalid = append(invalid, "appid")
	}
	if mpa.secret == "" {
		invalid = append(invalid, "secret")
	}
	if mpa.jscode == "" {
		invalid = append(invalid, "jscode")
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
	// url params
	params := url.Values{}
	params.Set("appid", mpa.appid)
	params.Set("secret", mpa.secret)
	params.Set("js_code", mpa.jscode)
	params.Set("grant_type", mpa.grantType)
	// PerformRequest
	res, err := mpa.client.PerformRequest(ctx, PerformRequestOptions{
		Method:   http.MethodGet,
		Params:   params,
		BaseURI:  MiniProgramBaseURI,
		Endpoint: MiniProgramAuthEndpoint,
	})
	if err != nil {
		return nil, err
	}
	// Return operation response
	ret := new(MiniProgramAuthResponse)
	if err := mpa.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// MiniProgramAuthResponse MiniProgramAuthResponse
type MiniProgramAuthResponse struct {
	CommonError
	AppID      string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}
