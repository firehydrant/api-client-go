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

// PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsReader is a Reader for the PostV1TicketingProjectsTicketingProjectIDInboundFieldMaps structure.
type PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated creates a PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated with default headers values
func NewPostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated() *PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated {
	return &PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated{}
}

/*
PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated describes a response with status code 201, with default header values.

Creates inbound field map for a ticketing project
*/
type PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated struct {
	Payload *models.TicketingProjectInboundFieldMapEntity
}

// IsSuccess returns true when this post v1 ticketing projects ticketing project Id inbound field maps created response has a 2xx status code
func (o *PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 ticketing projects ticketing project Id inbound field maps created response has a 3xx status code
func (o *PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 ticketing projects ticketing project Id inbound field maps created response has a 4xx status code
func (o *PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 ticketing projects ticketing project Id inbound field maps created response has a 5xx status code
func (o *PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 ticketing projects ticketing project Id inbound field maps created response a status code equal to that given
func (o *PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/ticketing/projects/{ticketing_project_id}/inbound_field_maps][%d] postV1TicketingProjectsTicketingProjectIdInboundFieldMapsCreated  %+v", 201, o.Payload)
}

func (o *PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated) String() string {
	return fmt.Sprintf("[POST /v1/ticketing/projects/{ticketing_project_id}/inbound_field_maps][%d] postV1TicketingProjectsTicketingProjectIdInboundFieldMapsCreated  %+v", 201, o.Payload)
}

func (o *PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated) GetPayload() *models.TicketingProjectInboundFieldMapEntity {
	return o.Payload
}

func (o *PostV1TicketingProjectsTicketingProjectIDInboundFieldMapsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TicketingProjectInboundFieldMapEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
