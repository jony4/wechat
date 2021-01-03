package wechat

import (
	"errors"
	"net/url"
)

const (
	// MPSubscribeMessageEndpoint https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
	MPSubscribeMessageEndpoint = "cgi-bin/message/subscribe/send"
)

// MPSubscribeMessage 实现 IBasicMessage 接口
type MPSubscribeMessage struct {
	MsgBody   *MPSubscribeMessageBody
	MsgParams url.Values
}

// MPSubscribeMessageBody 消息体
type MPSubscribeMessageBody struct {
	ToUser     string `json:"touser"`
	TemplateID string `json:"template_id"`
	Page       string `json:"page"`
	Data       map[string]struct {
		Value string `json:"value"`
	} `json:"data"`
	MiniprogramState string `json:"miniprogram_state"`
	Lang             string `json:"lang"`
}

// NewMPSubscribeMessage 订阅消息
func NewMPSubscribeMessage(sm *MPSubscribeMessage) *MPSubscribeMessage {
	return sm
}

// Body Body
func (mpum *MPSubscribeMessage) Body() interface{} {
	return mpum.MsgBody
}

// Validate Validate
func (mpum *MPSubscribeMessage) Validate() error {
	if mpum.MsgBody == nil {
		return errors.New("body is nil")
	}
	if mpum.MsgBody.ToUser == "" {
		return errors.New("接收人 openid 为空")
	}
	if mpum.MsgBody.TemplateID == "" {
		return errors.New("模板 id 为空")
	}
	if mpum.MsgParams == nil {
		mpum.MsgParams = url.Values{}
	}
	return nil
}

// BaseURI BaseURI
func (mpum *MPSubscribeMessage) BaseURI() string {
	return MiniProgramBaseHost
}

// Endpoint Endpoint
func (mpum *MPSubscribeMessage) Endpoint() string {
	return MPSubscribeMessageEndpoint
}

// Params Params
func (mpum *MPSubscribeMessage) Params() url.Values {
	return mpum.MsgParams
}
