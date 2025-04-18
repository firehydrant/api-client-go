// Code generated by go-swagger; DO NOT EDIT.

package audiences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new audiences API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for audiences API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1AudiencesAudienceID(params *DeleteV1AudiencesAudienceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1AudiencesAudienceIDOK, error)

	GetV1Audiences(params *GetV1AudiencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1AudiencesOK, error)

	GetV1AudiencesAudienceID(params *GetV1AudiencesAudienceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1AudiencesAudienceIDOK, error)

	GetV1AudiencesAudienceIDSummariesIncidentID(params *GetV1AudiencesAudienceIDSummariesIncidentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1AudiencesAudienceIDSummariesIncidentIDOK, error)

	GetV1AudiencesMemberMemberIDDefault(params *GetV1AudiencesMemberMemberIDDefaultParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1AudiencesMemberMemberIDDefaultOK, error)

	GetV1AudiencesSummariesIncidentID(params *GetV1AudiencesSummariesIncidentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1AudiencesSummariesIncidentIDOK, error)

	PatchV1AudiencesAudienceID(params *PatchV1AudiencesAudienceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1AudiencesAudienceIDOK, error)

	PatchV1AudiencesAudienceIDRestore(params *PatchV1AudiencesAudienceIDRestoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1AudiencesAudienceIDRestoreOK, error)

	PostV1Audiences(params *PostV1AudiencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1AudiencesCreated, error)

	PostV1AudiencesAudienceIDSummariesIncidentID(params *PostV1AudiencesAudienceIDSummariesIncidentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1AudiencesAudienceIDSummariesIncidentIDCreated, error)

	PutV1AudiencesMemberMemberIDDefault(params *PutV1AudiencesMemberMemberIDDefaultParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutV1AudiencesMemberMemberIDDefaultOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteV1AudiencesAudienceID archives audience

Archive an audience
*/
func (a *Client) DeleteV1AudiencesAudienceID(params *DeleteV1AudiencesAudienceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1AudiencesAudienceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1AudiencesAudienceIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1AudiencesAudienceId",
		Method:             "DELETE",
		PathPattern:        "/v1/audiences/{audience_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1AudiencesAudienceIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1AudiencesAudienceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1AudiencesAudienceId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1Audiences lists audiences

List all audiences
*/
func (a *Client) GetV1Audiences(params *GetV1AudiencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1AudiencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1AudiencesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1Audiences",
		Method:             "GET",
		PathPattern:        "/v1/audiences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1AudiencesReader{formats: a.formats},
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
	success, ok := result.(*GetV1AudiencesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1Audiences: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1AudiencesAudienceID gets audience

Get audience details
*/
func (a *Client) GetV1AudiencesAudienceID(params *GetV1AudiencesAudienceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1AudiencesAudienceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1AudiencesAudienceIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1AudiencesAudienceId",
		Method:             "GET",
		PathPattern:        "/v1/audiences/{audience_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1AudiencesAudienceIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1AudiencesAudienceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1AudiencesAudienceId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1AudiencesAudienceIDSummariesIncidentID gets latest summary

Get the latest audience-specific summary for an incident
*/
func (a *Client) GetV1AudiencesAudienceIDSummariesIncidentID(params *GetV1AudiencesAudienceIDSummariesIncidentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1AudiencesAudienceIDSummariesIncidentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1AudiencesAudienceIDSummariesIncidentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1AudiencesAudienceIdSummariesIncidentId",
		Method:             "GET",
		PathPattern:        "/v1/audiences/{audience_id}/summaries/{incident_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1AudiencesAudienceIDSummariesIncidentIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1AudiencesAudienceIDSummariesIncidentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1AudiencesAudienceIdSummariesIncidentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1AudiencesMemberMemberIDDefault gets default audience

Get member's default audience
*/
func (a *Client) GetV1AudiencesMemberMemberIDDefault(params *GetV1AudiencesMemberMemberIDDefaultParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1AudiencesMemberMemberIDDefaultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1AudiencesMemberMemberIDDefaultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1AudiencesMemberMemberIdDefault",
		Method:             "GET",
		PathPattern:        "/v1/audiences/member/{member_id}/default",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1AudiencesMemberMemberIDDefaultReader{formats: a.formats},
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
	success, ok := result.(*GetV1AudiencesMemberMemberIDDefaultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1AudiencesMemberMemberIdDefault: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1AudiencesSummariesIncidentID lists audience summaries

List all audience summaries for an incident
*/
func (a *Client) GetV1AudiencesSummariesIncidentID(params *GetV1AudiencesSummariesIncidentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1AudiencesSummariesIncidentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1AudiencesSummariesIncidentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1AudiencesSummariesIncidentId",
		Method:             "GET",
		PathPattern:        "/v1/audiences/summaries/{incident_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1AudiencesSummariesIncidentIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1AudiencesSummariesIncidentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1AudiencesSummariesIncidentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1AudiencesAudienceID updates audience

Update an existing audience
*/
func (a *Client) PatchV1AudiencesAudienceID(params *PatchV1AudiencesAudienceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1AudiencesAudienceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1AudiencesAudienceIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1AudiencesAudienceId",
		Method:             "PATCH",
		PathPattern:        "/v1/audiences/{audience_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1AudiencesAudienceIDReader{formats: a.formats},
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
	success, ok := result.(*PatchV1AudiencesAudienceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1AudiencesAudienceId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1AudiencesAudienceIDRestore restores audience

Restore a previously archived audience
*/
func (a *Client) PatchV1AudiencesAudienceIDRestore(params *PatchV1AudiencesAudienceIDRestoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1AudiencesAudienceIDRestoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1AudiencesAudienceIDRestoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1AudiencesAudienceIdRestore",
		Method:             "PATCH",
		PathPattern:        "/v1/audiences/{audience_id}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1AudiencesAudienceIDRestoreReader{formats: a.formats},
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
	success, ok := result.(*PatchV1AudiencesAudienceIDRestoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1AudiencesAudienceIdRestore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1Audiences creates audience

Create a new audience
*/
func (a *Client) PostV1Audiences(params *PostV1AudiencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1AudiencesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1AudiencesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1Audiences",
		Method:             "POST",
		PathPattern:        "/v1/audiences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1AudiencesReader{formats: a.formats},
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
	success, ok := result.(*PostV1AudiencesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1Audiences: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1AudiencesAudienceIDSummariesIncidentID generates summary

Generate a new audience-specific summary for an incident
*/
func (a *Client) PostV1AudiencesAudienceIDSummariesIncidentID(params *PostV1AudiencesAudienceIDSummariesIncidentIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1AudiencesAudienceIDSummariesIncidentIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1AudiencesAudienceIDSummariesIncidentIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1AudiencesAudienceIdSummariesIncidentId",
		Method:             "POST",
		PathPattern:        "/v1/audiences/{audience_id}/summaries/{incident_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1AudiencesAudienceIDSummariesIncidentIDReader{formats: a.formats},
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
	success, ok := result.(*PostV1AudiencesAudienceIDSummariesIncidentIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1AudiencesAudienceIdSummariesIncidentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutV1AudiencesMemberMemberIDDefault sets default audience

Set member's default audience
*/
func (a *Client) PutV1AudiencesMemberMemberIDDefault(params *PutV1AudiencesMemberMemberIDDefaultParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutV1AudiencesMemberMemberIDDefaultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutV1AudiencesMemberMemberIDDefaultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "putV1AudiencesMemberMemberIdDefault",
		Method:             "PUT",
		PathPattern:        "/v1/audiences/member/{member_id}/default",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutV1AudiencesMemberMemberIDDefaultReader{formats: a.formats},
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
	success, ok := result.(*PutV1AudiencesMemberMemberIDDefaultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putV1AudiencesMemberMemberIdDefault: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
