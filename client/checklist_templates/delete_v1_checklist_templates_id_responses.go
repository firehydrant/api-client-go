// Code generated by go-swagger; DO NOT EDIT.

package checklist_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// DeleteV1ChecklistTemplatesIDReader is a Reader for the DeleteV1ChecklistTemplatesID structure.
type DeleteV1ChecklistTemplatesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1ChecklistTemplatesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1ChecklistTemplatesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1ChecklistTemplatesIDOK creates a DeleteV1ChecklistTemplatesIDOK with default headers values
func NewDeleteV1ChecklistTemplatesIDOK() *DeleteV1ChecklistTemplatesIDOK {
	return &DeleteV1ChecklistTemplatesIDOK{}
}

/*
DeleteV1ChecklistTemplatesIDOK describes a response with status code 200, with default header values.

Archive a checklist template
*/
type DeleteV1ChecklistTemplatesIDOK struct {
	Payload *models.ChecklistTemplateEntity
}

// IsSuccess returns true when this delete v1 checklist templates Id o k response has a 2xx status code
func (o *DeleteV1ChecklistTemplatesIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 checklist templates Id o k response has a 3xx status code
func (o *DeleteV1ChecklistTemplatesIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 checklist templates Id o k response has a 4xx status code
func (o *DeleteV1ChecklistTemplatesIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 checklist templates Id o k response has a 5xx status code
func (o *DeleteV1ChecklistTemplatesIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 checklist templates Id o k response a status code equal to that given
func (o *DeleteV1ChecklistTemplatesIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteV1ChecklistTemplatesIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/checklist_templates/{id}][%d] deleteV1ChecklistTemplatesIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1ChecklistTemplatesIDOK) String() string {
	return fmt.Sprintf("[DELETE /v1/checklist_templates/{id}][%d] deleteV1ChecklistTemplatesIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1ChecklistTemplatesIDOK) GetPayload() *models.ChecklistTemplateEntity {
	return o.Payload
}

func (o *DeleteV1ChecklistTemplatesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChecklistTemplateEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}