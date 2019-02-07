// Code generated by go-swagger; DO NOT EDIT.

package incident_roles

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

	models "github.com/firehydrant/api-client-go/models"
)

// NewPatchV1IncidentRolesIncidentRoleIDParams creates a new PatchV1IncidentRolesIncidentRoleIDParams object
// with the default values initialized.
func NewPatchV1IncidentRolesIncidentRoleIDParams() *PatchV1IncidentRolesIncidentRoleIDParams {
	var ()
	return &PatchV1IncidentRolesIncidentRoleIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1IncidentRolesIncidentRoleIDParamsWithTimeout creates a new PatchV1IncidentRolesIncidentRoleIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchV1IncidentRolesIncidentRoleIDParamsWithTimeout(timeout time.Duration) *PatchV1IncidentRolesIncidentRoleIDParams {
	var ()
	return &PatchV1IncidentRolesIncidentRoleIDParams{

		timeout: timeout,
	}
}

// NewPatchV1IncidentRolesIncidentRoleIDParamsWithContext creates a new PatchV1IncidentRolesIncidentRoleIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchV1IncidentRolesIncidentRoleIDParamsWithContext(ctx context.Context) *PatchV1IncidentRolesIncidentRoleIDParams {
	var ()
	return &PatchV1IncidentRolesIncidentRoleIDParams{

		Context: ctx,
	}
}

// NewPatchV1IncidentRolesIncidentRoleIDParamsWithHTTPClient creates a new PatchV1IncidentRolesIncidentRoleIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchV1IncidentRolesIncidentRoleIDParamsWithHTTPClient(client *http.Client) *PatchV1IncidentRolesIncidentRoleIDParams {
	var ()
	return &PatchV1IncidentRolesIncidentRoleIDParams{
		HTTPClient: client,
	}
}

/*PatchV1IncidentRolesIncidentRoleIDParams contains all the parameters to send to the API endpoint
for the patch v1 incident roles incident role Id operation typically these are written to a http.Request
*/
type PatchV1IncidentRolesIncidentRoleIDParams struct {

	/*V1IncidentRoles*/
	V1IncidentRoles *models.PatchV1IncidentRoles
	/*IncidentRoleID*/
	IncidentRoleID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch v1 incident roles incident role Id params
func (o *PatchV1IncidentRolesIncidentRoleIDParams) WithTimeout(timeout time.Duration) *PatchV1IncidentRolesIncidentRoleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 incident roles incident role Id params
func (o *PatchV1IncidentRolesIncidentRoleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 incident roles incident role Id params
func (o *PatchV1IncidentRolesIncidentRoleIDParams) WithContext(ctx context.Context) *PatchV1IncidentRolesIncidentRoleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 incident roles incident role Id params
func (o *PatchV1IncidentRolesIncidentRoleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 incident roles incident role Id params
func (o *PatchV1IncidentRolesIncidentRoleIDParams) WithHTTPClient(client *http.Client) *PatchV1IncidentRolesIncidentRoleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 incident roles incident role Id params
func (o *PatchV1IncidentRolesIncidentRoleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1IncidentRoles adds the v1IncidentRoles to the patch v1 incident roles incident role Id params
func (o *PatchV1IncidentRolesIncidentRoleIDParams) WithV1IncidentRoles(v1IncidentRoles *models.PatchV1IncidentRoles) *PatchV1IncidentRolesIncidentRoleIDParams {
	o.SetV1IncidentRoles(v1IncidentRoles)
	return o
}

// SetV1IncidentRoles adds the v1IncidentRoles to the patch v1 incident roles incident role Id params
func (o *PatchV1IncidentRolesIncidentRoleIDParams) SetV1IncidentRoles(v1IncidentRoles *models.PatchV1IncidentRoles) {
	o.V1IncidentRoles = v1IncidentRoles
}

// WithIncidentRoleID adds the incidentRoleID to the patch v1 incident roles incident role Id params
func (o *PatchV1IncidentRolesIncidentRoleIDParams) WithIncidentRoleID(incidentRoleID int32) *PatchV1IncidentRolesIncidentRoleIDParams {
	o.SetIncidentRoleID(incidentRoleID)
	return o
}

// SetIncidentRoleID adds the incidentRoleId to the patch v1 incident roles incident role Id params
func (o *PatchV1IncidentRolesIncidentRoleIDParams) SetIncidentRoleID(incidentRoleID int32) {
	o.IncidentRoleID = incidentRoleID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1IncidentRolesIncidentRoleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.V1IncidentRoles != nil {
		if err := r.SetBodyParam(o.V1IncidentRoles); err != nil {
			return err
		}
	}

	// path param incident_role_id
	if err := r.SetPathParam("incident_role_id", swag.FormatInt32(o.IncidentRoleID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
