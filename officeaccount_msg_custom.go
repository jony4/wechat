package wechat

import (
	"errors"
	"net/url"
)

const (
	// OACustomMessageEndpoint https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
	OACustomMessageEndpoint = "cgi-bin/message/custom/send"
)

// OACustomMessage 实现 IBasicMessage 接口
type OACustomMessage struct {
	MsgBody   *OACustomMessageBody
	MsgParams url.Values
}

// OACustomMessageBody 消息体
type OACustomMessageBody struct {
	Touser          string `json:"touser"`
	Msgtype         string `json:"msgtype"`
	Miniprogrampage struct {
		Title        string `json:"title"`
		Appid        string `json:"appid"`
		Pagepath     string `json:"pagepath"`
		ThumbMediaID string `json:"thumb_media_id"`
	} `json:"miniprogrampage"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Customservice struct {
		KfAccount string `json:"kf_account"`
	} `json:"customservice"`
	Wxcard struct {
		CardID string `json:"card_id"`
	} `json:"wxcard"`
	Msgmenu struct {
		HeadContent string `json:"head_content"`
		List        []struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		} `json:"list"`
		TailContent string `json:"tail_content"`
	} `json:"msgmenu"`
	Mpnews struct {
		MediaID string `json:"media_id"`
	} `json:"mpnews"`
	News struct {
		Articles []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			URL         string `json:"url"`
			Picurl      string `json:"picurl"`
		} `json:"articles"`
	} `json:"news"`
	Music struct {
		Title        string `json:"title"`
		Description  string `json:"description"`
		Musicurl     string `json:"musicurl"`
		Hqmusicurl   string `json:"hqmusicurl"`
		ThumbMediaID string `json:"thumb_media_id"`
	} `json:"music"`
	Voice struct {
		MediaID string `json:"media_id"`
	} `json:"voice"`
	Video struct {
		MediaID      string `json:"media_id"`
		ThumbMediaID string `json:"thumb_media_id"`
		Title        string `json:"title"`
		Description  string `json:"description"`
	} `json:"video"`
	Image struct {
		MediaID string `json:"media_id"`
	} `json:"image"`
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
