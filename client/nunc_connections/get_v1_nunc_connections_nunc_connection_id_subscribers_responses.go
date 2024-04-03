// Code generated by go-swagger; DO NOT EDIT.

package nunc_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1NuncConnectionsNuncConnectionIDSubscribersReader is a Reader for the GetV1NuncConnectionsNuncConnectionIDSubscribers structure.
type GetV1NuncConnectionsNuncConnectionIDSubscribersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1NuncConnectionsNuncConnectionIDSubscribersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/nunc_connections/{nunc_connection_id}/subscribers] getV1NuncConnectionsNuncConnectionIdSubscribers", response, response.Code())
	}
}

// NewGetV1NuncConnectionsNuncConnectionIDSubscribersOK creates a GetV1NuncConnectionsNuncConnectionIDSubscribersOK with default headers values
func NewGetV1NuncConnectionsNuncConnectionIDSubscribersOK() *GetV1NuncConnectionsNuncConnectionIDSubscribersOK {
	return &GetV1NuncConnectionsNuncConnectionIDSubscribersOK{}
}

/*
GetV1NuncConnectionsNuncConnectionIDSubscribersOK describes a response with status code 200, with default header values.

Retrieves the list of subscribers for a status page.
*/
type GetV1NuncConnectionsNuncConnectionIDSubscribersOK struct {
	Payload *models.NuncEmailSubscribersEntity
}

// IsSuccess returns true when this get v1 nunc connections nunc connection Id subscribers o k response has a 2xx status code
func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 nunc connections nunc connection Id subscribers o k response has a 3xx status code
func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 nunc connections nunc connection Id subscribers o k response has a 4xx status code
func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 nunc connections nunc connection Id subscribers o k response has a 5xx status code
func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 nunc connections nunc connection Id subscribers o k response a status code equal to that given
func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 nunc connections nunc connection Id subscribers o k response
func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersOK) Code() int {
	return 200
}

func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersOK) Error() string {
	return fmt.Sprintf("[GET /v1/nunc_connections/{nunc_connection_id}/subscribers][%d] getV1NuncConnectionsNuncConnectionIdSubscribersOK  %+v", 200, o.Payload)
}

func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersOK) String() string {
	return fmt.Sprintf("[GET /v1/nunc_connections/{nunc_connection_id}/subscribers][%d] getV1NuncConnectionsNuncConnectionIdSubscribersOK  %+v", 200, o.Payload)
}

func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersOK) GetPayload() *models.NuncEmailSubscribersEntity {
	return o.Payload
}

func (o *GetV1NuncConnectionsNuncConnectionIDSubscribersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NuncEmailSubscribersEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
