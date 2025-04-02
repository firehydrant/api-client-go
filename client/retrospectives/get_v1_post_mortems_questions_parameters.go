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
	"github.com/go-openapi/swag"
)

// NewGetV1PostMortemsQuestionsParams creates a new GetV1PostMortemsQuestionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1PostMortemsQuestionsParams() *GetV1PostMortemsQuestionsParams {
	return &GetV1PostMortemsQuestionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1PostMortemsQuestionsParamsWithTimeout creates a new GetV1PostMortemsQuestionsParams object
// with the ability to set a timeout on a request.
func NewGetV1PostMortemsQuestionsParamsWithTimeout(timeout time.Duration) *GetV1PostMortemsQuestionsParams {
	return &GetV1PostMortemsQuestionsParams{
		timeout: timeout,
	}
}

// NewGetV1PostMortemsQuestionsParamsWithContext creates a new GetV1PostMortemsQuestionsParams object
// with the ability to set a context for a request.
func NewGetV1PostMortemsQuestionsParamsWithContext(ctx context.Context) *GetV1PostMortemsQuestionsParams {
	return &GetV1PostMortemsQuestionsParams{
		Context: ctx,
	}
}

// NewGetV1PostMortemsQuestionsParamsWithHTTPClient creates a new GetV1PostMortemsQuestionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1PostMortemsQuestionsParamsWithHTTPClient(client *http.Client) *GetV1PostMortemsQuestionsParams {
	return &GetV1PostMortemsQuestionsParams{
		HTTPClient: client,
	}
}

/*
GetV1PostMortemsQuestionsParams contains all the parameters to send to the API endpoint

	for the get v1 post mortems questions operation.

	Typically these are written to a http.Request.
*/
type GetV1PostMortemsQuestionsParams struct {

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 post mortems questions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1PostMortemsQuestionsParams) WithDefaults() *GetV1PostMortemsQuestionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 post mortems questions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1PostMortemsQuestionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 post mortems questions params
func (o *GetV1PostMortemsQuestionsParams) WithTimeout(timeout time.Duration) *GetV1PostMortemsQuestionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 post mortems questions params
func (o *GetV1PostMortemsQuestionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 post mortems questions params
func (o *GetV1PostMortemsQuestionsParams) WithContext(ctx context.Context) *GetV1PostMortemsQuestionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 post mortems questions params
func (o *GetV1PostMortemsQuestionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 post mortems questions params
func (o *GetV1PostMortemsQuestionsParams) WithHTTPClient(client *http.Client) *GetV1PostMortemsQuestionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 post mortems questions params
func (o *GetV1PostMortemsQuestionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 post mortems questions params
func (o *GetV1PostMortemsQuestionsParams) WithPage(page *int32) *GetV1PostMortemsQuestionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 post mortems questions params
func (o *GetV1PostMortemsQuestionsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 post mortems questions params
func (o *GetV1PostMortemsQuestionsParams) WithPerPage(perPage *int32) *GetV1PostMortemsQuestionsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 post mortems questions params
func (o *GetV1PostMortemsQuestionsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1PostMortemsQuestionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
