// Code generated by go-swagger; DO NOT EDIT.

package incidents

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

// NewDeleteV1IncidentsIncidentIDEventsEventIDStarsParams creates a new DeleteV1IncidentsIncidentIDEventsEventIDStarsParams object
// with the default values initialized.
func NewDeleteV1IncidentsIncidentIDEventsEventIDStarsParams() *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams {
	var ()
	return &DeleteV1IncidentsIncidentIDEventsEventIDStarsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1IncidentsIncidentIDEventsEventIDStarsParamsWithTimeout creates a new DeleteV1IncidentsIncidentIDEventsEventIDStarsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteV1IncidentsIncidentIDEventsEventIDStarsParamsWithTimeout(timeout time.Duration) *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams {
	var ()
	return &DeleteV1IncidentsIncidentIDEventsEventIDStarsParams{

		timeout: timeout,
	}
}

// NewDeleteV1IncidentsIncidentIDEventsEventIDStarsParamsWithContext creates a new DeleteV1IncidentsIncidentIDEventsEventIDStarsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteV1IncidentsIncidentIDEventsEventIDStarsParamsWithContext(ctx context.Context) *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams {
	var ()
	return &DeleteV1IncidentsIncidentIDEventsEventIDStarsParams{

		Context: ctx,
	}
}

// NewDeleteV1IncidentsIncidentIDEventsEventIDStarsParamsWithHTTPClient creates a new DeleteV1IncidentsIncidentIDEventsEventIDStarsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteV1IncidentsIncidentIDEventsEventIDStarsParamsWithHTTPClient(client *http.Client) *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams {
	var ()
	return &DeleteV1IncidentsIncidentIDEventsEventIDStarsParams{
		HTTPClient: client,
	}
}

/*DeleteV1IncidentsIncidentIDEventsEventIDStarsParams contains all the parameters to send to the API endpoint
for the delete v1 incidents incident Id events event Id stars operation typically these are written to a http.Request
*/
type DeleteV1IncidentsIncidentIDEventsEventIDStarsParams struct {

	/*EventID*/
	EventID int32
	/*IncidentID*/
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete v1 incidents incident Id events event Id stars params
func (o *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams) WithTimeout(timeout time.Duration) *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 incidents incident Id events event Id stars params
func (o *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 incidents incident Id events event Id stars params
func (o *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams) WithContext(ctx context.Context) *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 incidents incident Id events event Id stars params
func (o *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 incidents incident Id events event Id stars params
func (o *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams) WithHTTPClient(client *http.Client) *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 incidents incident Id events event Id stars params
func (o *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEventID adds the eventID to the delete v1 incidents incident Id events event Id stars params
func (o *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams) WithEventID(eventID int32) *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the delete v1 incidents incident Id events event Id stars params
func (o *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams) SetEventID(eventID int32) {
	o.EventID = eventID
}

// WithIncidentID adds the incidentID to the delete v1 incidents incident Id events event Id stars params
func (o *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams) WithIncidentID(incidentID string) *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the delete v1 incidents incident Id events event Id stars params
func (o *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1IncidentsIncidentIDEventsEventIDStarsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param event_id
	if err := r.SetPathParam("event_id", swag.FormatInt32(o.EventID)); err != nil {
		return err
	}

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
