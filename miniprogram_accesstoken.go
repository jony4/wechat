package wechat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	// MiniProgramAccessTokenEndpoint Endpoint
	MiniProgramAccessTokenEndpoint = "cgi-bin/token"
)

// MiniProgramAccessToken MiniProgramAccessToken
type MiniProgramAccessToken struct {
	client *Client

	appid     string
	secret    string
	grantType string
}

// NewMiniProgramAccessToken return instance of MiniProgramAccessToken
func NewMiniProgramAccessToken(client *Client) *MiniProgramAccessToken {
	mpat := &MiniProgramAccessToken{
		client: client,
	}
	mpat.SetGrantType()
	return mpat
}

// SetSecret SetSecret
func (mpat *MiniProgramAccessToken) SetSecret(secret string) *MiniProgramAccessToken {
	mpat.secret = secret
	return mpat
}

// SetAppID SetAppID
func (mpat *MiniProgramAccessToken) SetAppID(appid string) *MiniProgramAccessToken {
	mpat.appid = appid
	return mpat
}

// SetGrantType SetGrantType
func (mpat *MiniProgramAccessToken) SetGrantType() *MiniProgramAccessToken {
	mpat.grantType = "client_credential"
	return mpat
}

// Validate checks if the operation is valid.
func (mpat *MiniProgramAccessToken) Validate() error {
	var invalid []string
	if mpat.appid == "" {
		invalid = append(invalid, "appid")
	}
	if mpat.secret == "" {
		invalid = append(invalid, "secret")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Do Do
func (mpat *MiniProgramAccessToken) Do(ctx context.Context) (*MiniProgramAccessTokenResponse, error) {
	// Check pre-conditions
	if err := mpat.Validate(); err != nil {
		return nil, err
	}
	// url params
	params := url.Values{}
	params.Set("appid", mpat.appid)
	params.Set("secret", mpat.secret)
	params.Set("grant_type", mpat.grantType)
	// PerformRequest
	res, err := mpat.client.PerformRequest(ctx, PerformRequestOptions{
		Method:   http.MethodGet,
		Params:   params,
		BaseURI:  MiniProgramBaseURI,
		Endpoint: MiniProgramAccessTokenEndpoint,
	})
	if err != nil {
		return nil, err
	}
	// Return operation response
	ret := new(MiniProgramAccessTokenResponse)
	if err := mpat.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// Credentials Credentials
func (mpat *MiniProgramAccessToken) Credentials(ctx context.Context) (*AccessToken, error) {
	res, err := mpat.Do(ctx)
	if err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, fmt.Errorf("errcode: %v, errmsg: %s", res.ErrCode, res.ErrMsg)
	}
	at := &AccessToken{
		AccessToken: res.AccessToken,
		ExpiresIn:   res.ExpiresIn,
	}
	return at, nil
}

// ToString ToString
func (mpat *MiniProgramAccessToken) ToString() string {
	return fmt.Sprintf("%s_%s_%s", mpat.grantType, mpat.appid, mpat.secret)
}

// MiniProgramAccessTokenResponse MiniProgramAccessTokenResponse
type MiniProgramAccessTokenResponse struct {
	CommonError
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
