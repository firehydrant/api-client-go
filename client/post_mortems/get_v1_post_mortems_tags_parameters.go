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
)

// NewGetV1PostMortemsTagsParams creates a new GetV1PostMortemsTagsParams object
// with the default values initialized.
func NewGetV1PostMortemsTagsParams() *GetV1PostMortemsTagsParams {
	var ()
	return &GetV1PostMortemsTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1PostMortemsTagsParamsWithTimeout creates a new GetV1PostMortemsTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1PostMortemsTagsParamsWithTimeout(timeout time.Duration) *GetV1PostMortemsTagsParams {
	var ()
	return &GetV1PostMortemsTagsParams{

		timeout: timeout,
	}
}

// NewGetV1PostMortemsTagsParamsWithContext creates a new GetV1PostMortemsTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1PostMortemsTagsParamsWithContext(ctx context.Context) *GetV1PostMortemsTagsParams {
	var ()
	return &GetV1PostMortemsTagsParams{

		Context: ctx,
	}
}

// NewGetV1PostMortemsTagsParamsWithHTTPClient creates a new GetV1PostMortemsTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1PostMortemsTagsParamsWithHTTPClient(client *http.Client) *GetV1PostMortemsTagsParams {
	var ()
	return &GetV1PostMortemsTagsParams{
		HTTPClient: client,
	}
}

/*GetV1PostMortemsTagsParams contains all the parameters to send to the API endpoint
for the get v1 post mortems tags operation typically these are written to a http.Request
*/
type GetV1PostMortemsTagsParams struct {

	/*Page*/
	Page *int32
	/*PerPage*/
	PerPage *int32
	/*Prefix*/
	Prefix string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 post mortems tags params
func (o *GetV1PostMortemsTagsParams) WithTimeout(timeout time.Duration) *GetV1PostMortemsTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 post mortems tags params
func (o *GetV1PostMortemsTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 post mortems tags params
func (o *GetV1PostMortemsTagsParams) WithContext(ctx context.Context) *GetV1PostMortemsTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 post mortems tags params
func (o *GetV1PostMortemsTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 post mortems tags params
func (o *GetV1PostMortemsTagsParams) WithHTTPClient(client *http.Client) *GetV1PostMortemsTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 post mortems tags params
func (o *GetV1PostMortemsTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 post mortems tags params
func (o *GetV1PostMortemsTagsParams) WithPage(page *int32) *GetV1PostMortemsTagsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 post mortems tags params
func (o *GetV1PostMortemsTagsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 post mortems tags params
func (o *GetV1PostMortemsTagsParams) WithPerPage(perPage *int32) *GetV1PostMortemsTagsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 post mortems tags params
func (o *GetV1PostMortemsTagsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithPrefix adds the prefix to the get v1 post mortems tags params
func (o *GetV1PostMortemsTagsParams) WithPrefix(prefix string) *GetV1PostMortemsTagsParams {
	o.SetPrefix(prefix)
	return o
}

// SetPrefix adds the prefix to the get v1 post mortems tags params
func (o *GetV1PostMortemsTagsParams) SetPrefix(prefix string) {
	o.Prefix = prefix
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1PostMortemsTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param prefix
	qrPrefix := o.Prefix
	qPrefix := qrPrefix
	if qPrefix != "" {
		if err := r.SetQueryParam("prefix", qPrefix); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
