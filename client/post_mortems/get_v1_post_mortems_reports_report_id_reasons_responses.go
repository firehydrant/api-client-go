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

// GetV1PostMortemsReportsReportIDReasonsReader is a Reader for the GetV1PostMortemsReportsReportIDReasons structure.
type GetV1PostMortemsReportsReportIDReasonsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1PostMortemsReportsReportIDReasonsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1PostMortemsReportsReportIDReasonsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/post_mortems/reports/{report_id}/reasons] getV1PostMortemsReportsReportIdReasons", response, response.Code())
	}
}

// NewGetV1PostMortemsReportsReportIDReasonsOK creates a GetV1PostMortemsReportsReportIDReasonsOK with default headers values
func NewGetV1PostMortemsReportsReportIDReasonsOK() *GetV1PostMortemsReportsReportIDReasonsOK {
	return &GetV1PostMortemsReportsReportIDReasonsOK{}
}

/*
GetV1PostMortemsReportsReportIDReasonsOK describes a response with status code 200, with default header values.

List all contributing factors to an incident
*/
type GetV1PostMortemsReportsReportIDReasonsOK struct {
	Payload *models.PostMortemsReasonEntityPaginated
}

// IsSuccess returns true when this get v1 post mortems reports report Id reasons o k response has a 2xx status code
func (o *GetV1PostMortemsReportsReportIDReasonsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 post mortems reports report Id reasons o k response has a 3xx status code
func (o *GetV1PostMortemsReportsReportIDReasonsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 post mortems reports report Id reasons o k response has a 4xx status code
func (o *GetV1PostMortemsReportsReportIDReasonsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 post mortems reports report Id reasons o k response has a 5xx status code
func (o *GetV1PostMortemsReportsReportIDReasonsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 post mortems reports report Id reasons o k response a status code equal to that given
func (o *GetV1PostMortemsReportsReportIDReasonsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 post mortems reports report Id reasons o k response
func (o *GetV1PostMortemsReportsReportIDReasonsOK) Code() int {
	return 200
}

func (o *GetV1PostMortemsReportsReportIDReasonsOK) Error() string {
	return fmt.Sprintf("[GET /v1/post_mortems/reports/{report_id}/reasons][%d] getV1PostMortemsReportsReportIdReasonsOK  %+v", 200, o.Payload)
}

func (o *GetV1PostMortemsReportsReportIDReasonsOK) String() string {
	return fmt.Sprintf("[GET /v1/post_mortems/reports/{report_id}/reasons][%d] getV1PostMortemsReportsReportIdReasonsOK  %+v", 200, o.Payload)
}

func (o *GetV1PostMortemsReportsReportIDReasonsOK) GetPayload() *models.PostMortemsReasonEntityPaginated {
	return o.Payload
}

func (o *GetV1PostMortemsReportsReportIDReasonsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMortemsReasonEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
