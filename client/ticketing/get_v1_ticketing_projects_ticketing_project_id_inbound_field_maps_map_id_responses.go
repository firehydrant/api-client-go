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

// GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDReader is a Reader for the GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapID structure.
type GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK creates a GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK with default headers values
func NewGetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK() *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK {
	return &GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK{}
}

/*
GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK describes a response with status code 200, with default header values.

Retrieve inbound field map for a ticketing project
*/
type GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK struct {
	Payload *models.TicketingProjectInboundFieldMapEntity
}

// IsSuccess returns true when this get v1 ticketing projects ticketing project Id inbound field maps map Id o k response has a 2xx status code
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 ticketing projects ticketing project Id inbound field maps map Id o k response has a 3xx status code
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 ticketing projects ticketing project Id inbound field maps map Id o k response has a 4xx status code
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 ticketing projects ticketing project Id inbound field maps map Id o k response has a 5xx status code
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 ticketing projects ticketing project Id inbound field maps map Id o k response a status code equal to that given
func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/ticketing/projects/{ticketing_project_id}/inbound_field_maps/{map_id}][%d] getV1TicketingProjectsTicketingProjectIdInboundFieldMapsMapIdOK  %+v", 200, o.Payload)
}

func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK) String() string {
	return fmt.Sprintf("[GET /v1/ticketing/projects/{ticketing_project_id}/inbound_field_maps/{map_id}][%d] getV1TicketingProjectsTicketingProjectIdInboundFieldMapsMapIdOK  %+v", 200, o.Payload)
}

func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK) GetPayload() *models.TicketingProjectInboundFieldMapEntity {
	return o.Payload
}

func (o *GetV1TicketingProjectsTicketingProjectIDInboundFieldMapsMapIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TicketingProjectInboundFieldMapEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
