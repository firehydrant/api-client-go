// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1IntegrationsReader is a Reader for the GetV1Integrations structure.
type GetV1IntegrationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IntegrationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IntegrationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1IntegrationsOK creates a GetV1IntegrationsOK with default headers values
func NewGetV1IntegrationsOK() *GetV1IntegrationsOK {
	return &GetV1IntegrationsOK{}
}

/*
GetV1IntegrationsOK describes a response with status code 200, with default header values.

get Integration(s)
*/
type GetV1IntegrationsOK struct {
}

// IsSuccess returns true when this get v1 integrations o k response has a 2xx status code
func (o *GetV1IntegrationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 integrations o k response has a 3xx status code
func (o *GetV1IntegrationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 integrations o k response has a 4xx status code
func (o *GetV1IntegrationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 integrations o k response has a 5xx status code
func (o *GetV1IntegrationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 integrations o k response a status code equal to that given
func (o *GetV1IntegrationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1IntegrationsOK) Error() string {
	return fmt.Sprintf("[GET /v1/integrations][%d] getV1IntegrationsOK ", 200)
}

func (o *GetV1IntegrationsOK) String() string {
	return fmt.Sprintf("[GET /v1/integrations][%d] getV1IntegrationsOK ", 200)
}

func (o *GetV1IntegrationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
