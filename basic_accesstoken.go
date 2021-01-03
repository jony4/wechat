package wechat

import (
	"context"
	"fmt"
	"time"
)

const (
	safeDuration   = 600 * time.Second
	cachekeyPrefix = "jony4/wechat."
)

// IAccessToken AccessToken接口，不同类型应用只需要实现该接口即可管理 accesstoken
type IAccessToken interface {
	Credentials(ctx context.Context) (*AccessToken, error)
	ToString() string
}

// AccessToken AccessToken
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// BasicAccessToken BasicAccessToken
type BasicAccessToken struct {
	client *Client
	iat    IAccessToken
}

// NewBasicAccessToken NewBasicAccessToken
func NewBasicAccessToken(c *Client, accessToken IAccessToken) *BasicAccessToken {
	bat := &BasicAccessToken{
		client: c,
		iat:    accessToken,
	}
	return bat
}

func (bat *BasicAccessToken) cacheKey() string {
	return MD5Sum(fmt.Sprintf("%s%s", cachekeyPrefix, bat.iat.ToString()))
}

// SetToken SetToken
func (bat *BasicAccessToken) SetToken(ctx context.Context, at *AccessToken) error {
	lifetime := at.ExpiresIn - int64(safeDuration/time.Second)
	bat.client.cache.Set(ctx, bat.cacheKey(), at.AccessToken, time.Duration(lifetime)*time.Second)
	return nil
}

// GetToken GetToken
func (bat *BasicAccessToken) GetToken(ctx context.Context, refresh bool) (accessToken string) {
	if !refresh {
		value, err := bat.client.cache.Get(ctx, bat.cacheKey())
		bat.client.tracef("GetToken cache get err: %v", err)
		if err == nil {
			return value.(string)
		}
		refresh = true
	}
	if refresh {
		at, err := bat.iat.Credentials(ctx)
		if err != nil {
			bat.client.errorf("GetToken Credentials err: %v", err)
			return
		}
		if err := bat.SetToken(ctx, at); err != nil {
			bat.client.errorf("GetToken SetToken err: %v", err)
			return
		}
		accessToken = at.AccessToken
	}
	return
}

// RefreshToken RefreshToken
func (bat *BasicAccessToken) RefreshToken(ctx context.Context) string {
	return bat.GetToken(ctx, true)
}
