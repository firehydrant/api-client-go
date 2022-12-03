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

// GetV1PostMortemsReportsReportIDReader is a Reader for the GetV1PostMortemsReportsReportID structure.
type GetV1PostMortemsReportsReportIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1PostMortemsReportsReportIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1PostMortemsReportsReportIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1PostMortemsReportsReportIDOK creates a GetV1PostMortemsReportsReportIDOK with default headers values
func NewGetV1PostMortemsReportsReportIDOK() *GetV1PostMortemsReportsReportIDOK {
	return &GetV1PostMortemsReportsReportIDOK{}
}

/*
GetV1PostMortemsReportsReportIDOK describes a response with status code 200, with default header values.

Get a report
*/
type GetV1PostMortemsReportsReportIDOK struct {
	Payload *models.PostMortemReportEntity
}

// IsSuccess returns true when this get v1 post mortems reports report Id o k response has a 2xx status code
func (o *GetV1PostMortemsReportsReportIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 post mortems reports report Id o k response has a 3xx status code
func (o *GetV1PostMortemsReportsReportIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 post mortems reports report Id o k response has a 4xx status code
func (o *GetV1PostMortemsReportsReportIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 post mortems reports report Id o k response has a 5xx status code
func (o *GetV1PostMortemsReportsReportIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 post mortems reports report Id o k response a status code equal to that given
func (o *GetV1PostMortemsReportsReportIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1PostMortemsReportsReportIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/post_mortems/reports/{report_id}][%d] getV1PostMortemsReportsReportIdOK  %+v", 200, o.Payload)
}

func (o *GetV1PostMortemsReportsReportIDOK) String() string {
	return fmt.Sprintf("[GET /v1/post_mortems/reports/{report_id}][%d] getV1PostMortemsReportsReportIdOK  %+v", 200, o.Payload)
}

func (o *GetV1PostMortemsReportsReportIDOK) GetPayload() *models.PostMortemReportEntity {
	return o.Payload
}

func (o *GetV1PostMortemsReportsReportIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMortemReportEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
