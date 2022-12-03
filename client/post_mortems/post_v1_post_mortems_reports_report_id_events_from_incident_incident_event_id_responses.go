// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDReader is a Reader for the PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventID structure.
type PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated creates a PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated with default headers values
func NewPostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated() *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated {
	return &PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated{}
}

/*
PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated describes a response with status code 201, with default header values.

Creates an event on a report from an incident event
*/
type PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated struct {
	Payload *models.EventEntity
}

// IsSuccess returns true when this post v1 post mortems reports report Id events from incident incident event Id created response has a 2xx status code
func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 post mortems reports report Id events from incident incident event Id created response has a 3xx status code
func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 post mortems reports report Id events from incident incident event Id created response has a 4xx status code
func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 post mortems reports report Id events from incident incident event Id created response has a 5xx status code
func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 post mortems reports report Id events from incident incident event Id created response a status code equal to that given
func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated) Error() string {
	return fmt.Sprintf("[POST /v1/post_mortems/reports/{report_id}/events/from_incident/{incident_event_id}][%d] postV1PostMortemsReportsReportIdEventsFromIncidentIncidentEventIdCreated  %+v", 201, o.Payload)
}

func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated) String() string {
	return fmt.Sprintf("[POST /v1/post_mortems/reports/{report_id}/events/from_incident/{incident_event_id}][%d] postV1PostMortemsReportsReportIdEventsFromIncidentIncidentEventIdCreated  %+v", 201, o.Payload)
}

func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated) GetPayload() *models.EventEntity {
	return o.Payload
}

func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest creates a PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest with default headers values
func NewPostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest() *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest {
	return &PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest{}
}

/*
PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest struct {
	Payload *models.ErrorEntity
}

// IsSuccess returns true when this post v1 post mortems reports report Id events from incident incident event Id bad request response has a 2xx status code
func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 post mortems reports report Id events from incident incident event Id bad request response has a 3xx status code
func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 post mortems reports report Id events from incident incident event Id bad request response has a 4xx status code
func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 post mortems reports report Id events from incident incident event Id bad request response has a 5xx status code
func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 post mortems reports report Id events from incident incident event Id bad request response a status code equal to that given
func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/post_mortems/reports/{report_id}/events/from_incident/{incident_event_id}][%d] postV1PostMortemsReportsReportIdEventsFromIncidentIncidentEventIdBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/post_mortems/reports/{report_id}/events/from_incident/{incident_event_id}][%d] postV1PostMortemsReportsReportIdEventsFromIncidentIncidentEventIdBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *PostV1PostMortemsReportsReportIDEventsFromIncidentIncidentEventIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
