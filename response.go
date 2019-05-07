package wechat

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

var (
	// ErrResponseSize is raised if a response body exceeds the given max body size.
	ErrResponseSize = errors.New("wechat: response size too large")
)

// Response represents a response from wechat.
type Response struct {
	// StatusCode is the HTTP status code, e.g. 200.
	StatusCode int
	// Header is the HTTP header from the HTTP response.
	// Keys in the map are canonicalized (see http.CanonicalHeaderKey).
	Header http.Header
	// Body is the deserialized response body.
	Body json.RawMessage
}

// newResponse creates a new response from the HTTP response.
func (c *Client) newResponse(res *http.Response, maxBodySize int64) (*Response, error) {
	r := &Response{
		StatusCode: res.StatusCode,
		Header:     res.Header,
	}
	if res.Body != nil {
		body := io.Reader(res.Body)
		if maxBodySize > 0 {
			if res.ContentLength > maxBodySize {
				return nil, ErrResponseSize
			}
			body = io.LimitReader(body, maxBodySize+1)
		}
		slurp, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, errors.Wrap(err, "Response.newResponse")
		}
		if maxBodySize > 0 && int64(len(slurp)) > maxBodySize {
			return nil, ErrResponseSize
		}
		// HEAD requests return a body but no content
		if len(slurp) > 0 {
			r.Body = json.RawMessage(slurp)
		}
	}
	return r, nil
}
