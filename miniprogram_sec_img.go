package wechat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// Endpoint
const (
	MiniProgramSecImgEndpoint = "wxa/img_sec_check"
)

var (
	maxMediaSize = 1 << 20 // 1 MB
)

// MiniProgramSecImg MiniProgramSecImg
type MiniProgramSecImg struct {
	client *Client

	accessToken string
	media       string
}

// NewMiniProgramSecImg return instance of NewMiniProgramSecImg
func NewMiniProgramSecImg(client *Client) *MiniProgramSecImg {
	mpb := &MiniProgramSecImg{
		client: client,
	}
	return mpb
}

// SetAccessToken SetAccessToken
func (mpb *MiniProgramSecImg) SetAccessToken(accessToken string) *MiniProgramSecImg {
	mpb.accessToken = accessToken
	return mpb
}

// SetMedia SetMedia
func (mpb *MiniProgramSecImg) SetMedia(media string) *MiniProgramSecImg {
	mpb.media = media
	return mpb
}

// Validate checks if the operation is valid.
func (mpb *MiniProgramSecImg) Validate() error {
	var invalid []string
	if mpb.accessToken == "" {
		invalid = append(invalid, "access_token")
	}
	if len(mpb.media) > maxMediaSize || len(mpb.media) == 0 {
		invalid = append(invalid, "media")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	return nil
}

// Do Do
func (mpb *MiniProgramSecImg) Do(ctx context.Context) (*MiniProgramSecImgResponse, error) {
	// Check pre-conditions
	if err := mpb.Validate(); err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Set("access_token", mpb.accessToken)

	res, err := mpb.client.PerformRequest(ctx, PerformRequestOptions{
		Method:        http.MethodPost,
		Params:        params,
		FormValue:     []byte(mpb.media),
		FormFieldName: "media",
		BaseURI:       MiniProgramBaseURI,
		Endpoint:      MiniProgramSecImgEndpoint,
	})
	if err != nil {
		return nil, err
	}
	// Return operation response
	ret := new(MiniProgramSecImgResponse)
	if err := mpb.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// MiniProgramSecImgResponse MiniProgramSecImgResponse
type MiniProgramSecImgResponse struct {
	CommonError
}
