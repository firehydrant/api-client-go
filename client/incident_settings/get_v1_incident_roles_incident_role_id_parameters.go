// Code generated by go-swagger; DO NOT EDIT.

package incident_settings

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

// NewGetV1IncidentRolesIncidentRoleIDParams creates a new GetV1IncidentRolesIncidentRoleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IncidentRolesIncidentRoleIDParams() *GetV1IncidentRolesIncidentRoleIDParams {
	return &GetV1IncidentRolesIncidentRoleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IncidentRolesIncidentRoleIDParamsWithTimeout creates a new GetV1IncidentRolesIncidentRoleIDParams object
// with the ability to set a timeout on a request.
func NewGetV1IncidentRolesIncidentRoleIDParamsWithTimeout(timeout time.Duration) *GetV1IncidentRolesIncidentRoleIDParams {
	return &GetV1IncidentRolesIncidentRoleIDParams{
		timeout: timeout,
	}
}

// NewGetV1IncidentRolesIncidentRoleIDParamsWithContext creates a new GetV1IncidentRolesIncidentRoleIDParams object
// with the ability to set a context for a request.
func NewGetV1IncidentRolesIncidentRoleIDParamsWithContext(ctx context.Context) *GetV1IncidentRolesIncidentRoleIDParams {
	return &GetV1IncidentRolesIncidentRoleIDParams{
		Context: ctx,
	}
}

// NewGetV1IncidentRolesIncidentRoleIDParamsWithHTTPClient creates a new GetV1IncidentRolesIncidentRoleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IncidentRolesIncidentRoleIDParamsWithHTTPClient(client *http.Client) *GetV1IncidentRolesIncidentRoleIDParams {
	return &GetV1IncidentRolesIncidentRoleIDParams{
		HTTPClient: client,
	}
}

/*
GetV1IncidentRolesIncidentRoleIDParams contains all the parameters to send to the API endpoint

	for the get v1 incident roles incident role Id operation.

	Typically these are written to a http.Request.
*/
type GetV1IncidentRolesIncidentRoleIDParams struct {

	// IncidentRoleID.
	IncidentRoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 incident roles incident role Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentRolesIncidentRoleIDParams) WithDefaults() *GetV1IncidentRolesIncidentRoleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 incident roles incident role Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentRolesIncidentRoleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 incident roles incident role Id params
func (o *GetV1IncidentRolesIncidentRoleIDParams) WithTimeout(timeout time.Duration) *GetV1IncidentRolesIncidentRoleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 incident roles incident role Id params
func (o *GetV1IncidentRolesIncidentRoleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 incident roles incident role Id params
func (o *GetV1IncidentRolesIncidentRoleIDParams) WithContext(ctx context.Context) *GetV1IncidentRolesIncidentRoleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 incident roles incident role Id params
func (o *GetV1IncidentRolesIncidentRoleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 incident roles incident role Id params
func (o *GetV1IncidentRolesIncidentRoleIDParams) WithHTTPClient(client *http.Client) *GetV1IncidentRolesIncidentRoleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 incident roles incident role Id params
func (o *GetV1IncidentRolesIncidentRoleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentRoleID adds the incidentRoleID to the get v1 incident roles incident role Id params
func (o *GetV1IncidentRolesIncidentRoleIDParams) WithIncidentRoleID(incidentRoleID string) *GetV1IncidentRolesIncidentRoleIDParams {
	o.SetIncidentRoleID(incidentRoleID)
	return o
}

// SetIncidentRoleID adds the incidentRoleId to the get v1 incident roles incident role Id params
func (o *GetV1IncidentRolesIncidentRoleIDParams) SetIncidentRoleID(incidentRoleID string) {
	o.IncidentRoleID = incidentRoleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IncidentRolesIncidentRoleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_role_id
	if err := r.SetPathParam("incident_role_id", o.IncidentRoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
