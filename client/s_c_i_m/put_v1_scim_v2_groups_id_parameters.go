// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

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

	"github.com/firehydrant/api-client-go/models"
)

// NewPutV1ScimV2GroupsIDParams creates a new PutV1ScimV2GroupsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutV1ScimV2GroupsIDParams() *PutV1ScimV2GroupsIDParams {
	return &PutV1ScimV2GroupsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1ScimV2GroupsIDParamsWithTimeout creates a new PutV1ScimV2GroupsIDParams object
// with the ability to set a timeout on a request.
func NewPutV1ScimV2GroupsIDParamsWithTimeout(timeout time.Duration) *PutV1ScimV2GroupsIDParams {
	return &PutV1ScimV2GroupsIDParams{
		timeout: timeout,
	}
}

// NewPutV1ScimV2GroupsIDParamsWithContext creates a new PutV1ScimV2GroupsIDParams object
// with the ability to set a context for a request.
func NewPutV1ScimV2GroupsIDParamsWithContext(ctx context.Context) *PutV1ScimV2GroupsIDParams {
	return &PutV1ScimV2GroupsIDParams{
		Context: ctx,
	}
}

// NewPutV1ScimV2GroupsIDParamsWithHTTPClient creates a new PutV1ScimV2GroupsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutV1ScimV2GroupsIDParamsWithHTTPClient(client *http.Client) *PutV1ScimV2GroupsIDParams {
	return &PutV1ScimV2GroupsIDParams{
		HTTPClient: client,
	}
}

/*
PutV1ScimV2GroupsIDParams contains all the parameters to send to the API endpoint

	for the put v1 scim v2 groups Id operation.

	Typically these are written to a http.Request.
*/
type PutV1ScimV2GroupsIDParams struct {

	// ID.
	ID string

	// PutV1ScimV2GroupsID.
	PutV1ScimV2GroupsID *models.PutV1ScimV2GroupsID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put v1 scim v2 groups Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1ScimV2GroupsIDParams) WithDefaults() *PutV1ScimV2GroupsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put v1 scim v2 groups Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1ScimV2GroupsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put v1 scim v2 groups Id params
func (o *PutV1ScimV2GroupsIDParams) WithTimeout(timeout time.Duration) *PutV1ScimV2GroupsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 scim v2 groups Id params
func (o *PutV1ScimV2GroupsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 scim v2 groups Id params
func (o *PutV1ScimV2GroupsIDParams) WithContext(ctx context.Context) *PutV1ScimV2GroupsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 scim v2 groups Id params
func (o *PutV1ScimV2GroupsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put v1 scim v2 groups Id params
func (o *PutV1ScimV2GroupsIDParams) WithHTTPClient(client *http.Client) *PutV1ScimV2GroupsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put v1 scim v2 groups Id params
func (o *PutV1ScimV2GroupsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put v1 scim v2 groups Id params
func (o *PutV1ScimV2GroupsIDParams) WithID(id string) *PutV1ScimV2GroupsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put v1 scim v2 groups Id params
func (o *PutV1ScimV2GroupsIDParams) SetID(id string) {
	o.ID = id
}

// WithPutV1ScimV2GroupsID adds the putV1ScimV2GroupsID to the put v1 scim v2 groups Id params
func (o *PutV1ScimV2GroupsIDParams) WithPutV1ScimV2GroupsID(putV1ScimV2GroupsID *models.PutV1ScimV2GroupsID) *PutV1ScimV2GroupsIDParams {
	o.SetPutV1ScimV2GroupsID(putV1ScimV2GroupsID)
	return o
}

// SetPutV1ScimV2GroupsID adds the putV1ScimV2GroupsId to the put v1 scim v2 groups Id params
func (o *PutV1ScimV2GroupsIDParams) SetPutV1ScimV2GroupsID(putV1ScimV2GroupsID *models.PutV1ScimV2GroupsID) {
	o.PutV1ScimV2GroupsID = putV1ScimV2GroupsID
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1ScimV2GroupsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.PutV1ScimV2GroupsID != nil {
		if err := r.SetBodyParam(o.PutV1ScimV2GroupsID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
