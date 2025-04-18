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

// NewPutV1ScimV2UsersIDParams creates a new PutV1ScimV2UsersIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutV1ScimV2UsersIDParams() *PutV1ScimV2UsersIDParams {
	return &PutV1ScimV2UsersIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1ScimV2UsersIDParamsWithTimeout creates a new PutV1ScimV2UsersIDParams object
// with the ability to set a timeout on a request.
func NewPutV1ScimV2UsersIDParamsWithTimeout(timeout time.Duration) *PutV1ScimV2UsersIDParams {
	return &PutV1ScimV2UsersIDParams{
		timeout: timeout,
	}
}

// NewPutV1ScimV2UsersIDParamsWithContext creates a new PutV1ScimV2UsersIDParams object
// with the ability to set a context for a request.
func NewPutV1ScimV2UsersIDParamsWithContext(ctx context.Context) *PutV1ScimV2UsersIDParams {
	return &PutV1ScimV2UsersIDParams{
		Context: ctx,
	}
}

// NewPutV1ScimV2UsersIDParamsWithHTTPClient creates a new PutV1ScimV2UsersIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutV1ScimV2UsersIDParamsWithHTTPClient(client *http.Client) *PutV1ScimV2UsersIDParams {
	return &PutV1ScimV2UsersIDParams{
		HTTPClient: client,
	}
}

/*
PutV1ScimV2UsersIDParams contains all the parameters to send to the API endpoint

	for the put v1 scim v2 users Id operation.

	Typically these are written to a http.Request.
*/
type PutV1ScimV2UsersIDParams struct {

	// ID.
	ID string

	// PutV1ScimV2UsersID.
	PutV1ScimV2UsersID *models.PutV1ScimV2UsersID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put v1 scim v2 users Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1ScimV2UsersIDParams) WithDefaults() *PutV1ScimV2UsersIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put v1 scim v2 users Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1ScimV2UsersIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put v1 scim v2 users Id params
func (o *PutV1ScimV2UsersIDParams) WithTimeout(timeout time.Duration) *PutV1ScimV2UsersIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 scim v2 users Id params
func (o *PutV1ScimV2UsersIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 scim v2 users Id params
func (o *PutV1ScimV2UsersIDParams) WithContext(ctx context.Context) *PutV1ScimV2UsersIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 scim v2 users Id params
func (o *PutV1ScimV2UsersIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put v1 scim v2 users Id params
func (o *PutV1ScimV2UsersIDParams) WithHTTPClient(client *http.Client) *PutV1ScimV2UsersIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put v1 scim v2 users Id params
func (o *PutV1ScimV2UsersIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put v1 scim v2 users Id params
func (o *PutV1ScimV2UsersIDParams) WithID(id string) *PutV1ScimV2UsersIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put v1 scim v2 users Id params
func (o *PutV1ScimV2UsersIDParams) SetID(id string) {
	o.ID = id
}

// WithPutV1ScimV2UsersID adds the putV1ScimV2UsersID to the put v1 scim v2 users Id params
func (o *PutV1ScimV2UsersIDParams) WithPutV1ScimV2UsersID(putV1ScimV2UsersID *models.PutV1ScimV2UsersID) *PutV1ScimV2UsersIDParams {
	o.SetPutV1ScimV2UsersID(putV1ScimV2UsersID)
	return o
}

// SetPutV1ScimV2UsersID adds the putV1ScimV2UsersId to the put v1 scim v2 users Id params
func (o *PutV1ScimV2UsersIDParams) SetPutV1ScimV2UsersID(putV1ScimV2UsersID *models.PutV1ScimV2UsersID) {
	o.PutV1ScimV2UsersID = putV1ScimV2UsersID
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1ScimV2UsersIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.PutV1ScimV2UsersID != nil {
		if err := r.SetBodyParam(o.PutV1ScimV2UsersID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
