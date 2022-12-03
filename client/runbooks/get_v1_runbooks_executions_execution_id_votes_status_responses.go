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

// GetV1RunbooksExecutionsExecutionIDVotesStatusReader is a Reader for the GetV1RunbooksExecutionsExecutionIDVotesStatus structure.
type GetV1RunbooksExecutionsExecutionIDVotesStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1RunbooksExecutionsExecutionIDVotesStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1RunbooksExecutionsExecutionIDVotesStatusOK creates a GetV1RunbooksExecutionsExecutionIDVotesStatusOK with default headers values
func NewGetV1RunbooksExecutionsExecutionIDVotesStatusOK() *GetV1RunbooksExecutionsExecutionIDVotesStatusOK {
	return &GetV1RunbooksExecutionsExecutionIDVotesStatusOK{}
}

/*
GetV1RunbooksExecutionsExecutionIDVotesStatusOK describes a response with status code 200, with default header values.

Returns the current vote counts for an object
*/
type GetV1RunbooksExecutionsExecutionIDVotesStatusOK struct {
	Payload *models.VotesEntity
}

// IsSuccess returns true when this get v1 runbooks executions execution Id votes status o k response has a 2xx status code
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 runbooks executions execution Id votes status o k response has a 3xx status code
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 runbooks executions execution Id votes status o k response has a 4xx status code
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 runbooks executions execution Id votes status o k response has a 5xx status code
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 runbooks executions execution Id votes status o k response a status code equal to that given
func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/runbooks/executions/{execution_id}/votes/status][%d] getV1RunbooksExecutionsExecutionIdVotesStatusOK  %+v", 200, o.Payload)
}

func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusOK) String() string {
	return fmt.Sprintf("[GET /v1/runbooks/executions/{execution_id}/votes/status][%d] getV1RunbooksExecutionsExecutionIdVotesStatusOK  %+v", 200, o.Payload)
}

func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusOK) GetPayload() *models.VotesEntity {
	return o.Payload
}

func (o *GetV1RunbooksExecutionsExecutionIDVotesStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VotesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
