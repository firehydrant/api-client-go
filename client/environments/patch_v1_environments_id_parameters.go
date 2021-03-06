// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firehydrant/api-client-go/models"
)

// NewPatchV1EnvironmentsIDParams creates a new PatchV1EnvironmentsIDParams object
// with the default values initialized.
func NewPatchV1EnvironmentsIDParams() *PatchV1EnvironmentsIDParams {
	var ()
	return &PatchV1EnvironmentsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1EnvironmentsIDParamsWithTimeout creates a new PatchV1EnvironmentsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchV1EnvironmentsIDParamsWithTimeout(timeout time.Duration) *PatchV1EnvironmentsIDParams {
	var ()
	return &PatchV1EnvironmentsIDParams{

		timeout: timeout,
	}
}

// NewPatchV1EnvironmentsIDParamsWithContext creates a new PatchV1EnvironmentsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchV1EnvironmentsIDParamsWithContext(ctx context.Context) *PatchV1EnvironmentsIDParams {
	var ()
	return &PatchV1EnvironmentsIDParams{

		Context: ctx,
	}
}

// NewPatchV1EnvironmentsIDParamsWithHTTPClient creates a new PatchV1EnvironmentsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchV1EnvironmentsIDParamsWithHTTPClient(client *http.Client) *PatchV1EnvironmentsIDParams {
	var ()
	return &PatchV1EnvironmentsIDParams{
		HTTPClient: client,
	}
}

/*PatchV1EnvironmentsIDParams contains all the parameters to send to the API endpoint
for the patch v1 environments Id operation typically these are written to a http.Request
*/
type PatchV1EnvironmentsIDParams struct {

	/*V1Environments*/
	V1Environments *models.PatchV1Environments
	/*ID
	  Environment UUID

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch v1 environments Id params
func (o *PatchV1EnvironmentsIDParams) WithTimeout(timeout time.Duration) *PatchV1EnvironmentsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 environments Id params
func (o *PatchV1EnvironmentsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 environments Id params
func (o *PatchV1EnvironmentsIDParams) WithContext(ctx context.Context) *PatchV1EnvironmentsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 environments Id params
func (o *PatchV1EnvironmentsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 environments Id params
func (o *PatchV1EnvironmentsIDParams) WithHTTPClient(client *http.Client) *PatchV1EnvironmentsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 environments Id params
func (o *PatchV1EnvironmentsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1Environments adds the v1Environments to the patch v1 environments Id params
func (o *PatchV1EnvironmentsIDParams) WithV1Environments(v1Environments *models.PatchV1Environments) *PatchV1EnvironmentsIDParams {
	o.SetV1Environments(v1Environments)
	return o
}

// SetV1Environments adds the v1Environments to the patch v1 environments Id params
func (o *PatchV1EnvironmentsIDParams) SetV1Environments(v1Environments *models.PatchV1Environments) {
	o.V1Environments = v1Environments
}

// WithID adds the id to the patch v1 environments Id params
func (o *PatchV1EnvironmentsIDParams) WithID(id string) *PatchV1EnvironmentsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch v1 environments Id params
func (o *PatchV1EnvironmentsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1EnvironmentsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.V1Environments != nil {
		if err := r.SetBodyParam(o.V1Environments); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
