// Code generated by go-swagger; DO NOT EDIT.

package runbook_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1RunbookTemplatesReader is a Reader for the GetV1RunbookTemplates structure.
type GetV1RunbookTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1RunbookTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1RunbookTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/runbook_templates] getV1RunbookTemplates", response, response.Code())
	}
}

// NewGetV1RunbookTemplatesOK creates a GetV1RunbookTemplatesOK with default headers values
func NewGetV1RunbookTemplatesOK() *GetV1RunbookTemplatesOK {
	return &GetV1RunbookTemplatesOK{}
}

/*
GetV1RunbookTemplatesOK describes a response with status code 200, with default header values.

List all available runbook templates
*/
type GetV1RunbookTemplatesOK struct {
}

// IsSuccess returns true when this get v1 runbook templates o k response has a 2xx status code
func (o *GetV1RunbookTemplatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 runbook templates o k response has a 3xx status code
func (o *GetV1RunbookTemplatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 runbook templates o k response has a 4xx status code
func (o *GetV1RunbookTemplatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 runbook templates o k response has a 5xx status code
func (o *GetV1RunbookTemplatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 runbook templates o k response a status code equal to that given
func (o *GetV1RunbookTemplatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 runbook templates o k response
func (o *GetV1RunbookTemplatesOK) Code() int {
	return 200
}

func (o *GetV1RunbookTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /v1/runbook_templates][%d] getV1RunbookTemplatesOK ", 200)
}

func (o *GetV1RunbookTemplatesOK) String() string {
	return fmt.Sprintf("[GET /v1/runbook_templates][%d] getV1RunbookTemplatesOK ", 200)
}

func (o *GetV1RunbookTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
