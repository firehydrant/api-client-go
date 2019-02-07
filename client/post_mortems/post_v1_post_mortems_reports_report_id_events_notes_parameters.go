// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

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

// NewPostV1PostMortemsReportsReportIDEventsNotesParams creates a new PostV1PostMortemsReportsReportIDEventsNotesParams object
// with the default values initialized.
func NewPostV1PostMortemsReportsReportIDEventsNotesParams() *PostV1PostMortemsReportsReportIDEventsNotesParams {
	var ()
	return &PostV1PostMortemsReportsReportIDEventsNotesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1PostMortemsReportsReportIDEventsNotesParamsWithTimeout creates a new PostV1PostMortemsReportsReportIDEventsNotesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostV1PostMortemsReportsReportIDEventsNotesParamsWithTimeout(timeout time.Duration) *PostV1PostMortemsReportsReportIDEventsNotesParams {
	var ()
	return &PostV1PostMortemsReportsReportIDEventsNotesParams{

		timeout: timeout,
	}
}

// NewPostV1PostMortemsReportsReportIDEventsNotesParamsWithContext creates a new PostV1PostMortemsReportsReportIDEventsNotesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostV1PostMortemsReportsReportIDEventsNotesParamsWithContext(ctx context.Context) *PostV1PostMortemsReportsReportIDEventsNotesParams {
	var ()
	return &PostV1PostMortemsReportsReportIDEventsNotesParams{

		Context: ctx,
	}
}

// NewPostV1PostMortemsReportsReportIDEventsNotesParamsWithHTTPClient creates a new PostV1PostMortemsReportsReportIDEventsNotesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostV1PostMortemsReportsReportIDEventsNotesParamsWithHTTPClient(client *http.Client) *PostV1PostMortemsReportsReportIDEventsNotesParams {
	var ()
	return &PostV1PostMortemsReportsReportIDEventsNotesParams{
		HTTPClient: client,
	}
}

/*PostV1PostMortemsReportsReportIDEventsNotesParams contains all the parameters to send to the API endpoint
for the post v1 post mortems reports report Id events notes operation typically these are written to a http.Request
*/
type PostV1PostMortemsReportsReportIDEventsNotesParams struct {

	/*V1PostMortemsReportsReportIDEventsNotes*/
	V1PostMortemsReportsReportIDEventsNotes *models.PostV1PostMortemsReportsReportIDEventsNotes
	/*ReportID*/
	ReportID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post v1 post mortems reports report Id events notes params
func (o *PostV1PostMortemsReportsReportIDEventsNotesParams) WithTimeout(timeout time.Duration) *PostV1PostMortemsReportsReportIDEventsNotesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 post mortems reports report Id events notes params
func (o *PostV1PostMortemsReportsReportIDEventsNotesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 post mortems reports report Id events notes params
func (o *PostV1PostMortemsReportsReportIDEventsNotesParams) WithContext(ctx context.Context) *PostV1PostMortemsReportsReportIDEventsNotesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 post mortems reports report Id events notes params
func (o *PostV1PostMortemsReportsReportIDEventsNotesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 post mortems reports report Id events notes params
func (o *PostV1PostMortemsReportsReportIDEventsNotesParams) WithHTTPClient(client *http.Client) *PostV1PostMortemsReportsReportIDEventsNotesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 post mortems reports report Id events notes params
func (o *PostV1PostMortemsReportsReportIDEventsNotesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1PostMortemsReportsReportIDEventsNotes adds the v1PostMortemsReportsReportIDEventsNotes to the post v1 post mortems reports report Id events notes params
func (o *PostV1PostMortemsReportsReportIDEventsNotesParams) WithV1PostMortemsReportsReportIDEventsNotes(v1PostMortemsReportsReportIDEventsNotes *models.PostV1PostMortemsReportsReportIDEventsNotes) *PostV1PostMortemsReportsReportIDEventsNotesParams {
	o.SetV1PostMortemsReportsReportIDEventsNotes(v1PostMortemsReportsReportIDEventsNotes)
	return o
}

// SetV1PostMortemsReportsReportIDEventsNotes adds the v1PostMortemsReportsReportIdEventsNotes to the post v1 post mortems reports report Id events notes params
func (o *PostV1PostMortemsReportsReportIDEventsNotesParams) SetV1PostMortemsReportsReportIDEventsNotes(v1PostMortemsReportsReportIDEventsNotes *models.PostV1PostMortemsReportsReportIDEventsNotes) {
	o.V1PostMortemsReportsReportIDEventsNotes = v1PostMortemsReportsReportIDEventsNotes
}

// WithReportID adds the reportID to the post v1 post mortems reports report Id events notes params
func (o *PostV1PostMortemsReportsReportIDEventsNotesParams) WithReportID(reportID int32) *PostV1PostMortemsReportsReportIDEventsNotesParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the post v1 post mortems reports report Id events notes params
func (o *PostV1PostMortemsReportsReportIDEventsNotesParams) SetReportID(reportID int32) {
	o.ReportID = reportID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1PostMortemsReportsReportIDEventsNotesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.V1PostMortemsReportsReportIDEventsNotes != nil {
		if err := r.SetBodyParam(o.V1PostMortemsReportsReportIDEventsNotes); err != nil {
			return err
		}
	}

	// path param report_id
	if err := r.SetPathParam("report_id", swag.FormatInt32(o.ReportID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
