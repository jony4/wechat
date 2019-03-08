package wechat

import (
	"context"
	"fmt"
	"time"
)

const (
	safeDuration   time.Duration = 600 * time.Second
	cachekeyPrefix               = "jony4wechat."
)

// IAccessToken IAccessToken
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
		if err != nil && err != ErrCacheKeyNotExist {
			return
		} else if err == ErrCacheKeyNotExist {
			refresh = true
		} else {
			accessToken = value.(string)
		}
	}
	if refresh {
		at, err := bat.iat.Credentials(ctx)
		if err != nil {
			return
		}
		if err := bat.SetToken(ctx, at); err != nil {
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
