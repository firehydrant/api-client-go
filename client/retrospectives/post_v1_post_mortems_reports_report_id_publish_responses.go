// Code generated by go-swagger; DO NOT EDIT.

package retrospectives

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1PostMortemsReportsReportIDPublishReader is a Reader for the PostV1PostMortemsReportsReportIDPublish structure.
type PostV1PostMortemsReportsReportIDPublishReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1PostMortemsReportsReportIDPublishReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1PostMortemsReportsReportIDPublishCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1PostMortemsReportsReportIDPublishBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1PostMortemsReportsReportIDPublishCreated creates a PostV1PostMortemsReportsReportIDPublishCreated with default headers values
func NewPostV1PostMortemsReportsReportIDPublishCreated() *PostV1PostMortemsReportsReportIDPublishCreated {
	return &PostV1PostMortemsReportsReportIDPublishCreated{}
}

/*
PostV1PostMortemsReportsReportIDPublishCreated describes a response with status code 201, with default header values.

Marks an incident retrospective as published and emails all of the participants in the report the summary
*/
type PostV1PostMortemsReportsReportIDPublishCreated struct {
	Payload *models.PostMortemsPostMortemReportEntity
}

// IsSuccess returns true when this post v1 post mortems reports report Id publish created response has a 2xx status code
func (o *PostV1PostMortemsReportsReportIDPublishCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 post mortems reports report Id publish created response has a 3xx status code
func (o *PostV1PostMortemsReportsReportIDPublishCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 post mortems reports report Id publish created response has a 4xx status code
func (o *PostV1PostMortemsReportsReportIDPublishCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 post mortems reports report Id publish created response has a 5xx status code
func (o *PostV1PostMortemsReportsReportIDPublishCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 post mortems reports report Id publish created response a status code equal to that given
func (o *PostV1PostMortemsReportsReportIDPublishCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1PostMortemsReportsReportIDPublishCreated) Error() string {
	return fmt.Sprintf("[POST /v1/post_mortems/reports/{report_id}/publish][%d] postV1PostMortemsReportsReportIdPublishCreated  %+v", 201, o.Payload)
}

func (o *PostV1PostMortemsReportsReportIDPublishCreated) String() string {
	return fmt.Sprintf("[POST /v1/post_mortems/reports/{report_id}/publish][%d] postV1PostMortemsReportsReportIdPublishCreated  %+v", 201, o.Payload)
}

func (o *PostV1PostMortemsReportsReportIDPublishCreated) GetPayload() *models.PostMortemsPostMortemReportEntity {
	return o.Payload
}

func (o *PostV1PostMortemsReportsReportIDPublishCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMortemsPostMortemReportEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1PostMortemsReportsReportIDPublishBadRequest creates a PostV1PostMortemsReportsReportIDPublishBadRequest with default headers values
func NewPostV1PostMortemsReportsReportIDPublishBadRequest() *PostV1PostMortemsReportsReportIDPublishBadRequest {
	return &PostV1PostMortemsReportsReportIDPublishBadRequest{}
}

/*
PostV1PostMortemsReportsReportIDPublishBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostV1PostMortemsReportsReportIDPublishBadRequest struct {
	Payload *models.ErrorEntity
}

// IsSuccess returns true when this post v1 post mortems reports report Id publish bad request response has a 2xx status code
func (o *PostV1PostMortemsReportsReportIDPublishBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 post mortems reports report Id publish bad request response has a 3xx status code
func (o *PostV1PostMortemsReportsReportIDPublishBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 post mortems reports report Id publish bad request response has a 4xx status code
func (o *PostV1PostMortemsReportsReportIDPublishBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 post mortems reports report Id publish bad request response has a 5xx status code
func (o *PostV1PostMortemsReportsReportIDPublishBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 post mortems reports report Id publish bad request response a status code equal to that given
func (o *PostV1PostMortemsReportsReportIDPublishBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostV1PostMortemsReportsReportIDPublishBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/post_mortems/reports/{report_id}/publish][%d] postV1PostMortemsReportsReportIdPublishBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1PostMortemsReportsReportIDPublishBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/post_mortems/reports/{report_id}/publish][%d] postV1PostMortemsReportsReportIdPublishBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1PostMortemsReportsReportIDPublishBadRequest) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *PostV1PostMortemsReportsReportIDPublishBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
