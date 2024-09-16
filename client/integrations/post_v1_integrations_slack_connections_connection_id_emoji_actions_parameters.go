// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams creates a new PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams() *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams {
	return &PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParamsWithTimeout creates a new PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams object
// with the ability to set a timeout on a request.
func NewPostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParamsWithTimeout(timeout time.Duration) *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams {
	return &PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams{
		timeout: timeout,
	}
}

// NewPostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParamsWithContext creates a new PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams object
// with the ability to set a context for a request.
func NewPostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParamsWithContext(ctx context.Context) *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams {
	return &PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams{
		Context: ctx,
	}
}

// NewPostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParamsWithHTTPClient creates a new PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParamsWithHTTPClient(client *http.Client) *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams {
	return &PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams{
		HTTPClient: client,
	}
}

/*
PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams contains all the parameters to send to the API endpoint

	for the post v1 integrations slack connections connection Id emoji actions operation.

	Typically these are written to a http.Request.
*/
type PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams struct {

	/* ConnectionID.

	   Slack Connection UUID
	*/
	ConnectionID string

	/* EmojiName.

	   The name of the emoji to associate with this action
	*/
	EmojiName string

	/* IncidentTypeID.

	   The ID of the incident type to associate with this emoji action
	*/
	IncidentTypeID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 integrations slack connections connection Id emoji actions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) WithDefaults() *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 integrations slack connections connection Id emoji actions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 integrations slack connections connection Id emoji actions params
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) WithTimeout(timeout time.Duration) *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 integrations slack connections connection Id emoji actions params
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 integrations slack connections connection Id emoji actions params
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) WithContext(ctx context.Context) *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 integrations slack connections connection Id emoji actions params
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 integrations slack connections connection Id emoji actions params
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) WithHTTPClient(client *http.Client) *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 integrations slack connections connection Id emoji actions params
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the post v1 integrations slack connections connection Id emoji actions params
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) WithConnectionID(connectionID string) *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the post v1 integrations slack connections connection Id emoji actions params
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) SetConnectionID(connectionID string) {
	o.ConnectionID = connectionID
}

// WithEmojiName adds the emojiName to the post v1 integrations slack connections connection Id emoji actions params
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) WithEmojiName(emojiName string) *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams {
	o.SetEmojiName(emojiName)
	return o
}

// SetEmojiName adds the emojiName to the post v1 integrations slack connections connection Id emoji actions params
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) SetEmojiName(emojiName string) {
	o.EmojiName = emojiName
}

// WithIncidentTypeID adds the incidentTypeID to the post v1 integrations slack connections connection Id emoji actions params
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) WithIncidentTypeID(incidentTypeID *string) *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams {
	o.SetIncidentTypeID(incidentTypeID)
	return o
}

// SetIncidentTypeID adds the incidentTypeId to the post v1 integrations slack connections connection Id emoji actions params
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) SetIncidentTypeID(incidentTypeID *string) {
	o.IncidentTypeID = incidentTypeID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1IntegrationsSlackConnectionsConnectionIDEmojiActionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_id
	if err := r.SetPathParam("connection_id", o.ConnectionID); err != nil {
		return err
	}

	// form param emoji_name
	frEmojiName := o.EmojiName
	fEmojiName := frEmojiName
	if fEmojiName != "" {
		if err := r.SetFormParam("emoji_name", fEmojiName); err != nil {
			return err
		}
	}

	if o.IncidentTypeID != nil {

		// form param incident_type_id
		var frIncidentTypeID string
		if o.IncidentTypeID != nil {
			frIncidentTypeID = *o.IncidentTypeID
		}
		fIncidentTypeID := frIncidentTypeID
		if fIncidentTypeID != "" {
			if err := r.SetFormParam("incident_type_id", fIncidentTypeID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}