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
	"github.com/go-openapi/swag"

	"github.com/firehydrant/api-client-go/models"
)

// NewPostV1ConversationsConversationIDCommentsCommentIDReactionsParams creates a new PostV1ConversationsConversationIDCommentsCommentIDReactionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1ConversationsConversationIDCommentsCommentIDReactionsParams() *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	return &PostV1ConversationsConversationIDCommentsCommentIDReactionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1ConversationsConversationIDCommentsCommentIDReactionsParamsWithTimeout creates a new PostV1ConversationsConversationIDCommentsCommentIDReactionsParams object
// with the ability to set a timeout on a request.
func NewPostV1ConversationsConversationIDCommentsCommentIDReactionsParamsWithTimeout(timeout time.Duration) *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	return &PostV1ConversationsConversationIDCommentsCommentIDReactionsParams{
		timeout: timeout,
	}
}

// NewPostV1ConversationsConversationIDCommentsCommentIDReactionsParamsWithContext creates a new PostV1ConversationsConversationIDCommentsCommentIDReactionsParams object
// with the ability to set a context for a request.
func NewPostV1ConversationsConversationIDCommentsCommentIDReactionsParamsWithContext(ctx context.Context) *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	return &PostV1ConversationsConversationIDCommentsCommentIDReactionsParams{
		Context: ctx,
	}
}

// NewPostV1ConversationsConversationIDCommentsCommentIDReactionsParamsWithHTTPClient creates a new PostV1ConversationsConversationIDCommentsCommentIDReactionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1ConversationsConversationIDCommentsCommentIDReactionsParamsWithHTTPClient(client *http.Client) *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	return &PostV1ConversationsConversationIDCommentsCommentIDReactionsParams{
		HTTPClient: client,
	}
}

/*
PostV1ConversationsConversationIDCommentsCommentIDReactionsParams contains all the parameters to send to the API endpoint

	for the post v1 conversations conversation Id comments comment Id reactions operation.

	Typically these are written to a http.Request.
*/
type PostV1ConversationsConversationIDCommentsCommentIDReactionsParams struct {

	// CommentID.
	//
	// Format: int32
	CommentID int32

	// ConversationID.
	//
	// Format: int32
	ConversationID int32

	// PostV1ConversationsConversationIDCommentsCommentIDReactions.
	PostV1ConversationsConversationIDCommentsCommentIDReactions *models.PostV1ConversationsConversationIDCommentsCommentIDReactions

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 conversations conversation Id comments comment Id reactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) WithDefaults() *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 conversations conversation Id comments comment Id reactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 conversations conversation Id comments comment Id reactions params
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) WithTimeout(timeout time.Duration) *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 conversations conversation Id comments comment Id reactions params
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 conversations conversation Id comments comment Id reactions params
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) WithContext(ctx context.Context) *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 conversations conversation Id comments comment Id reactions params
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 conversations conversation Id comments comment Id reactions params
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) WithHTTPClient(client *http.Client) *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 conversations conversation Id comments comment Id reactions params
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommentID adds the commentID to the post v1 conversations conversation Id comments comment Id reactions params
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) WithCommentID(commentID int32) *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	o.SetCommentID(commentID)
	return o
}

// SetCommentID adds the commentId to the post v1 conversations conversation Id comments comment Id reactions params
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) SetCommentID(commentID int32) {
	o.CommentID = commentID
}

// WithConversationID adds the conversationID to the post v1 conversations conversation Id comments comment Id reactions params
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) WithConversationID(conversationID int32) *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the post v1 conversations conversation Id comments comment Id reactions params
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) SetConversationID(conversationID int32) {
	o.ConversationID = conversationID
}

// WithPostV1ConversationsConversationIDCommentsCommentIDReactions adds the postV1ConversationsConversationIDCommentsCommentIDReactions to the post v1 conversations conversation Id comments comment Id reactions params
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) WithPostV1ConversationsConversationIDCommentsCommentIDReactions(postV1ConversationsConversationIDCommentsCommentIDReactions *models.PostV1ConversationsConversationIDCommentsCommentIDReactions) *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams {
	o.SetPostV1ConversationsConversationIDCommentsCommentIDReactions(postV1ConversationsConversationIDCommentsCommentIDReactions)
	return o
}

// SetPostV1ConversationsConversationIDCommentsCommentIDReactions adds the postV1ConversationsConversationIdCommentsCommentIdReactions to the post v1 conversations conversation Id comments comment Id reactions params
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) SetPostV1ConversationsConversationIDCommentsCommentIDReactions(postV1ConversationsConversationIDCommentsCommentIDReactions *models.PostV1ConversationsConversationIDCommentsCommentIDReactions) {
	o.PostV1ConversationsConversationIDCommentsCommentIDReactions = postV1ConversationsConversationIDCommentsCommentIDReactions
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param comment_id
	if err := r.SetPathParam("comment_id", swag.FormatInt32(o.CommentID)); err != nil {
		return err
	}

	// path param conversation_id
	if err := r.SetPathParam("conversation_id", swag.FormatInt32(o.ConversationID)); err != nil {
		return err
	}
	if o.PostV1ConversationsConversationIDCommentsCommentIDReactions != nil {
		if err := r.SetBodyParam(o.PostV1ConversationsConversationIDCommentsCommentIDReactions); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
