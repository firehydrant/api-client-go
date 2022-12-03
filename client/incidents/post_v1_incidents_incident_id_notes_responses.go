// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1IncidentsIncidentIDNotesReader is a Reader for the PostV1IncidentsIncidentIDNotes structure.
type PostV1IncidentsIncidentIDNotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1IncidentsIncidentIDNotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1IncidentsIncidentIDNotesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1IncidentsIncidentIDNotesCreated creates a PostV1IncidentsIncidentIDNotesCreated with default headers values
func NewPostV1IncidentsIncidentIDNotesCreated() *PostV1IncidentsIncidentIDNotesCreated {
	return &PostV1IncidentsIncidentIDNotesCreated{}
}

/*
PostV1IncidentsIncidentIDNotesCreated describes a response with status code 201, with default header values.

Create a new note on for an incident. The visibility field on a note determines where it gets posted.
*/
type PostV1IncidentsIncidentIDNotesCreated struct {
	Payload *models.NoteEntity
}

// IsSuccess returns true when this post v1 incidents incident Id notes created response has a 2xx status code
func (o *PostV1IncidentsIncidentIDNotesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 incidents incident Id notes created response has a 3xx status code
func (o *PostV1IncidentsIncidentIDNotesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 incidents incident Id notes created response has a 4xx status code
func (o *PostV1IncidentsIncidentIDNotesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 incidents incident Id notes created response has a 5xx status code
func (o *PostV1IncidentsIncidentIDNotesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 incidents incident Id notes created response a status code equal to that given
func (o *PostV1IncidentsIncidentIDNotesCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1IncidentsIncidentIDNotesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/incidents/{incident_id}/notes][%d] postV1IncidentsIncidentIdNotesCreated  %+v", 201, o.Payload)
}

func (o *PostV1IncidentsIncidentIDNotesCreated) String() string {
	return fmt.Sprintf("[POST /v1/incidents/{incident_id}/notes][%d] postV1IncidentsIncidentIdNotesCreated  %+v", 201, o.Payload)
}

func (o *PostV1IncidentsIncidentIDNotesCreated) GetPayload() *models.NoteEntity {
	return o.Payload
}

func (o *PostV1IncidentsIncidentIDNotesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NoteEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
