// Code generated by go-swagger; DO NOT EDIT.

package conversations

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

// NewPatchV1ConversationsConversationIDCommentsCommentIDParams creates a new PatchV1ConversationsConversationIDCommentsCommentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1ConversationsConversationIDCommentsCommentIDParams() *PatchV1ConversationsConversationIDCommentsCommentIDParams {
	return &PatchV1ConversationsConversationIDCommentsCommentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1ConversationsConversationIDCommentsCommentIDParamsWithTimeout creates a new PatchV1ConversationsConversationIDCommentsCommentIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1ConversationsConversationIDCommentsCommentIDParamsWithTimeout(timeout time.Duration) *PatchV1ConversationsConversationIDCommentsCommentIDParams {
	return &PatchV1ConversationsConversationIDCommentsCommentIDParams{
		timeout: timeout,
	}
}

// NewPatchV1ConversationsConversationIDCommentsCommentIDParamsWithContext creates a new PatchV1ConversationsConversationIDCommentsCommentIDParams object
// with the ability to set a context for a request.
func NewPatchV1ConversationsConversationIDCommentsCommentIDParamsWithContext(ctx context.Context) *PatchV1ConversationsConversationIDCommentsCommentIDParams {
	return &PatchV1ConversationsConversationIDCommentsCommentIDParams{
		Context: ctx,
	}
}

// NewPatchV1ConversationsConversationIDCommentsCommentIDParamsWithHTTPClient creates a new PatchV1ConversationsConversationIDCommentsCommentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1ConversationsConversationIDCommentsCommentIDParamsWithHTTPClient(client *http.Client) *PatchV1ConversationsConversationIDCommentsCommentIDParams {
	return &PatchV1ConversationsConversationIDCommentsCommentIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1ConversationsConversationIDCommentsCommentIDParams contains all the parameters to send to the API endpoint

	for the patch v1 conversations conversation Id comments comment Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1ConversationsConversationIDCommentsCommentIDParams struct {

	// CommentID.
	CommentID string

	// ConversationID.
	ConversationID string

	// PatchV1ConversationsConversationIDCommentsCommentID.
	PatchV1ConversationsConversationIDCommentsCommentID *models.PatchV1ConversationsConversationIDCommentsCommentID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 conversations conversation Id comments comment Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) WithDefaults() *PatchV1ConversationsConversationIDCommentsCommentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 conversations conversation Id comments comment Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 conversations conversation Id comments comment Id params
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) WithTimeout(timeout time.Duration) *PatchV1ConversationsConversationIDCommentsCommentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 conversations conversation Id comments comment Id params
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 conversations conversation Id comments comment Id params
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) WithContext(ctx context.Context) *PatchV1ConversationsConversationIDCommentsCommentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 conversations conversation Id comments comment Id params
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 conversations conversation Id comments comment Id params
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) WithHTTPClient(client *http.Client) *PatchV1ConversationsConversationIDCommentsCommentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 conversations conversation Id comments comment Id params
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommentID adds the commentID to the patch v1 conversations conversation Id comments comment Id params
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) WithCommentID(commentID string) *PatchV1ConversationsConversationIDCommentsCommentIDParams {
	o.SetCommentID(commentID)
	return o
}

// SetCommentID adds the commentId to the patch v1 conversations conversation Id comments comment Id params
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) SetCommentID(commentID string) {
	o.CommentID = commentID
}

// WithConversationID adds the conversationID to the patch v1 conversations conversation Id comments comment Id params
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) WithConversationID(conversationID string) *PatchV1ConversationsConversationIDCommentsCommentIDParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the patch v1 conversations conversation Id comments comment Id params
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithPatchV1ConversationsConversationIDCommentsCommentID adds the patchV1ConversationsConversationIDCommentsCommentID to the patch v1 conversations conversation Id comments comment Id params
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) WithPatchV1ConversationsConversationIDCommentsCommentID(patchV1ConversationsConversationIDCommentsCommentID *models.PatchV1ConversationsConversationIDCommentsCommentID) *PatchV1ConversationsConversationIDCommentsCommentIDParams {
	o.SetPatchV1ConversationsConversationIDCommentsCommentID(patchV1ConversationsConversationIDCommentsCommentID)
	return o
}

// SetPatchV1ConversationsConversationIDCommentsCommentID adds the patchV1ConversationsConversationIdCommentsCommentId to the patch v1 conversations conversation Id comments comment Id params
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) SetPatchV1ConversationsConversationIDCommentsCommentID(patchV1ConversationsConversationIDCommentsCommentID *models.PatchV1ConversationsConversationIDCommentsCommentID) {
	o.PatchV1ConversationsConversationIDCommentsCommentID = patchV1ConversationsConversationIDCommentsCommentID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1ConversationsConversationIDCommentsCommentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param comment_id
	if err := r.SetPathParam("comment_id", o.CommentID); err != nil {
		return err
	}

	// path param conversation_id
	if err := r.SetPathParam("conversation_id", o.ConversationID); err != nil {
		return err
	}
	if o.PatchV1ConversationsConversationIDCommentsCommentID != nil {
		if err := r.SetBodyParam(o.PatchV1ConversationsConversationIDCommentsCommentID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
