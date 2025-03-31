// Code generated by go-swagger; DO NOT EDIT.

package audiences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1AudiencesSummariesIncidentIDReader is a Reader for the GetV1AudiencesSummariesIncidentID structure.
type GetV1AudiencesSummariesIncidentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1AudiencesSummariesIncidentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1AudiencesSummariesIncidentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1AudiencesSummariesIncidentIDOK creates a GetV1AudiencesSummariesIncidentIDOK with default headers values
func NewGetV1AudiencesSummariesIncidentIDOK() *GetV1AudiencesSummariesIncidentIDOK {
	return &GetV1AudiencesSummariesIncidentIDOK{}
}

/*
GetV1AudiencesSummariesIncidentIDOK describes a response with status code 200, with default header values.

List all audience summaries for an incident
*/
type GetV1AudiencesSummariesIncidentIDOK struct {
	Payload *models.AudiencesEntitiesAudienceSummariesEntity
}

// IsSuccess returns true when this get v1 audiences summaries incident Id o k response has a 2xx status code
func (o *GetV1AudiencesSummariesIncidentIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 audiences summaries incident Id o k response has a 3xx status code
func (o *GetV1AudiencesSummariesIncidentIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 audiences summaries incident Id o k response has a 4xx status code
func (o *GetV1AudiencesSummariesIncidentIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 audiences summaries incident Id o k response has a 5xx status code
func (o *GetV1AudiencesSummariesIncidentIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 audiences summaries incident Id o k response a status code equal to that given
func (o *GetV1AudiencesSummariesIncidentIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1AudiencesSummariesIncidentIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/audiences/summaries/{incident_id}][%d] getV1AudiencesSummariesIncidentIdOK  %+v", 200, o.Payload)
}

func (o *GetV1AudiencesSummariesIncidentIDOK) String() string {
	return fmt.Sprintf("[GET /v1/audiences/summaries/{incident_id}][%d] getV1AudiencesSummariesIncidentIdOK  %+v", 200, o.Payload)
}

func (o *GetV1AudiencesSummariesIncidentIDOK) GetPayload() *models.AudiencesEntitiesAudienceSummariesEntity {
	return o.Payload
}

func (o *GetV1AudiencesSummariesIncidentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AudiencesEntitiesAudienceSummariesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
