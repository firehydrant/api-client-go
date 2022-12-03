// Code generated by go-swagger; DO NOT EDIT.

package runbooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesReader is a Reader for the PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotes structure.
type PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK creates a PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK with default headers values
func NewPatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK() *PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK {
	return &PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK{}
}

/*
PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK describes a response with status code 200, with default header values.

Allows for upvoting or downvoting an event
*/
type PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK struct {
	Payload *models.VotesEntity
}

// IsSuccess returns true when this patch v1 runbooks executions execution Id steps step Id votes o k response has a 2xx status code
func (o *PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 runbooks executions execution Id steps step Id votes o k response has a 3xx status code
func (o *PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 runbooks executions execution Id steps step Id votes o k response has a 4xx status code
func (o *PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 runbooks executions execution Id steps step Id votes o k response has a 5xx status code
func (o *PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 runbooks executions execution Id steps step Id votes o k response a status code equal to that given
func (o *PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/runbooks/executions/{execution_id}/steps/{step_id}/votes][%d] patchV1RunbooksExecutionsExecutionIdStepsStepIdVotesOK  %+v", 200, o.Payload)
}

func (o *PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK) String() string {
	return fmt.Sprintf("[PATCH /v1/runbooks/executions/{execution_id}/steps/{step_id}/votes][%d] patchV1RunbooksExecutionsExecutionIdStepsStepIdVotesOK  %+v", 200, o.Payload)
}

func (o *PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK) GetPayload() *models.VotesEntity {
	return o.Payload
}

func (o *PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VotesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
