package wechat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	// MiniProgramTemplateUniformMessageEndpoint Endpoint
	MiniProgramTemplateUniformMessageEndpoint = "cgi-bin/message/wxopen/template/uniform_send"
)

// MiniProgramTemplateUniformMessage TemplateUniform
type MiniProgramTemplateUniformMessage struct {
	client *Client

	accessToken string
	body        *MiniProgramTemplateUniformMessageBody
}

// MiniProgramTemplateUniformMessageBody MiniProgramTemplateUniformMessageBody
type MiniProgramTemplateUniformMessageBody struct {
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

// NewMiniProgramTemplateUniformMessage return instance of NewMiniProgramTemplateUniformMessage
func NewMiniProgramTemplateUniformMessage(client *Client) *MiniProgramTemplateUniformMessage {
	mptu := &MiniProgramTemplateUniformMessage{
		client: client,
	}
	return mptu
}

// SetAccessToken SetAccessToken
func (mptu *MiniProgramTemplateUniformMessage) SetAccessToken(accessToken string) *MiniProgramTemplateUniformMessage {
	mptu.accessToken = accessToken
	return mptu
}

// SetBody SetBody
func (mptu *MiniProgramTemplateUniformMessage) SetBody(body *MiniProgramTemplateUniformMessageBody) *MiniProgramTemplateUniformMessage {
	mptu.body = body
	return mptu
}

// Validate checks if the operation is valid.
func (mptu *MiniProgramTemplateUniformMessage) Validate() error {
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
func (mptu *MiniProgramTemplateUniformMessage) Do(ctx context.Context) (*MiniProgramTemplateUniformMessageResponse, error) {
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
		Endpoint: MiniProgramTemplateUniformMessageEndpoint,
	})
	if err != nil {
		return nil, err
	}
	// Return operation response
	ret := new(MiniProgramTemplateUniformMessageResponse)
	if err := mptu.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// MiniProgramTemplateUniformMessageResponse MiniProgramTemplateUniformMessageResponse
type MiniProgramTemplateUniformMessageResponse struct {
	CommonError
}
