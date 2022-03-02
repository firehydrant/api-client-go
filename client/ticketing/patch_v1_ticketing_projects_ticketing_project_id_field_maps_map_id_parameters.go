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
	"github.com/go-openapi/swag"
)

// NewPatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams creates a new PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams() *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	return &PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParamsWithTimeout creates a new PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParamsWithTimeout(timeout time.Duration) *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	return &PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams{
		timeout: timeout,
	}
}

// NewPatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParamsWithContext creates a new PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams object
// with the ability to set a context for a request.
func NewPatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParamsWithContext(ctx context.Context) *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	return &PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams{
		Context: ctx,
	}
}

// NewPatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParamsWithHTTPClient creates a new PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParamsWithHTTPClient(client *http.Client) *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	return &PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams{
		HTTPClient: client,
	}
}

/* PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams contains all the parameters to send to the API endpoint
   for the patch v1 ticketing projects ticketing project Id field maps map Id operation.

   Typically these are written to a http.Request.
*/
type PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams struct {

	/* MapID.

	   ID of field map
	*/
	MapID string

	// TicketingProjectID.
	//
	// Format: int32
	TicketingProjectID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 ticketing projects ticketing project Id field maps map Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) WithDefaults() *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 ticketing projects ticketing project Id field maps map Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 ticketing projects ticketing project Id field maps map Id params
func (o *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) WithTimeout(timeout time.Duration) *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 ticketing projects ticketing project Id field maps map Id params
func (o *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 ticketing projects ticketing project Id field maps map Id params
func (o *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) WithContext(ctx context.Context) *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 ticketing projects ticketing project Id field maps map Id params
func (o *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 ticketing projects ticketing project Id field maps map Id params
func (o *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) WithHTTPClient(client *http.Client) *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 ticketing projects ticketing project Id field maps map Id params
func (o *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMapID adds the mapID to the patch v1 ticketing projects ticketing project Id field maps map Id params
func (o *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) WithMapID(mapID string) *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	o.SetMapID(mapID)
	return o
}

// SetMapID adds the mapId to the patch v1 ticketing projects ticketing project Id field maps map Id params
func (o *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) SetMapID(mapID string) {
	o.MapID = mapID
}

// WithTicketingProjectID adds the ticketingProjectID to the patch v1 ticketing projects ticketing project Id field maps map Id params
func (o *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) WithTicketingProjectID(ticketingProjectID int32) *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams {
	o.SetTicketingProjectID(ticketingProjectID)
	return o
}

// SetTicketingProjectID adds the ticketingProjectId to the patch v1 ticketing projects ticketing project Id field maps map Id params
func (o *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) SetTicketingProjectID(ticketingProjectID int32) {
	o.TicketingProjectID = ticketingProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1TicketingProjectsTicketingProjectIDFieldMapsMapIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param map_id
	if err := r.SetPathParam("map_id", o.MapID); err != nil {
		return err
	}

	// path param ticketing_project_id
	if err := r.SetPathParam("ticketing_project_id", swag.FormatInt32(o.TicketingProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
