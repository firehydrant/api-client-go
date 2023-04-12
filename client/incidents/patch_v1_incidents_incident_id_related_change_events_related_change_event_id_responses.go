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

// PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDReader is a Reader for the PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID structure.
type PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK creates a PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK with default headers values
func NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK() *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK {
	return &PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK{}
}

/*
PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK describes a response with status code 200, with default header values.

Update a change attached to an incident
*/
type PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK struct {
	Payload *models.IncidentsRelatedChangeEventEntity
}

// IsSuccess returns true when this patch v1 incidents incident Id related change events related change event Id o k response has a 2xx status code
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 incidents incident Id related change events related change event Id o k response has a 3xx status code
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 incidents incident Id related change events related change event Id o k response has a 4xx status code
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 incidents incident Id related change events related change event Id o k response has a 5xx status code
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 incidents incident Id related change events related change event Id o k response a status code equal to that given
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/incidents/{incident_id}/related_change_events/{related_change_event_id}][%d] patchV1IncidentsIncidentIdRelatedChangeEventsRelatedChangeEventIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/incidents/{incident_id}/related_change_events/{related_change_event_id}][%d] patchV1IncidentsIncidentIdRelatedChangeEventsRelatedChangeEventIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK) GetPayload() *models.IncidentsRelatedChangeEventEntity {
	return o.Payload
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncidentsRelatedChangeEventEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest creates a PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest with default headers values
func NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest() *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest {
	return &PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest{}
}

/*
PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest struct {
	Payload *models.ErrorEntity
}

// IsSuccess returns true when this patch v1 incidents incident Id related change events related change event Id bad request response has a 2xx status code
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch v1 incidents incident Id related change events related change event Id bad request response has a 3xx status code
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 incidents incident Id related change events related change event Id bad request response has a 4xx status code
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch v1 incidents incident Id related change events related change event Id bad request response has a 5xx status code
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 incidents incident Id related change events related change event Id bad request response a status code equal to that given
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/incidents/{incident_id}/related_change_events/{related_change_event_id}][%d] patchV1IncidentsIncidentIdRelatedChangeEventsRelatedChangeEventIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest) String() string {
	return fmt.Sprintf("[PATCH /v1/incidents/{incident_id}/related_change_events/{related_change_event_id}][%d] patchV1IncidentsIncidentIdRelatedChangeEventsRelatedChangeEventIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict creates a PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict with default headers values
func NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict() *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict {
	return &PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict{}
}

/*
PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict describes a response with status code 409, with default header values.

Already Added
*/
type PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict struct {
	Payload *models.ErrorEntity
}

// IsSuccess returns true when this patch v1 incidents incident Id related change events related change event Id conflict response has a 2xx status code
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch v1 incidents incident Id related change events related change event Id conflict response has a 3xx status code
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 incidents incident Id related change events related change event Id conflict response has a 4xx status code
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch v1 incidents incident Id related change events related change event Id conflict response has a 5xx status code
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 incidents incident Id related change events related change event Id conflict response a status code equal to that given
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict) Error() string {
	return fmt.Sprintf("[PATCH /v1/incidents/{incident_id}/related_change_events/{related_change_event_id}][%d] patchV1IncidentsIncidentIdRelatedChangeEventsRelatedChangeEventIdConflict  %+v", 409, o.Payload)
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict) String() string {
	return fmt.Sprintf("[PATCH /v1/incidents/{incident_id}/related_change_events/{related_change_event_id}][%d] patchV1IncidentsIncidentIdRelatedChangeEventsRelatedChangeEventIdConflict  %+v", 409, o.Payload)
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
