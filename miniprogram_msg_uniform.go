package wechat

import (
	"errors"
	"net/url"
)

const (
	// MiniProgramUniformMessageEndpoint ..
	// 这个接口的初衷就是想大家在开发小程序的时候，如果要发送公众号消息直接使用这个接口就可以了，无需再去调用公众号的模板消息接口
	MiniProgramUniformMessageEndpoint = "cgi-bin/message/wxopen/template/uniform_send"
)

// MiniProgramUniformMessage MiniProgramUniformMessage
type MiniProgramUniformMessage struct {
	MsgBody   *MiniProgramUniformMessageBody
	MsgParams url.Values
}

// MiniProgramUniformMessageBody MiniProgramUniformMessageBody
type MiniProgramUniformMessageBody struct {
	// 用户openid，可以是小程序的openid，也可以是mp_template_msg.appid对应的公众号的openid
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
	return MiniProgramBaseHost
}

// Endpoint Endpoint
func (mpum *MiniProgramUniformMessage) Endpoint() string {
	return MiniProgramUniformMessageEndpoint
}

// Params Params
func (mpum *MiniProgramUniformMessage) Params() url.Values {
	return mpum.MsgParams
}
