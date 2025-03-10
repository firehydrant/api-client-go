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

// PostV1AudiencesAudienceIDSummariesIncidentIDReader is a Reader for the PostV1AudiencesAudienceIDSummariesIncidentID structure.
type PostV1AudiencesAudienceIDSummariesIncidentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1AudiencesAudienceIDSummariesIncidentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1AudiencesAudienceIDSummariesIncidentIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1AudiencesAudienceIDSummariesIncidentIDCreated creates a PostV1AudiencesAudienceIDSummariesIncidentIDCreated with default headers values
func NewPostV1AudiencesAudienceIDSummariesIncidentIDCreated() *PostV1AudiencesAudienceIDSummariesIncidentIDCreated {
	return &PostV1AudiencesAudienceIDSummariesIncidentIDCreated{}
}

/*
PostV1AudiencesAudienceIDSummariesIncidentIDCreated describes a response with status code 201, with default header values.

Generate a new audience-specific summary for an incident
*/
type PostV1AudiencesAudienceIDSummariesIncidentIDCreated struct {
	Payload *models.AIEntitiesIncidentSummaryEntity
}

// IsSuccess returns true when this post v1 audiences audience Id summaries incident Id created response has a 2xx status code
func (o *PostV1AudiencesAudienceIDSummariesIncidentIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 audiences audience Id summaries incident Id created response has a 3xx status code
func (o *PostV1AudiencesAudienceIDSummariesIncidentIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 audiences audience Id summaries incident Id created response has a 4xx status code
func (o *PostV1AudiencesAudienceIDSummariesIncidentIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 audiences audience Id summaries incident Id created response has a 5xx status code
func (o *PostV1AudiencesAudienceIDSummariesIncidentIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 audiences audience Id summaries incident Id created response a status code equal to that given
func (o *PostV1AudiencesAudienceIDSummariesIncidentIDCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1AudiencesAudienceIDSummariesIncidentIDCreated) Error() string {
	return fmt.Sprintf("[POST /v1/audiences/{audience_id}/summaries/{incident_id}][%d] postV1AudiencesAudienceIdSummariesIncidentIdCreated  %+v", 201, o.Payload)
}

func (o *PostV1AudiencesAudienceIDSummariesIncidentIDCreated) String() string {
	return fmt.Sprintf("[POST /v1/audiences/{audience_id}/summaries/{incident_id}][%d] postV1AudiencesAudienceIdSummariesIncidentIdCreated  %+v", 201, o.Payload)
}

func (o *PostV1AudiencesAudienceIDSummariesIncidentIDCreated) GetPayload() *models.AIEntitiesIncidentSummaryEntity {
	return o.Payload
}

func (o *PostV1AudiencesAudienceIDSummariesIncidentIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AIEntitiesIncidentSummaryEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
