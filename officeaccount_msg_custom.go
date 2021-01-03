package wechat

import (
	"errors"
	"net/url"
)

const (
	// OACustomMessageEndpoint https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html
	OACustomMessageEndpoint = "cgi-bin/message/custom/send"
)

// OACustomMessage 实现 IBasicMessage 接口
type OACustomMessage struct {
	MsgBody   *OACustomMessageBody
	MsgParams url.Values
}

// OACustomMessageBody 消息体
type OACustomMessageBody struct {
	Touser          string                   `json:"touser"`
	Msgtype         string                   `json:"msgtype"`
	Miniprogrampage *OACustomMiniprogrampage `json:"miniprogrampage,omitempty"`
	Text            *OACustomText            `json:"text,omitempty"`
	Customservice   *OACustomservice         `json:"customservice,omitempty"`
	Wxcard          *OACustomWXCard          `json:"wxcard,omitempty"`
	Msgmenu         *OACustomMsgmenu         `json:"msgmenu,omitempty"`
	Mpnews          *OACustomMpnews          `json:"mpnews,omitempty"`
	News            *OACustomNews            `json:"news,omitempty"`
	Music           *OACustomMusic           `json:"music,omitempty"`
	Voice           *OACustomVoice           `json:"voice,omitempty"`
	Video           *OACustomVideo           `json:"video,omitempty"`
	Image           *OACustomImage           `json:"image,omitempty"`
}

type OACustomImage struct {
	MediaID string `json:"media_id"`
}

type OACustomVideo struct {
	MediaID      string `json:"media_id"`
	ThumbMediaID string `json:"thumb_media_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

type OACustomVoice struct {
	MediaID string `json:"media_id"`
}

type OACustomMusic struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Musicurl     string `json:"musicurl"`
	Hqmusicurl   string `json:"hqmusicurl"`
	ThumbMediaID string `json:"thumb_media_id"`
}

type OACustomNews struct {
	Articles []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Picurl      string `json:"picurl"`
	} `json:"articles,omitempty"`
}

type OACustomMpnews struct {
	MediaID string `json:"media_id"`
}

type OACustomMsgmenu struct {
	HeadContent string `json:"head_content"`
	List        []struct {
		ID      string `json:"id"`
		Content string `json:"content"`
	} `json:"list,omitempty"`
	TailContent string `json:"tail_content"`
}

type OACustomWXCard struct {
	CardID string `json:"card_id"`
}

type OACustomservice struct {
	KfAccount string `json:"kf_account"`
}

type OACustomText struct {
	Content string `json:"content"`
}

type OACustomMiniprogrampage struct {
	Title        string `json:"title"`
	Appid        string `json:"appid"`
	Pagepath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}

// NewOACustomMessage 订阅消息
func NewOACustomMessage(sm *OACustomMessage) *OACustomMessage {
	return sm
}

// Body Body
func (mpum *OACustomMessage) Body() interface{} {
	return mpum.MsgBody
}

// Validate Validate
func (mpum *OACustomMessage) Validate() error {
	if mpum.MsgBody == nil {
		return errors.New("body is nil")
	}
	if mpum.MsgBody.Touser == "" {
		return errors.New("接收人 openid 为空")
	}
	if mpum.MsgParams == nil {
		mpum.MsgParams = url.Values{}
	}
	return nil
}

// BaseURI BaseURI
func (mpum *OACustomMessage) BaseURI() string {
	return MiniProgramBaseHost
}

// Endpoint Endpoint
func (mpum *OACustomMessage) Endpoint() string {
	return OACustomMessageEndpoint
}

// Params Params
func (mpum *OACustomMessage) Params() url.Values {
	return mpum.MsgParams
}
