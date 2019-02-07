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

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostV1IncidentsIncidentIDAttachmentsParams creates a new PostV1IncidentsIncidentIDAttachmentsParams object
// with the default values initialized.
func NewPostV1IncidentsIncidentIDAttachmentsParams() *PostV1IncidentsIncidentIDAttachmentsParams {
	var ()
	return &PostV1IncidentsIncidentIDAttachmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1IncidentsIncidentIDAttachmentsParamsWithTimeout creates a new PostV1IncidentsIncidentIDAttachmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostV1IncidentsIncidentIDAttachmentsParamsWithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDAttachmentsParams {
	var ()
	return &PostV1IncidentsIncidentIDAttachmentsParams{

		timeout: timeout,
	}
}

// NewPostV1IncidentsIncidentIDAttachmentsParamsWithContext creates a new PostV1IncidentsIncidentIDAttachmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostV1IncidentsIncidentIDAttachmentsParamsWithContext(ctx context.Context) *PostV1IncidentsIncidentIDAttachmentsParams {
	var ()
	return &PostV1IncidentsIncidentIDAttachmentsParams{

		Context: ctx,
	}
}

// NewPostV1IncidentsIncidentIDAttachmentsParamsWithHTTPClient creates a new PostV1IncidentsIncidentIDAttachmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostV1IncidentsIncidentIDAttachmentsParamsWithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDAttachmentsParams {
	var ()
	return &PostV1IncidentsIncidentIDAttachmentsParams{
		HTTPClient: client,
	}
}

/*PostV1IncidentsIncidentIDAttachmentsParams contains all the parameters to send to the API endpoint
for the post v1 incidents incident Id attachments operation typically these are written to a http.Request
*/
type PostV1IncidentsIncidentIDAttachmentsParams struct {

	/*Description*/
	Description *string
	/*File*/
	File runtime.NamedReadCloser
	/*IncidentID*/
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post v1 incidents incident Id attachments params
func (o *PostV1IncidentsIncidentIDAttachmentsParams) WithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDAttachmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 incidents incident Id attachments params
func (o *PostV1IncidentsIncidentIDAttachmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 incidents incident Id attachments params
func (o *PostV1IncidentsIncidentIDAttachmentsParams) WithContext(ctx context.Context) *PostV1IncidentsIncidentIDAttachmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 incidents incident Id attachments params
func (o *PostV1IncidentsIncidentIDAttachmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 incidents incident Id attachments params
func (o *PostV1IncidentsIncidentIDAttachmentsParams) WithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDAttachmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 incidents incident Id attachments params
func (o *PostV1IncidentsIncidentIDAttachmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the post v1 incidents incident Id attachments params
func (o *PostV1IncidentsIncidentIDAttachmentsParams) WithDescription(description *string) *PostV1IncidentsIncidentIDAttachmentsParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the post v1 incidents incident Id attachments params
func (o *PostV1IncidentsIncidentIDAttachmentsParams) SetDescription(description *string) {
	o.Description = description
}

// WithFile adds the file to the post v1 incidents incident Id attachments params
func (o *PostV1IncidentsIncidentIDAttachmentsParams) WithFile(file runtime.NamedReadCloser) *PostV1IncidentsIncidentIDAttachmentsParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the post v1 incidents incident Id attachments params
func (o *PostV1IncidentsIncidentIDAttachmentsParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithIncidentID adds the incidentID to the post v1 incidents incident Id attachments params
func (o *PostV1IncidentsIncidentIDAttachmentsParams) WithIncidentID(incidentID string) *PostV1IncidentsIncidentIDAttachmentsParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the post v1 incidents incident Id attachments params
func (o *PostV1IncidentsIncidentIDAttachmentsParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1IncidentsIncidentIDAttachmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// form param description
		var frDescription string
		if o.Description != nil {
			frDescription = *o.Description
		}
		fDescription := frDescription
		if fDescription != "" {
			if err := r.SetFormParam("description", fDescription); err != nil {
				return err
			}
		}

	}

	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
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
