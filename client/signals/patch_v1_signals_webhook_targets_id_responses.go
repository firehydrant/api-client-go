// Code generated by go-swagger; DO NOT EDIT.

package signals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchV1SignalsWebhookTargetsIDReader is a Reader for the PatchV1SignalsWebhookTargetsID structure.
type PatchV1SignalsWebhookTargetsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1SignalsWebhookTargetsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1SignalsWebhookTargetsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1SignalsWebhookTargetsIDOK creates a PatchV1SignalsWebhookTargetsIDOK with default headers values
func NewPatchV1SignalsWebhookTargetsIDOK() *PatchV1SignalsWebhookTargetsIDOK {
	return &PatchV1SignalsWebhookTargetsIDOK{}
}

/*
PatchV1SignalsWebhookTargetsIDOK describes a response with status code 200, with default header values.

Update a Signals webhook target by ID
*/
type PatchV1SignalsWebhookTargetsIDOK struct {
}

// IsSuccess returns true when this patch v1 signals webhook targets Id o k response has a 2xx status code
func (o *PatchV1SignalsWebhookTargetsIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 signals webhook targets Id o k response has a 3xx status code
func (o *PatchV1SignalsWebhookTargetsIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 signals webhook targets Id o k response has a 4xx status code
func (o *PatchV1SignalsWebhookTargetsIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 signals webhook targets Id o k response has a 5xx status code
func (o *PatchV1SignalsWebhookTargetsIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 signals webhook targets Id o k response a status code equal to that given
func (o *PatchV1SignalsWebhookTargetsIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchV1SignalsWebhookTargetsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/signals/webhook_targets/{id}][%d] patchV1SignalsWebhookTargetsIdOK ", 200)
}

func (o *PatchV1SignalsWebhookTargetsIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/signals/webhook_targets/{id}][%d] patchV1SignalsWebhookTargetsIdOK ", 200)
}

func (o *PatchV1SignalsWebhookTargetsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}