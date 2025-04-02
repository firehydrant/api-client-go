// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1ServiceDependenciesServiceDependencyID(params *DeleteV1ServiceDependenciesServiceDependencyIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ServiceDependenciesServiceDependencyIDOK, error)

	DeleteV1ServicesServiceID(params *DeleteV1ServicesServiceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ServicesServiceIDOK, error)

	DeleteV1ServicesServiceIDServiceLinksRemoteID(params *DeleteV1ServicesServiceIDServiceLinksRemoteIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ServicesServiceIDServiceLinksRemoteIDNoContent, error)

	GetV1CatalogsCatalogIDRefresh(params *GetV1CatalogsCatalogIDRefreshParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1CatalogsCatalogIDRefreshOK, error)

	GetV1Infrastructures(params *GetV1InfrastructuresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1InfrastructuresOK, error)

	GetV1ServiceDependenciesServiceDependencyID(params *GetV1ServiceDependenciesServiceDependencyIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServiceDependenciesServiceDependencyIDOK, error)

	GetV1Services(params *GetV1ServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServicesOK, error)

	GetV1ServicesServiceID(params *GetV1ServicesServiceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServicesServiceIDOK, error)

	GetV1ServicesServiceIDAvailableDownstreamDependencies(params *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServicesServiceIDAvailableDownstreamDependenciesOK, error)

	GetV1ServicesServiceIDAvailableUpstreamDependencies(params *GetV1ServicesServiceIDAvailableUpstreamDependenciesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServicesServiceIDAvailableUpstreamDependenciesOK, error)

	GetV1ServicesServiceIDDependencies(params *GetV1ServicesServiceIDDependenciesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServicesServiceIDDependenciesOK, error)

	GetV1UsersIDServices(params *GetV1UsersIDServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1UsersIDServicesOK, error)

	PatchV1ServiceDependenciesServiceDependencyID(params *PatchV1ServiceDependenciesServiceDependencyIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ServiceDependenciesServiceDependencyIDOK, error)

	PatchV1ServicesServiceID(params *PatchV1ServicesServiceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ServicesServiceIDOK, error)

	PostV1CatalogsCatalogIDIngest(params *PostV1CatalogsCatalogIDIngestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1CatalogsCatalogIDIngestCreated, error)

	PostV1ServiceDependencies(params *PostV1ServiceDependenciesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ServiceDependenciesCreated, error)

	PostV1Services(params *PostV1ServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ServicesCreated, error)

	PostV1ServicesServiceIDChecklistResponseChecklistID(params *PostV1ServicesServiceIDChecklistResponseChecklistIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ServicesServiceIDChecklistResponseChecklistIDCreated, error)

	PostV1ServicesServiceLinks(params *PostV1ServicesServiceLinksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ServicesServiceLinksCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteV1ServiceDependenciesServiceDependencyID deletes a service dependency

Deletes a single service dependency
*/
func (a *Client) DeleteV1ServiceDependenciesServiceDependencyID(params *DeleteV1ServiceDependenciesServiceDependencyIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ServiceDependenciesServiceDependencyIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1ServiceDependenciesServiceDependencyIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1ServiceDependenciesServiceDependencyId",
		Method:             "DELETE",
		PathPattern:        "/v1/service_dependencies/{service_dependency_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1ServiceDependenciesServiceDependencyIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1ServiceDependenciesServiceDependencyIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1ServiceDependenciesServiceDependencyId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteV1ServicesServiceID deletes a service

Deletes the service from FireHydrant.
*/
func (a *Client) DeleteV1ServicesServiceID(params *DeleteV1ServicesServiceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ServicesServiceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1ServicesServiceIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1ServicesServiceId",
		Method:             "DELETE",
		PathPattern:        "/v1/services/{service_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1ServicesServiceIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1ServicesServiceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1ServicesServiceId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteV1ServicesServiceIDServiceLinksRemoteID deletes a service link

Deletes a service link from FireHydrant.
*/
func (a *Client) DeleteV1ServicesServiceIDServiceLinksRemoteID(params *DeleteV1ServicesServiceIDServiceLinksRemoteIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1ServicesServiceIDServiceLinksRemoteIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1ServicesServiceIDServiceLinksRemoteIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1ServicesServiceIdServiceLinksRemoteId",
		Method:             "DELETE",
		PathPattern:        "/v1/services/{service_id}/service_links/{remote_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1ServicesServiceIDServiceLinksRemoteIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1ServicesServiceIDServiceLinksRemoteIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1ServicesServiceIdServiceLinksRemoteId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1CatalogsCatalogIDRefresh refreshes a service catalog

Schedules an async task to re-import catalog info and update catalog data accordingly.
*/
func (a *Client) GetV1CatalogsCatalogIDRefresh(params *GetV1CatalogsCatalogIDRefreshParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1CatalogsCatalogIDRefreshOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1CatalogsCatalogIDRefreshParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1CatalogsCatalogIdRefresh",
		Method:             "GET",
		PathPattern:        "/v1/catalogs/{catalog_id}/refresh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1CatalogsCatalogIDRefreshReader{formats: a.formats},
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
	success, ok := result.(*GetV1CatalogsCatalogIDRefreshOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1CatalogsCatalogIdRefresh: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1Infrastructures lists functionality service and environment objects

Lists functionality, service and environment objects
*/
func (a *Client) GetV1Infrastructures(params *GetV1InfrastructuresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1InfrastructuresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1InfrastructuresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1Infrastructures",
		Method:             "GET",
		PathPattern:        "/v1/infrastructures",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1InfrastructuresReader{formats: a.formats},
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
	success, ok := result.(*GetV1InfrastructuresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1Infrastructures: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ServiceDependenciesServiceDependencyID gets a service dependency

Retrieves a single service dependency by ID
*/
func (a *Client) GetV1ServiceDependenciesServiceDependencyID(params *GetV1ServiceDependenciesServiceDependencyIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServiceDependenciesServiceDependencyIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ServiceDependenciesServiceDependencyIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1ServiceDependenciesServiceDependencyId",
		Method:             "GET",
		PathPattern:        "/v1/service_dependencies/{service_dependency_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ServiceDependenciesServiceDependencyIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1ServiceDependenciesServiceDependencyIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1ServiceDependenciesServiceDependencyId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1Services lists services

List all of the services that have been added to the organization.
*/
func (a *Client) GetV1Services(params *GetV1ServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ServicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1Services",
		Method:             "GET",
		PathPattern:        "/v1/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ServicesReader{formats: a.formats},
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
	success, ok := result.(*GetV1ServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1Services: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ServicesServiceID gets a service

Retrieves a single service by ID
*/
func (a *Client) GetV1ServicesServiceID(params *GetV1ServicesServiceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServicesServiceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ServicesServiceIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1ServicesServiceId",
		Method:             "GET",
		PathPattern:        "/v1/services/{service_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ServicesServiceIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1ServicesServiceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1ServicesServiceId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ServicesServiceIDAvailableDownstreamDependencies lists available downstream service dependencies

Retrieves all services that are available to be downstream dependencies
*/
func (a *Client) GetV1ServicesServiceIDAvailableDownstreamDependencies(params *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServicesServiceIDAvailableDownstreamDependenciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ServicesServiceIDAvailableDownstreamDependenciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1ServicesServiceIdAvailableDownstreamDependencies",
		Method:             "GET",
		PathPattern:        "/v1/services/{service_id}/available_downstream_dependencies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ServicesServiceIDAvailableDownstreamDependenciesReader{formats: a.formats},
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
	success, ok := result.(*GetV1ServicesServiceIDAvailableDownstreamDependenciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1ServicesServiceIdAvailableDownstreamDependencies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ServicesServiceIDAvailableUpstreamDependencies lists available upstream service dependencies

Retrieves all services that are available to be upstream dependencies
*/
func (a *Client) GetV1ServicesServiceIDAvailableUpstreamDependencies(params *GetV1ServicesServiceIDAvailableUpstreamDependenciesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServicesServiceIDAvailableUpstreamDependenciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ServicesServiceIDAvailableUpstreamDependenciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1ServicesServiceIdAvailableUpstreamDependencies",
		Method:             "GET",
		PathPattern:        "/v1/services/{service_id}/available_upstream_dependencies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ServicesServiceIDAvailableUpstreamDependenciesReader{formats: a.formats},
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
	success, ok := result.(*GetV1ServicesServiceIDAvailableUpstreamDependenciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1ServicesServiceIdAvailableUpstreamDependencies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ServicesServiceIDDependencies lists dependencies for a service

Retrieves a service's dependencies
*/
func (a *Client) GetV1ServicesServiceIDDependencies(params *GetV1ServicesServiceIDDependenciesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ServicesServiceIDDependenciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ServicesServiceIDDependenciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1ServicesServiceIdDependencies",
		Method:             "GET",
		PathPattern:        "/v1/services/{service_id}/dependencies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ServicesServiceIDDependenciesReader{formats: a.formats},
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
	success, ok := result.(*GetV1ServicesServiceIDDependenciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1ServicesServiceIdDependencies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1UsersIDServices lists services owned by a user s teams

Retrieves a list of services owned by the teams a user is on
*/
func (a *Client) GetV1UsersIDServices(params *GetV1UsersIDServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1UsersIDServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1UsersIDServicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1UsersIdServices",
		Method:             "GET",
		PathPattern:        "/v1/users/{id}/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1UsersIDServicesReader{formats: a.formats},
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
	success, ok := result.(*GetV1UsersIDServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1UsersIdServices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1ServiceDependenciesServiceDependencyID updates a service dependency

Update the notes of the service dependency
*/
func (a *Client) PatchV1ServiceDependenciesServiceDependencyID(params *PatchV1ServiceDependenciesServiceDependencyIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ServiceDependenciesServiceDependencyIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1ServiceDependenciesServiceDependencyIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1ServiceDependenciesServiceDependencyId",
		Method:             "PATCH",
		PathPattern:        "/v1/service_dependencies/{service_dependency_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1ServiceDependenciesServiceDependencyIDReader{formats: a.formats},
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
	success, ok := result.(*PatchV1ServiceDependenciesServiceDependencyIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1ServiceDependenciesServiceDependencyId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PatchV1ServicesServiceID updates a service

	Update a services attributes, you may also add or remove functionalities from the service as well.

Note: You may not remove or add individual label key/value pairs. You must include the entire object to override label values.
*/
func (a *Client) PatchV1ServicesServiceID(params *PatchV1ServicesServiceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1ServicesServiceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1ServicesServiceIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1ServicesServiceId",
		Method:             "PATCH",
		PathPattern:        "/v1/services/{service_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1ServicesServiceIDReader{formats: a.formats},
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
	success, ok := result.(*PatchV1ServicesServiceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1ServicesServiceId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1CatalogsCatalogIDIngest ingests service catalog data

Accepts catalog data in the configured format and asyncronously processes the data to incorporate changes into service catalog.
*/
func (a *Client) PostV1CatalogsCatalogIDIngest(params *PostV1CatalogsCatalogIDIngestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1CatalogsCatalogIDIngestCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1CatalogsCatalogIDIngestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1CatalogsCatalogIdIngest",
		Method:             "POST",
		PathPattern:        "/v1/catalogs/{catalog_id}/ingest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1CatalogsCatalogIDIngestReader{formats: a.formats},
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
	success, ok := result.(*PostV1CatalogsCatalogIDIngestCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1CatalogsCatalogIdIngest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1ServiceDependencies creates a service dependency

Creates a service dependency relationship between two services
*/
func (a *Client) PostV1ServiceDependencies(params *PostV1ServiceDependenciesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ServiceDependenciesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1ServiceDependenciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1ServiceDependencies",
		Method:             "POST",
		PathPattern:        "/v1/service_dependencies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1ServiceDependenciesReader{formats: a.formats},
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
	success, ok := result.(*PostV1ServiceDependenciesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1ServiceDependencies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1Services creates a service

Creates a service for the organization, you may also create or attach functionalities to the service on create.
*/
func (a *Client) PostV1Services(params *PostV1ServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ServicesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1ServicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1Services",
		Method:             "POST",
		PathPattern:        "/v1/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1ServicesReader{formats: a.formats},
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
	success, ok := result.(*PostV1ServicesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1Services: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1ServicesServiceIDChecklistResponseChecklistID records a response for a checklist item

Creates a response for a checklist item
*/
func (a *Client) PostV1ServicesServiceIDChecklistResponseChecklistID(params *PostV1ServicesServiceIDChecklistResponseChecklistIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ServicesServiceIDChecklistResponseChecklistIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1ServicesServiceIDChecklistResponseChecklistIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1ServicesServiceIdChecklistResponseChecklistId",
		Method:             "POST",
		PathPattern:        "/v1/services/{service_id}/checklist_response/{checklist_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1ServicesServiceIDChecklistResponseChecklistIDReader{formats: a.formats},
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
	success, ok := result.(*PostV1ServicesServiceIDChecklistResponseChecklistIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1ServicesServiceIdChecklistResponseChecklistId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1ServicesServiceLinks creates multiple services linked to external services

Creates a service with the appropriate integration for each external service ID passed
*/
func (a *Client) PostV1ServicesServiceLinks(params *PostV1ServicesServiceLinksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1ServicesServiceLinksCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1ServicesServiceLinksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1ServicesServiceLinks",
		Method:             "POST",
		PathPattern:        "/v1/services/service_links",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1ServicesServiceLinksReader{formats: a.formats},
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
	success, ok := result.(*PostV1ServicesServiceLinksCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1ServicesServiceLinks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
