// Code generated by go-swagger; DO NOT EDIT.

package status_pages

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

// NewPatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams creates a new PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams() *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams {
	return &PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParamsWithTimeout creates a new PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParamsWithTimeout(timeout time.Duration) *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams {
	return &PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams{
		timeout: timeout,
	}
}

// NewPatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParamsWithContext creates a new PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams object
// with the ability to set a context for a request.
func NewPatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParamsWithContext(ctx context.Context) *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams {
	return &PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams{
		Context: ctx,
	}
}

// NewPatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParamsWithHTTPClient creates a new PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParamsWithHTTPClient(client *http.Client) *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams {
	return &PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams contains all the parameters to send to the API endpoint

	for the patch v1 nunc connections nunc connection Id links link Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams struct {

	// LinkID.
	LinkID string

	// NuncConnectionID.
	NuncConnectionID string

	// PatchV1NuncConnectionsNuncConnectionIDLinksLinkID.
	PatchV1NuncConnectionsNuncConnectionIDLinksLinkID *models.PatchV1NuncConnectionsNuncConnectionIDLinksLinkID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 nunc connections nunc connection Id links link Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) WithDefaults() *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 nunc connections nunc connection Id links link Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 nunc connections nunc connection Id links link Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) WithTimeout(timeout time.Duration) *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 nunc connections nunc connection Id links link Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 nunc connections nunc connection Id links link Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) WithContext(ctx context.Context) *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 nunc connections nunc connection Id links link Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 nunc connections nunc connection Id links link Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) WithHTTPClient(client *http.Client) *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 nunc connections nunc connection Id links link Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLinkID adds the linkID to the patch v1 nunc connections nunc connection Id links link Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) WithLinkID(linkID string) *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams {
	o.SetLinkID(linkID)
	return o
}

// SetLinkID adds the linkId to the patch v1 nunc connections nunc connection Id links link Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) SetLinkID(linkID string) {
	o.LinkID = linkID
}

// WithNuncConnectionID adds the nuncConnectionID to the patch v1 nunc connections nunc connection Id links link Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) WithNuncConnectionID(nuncConnectionID string) *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams {
	o.SetNuncConnectionID(nuncConnectionID)
	return o
}

// SetNuncConnectionID adds the nuncConnectionId to the patch v1 nunc connections nunc connection Id links link Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) SetNuncConnectionID(nuncConnectionID string) {
	o.NuncConnectionID = nuncConnectionID
}

// WithPatchV1NuncConnectionsNuncConnectionIDLinksLinkID adds the patchV1NuncConnectionsNuncConnectionIDLinksLinkID to the patch v1 nunc connections nunc connection Id links link Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) WithPatchV1NuncConnectionsNuncConnectionIDLinksLinkID(patchV1NuncConnectionsNuncConnectionIDLinksLinkID *models.PatchV1NuncConnectionsNuncConnectionIDLinksLinkID) *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams {
	o.SetPatchV1NuncConnectionsNuncConnectionIDLinksLinkID(patchV1NuncConnectionsNuncConnectionIDLinksLinkID)
	return o
}

// SetPatchV1NuncConnectionsNuncConnectionIDLinksLinkID adds the patchV1NuncConnectionsNuncConnectionIdLinksLinkId to the patch v1 nunc connections nunc connection Id links link Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) SetPatchV1NuncConnectionsNuncConnectionIDLinksLinkID(patchV1NuncConnectionsNuncConnectionIDLinksLinkID *models.PatchV1NuncConnectionsNuncConnectionIDLinksLinkID) {
	o.PatchV1NuncConnectionsNuncConnectionIDLinksLinkID = patchV1NuncConnectionsNuncConnectionIDLinksLinkID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param link_id
	if err := r.SetPathParam("link_id", o.LinkID); err != nil {
		return err
	}

	// path param nunc_connection_id
	if err := r.SetPathParam("nunc_connection_id", o.NuncConnectionID); err != nil {
		return err
	}
	if o.PatchV1NuncConnectionsNuncConnectionIDLinksLinkID != nil {
		if err := r.SetBodyParam(o.PatchV1NuncConnectionsNuncConnectionIDLinksLinkID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
