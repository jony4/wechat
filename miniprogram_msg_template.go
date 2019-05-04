package wechat

import (
	"errors"
	"net/url"
)

const (
	// MiniProgramTemplateMessageEndpoint Endpoint
	MiniProgramTemplateMessageEndpoint = "cgi-bin/message/wxopen/template/send"
)

// MiniProgramTemplateMessage TemplateUniform
type MiniProgramTemplateMessage struct {
	MsgBody   *MiniProgramUniformMessageBody
	MsgParams url.Values
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

// NewMiniProgramTemplateMessage NewMiniProgramTemplateMessage
func NewMiniProgramTemplateMessage() *MiniProgramTemplateMessage {
	return &MiniProgramTemplateMessage{}
}

// Body Body
func (mptm *MiniProgramTemplateMessage) Body() interface{} {
	return mptm.MsgBody
}

// Validate Validate
func (mptm *MiniProgramTemplateMessage) Validate() error {
	if mptm.MsgBody == nil {
		return errors.New("body is nil")
	}
	if mptm.MsgParams == nil {
		mptm.MsgParams = url.Values{}
	}
	return nil
}

// BaseURI BaseURI
func (mptm *MiniProgramTemplateMessage) BaseURI() string {
	return MiniProgramBaseURI
}

// Endpoint Endpoint
func (mptm *MiniProgramTemplateMessage) Endpoint() string {
	return MiniProgramTemplateMessageEndpoint
}

// Params Params
func (mptm *MiniProgramTemplateMessage) Params() url.Values {
	return mptm.MsgParams
}
