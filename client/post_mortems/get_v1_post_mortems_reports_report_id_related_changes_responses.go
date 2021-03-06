// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firehydrant/api-client-go/models"
)

// GetV1PostMortemsReportsReportIDRelatedChangesReader is a Reader for the GetV1PostMortemsReportsReportIDRelatedChanges structure.
type GetV1PostMortemsReportsReportIDRelatedChangesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1PostMortemsReportsReportIDRelatedChangesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetV1PostMortemsReportsReportIDRelatedChangesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetV1PostMortemsReportsReportIDRelatedChangesOK creates a GetV1PostMortemsReportsReportIDRelatedChangesOK with default headers values
func NewGetV1PostMortemsReportsReportIDRelatedChangesOK() *GetV1PostMortemsReportsReportIDRelatedChangesOK {
	return &GetV1PostMortemsReportsReportIDRelatedChangesOK{}
}

/*GetV1PostMortemsReportsReportIDRelatedChangesOK handles this case with default header values.

Retrieve all related changes for the report
*/
type GetV1PostMortemsReportsReportIDRelatedChangesOK struct {
	Payload *models.RelatedChangeEntityPaginated
}

func (o *GetV1PostMortemsReportsReportIDRelatedChangesOK) Error() string {
	return fmt.Sprintf("[GET /v1/post_mortems/reports/{report_id}/related_changes][%d] getV1PostMortemsReportsReportIdRelatedChangesOK  %+v", 200, o.Payload)
}

func (o *GetV1PostMortemsReportsReportIDRelatedChangesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RelatedChangeEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
