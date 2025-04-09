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

// PostV1PostMortemsReportsReportIDReasonsReader is a Reader for the PostV1PostMortemsReportsReportIDReasons structure.
type PostV1PostMortemsReportsReportIDReasonsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1PostMortemsReportsReportIDReasonsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1PostMortemsReportsReportIDReasonsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1PostMortemsReportsReportIDReasonsCreated creates a PostV1PostMortemsReportsReportIDReasonsCreated with default headers values
func NewPostV1PostMortemsReportsReportIDReasonsCreated() *PostV1PostMortemsReportsReportIDReasonsCreated {
	return &PostV1PostMortemsReportsReportIDReasonsCreated{}
}

/*
PostV1PostMortemsReportsReportIDReasonsCreated describes a response with status code 201, with default header values.

Add a new contributing factor to an incident
*/
type PostV1PostMortemsReportsReportIDReasonsCreated struct {
	Payload *models.PostMortemsReasonEntity
}

// IsSuccess returns true when this post v1 post mortems reports report Id reasons created response has a 2xx status code
func (o *PostV1PostMortemsReportsReportIDReasonsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 post mortems reports report Id reasons created response has a 3xx status code
func (o *PostV1PostMortemsReportsReportIDReasonsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 post mortems reports report Id reasons created response has a 4xx status code
func (o *PostV1PostMortemsReportsReportIDReasonsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 post mortems reports report Id reasons created response has a 5xx status code
func (o *PostV1PostMortemsReportsReportIDReasonsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 post mortems reports report Id reasons created response a status code equal to that given
func (o *PostV1PostMortemsReportsReportIDReasonsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1PostMortemsReportsReportIDReasonsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/post_mortems/reports/{report_id}/reasons][%d] postV1PostMortemsReportsReportIdReasonsCreated  %+v", 201, o.Payload)
}

func (o *PostV1PostMortemsReportsReportIDReasonsCreated) String() string {
	return fmt.Sprintf("[POST /v1/post_mortems/reports/{report_id}/reasons][%d] postV1PostMortemsReportsReportIdReasonsCreated  %+v", 201, o.Payload)
}

func (o *PostV1PostMortemsReportsReportIDReasonsCreated) GetPayload() *models.PostMortemsReasonEntity {
	return o.Payload
}

func (o *PostV1PostMortemsReportsReportIDReasonsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMortemsReasonEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
