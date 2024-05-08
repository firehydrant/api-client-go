// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new metrics API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for metrics API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1MetricsIncidents(params *GetV1MetricsIncidentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsIncidentsOK, error)

	GetV1MetricsInfraType(params *GetV1MetricsInfraTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsInfraTypeOK, error)

	GetV1MetricsInfraTypeInfraID(params *GetV1MetricsInfraTypeInfraIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsInfraTypeInfraIDOK, error)

	GetV1MetricsMttx(params *GetV1MetricsMttxParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsMttxOK, error)

	GetV1MetricsRetrospectives(params *GetV1MetricsRetrospectivesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsRetrospectivesOK, error)

	GetV1MetricsUserInvolvements(params *GetV1MetricsUserInvolvementsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsUserInvolvementsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetV1MetricsIncidents lists incident metrics

Returns a report with time bucketed analytics data
*/
func (a *Client) GetV1MetricsIncidents(params *GetV1MetricsIncidentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsIncidentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1MetricsIncidentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1MetricsIncidents",
		Method:             "GET",
		PathPattern:        "/v1/metrics/incidents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1MetricsIncidentsReader{formats: a.formats},
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
	success, ok := result.(*GetV1MetricsIncidentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1MetricsIncidents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1MetricsInfraType lists metrics for a component type

Returns metrics for all components of a given type
*/
func (a *Client) GetV1MetricsInfraType(params *GetV1MetricsInfraTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsInfraTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1MetricsInfraTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1MetricsInfraType",
		Method:             "GET",
		PathPattern:        "/v1/metrics/{infra_type}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1MetricsInfraTypeReader{formats: a.formats},
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
	success, ok := result.(*GetV1MetricsInfraTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1MetricsInfraType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1MetricsInfraTypeInfraID shows metrics for a component

Return metrics for a specific component
*/
func (a *Client) GetV1MetricsInfraTypeInfraID(params *GetV1MetricsInfraTypeInfraIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsInfraTypeInfraIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1MetricsInfraTypeInfraIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1MetricsInfraTypeInfraId",
		Method:             "GET",
		PathPattern:        "/v1/metrics/{infra_type}/{infra_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1MetricsInfraTypeInfraIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1MetricsInfraTypeInfraIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1MetricsInfraTypeInfraId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1MetricsMttx Fetch infrastructure metrics based on custom query
*/
func (a *Client) GetV1MetricsMttx(params *GetV1MetricsMttxParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsMttxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1MetricsMttxParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1MetricsMttx",
		Method:             "GET",
		PathPattern:        "/v1/metrics/mttx",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1MetricsMttxReader{formats: a.formats},
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
	success, ok := result.(*GetV1MetricsMttxOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1MetricsMttx: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1MetricsRetrospectives lists retrospective metrics

Returns a report with retrospective analytics data
*/
func (a *Client) GetV1MetricsRetrospectives(params *GetV1MetricsRetrospectivesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsRetrospectivesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1MetricsRetrospectivesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1MetricsRetrospectives",
		Method:             "GET",
		PathPattern:        "/v1/metrics/retrospectives",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1MetricsRetrospectivesReader{formats: a.formats},
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
	success, ok := result.(*GetV1MetricsRetrospectivesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1MetricsRetrospectives: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1MetricsUserInvolvements lists user metrics

Returns a report with time bucketed analytics data
*/
func (a *Client) GetV1MetricsUserInvolvements(params *GetV1MetricsUserInvolvementsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsUserInvolvementsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1MetricsUserInvolvementsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1MetricsUserInvolvements",
		Method:             "GET",
		PathPattern:        "/v1/metrics/user_involvements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1MetricsUserInvolvementsReader{formats: a.formats},
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
	success, ok := result.(*GetV1MetricsUserInvolvementsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1MetricsUserInvolvements: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
