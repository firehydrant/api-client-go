// Code generated by go-swagger; DO NOT EDIT.

package audiences

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

// NewGetV1AudiencesMemberMemberIDDefaultParams creates a new GetV1AudiencesMemberMemberIDDefaultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1AudiencesMemberMemberIDDefaultParams() *GetV1AudiencesMemberMemberIDDefaultParams {
	return &GetV1AudiencesMemberMemberIDDefaultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1AudiencesMemberMemberIDDefaultParamsWithTimeout creates a new GetV1AudiencesMemberMemberIDDefaultParams object
// with the ability to set a timeout on a request.
func NewGetV1AudiencesMemberMemberIDDefaultParamsWithTimeout(timeout time.Duration) *GetV1AudiencesMemberMemberIDDefaultParams {
	return &GetV1AudiencesMemberMemberIDDefaultParams{
		timeout: timeout,
	}
}

// NewGetV1AudiencesMemberMemberIDDefaultParamsWithContext creates a new GetV1AudiencesMemberMemberIDDefaultParams object
// with the ability to set a context for a request.
func NewGetV1AudiencesMemberMemberIDDefaultParamsWithContext(ctx context.Context) *GetV1AudiencesMemberMemberIDDefaultParams {
	return &GetV1AudiencesMemberMemberIDDefaultParams{
		Context: ctx,
	}
}

// NewGetV1AudiencesMemberMemberIDDefaultParamsWithHTTPClient creates a new GetV1AudiencesMemberMemberIDDefaultParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1AudiencesMemberMemberIDDefaultParamsWithHTTPClient(client *http.Client) *GetV1AudiencesMemberMemberIDDefaultParams {
	return &GetV1AudiencesMemberMemberIDDefaultParams{
		HTTPClient: client,
	}
}

/*
GetV1AudiencesMemberMemberIDDefaultParams contains all the parameters to send to the API endpoint

	for the get v1 audiences member member Id default operation.

	Typically these are written to a http.Request.
*/
type GetV1AudiencesMemberMemberIDDefaultParams struct {

	// MemberID.
	//
	// Format: int32
	MemberID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 audiences member member Id default params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AudiencesMemberMemberIDDefaultParams) WithDefaults() *GetV1AudiencesMemberMemberIDDefaultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 audiences member member Id default params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AudiencesMemberMemberIDDefaultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 audiences member member Id default params
func (o *GetV1AudiencesMemberMemberIDDefaultParams) WithTimeout(timeout time.Duration) *GetV1AudiencesMemberMemberIDDefaultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 audiences member member Id default params
func (o *GetV1AudiencesMemberMemberIDDefaultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 audiences member member Id default params
func (o *GetV1AudiencesMemberMemberIDDefaultParams) WithContext(ctx context.Context) *GetV1AudiencesMemberMemberIDDefaultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 audiences member member Id default params
func (o *GetV1AudiencesMemberMemberIDDefaultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 audiences member member Id default params
func (o *GetV1AudiencesMemberMemberIDDefaultParams) WithHTTPClient(client *http.Client) *GetV1AudiencesMemberMemberIDDefaultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 audiences member member Id default params
func (o *GetV1AudiencesMemberMemberIDDefaultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMemberID adds the memberID to the get v1 audiences member member Id default params
func (o *GetV1AudiencesMemberMemberIDDefaultParams) WithMemberID(memberID int32) *GetV1AudiencesMemberMemberIDDefaultParams {
	o.SetMemberID(memberID)
	return o
}

// SetMemberID adds the memberId to the get v1 audiences member member Id default params
func (o *GetV1AudiencesMemberMemberIDDefaultParams) SetMemberID(memberID int32) {
	o.MemberID = memberID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1AudiencesMemberMemberIDDefaultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param member_id
	if err := r.SetPathParam("member_id", swag.FormatInt32(o.MemberID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
