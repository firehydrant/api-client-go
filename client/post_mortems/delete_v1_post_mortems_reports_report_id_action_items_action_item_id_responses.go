// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteV1PostMortemsReportsReportIDActionItemsActionItemIDReader is a Reader for the DeleteV1PostMortemsReportsReportIDActionItemsActionItemID structure.
type DeleteV1PostMortemsReportsReportIDActionItemsActionItemIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1PostMortemsReportsReportIDActionItemsActionItemIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteV1PostMortemsReportsReportIDActionItemsActionItemIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteV1PostMortemsReportsReportIDActionItemsActionItemIDNoContent creates a DeleteV1PostMortemsReportsReportIDActionItemsActionItemIDNoContent with default headers values
func NewDeleteV1PostMortemsReportsReportIDActionItemsActionItemIDNoContent() *DeleteV1PostMortemsReportsReportIDActionItemsActionItemIDNoContent {
	return &DeleteV1PostMortemsReportsReportIDActionItemsActionItemIDNoContent{}
}

/*DeleteV1PostMortemsReportsReportIDActionItemsActionItemIDNoContent handles this case with default header values.

Delete an action item
*/
type DeleteV1PostMortemsReportsReportIDActionItemsActionItemIDNoContent struct {
}

func (o *DeleteV1PostMortemsReportsReportIDActionItemsActionItemIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/post_mortems/reports/{report_id}/action_items/{action_item_id}][%d] deleteV1PostMortemsReportsReportIdActionItemsActionItemIdNoContent ", 204)
}

func (o *DeleteV1PostMortemsReportsReportIDActionItemsActionItemIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
