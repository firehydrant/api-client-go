// Code generated by go-swagger; DO NOT EDIT.

package signals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1SignalsWebhookTargetsIDReader is a Reader for the DeleteV1SignalsWebhookTargetsID structure.
type DeleteV1SignalsWebhookTargetsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1SignalsWebhookTargetsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1SignalsWebhookTargetsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1SignalsWebhookTargetsIDNoContent creates a DeleteV1SignalsWebhookTargetsIDNoContent with default headers values
func NewDeleteV1SignalsWebhookTargetsIDNoContent() *DeleteV1SignalsWebhookTargetsIDNoContent {
	return &DeleteV1SignalsWebhookTargetsIDNoContent{}
}

/*
DeleteV1SignalsWebhookTargetsIDNoContent describes a response with status code 204, with default header values.

Delete a Signals webhook target by ID
*/
type DeleteV1SignalsWebhookTargetsIDNoContent struct {
}

// IsSuccess returns true when this delete v1 signals webhook targets Id no content response has a 2xx status code
func (o *DeleteV1SignalsWebhookTargetsIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 signals webhook targets Id no content response has a 3xx status code
func (o *DeleteV1SignalsWebhookTargetsIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 signals webhook targets Id no content response has a 4xx status code
func (o *DeleteV1SignalsWebhookTargetsIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 signals webhook targets Id no content response has a 5xx status code
func (o *DeleteV1SignalsWebhookTargetsIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 signals webhook targets Id no content response a status code equal to that given
func (o *DeleteV1SignalsWebhookTargetsIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteV1SignalsWebhookTargetsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/signals/webhook_targets/{id}][%d] deleteV1SignalsWebhookTargetsIdNoContent ", 204)
}

func (o *DeleteV1SignalsWebhookTargetsIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/signals/webhook_targets/{id}][%d] deleteV1SignalsWebhookTargetsIdNoContent ", 204)
}

func (o *DeleteV1SignalsWebhookTargetsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
