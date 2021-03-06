// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteV1PostMortemsReportsReportIDEventsReportEventIDReader is a Reader for the DeleteV1PostMortemsReportsReportIDEventsReportEventID structure.
type DeleteV1PostMortemsReportsReportIDEventsReportEventIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1PostMortemsReportsReportIDEventsReportEventIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteV1PostMortemsReportsReportIDEventsReportEventIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteV1PostMortemsReportsReportIDEventsReportEventIDNoContent creates a DeleteV1PostMortemsReportsReportIDEventsReportEventIDNoContent with default headers values
func NewDeleteV1PostMortemsReportsReportIDEventsReportEventIDNoContent() *DeleteV1PostMortemsReportsReportIDEventsReportEventIDNoContent {
	return &DeleteV1PostMortemsReportsReportIDEventsReportEventIDNoContent{}
}

/*DeleteV1PostMortemsReportsReportIDEventsReportEventIDNoContent handles this case with default header values.

Remove an event on the report timeline by its ID
*/
type DeleteV1PostMortemsReportsReportIDEventsReportEventIDNoContent struct {
}

func (o *DeleteV1PostMortemsReportsReportIDEventsReportEventIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/post_mortems/reports/{report_id}/events/{report_event_id}][%d] deleteV1PostMortemsReportsReportIdEventsReportEventIdNoContent ", 204)
}

func (o *DeleteV1PostMortemsReportsReportIDEventsReportEventIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
