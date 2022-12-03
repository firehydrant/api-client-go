// Code generated by go-swagger; DO NOT EDIT.

package scim

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

// NewGetV1ScimV2UsersParams creates a new GetV1ScimV2UsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ScimV2UsersParams() *GetV1ScimV2UsersParams {
	return &GetV1ScimV2UsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ScimV2UsersParamsWithTimeout creates a new GetV1ScimV2UsersParams object
// with the ability to set a timeout on a request.
func NewGetV1ScimV2UsersParamsWithTimeout(timeout time.Duration) *GetV1ScimV2UsersParams {
	return &GetV1ScimV2UsersParams{
		timeout: timeout,
	}
}

// NewGetV1ScimV2UsersParamsWithContext creates a new GetV1ScimV2UsersParams object
// with the ability to set a context for a request.
func NewGetV1ScimV2UsersParamsWithContext(ctx context.Context) *GetV1ScimV2UsersParams {
	return &GetV1ScimV2UsersParams{
		Context: ctx,
	}
}

// NewGetV1ScimV2UsersParamsWithHTTPClient creates a new GetV1ScimV2UsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ScimV2UsersParamsWithHTTPClient(client *http.Client) *GetV1ScimV2UsersParams {
	return &GetV1ScimV2UsersParams{
		HTTPClient: client,
	}
}

/*
GetV1ScimV2UsersParams contains all the parameters to send to the API endpoint

	for the get v1 scim v2 users operation.

	Typically these are written to a http.Request.
*/
type GetV1ScimV2UsersParams struct {

	/* Count.

	   This is an integer which represents the number of items per page in the response

	   Format: int32
	*/
	Count *int32

	/* Filter.

	   This is a string used to query users by either userName or email.
	      Proper example syntax for this would be `?filter=userName eq john` or `?filter=userName eq "john@firehydrant.com"`.
	      Currently we only support the `eq` operator
	*/
	Filter *string

	/* StartIndex.

	   This is an integer which represents a pagination offset

	   Format: int32
	*/
	StartIndex *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 scim v2 users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ScimV2UsersParams) WithDefaults() *GetV1ScimV2UsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 scim v2 users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ScimV2UsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 scim v2 users params
func (o *GetV1ScimV2UsersParams) WithTimeout(timeout time.Duration) *GetV1ScimV2UsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 scim v2 users params
func (o *GetV1ScimV2UsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 scim v2 users params
func (o *GetV1ScimV2UsersParams) WithContext(ctx context.Context) *GetV1ScimV2UsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 scim v2 users params
func (o *GetV1ScimV2UsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 scim v2 users params
func (o *GetV1ScimV2UsersParams) WithHTTPClient(client *http.Client) *GetV1ScimV2UsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 scim v2 users params
func (o *GetV1ScimV2UsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get v1 scim v2 users params
func (o *GetV1ScimV2UsersParams) WithCount(count *int32) *GetV1ScimV2UsersParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get v1 scim v2 users params
func (o *GetV1ScimV2UsersParams) SetCount(count *int32) {
	o.Count = count
}

// WithFilter adds the filter to the get v1 scim v2 users params
func (o *GetV1ScimV2UsersParams) WithFilter(filter *string) *GetV1ScimV2UsersParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get v1 scim v2 users params
func (o *GetV1ScimV2UsersParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithStartIndex adds the startIndex to the get v1 scim v2 users params
func (o *GetV1ScimV2UsersParams) WithStartIndex(startIndex *int32) *GetV1ScimV2UsersParams {
	o.SetStartIndex(startIndex)
	return o
}

// SetStartIndex adds the startIndex to the get v1 scim v2 users params
func (o *GetV1ScimV2UsersParams) SetStartIndex(startIndex *int32) {
	o.StartIndex = startIndex
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ScimV2UsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int32

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}
	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.StartIndex != nil {

		// query param startIndex
		var qrStartIndex int32

		if o.StartIndex != nil {
			qrStartIndex = *o.StartIndex
		}
		qStartIndex := swag.FormatInt32(qrStartIndex)
		if qStartIndex != "" {

			if err := r.SetQueryParam("startIndex", qStartIndex); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
