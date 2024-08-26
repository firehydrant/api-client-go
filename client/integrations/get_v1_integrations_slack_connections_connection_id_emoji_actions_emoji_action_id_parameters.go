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

// NewGetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams creates a new GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams() *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams {
	return &GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParamsWithTimeout creates a new GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams object
// with the ability to set a timeout on a request.
func NewGetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParamsWithTimeout(timeout time.Duration) *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams {
	return &GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams{
		timeout: timeout,
	}
}

// NewGetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParamsWithContext creates a new GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams object
// with the ability to set a context for a request.
func NewGetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParamsWithContext(ctx context.Context) *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams {
	return &GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams{
		Context: ctx,
	}
}

// NewGetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParamsWithHTTPClient creates a new GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParamsWithHTTPClient(client *http.Client) *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams {
	return &GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams{
		HTTPClient: client,
	}
}

/*
GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams contains all the parameters to send to the API endpoint

	for the get v1 integrations slack connections connection Id emoji actions emoji action Id operation.

	Typically these are written to a http.Request.
*/
type GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams struct {

	/* ConnectionID.

	   Slack Connection UUID
	*/
	ConnectionID string

	// EmojiActionID.
	EmojiActionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 integrations slack connections connection Id emoji actions emoji action Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams) WithDefaults() *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 integrations slack connections connection Id emoji actions emoji action Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 integrations slack connections connection Id emoji actions emoji action Id params
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams) WithTimeout(timeout time.Duration) *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 integrations slack connections connection Id emoji actions emoji action Id params
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 integrations slack connections connection Id emoji actions emoji action Id params
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams) WithContext(ctx context.Context) *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 integrations slack connections connection Id emoji actions emoji action Id params
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 integrations slack connections connection Id emoji actions emoji action Id params
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams) WithHTTPClient(client *http.Client) *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 integrations slack connections connection Id emoji actions emoji action Id params
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the get v1 integrations slack connections connection Id emoji actions emoji action Id params
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams) WithConnectionID(connectionID string) *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the get v1 integrations slack connections connection Id emoji actions emoji action Id params
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams) SetConnectionID(connectionID string) {
	o.ConnectionID = connectionID
}

// WithEmojiActionID adds the emojiActionID to the get v1 integrations slack connections connection Id emoji actions emoji action Id params
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams) WithEmojiActionID(emojiActionID string) *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams {
	o.SetEmojiActionID(emojiActionID)
	return o
}

// SetEmojiActionID adds the emojiActionId to the get v1 integrations slack connections connection Id emoji actions emoji action Id params
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams) SetEmojiActionID(emojiActionID string) {
	o.EmojiActionID = emojiActionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_id
	if err := r.SetPathParam("connection_id", o.ConnectionID); err != nil {
		return err
	}

	// path param emoji_action_id
	if err := r.SetPathParam("emoji_action_id", o.EmojiActionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
