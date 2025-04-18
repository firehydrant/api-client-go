// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new alerts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alerts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteV1IncidentsIncidentIDAlertsIncidentAlertID(params *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDNoContent, error)

	GetV1Alerts(params *GetV1AlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1AlertsOK, error)

	GetV1AlertsAlertID(params *GetV1AlertsAlertIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1AlertsAlertIDOK, error)

	GetV1IncidentsIncidentIDAlerts(params *GetV1IncidentsIncidentIDAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IncidentsIncidentIDAlertsOK, error)

	GetV1ProcessingLogEntries(params *GetV1ProcessingLogEntriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ProcessingLogEntriesOK, error)

	PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary(params *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK, error)

	PostV1IncidentsIncidentIDAlerts(params *PostV1IncidentsIncidentIDAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1IncidentsIncidentIDAlertsNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteV1IncidentsIncidentIDAlertsIncidentAlertID removes an alert from an incident

Remove an alert from an incident
*/
func (a *Client) DeleteV1IncidentsIncidentIDAlertsIncidentAlertID(params *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteV1IncidentsIncidentIdAlertsIncidentAlertId",
		Method:             "DELETE",
		PathPattern:        "/v1/incidents/{incident_id}/alerts/{incident_alert_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteV1IncidentsIncidentIdAlertsIncidentAlertId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1Alerts lists alerts

Retrieve all alerts from third parties
*/
func (a *Client) GetV1Alerts(params *GetV1AlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1AlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1AlertsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1Alerts",
		Method:             "GET",
		PathPattern:        "/v1/alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1AlertsReader{formats: a.formats},
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
	success, ok := result.(*GetV1AlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1Alerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1AlertsAlertID gets an alert

Retrieve a single alert
*/
func (a *Client) GetV1AlertsAlertID(params *GetV1AlertsAlertIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1AlertsAlertIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1AlertsAlertIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1AlertsAlertId",
		Method:             "GET",
		PathPattern:        "/v1/alerts/{alert_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1AlertsAlertIDReader{formats: a.formats},
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
	success, ok := result.(*GetV1AlertsAlertIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1AlertsAlertId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1IncidentsIncidentIDAlerts lists alerts for an incident

List alerts that have been attached to an incident
*/
func (a *Client) GetV1IncidentsIncidentIDAlerts(params *GetV1IncidentsIncidentIDAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1IncidentsIncidentIDAlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1IncidentsIncidentIDAlertsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1IncidentsIncidentIdAlerts",
		Method:             "GET",
		PathPattern:        "/v1/incidents/{incident_id}/alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1IncidentsIncidentIDAlertsReader{formats: a.formats},
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
	success, ok := result.(*GetV1IncidentsIncidentIDAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1IncidentsIncidentIdAlerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1ProcessingLogEntries lists alert processing log entries

Processing Log Entries for a specific alert
*/
func (a *Client) GetV1ProcessingLogEntries(params *GetV1ProcessingLogEntriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1ProcessingLogEntriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1ProcessingLogEntriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getV1ProcessingLogEntries",
		Method:             "GET",
		PathPattern:        "/v1/processing_log_entries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1ProcessingLogEntriesReader{formats: a.formats},
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
	success, ok := result.(*GetV1ProcessingLogEntriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getV1ProcessingLogEntries: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary sets an alert as primary for an incident

Setting an alert as primary will overwrite milestone times in the FireHydrant incident with times included in the primary alert. Services attached to the primary alert will also be automatically added to the incident.
*/
func (a *Client) PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary(params *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchV1IncidentsIncidentIdAlertsIncidentAlertIdPrimary",
		Method:             "PATCH",
		PathPattern:        "/v1/incidents/{incident_id}/alerts/{incident_alert_id}/primary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryReader{formats: a.formats},
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
	success, ok := result.(*PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchV1IncidentsIncidentIdAlertsIncidentAlertIdPrimary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1IncidentsIncidentIDAlerts attaches an alert to an incident

Add an alert to an incident. FireHydrant needs to have ingested the alert from a third party system in order to attach it to the incident.
*/
func (a *Client) PostV1IncidentsIncidentIDAlerts(params *PostV1IncidentsIncidentIDAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1IncidentsIncidentIDAlertsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1IncidentsIncidentIDAlertsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postV1IncidentsIncidentIdAlerts",
		Method:             "POST",
		PathPattern:        "/v1/incidents/{incident_id}/alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1IncidentsIncidentIDAlertsReader{formats: a.formats},
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
	success, ok := result.(*PostV1IncidentsIncidentIDAlertsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postV1IncidentsIncidentIdAlerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
