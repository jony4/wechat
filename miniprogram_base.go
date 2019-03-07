package wechat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	// MiniProgramBaseEndpoint Endpoint
	MiniProgramBaseEndpoint = "wxa/getpaidunionid"
)

// MiniProgramBase mini program auth.
type MiniProgramBase struct {
	client *Client

	accessToken string
	appid       string

	usingTransID bool

	transactionID string

	mchID      string
	outTradeNo string
}

// NewMiniProgramBase return instance of mini program auth
func NewMiniProgramBase(client *Client) *MiniProgramBase {
	mpb := &MiniProgramBase{
		client: client,
	}
	return mpb
}

// SetAppID SetAppID
func (mpb *MiniProgramBase) SetAppID(appid string) *MiniProgramBase {
	mpb.appid = appid
	return mpb
}

// SetAccessToken SetAccessToken
func (mpb *MiniProgramBase) SetAccessToken(accessToken string) *MiniProgramBase {
	mpb.accessToken = accessToken
	return mpb
}

// SetTransactionID SetTransactionID
func (mpb *MiniProgramBase) SetTransactionID(transactionID string) *MiniProgramBase {
	mpb.transactionID = transactionID
	return mpb
}

// SetMchID SetMchID
func (mpb *MiniProgramBase) SetMchID(mchID string) *MiniProgramBase {
	mpb.mchID = mchID
	return mpb
}

// SetOutTradeNo SetOutTradeNo
func (mpb *MiniProgramBase) SetOutTradeNo(outTradeNo string) *MiniProgramBase {
	mpb.outTradeNo = outTradeNo
	return mpb
}

// SetUsingTransID SetUsingTransID
func (mpb *MiniProgramBase) SetUsingTransID() *MiniProgramBase {
	mpb.usingTransID = true
	return mpb
}

// Validate checks if the operation is valid.
func (mpb *MiniProgramBase) Validate() error {
	var invalid []string
	if mpb.appid == "" {
		invalid = append(invalid, "AppID")
	}
	if mpb.accessToken == "" {
		invalid = append(invalid, "AccessToken")
	}
	if len(invalid) > 0 {
		return fmt.Errorf("missing required fields: %v", invalid)
	}
	if mpb.usingTransID && mpb.transactionID == "" {
		return fmt.Errorf("transaction_id is nil")
	}
	if !mpb.usingTransID && (mpb.mchID == "" || mpb.outTradeNo == "") {
		return fmt.Errorf("mch_id or out_trade_no is nil")
	}
	return nil
}

// Do Do
func (mpb *MiniProgramBase) Do(ctx context.Context) (*MiniProgramBaseResponse, error) {
	// Check pre-conditions
	if err := mpb.Validate(); err != nil {
		return nil, err
	}
	// url params
	params := url.Values{}
	params.Set("appid", mpb.appid)
	params.Set("access_token", mpb.accessToken)
	if mpb.usingTransID {
		params.Set("transaction_id", mpb.transactionID)
	} else {
		params.Set("mch_id", mpb.mchID)
		params.Set("out_trade_no", mpb.outTradeNo)
	}
	// PerformRequest
	res, err := mpb.client.PerformRequest(ctx, PerformRequestOptions{
		Method:   http.MethodGet,
		Params:   params,
		BaseURI:  MiniProgramBaseURI,
		Endpoint: MiniProgramBaseEndpoint,
	})
	if err != nil {
		return nil, err
	}
	// Return operation response
	ret := new(MiniProgramBaseResponse)
	if err := mpb.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// MiniProgramBaseResponse MiniProgramBaseResponse
type MiniProgramBaseResponse struct {
	CommonError
	UnionID string `json:"unionid"`
}
