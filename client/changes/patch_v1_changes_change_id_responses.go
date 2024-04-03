// Code generated by go-swagger; DO NOT EDIT.

package changes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PatchV1ChangesChangeIDReader is a Reader for the PatchV1ChangesChangeID structure.
type PatchV1ChangesChangeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1ChangesChangeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1ChangesChangeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/changes/{change_id}] patchV1ChangesChangeId", response, response.Code())
	}
}

// NewPatchV1ChangesChangeIDOK creates a PatchV1ChangesChangeIDOK with default headers values
func NewPatchV1ChangesChangeIDOK() *PatchV1ChangesChangeIDOK {
	return &PatchV1ChangesChangeIDOK{}
}

/*
PatchV1ChangesChangeIDOK describes a response with status code 200, with default header values.

Update a change entry
*/
type PatchV1ChangesChangeIDOK struct {
	Payload *models.ChangeEntity
}

// IsSuccess returns true when this patch v1 changes change Id o k response has a 2xx status code
func (o *PatchV1ChangesChangeIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 changes change Id o k response has a 3xx status code
func (o *PatchV1ChangesChangeIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 changes change Id o k response has a 4xx status code
func (o *PatchV1ChangesChangeIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 changes change Id o k response has a 5xx status code
func (o *PatchV1ChangesChangeIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 changes change Id o k response a status code equal to that given
func (o *PatchV1ChangesChangeIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch v1 changes change Id o k response
func (o *PatchV1ChangesChangeIDOK) Code() int {
	return 200
}

func (o *PatchV1ChangesChangeIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/changes/{change_id}][%d] patchV1ChangesChangeIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1ChangesChangeIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/changes/{change_id}][%d] patchV1ChangesChangeIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1ChangesChangeIDOK) GetPayload() *models.ChangeEntity {
	return o.Payload
}

func (o *PatchV1ChangesChangeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChangeEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
