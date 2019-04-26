package wechat

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

// IBasicMessage IBasicMessage
type IBasicMessage interface {
	Body() interface{}
	Validate() error
	BaseURI() string
	Endpoint() string
	Params() url.Values
}

// BasicMessage BasicMessage
type BasicMessage struct {
	client      *Client
	message     IBasicMessage
	accessToken IAccessToken
}

// NewBasicMessage NewBasicMessage
func NewBasicMessage(c *Client, message IBasicMessage, accessToken IAccessToken) *BasicMessage {
	bm := &BasicMessage{
		client:      c,
		message:     message,
		accessToken: accessToken,
	}
	return bm
}

// Send message
func (bm *BasicMessage) Send(ctx context.Context) error {
	// Check pre-conditions
	if err := bm.message.Validate(); err != nil {
		return err
	}
	// accessToken
	at := bm.client.BasicAccessToken(bm.accessToken).GetToken(ctx, false)
	// url params
	params := bm.message.Params()
	params.Set("access_token", at)
	// body
	bodybyte, err := json.Marshal(bm.message.Body())
	if err != nil {
		return err
	}
	// PerformRequest
	res, err := bm.client.PerformRequest(ctx, PerformRequestOptions{
		Method:   http.MethodPost,
		Params:   params,
		Body:     string(bodybyte),
		BaseURI:  bm.message.BaseURI(),
		Endpoint: bm.message.Endpoint(),
	})
	if err != nil {
		return err
	}
	// Return operation response
	ret := new(CommonError)
	if err := bm.client.decoder.Decode(res.Body, ret); err != nil {
		return err
	}
	if ret.ErrCode != 0 {
		err = errors.New(ret.ErrMsg)
	}
	return err
}
