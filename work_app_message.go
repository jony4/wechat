package wechat

import (
	"errors"
	"net/url"
)

const (
	// WorkAppMessageEndpoint Endpoint
	WorkAppMessageEndpoint = "cgi-bin/message/send"
)

// WorkAppMessage WorkAppMessage
type WorkAppMessage struct {
	MsgBody   *WorkAppMessageBody
	MsgParams url.Values
}

// WorkAppMessageBody WorkAppMessageBody
type WorkAppMessageBody struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Safe    int    `json:"safe,omitempty"`

	Textcard          *Textcard          `json:"textcard,omitempty"`
	Text              *Text              `json:"text,omitempty"`
	Image             *Image             `json:"image,omitempty"`
	Voice             *Voice             `json:"voice,omitempty"`
	Video             *Video             `json:"video,omitempty"`
	File              *File              `json:"file,omitempty"`
	News              *News              `json:"news,omitempty"`
	Mpnews            *Mpnews            `json:"mpnews,omitempty"`
	Markdown          *Markdown          `json:"markdown,omitempty"`
	MiniprogramNotice *MiniprogramNotice `json:"miniprogram_notice,omitempty"`
	Taskcard          *Taskcard          `json:"taskcard,omitempty"`
}

// NewWorkAppMessage NewWorkAppMessage
func NewWorkAppMessage() *WorkAppMessage {
	return &WorkAppMessage{}
}

// Body Body
func (wam *WorkAppMessage) Body() interface{} {
	return wam.MsgBody
}

// Validate Validate
func (wam *WorkAppMessage) Validate() error {
	if wam.MsgBody == nil {
		return errors.New("body is nil")
	}
	if wam.MsgParams == nil {
		wam.MsgParams = url.Values{}
	}
	return nil
}

// BaseURI BaseURI
func (wam *WorkAppMessage) BaseURI() string {
	return WorkBaseHost
}

// Endpoint Endpoint
func (wam *WorkAppMessage) Endpoint() string {
	return WorkAppMessageEndpoint
}

// Params Params
func (wam *WorkAppMessage) Params() url.Values {
	return wam.MsgParams
}

// Textcard Textcard
type Textcard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Btntxt      string `json:"btntxt"`
}

// Text Text
type Text struct {
	Content string `json:"content"`
}

// Image Image
type Image struct {
	MediaID string `json:"media_id"`
}

// Voice Voice
type Voice struct {
	MediaID string `json:"media_id"`
}

//Video Video
type Video struct {
	MediaID     string `json:"media_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// File File
type File struct {
	MediaID string `json:"media_id"`
}

// News News
type News struct {
	Articles []Article `json:"articles"`
}

// Article Article
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Picurl      string `json:"picurl"`
}

// Mpnews Mpnews
type Mpnews struct {
	Articles []MpArticle `json:"articles"`
}

// MpArticle MpArticle
type MpArticle struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	ContentSourceURL string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
}

// Markdown Markdown
type Markdown struct {
	Content string `json:"content"`
}

// MiniprogramNotice MiniprogramNotice
type MiniprogramNotice struct {
	Appid             string                   `json:"appid"`
	Page              string                   `json:"page"`
	Title             string                   `json:"title"`
	Description       string                   `json:"description"`
	EmphasisFirstItem bool                     `json:"emphasis_first_item"`
	ContentItem       []MiniprogramContentItem `json:"content_item"`
}

// MiniprogramContentItem MiniprogramContentItem
type MiniprogramContentItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Taskcard Taskcard
type Taskcard struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	URL         string        `json:"url"`
	TaskID      string        `json:"task_id"`
	Btn         []TaskcardBtn `json:"btn"`
}

// TaskcardBtn TaskcardBtn
type TaskcardBtn struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	ReplaceName string `json:"replace_name"`
	Color       string `json:"color,omitempty"`
	IsBold      bool   `json:"is_bold,omitempty"`
}
