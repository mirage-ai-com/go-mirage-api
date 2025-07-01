// go-mirage-api
//
// Copyright 2023, Valerian Saliou
// Author: Valerian Saliou <valerian@valeriansaliou.name>

package mirage

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	libraryVersion            = "1.12.0"
	defaultRestEndpointURL    = "https://api.mirage-ai.com/v1/"
	userAgent                 = "go-mirage-api/" + libraryVersion
	acceptContentType         = "application/json"
	clientTimeout             = 40
	clientIdleConnTimeout     = 45
	clientMaxIdleConns        = 16
	clientMaxConnsPerHost     = 64
	clientMaxIdleConnsPerHost = 4
)

var errorDoNilRequest = errors.New("request could not be constructed")

// ClientConfig mapping
type ClientConfig struct {
	HTTPClient      *http.Client
	RestEndpointURL string
}

type auth struct {
	Username string
	Password string
}

// Client maps an API client
type Client struct {
	config *ClientConfig
	client *http.Client
	auth   *auth

	BaseURL   *url.URL
	UserAgent string

	common service

	Task *TaskService
	Data *DataService
}

type service struct {
	client *Client
}

// RequestContext maps custom context for a request
type RequestContext struct {
	Trace string `json:"string"`
}

// Response maps an API HTTP response
type Response struct {
	*http.Response
}

type errorResponse struct {
	Response *http.Response
	Reason   string `json:"reason,omitempty"`
	Message  string `json:"message,omitempty"`
}

// Error prints an error response
func (response *errorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		response.Response.Request.Method, response.Response.Request.URL,
		response.Response.StatusCode, response.Reason)
}

// NewWithConfig returns a new API client
func NewWithConfig(userID string, secretKey string, config ClientConfig) *Client {
	// Defaults
	if config.HTTPClient == nil {
		// Build client transport
		clientTransport := http.DefaultTransport.(*http.Transport).Clone()
		clientTransport.IdleConnTimeout = time.Duration(clientIdleConnTimeout * time.Second)
		clientTransport.MaxIdleConns = clientMaxIdleConns
		clientTransport.MaxConnsPerHost = clientMaxConnsPerHost
		clientTransport.MaxIdleConnsPerHost = clientMaxIdleConnsPerHost

		// Create client
		config.HTTPClient = &http.Client{
			Timeout:   time.Duration(clientTimeout * time.Second),
			Transport: clientTransport,
		}
	}
	if config.RestEndpointURL == "" {
		config.RestEndpointURL = defaultRestEndpointURL
	}

	// Create client
	baseURL, _ := url.Parse(config.RestEndpointURL)

	client := &Client{config: &config, client: config.HTTPClient, auth: &auth{userID, secretKey}, BaseURL: baseURL, UserAgent: userAgent}
	client.common.client = client

	// Map services
	client.Task = (*TaskService)(&client.common)
	client.Data = (*DataService)(&client.common)

	return client
}

// New returns a new API client
func New(userID string, secretKey string) *Client {
	return NewWithConfig(userID, secretKey, ClientConfig{})
}

// NewRequest creates an API request
func (client *Client) NewRequest(method, urlStr string, body interface{}, ctx RequestContext) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	url := client.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(client.auth.Username, client.auth.Password)

	req.Header.Add("Accept", acceptContentType)
	req.Header.Add("Content-Type", acceptContentType)

	if client.UserAgent != "" {
		req.Header.Add("User-Agent", client.UserAgent)
	}

	if ctx.Trace != "" {
		// Stamp request with provided trace identifier (this is optional, but \
		//   can be used to track request flows across Mirage backend systems, for \
		//   debugging purposes)
		req.Header.Add("X-Request-ID", ctx.Trace)
	}

	return req, nil
}

// Do sends an API request
func (client *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	if req == nil {
		return nil, errorDoNilRequest
	}

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	response := newResponse(resp)

	err = checkResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}

	return response, err
}

// newResponse creates an HTTP response
func newResponse(httpResponse *http.Response) *Response {
	response := &Response{Response: httpResponse}

	return response
}

// checkResponse checks response for errors
func checkResponse(response *http.Response) error {
	// No error in response? (HTTP 2xx)
	if code := response.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	// Map response error data (eg. HTTP 4xx)
	errorResponse := &errorResponse{Response: response}

	data, err := ioutil.ReadAll(response.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}
