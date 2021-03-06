package wechat

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
)

const (
	// Version is the current version of wechat.
	Version = "0.0.3"

	// DefaultScheme is the default protocol scheme to use.
	DefaultScheme = "https"

	// DefaultSendGetBodyAs is the HTTP method to use when wechat is sending
	// a GET request with a body.
	DefaultSendGetBodyAs = "GET"

	// DefaultGzipEnabled specifies if gzip compression is enabled by default.
	DefaultGzipEnabled = false

	// DefaultCacheExpiration of access token
	DefaultCacheExpiration = 7200 * time.Second

	// DefaultCacheInterval cleanup cache
	DefaultCacheInterval = 240 * time.Minute
)

var (
	// ErrNoClient is raised when no Wechat node is available.
	ErrNoClient = errors.New("no Wechat node available")

	// ErrTimeout is raised when a request timed out, e.g. when WaitForStatus
	// didn't return in time.
	ErrTimeout = errors.New("timeout")

	// ErrNoBaseURI is raised when no base uri is available
	ErrNoBaseURI = errors.New("no base uri available")

	// ErrNoEndpoint is raised when no endpoint is available
	ErrNoEndpoint = errors.New("no endpoint available")
)

// ClientOptionFunc is a function that configures a Client.
// It is used in NewClient.
type ClientOptionFunc func(*Client) error

// Client is a common wechat client.
// Create one by calling NewClient.
type Client struct {
	httpClient *http.Client

	running bool         // true if the client's background processes are running
	mu      sync.RWMutex // guards the next block

	errorlog      Logger  // error log for critical messages
	infolog       Logger  // information log for e.g. response times
	tracelog      Logger  // trace log for debugging
	scheme        string  // http or https
	decoder       Decoder // used to decode data sent from wechat
	sendGetBodyAs string  // override for when sending a GET with a body
	gzipEnabled   bool    // gzip compression enabled or disabled (default)

	cache Cache // Cache backend, used for saving access token etc.
}

// NewClient creates a new short-lived Client that can be used in
// use cases where you need e.g. one client per request.
func NewClient(options ...ClientOptionFunc) (*Client, error) {
	c := &Client{
		httpClient:    http.DefaultClient,
		scheme:        DefaultScheme,
		decoder:       &DefaultDecoder{},
		sendGetBodyAs: DefaultSendGetBodyAs,
		gzipEnabled:   DefaultGzipEnabled,
	}
	// Run the options on it
	for _, option := range options {
		if err := option(c); err != nil {
			c.errorf("NewClient err %v", err)
			return nil, err
		}
	}
	if c.cache == nil {
		opts := []MemCacheOptFunc{
			SetDefaultExpiration(DefaultCacheExpiration),
			SetDefaultInterval(DefaultCacheInterval),
		}
		cache, err := NewMemCache(opts...)
		if err != nil {
			c.errorf("NewClient.NewMemCache err %v", err)
			return nil, err
		}
		c.cache = cache
	}
	c.mu.Lock()
	c.running = true
	c.mu.Unlock()
	return c, nil
}

// SetCacheBackend SetCacheBackend which implentment Cache interface
func SetCacheBackend(cache Cache) ClientOptionFunc {
	return func(c *Client) error {
		c.cache = cache
		return nil
	}
}

// SetHTTPClient can be used to specify the http.Client to use when making
// HTTP requests to wechat.
func SetHTTPClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *Client) error {
		if httpClient != nil {
			c.httpClient = httpClient
		} else {
			c.httpClient = http.DefaultClient
		}
		return nil
	}
}

// SetHTTPClientInsecureSkipVerify SetHTTPClientInsecureSkipVerify
func SetHTTPClientInsecureSkipVerify() ClientOptionFunc {
	return func(c *Client) error {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		c.httpClient.Transport = transport
		return nil
	}
}

// SetErrorLog sets the logger for critical messages like nodes joining
// or leaving the cluster or failing requests. It is nil by default.
func SetErrorLog(logger Logger) ClientOptionFunc {
	return func(c *Client) error {
		c.errorlog = logger
		return nil
	}
}

// SetInfoLog sets the logger for informational messages, e.g. requests
// and their response times. It is nil by default.
func SetInfoLog(logger Logger) ClientOptionFunc {
	return func(c *Client) error {
		c.infolog = logger
		return nil
	}
}

// SetTraceLog specifies the log.Logger to use for output of HTTP requests
// and responses which is helpful during debugging. It is nil by default.
func SetTraceLog(logger Logger) ClientOptionFunc {
	return func(c *Client) error {
		c.tracelog = logger
		return nil
	}
}

// SetScheme sets the HTTP scheme to look for when sniffing (http or https).
// This is http by default.
func SetScheme(scheme string) ClientOptionFunc {
	return func(c *Client) error {
		c.scheme = scheme
		return nil
	}
}

// SetSendGetBodyAs specifies the HTTP method to use when sending a GET request
// with a body. It is GET by default.
func SetSendGetBodyAs(httpMethod string) ClientOptionFunc {
	return func(c *Client) error {
		c.sendGetBodyAs = httpMethod
		return nil
	}
}

// SetGzip enables or disables gzip compression (disabled by default).
func SetGzip(enabled bool) ClientOptionFunc {
	return func(c *Client) error {
		c.gzipEnabled = enabled
		return nil
	}
}

// SetDecoder sets the Decoder to use when decoding data from Wechat.
// DefaultDecoder is used by default.
func SetDecoder(decoder Decoder) ClientOptionFunc {
	return func(c *Client) error {
		if decoder != nil {
			c.decoder = decoder
		} else {
			c.decoder = &DefaultDecoder{}
		}
		return nil
	}
}

// IsRunning returns true if the background processes of the client are
// running, false otherwise.
func (c *Client) IsRunning() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.running
}

// Start starts the background processes like periodic health checks.
// You don't need to run Start when creating a client with NewClient;
// the background processes are run by default.
//
// If the background processes are already running, this is a no-op.
func (c *Client) Start() {
	c.mu.RLock()
	if c.running {
		c.mu.RUnlock()
		return
	}
	c.mu.RUnlock()

	c.mu.Lock()
	c.running = true
	c.mu.Unlock()
}

// Stop stops the background processes that the client is running,
//
// If the background processes are not running, this is a no-op.
func (c *Client) Stop() {
	c.mu.RLock()
	if !c.running {
		c.mu.RUnlock()
		return
	}
	c.mu.RUnlock()

	c.mu.Lock()
	c.running = false
	c.mu.Unlock()
}

// errorf logs to the error log.
func (c *Client) errorf(format string, args ...interface{}) {
	if c.errorlog != nil {
		c.errorlog.Printf(format, args...)
	}
}

// infof logs informational messages.
func (c *Client) infof(format string, args ...interface{}) {
	if c.infolog != nil {
		c.infolog.Printf(format, args...)
	}
}

// tracef logs to the trace log.
func (c *Client) tracef(format string, args ...interface{}) {
	if c.tracelog != nil {
		c.tracelog.Printf(format, args...)
	}
}

// dumpRequest dumps the given HTTP request to the trace log.
func (c *Client) dumpRequest(r *http.Request) {
	if c.tracelog != nil {
		out, err := httputil.DumpRequestOut(r, true)
		if err == nil {
			c.tracef("%s\n", string(out))
		}
	}
}

// dumpResponse dumps the given HTTP response to the trace log.
func (c *Client) dumpResponse(resp *http.Response) {
	if c.tracelog != nil {
		out, err := httputil.DumpResponse(resp, true)
		if err == nil {
			c.tracef("%s\n", string(out))
		}
	}
}

// -- PerformRequest --

// PerformRequestOptions must be passed into PerformRequest.
type PerformRequestOptions struct {
	Method          string
	Params          url.Values
	Body            interface{}
	FormValue       []byte
	FormFieldName   string
	FormFileName    string
	ContentType     string
	IgnoreErrors    []int
	Headers         http.Header
	MaxResponseSize int64
	BaseURI         string
	Endpoint        string
}

// PerformRequest does a HTTP request to wechat.
func (c *Client) PerformRequest(ctx context.Context, opt PerformRequestOptions) (*Response, error) {
	start := time.Now().UTC()

	c.mu.Lock()
	sendGetBodyAs := c.sendGetBodyAs
	gzipEnabled := c.gzipEnabled
	pathWithParams := fmt.Sprintf("%s://%s/%s", c.scheme, opt.BaseURI, opt.Endpoint)
	if len(opt.Params) > 0 {
		pathWithParams += "?" + opt.Params.Encode()
	}
	c.mu.Unlock()

	var (
		err  error
		req  *Request
		resp *Response
	)

	// Change method if sendGetBodyAs is specified.
	if opt.Method == "GET" && opt.Body != nil && sendGetBodyAs != "GET" {
		opt.Method = sendGetBodyAs
	}

	req, err = NewRequest(opt.Method, pathWithParams, nil)
	if err != nil {
		c.errorf("wechat: cannot create request for %s %s: %v", strings.ToUpper(opt.Method), pathWithParams, err)
		return nil, err
	}

	if opt.ContentType != "" {
		req.Header.Set("Content-Type", opt.ContentType)
	}

	if len(opt.Headers) > 0 {
		for key, value := range opt.Headers {
			for _, v := range value {
				req.Header.Add(key, v)
			}
		}
	}

	// Set body
	if opt.Body != nil {
		err = req.SetBody(opt.Body, gzipEnabled)
		if err != nil {
			c.errorf("wechat: couldn't set body %+v for request: %v", opt.Body, err)
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
	}

	// Tracing
	c.dumpRequest((*http.Request)(req))

	// Get response
	res, err := c.httpClient.Do((*http.Request)(req).WithContext(ctx))

	if err != nil {
		c.errorf("wechat: couldn't do request body %+v for request: %v", opt.Body, err)
		return nil, err
	}

	if IsContextErr(err) {
		c.errorf("PerformRequest.IsContextErr err %v", err)
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Tracing
	c.dumpResponse(res)

	// Check for errors
	if err := checkResponse((*http.Request)(req), res, opt.IgnoreErrors...); err != nil {
		// We still try to return a response.
		resp, err1 := c.newResponse(res, opt.MaxResponseSize)
		c.tracef("PerformRequest.newResponse try again err %v", err1)
		return resp, err1
	}

	resp, err = c.newResponse(res, opt.MaxResponseSize)
	if err != nil {
		c.tracef("PerformRequest.newResponse err %v", err)
		return nil, err
	}

	duration := time.Now().UTC().Sub(start)
	c.infof("%s %s [status:%d, request:%.3fs]",
		strings.ToUpper(opt.Method),
		req.URL,
		resp.StatusCode,
		float64(int64(duration/time.Millisecond))/1000)

	return resp, nil
}

// PerformFormRequest does a HTTP request to wechat.
func (c *Client) PerformFormRequest(ctx context.Context, opt PerformRequestOptions) (*Response, error) {
	start := time.Now().UTC()

	c.mu.Lock()
	sendGetBodyAs := c.sendGetBodyAs
	pathWithParams := fmt.Sprintf("%s://%s/%s", c.scheme, opt.BaseURI, opt.Endpoint)
	if len(opt.Params) > 0 {
		pathWithParams += "?" + opt.Params.Encode()
	}
	c.mu.Unlock()

	// Change method if sendGetBodyAs is specified.
	if opt.Method == "GET" && opt.Body != nil && sendGetBodyAs != "GET" {
		opt.Method = sendGetBodyAs
	}

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile(opt.FormFieldName, opt.FormFileName)
	if err != nil {
		return nil, errors.Wrap(err, "CreateFormFile")
	}
	if _, err := io.Copy(fileWriter, bytes.NewReader(opt.FormValue)); err != nil {
		return nil, errors.Wrap(err, "io.Copy")
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	var (
		req  *Request
		resp *Response
	)

	req, err = NewRequest(opt.Method, pathWithParams, bodyBuf)
	if err != nil {
		c.errorf("wechat: cannot create request for %s %s: %v", strings.ToUpper(opt.Method), pathWithParams, err)
		return nil, err
	}

	if len(opt.Headers) > 0 {
		for key, value := range opt.Headers {
			for _, v := range value {
				req.Header.Add(key, v)
			}
		}
	}
	req.Header.Set("Content-Type", contentType)

	// Tracing
	c.dumpRequest((*http.Request)(req))

	// Get response
	res, err := c.httpClient.Do((*http.Request)(req).WithContext(ctx))

	if err != nil {
		c.errorf("wechat: couldn't do request body %+v for request: %v", opt.Body, err)
		return nil, err
	}

	if IsContextErr(err) {
		c.errorf("PerformRequest.IsContextErr err %v", err)
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// // Tracing
	c.dumpResponse(res)

	// // Check for errors
	if err := checkResponse(nil, res, opt.IgnoreErrors...); err != nil {
		// We still try to return a response.
		resp, err1 := c.newResponse(res, opt.MaxResponseSize)
		c.tracef("PerformRequest.newResponse try again err %v", err1)
		return resp, err1
	}

	resp, err = c.newResponse(res, opt.MaxResponseSize)
	if err != nil {
		c.tracef("PerformRequest.newResponse err %v", err)
		return nil, err
	}

	duration := time.Now().UTC().Sub(start)
	c.infof("%s %s [status:%d, request:%.3fs]",
		strings.ToUpper(opt.Method),
		req.URL,
		resp.StatusCode,
		float64(int64(duration/time.Millisecond))/1000)

	return resp, nil
}
