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

// GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDReader is a Reader for the GetV1TicketingProjectsTicketingProjectIDFieldMapsMapID structure.
type GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK creates a GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK with default headers values
func NewGetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK() *GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK {
	return &GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK{}
}

/*
GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK describes a response with status code 200, with default header values.

Retrieve field map for a ticketing project
*/
type GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK struct {
	Payload *models.TicketingProjectFieldMapEntity
}

// IsSuccess returns true when this get v1 ticketing projects ticketing project Id field maps map Id o k response has a 2xx status code
func (o *GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 ticketing projects ticketing project Id field maps map Id o k response has a 3xx status code
func (o *GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 ticketing projects ticketing project Id field maps map Id o k response has a 4xx status code
func (o *GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 ticketing projects ticketing project Id field maps map Id o k response has a 5xx status code
func (o *GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 ticketing projects ticketing project Id field maps map Id o k response a status code equal to that given
func (o *GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/ticketing/projects/{ticketing_project_id}/field_maps/{map_id}][%d] getV1TicketingProjectsTicketingProjectIdFieldMapsMapIdOK  %+v", 200, o.Payload)
}

func (o *GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK) String() string {
	return fmt.Sprintf("[GET /v1/ticketing/projects/{ticketing_project_id}/field_maps/{map_id}][%d] getV1TicketingProjectsTicketingProjectIdFieldMapsMapIdOK  %+v", 200, o.Payload)
}

func (o *GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK) GetPayload() *models.TicketingProjectFieldMapEntity {
	return o.Payload
}

func (o *GetV1TicketingProjectsTicketingProjectIDFieldMapsMapIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TicketingProjectFieldMapEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
