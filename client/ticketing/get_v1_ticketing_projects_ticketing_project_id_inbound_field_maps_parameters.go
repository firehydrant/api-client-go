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

// NewGetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams creates a new GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams() *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams {
	return &GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParamsWithTimeout creates a new GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams object
// with the ability to set a timeout on a request.
func NewGetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParamsWithTimeout(timeout time.Duration) *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams {
	return &GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams{
		timeout: timeout,
	}
}

// NewGetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParamsWithContext creates a new GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams object
// with the ability to set a context for a request.
func NewGetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParamsWithContext(ctx context.Context) *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams {
	return &GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams{
		Context: ctx,
	}
}

// NewGetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParamsWithHTTPClient creates a new GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParamsWithHTTPClient(client *http.Client) *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams {
	return &GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams{
		HTTPClient: client,
	}
}

/*
GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams contains all the parameters to send to the API endpoint

	for the get v1 ticketing projects ticketing project Id inbound field maps operation.

	Typically these are written to a http.Request.
*/
type GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams struct {

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	/* TicketType.

	   Filter by ticket type. Values: incident, follow_up
	*/
	TicketType *string

	// TicketingProjectID.
	TicketingProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 ticketing projects ticketing project Id inbound field maps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) WithDefaults() *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 ticketing projects ticketing project Id inbound field maps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 ticketing projects ticketing project Id inbound field maps params
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) WithTimeout(timeout time.Duration) *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 ticketing projects ticketing project Id inbound field maps params
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 ticketing projects ticketing project Id inbound field maps params
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) WithContext(ctx context.Context) *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 ticketing projects ticketing project Id inbound field maps params
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 ticketing projects ticketing project Id inbound field maps params
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) WithHTTPClient(client *http.Client) *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 ticketing projects ticketing project Id inbound field maps params
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 ticketing projects ticketing project Id inbound field maps params
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) WithPage(page *int32) *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 ticketing projects ticketing project Id inbound field maps params
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 ticketing projects ticketing project Id inbound field maps params
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) WithPerPage(perPage *int32) *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 ticketing projects ticketing project Id inbound field maps params
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithTicketType adds the ticketType to the get v1 ticketing projects ticketing project Id inbound field maps params
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) WithTicketType(ticketType *string) *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams {
	o.SetTicketType(ticketType)
	return o
}

// SetTicketType adds the ticketType to the get v1 ticketing projects ticketing project Id inbound field maps params
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) SetTicketType(ticketType *string) {
	o.TicketType = ticketType
}

// WithTicketingProjectID adds the ticketingProjectID to the get v1 ticketing projects ticketing project Id inbound field maps params
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) WithTicketingProjectID(ticketingProjectID string) *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams {
	o.SetTicketingProjectID(ticketingProjectID)
	return o
}

// SetTicketingProjectID adds the ticketingProjectId to the get v1 ticketing projects ticketing project Id inbound field maps params
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) SetTicketingProjectID(ticketingProjectID string) {
	o.TicketingProjectID = ticketingProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if o.TicketType != nil {

		// query param ticket_type
		var qrTicketType string

		if o.TicketType != nil {
			qrTicketType = *o.TicketType
		}
		qTicketType := qrTicketType
		if qTicketType != "" {

			if err := r.SetQueryParam("ticket_type", qTicketType); err != nil {
				return err
			}
		}
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
