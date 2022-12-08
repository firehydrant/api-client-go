// Code generated by go-swagger; DO NOT EDIT.

package ticketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDReader is a Reader for the DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapID structure.
type DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent creates a DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent with default headers values
func NewDeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent() *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent {
	return &DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent{}
}

/*
DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent describes a response with status code 204, with default header values.

Archive field map for a ticketing project
*/
type DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent struct {
}

// IsSuccess returns true when this delete v1 ticketing projects ticketing project Id field maps map Id no content response has a 2xx status code
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 ticketing projects ticketing project Id field maps map Id no content response has a 3xx status code
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 ticketing projects ticketing project Id field maps map Id no content response has a 4xx status code
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 ticketing projects ticketing project Id field maps map Id no content response has a 5xx status code
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 ticketing projects ticketing project Id field maps map Id no content response a status code equal to that given
func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/ticketing/projects/{ticketing_project_id}/field_maps/{map_id}][%d] deleteV1TicketingProjectsTicketingProjectIdFieldMapsMapIdNoContent ", 204)
}

func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/ticketing/projects/{ticketing_project_id}/field_maps/{map_id}][%d] deleteV1TicketingProjectsTicketingProjectIdFieldMapsMapIdNoContent ", 204)
}

func (o *DeleteV1TicketingProjectsTicketingProjectIDFieldMapsMapIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}