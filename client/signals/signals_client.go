// Code generated by go-swagger; DO NOT EDIT.

package signals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new signals API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for signals API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1SignalsEmailTargetsID(params *DeleteV1SignalsEmailTargetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1SignalsEmailTargetsIDNoContent, error)

	DeleteV1SignalsWebhookTargetsID(params *DeleteV1SignalsWebhookTargetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1SignalsWebhookTargetsIDNoContent, error)

	GetV1SignalsAnalyticsGroupedMetrics(params *GetV1SignalsAnalyticsGroupedMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsAnalyticsGroupedMetricsOK, error)

	GetV1SignalsAnalyticsMttx(params *GetV1SignalsAnalyticsMttxParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsAnalyticsMttxOK, error)

	GetV1SignalsAnalyticsShiftsExport(params *GetV1SignalsAnalyticsShiftsExportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsAnalyticsShiftsExportOK, error)

	GetV1SignalsAnalyticsTimeseries(params *GetV1SignalsAnalyticsTimeseriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsAnalyticsTimeseriesOK, error)

	GetV1SignalsEmailTargets(params *GetV1SignalsEmailTargetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsEmailTargetsOK, error)

	GetV1SignalsEmailTargetsID(params *GetV1SignalsEmailTargetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsEmailTargetsIDOK, error)

	GetV1SignalsEventSources(params *GetV1SignalsEventSourcesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsEventSourcesOK, error)

	GetV1SignalsIngestURL(params *GetV1SignalsIngestURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsIngestURLOK, error)

	GetV1SignalsTransposers(params *GetV1SignalsTransposersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsTransposersOK, error)

	GetV1SignalsWebhookTargets(params *GetV1SignalsWebhookTargetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsWebhookTargetsOK, error)

	GetV1SignalsWebhookTargetsID(params *GetV1SignalsWebhookTargetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsWebhookTargetsIDOK, error)

	PatchV1SignalsEmailTargetsID(params *PatchV1SignalsEmailTargetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1SignalsEmailTargetsIDOK, error)

	PatchV1SignalsWebhookTargetsID(params *PatchV1SignalsWebhookTargetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1SignalsWebhookTargetsIDOK, error)

	PostV1SignalsDebugger(params *PostV1SignalsDebuggerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1SignalsDebuggerCreated, error)

	PostV1SignalsEmailTargets(params *PostV1SignalsEmailTargetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1SignalsEmailTargetsCreated, error)

	PostV1SignalsWebhookTargets(params *PostV1SignalsWebhookTargetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1SignalsWebhookTargetsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteV1SignalsEmailTargetsID deletes an email target

Delete a Signals email target by ID
*/
func (a *Client) DeleteV1SignalsEmailTargetsID(params *DeleteV1SignalsEmailTargetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1SignalsEmailTargetsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1SignalsEmailTargetsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1SignalsEmailTargetsId",
		Method:             "DELETE",
		PathPattern:        "/v1/signals/email_targets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1SignalsEmailTargetsIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1SignalsEmailTargetsIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1SignalsEmailTargetsId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteV1SignalsWebhookTargetsID deletes an webhook target

Delete a Signals webhook target by ID
*/
func (a *Client) DeleteV1SignalsWebhookTargetsID(params *DeleteV1SignalsWebhookTargetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1SignalsWebhookTargetsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1SignalsWebhookTargetsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1SignalsWebhookTargetsId",
		Method:             "DELETE",
		PathPattern:        "/v1/signals/webhook_targets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1SignalsWebhookTargetsIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1SignalsWebhookTargetsIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1SignalsWebhookTargetsId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1SignalsAnalyticsGroupedMetrics generates grouped alert metrics

Generate a report of grouped metrics for Signals alerts
*/
func (a *Client) GetV1SignalsAnalyticsGroupedMetrics(params *GetV1SignalsAnalyticsGroupedMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsAnalyticsGroupedMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SignalsAnalyticsGroupedMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1SignalsAnalyticsGroupedMetrics",
		Method:             "GET",
		PathPattern:        "/v1/signals/analytics/grouped_metrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SignalsAnalyticsGroupedMetricsReader{formats: a.formats},
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
	success, ok := result.(*GetV1SignalsAnalyticsGroupedMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1SignalsAnalyticsGroupedMetrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1SignalsAnalyticsMttx gets m t t x metrics for signals alerts

Get mean-time-to-acknowledged (MTTA) and mean-time-to-resolved (MTTR) metrics for Signals alerts
*/
func (a *Client) GetV1SignalsAnalyticsMttx(params *GetV1SignalsAnalyticsMttxParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsAnalyticsMttxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SignalsAnalyticsMttxParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1SignalsAnalyticsMttx",
		Method:             "GET",
		PathPattern:        "/v1/signals/analytics/mttx",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SignalsAnalyticsMttxReader{formats: a.formats},
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
	success, ok := result.(*GetV1SignalsAnalyticsMttxOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1SignalsAnalyticsMttx: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1SignalsAnalyticsShiftsExport Export oncall hours report for given users within a time period
*/
func (a *Client) GetV1SignalsAnalyticsShiftsExport(params *GetV1SignalsAnalyticsShiftsExportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsAnalyticsShiftsExportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SignalsAnalyticsShiftsExportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1SignalsAnalyticsShiftsExport",
		Method:             "GET",
		PathPattern:        "/v1/signals/analytics/shifts/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SignalsAnalyticsShiftsExportReader{formats: a.formats},
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
	success, ok := result.(*GetV1SignalsAnalyticsShiftsExportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1SignalsAnalyticsShiftsExport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1SignalsAnalyticsTimeseries generates timeseries alert metrics

Generate a timeseries-based report of metrics for Signals alerts
*/
func (a *Client) GetV1SignalsAnalyticsTimeseries(params *GetV1SignalsAnalyticsTimeseriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsAnalyticsTimeseriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SignalsAnalyticsTimeseriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1SignalsAnalyticsTimeseries",
		Method:             "GET",
		PathPattern:        "/v1/signals/analytics/timeseries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SignalsAnalyticsTimeseriesReader{formats: a.formats},
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
	success, ok := result.(*GetV1SignalsAnalyticsTimeseriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1SignalsAnalyticsTimeseries: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1SignalsEmailTargets lists email targets

List all Signals email targets for a team.
*/
func (a *Client) GetV1SignalsEmailTargets(params *GetV1SignalsEmailTargetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsEmailTargetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SignalsEmailTargetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1SignalsEmailTargets",
		Method:             "GET",
		PathPattern:        "/v1/signals/email_targets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SignalsEmailTargetsReader{formats: a.formats},
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
	success, ok := result.(*GetV1SignalsEmailTargetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1SignalsEmailTargets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1SignalsEmailTargetsID gets an email target

Get a Signals email target by ID
*/
func (a *Client) GetV1SignalsEmailTargetsID(params *GetV1SignalsEmailTargetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsEmailTargetsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SignalsEmailTargetsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1SignalsEmailTargetsId",
		Method:             "GET",
		PathPattern:        "/v1/signals/email_targets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SignalsEmailTargetsIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1SignalsEmailTargetsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1SignalsEmailTargetsId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1SignalsEventSources get v1 signals event sources API
*/
func (a *Client) GetV1SignalsEventSources(params *GetV1SignalsEventSourcesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsEventSourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SignalsEventSourcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1SignalsEventSources",
		Method:             "GET",
		PathPattern:        "/v1/signals/event_sources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SignalsEventSourcesReader{formats: a.formats},
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
	success, ok := result.(*GetV1SignalsEventSourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1SignalsEventSources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1SignalsIngestURL retrieves the url for ingesting signals

Retrieve the url for ingesting signals for your organization
*/
func (a *Client) GetV1SignalsIngestURL(params *GetV1SignalsIngestURLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsIngestURLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SignalsIngestURLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1SignalsIngestUrl",
		Method:             "GET",
		PathPattern:        "/v1/signals/ingest_url",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SignalsIngestURLReader{formats: a.formats},
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
	success, ok := result.(*GetV1SignalsIngestURLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1SignalsIngestUrl: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1SignalsTransposers get v1 signals transposers API
*/
func (a *Client) GetV1SignalsTransposers(params *GetV1SignalsTransposersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsTransposersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SignalsTransposersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1SignalsTransposers",
		Method:             "GET",
		PathPattern:        "/v1/signals/transposers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SignalsTransposersReader{formats: a.formats},
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
	success, ok := result.(*GetV1SignalsTransposersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1SignalsTransposers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1SignalsWebhookTargets lists webhook targets

List all Signals webhook targets.
*/
func (a *Client) GetV1SignalsWebhookTargets(params *GetV1SignalsWebhookTargetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsWebhookTargetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SignalsWebhookTargetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1SignalsWebhookTargets",
		Method:             "GET",
		PathPattern:        "/v1/signals/webhook_targets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SignalsWebhookTargetsReader{formats: a.formats},
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
	success, ok := result.(*GetV1SignalsWebhookTargetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1SignalsWebhookTargets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1SignalsWebhookTargetsID gets an webhook target

Get a Signals webhook target by ID
*/
func (a *Client) GetV1SignalsWebhookTargetsID(params *GetV1SignalsWebhookTargetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsWebhookTargetsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SignalsWebhookTargetsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1SignalsWebhookTargetsId",
		Method:             "GET",
		PathPattern:        "/v1/signals/webhook_targets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SignalsWebhookTargetsIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1SignalsWebhookTargetsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1SignalsWebhookTargetsId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1SignalsEmailTargetsID updates an email target

Update a Signals email target by ID
*/
func (a *Client) PatchV1SignalsEmailTargetsID(params *PatchV1SignalsEmailTargetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1SignalsEmailTargetsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1SignalsEmailTargetsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1SignalsEmailTargetsId",
		Method:             "PATCH",
		PathPattern:        "/v1/signals/email_targets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1SignalsEmailTargetsIDReader{formats: a.formats},
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
	success, ok := result.(*PatchV1SignalsEmailTargetsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1SignalsEmailTargetsId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1SignalsWebhookTargetsID updates an webhook target

Update a Signals webhook target by ID
*/
func (a *Client) PatchV1SignalsWebhookTargetsID(params *PatchV1SignalsWebhookTargetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1SignalsWebhookTargetsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1SignalsWebhookTargetsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1SignalsWebhookTargetsId",
		Method:             "PATCH",
		PathPattern:        "/v1/signals/webhook_targets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1SignalsWebhookTargetsIDReader{formats: a.formats},
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
	success, ok := result.(*PatchV1SignalsWebhookTargetsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1SignalsWebhookTargetsId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1SignalsDebugger post v1 signals debugger API
*/
func (a *Client) PostV1SignalsDebugger(params *PostV1SignalsDebuggerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1SignalsDebuggerCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1SignalsDebuggerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1SignalsDebugger",
		Method:             "POST",
		PathPattern:        "/v1/signals/debugger",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1SignalsDebuggerReader{formats: a.formats},
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
	success, ok := result.(*PostV1SignalsDebuggerCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1SignalsDebugger: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1SignalsEmailTargets creates an email target

Create a Signals email target for a team.
*/
func (a *Client) PostV1SignalsEmailTargets(params *PostV1SignalsEmailTargetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1SignalsEmailTargetsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1SignalsEmailTargetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1SignalsEmailTargets",
		Method:             "POST",
		PathPattern:        "/v1/signals/email_targets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1SignalsEmailTargetsReader{formats: a.formats},
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
	success, ok := result.(*PostV1SignalsEmailTargetsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1SignalsEmailTargets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1SignalsWebhookTargets creates an webhook target

Create a Signals webhook target.
*/
func (a *Client) PostV1SignalsWebhookTargets(params *PostV1SignalsWebhookTargetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1SignalsWebhookTargetsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1SignalsWebhookTargetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1SignalsWebhookTargets",
		Method:             "POST",
		PathPattern:        "/v1/signals/webhook_targets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1SignalsWebhookTargetsReader{formats: a.formats},
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
	success, ok := result.(*PostV1SignalsWebhookTargetsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1SignalsWebhookTargets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
