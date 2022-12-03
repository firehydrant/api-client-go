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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/firehydrant/api-client-go/models"
)

// NewPatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams creates a new PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams() *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams {
	return &PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1PostMortemsReportsReportIDEventsNotesNoteIDParamsWithTimeout creates a new PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1PostMortemsReportsReportIDEventsNotesNoteIDParamsWithTimeout(timeout time.Duration) *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams {
	return &PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams{
		timeout: timeout,
	}
}

// NewPatchV1PostMortemsReportsReportIDEventsNotesNoteIDParamsWithContext creates a new PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams object
// with the ability to set a context for a request.
func NewPatchV1PostMortemsReportsReportIDEventsNotesNoteIDParamsWithContext(ctx context.Context) *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams {
	return &PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams{
		Context: ctx,
	}
}

// NewPatchV1PostMortemsReportsReportIDEventsNotesNoteIDParamsWithHTTPClient creates a new PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1PostMortemsReportsReportIDEventsNotesNoteIDParamsWithHTTPClient(client *http.Client) *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams {
	return &PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams contains all the parameters to send to the API endpoint

	for the patch v1 post mortems reports report Id events notes note Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams struct {

	// V1PostMortemsReportsReportIDEventsNotes.
	V1PostMortemsReportsReportIDEventsNotes *models.PatchV1PostMortemsReportsReportIDEventsNotes

	// NoteID.
	NoteID string

	// ReportID.
	//
	// Format: int32
	ReportID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 post mortems reports report Id events notes note Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) WithDefaults() *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 post mortems reports report Id events notes note Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 post mortems reports report Id events notes note Id params
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) WithTimeout(timeout time.Duration) *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 post mortems reports report Id events notes note Id params
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 post mortems reports report Id events notes note Id params
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) WithContext(ctx context.Context) *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 post mortems reports report Id events notes note Id params
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 post mortems reports report Id events notes note Id params
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) WithHTTPClient(client *http.Client) *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 post mortems reports report Id events notes note Id params
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1PostMortemsReportsReportIDEventsNotes adds the v1PostMortemsReportsReportIDEventsNotes to the patch v1 post mortems reports report Id events notes note Id params
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) WithV1PostMortemsReportsReportIDEventsNotes(v1PostMortemsReportsReportIDEventsNotes *models.PatchV1PostMortemsReportsReportIDEventsNotes) *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams {
	o.SetV1PostMortemsReportsReportIDEventsNotes(v1PostMortemsReportsReportIDEventsNotes)
	return o
}

// SetV1PostMortemsReportsReportIDEventsNotes adds the v1PostMortemsReportsReportIdEventsNotes to the patch v1 post mortems reports report Id events notes note Id params
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) SetV1PostMortemsReportsReportIDEventsNotes(v1PostMortemsReportsReportIDEventsNotes *models.PatchV1PostMortemsReportsReportIDEventsNotes) {
	o.V1PostMortemsReportsReportIDEventsNotes = v1PostMortemsReportsReportIDEventsNotes
}

// WithNoteID adds the noteID to the patch v1 post mortems reports report Id events notes note Id params
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) WithNoteID(noteID string) *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams {
	o.SetNoteID(noteID)
	return o
}

// SetNoteID adds the noteId to the patch v1 post mortems reports report Id events notes note Id params
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) SetNoteID(noteID string) {
	o.NoteID = noteID
}

// WithReportID adds the reportID to the patch v1 post mortems reports report Id events notes note Id params
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) WithReportID(reportID int32) *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the patch v1 post mortems reports report Id events notes note Id params
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) SetReportID(reportID int32) {
	o.ReportID = reportID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1PostMortemsReportsReportIDEventsNotesNoteIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1PostMortemsReportsReportIDEventsNotes != nil {
		if err := r.SetBodyParam(o.V1PostMortemsReportsReportIDEventsNotes); err != nil {
			return err
		}
	}

	// path param note_id
	if err := r.SetPathParam("note_id", o.NoteID); err != nil {
		return err
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
