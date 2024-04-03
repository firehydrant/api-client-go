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

// PatchV1IncidentsIncidentIDTasksTaskIDReader is a Reader for the PatchV1IncidentsIncidentIDTasksTaskID structure.
type PatchV1IncidentsIncidentIDTasksTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1IncidentsIncidentIDTasksTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1IncidentsIncidentIDTasksTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/incidents/{incident_id}/tasks/{task_id}] patchV1IncidentsIncidentIdTasksTaskId", response, response.Code())
	}
}

// NewPatchV1IncidentsIncidentIDTasksTaskIDOK creates a PatchV1IncidentsIncidentIDTasksTaskIDOK with default headers values
func NewPatchV1IncidentsIncidentIDTasksTaskIDOK() *PatchV1IncidentsIncidentIDTasksTaskIDOK {
	return &PatchV1IncidentsIncidentIDTasksTaskIDOK{}
}

/*
PatchV1IncidentsIncidentIDTasksTaskIDOK describes a response with status code 200, with default header values.

Update a task's attributes
*/
type PatchV1IncidentsIncidentIDTasksTaskIDOK struct {
	Payload *models.TaskEntity
}

// IsSuccess returns true when this patch v1 incidents incident Id tasks task Id o k response has a 2xx status code
func (o *PatchV1IncidentsIncidentIDTasksTaskIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 incidents incident Id tasks task Id o k response has a 3xx status code
func (o *PatchV1IncidentsIncidentIDTasksTaskIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 incidents incident Id tasks task Id o k response has a 4xx status code
func (o *PatchV1IncidentsIncidentIDTasksTaskIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 incidents incident Id tasks task Id o k response has a 5xx status code
func (o *PatchV1IncidentsIncidentIDTasksTaskIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 incidents incident Id tasks task Id o k response a status code equal to that given
func (o *PatchV1IncidentsIncidentIDTasksTaskIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch v1 incidents incident Id tasks task Id o k response
func (o *PatchV1IncidentsIncidentIDTasksTaskIDOK) Code() int {
	return 200
}

func (o *PatchV1IncidentsIncidentIDTasksTaskIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/incidents/{incident_id}/tasks/{task_id}][%d] patchV1IncidentsIncidentIdTasksTaskIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1IncidentsIncidentIDTasksTaskIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/incidents/{incident_id}/tasks/{task_id}][%d] patchV1IncidentsIncidentIdTasksTaskIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1IncidentsIncidentIDTasksTaskIDOK) GetPayload() *models.TaskEntity {
	return o.Payload
}

func (o *PatchV1IncidentsIncidentIDTasksTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
