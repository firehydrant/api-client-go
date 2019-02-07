// Code generated by go-swagger; DO NOT EDIT.

package changes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteV1ChangesChangeIDIdentitiesIdentityIDParams creates a new DeleteV1ChangesChangeIDIdentitiesIdentityIDParams object
// with the default values initialized.
func NewDeleteV1ChangesChangeIDIdentitiesIdentityIDParams() *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams {
	var ()
	return &DeleteV1ChangesChangeIDIdentitiesIdentityIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1ChangesChangeIDIdentitiesIdentityIDParamsWithTimeout creates a new DeleteV1ChangesChangeIDIdentitiesIdentityIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteV1ChangesChangeIDIdentitiesIdentityIDParamsWithTimeout(timeout time.Duration) *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams {
	var ()
	return &DeleteV1ChangesChangeIDIdentitiesIdentityIDParams{

		timeout: timeout,
	}
}

// NewDeleteV1ChangesChangeIDIdentitiesIdentityIDParamsWithContext creates a new DeleteV1ChangesChangeIDIdentitiesIdentityIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteV1ChangesChangeIDIdentitiesIdentityIDParamsWithContext(ctx context.Context) *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams {
	var ()
	return &DeleteV1ChangesChangeIDIdentitiesIdentityIDParams{

		Context: ctx,
	}
}

// NewDeleteV1ChangesChangeIDIdentitiesIdentityIDParamsWithHTTPClient creates a new DeleteV1ChangesChangeIDIdentitiesIdentityIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteV1ChangesChangeIDIdentitiesIdentityIDParamsWithHTTPClient(client *http.Client) *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams {
	var ()
	return &DeleteV1ChangesChangeIDIdentitiesIdentityIDParams{
		HTTPClient: client,
	}
}

/*DeleteV1ChangesChangeIDIdentitiesIdentityIDParams contains all the parameters to send to the API endpoint
for the delete v1 changes change Id identities identity Id operation typically these are written to a http.Request
*/
type DeleteV1ChangesChangeIDIdentitiesIdentityIDParams struct {

	/*ChangeID*/
	ChangeID int32
	/*IdentityID*/
	IdentityID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete v1 changes change Id identities identity Id params
func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams) WithTimeout(timeout time.Duration) *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 changes change Id identities identity Id params
func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 changes change Id identities identity Id params
func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams) WithContext(ctx context.Context) *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 changes change Id identities identity Id params
func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 changes change Id identities identity Id params
func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams) WithHTTPClient(client *http.Client) *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 changes change Id identities identity Id params
func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangeID adds the changeID to the delete v1 changes change Id identities identity Id params
func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams) WithChangeID(changeID int32) *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams {
	o.SetChangeID(changeID)
	return o
}

// SetChangeID adds the changeId to the delete v1 changes change Id identities identity Id params
func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams) SetChangeID(changeID int32) {
	o.ChangeID = changeID
}

// WithIdentityID adds the identityID to the delete v1 changes change Id identities identity Id params
func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams) WithIdentityID(identityID int32) *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams {
	o.SetIdentityID(identityID)
	return o
}

// SetIdentityID adds the identityId to the delete v1 changes change Id identities identity Id params
func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams) SetIdentityID(identityID int32) {
	o.IdentityID = identityID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param change_id
	if err := r.SetPathParam("change_id", swag.FormatInt32(o.ChangeID)); err != nil {
		return err
	}

	// path param identity_id
	if err := r.SetPathParam("identity_id", swag.FormatInt32(o.IdentityID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
