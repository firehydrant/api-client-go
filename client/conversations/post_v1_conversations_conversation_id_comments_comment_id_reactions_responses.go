// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostV1ConversationsConversationIDCommentsCommentIDReactionsReader is a Reader for the PostV1ConversationsConversationIDCommentsCommentIDReactions structure.
type PostV1ConversationsConversationIDCommentsCommentIDReactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1ConversationsConversationIDCommentsCommentIDReactionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/conversations/{conversation_id}/comments/{comment_id}/reactions] postV1ConversationsConversationIdCommentsCommentIdReactions", response, response.Code())
	}
}

// NewPostV1ConversationsConversationIDCommentsCommentIDReactionsCreated creates a PostV1ConversationsConversationIDCommentsCommentIDReactionsCreated with default headers values
func NewPostV1ConversationsConversationIDCommentsCommentIDReactionsCreated() *PostV1ConversationsConversationIDCommentsCommentIDReactionsCreated {
	return &PostV1ConversationsConversationIDCommentsCommentIDReactionsCreated{}
}

/*
PostV1ConversationsConversationIDCommentsCommentIDReactionsCreated describes a response with status code 201, with default header values.

ALPHA - Create a reaction on a comment
*/
type PostV1ConversationsConversationIDCommentsCommentIDReactionsCreated struct {
}

// IsSuccess returns true when this post v1 conversations conversation Id comments comment Id reactions created response has a 2xx status code
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 conversations conversation Id comments comment Id reactions created response has a 3xx status code
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 conversations conversation Id comments comment Id reactions created response has a 4xx status code
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 conversations conversation Id comments comment Id reactions created response has a 5xx status code
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 conversations conversation Id comments comment Id reactions created response a status code equal to that given
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post v1 conversations conversation Id comments comment Id reactions created response
func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsCreated) Code() int {
	return 201
}

func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/conversations/{conversation_id}/comments/{comment_id}/reactions][%d] postV1ConversationsConversationIdCommentsCommentIdReactionsCreated ", 201)
}

func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsCreated) String() string {
	return fmt.Sprintf("[POST /v1/conversations/{conversation_id}/comments/{comment_id}/reactions][%d] postV1ConversationsConversationIdCommentsCommentIdReactionsCreated ", 201)
}

func (o *PostV1ConversationsConversationIDCommentsCommentIDReactionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
