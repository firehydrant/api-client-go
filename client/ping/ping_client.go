// Code generated by go-swagger; DO NOT EDIT.

package ping

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new ping API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ping API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetV1Ping Simple endpoint to verify your API connection is working
*/
func (a *Client) GetV1Ping(params *GetV1PingParams, authInfo runtime.ClientAuthInfoWriter) (*GetV1PingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1PingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getV1Ping",
		Method:             "GET",
		PathPattern:        "/v1/ping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1PingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetV1PingOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
