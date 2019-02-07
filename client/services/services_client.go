// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetV1Services get v1 services API
*/
func (a *Client) GetV1Services(params *GetV1ServicesParams, authInfo runtime.ClientAuthInfoWriter) (*GetV1ServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getV1Services",
		Method:             "GET",
		PathPattern:        "/v1/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1ServicesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetV1ServicesOK), nil

}

/*
PostV1Services Create a service
*/
func (a *Client) PostV1Services(params *PostV1ServicesParams, authInfo runtime.ClientAuthInfoWriter) (*PostV1ServicesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1ServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postV1Services",
		Method:             "POST",
		PathPattern:        "/v1/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostV1ServicesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostV1ServicesCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
