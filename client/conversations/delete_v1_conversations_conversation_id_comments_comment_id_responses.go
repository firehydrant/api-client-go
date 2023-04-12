// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1ConversationsConversationIDCommentsCommentIDReader is a Reader for the DeleteV1ConversationsConversationIDCommentsCommentID structure.
type DeleteV1ConversationsConversationIDCommentsCommentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1ConversationsConversationIDCommentsCommentIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1ConversationsConversationIDCommentsCommentIDNoContent creates a DeleteV1ConversationsConversationIDCommentsCommentIDNoContent with default headers values
func NewDeleteV1ConversationsConversationIDCommentsCommentIDNoContent() *DeleteV1ConversationsConversationIDCommentsCommentIDNoContent {
	return &DeleteV1ConversationsConversationIDCommentsCommentIDNoContent{}
}

/*
DeleteV1ConversationsConversationIDCommentsCommentIDNoContent describes a response with status code 204, with default header values.

ALPHA - Archive a comment
*/
type DeleteV1ConversationsConversationIDCommentsCommentIDNoContent struct {
}

// IsSuccess returns true when this delete v1 conversations conversation Id comments comment Id no content response has a 2xx status code
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 conversations conversation Id comments comment Id no content response has a 3xx status code
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 conversations conversation Id comments comment Id no content response has a 4xx status code
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 conversations conversation Id comments comment Id no content response has a 5xx status code
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 conversations conversation Id comments comment Id no content response a status code equal to that given
func (o *DeleteV1ConversationsConversationIDCommentsCommentIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteV1ConversationsConversationIDCommentsCommentIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/conversations/{conversation_id}/comments/{comment_id}][%d] deleteV1ConversationsConversationIdCommentsCommentIdNoContent ", 204)
}

func (o *DeleteV1ConversationsConversationIDCommentsCommentIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/conversations/{conversation_id}/comments/{comment_id}][%d] deleteV1ConversationsConversationIdCommentsCommentIdNoContent ", 204)
}

func (o *DeleteV1ConversationsConversationIDCommentsCommentIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
