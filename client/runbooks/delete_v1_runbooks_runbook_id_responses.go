// Code generated by go-swagger; DO NOT EDIT.

package runbooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// DeleteV1RunbooksRunbookIDReader is a Reader for the DeleteV1RunbooksRunbookID structure.
type DeleteV1RunbooksRunbookIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1RunbooksRunbookIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1RunbooksRunbookIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/runbooks/{runbook_id}] deleteV1RunbooksRunbookId", response, response.Code())
	}
}

// NewDeleteV1RunbooksRunbookIDOK creates a DeleteV1RunbooksRunbookIDOK with default headers values
func NewDeleteV1RunbooksRunbookIDOK() *DeleteV1RunbooksRunbookIDOK {
	return &DeleteV1RunbooksRunbookIDOK{}
}

/*
DeleteV1RunbooksRunbookIDOK describes a response with status code 200, with default header values.

Delete a runbook and make it unavailable for any future incidents.
*/
type DeleteV1RunbooksRunbookIDOK struct {
	Payload *models.RunbookEntity
}

// IsSuccess returns true when this delete v1 runbooks runbook Id o k response has a 2xx status code
func (o *DeleteV1RunbooksRunbookIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 runbooks runbook Id o k response has a 3xx status code
func (o *DeleteV1RunbooksRunbookIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 runbooks runbook Id o k response has a 4xx status code
func (o *DeleteV1RunbooksRunbookIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 runbooks runbook Id o k response has a 5xx status code
func (o *DeleteV1RunbooksRunbookIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 runbooks runbook Id o k response a status code equal to that given
func (o *DeleteV1RunbooksRunbookIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete v1 runbooks runbook Id o k response
func (o *DeleteV1RunbooksRunbookIDOK) Code() int {
	return 200
}

func (o *DeleteV1RunbooksRunbookIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/runbooks/{runbook_id}][%d] deleteV1RunbooksRunbookIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1RunbooksRunbookIDOK) String() string {
	return fmt.Sprintf("[DELETE /v1/runbooks/{runbook_id}][%d] deleteV1RunbooksRunbookIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1RunbooksRunbookIDOK) GetPayload() *models.RunbookEntity {
	return o.Payload
}

func (o *DeleteV1RunbooksRunbookIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
