// Code generated by go-swagger; DO NOT EDIT.

package signals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1SignalsEmailTargetsIDReader is a Reader for the DeleteV1SignalsEmailTargetsID structure.
type DeleteV1SignalsEmailTargetsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1SignalsEmailTargetsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1SignalsEmailTargetsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1SignalsEmailTargetsIDNoContent creates a DeleteV1SignalsEmailTargetsIDNoContent with default headers values
func NewDeleteV1SignalsEmailTargetsIDNoContent() *DeleteV1SignalsEmailTargetsIDNoContent {
	return &DeleteV1SignalsEmailTargetsIDNoContent{}
}

/*
DeleteV1SignalsEmailTargetsIDNoContent describes a response with status code 204, with default header values.

Delete a Signals email target by ID
*/
type DeleteV1SignalsEmailTargetsIDNoContent struct {
}

// IsSuccess returns true when this delete v1 signals email targets Id no content response has a 2xx status code
func (o *DeleteV1SignalsEmailTargetsIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 signals email targets Id no content response has a 3xx status code
func (o *DeleteV1SignalsEmailTargetsIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 signals email targets Id no content response has a 4xx status code
func (o *DeleteV1SignalsEmailTargetsIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 signals email targets Id no content response has a 5xx status code
func (o *DeleteV1SignalsEmailTargetsIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 signals email targets Id no content response a status code equal to that given
func (o *DeleteV1SignalsEmailTargetsIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteV1SignalsEmailTargetsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/signals/email_targets/{id}][%d] deleteV1SignalsEmailTargetsIdNoContent ", 204)
}

func (o *DeleteV1SignalsEmailTargetsIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/signals/email_targets/{id}][%d] deleteV1SignalsEmailTargetsIdNoContent ", 204)
}

func (o *DeleteV1SignalsEmailTargetsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
