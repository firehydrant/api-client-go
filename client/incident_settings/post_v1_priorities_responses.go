// Code generated by go-swagger; DO NOT EDIT.

package incident_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1PrioritiesReader is a Reader for the PostV1Priorities structure.
type PostV1PrioritiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1PrioritiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1PrioritiesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1PrioritiesCreated creates a PostV1PrioritiesCreated with default headers values
func NewPostV1PrioritiesCreated() *PostV1PrioritiesCreated {
	return &PostV1PrioritiesCreated{}
}

/*
PostV1PrioritiesCreated describes a response with status code 201, with default header values.

Create a new priority
*/
type PostV1PrioritiesCreated struct {
	Payload *models.PriorityEntity
}

// IsSuccess returns true when this post v1 priorities created response has a 2xx status code
func (o *PostV1PrioritiesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 priorities created response has a 3xx status code
func (o *PostV1PrioritiesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 priorities created response has a 4xx status code
func (o *PostV1PrioritiesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 priorities created response has a 5xx status code
func (o *PostV1PrioritiesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 priorities created response a status code equal to that given
func (o *PostV1PrioritiesCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1PrioritiesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/priorities][%d] postV1PrioritiesCreated  %+v", 201, o.Payload)
}

func (o *PostV1PrioritiesCreated) String() string {
	return fmt.Sprintf("[POST /v1/priorities][%d] postV1PrioritiesCreated  %+v", 201, o.Payload)
}

func (o *PostV1PrioritiesCreated) GetPayload() *models.PriorityEntity {
	return o.Payload
}

func (o *PostV1PrioritiesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PriorityEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
