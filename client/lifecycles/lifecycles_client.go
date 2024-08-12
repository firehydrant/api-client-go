// Code generated by go-swagger; DO NOT EDIT.

package lifecycles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new lifecycles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for lifecycles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1LifecyclesMilestonesMilestoneID(params *DeleteV1LifecyclesMilestonesMilestoneIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1LifecyclesMilestonesMilestoneIDOK, error)

	PatchV1LifecyclesMilestonesMilestoneID(params *PatchV1LifecyclesMilestonesMilestoneIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1LifecyclesMilestonesMilestoneIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteV1LifecyclesMilestonesMilestoneID deletes a milestone

Delete a milestone
*/
func (a *Client) DeleteV1LifecyclesMilestonesMilestoneID(params *DeleteV1LifecyclesMilestonesMilestoneIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1LifecyclesMilestonesMilestoneIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1LifecyclesMilestonesMilestoneIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1LifecyclesMilestonesMilestoneId",
		Method:             "DELETE",
		PathPattern:        "/v1/lifecycles/milestones/{milestone_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1LifecyclesMilestonesMilestoneIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1LifecyclesMilestonesMilestoneIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1LifecyclesMilestonesMilestoneId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1LifecyclesMilestonesMilestoneID updates a milestone

Update a milestone
*/
func (a *Client) PatchV1LifecyclesMilestonesMilestoneID(params *PatchV1LifecyclesMilestonesMilestoneIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1LifecyclesMilestonesMilestoneIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1LifecyclesMilestonesMilestoneIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1LifecyclesMilestonesMilestoneId",
		Method:             "PATCH",
		PathPattern:        "/v1/lifecycles/milestones/{milestone_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1LifecyclesMilestonesMilestoneIDReader{formats: a.formats},
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
	success, ok := result.(*PatchV1LifecyclesMilestonesMilestoneIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1LifecyclesMilestonesMilestoneId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
