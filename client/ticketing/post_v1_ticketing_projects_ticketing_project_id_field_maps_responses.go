// Code generated by go-swagger; DO NOT EDIT.

package ticketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1TicketingProjectsTicketingProjectIDFieldMapsReader is a Reader for the PostV1TicketingProjectsTicketingProjectIDFieldMaps structure.
type PostV1TicketingProjectsTicketingProjectIDFieldMapsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1TicketingProjectsTicketingProjectIDFieldMapsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1TicketingProjectsTicketingProjectIDFieldMapsCreated creates a PostV1TicketingProjectsTicketingProjectIDFieldMapsCreated with default headers values
func NewPostV1TicketingProjectsTicketingProjectIDFieldMapsCreated() *PostV1TicketingProjectsTicketingProjectIDFieldMapsCreated {
	return &PostV1TicketingProjectsTicketingProjectIDFieldMapsCreated{}
}

/*
PostV1TicketingProjectsTicketingProjectIDFieldMapsCreated describes a response with status code 201, with default header values.

Creates field map for a ticketing project
*/
type PostV1TicketingProjectsTicketingProjectIDFieldMapsCreated struct {
	Payload *models.TicketingProjectFieldMapEntity
}

// IsSuccess returns true when this post v1 ticketing projects ticketing project Id field maps created response has a 2xx status code
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 ticketing projects ticketing project Id field maps created response has a 3xx status code
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 ticketing projects ticketing project Id field maps created response has a 4xx status code
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 ticketing projects ticketing project Id field maps created response has a 5xx status code
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 ticketing projects ticketing project Id field maps created response a status code equal to that given
func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/ticketing/projects/{ticketing_project_id}/field_maps][%d] postV1TicketingProjectsTicketingProjectIdFieldMapsCreated  %+v", 201, o.Payload)
}

func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsCreated) String() string {
	return fmt.Sprintf("[POST /v1/ticketing/projects/{ticketing_project_id}/field_maps][%d] postV1TicketingProjectsTicketingProjectIdFieldMapsCreated  %+v", 201, o.Payload)
}

func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsCreated) GetPayload() *models.TicketingProjectFieldMapEntity {
	return o.Payload
}

func (o *PostV1TicketingProjectsTicketingProjectIDFieldMapsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TicketingProjectFieldMapEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
