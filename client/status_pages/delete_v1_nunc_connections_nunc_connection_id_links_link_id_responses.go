// Code generated by go-swagger; DO NOT EDIT.

package status_pages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDReader is a Reader for the DeleteV1NuncConnectionsNuncConnectionIDLinksLinkID structure.
type DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent creates a DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent with default headers values
func NewDeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent() *DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent {
	return &DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent{}
}

/*
DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent describes a response with status code 204, with default header values.

Delete a link displayed on a FireHydrant status page
*/
type DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent struct {
}

// IsSuccess returns true when this delete v1 nunc connections nunc connection Id links link Id no content response has a 2xx status code
func (o *DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 nunc connections nunc connection Id links link Id no content response has a 3xx status code
func (o *DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 nunc connections nunc connection Id links link Id no content response has a 4xx status code
func (o *DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 nunc connections nunc connection Id links link Id no content response has a 5xx status code
func (o *DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 nunc connections nunc connection Id links link Id no content response a status code equal to that given
func (o *DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/nunc_connections/{nunc_connection_id}/links/{link_id}][%d] deleteV1NuncConnectionsNuncConnectionIdLinksLinkIdNoContent ", 204)
}

func (o *DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/nunc_connections/{nunc_connection_id}/links/{link_id}][%d] deleteV1NuncConnectionsNuncConnectionIdLinksLinkIdNoContent ", 204)
}

func (o *DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
