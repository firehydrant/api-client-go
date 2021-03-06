// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteV1ComponentsIDReader is a Reader for the DeleteV1ComponentsID structure.
type DeleteV1ComponentsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1ComponentsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteV1ComponentsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteV1ComponentsIDNoContent creates a DeleteV1ComponentsIDNoContent with default headers values
func NewDeleteV1ComponentsIDNoContent() *DeleteV1ComponentsIDNoContent {
	return &DeleteV1ComponentsIDNoContent{}
}

/*DeleteV1ComponentsIDNoContent handles this case with default header values.

Archive an component
*/
type DeleteV1ComponentsIDNoContent struct {
}

func (o *DeleteV1ComponentsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/components/{id}][%d] deleteV1ComponentsIdNoContent ", 204)
}

func (o *DeleteV1ComponentsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
