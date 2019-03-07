package wechat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// Endpoint
const (
	MiniProgramAppCodeGetEndpoint        = "wxa/getwxacode"
	MiniProgramAppCodeGetUnlimitEndpoint = "wxa/getwxacodeunlimit"
	MiniProgramAppCodeCreateEndpoint     = "cgi-bin/wxaapp/createwxaqrcode"
)

// MiniProgramAppCodeGet MiniProgramAppCodeGet
type MiniProgramAppCodeGet struct {
	client *Client

	accessToken   string
	openid        string
	usingTransID  bool
	transactionID string
	mchID         string
	outTradeNo    string
}

// NewMiniProgramAppCodeGet return instance of NewMiniProgramAppCodeGet
func NewMiniProgramAppCodeGet(client *Client) *MiniProgramAppCodeGet {
	mpb := &MiniProgramAppCodeGet{
		client: client,
	}
	return mpb
}

// SetOpenID SetOpenID
func (mpb *MiniProgramAppCodeGet) SetOpenID(openid string) *MiniProgramAppCodeGet {
	mpb.openid = openid
	return mpb
}

// SetAccessToken SetAccessToken
func (mpb *MiniProgramAppCodeGet) SetAccessToken(accessToken string) *MiniProgramAppCodeGet {
	mpb.accessToken = accessToken
	return mpb
}

// SetTransactionID SetTransactionID
func (mpb *MiniProgramAppCodeGet) SetTransactionID(transactionID string) *MiniProgramAppCodeGet {
	mpb.transactionID = transactionID
	return mpb
}

// SetMchID SetMchID
func (mpb *MiniProgramAppCodeGet) SetMchID(mchID string) *MiniProgramAppCodeGet {
	mpb.mchID = mchID
	return mpb
}

// SetOutTradeNo SetOutTradeNo
func (mpb *MiniProgramAppCodeGet) SetOutTradeNo(outTradeNo string) *MiniProgramAppCodeGet {
	mpb.outTradeNo = outTradeNo
	return mpb
}

// SetUsingTransID SetUsingTransID
func (mpb *MiniProgramAppCodeGet) SetUsingTransID() *MiniProgramAppCodeGet {
	mpb.usingTransID = true
	return mpb
}

// Validate checks if the operation is valid.
func (mpb *MiniProgramAppCodeGet) Validate() error {
	var invalid []string
	if mpb.openid == "" {
		invalid = append(invalid, "openid")
	}
	if mpb.accessToken == "" {
		invalid = append(invalid, "access_token")
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
func (mpb *MiniProgramAppCodeGet) Do(ctx context.Context) (*MiniProgramAppCodeGetResponse, error) {
	// Check pre-conditions
	if err := mpb.Validate(); err != nil {
		return nil, err
	}
	// url params
	params := url.Values{}
	params.Set("openid", mpb.openid)
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
		Endpoint: MiniProgramAppCodeGetEndpoint,
	})
	if err != nil {
		return nil, err
	}
	// Return operation response
	ret := new(MiniProgramAppCodeGetResponse)
	if err := mpb.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// MiniProgramAppCodeGetResponse MiniProgramAppCodeGetResponse
type MiniProgramAppCodeGetResponse struct {
	CommonError
	UnionID string `json:"unionid"`
}

// MiniProgramAppCodeGetUnlimit MiniProgramAppCodeGetUnlimit
type MiniProgramAppCodeGetUnlimit struct {
	client *Client

	accessToken   string
	openid        string
	usingTransID  bool
	transactionID string
	mchID         string
	outTradeNo    string
}

// NewMiniProgramAppCodeGetUnlimit return instance of NewMiniProgramAppCodeGetUnlimit
func NewMiniProgramAppCodeGetUnlimit(client *Client) *MiniProgramAppCodeGetUnlimit {
	mpb := &MiniProgramAppCodeGetUnlimit{
		client: client,
	}
	return mpb
}

// SetOpenID SetOpenID
func (mpb *MiniProgramAppCodeGetUnlimit) SetOpenID(openid string) *MiniProgramAppCodeGetUnlimit {
	mpb.openid = openid
	return mpb
}

// SetAccessToken SetAccessToken
func (mpb *MiniProgramAppCodeGetUnlimit) SetAccessToken(accessToken string) *MiniProgramAppCodeGetUnlimit {
	mpb.accessToken = accessToken
	return mpb
}

// SetTransactionID SetTransactionID
func (mpb *MiniProgramAppCodeGetUnlimit) SetTransactionID(transactionID string) *MiniProgramAppCodeGetUnlimit {
	mpb.transactionID = transactionID
	return mpb
}

// SetMchID SetMchID
func (mpb *MiniProgramAppCodeGetUnlimit) SetMchID(mchID string) *MiniProgramAppCodeGetUnlimit {
	mpb.mchID = mchID
	return mpb
}

// SetOutTradeNo SetOutTradeNo
func (mpb *MiniProgramAppCodeGetUnlimit) SetOutTradeNo(outTradeNo string) *MiniProgramAppCodeGetUnlimit {
	mpb.outTradeNo = outTradeNo
	return mpb
}

// SetUsingTransID SetUsingTransID
func (mpb *MiniProgramAppCodeGetUnlimit) SetUsingTransID() *MiniProgramAppCodeGetUnlimit {
	mpb.usingTransID = true
	return mpb
}

// Validate checks if the operation is valid.
func (mpb *MiniProgramAppCodeGetUnlimit) Validate() error {
	var invalid []string
	if mpb.openid == "" {
		invalid = append(invalid, "openid")
	}
	if mpb.accessToken == "" {
		invalid = append(invalid, "access_token")
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
func (mpb *MiniProgramAppCodeGetUnlimit) Do(ctx context.Context) (*MiniProgramAppCodeGetUnlimitResponse, error) {
	// Check pre-conditions
	if err := mpb.Validate(); err != nil {
		return nil, err
	}
	// url params
	params := url.Values{}
	params.Set("openid", mpb.openid)
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
		Endpoint: MiniProgramAppCodeGetUnlimitEndpoint,
	})
	if err != nil {
		return nil, err
	}
	// Return operation response
	ret := new(MiniProgramAppCodeGetUnlimitResponse)
	if err := mpb.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// MiniProgramAppCodeGetUnlimitResponse MiniProgramAppCodeGetUnlimitResponse
type MiniProgramAppCodeGetUnlimitResponse struct {
	CommonError
	UnionID string `json:"unionid"`
}

// MiniProgramAppCodeCreate MiniProgramAppCodeCreate
type MiniProgramAppCodeCreate struct {
	client *Client

	accessToken   string
	openid        string
	usingTransID  bool
	transactionID string
	mchID         string
	outTradeNo    string
}

// NewMiniProgramAppCodeCreate return instance of NewMiniProgramAppCodeCreate
func NewMiniProgramAppCodeCreate(client *Client) *MiniProgramAppCodeCreate {
	mpb := &MiniProgramAppCodeCreate{
		client: client,
	}
	return mpb
}

// SetOpenID SetOpenID
func (mpb *MiniProgramAppCodeCreate) SetOpenID(openid string) *MiniProgramAppCodeCreate {
	mpb.openid = openid
	return mpb
}

// SetAccessToken SetAccessToken
func (mpb *MiniProgramAppCodeCreate) SetAccessToken(accessToken string) *MiniProgramAppCodeCreate {
	mpb.accessToken = accessToken
	return mpb
}

// SetTransactionID SetTransactionID
func (mpb *MiniProgramAppCodeCreate) SetTransactionID(transactionID string) *MiniProgramAppCodeCreate {
	mpb.transactionID = transactionID
	return mpb
}

// SetMchID SetMchID
func (mpb *MiniProgramAppCodeCreate) SetMchID(mchID string) *MiniProgramAppCodeCreate {
	mpb.mchID = mchID
	return mpb
}

// SetOutTradeNo SetOutTradeNo
func (mpb *MiniProgramAppCodeCreate) SetOutTradeNo(outTradeNo string) *MiniProgramAppCodeCreate {
	mpb.outTradeNo = outTradeNo
	return mpb
}

// SetUsingTransID SetUsingTransID
func (mpb *MiniProgramAppCodeCreate) SetUsingTransID() *MiniProgramAppCodeCreate {
	mpb.usingTransID = true
	return mpb
}

// Validate checks if the operation is valid.
func (mpb *MiniProgramAppCodeCreate) Validate() error {
	var invalid []string
	if mpb.openid == "" {
		invalid = append(invalid, "openid")
	}
	if mpb.accessToken == "" {
		invalid = append(invalid, "access_token")
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
func (mpb *MiniProgramAppCodeCreate) Do(ctx context.Context) (*MiniProgramAppCodeCreateResponse, error) {
	// Check pre-conditions
	if err := mpb.Validate(); err != nil {
		return nil, err
	}
	// url params
	params := url.Values{}
	params.Set("openid", mpb.openid)
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
		Endpoint: MiniProgramAppCodeCreateEndpoint,
	})
	if err != nil {
		return nil, err
	}
	// Return operation response
	ret := new(MiniProgramAppCodeCreateResponse)
	if err := mpb.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// MiniProgramAppCodeCreateResponse MiniProgramAppCodeCreateResponse
type MiniProgramAppCodeCreateResponse struct {
	CommonError
	UnionID string `json:"unionid"`
}
