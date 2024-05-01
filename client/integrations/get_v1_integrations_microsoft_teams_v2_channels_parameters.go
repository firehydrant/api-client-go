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

// NewGetV1IntegrationsMicrosoftTeamsV2ChannelsParams creates a new GetV1IntegrationsMicrosoftTeamsV2ChannelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IntegrationsMicrosoftTeamsV2ChannelsParams() *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams {
	return &GetV1IntegrationsMicrosoftTeamsV2ChannelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IntegrationsMicrosoftTeamsV2ChannelsParamsWithTimeout creates a new GetV1IntegrationsMicrosoftTeamsV2ChannelsParams object
// with the ability to set a timeout on a request.
func NewGetV1IntegrationsMicrosoftTeamsV2ChannelsParamsWithTimeout(timeout time.Duration) *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams {
	return &GetV1IntegrationsMicrosoftTeamsV2ChannelsParams{
		timeout: timeout,
	}
}

// NewGetV1IntegrationsMicrosoftTeamsV2ChannelsParamsWithContext creates a new GetV1IntegrationsMicrosoftTeamsV2ChannelsParams object
// with the ability to set a context for a request.
func NewGetV1IntegrationsMicrosoftTeamsV2ChannelsParamsWithContext(ctx context.Context) *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams {
	return &GetV1IntegrationsMicrosoftTeamsV2ChannelsParams{
		Context: ctx,
	}
}

// NewGetV1IntegrationsMicrosoftTeamsV2ChannelsParamsWithHTTPClient creates a new GetV1IntegrationsMicrosoftTeamsV2ChannelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IntegrationsMicrosoftTeamsV2ChannelsParamsWithHTTPClient(client *http.Client) *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams {
	return &GetV1IntegrationsMicrosoftTeamsV2ChannelsParams{
		HTTPClient: client,
	}
}

/*
GetV1IntegrationsMicrosoftTeamsV2ChannelsParams contains all the parameters to send to the API endpoint

	for the get v1 integrations microsoft teams v2 channels operation.

	Typically these are written to a http.Request.
*/
type GetV1IntegrationsMicrosoftTeamsV2ChannelsParams struct {

	/* MsChannelID.

	   Microsoft Teams Channel ID
	*/
	MsChannelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 integrations microsoft teams v2 channels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams) WithDefaults() *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 integrations microsoft teams v2 channels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 integrations microsoft teams v2 channels params
func (o *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams) WithTimeout(timeout time.Duration) *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 integrations microsoft teams v2 channels params
func (o *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 integrations microsoft teams v2 channels params
func (o *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams) WithContext(ctx context.Context) *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 integrations microsoft teams v2 channels params
func (o *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 integrations microsoft teams v2 channels params
func (o *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams) WithHTTPClient(client *http.Client) *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 integrations microsoft teams v2 channels params
func (o *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsChannelID adds the msChannelID to the get v1 integrations microsoft teams v2 channels params
func (o *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams) WithMsChannelID(msChannelID string) *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams {
	o.SetMsChannelID(msChannelID)
	return o
}

// SetMsChannelID adds the msChannelId to the get v1 integrations microsoft teams v2 channels params
func (o *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams) SetMsChannelID(msChannelID string) {
	o.MsChannelID = msChannelID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IntegrationsMicrosoftTeamsV2ChannelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param ms_channel_id
	qrMsChannelID := o.MsChannelID
	qMsChannelID := qrMsChannelID
	if qMsChannelID != "" {

		if err := r.SetQueryParam("ms_channel_id", qMsChannelID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}