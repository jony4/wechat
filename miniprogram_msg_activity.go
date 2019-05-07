package wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// Endpoint
const (
	MiniProgramActivityMessageCreateEndpoint = "cgi-bin/message/wxopen/activityid/create"
	MiniProgramActivityMessageUpdateEndpoint = "cgi-bin/message/wxopen/updatablemsg/send"
)

// MiniProgramActivityMessageCreate MiniProgramActivityMessageCreateCreate
type MiniProgramActivityMessageCreate struct {
	client *Client

	accessToken string
}

// NewMiniProgramActivityMessageCreate return instance of mini program auth
func NewMiniProgramActivityMessageCreate(client *Client) *MiniProgramActivityMessageCreate {
	mpam := &MiniProgramActivityMessageCreate{
		client: client,
	}
	return mpam
}

// SetAccessToken SetAccessToken
func (mpam *MiniProgramActivityMessageCreate) SetAccessToken(accessToken string) *MiniProgramActivityMessageCreate {
	mpam.accessToken = accessToken
	return mpam
}

// Validate checks if the operation is valid.
func (mpam *MiniProgramActivityMessageCreate) Validate() error {
	var invalid []string
	if mpam.accessToken == "" {
		invalid = append(invalid, "AccessToken")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Do Do
func (mpam *MiniProgramActivityMessageCreate) Do(ctx context.Context) (*MiniProgramActivityMessageCreateResponse, error) {
	// Check pre-conditions
	if err := mpam.Validate(); err != nil {
		return nil, errors.Wrap(err, "MiniProgramActivityMessageCreate.Do")
	}
	// url params
	params := url.Values{}
	params.Set("access_token", mpam.accessToken)
	// PerformRequest
	res, err := mpam.client.PerformRequest(ctx, PerformRequestOptions{
		Method:   http.MethodGet,
		Params:   params,
		BaseURI:  MiniProgramBaseURI,
		Endpoint: MiniProgramActivityMessageCreateEndpoint,
	})
	if err != nil {
		return nil, errors.Wrap(err, "MiniProgramActivityMessageCreate.Do")
	}
	// Return operation response
	ret := new(MiniProgramActivityMessageCreateResponse)
	if err := mpam.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, errors.Wrap(err, "MiniProgramActivityMessageCreate.Do")
	}
	return ret, nil
}

// MiniProgramActivityMessageCreateResponse MiniProgramActivityMessageCreateResponse
type MiniProgramActivityMessageCreateResponse struct {
	CommonError
	ActivityID     string `json:"activity_id"`
	ExpirationTime string `json:"expiration_time"`
}

// MiniProgramActivityMessageUpdate MiniProgramActivityMessageUpdateCreate
type MiniProgramActivityMessageUpdate struct {
	client *Client

	accessToken string
	body        *MiniProgramActivityMessageUpdateBody
}

var (
	allowedTemplateName = map[string]bool{
		"member_count": true,
		"room_limit":   true,
		"path":         true,
		"version_type": true,
	}
	allowedVersionType = map[string]bool{
		"develop": true,
		"trial":   true,
		"release": true,
	}
)

// MiniProgramActivityMessageUpdateBody MiniProgramActivityMessageUpdateBody
type MiniProgramActivityMessageUpdateBody struct {
	ActivityID   string                      `json:"activity_id"`
	TargetState  int64                       `json:"target_state"`
	TemplateInfo *MPAMUpdateBodyTemplateInfo `json:"template_info"`
}

// MPAMUpdateBodyTemplateInfo MPAMUpdateBodyTemplateInfo
type MPAMUpdateBodyTemplateInfo struct {
	ParameterList []*MPAMUpdateBodyParameterList `json:"parameter_list"`
}

// MPAMUpdateBodyParameterList MPAMUpdateBodyParameterList
type MPAMUpdateBodyParameterList struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Validate Validate
func (mpamub *MiniProgramActivityMessageUpdateBody) Validate() error {
	if mpamub.TemplateInfo == nil {
		return fmt.Errorf("missing required fields: %v", "parameter_list")
	}
	var invalid []string
	if mpamub.ActivityID == "" {
		invalid = append(invalid, "activity_id")
	}
	if mpamub.TargetState != 1 && mpamub.TargetState != 0 {
		invalid = append(invalid, "target_state")
	}
	for _, v := range mpamub.TemplateInfo.ParameterList {
		if _, ok := allowedTemplateName[v.Name]; !ok {
			continue
		}
		if v.Name == "version_type" {
			if _, ok := allowedVersionType[v.Value]; !ok {
				return fmt.Errorf("not allowed Version Type")
			}
		}
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// NewMiniProgramActivityMessageUpdate return instance of mini program auth
func NewMiniProgramActivityMessageUpdate(client *Client) *MiniProgramActivityMessageUpdate {
	mpamu := &MiniProgramActivityMessageUpdate{
		client: client,
	}
	return mpamu
}

// SetAccessToken SetAccessToken
func (mpamu *MiniProgramActivityMessageUpdate) SetAccessToken(accessToken string) *MiniProgramActivityMessageUpdate {
	mpamu.accessToken = accessToken
	return mpamu
}

// SetBody SetBody
func (mpamu *MiniProgramActivityMessageUpdate) SetBody(body *MiniProgramActivityMessageUpdateBody) *MiniProgramActivityMessageUpdate {
	mpamu.body = body
	return mpamu
}

// Validate checks if the operation is valid.
func (mpamu *MiniProgramActivityMessageUpdate) Validate() error {
	var invalid []string
	if mpamu.body == nil {
		return fmt.Errorf("missing required fields: %v", "Body")
	}
	if err := mpamu.body.Validate(); err != nil {
		invalid = append(invalid, err.Error())
	}
	if mpamu.accessToken == "" {
		invalid = append(invalid, "AccessToken")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Do Do
func (mpamu *MiniProgramActivityMessageUpdate) Do(ctx context.Context) (*MiniProgramActivityMessageUpdateResponse, error) {
	// Check pre-conditions
	if err := mpamu.Validate(); err != nil {
		return nil, errors.Wrap(err, "MiniProgramActivityMessageUpdate.Do")
	}
	bodybyte, err := json.Marshal(mpamu.body)
	if err != nil {
		return nil, errors.Wrap(err, "MiniProgramActivityMessageUpdate.Do")
	}
	// url params
	params := url.Values{}
	params.Set("access_token", mpamu.accessToken)
	// PerformRequest
	res, err := mpamu.client.PerformRequest(ctx, PerformRequestOptions{
		Method:   http.MethodPost,
		Params:   params,
		Body:     string(bodybyte),
		BaseURI:  MiniProgramBaseURI,
		Endpoint: MiniProgramActivityMessageUpdateEndpoint,
	})
	if err != nil {
		return nil, errors.Wrap(err, "MiniProgramActivityMessageUpdate.Do")
	}
	// Return operation response
	ret := new(MiniProgramActivityMessageUpdateResponse)
	if err := mpamu.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, errors.Wrap(err, "MiniProgramActivityMessageUpdate.Do")
	}
	return ret, nil
}

// MiniProgramActivityMessageUpdateResponse MiniProgramActivityMessageUpdateResponse
type MiniProgramActivityMessageUpdateResponse struct {
	CommonError
}
