// Code generated by go-swagger; DO NOT EDIT.

package retrospective_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1RetrospectiveTemplatesRetrospectiveTemplateIDReader is a Reader for the GetV1RetrospectiveTemplatesRetrospectiveTemplateID structure.
type GetV1RetrospectiveTemplatesRetrospectiveTemplateIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK creates a GetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK with default headers values
func NewGetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK() *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK {
	return &GetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK{}
}

/*
GetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK describes a response with status code 200, with default header values.

Retrieve a single retrospective template by ID
*/
type GetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK struct {
	Payload *models.RetrospectivesTemplateEntity
}

// IsSuccess returns true when this get v1 retrospective templates retrospective template Id o k response has a 2xx status code
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 retrospective templates retrospective template Id o k response has a 3xx status code
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 retrospective templates retrospective template Id o k response has a 4xx status code
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 retrospective templates retrospective template Id o k response has a 5xx status code
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 retrospective templates retrospective template Id o k response a status code equal to that given
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/retrospective_templates/{retrospective_template_id}][%d] getV1RetrospectiveTemplatesRetrospectiveTemplateIdOK  %+v", 200, o.Payload)
}

func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK) String() string {
	return fmt.Sprintf("[GET /v1/retrospective_templates/{retrospective_template_id}][%d] getV1RetrospectiveTemplatesRetrospectiveTemplateIdOK  %+v", 200, o.Payload)
}

func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK) GetPayload() *models.RetrospectivesTemplateEntity {
	return o.Payload
}

func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RetrospectivesTemplateEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
