package wechat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const (
	// OfficeAccountAccessTokenEndpoint Endpoint
	OfficeAccountAccessTokenEndpoint = "cgi-bin/token"
)

// OfficeAccountAccessToken OfficeAccountAccessToken
type OfficeAccountAccessToken struct {
	client *Client

	appid     string
	secret    string
	grantType string
}

// NewOfficeAccountAccessToken return instance of OfficeAccountAccessToken
func NewOfficeAccountAccessToken(client *Client) *OfficeAccountAccessToken {
	mpat := &OfficeAccountAccessToken{
		client: client,
	}
	mpat.SetGrantType()
	return mpat
}

// SetSecret SetSecret
func (mpat *OfficeAccountAccessToken) SetSecret(secret string) *OfficeAccountAccessToken {
	mpat.secret = secret
	return mpat
}

// SetAppID SetAppID
func (mpat *OfficeAccountAccessToken) SetAppID(appid string) *OfficeAccountAccessToken {
	mpat.appid = appid
	return mpat
}

// SetGrantType SetGrantType
func (mpat *OfficeAccountAccessToken) SetGrantType() *OfficeAccountAccessToken {
	mpat.grantType = "client_credential"
	return mpat
}

// Validate checks if the operation is valid.
func (mpat *OfficeAccountAccessToken) Validate() error {
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
func (mpat *OfficeAccountAccessToken) Do(ctx context.Context) (*OfficeAccountAccessTokenResponse, error) {
	// Check pre-conditions
	if err := mpat.Validate(); err != nil {
		return nil, errors.Wrap(err, "OfficeAccountAccessToken.Do")
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
		BaseURI:  OfficeAccountBaseHost,
		Endpoint: OfficeAccountAccessTokenEndpoint,
	})
	if err != nil {
		return nil, errors.Wrap(err, "OfficeAccountAccessToken.Do")
	}
	// Return operation response
	ret := new(OfficeAccountAccessTokenResponse)
	if err := mpat.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, errors.Wrap(err, "OfficeAccountAccessToken.Do")
	}
	return ret, nil
}

// Credentials Credentials
func (mpat *OfficeAccountAccessToken) Credentials(ctx context.Context) (*AccessToken, error) {
	res, err := mpat.Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "OfficeAccountAccessToken.Credentials")
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("errcode: %v, errmsg: %s", res.ErrCode, res.ErrMsg)
		return nil, errors.Wrap(err, "OfficeAccountAccessToken.Credentials")
	}
	at := &AccessToken{
		AccessToken: res.AccessToken,
		ExpiresIn:   res.ExpiresIn,
	}
	return at, nil
}

// ToString ToString
func (mpat *OfficeAccountAccessToken) ToString() string {
	return fmt.Sprintf("%s_%s_%s", mpat.grantType, mpat.appid, mpat.secret)
}

// OfficeAccountAccessTokenResponse OfficeAccountAccessTokenResponse
type OfficeAccountAccessTokenResponse struct {
	CommonError
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
