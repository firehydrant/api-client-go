// Code generated by go-swagger; DO NOT EDIT.

package runbooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1RunbooksExecutionsReader is a Reader for the PostV1RunbooksExecutions structure.
type PostV1RunbooksExecutionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1RunbooksExecutionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1RunbooksExecutionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1RunbooksExecutionsCreated creates a PostV1RunbooksExecutionsCreated with default headers values
func NewPostV1RunbooksExecutionsCreated() *PostV1RunbooksExecutionsCreated {
	return &PostV1RunbooksExecutionsCreated{}
}

/*
PostV1RunbooksExecutionsCreated describes a response with status code 201, with default header values.

Attaches a runbook to an incident and executes it
*/
type PostV1RunbooksExecutionsCreated struct {
	Payload *models.RunbooksExecutionEntity
}

// IsSuccess returns true when this post v1 runbooks executions created response has a 2xx status code
func (o *PostV1RunbooksExecutionsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 runbooks executions created response has a 3xx status code
func (o *PostV1RunbooksExecutionsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 runbooks executions created response has a 4xx status code
func (o *PostV1RunbooksExecutionsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 runbooks executions created response has a 5xx status code
func (o *PostV1RunbooksExecutionsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 runbooks executions created response a status code equal to that given
func (o *PostV1RunbooksExecutionsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1RunbooksExecutionsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/runbooks/executions][%d] postV1RunbooksExecutionsCreated  %+v", 201, o.Payload)
}

func (o *PostV1RunbooksExecutionsCreated) String() string {
	return fmt.Sprintf("[POST /v1/runbooks/executions][%d] postV1RunbooksExecutionsCreated  %+v", 201, o.Payload)
}

func (o *PostV1RunbooksExecutionsCreated) GetPayload() *models.RunbooksExecutionEntity {
	return o.Payload
}

func (o *PostV1RunbooksExecutionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbooksExecutionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
