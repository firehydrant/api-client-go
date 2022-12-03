// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1UsersIDServicesReader is a Reader for the GetV1UsersIDServices structure.
type GetV1UsersIDServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1UsersIDServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1UsersIDServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1UsersIDServicesOK creates a GetV1UsersIDServicesOK with default headers values
func NewGetV1UsersIDServicesOK() *GetV1UsersIDServicesOK {
	return &GetV1UsersIDServicesOK{}
}

/*
GetV1UsersIDServicesOK describes a response with status code 200, with default header values.

Retrieves a list of services owned by the teams a user is on
*/
type GetV1UsersIDServicesOK struct {
	Payload []*models.TeamEntityPaginated
}

// IsSuccess returns true when this get v1 users Id services o k response has a 2xx status code
func (o *GetV1UsersIDServicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 users Id services o k response has a 3xx status code
func (o *GetV1UsersIDServicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 users Id services o k response has a 4xx status code
func (o *GetV1UsersIDServicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 users Id services o k response has a 5xx status code
func (o *GetV1UsersIDServicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 users Id services o k response a status code equal to that given
func (o *GetV1UsersIDServicesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1UsersIDServicesOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/{id}/services][%d] getV1UsersIdServicesOK  %+v", 200, o.Payload)
}

func (o *GetV1UsersIDServicesOK) String() string {
	return fmt.Sprintf("[GET /v1/users/{id}/services][%d] getV1UsersIdServicesOK  %+v", 200, o.Payload)
}

func (o *GetV1UsersIDServicesOK) GetPayload() []*models.TeamEntityPaginated {
	return o.Payload
}

func (o *GetV1UsersIDServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
