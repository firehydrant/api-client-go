// Code generated by go-swagger; DO NOT EDIT.

package reporting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new reporting API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reporting API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1SavedSearchesResourceTypeSavedSearchID(params *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1SavedSearchesResourceTypeSavedSearchIDOK, error)

	GetV1MetricsIncidents(params *GetV1MetricsIncidentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsIncidentsOK, error)

	GetV1MetricsInfraType(params *GetV1MetricsInfraTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsInfraTypeOK, error)

	GetV1MetricsInfraTypeInfraID(params *GetV1MetricsInfraTypeInfraIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsInfraTypeInfraIDOK, error)

	GetV1MetricsMilestoneFunnel(params *GetV1MetricsMilestoneFunnelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsMilestoneFunnelOK, error)

	GetV1MetricsMttx(params *GetV1MetricsMttxParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsMttxOK, error)

	GetV1MetricsRetrospectives(params *GetV1MetricsRetrospectivesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsRetrospectivesOK, error)

	GetV1MetricsTicketFunnel(params *GetV1MetricsTicketFunnelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsTicketFunnelOK, error)

	GetV1MetricsUserInvolvements(params *GetV1MetricsUserInvolvementsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsUserInvolvementsOK, error)

	GetV1ReportsMeanTime(params *GetV1ReportsMeanTimeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ReportsMeanTimeOK, error)

	GetV1SavedSearchesResourceType(params *GetV1SavedSearchesResourceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SavedSearchesResourceTypeOK, error)

	GetV1SavedSearchesResourceTypeSavedSearchID(params *GetV1SavedSearchesResourceTypeSavedSearchIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SavedSearchesResourceTypeSavedSearchIDOK, error)

	GetV1SignalsAnalyticsGroupedMetrics(params *GetV1SignalsAnalyticsGroupedMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsAnalyticsGroupedMetricsOK, error)

	GetV1SignalsAnalyticsMttx(params *GetV1SignalsAnalyticsMttxParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsAnalyticsMttxOK, error)

	GetV1SignalsAnalyticsShiftsExport(params *GetV1SignalsAnalyticsShiftsExportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsAnalyticsShiftsExportOK, error)

	GetV1SignalsAnalyticsTimeseries(params *GetV1SignalsAnalyticsTimeseriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SignalsAnalyticsTimeseriesOK, error)

	PatchV1SavedSearchesResourceTypeSavedSearchID(params *PatchV1SavedSearchesResourceTypeSavedSearchIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1SavedSearchesResourceTypeSavedSearchIDOK, error)

	PostV1SavedSearchesResourceType(params *PostV1SavedSearchesResourceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1SavedSearchesResourceTypeCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteV1SavedSearchesResourceTypeSavedSearchID deletes a saved search

Delete a specific saved search
*/
func (a *Client) DeleteV1SavedSearchesResourceTypeSavedSearchID(params *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1SavedSearchesResourceTypeSavedSearchIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1SavedSearchesResourceTypeSavedSearchIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1SavedSearchesResourceTypeSavedSearchId",
		Method:             "DELETE",
		PathPattern:        "/v1/saved_searches/{resource_type}/{saved_search_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1SavedSearchesResourceTypeSavedSearchIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1SavedSearchesResourceTypeSavedSearchIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1SavedSearchesResourceTypeSavedSearchId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1MetricsIncidents lists incident metrics and analytics

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
GetV1MetricsInfraTypeInfraID gets metrics for a component

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
GetV1MetricsMilestoneFunnel lists milestone funnel metrics

Returns a report with time bucketed milestone data
*/
func (a *Client) GetV1MetricsMilestoneFunnel(params *GetV1MetricsMilestoneFunnelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsMilestoneFunnelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1MetricsMilestoneFunnelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1MetricsMilestoneFunnel",
		Method:             "GET",
		PathPattern:        "/v1/metrics/milestone_funnel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1MetricsMilestoneFunnelReader{formats: a.formats},
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
	success, ok := result.(*GetV1MetricsMilestoneFunnelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1MetricsMilestoneFunnel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1MetricsMttx gets infrastructure metrics

Fetch infrastructure metrics based on custom query
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
GetV1MetricsTicketFunnel lists ticket task and follow up creation and completion metrics

Returns a report with task and follow up creation and completion data
*/
func (a *Client) GetV1MetricsTicketFunnel(params *GetV1MetricsTicketFunnelParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1MetricsTicketFunnelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1MetricsTicketFunnelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1MetricsTicketFunnel",
		Method:             "GET",
		PathPattern:        "/v1/metrics/ticket_funnel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1MetricsTicketFunnelReader{formats: a.formats},
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
	success, ok := result.(*GetV1MetricsTicketFunnelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1MetricsTicketFunnel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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

/*
GetV1ReportsMeanTime gets mean time metrics for incidents

Returns a report with time bucketed analytics data
*/
func (a *Client) GetV1ReportsMeanTime(params *GetV1ReportsMeanTimeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ReportsMeanTimeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ReportsMeanTimeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1ReportsMeanTime",
		Method:             "GET",
		PathPattern:        "/v1/reports/mean_time",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ReportsMeanTimeReader{formats: a.formats},
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
	success, ok := result.(*GetV1ReportsMeanTimeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1ReportsMeanTime: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1SavedSearchesResourceType lists saved searches

Lists saved searches
*/
func (a *Client) GetV1SavedSearchesResourceType(params *GetV1SavedSearchesResourceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SavedSearchesResourceTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SavedSearchesResourceTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1SavedSearchesResourceType",
		Method:             "GET",
		PathPattern:        "/v1/saved_searches/{resource_type}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SavedSearchesResourceTypeReader{formats: a.formats},
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
	success, ok := result.(*GetV1SavedSearchesResourceTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1SavedSearchesResourceType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1SavedSearchesResourceTypeSavedSearchID gets a saved search

Retrieve a specific save search
*/
func (a *Client) GetV1SavedSearchesResourceTypeSavedSearchID(params *GetV1SavedSearchesResourceTypeSavedSearchIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1SavedSearchesResourceTypeSavedSearchIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1SavedSearchesResourceTypeSavedSearchIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1SavedSearchesResourceTypeSavedSearchId",
		Method:             "GET",
		PathPattern:        "/v1/saved_searches/{resource_type}/{saved_search_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1SavedSearchesResourceTypeSavedSearchIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1SavedSearchesResourceTypeSavedSearchIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1SavedSearchesResourceTypeSavedSearchId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
GetV1SignalsAnalyticsMttx gets m t t x analytics for signals

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
GetV1SignalsAnalyticsShiftsExport exports on call hours report

Export on-call hours report for users/teams during a time period
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
PatchV1SavedSearchesResourceTypeSavedSearchID updates a saved search

Update a specific saved search
*/
func (a *Client) PatchV1SavedSearchesResourceTypeSavedSearchID(params *PatchV1SavedSearchesResourceTypeSavedSearchIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1SavedSearchesResourceTypeSavedSearchIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1SavedSearchesResourceTypeSavedSearchIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1SavedSearchesResourceTypeSavedSearchId",
		Method:             "PATCH",
		PathPattern:        "/v1/saved_searches/{resource_type}/{saved_search_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1SavedSearchesResourceTypeSavedSearchIDReader{formats: a.formats},
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
	success, ok := result.(*PatchV1SavedSearchesResourceTypeSavedSearchIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1SavedSearchesResourceTypeSavedSearchId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1SavedSearchesResourceType creates a saved search

Create a new saved search for a particular resource type
*/
func (a *Client) PostV1SavedSearchesResourceType(params *PostV1SavedSearchesResourceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1SavedSearchesResourceTypeCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1SavedSearchesResourceTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1SavedSearchesResourceType",
		Method:             "POST",
		PathPattern:        "/v1/saved_searches/{resource_type}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1SavedSearchesResourceTypeReader{formats: a.formats},
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
	success, ok := result.(*PostV1SavedSearchesResourceTypeCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1SavedSearchesResourceType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
