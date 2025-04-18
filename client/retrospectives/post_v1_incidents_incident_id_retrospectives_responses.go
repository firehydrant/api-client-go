// Code generated by go-swagger; DO NOT EDIT.

package retrospectives

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1IncidentsIncidentIDRetrospectivesReader is a Reader for the PostV1IncidentsIncidentIDRetrospectives structure.
type PostV1IncidentsIncidentIDRetrospectivesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1IncidentsIncidentIDRetrospectivesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1IncidentsIncidentIDRetrospectivesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1IncidentsIncidentIDRetrospectivesCreated creates a PostV1IncidentsIncidentIDRetrospectivesCreated with default headers values
func NewPostV1IncidentsIncidentIDRetrospectivesCreated() *PostV1IncidentsIncidentIDRetrospectivesCreated {
	return &PostV1IncidentsIncidentIDRetrospectivesCreated{}
}

/*
PostV1IncidentsIncidentIDRetrospectivesCreated describes a response with status code 201, with default header values.

Create a new retrospective for an incident
*/
type PostV1IncidentsIncidentIDRetrospectivesCreated struct {
	Payload *models.IncidentsRetrospectiveEntity
}

// IsSuccess returns true when this post v1 incidents incident Id retrospectives created response has a 2xx status code
func (o *PostV1IncidentsIncidentIDRetrospectivesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 incidents incident Id retrospectives created response has a 3xx status code
func (o *PostV1IncidentsIncidentIDRetrospectivesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 incidents incident Id retrospectives created response has a 4xx status code
func (o *PostV1IncidentsIncidentIDRetrospectivesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 incidents incident Id retrospectives created response has a 5xx status code
func (o *PostV1IncidentsIncidentIDRetrospectivesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 incidents incident Id retrospectives created response a status code equal to that given
func (o *PostV1IncidentsIncidentIDRetrospectivesCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1IncidentsIncidentIDRetrospectivesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/incidents/{incident_id}/retrospectives][%d] postV1IncidentsIncidentIdRetrospectivesCreated  %+v", 201, o.Payload)
}

func (o *PostV1IncidentsIncidentIDRetrospectivesCreated) String() string {
	return fmt.Sprintf("[POST /v1/incidents/{incident_id}/retrospectives][%d] postV1IncidentsIncidentIdRetrospectivesCreated  %+v", 201, o.Payload)
}

func (o *PostV1IncidentsIncidentIDRetrospectivesCreated) GetPayload() *models.IncidentsRetrospectiveEntity {
	return o.Payload
}

func (o *PostV1IncidentsIncidentIDRetrospectivesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncidentsRetrospectiveEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
