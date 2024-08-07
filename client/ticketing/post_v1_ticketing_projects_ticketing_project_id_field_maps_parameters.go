// Code generated by go-swagger; DO NOT EDIT.

package ticketing

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

// NewPostV1TicketingProjectsTicketingProjectIDFieldMapsParams creates a new PostV1TicketingProjectsTicketingProjectIDFieldMapsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1TicketingProjectsTicketingProjectIDFieldMapsParams() *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams {
	return &PostV1TicketingProjectsTicketingProjectIDFieldMapsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1TicketingProjectsTicketingProjectIDFieldMapsParamsWithTimeout creates a new PostV1TicketingProjectsTicketingProjectIDFieldMapsParams object
// with the ability to set a timeout on a request.
func NewPostV1TicketingProjectsTicketingProjectIDFieldMapsParamsWithTimeout(timeout time.Duration) *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams {
	return &PostV1TicketingProjectsTicketingProjectIDFieldMapsParams{
		timeout: timeout,
	}
}

// NewPostV1TicketingProjectsTicketingProjectIDFieldMapsParamsWithContext creates a new PostV1TicketingProjectsTicketingProjectIDFieldMapsParams object
// with the ability to set a context for a request.
func NewPostV1TicketingProjectsTicketingProjectIDFieldMapsParamsWithContext(ctx context.Context) *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams {
	return &PostV1TicketingProjectsTicketingProjectIDFieldMapsParams{
		Context: ctx,
	}
}

// NewPostV1TicketingProjectsTicketingProjectIDFieldMapsParamsWithHTTPClient creates a new PostV1TicketingProjectsTicketingProjectIDFieldMapsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1TicketingProjectsTicketingProjectIDFieldMapsParamsWithHTTPClient(client *http.Client) *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams {
	return &PostV1TicketingProjectsTicketingProjectIDFieldMapsParams{
		HTTPClient: client,
	}
}

/*
PostV1TicketingProjectsTicketingProjectIDFieldMapsParams contains all the parameters to send to the API endpoint

	for the post v1 ticketing projects ticketing project Id field maps operation.

	Typically these are written to a http.Request.
*/
type PostV1TicketingProjectsTicketingProjectIDFieldMapsParams struct {

	// TicketingProjectID.
	TicketingProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 ticketing projects ticketing project Id field maps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams) WithDefaults() *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 ticketing projects ticketing project Id field maps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 ticketing projects ticketing project Id field maps params
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams) WithTimeout(timeout time.Duration) *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 ticketing projects ticketing project Id field maps params
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 ticketing projects ticketing project Id field maps params
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams) WithContext(ctx context.Context) *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 ticketing projects ticketing project Id field maps params
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 ticketing projects ticketing project Id field maps params
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams) WithHTTPClient(client *http.Client) *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 ticketing projects ticketing project Id field maps params
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTicketingProjectID adds the ticketingProjectID to the post v1 ticketing projects ticketing project Id field maps params
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams) WithTicketingProjectID(ticketingProjectID string) *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams {
	o.SetTicketingProjectID(ticketingProjectID)
	return o
}

// SetTicketingProjectID adds the ticketingProjectId to the post v1 ticketing projects ticketing project Id field maps params
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams) SetTicketingProjectID(ticketingProjectID string) {
	o.TicketingProjectID = ticketingProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ticketing_project_id
	if err := r.SetPathParam("ticketing_project_id", o.TicketingProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
