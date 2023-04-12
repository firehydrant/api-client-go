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

// PatchV1PostMortemsReportsReportIDReader is a Reader for the PatchV1PostMortemsReportsReportID structure.
type PatchV1PostMortemsReportsReportIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1PostMortemsReportsReportIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1PostMortemsReportsReportIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1PostMortemsReportsReportIDOK creates a PatchV1PostMortemsReportsReportIDOK with default headers values
func NewPatchV1PostMortemsReportsReportIDOK() *PatchV1PostMortemsReportsReportIDOK {
	return &PatchV1PostMortemsReportsReportIDOK{}
}

/*
PatchV1PostMortemsReportsReportIDOK describes a response with status code 200, with default header values.

Update a report
*/
type PatchV1PostMortemsReportsReportIDOK struct {
	Payload *models.PostMortemsPostMortemReportEntity
}

// IsSuccess returns true when this patch v1 post mortems reports report Id o k response has a 2xx status code
func (o *PatchV1PostMortemsReportsReportIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 post mortems reports report Id o k response has a 3xx status code
func (o *PatchV1PostMortemsReportsReportIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 post mortems reports report Id o k response has a 4xx status code
func (o *PatchV1PostMortemsReportsReportIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 post mortems reports report Id o k response has a 5xx status code
func (o *PatchV1PostMortemsReportsReportIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 post mortems reports report Id o k response a status code equal to that given
func (o *PatchV1PostMortemsReportsReportIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchV1PostMortemsReportsReportIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/post_mortems/reports/{report_id}][%d] patchV1PostMortemsReportsReportIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1PostMortemsReportsReportIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/post_mortems/reports/{report_id}][%d] patchV1PostMortemsReportsReportIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1PostMortemsReportsReportIDOK) GetPayload() *models.PostMortemsPostMortemReportEntity {
	return o.Payload
}

func (o *PatchV1PostMortemsReportsReportIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMortemsPostMortemReportEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
