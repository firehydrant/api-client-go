// Code generated by go-swagger; DO NOT EDIT.

package retrospectives

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

// NewPatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams creates a new PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams() *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams {
	return &PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParamsWithTimeout creates a new PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParamsWithTimeout(timeout time.Duration) *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams {
	return &PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams{
		timeout: timeout,
	}
}

// NewPatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParamsWithContext creates a new PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams object
// with the ability to set a context for a request.
func NewPatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParamsWithContext(ctx context.Context) *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams {
	return &PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams{
		Context: ctx,
	}
}

// NewPatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParamsWithHTTPClient creates a new PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParamsWithHTTPClient(client *http.Client) *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams {
	return &PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams contains all the parameters to send to the API endpoint

	for the patch v1 incidents incident Id retrospectives retrospective Id fields field Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams struct {

	// FieldID.
	FieldID string

	// IncidentID.
	IncidentID string

	// PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID.
	PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID *models.PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID

	// RetrospectiveID.
	RetrospectiveID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) WithDefaults() *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) WithTimeout(timeout time.Duration) *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) WithContext(ctx context.Context) *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) WithHTTPClient(client *http.Client) *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldID adds the fieldID to the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) WithFieldID(fieldID string) *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams {
	o.SetFieldID(fieldID)
	return o
}

// SetFieldID adds the fieldId to the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) SetFieldID(fieldID string) {
	o.FieldID = fieldID
}

// WithIncidentID adds the incidentID to the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) WithIncidentID(incidentID string) *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithPatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID adds the patchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID to the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) WithPatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID(patchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID *models.PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID) *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams {
	o.SetPatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID(patchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID)
	return o
}

// SetPatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID adds the patchV1IncidentsIncidentIdRetrospectivesRetrospectiveIdFieldsFieldId to the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) SetPatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID(patchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID *models.PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID) {
	o.PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID = patchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID
}

// WithRetrospectiveID adds the retrospectiveID to the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) WithRetrospectiveID(retrospectiveID string) *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams {
	o.SetRetrospectiveID(retrospectiveID)
	return o
}

// SetRetrospectiveID adds the retrospectiveId to the patch v1 incidents incident Id retrospectives retrospective Id fields field Id params
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) SetRetrospectiveID(retrospectiveID string) {
	o.RetrospectiveID = retrospectiveID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param field_id
	if err := r.SetPathParam("field_id", o.FieldID); err != nil {
		return err
	}

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}
	if o.PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID != nil {
		if err := r.SetBodyParam(o.PatchV1IncidentsIncidentIDRetrospectivesRetrospectiveIDFieldsFieldID); err != nil {
			return err
		}
	}

	// path param retrospective_id
	if err := r.SetPathParam("retrospective_id", o.RetrospectiveID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
