// Code generated by go-swagger; DO NOT EDIT.

package communication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1StatusUpdateTemplatesReader is a Reader for the PostV1StatusUpdateTemplates structure.
type PostV1StatusUpdateTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1StatusUpdateTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1StatusUpdateTemplatesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1StatusUpdateTemplatesCreated creates a PostV1StatusUpdateTemplatesCreated with default headers values
func NewPostV1StatusUpdateTemplatesCreated() *PostV1StatusUpdateTemplatesCreated {
	return &PostV1StatusUpdateTemplatesCreated{}
}

/*
PostV1StatusUpdateTemplatesCreated describes a response with status code 201, with default header values.

Create a status update template for your organization
*/
type PostV1StatusUpdateTemplatesCreated struct {
	Payload *models.StatusUpdateTemplateEntity
}

// IsSuccess returns true when this post v1 status update templates created response has a 2xx status code
func (o *PostV1StatusUpdateTemplatesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 status update templates created response has a 3xx status code
func (o *PostV1StatusUpdateTemplatesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 status update templates created response has a 4xx status code
func (o *PostV1StatusUpdateTemplatesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 status update templates created response has a 5xx status code
func (o *PostV1StatusUpdateTemplatesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 status update templates created response a status code equal to that given
func (o *PostV1StatusUpdateTemplatesCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1StatusUpdateTemplatesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/status_update_templates][%d] postV1StatusUpdateTemplatesCreated  %+v", 201, o.Payload)
}

func (o *PostV1StatusUpdateTemplatesCreated) String() string {
	return fmt.Sprintf("[POST /v1/status_update_templates][%d] postV1StatusUpdateTemplatesCreated  %+v", 201, o.Payload)
}

func (o *PostV1StatusUpdateTemplatesCreated) GetPayload() *models.StatusUpdateTemplateEntity {
	return o.Payload
}

func (o *PostV1StatusUpdateTemplatesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusUpdateTemplateEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
