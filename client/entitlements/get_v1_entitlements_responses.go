// Code generated by go-swagger; DO NOT EDIT.

package entitlements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1EntitlementsReader is a Reader for the GetV1Entitlements structure.
type GetV1EntitlementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1EntitlementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1EntitlementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/entitlements] getV1Entitlements", response, response.Code())
	}
}

// NewGetV1EntitlementsOK creates a GetV1EntitlementsOK with default headers values
func NewGetV1EntitlementsOK() *GetV1EntitlementsOK {
	return &GetV1EntitlementsOK{}
}

/*
GetV1EntitlementsOK describes a response with status code 200, with default header values.

Retrieve all entitlements
*/
type GetV1EntitlementsOK struct {
	Payload *models.EntitlementEntityPaginated
}

// IsSuccess returns true when this get v1 entitlements o k response has a 2xx status code
func (o *GetV1EntitlementsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 entitlements o k response has a 3xx status code
func (o *GetV1EntitlementsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 entitlements o k response has a 4xx status code
func (o *GetV1EntitlementsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 entitlements o k response has a 5xx status code
func (o *GetV1EntitlementsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 entitlements o k response a status code equal to that given
func (o *GetV1EntitlementsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 entitlements o k response
func (o *GetV1EntitlementsOK) Code() int {
	return 200
}

func (o *GetV1EntitlementsOK) Error() string {
	return fmt.Sprintf("[GET /v1/entitlements][%d] getV1EntitlementsOK  %+v", 200, o.Payload)
}

func (o *GetV1EntitlementsOK) String() string {
	return fmt.Sprintf("[GET /v1/entitlements][%d] getV1EntitlementsOK  %+v", 200, o.Payload)
}

func (o *GetV1EntitlementsOK) GetPayload() *models.EntitlementEntityPaginated {
	return o.Payload
}

func (o *GetV1EntitlementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EntitlementEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
