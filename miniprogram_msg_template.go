package wechat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	// MiniProgramTemplateMessageEndpoint Endpoint
	MiniProgramTemplateMessageEndpoint = "cgi-bin/message/wxopen/template/send"
)

// MiniProgramTemplateMessage TemplateUniform
type MiniProgramTemplateMessage struct {
	client *Client

	accessToken string
	body        *MiniProgramTemplateMessageBody
}

// MiniProgramTemplateMessageBody MiniProgramTemplateMessageBody
type MiniProgramTemplateMessageBody struct {
	Touser     string `json:"touser"`
	TemplateID string `json:"template_id"`
	Page       string `json:"page"`
	FormID     string `json:"form_id"`
	Data       map[string]struct {
		Value string `json:"value"`
	} `json:"data"`
	EmphasisKeyword string `json:"emphasis_keyword"`
}

// NewMiniProgramTemplateMessage return instance of NewMiniProgramTemplateMessage
func NewMiniProgramTemplateMessage(client *Client) *MiniProgramTemplateMessage {
	mpt := &MiniProgramTemplateMessage{
		client: client,
	}
	return mpt
}

// SetAccessToken SetAccessToken
func (mpt *MiniProgramTemplateMessage) SetAccessToken(accessToken string) *MiniProgramTemplateMessage {
	mpt.accessToken = accessToken
	return mpt
}

// SetBody SetBody
func (mpt *MiniProgramTemplateMessage) SetBody(body *MiniProgramTemplateMessageBody) *MiniProgramTemplateMessage {
	mpt.body = body
	return mpt
}

// Validate checks if the operation is valid.
func (mpt *MiniProgramTemplateMessage) Validate() error {
	var invalid []string
	if mpt.accessToken == "" {
		invalid = append(invalid, "access_token")
	}
	if mpt.body == nil {
		invalid = append(invalid, "body")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Do Do
func (mpt *MiniProgramTemplateMessage) Do(ctx context.Context) (*MiniProgramTemplateMessageResponse, error) {
	// Check pre-conditions
	if err := mpt.Validate(); err != nil {
		return nil, err
	}
	// url params
	params := url.Values{}
	params.Set("access_token", mpt.accessToken)
	// PerformRequest
	res, err := mpt.client.PerformRequest(ctx, PerformRequestOptions{
		Method:   http.MethodPost,
		Params:   params,
		Body:     mpt.body,
		BaseURI:  MiniProgramBaseURI,
		Endpoint: MiniProgramTemplateMessageEndpoint,
	})
	if err != nil {
		return nil, err
	}
	// Return operation response
	ret := new(MiniProgramTemplateMessageResponse)
	if err := mpt.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// MiniProgramTemplateMessageResponse MiniProgramTemplateMessageResponse
type MiniProgramTemplateMessageResponse struct {
	CommonError
}
