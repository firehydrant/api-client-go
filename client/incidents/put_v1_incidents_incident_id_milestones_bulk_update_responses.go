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

// PutV1IncidentsIncidentIDMilestonesBulkUpdateReader is a Reader for the PutV1IncidentsIncidentIDMilestonesBulkUpdate structure.
type PutV1IncidentsIncidentIDMilestonesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutV1IncidentsIncidentIDMilestonesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutV1IncidentsIncidentIDMilestonesBulkUpdateOK creates a PutV1IncidentsIncidentIDMilestonesBulkUpdateOK with default headers values
func NewPutV1IncidentsIncidentIDMilestonesBulkUpdateOK() *PutV1IncidentsIncidentIDMilestonesBulkUpdateOK {
	return &PutV1IncidentsIncidentIDMilestonesBulkUpdateOK{}
}

/*
	PutV1IncidentsIncidentIDMilestonesBulkUpdateOK describes a response with status code 200, with default header values.

	Update milestone times in bulk for a given incident. All milestone

times for an incident must occur in chronological order
corresponding to the configured order of milestones. If the result
of this request would cause any milestone(s) to appear out of place,
a 422 response will instead be returned. This includes milestones
not explicitly submitted or updated in this request.
*/
type PutV1IncidentsIncidentIDMilestonesBulkUpdateOK struct {
	Payload *models.IncidentsMilestoneEntityPaginated
}

// IsSuccess returns true when this put v1 incidents incident Id milestones bulk update o k response has a 2xx status code
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put v1 incidents incident Id milestones bulk update o k response has a 3xx status code
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 incidents incident Id milestones bulk update o k response has a 4xx status code
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put v1 incidents incident Id milestones bulk update o k response has a 5xx status code
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 incidents incident Id milestones bulk update o k response a status code equal to that given
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1/incidents/{incident_id}/milestones/bulk_update][%d] putV1IncidentsIncidentIdMilestonesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /v1/incidents/{incident_id}/milestones/bulk_update][%d] putV1IncidentsIncidentIdMilestonesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateOK) GetPayload() *models.IncidentsMilestoneEntityPaginated {
	return o.Payload
}

func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncidentsMilestoneEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
