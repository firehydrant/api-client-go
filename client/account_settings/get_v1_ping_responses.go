// Code generated by go-swagger; DO NOT EDIT.

package account_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1PingReader is a Reader for the GetV1Ping structure.
type GetV1PingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1PingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1PingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1PingOK creates a GetV1PingOK with default headers values
func NewGetV1PingOK() *GetV1PingOK {
	return &GetV1PingOK{}
}

/*
GetV1PingOK describes a response with status code 200, with default header values.

Simple endpoint to verify your API connection is working
*/
type GetV1PingOK struct {
	Payload *models.PongEntity
}

// IsSuccess returns true when this get v1 ping o k response has a 2xx status code
func (o *GetV1PingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 ping o k response has a 3xx status code
func (o *GetV1PingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 ping o k response has a 4xx status code
func (o *GetV1PingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 ping o k response has a 5xx status code
func (o *GetV1PingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 ping o k response a status code equal to that given
func (o *GetV1PingOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1PingOK) Error() string {
	return fmt.Sprintf("[GET /v1/ping][%d] getV1PingOK  %+v", 200, o.Payload)
}

func (o *GetV1PingOK) String() string {
	return fmt.Sprintf("[GET /v1/ping][%d] getV1PingOK  %+v", 200, o.Payload)
}

func (o *GetV1PingOK) GetPayload() *models.PongEntity {
	return o.Payload
}

func (o *GetV1PingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PongEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
