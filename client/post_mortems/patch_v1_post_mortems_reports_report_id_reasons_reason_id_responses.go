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

// PatchV1PostMortemsReportsReportIDReasonsReasonIDReader is a Reader for the PatchV1PostMortemsReportsReportIDReasonsReasonID structure.
type PatchV1PostMortemsReportsReportIDReasonsReasonIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1PostMortemsReportsReportIDReasonsReasonIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/post_mortems/reports/{report_id}/reasons/{reason_id}] patchV1PostMortemsReportsReportIdReasonsReasonId", response, response.Code())
	}
}

// NewPatchV1PostMortemsReportsReportIDReasonsReasonIDOK creates a PatchV1PostMortemsReportsReportIDReasonsReasonIDOK with default headers values
func NewPatchV1PostMortemsReportsReportIDReasonsReasonIDOK() *PatchV1PostMortemsReportsReportIDReasonsReasonIDOK {
	return &PatchV1PostMortemsReportsReportIDReasonsReasonIDOK{}
}

/*
PatchV1PostMortemsReportsReportIDReasonsReasonIDOK describes a response with status code 200, with default header values.

Update a contributing factor
*/
type PatchV1PostMortemsReportsReportIDReasonsReasonIDOK struct {
	Payload *models.PostMortemsReasonEntity
}

// IsSuccess returns true when this patch v1 post mortems reports report Id reasons reason Id o k response has a 2xx status code
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 post mortems reports report Id reasons reason Id o k response has a 3xx status code
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 post mortems reports report Id reasons reason Id o k response has a 4xx status code
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 post mortems reports report Id reasons reason Id o k response has a 5xx status code
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 post mortems reports report Id reasons reason Id o k response a status code equal to that given
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch v1 post mortems reports report Id reasons reason Id o k response
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDOK) Code() int {
	return 200
}

func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/post_mortems/reports/{report_id}/reasons/{reason_id}][%d] patchV1PostMortemsReportsReportIdReasonsReasonIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/post_mortems/reports/{report_id}/reasons/{reason_id}][%d] patchV1PostMortemsReportsReportIdReasonsReasonIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDOK) GetPayload() *models.PostMortemsReasonEntity {
	return o.Payload
}

func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMortemsReasonEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
