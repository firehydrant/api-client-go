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

// DeleteV1IncidentsIncidentIDImpactTypeIDReader is a Reader for the DeleteV1IncidentsIncidentIDImpactTypeID structure.
type DeleteV1IncidentsIncidentIDImpactTypeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1IncidentsIncidentIDImpactTypeIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteV1IncidentsIncidentIDImpactTypeIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1IncidentsIncidentIDImpactTypeIDNoContent creates a DeleteV1IncidentsIncidentIDImpactTypeIDNoContent with default headers values
func NewDeleteV1IncidentsIncidentIDImpactTypeIDNoContent() *DeleteV1IncidentsIncidentIDImpactTypeIDNoContent {
	return &DeleteV1IncidentsIncidentIDImpactTypeIDNoContent{}
}

/*
DeleteV1IncidentsIncidentIDImpactTypeIDNoContent describes a response with status code 204, with default header values.

Remove impacted infrastructure from an incident
*/
type DeleteV1IncidentsIncidentIDImpactTypeIDNoContent struct {
}

// IsSuccess returns true when this delete v1 incidents incident Id impact type Id no content response has a 2xx status code
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 incidents incident Id impact type Id no content response has a 3xx status code
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 incidents incident Id impact type Id no content response has a 4xx status code
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 incidents incident Id impact type Id no content response has a 5xx status code
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 incidents incident Id impact type Id no content response a status code equal to that given
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteV1IncidentsIncidentIDImpactTypeIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/incidents/{incident_id}/impact/{type}/{id}][%d] deleteV1IncidentsIncidentIdImpactTypeIdNoContent ", 204)
}

func (o *DeleteV1IncidentsIncidentIDImpactTypeIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/incidents/{incident_id}/impact/{type}/{id}][%d] deleteV1IncidentsIncidentIdImpactTypeIdNoContent ", 204)
}

func (o *DeleteV1IncidentsIncidentIDImpactTypeIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteV1IncidentsIncidentIDImpactTypeIDBadRequest creates a DeleteV1IncidentsIncidentIDImpactTypeIDBadRequest with default headers values
func NewDeleteV1IncidentsIncidentIDImpactTypeIDBadRequest() *DeleteV1IncidentsIncidentIDImpactTypeIDBadRequest {
	return &DeleteV1IncidentsIncidentIDImpactTypeIDBadRequest{}
}

/*
DeleteV1IncidentsIncidentIDImpactTypeIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteV1IncidentsIncidentIDImpactTypeIDBadRequest struct {
	Payload *models.ErrorEntity
}

// IsSuccess returns true when this delete v1 incidents incident Id impact type Id bad request response has a 2xx status code
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete v1 incidents incident Id impact type Id bad request response has a 3xx status code
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 incidents incident Id impact type Id bad request response has a 4xx status code
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete v1 incidents incident Id impact type Id bad request response has a 5xx status code
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 incidents incident Id impact type Id bad request response a status code equal to that given
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteV1IncidentsIncidentIDImpactTypeIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/incidents/{incident_id}/impact/{type}/{id}][%d] deleteV1IncidentsIncidentIdImpactTypeIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteV1IncidentsIncidentIDImpactTypeIDBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/incidents/{incident_id}/impact/{type}/{id}][%d] deleteV1IncidentsIncidentIdImpactTypeIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteV1IncidentsIncidentIDImpactTypeIDBadRequest) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *DeleteV1IncidentsIncidentIDImpactTypeIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
