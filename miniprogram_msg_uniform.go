package wechat

import (
	"errors"
	"net/url"
)

const (
	// MiniProgramUniformMessageEndpoint Endpoint
	MiniProgramUniformMessageEndpoint = "cgi-bin/message/wxopen/template/uniform_send"
)

// MiniProgramUniformMessage MiniProgramUniformMessage
type MiniProgramUniformMessage struct {
	MsgBody   *MiniProgramUniformMessageBody
	MsgParams url.Values
}

// MiniProgramUniformMessageBody MiniProgramUniformMessageBody
type MiniProgramUniformMessageBody struct {
	Touser           string           `json:"touser"`
	WeappTemplateMsg WeappTemplateMsg `json:"weapp_template_msg"` // 小程序模板消息相关的信息
	MpTemplateMsg    MpTemplateMsg    `json:"mp_template_msg"`    // 公众号模板消息相关的信息
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

// NewMiniProgramUniformMessage NewMiniProgramUniformMessage
func NewMiniProgramUniformMessage() *MiniProgramUniformMessage {
	return &MiniProgramUniformMessage{}
}

// Body Body
func (mpum *MiniProgramUniformMessage) Body() interface{} {
	return mpum.MsgBody
}

// Validate Validate
func (mpum *MiniProgramUniformMessage) Validate() error {
	if mpum.MsgBody == nil {
		return errors.New("body is nil")
	}
	if mpum.MsgParams == nil {
		mpum.MsgParams = url.Values{}
	}
	return nil
}

// BaseURI BaseURI
func (mpum *MiniProgramUniformMessage) BaseURI() string {
	return MiniProgramBaseURI
}

// Endpoint Endpoint
func (mpum *MiniProgramUniformMessage) Endpoint() string {
	return MiniProgramUniformMessageEndpoint
}

// Params Params
func (mpum *MiniProgramUniformMessage) Params() url.Values {
	return mpum.MsgParams
}
