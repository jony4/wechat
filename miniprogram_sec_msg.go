package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// Endpoint
const (
	MiniProgramSecMsgEndpoint = "wxa/msg_sec_check"
)

// MiniProgramSecMsg MiniProgramSecMsg
type MiniProgramSecMsg struct {
	client *Client

	accessToken string
	message     string
}

// NewMiniProgramSecMsg return instance of NewMiniProgramSecMsg
func NewMiniProgramSecMsg(client *Client) *MiniProgramSecMsg {
	mpb := &MiniProgramSecMsg{
		client: client,
	}
	return mpb
}

// SetAccessToken SetAccessToken
func (mpb *MiniProgramSecMsg) SetAccessToken(accessToken string) *MiniProgramSecMsg {
	mpb.accessToken = accessToken
	return mpb
}

// SetMessage SetMessage
func (mpb *MiniProgramSecMsg) SetMessage(message string) *MiniProgramSecMsg {
	mpb.message = message
	return mpb
}

// Validate checks if the operation is valid.
func (mpb *MiniProgramSecMsg) Validate() error {
	var invalid []string
	if mpb.accessToken == "" {
		invalid = append(invalid, "access_token")
	}
	if len(mpb.message) == 0 {
		invalid = append(invalid, "media")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Do Do
func (mpb *MiniProgramSecMsg) Do(ctx context.Context) (*MiniProgramSecMsgResponse, error) {
	// Check pre-conditions
	if err := mpb.Validate(); err != nil {
		return nil, errors.Wrap(err, "MiniProgramSecMsg.Do")
	}
	params := url.Values{}
	params.Set("access_token", mpb.accessToken)
	bodybyte, err := json.Marshal(map[string]string{
		"content": mpb.message,
	})
	if err != nil {
		return nil, errors.Wrap(err, "MiniProgramSecMsg.Do")
	}
	res, err := mpb.client.PerformRequest(ctx, PerformRequestOptions{
		Method:   http.MethodPost,
		Params:   params,
		Body:     string(bodybyte),
		BaseURI:  MiniProgramBaseHost,
		Endpoint: MiniProgramSecMsgEndpoint,
	})
	if err != nil {
		return nil, errors.Wrap(err, "MiniProgramSecMsg.Do")
	}
	// Return operation response
	ret := new(MiniProgramSecMsgResponse)
	if err := mpb.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, errors.Wrap(err, "MiniProgramSecMsg.Do")
	}
	return ret, nil
}

// MiniProgramSecMsgResponse MiniProgramSecMsgResponse
type MiniProgramSecMsgResponse struct {
	CommonError
}
