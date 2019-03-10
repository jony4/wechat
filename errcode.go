package wechat

import (
	"fmt"
)

// CommonError CommonError
type CommonError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// DecodeWithCommonError DecodeWithCommonError
func DecodeWithCommonError(apiName string, ce CommonError) (err error) {
	if ce.ErrCode != 0 {
		return fmt.Errorf("%s Error , errcode=%d , errmsg=%s", apiName, ce.ErrCode, ce.ErrMsg)
	}
	return nil
}
