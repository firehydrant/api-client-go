// Code generated by go-swagger; DO NOT EDIT.

package status_pages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PutV1NuncConnectionsNuncConnectionIDReader is a Reader for the PutV1NuncConnectionsNuncConnectionID structure.
type PutV1NuncConnectionsNuncConnectionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1NuncConnectionsNuncConnectionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutV1NuncConnectionsNuncConnectionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutV1NuncConnectionsNuncConnectionIDOK creates a PutV1NuncConnectionsNuncConnectionIDOK with default headers values
func NewPutV1NuncConnectionsNuncConnectionIDOK() *PutV1NuncConnectionsNuncConnectionIDOK {
	return &PutV1NuncConnectionsNuncConnectionIDOK{}
}

/*
PutV1NuncConnectionsNuncConnectionIDOK describes a response with status code 200, with default header values.

Update your company's information and other components in the specified FireHydrant hosted status page.
*/
type PutV1NuncConnectionsNuncConnectionIDOK struct {
	Payload *models.NuncConnectionEntity
}

// IsSuccess returns true when this put v1 nunc connections nunc connection Id o k response has a 2xx status code
func (o *PutV1NuncConnectionsNuncConnectionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put v1 nunc connections nunc connection Id o k response has a 3xx status code
func (o *PutV1NuncConnectionsNuncConnectionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 nunc connections nunc connection Id o k response has a 4xx status code
func (o *PutV1NuncConnectionsNuncConnectionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put v1 nunc connections nunc connection Id o k response has a 5xx status code
func (o *PutV1NuncConnectionsNuncConnectionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 nunc connections nunc connection Id o k response a status code equal to that given
func (o *PutV1NuncConnectionsNuncConnectionIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutV1NuncConnectionsNuncConnectionIDOK) Error() string {
	return fmt.Sprintf("[PUT /v1/nunc_connections/{nunc_connection_id}][%d] putV1NuncConnectionsNuncConnectionIdOK  %+v", 200, o.Payload)
}

func (o *PutV1NuncConnectionsNuncConnectionIDOK) String() string {
	return fmt.Sprintf("[PUT /v1/nunc_connections/{nunc_connection_id}][%d] putV1NuncConnectionsNuncConnectionIdOK  %+v", 200, o.Payload)
}

func (o *PutV1NuncConnectionsNuncConnectionIDOK) GetPayload() *models.NuncConnectionEntity {
	return o.Payload
}

func (o *PutV1NuncConnectionsNuncConnectionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NuncConnectionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
