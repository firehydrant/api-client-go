// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firehydrant/api-client-go/models"
)

// PatchV1EnvironmentsIDReader is a Reader for the PatchV1EnvironmentsID structure.
type PatchV1EnvironmentsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1EnvironmentsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchV1EnvironmentsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchV1EnvironmentsIDOK creates a PatchV1EnvironmentsIDOK with default headers values
func NewPatchV1EnvironmentsIDOK() *PatchV1EnvironmentsIDOK {
	return &PatchV1EnvironmentsIDOK{}
}

/*PatchV1EnvironmentsIDOK handles this case with default header values.

Update an environment
*/
type PatchV1EnvironmentsIDOK struct {
	Payload *models.EnvironmentEntity
}

func (o *PatchV1EnvironmentsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/environments/{id}][%d] patchV1EnvironmentsIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1EnvironmentsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
