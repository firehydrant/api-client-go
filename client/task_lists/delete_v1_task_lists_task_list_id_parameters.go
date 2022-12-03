// Code generated by go-swagger; DO NOT EDIT.

package task_lists

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

// NewDeleteV1TaskListsTaskListIDParams creates a new DeleteV1TaskListsTaskListIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1TaskListsTaskListIDParams() *DeleteV1TaskListsTaskListIDParams {
	return &DeleteV1TaskListsTaskListIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1TaskListsTaskListIDParamsWithTimeout creates a new DeleteV1TaskListsTaskListIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1TaskListsTaskListIDParamsWithTimeout(timeout time.Duration) *DeleteV1TaskListsTaskListIDParams {
	return &DeleteV1TaskListsTaskListIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1TaskListsTaskListIDParamsWithContext creates a new DeleteV1TaskListsTaskListIDParams object
// with the ability to set a context for a request.
func NewDeleteV1TaskListsTaskListIDParamsWithContext(ctx context.Context) *DeleteV1TaskListsTaskListIDParams {
	return &DeleteV1TaskListsTaskListIDParams{
		Context: ctx,
	}
}

// NewDeleteV1TaskListsTaskListIDParamsWithHTTPClient creates a new DeleteV1TaskListsTaskListIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1TaskListsTaskListIDParamsWithHTTPClient(client *http.Client) *DeleteV1TaskListsTaskListIDParams {
	return &DeleteV1TaskListsTaskListIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1TaskListsTaskListIDParams contains all the parameters to send to the API endpoint

	for the delete v1 task lists task list Id operation.

	Typically these are written to a http.Request.
*/
type DeleteV1TaskListsTaskListIDParams struct {

	// TaskListID.
	//
	// Format: int32
	TaskListID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 task lists task list Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1TaskListsTaskListIDParams) WithDefaults() *DeleteV1TaskListsTaskListIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 task lists task list Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1TaskListsTaskListIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 task lists task list Id params
func (o *DeleteV1TaskListsTaskListIDParams) WithTimeout(timeout time.Duration) *DeleteV1TaskListsTaskListIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 task lists task list Id params
func (o *DeleteV1TaskListsTaskListIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 task lists task list Id params
func (o *DeleteV1TaskListsTaskListIDParams) WithContext(ctx context.Context) *DeleteV1TaskListsTaskListIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 task lists task list Id params
func (o *DeleteV1TaskListsTaskListIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 task lists task list Id params
func (o *DeleteV1TaskListsTaskListIDParams) WithHTTPClient(client *http.Client) *DeleteV1TaskListsTaskListIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 task lists task list Id params
func (o *DeleteV1TaskListsTaskListIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTaskListID adds the taskListID to the delete v1 task lists task list Id params
func (o *DeleteV1TaskListsTaskListIDParams) WithTaskListID(taskListID int32) *DeleteV1TaskListsTaskListIDParams {
	o.SetTaskListID(taskListID)
	return o
}

// SetTaskListID adds the taskListId to the delete v1 task lists task list Id params
func (o *DeleteV1TaskListsTaskListIDParams) SetTaskListID(taskListID int32) {
	o.TaskListID = taskListID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1TaskListsTaskListIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param task_list_id
	if err := r.SetPathParam("task_list_id", swag.FormatInt32(o.TaskListID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
