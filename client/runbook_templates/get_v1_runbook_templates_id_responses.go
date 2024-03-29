// Code generated by go-swagger; DO NOT EDIT.

package runbook_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1RunbookTemplatesIDReader is a Reader for the GetV1RunbookTemplatesID structure.
type GetV1RunbookTemplatesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1RunbookTemplatesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1RunbookTemplatesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1RunbookTemplatesIDOK creates a GetV1RunbookTemplatesIDOK with default headers values
func NewGetV1RunbookTemplatesIDOK() *GetV1RunbookTemplatesIDOK {
	return &GetV1RunbookTemplatesIDOK{}
}

/*
GetV1RunbookTemplatesIDOK describes a response with status code 200, with default header values.

Retrieve a single runbook template
*/
type GetV1RunbookTemplatesIDOK struct {
}

// IsSuccess returns true when this get v1 runbook templates Id o k response has a 2xx status code
func (o *GetV1RunbookTemplatesIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 runbook templates Id o k response has a 3xx status code
func (o *GetV1RunbookTemplatesIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 runbook templates Id o k response has a 4xx status code
func (o *GetV1RunbookTemplatesIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 runbook templates Id o k response has a 5xx status code
func (o *GetV1RunbookTemplatesIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 runbook templates Id o k response a status code equal to that given
func (o *GetV1RunbookTemplatesIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1RunbookTemplatesIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/runbook_templates/{id}][%d] getV1RunbookTemplatesIdOK ", 200)
}

func (o *GetV1RunbookTemplatesIDOK) String() string {
	return fmt.Sprintf("[GET /v1/runbook_templates/{id}][%d] getV1RunbookTemplatesIdOK ", 200)
}

func (o *GetV1RunbookTemplatesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
