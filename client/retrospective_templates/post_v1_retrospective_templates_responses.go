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

// PostV1RetrospectiveTemplatesReader is a Reader for the PostV1RetrospectiveTemplates structure.
type PostV1RetrospectiveTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1RetrospectiveTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1RetrospectiveTemplatesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1RetrospectiveTemplatesCreated creates a PostV1RetrospectiveTemplatesCreated with default headers values
func NewPostV1RetrospectiveTemplatesCreated() *PostV1RetrospectiveTemplatesCreated {
	return &PostV1RetrospectiveTemplatesCreated{}
}

/*
PostV1RetrospectiveTemplatesCreated describes a response with status code 201, with default header values.

Create a new retrospective template
*/
type PostV1RetrospectiveTemplatesCreated struct {
	Payload *models.RetrospectivesTemplateEntity
}

// IsSuccess returns true when this post v1 retrospective templates created response has a 2xx status code
func (o *PostV1RetrospectiveTemplatesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 retrospective templates created response has a 3xx status code
func (o *PostV1RetrospectiveTemplatesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 retrospective templates created response has a 4xx status code
func (o *PostV1RetrospectiveTemplatesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 retrospective templates created response has a 5xx status code
func (o *PostV1RetrospectiveTemplatesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 retrospective templates created response a status code equal to that given
func (o *PostV1RetrospectiveTemplatesCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1RetrospectiveTemplatesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/retrospective_templates][%d] postV1RetrospectiveTemplatesCreated  %+v", 201, o.Payload)
}

func (o *PostV1RetrospectiveTemplatesCreated) String() string {
	return fmt.Sprintf("[POST /v1/retrospective_templates][%d] postV1RetrospectiveTemplatesCreated  %+v", 201, o.Payload)
}

func (o *PostV1RetrospectiveTemplatesCreated) GetPayload() *models.RetrospectivesTemplateEntity {
	return o.Payload
}

func (o *PostV1RetrospectiveTemplatesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RetrospectivesTemplateEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
