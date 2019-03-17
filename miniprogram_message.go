package wechat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	// MiniProgramTemplateUniformEndpoint Endpoint
	MiniProgramTemplateUniformEndpoint = "cgi-bin/message/wxopen/template/uniform_send"
)

// MiniProgramTemplateUniform TemplateUniform
type MiniProgramTemplateUniform struct {
	client *Client

	accessToken string
	body        *MiniProgramTemplateUniformBody
}

// MiniProgramTemplateUniformBody MiniProgramTemplateUniformBody
type MiniProgramTemplateUniformBody struct {
	Touser           string           `json:"touser"`
	WeappTemplateMsg WeappTemplateMsg `json:"weapp_template_msg"`
	MpTemplateMsg    MpTemplateMsg    `json:"mp_template_msg"`
}

// WeappTemplateMsg WeappTemplateMsg
type WeappTemplateMsg struct {
	TemplateID string `json:"template_id"`
	Page       string `json:"page"`
	FormID     string `json:"form_id"`
	Data       map[string]struct {
		Value string `json:"value"`
	} `json:"data"`
	EmphasisKeyword string `json:"emphasis_keyword"`
}

// MpTemplateMsg MpTemplateMsg
type MpTemplateMsg struct {
	Appid       string `json:"appid"`
	TemplateID  string `json:"template_id"`
	URL         string `json:"url"`
	Miniprogram struct {
		Appid    string `json:"appid"`
		Pagepath string `json:"pagepath"`
	} `json:"miniprogram"`
	Data map[string]map[string]string `json:"data"`
}

// NewMiniProgramTemplateUniform return instance of NewMiniProgramTemplateUniform
func NewMiniProgramTemplateUniform(client *Client) *MiniProgramTemplateUniform {
	mptu := &MiniProgramTemplateUniform{
		client: client,
	}
	return mptu
}

// SetAccessToken SetAccessToken
func (mptu *MiniProgramTemplateUniform) SetAccessToken(accessToken string) *MiniProgramTemplateUniform {
	mptu.accessToken = accessToken
	return mptu
}

// SetBody SetBody
func (mptu *MiniProgramTemplateUniform) SetBody(body *MiniProgramTemplateUniformBody) *MiniProgramTemplateUniform {
	mptu.body = body
	return mptu
}

// Validate checks if the operation is valid.
func (mptu *MiniProgramTemplateUniform) Validate() error {
	var invalid []string
	if mptu.accessToken == "" {
		invalid = append(invalid, "access_token")
	}
	if mptu.body == nil {
		invalid = append(invalid, "body")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Do Do
func (mptu *MiniProgramTemplateUniform) Do(ctx context.Context) (*MiniProgramTemplateUniformResponse, error) {
	// Check pre-conditions
	if err := mptu.Validate(); err != nil {
		return nil, err
	}
	// url params
	params := url.Values{}
	params.Set("access_token", mptu.accessToken)
	// PerformRequest
	res, err := mptu.client.PerformRequest(ctx, PerformRequestOptions{
		Method:   http.MethodPost,
		Params:   params,
		Body:     mptu.body,
		BaseURI:  MiniProgramBaseURI,
		Endpoint: MiniProgramTemplateUniformEndpoint,
	})
	if err != nil {
		return nil, err
	}
	// Return operation response
	ret := new(MiniProgramTemplateUniformResponse)
	if err := mptu.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// MiniProgramTemplateUniformResponse MiniProgramTemplateUniformResponse
type MiniProgramTemplateUniformResponse struct {
	CommonError
}
