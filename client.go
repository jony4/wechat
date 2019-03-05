package wechat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	c *http.Client
}

func NewClient() *Client {
	return &Client{
		c: http.DefaultClient,
	}
}

type Request http.Request

func NewRequest(method, url string) (*Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	return (*Request)(req), nil
}

type PerformRequestOptions struct {
	Method          string
	Path            string
	Params          url.Values
	Body            interface{}
	ContentType     string
	IgnoreErrors    []int
	Headers         http.Header
	MaxResponseSize int64
}

type Response struct {
}

func (client *Client) PerformRequest(ctx context.Context, opt PerformRequestOptions) (*Response, error) {
	req, err := NewRequest("GET", "https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	res, err := client.c.Do((*http.Request)(req).WithContext(ctx))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	return &Response{}, nil
}

func (client *Client) OfficialAccountAuth() *OfficialAccountAuth {
	return NewOfficialAccountAuth(client)
}

type OfficialAccountAuth struct {
	client *Client
	params map[string]string
}

func NewOfficialAccountAuth(client *Client) *OfficialAccountAuth {
	return &OfficialAccountAuth{
		client: client,
		params: map[string]string{},
	}
}

type OfficialAccountAuthRequest struct {
}

type OfficialAccountAuthResponse struct {
}

func (oaAuth *OfficialAccountAuth) Set(key, value string) *OfficialAccountAuth {
	oaAuth.params[key] = value
	return oaAuth
}

func (oaAuth *OfficialAccountAuth) Do(ctx context.Context) (*OfficialAccountAuthResponse, error) {
	oaAuth.client.PerformRequest(ctx, PerformRequestOptions{})
	return &OfficialAccountAuthResponse{}, nil
}
