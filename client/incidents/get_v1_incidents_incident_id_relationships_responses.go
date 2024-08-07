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

// GetV1IncidentsIncidentIDRelationshipsReader is a Reader for the GetV1IncidentsIncidentIDRelationships structure.
type GetV1IncidentsIncidentIDRelationshipsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IncidentsIncidentIDRelationshipsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IncidentsIncidentIDRelationshipsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1IncidentsIncidentIDRelationshipsOK creates a GetV1IncidentsIncidentIDRelationshipsOK with default headers values
func NewGetV1IncidentsIncidentIDRelationshipsOK() *GetV1IncidentsIncidentIDRelationshipsOK {
	return &GetV1IncidentsIncidentIDRelationshipsOK{}
}

/*
GetV1IncidentsIncidentIDRelationshipsOK describes a response with status code 200, with default header values.

List any parent/child relationships for an incident
*/
type GetV1IncidentsIncidentIDRelationshipsOK struct {
	Payload *models.IncidentsRelationshipsEntity
}

// IsSuccess returns true when this get v1 incidents incident Id relationships o k response has a 2xx status code
func (o *GetV1IncidentsIncidentIDRelationshipsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 incidents incident Id relationships o k response has a 3xx status code
func (o *GetV1IncidentsIncidentIDRelationshipsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 incidents incident Id relationships o k response has a 4xx status code
func (o *GetV1IncidentsIncidentIDRelationshipsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 incidents incident Id relationships o k response has a 5xx status code
func (o *GetV1IncidentsIncidentIDRelationshipsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 incidents incident Id relationships o k response a status code equal to that given
func (o *GetV1IncidentsIncidentIDRelationshipsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1IncidentsIncidentIDRelationshipsOK) Error() string {
	return fmt.Sprintf("[GET /v1/incidents/{incident_id}/relationships][%d] getV1IncidentsIncidentIdRelationshipsOK  %+v", 200, o.Payload)
}

func (o *GetV1IncidentsIncidentIDRelationshipsOK) String() string {
	return fmt.Sprintf("[GET /v1/incidents/{incident_id}/relationships][%d] getV1IncidentsIncidentIdRelationshipsOK  %+v", 200, o.Payload)
}

func (o *GetV1IncidentsIncidentIDRelationshipsOK) GetPayload() *models.IncidentsRelationshipsEntity {
	return o.Payload
}

func (o *GetV1IncidentsIncidentIDRelationshipsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncidentsRelationshipsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
