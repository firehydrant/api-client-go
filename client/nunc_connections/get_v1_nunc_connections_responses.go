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

// GetV1NuncConnectionsReader is a Reader for the GetV1NuncConnections structure.
type GetV1NuncConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1NuncConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1NuncConnectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1NuncConnectionsOK creates a GetV1NuncConnectionsOK with default headers values
func NewGetV1NuncConnectionsOK() *GetV1NuncConnectionsOK {
	return &GetV1NuncConnectionsOK{}
}

/*
GetV1NuncConnectionsOK describes a response with status code 200, with default header values.

Lists the information displayed as part of your FireHydrant hosted status pages.
*/
type GetV1NuncConnectionsOK struct {
	Payload *models.NuncConnectionEntityPaginated
}

// IsSuccess returns true when this get v1 nunc connections o k response has a 2xx status code
func (o *GetV1NuncConnectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 nunc connections o k response has a 3xx status code
func (o *GetV1NuncConnectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 nunc connections o k response has a 4xx status code
func (o *GetV1NuncConnectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 nunc connections o k response has a 5xx status code
func (o *GetV1NuncConnectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 nunc connections o k response a status code equal to that given
func (o *GetV1NuncConnectionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1NuncConnectionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/nunc_connections][%d] getV1NuncConnectionsOK  %+v", 200, o.Payload)
}

func (o *GetV1NuncConnectionsOK) String() string {
	return fmt.Sprintf("[GET /v1/nunc_connections][%d] getV1NuncConnectionsOK  %+v", 200, o.Payload)
}

func (o *GetV1NuncConnectionsOK) GetPayload() *models.NuncConnectionEntityPaginated {
	return o.Payload
}

func (o *GetV1NuncConnectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NuncConnectionEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}