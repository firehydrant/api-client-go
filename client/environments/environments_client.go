// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new environments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for environments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1EnvironmentsEnvironmentID(params *DeleteV1EnvironmentsEnvironmentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1EnvironmentsEnvironmentIDOK, error)

	GetV1Environments(params *GetV1EnvironmentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1EnvironmentsOK, error)

	GetV1EnvironmentsEnvironmentID(params *GetV1EnvironmentsEnvironmentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1EnvironmentsEnvironmentIDOK, error)

	PatchV1EnvironmentsEnvironmentID(params *PatchV1EnvironmentsEnvironmentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1EnvironmentsEnvironmentIDOK, error)

	PostV1Environments(params *PostV1EnvironmentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1EnvironmentsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteV1EnvironmentsEnvironmentID archives an environment

Archive an environment
*/
func (a *Client) DeleteV1EnvironmentsEnvironmentID(params *DeleteV1EnvironmentsEnvironmentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1EnvironmentsEnvironmentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1EnvironmentsEnvironmentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1EnvironmentsEnvironmentId",
		Method:             "DELETE",
		PathPattern:        "/v1/environments/{environment_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1EnvironmentsEnvironmentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteV1EnvironmentsEnvironmentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1EnvironmentsEnvironmentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1Environments lists environments

List all of the environments that have been added to the organiation
*/
func (a *Client) GetV1Environments(params *GetV1EnvironmentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1EnvironmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1EnvironmentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1Environments",
		Method:             "GET",
		PathPattern:        "/v1/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1EnvironmentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1EnvironmentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1Environments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1EnvironmentsEnvironmentID gets an environment

Retrieves a single environment by ID
*/
func (a *Client) GetV1EnvironmentsEnvironmentID(params *GetV1EnvironmentsEnvironmentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1EnvironmentsEnvironmentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1EnvironmentsEnvironmentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1EnvironmentsEnvironmentId",
		Method:             "GET",
		PathPattern:        "/v1/environments/{environment_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1EnvironmentsEnvironmentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1EnvironmentsEnvironmentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1EnvironmentsEnvironmentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1EnvironmentsEnvironmentID updates an environment

Update a environments attributes
*/
func (a *Client) PatchV1EnvironmentsEnvironmentID(params *PatchV1EnvironmentsEnvironmentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1EnvironmentsEnvironmentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1EnvironmentsEnvironmentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1EnvironmentsEnvironmentId",
		Method:             "PATCH",
		PathPattern:        "/v1/environments/{environment_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1EnvironmentsEnvironmentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchV1EnvironmentsEnvironmentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1EnvironmentsEnvironmentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1Environments creates an environment

Creates an environment for the organization
*/
func (a *Client) PostV1Environments(params *PostV1EnvironmentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1EnvironmentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1EnvironmentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1Environments",
		Method:             "POST",
		PathPattern:        "/v1/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1EnvironmentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV1EnvironmentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1Environments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
