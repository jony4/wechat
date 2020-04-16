package wechat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const (
	// WorkAccessTokenEndpoint Endpoint
	WorkAccessTokenEndpoint = "cgi-bin/gettoken"
)

// WorkAccessToken WorkAccessToken
type WorkAccessToken struct {
	client *Client

	corpid     string
	corpsecret string
	agentid    string
}

// NewWorkAccessToken return instance of WorkAccessToken
func NewWorkAccessToken(client *Client) *WorkAccessToken {
	wat := &WorkAccessToken{
		client: client,
	}
	return wat
}

// SetSecret SetSecret
func (wat *WorkAccessToken) SetSecret(corpsecret string) *WorkAccessToken {
	wat.corpsecret = corpsecret
	return wat
}

// SetAppID SetAppID
func (wat *WorkAccessToken) SetAppID(corpid string) *WorkAccessToken {
	wat.corpid = corpid
	return wat
}

// Validate checks if the operation is valid.
func (wat *WorkAccessToken) Validate() error {
	var invalid []string
	if wat.corpid == "" {
		invalid = append(invalid, "corpid")
	}
	if wat.corpsecret == "" {
		invalid = append(invalid, "corpsecret")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Do Do
func (wat *WorkAccessToken) Do(ctx context.Context) (*WorkAccessTokenResponse, error) {
	// Check pre-conditions
	if err := wat.Validate(); err != nil {
		return nil, errors.Wrap(err, "WorkAccessToken.Do")
	}
	// url params
	params := url.Values{}
	params.Set("corpid", wat.corpid)
	params.Set("corpsecret", wat.corpsecret)
	// PerformRequest
	res, err := wat.client.PerformRequest(ctx, PerformRequestOptions{
		Method:   http.MethodGet,
		Params:   params,
		BaseURI:  WorkBaseURI,
		Endpoint: WorkAccessTokenEndpoint,
	})
	if err != nil {
		return nil, errors.Wrap(err, "WorkAccessToken.Do")
	}
	// Return operation response
	ret := new(WorkAccessTokenResponse)
	if err := wat.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, errors.Wrap(err, "WorkAccessToken.Do")
	}
	return ret, nil
}

// Credentials Credentials
func (wat *WorkAccessToken) Credentials(ctx context.Context) (*AccessToken, error) {
	res, err := wat.Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "WorkAccessToken.Credentials")
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
func (wat *WorkAccessToken) ToString() string {
	return fmt.Sprintf("%s_%s_%s", wat.corpid, wat.corpsecret, wat.agentid)
}

// WorkAccessTokenResponse WorkAccessTokenResponse
type WorkAccessTokenResponse struct {
	CommonError
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
