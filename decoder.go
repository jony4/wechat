package wechat

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
)

// Decoder is used to decode responses from wechat.
// Users of wechat can implement their own marshaler for advanced purposes
// and set them per Client (see SetDecoder). If none is specified,
// DefaultDecoder is used.
type Decoder interface {
	Decode(data []byte, v interface{}) error
}

// DefaultDecoder uses json.Unmarshal from the Go standard library
// to decode JSON data.
type DefaultDecoder struct{}

// Decode decodes with json.Unmarshal from the Go standard library.
func (u *DefaultDecoder) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// NumberDecoder uses json.NewDecoder, with UseNumber() enabled, from
// the Go standard library to decode JSON data.
type NumberDecoder struct{}

// Decode decodes with json.Unmarshal from the Go standard library.
func (u *NumberDecoder) Decode(data []byte, v interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	return dec.Decode(v)
}

// XMLDecoder uses xml.NewDecoder from the Go standard library
// to decode XML data.
type XMLDecoder struct{}

// Decode decodes with json.Unmarshal from the Go standard library.
func (u *XMLDecoder) Decode(data []byte, v interface{}) error {
	dec := xml.NewDecoder(bytes.NewReader(data))
	return dec.Decode(v)
}
