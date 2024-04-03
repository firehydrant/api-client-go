// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1EnvironmentsReader is a Reader for the PostV1Environments structure.
type PostV1EnvironmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1EnvironmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1EnvironmentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/environments] postV1Environments", response, response.Code())
	}
}

// NewPostV1EnvironmentsCreated creates a PostV1EnvironmentsCreated with default headers values
func NewPostV1EnvironmentsCreated() *PostV1EnvironmentsCreated {
	return &PostV1EnvironmentsCreated{}
}

/*
PostV1EnvironmentsCreated describes a response with status code 201, with default header values.

Creates an environment for the organization
*/
type PostV1EnvironmentsCreated struct {
	Payload *models.EnvironmentEntryEntity
}

// IsSuccess returns true when this post v1 environments created response has a 2xx status code
func (o *PostV1EnvironmentsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 environments created response has a 3xx status code
func (o *PostV1EnvironmentsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 environments created response has a 4xx status code
func (o *PostV1EnvironmentsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 environments created response has a 5xx status code
func (o *PostV1EnvironmentsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 environments created response a status code equal to that given
func (o *PostV1EnvironmentsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post v1 environments created response
func (o *PostV1EnvironmentsCreated) Code() int {
	return 201
}

func (o *PostV1EnvironmentsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/environments][%d] postV1EnvironmentsCreated  %+v", 201, o.Payload)
}

func (o *PostV1EnvironmentsCreated) String() string {
	return fmt.Sprintf("[POST /v1/environments][%d] postV1EnvironmentsCreated  %+v", 201, o.Payload)
}

func (o *PostV1EnvironmentsCreated) GetPayload() *models.EnvironmentEntryEntity {
	return o.Payload
}

func (o *PostV1EnvironmentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentEntryEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
