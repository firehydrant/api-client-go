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
	"github.com/go-openapi/swag"
)

// NewPatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams creates a new PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams() *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams {
	return &PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParamsWithTimeout creates a new PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParamsWithTimeout(timeout time.Duration) *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams {
	return &PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams{
		timeout: timeout,
	}
}

// NewPatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParamsWithContext creates a new PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams object
// with the ability to set a context for a request.
func NewPatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParamsWithContext(ctx context.Context) *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams {
	return &PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams{
		Context: ctx,
	}
}

// NewPatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParamsWithHTTPClient creates a new PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParamsWithHTTPClient(client *http.Client) *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams {
	return &PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams contains all the parameters to send to the API endpoint

	for the patch v1 nunc connections nunc connection Id component groups group Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams struct {

	// ComponentGroupID.
	ComponentGroupID *string

	// GroupID.
	GroupID string

	// Name.
	Name *string

	// NuncConnectionID.
	NuncConnectionID string

	// Position.
	//
	// Format: int32
	Position *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 nunc connections nunc connection Id component groups group Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) WithDefaults() *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 nunc connections nunc connection Id component groups group Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) WithTimeout(timeout time.Duration) *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) WithContext(ctx context.Context) *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) WithHTTPClient(client *http.Client) *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentGroupID adds the componentGroupID to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) WithComponentGroupID(componentGroupID *string) *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams {
	o.SetComponentGroupID(componentGroupID)
	return o
}

// SetComponentGroupID adds the componentGroupId to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) SetComponentGroupID(componentGroupID *string) {
	o.ComponentGroupID = componentGroupID
}

// WithGroupID adds the groupID to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) WithGroupID(groupID string) *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithName adds the name to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) WithName(name *string) *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) SetName(name *string) {
	o.Name = name
}

// WithNuncConnectionID adds the nuncConnectionID to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) WithNuncConnectionID(nuncConnectionID string) *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams {
	o.SetNuncConnectionID(nuncConnectionID)
	return o
}

// SetNuncConnectionID adds the nuncConnectionId to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) SetNuncConnectionID(nuncConnectionID string) {
	o.NuncConnectionID = nuncConnectionID
}

// WithPosition adds the position to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) WithPosition(position *int32) *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams {
	o.SetPosition(position)
	return o
}

// SetPosition adds the position to the patch v1 nunc connections nunc connection Id component groups group Id params
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) SetPosition(position *int32) {
	o.Position = position
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ComponentGroupID != nil {

		// form param component_group_id
		var frComponentGroupID string
		if o.ComponentGroupID != nil {
			frComponentGroupID = *o.ComponentGroupID
		}
		fComponentGroupID := frComponentGroupID
		if fComponentGroupID != "" {
			if err := r.SetFormParam("component_group_id", fComponentGroupID); err != nil {
				return err
			}
		}
	}

	// path param group_id
	if err := r.SetPathParam("group_id", o.GroupID); err != nil {
		return err
	}

	if o.Name != nil {

		// form param name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("name", fName); err != nil {
				return err
			}
		}
	}

	// path param nunc_connection_id
	if err := r.SetPathParam("nunc_connection_id", o.NuncConnectionID); err != nil {
		return err
	}

	if o.Position != nil {

		// form param position
		var frPosition int32
		if o.Position != nil {
			frPosition = *o.Position
		}
		fPosition := swag.FormatInt32(frPosition)
		if fPosition != "" {
			if err := r.SetFormParam("position", fPosition); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
