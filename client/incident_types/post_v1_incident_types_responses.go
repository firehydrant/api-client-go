// Code generated by go-swagger; DO NOT EDIT.

package incident_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1IncidentTypesReader is a Reader for the PostV1IncidentTypes structure.
type PostV1IncidentTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1IncidentTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1IncidentTypesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1IncidentTypesCreated creates a PostV1IncidentTypesCreated with default headers values
func NewPostV1IncidentTypesCreated() *PostV1IncidentTypesCreated {
	return &PostV1IncidentTypesCreated{}
}

/*
PostV1IncidentTypesCreated describes a response with status code 201, with default header values.

Create a new incident type
*/
type PostV1IncidentTypesCreated struct {
	Payload *models.IncidentTypeEntity
}

// IsSuccess returns true when this post v1 incident types created response has a 2xx status code
func (o *PostV1IncidentTypesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 incident types created response has a 3xx status code
func (o *PostV1IncidentTypesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 incident types created response has a 4xx status code
func (o *PostV1IncidentTypesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 incident types created response has a 5xx status code
func (o *PostV1IncidentTypesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 incident types created response a status code equal to that given
func (o *PostV1IncidentTypesCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1IncidentTypesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/incident_types][%d] postV1IncidentTypesCreated  %+v", 201, o.Payload)
}

func (o *PostV1IncidentTypesCreated) String() string {
	return fmt.Sprintf("[POST /v1/incident_types][%d] postV1IncidentTypesCreated  %+v", 201, o.Payload)
}

func (o *PostV1IncidentTypesCreated) GetPayload() *models.IncidentTypeEntity {
	return o.Payload
}

func (o *PostV1IncidentTypesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncidentTypeEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
