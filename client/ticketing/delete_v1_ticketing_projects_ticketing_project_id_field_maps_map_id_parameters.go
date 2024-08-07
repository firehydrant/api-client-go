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

// NewDeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams creates a new DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams() *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	return &DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParamsWithTimeout creates a new DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParamsWithTimeout(timeout time.Duration) *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	return &DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParamsWithContext creates a new DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams object
// with the ability to set a context for a request.
func NewDeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParamsWithContext(ctx context.Context) *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	return &DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams{
		Context: ctx,
	}
}

// NewDeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParamsWithHTTPClient creates a new DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParamsWithHTTPClient(client *http.Client) *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	return &DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams contains all the parameters to send to the API endpoint

	for the delete v1 ticketing projects ticketing project Id field maps map Id operation.

	Typically these are written to a http.Request.
*/
type DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams struct {

	// MapID.
	MapID string

	// TicketingProjectID.
	TicketingProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 ticketing projects ticketing project Id field maps map Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) WithDefaults() *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 ticketing projects ticketing project Id field maps map Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 ticketing projects ticketing project Id field maps map Id params
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) WithTimeout(timeout time.Duration) *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 ticketing projects ticketing project Id field maps map Id params
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 ticketing projects ticketing project Id field maps map Id params
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) WithContext(ctx context.Context) *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 ticketing projects ticketing project Id field maps map Id params
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 ticketing projects ticketing project Id field maps map Id params
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) WithHTTPClient(client *http.Client) *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 ticketing projects ticketing project Id field maps map Id params
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMapID adds the mapID to the delete v1 ticketing projects ticketing project Id field maps map Id params
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) WithMapID(mapID string) *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	o.SetMapID(mapID)
	return o
}

// SetMapID adds the mapId to the delete v1 ticketing projects ticketing project Id field maps map Id params
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) SetMapID(mapID string) {
	o.MapID = mapID
}

// WithTicketingProjectID adds the ticketingProjectID to the delete v1 ticketing projects ticketing project Id field maps map Id params
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) WithTicketingProjectID(ticketingProjectID string) *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	o.SetTicketingProjectID(ticketingProjectID)
	return o
}

// SetTicketingProjectID adds the ticketingProjectId to the delete v1 ticketing projects ticketing project Id field maps map Id params
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) SetTicketingProjectID(ticketingProjectID string) {
	o.TicketingProjectID = ticketingProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param map_id
	if err := r.SetPathParam("map_id", o.MapID); err != nil {
		return err
	}

	// path param ticketing_project_id
	if err := r.SetPathParam("ticketing_project_id", o.TicketingProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
