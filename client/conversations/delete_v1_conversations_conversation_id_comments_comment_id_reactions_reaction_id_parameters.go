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
)

// NewDeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams creates a new DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams() *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams {
	return &DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParamsWithTimeout creates a new DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParamsWithTimeout(timeout time.Duration) *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams {
	return &DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParamsWithContext creates a new DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams object
// with the ability to set a context for a request.
func NewDeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParamsWithContext(ctx context.Context) *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams {
	return &DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams{
		Context: ctx,
	}
}

// NewDeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParamsWithHTTPClient creates a new DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParamsWithHTTPClient(client *http.Client) *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams {
	return &DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams contains all the parameters to send to the API endpoint

	for the delete v1 conversations conversation Id comments comment Id reactions reaction Id operation.

	Typically these are written to a http.Request.
*/
type DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams struct {

	// CommentID.
	CommentID string

	// ConversationID.
	ConversationID string

	// ReactionID.
	ReactionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 conversations conversation Id comments comment Id reactions reaction Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) WithDefaults() *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 conversations conversation Id comments comment Id reactions reaction Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 conversations conversation Id comments comment Id reactions reaction Id params
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) WithTimeout(timeout time.Duration) *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 conversations conversation Id comments comment Id reactions reaction Id params
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 conversations conversation Id comments comment Id reactions reaction Id params
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) WithContext(ctx context.Context) *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 conversations conversation Id comments comment Id reactions reaction Id params
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 conversations conversation Id comments comment Id reactions reaction Id params
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) WithHTTPClient(client *http.Client) *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 conversations conversation Id comments comment Id reactions reaction Id params
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommentID adds the commentID to the delete v1 conversations conversation Id comments comment Id reactions reaction Id params
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) WithCommentID(commentID string) *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams {
	o.SetCommentID(commentID)
	return o
}

// SetCommentID adds the commentId to the delete v1 conversations conversation Id comments comment Id reactions reaction Id params
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) SetCommentID(commentID string) {
	o.CommentID = commentID
}

// WithConversationID adds the conversationID to the delete v1 conversations conversation Id comments comment Id reactions reaction Id params
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) WithConversationID(conversationID string) *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the delete v1 conversations conversation Id comments comment Id reactions reaction Id params
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WithReactionID adds the reactionID to the delete v1 conversations conversation Id comments comment Id reactions reaction Id params
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) WithReactionID(reactionID string) *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams {
	o.SetReactionID(reactionID)
	return o
}

// SetReactionID adds the reactionId to the delete v1 conversations conversation Id comments comment Id reactions reaction Id params
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) SetReactionID(reactionID string) {
	o.ReactionID = reactionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReactionsReactionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param reaction_id
	if err := r.SetPathParam("reaction_id", o.ReactionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
