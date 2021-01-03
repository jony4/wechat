package wechat

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const (
	// MiniProgramPaidEndpoint Endpoint
	MiniProgramPaidEndpoint = "wxa/getpaidunionid"
)

// MiniProgramPaid MiniProgramPaid
type MiniProgramPaid struct {
	client *Client

	accessToken   string
	openid        string
	usingTransID  bool
	transactionID string
	mchID         string
	outTradeNo    string
}

// NewMiniProgramPaid return instance of NewMiniProgramPaid
func NewMiniProgramPaid(client *Client) *MiniProgramPaid {
	mpb := &MiniProgramPaid{
		client: client,
	}
	return mpb
}

// SetOpenID SetOpenID
func (mpb *MiniProgramPaid) SetOpenID(openid string) *MiniProgramPaid {
	mpb.openid = openid
	return mpb
}

// SetAccessToken SetAccessToken
func (mpb *MiniProgramPaid) SetAccessToken(accessToken string) *MiniProgramPaid {
	mpb.accessToken = accessToken
	return mpb
}

// SetTransactionID SetTransactionID
func (mpb *MiniProgramPaid) SetTransactionID(transactionID string) *MiniProgramPaid {
	mpb.transactionID = transactionID
	return mpb
}

// SetMchID SetMchID
func (mpb *MiniProgramPaid) SetMchID(mchID string) *MiniProgramPaid {
	mpb.mchID = mchID
	return mpb
}

// SetOutTradeNo SetOutTradeNo
func (mpb *MiniProgramPaid) SetOutTradeNo(outTradeNo string) *MiniProgramPaid {
	mpb.outTradeNo = outTradeNo
	return mpb
}

// SetUsingTransID SetUsingTransID
func (mpb *MiniProgramPaid) SetUsingTransID() *MiniProgramPaid {
	mpb.usingTransID = true
	return mpb
}

// Validate checks if the operation is valid.
func (mpb *MiniProgramPaid) Validate() error {
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
func (mpb *MiniProgramPaid) Do(ctx context.Context) (*MiniProgramPaidResponse, error) {
	// Check pre-conditions
	if err := mpb.Validate(); err != nil {
		return nil, errors.Wrap(err, "MiniProgramPaid.Do")
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
		BaseURI:  MiniProgramBaseHost,
		Endpoint: MiniProgramPaidEndpoint,
	})
	if err != nil {
		return nil, errors.Wrap(err, "MiniProgramPaid.Do")
	}
	// Return operation response
	ret := new(MiniProgramPaidResponse)
	if err := mpb.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, errors.Wrap(err, "MiniProgramPaid.Do")
	}
	return ret, nil
}

// MiniProgramPaidResponse MiniProgramPaidResponse
type MiniProgramPaidResponse struct {
	CommonError
	UnionID string `json:"unionid"`
}
