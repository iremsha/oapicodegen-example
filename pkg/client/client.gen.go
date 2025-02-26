// Package client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// BankRequest defines model for BankRequest.
type BankRequest struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

// BankResponse defines model for BankResponse.
type BankResponse struct {
	Address string `json:"address"`
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Owner   Owner  `json:"owner"`
}

// BanksCardResponse defines model for BanksCardResponse.
type BanksCardResponse struct {
	Cvv  int    `json:"cvv"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// CardRequest defines model for CardRequest.
type CardRequest struct {
	Cvv  int    `json:"cvv"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// CardResponse defines model for CardResponse.
type CardResponse struct {
	Cvv  *int   `json:"cvv,omitempty"`
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Owner defines model for Owner.
type Owner struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// CreateApiV1BankJSONRequestBody defines body for CreateApiV1Bank for application/json ContentType.
type CreateApiV1BankJSONRequestBody = BankRequest

// UpdateApiV1BankJSONRequestBody defines body for UpdateApiV1Bank for application/json ContentType.
type UpdateApiV1BankJSONRequestBody = BankRequest

// CreateApiV1CardJSONRequestBody defines body for CreateApiV1Card for application/json ContentType.
type CreateApiV1CardJSONRequestBody = CardRequest

// UpdateApiV1CardJSONRequestBody defines body for UpdateApiV1Card for application/json ContentType.
type UpdateApiV1CardJSONRequestBody = CardRequest

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// CreateApiV1BankWithBody request with any body
	CreateApiV1BankWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateApiV1Bank(ctx context.Context, body CreateApiV1BankJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1Bank request
	GetApiV1Bank(ctx context.Context, bankId int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateApiV1BankWithBody request with any body
	UpdateApiV1BankWithBody(ctx context.Context, bankId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateApiV1Bank(ctx context.Context, bankId int, body UpdateApiV1BankJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1BankCards request
	GetApiV1BankCards(ctx context.Context, bankId int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1Cards request
	GetApiV1Cards(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateApiV1CardWithBody request with any body
	CreateApiV1CardWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateApiV1Card(ctx context.Context, body CreateApiV1CardJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApiV1Card request
	GetApiV1Card(ctx context.Context, cardId int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateApiV1CardWithBody request with any body
	UpdateApiV1CardWithBody(ctx context.Context, cardId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateApiV1Card(ctx context.Context, cardId int, body UpdateApiV1CardJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) CreateApiV1BankWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateApiV1BankRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateApiV1Bank(ctx context.Context, body CreateApiV1BankJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateApiV1BankRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1Bank(ctx context.Context, bankId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1BankRequest(c.Server, bankId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateApiV1BankWithBody(ctx context.Context, bankId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateApiV1BankRequestWithBody(c.Server, bankId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateApiV1Bank(ctx context.Context, bankId int, body UpdateApiV1BankJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateApiV1BankRequest(c.Server, bankId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1BankCards(ctx context.Context, bankId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1BankCardsRequest(c.Server, bankId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1Cards(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1CardsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateApiV1CardWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateApiV1CardRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateApiV1Card(ctx context.Context, body CreateApiV1CardJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateApiV1CardRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApiV1Card(ctx context.Context, cardId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV1CardRequest(c.Server, cardId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateApiV1CardWithBody(ctx context.Context, cardId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateApiV1CardRequestWithBody(c.Server, cardId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateApiV1Card(ctx context.Context, cardId int, body UpdateApiV1CardJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateApiV1CardRequest(c.Server, cardId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewCreateApiV1BankRequest calls the generic CreateApiV1Bank builder with application/json body
func NewCreateApiV1BankRequest(server string, body CreateApiV1BankJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateApiV1BankRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateApiV1BankRequestWithBody generates requests for CreateApiV1Bank with any type of body
func NewCreateApiV1BankRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/banks/")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetApiV1BankRequest generates requests for GetApiV1Bank
func NewGetApiV1BankRequest(server string, bankId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "bank_id", runtime.ParamLocationPath, bankId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/banks/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateApiV1BankRequest calls the generic UpdateApiV1Bank builder with application/json body
func NewUpdateApiV1BankRequest(server string, bankId int, body UpdateApiV1BankJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateApiV1BankRequestWithBody(server, bankId, "application/json", bodyReader)
}

// NewUpdateApiV1BankRequestWithBody generates requests for UpdateApiV1Bank with any type of body
func NewUpdateApiV1BankRequestWithBody(server string, bankId int, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "bank_id", runtime.ParamLocationPath, bankId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/banks/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetApiV1BankCardsRequest generates requests for GetApiV1BankCards
func NewGetApiV1BankCardsRequest(server string, bankId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "bank_id", runtime.ParamLocationPath, bankId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/banks/%s/cards", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiV1CardsRequest generates requests for GetApiV1Cards
func NewGetApiV1CardsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/cards/")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateApiV1CardRequest calls the generic CreateApiV1Card builder with application/json body
func NewCreateApiV1CardRequest(server string, body CreateApiV1CardJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateApiV1CardRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateApiV1CardRequestWithBody generates requests for CreateApiV1Card with any type of body
func NewCreateApiV1CardRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/cards/")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetApiV1CardRequest generates requests for GetApiV1Card
func NewGetApiV1CardRequest(server string, cardId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "card_id", runtime.ParamLocationPath, cardId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/cards/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateApiV1CardRequest calls the generic UpdateApiV1Card builder with application/json body
func NewUpdateApiV1CardRequest(server string, cardId int, body UpdateApiV1CardJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateApiV1CardRequestWithBody(server, cardId, "application/json", bodyReader)
}

// NewUpdateApiV1CardRequestWithBody generates requests for UpdateApiV1Card with any type of body
func NewUpdateApiV1CardRequestWithBody(server string, cardId int, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "card_id", runtime.ParamLocationPath, cardId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v1/cards/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// CreateApiV1BankWithBodyWithResponse request with any body
	CreateApiV1BankWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateApiV1BankResponse, error)

	CreateApiV1BankWithResponse(ctx context.Context, body CreateApiV1BankJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateApiV1BankResponse, error)

	// GetApiV1BankWithResponse request
	GetApiV1BankWithResponse(ctx context.Context, bankId int, reqEditors ...RequestEditorFn) (*GetApiV1BankResponse, error)

	// UpdateApiV1BankWithBodyWithResponse request with any body
	UpdateApiV1BankWithBodyWithResponse(ctx context.Context, bankId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateApiV1BankResponse, error)

	UpdateApiV1BankWithResponse(ctx context.Context, bankId int, body UpdateApiV1BankJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateApiV1BankResponse, error)

	// GetApiV1BankCardsWithResponse request
	GetApiV1BankCardsWithResponse(ctx context.Context, bankId int, reqEditors ...RequestEditorFn) (*GetApiV1BankCardsResponse, error)

	// GetApiV1CardsWithResponse request
	GetApiV1CardsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1CardsResponse, error)

	// CreateApiV1CardWithBodyWithResponse request with any body
	CreateApiV1CardWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateApiV1CardResponse, error)

	CreateApiV1CardWithResponse(ctx context.Context, body CreateApiV1CardJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateApiV1CardResponse, error)

	// GetApiV1CardWithResponse request
	GetApiV1CardWithResponse(ctx context.Context, cardId int, reqEditors ...RequestEditorFn) (*GetApiV1CardResponse, error)

	// UpdateApiV1CardWithBodyWithResponse request with any body
	UpdateApiV1CardWithBodyWithResponse(ctx context.Context, cardId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateApiV1CardResponse, error)

	UpdateApiV1CardWithResponse(ctx context.Context, cardId int, body UpdateApiV1CardJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateApiV1CardResponse, error)
}

type CreateApiV1BankResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *BankResponse
}

// Status returns HTTPResponse.Status
func (r CreateApiV1BankResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateApiV1BankResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1BankResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *BankResponse
}

// Status returns HTTPResponse.Status
func (r GetApiV1BankResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1BankResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateApiV1BankResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *BankResponse
}

// Status returns HTTPResponse.Status
func (r UpdateApiV1BankResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateApiV1BankResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1BankCardsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]BanksCardResponse
}

// Status returns HTTPResponse.Status
func (r GetApiV1BankCardsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1BankCardsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1CardsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]CardResponse
}

// Status returns HTTPResponse.Status
func (r GetApiV1CardsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1CardsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateApiV1CardResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CardResponse
}

// Status returns HTTPResponse.Status
func (r CreateApiV1CardResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateApiV1CardResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiV1CardResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CardResponse
}

// Status returns HTTPResponse.Status
func (r GetApiV1CardResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiV1CardResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateApiV1CardResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CardResponse
}

// Status returns HTTPResponse.Status
func (r UpdateApiV1CardResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateApiV1CardResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// CreateApiV1BankWithBodyWithResponse request with arbitrary body returning *CreateApiV1BankResponse
func (c *ClientWithResponses) CreateApiV1BankWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateApiV1BankResponse, error) {
	rsp, err := c.CreateApiV1BankWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateApiV1BankResponse(rsp)
}

func (c *ClientWithResponses) CreateApiV1BankWithResponse(ctx context.Context, body CreateApiV1BankJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateApiV1BankResponse, error) {
	rsp, err := c.CreateApiV1Bank(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateApiV1BankResponse(rsp)
}

// GetApiV1BankWithResponse request returning *GetApiV1BankResponse
func (c *ClientWithResponses) GetApiV1BankWithResponse(ctx context.Context, bankId int, reqEditors ...RequestEditorFn) (*GetApiV1BankResponse, error) {
	rsp, err := c.GetApiV1Bank(ctx, bankId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1BankResponse(rsp)
}

// UpdateApiV1BankWithBodyWithResponse request with arbitrary body returning *UpdateApiV1BankResponse
func (c *ClientWithResponses) UpdateApiV1BankWithBodyWithResponse(ctx context.Context, bankId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateApiV1BankResponse, error) {
	rsp, err := c.UpdateApiV1BankWithBody(ctx, bankId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateApiV1BankResponse(rsp)
}

func (c *ClientWithResponses) UpdateApiV1BankWithResponse(ctx context.Context, bankId int, body UpdateApiV1BankJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateApiV1BankResponse, error) {
	rsp, err := c.UpdateApiV1Bank(ctx, bankId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateApiV1BankResponse(rsp)
}

// GetApiV1BankCardsWithResponse request returning *GetApiV1BankCardsResponse
func (c *ClientWithResponses) GetApiV1BankCardsWithResponse(ctx context.Context, bankId int, reqEditors ...RequestEditorFn) (*GetApiV1BankCardsResponse, error) {
	rsp, err := c.GetApiV1BankCards(ctx, bankId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1BankCardsResponse(rsp)
}

// GetApiV1CardsWithResponse request returning *GetApiV1CardsResponse
func (c *ClientWithResponses) GetApiV1CardsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV1CardsResponse, error) {
	rsp, err := c.GetApiV1Cards(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1CardsResponse(rsp)
}

// CreateApiV1CardWithBodyWithResponse request with arbitrary body returning *CreateApiV1CardResponse
func (c *ClientWithResponses) CreateApiV1CardWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateApiV1CardResponse, error) {
	rsp, err := c.CreateApiV1CardWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateApiV1CardResponse(rsp)
}

func (c *ClientWithResponses) CreateApiV1CardWithResponse(ctx context.Context, body CreateApiV1CardJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateApiV1CardResponse, error) {
	rsp, err := c.CreateApiV1Card(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateApiV1CardResponse(rsp)
}

// GetApiV1CardWithResponse request returning *GetApiV1CardResponse
func (c *ClientWithResponses) GetApiV1CardWithResponse(ctx context.Context, cardId int, reqEditors ...RequestEditorFn) (*GetApiV1CardResponse, error) {
	rsp, err := c.GetApiV1Card(ctx, cardId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV1CardResponse(rsp)
}

// UpdateApiV1CardWithBodyWithResponse request with arbitrary body returning *UpdateApiV1CardResponse
func (c *ClientWithResponses) UpdateApiV1CardWithBodyWithResponse(ctx context.Context, cardId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateApiV1CardResponse, error) {
	rsp, err := c.UpdateApiV1CardWithBody(ctx, cardId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateApiV1CardResponse(rsp)
}

func (c *ClientWithResponses) UpdateApiV1CardWithResponse(ctx context.Context, cardId int, body UpdateApiV1CardJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateApiV1CardResponse, error) {
	rsp, err := c.UpdateApiV1Card(ctx, cardId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateApiV1CardResponse(rsp)
}

// ParseCreateApiV1BankResponse parses an HTTP response from a CreateApiV1BankWithResponse call
func ParseCreateApiV1BankResponse(rsp *http.Response) (*CreateApiV1BankResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateApiV1BankResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest BankResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetApiV1BankResponse parses an HTTP response from a GetApiV1BankWithResponse call
func ParseGetApiV1BankResponse(rsp *http.Response) (*GetApiV1BankResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1BankResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest BankResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseUpdateApiV1BankResponse parses an HTTP response from a UpdateApiV1BankWithResponse call
func ParseUpdateApiV1BankResponse(rsp *http.Response) (*UpdateApiV1BankResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateApiV1BankResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest BankResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetApiV1BankCardsResponse parses an HTTP response from a GetApiV1BankCardsWithResponse call
func ParseGetApiV1BankCardsResponse(rsp *http.Response) (*GetApiV1BankCardsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1BankCardsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []BanksCardResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetApiV1CardsResponse parses an HTTP response from a GetApiV1CardsWithResponse call
func ParseGetApiV1CardsResponse(rsp *http.Response) (*GetApiV1CardsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1CardsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []CardResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateApiV1CardResponse parses an HTTP response from a CreateApiV1CardWithResponse call
func ParseCreateApiV1CardResponse(rsp *http.Response) (*CreateApiV1CardResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateApiV1CardResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CardResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetApiV1CardResponse parses an HTTP response from a GetApiV1CardWithResponse call
func ParseGetApiV1CardResponse(rsp *http.Response) (*GetApiV1CardResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV1CardResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CardResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseUpdateApiV1CardResponse parses an HTTP response from a UpdateApiV1CardWithResponse call
func ParseUpdateApiV1CardResponse(rsp *http.Response) (*UpdateApiV1CardResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateApiV1CardResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CardResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
