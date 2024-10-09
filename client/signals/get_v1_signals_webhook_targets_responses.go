// Code generated by go-swagger; DO NOT EDIT.

package signals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1SignalsWebhookTargetsReader is a Reader for the GetV1SignalsWebhookTargets structure.
type GetV1SignalsWebhookTargetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1SignalsWebhookTargetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1SignalsWebhookTargetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1SignalsWebhookTargetsOK creates a GetV1SignalsWebhookTargetsOK with default headers values
func NewGetV1SignalsWebhookTargetsOK() *GetV1SignalsWebhookTargetsOK {
	return &GetV1SignalsWebhookTargetsOK{}
}

/*
GetV1SignalsWebhookTargetsOK describes a response with status code 200, with default header values.

List all Signals webhook targets.
*/
type GetV1SignalsWebhookTargetsOK struct {
}

// IsSuccess returns true when this get v1 signals webhook targets o k response has a 2xx status code
func (o *GetV1SignalsWebhookTargetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 signals webhook targets o k response has a 3xx status code
func (o *GetV1SignalsWebhookTargetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 signals webhook targets o k response has a 4xx status code
func (o *GetV1SignalsWebhookTargetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 signals webhook targets o k response has a 5xx status code
func (o *GetV1SignalsWebhookTargetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 signals webhook targets o k response a status code equal to that given
func (o *GetV1SignalsWebhookTargetsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1SignalsWebhookTargetsOK) Error() string {
	return fmt.Sprintf("[GET /v1/signals/webhook_targets][%d] getV1SignalsWebhookTargetsOK ", 200)
}

func (o *GetV1SignalsWebhookTargetsOK) String() string {
	return fmt.Sprintf("[GET /v1/signals/webhook_targets][%d] getV1SignalsWebhookTargetsOK ", 200)
}

func (o *GetV1SignalsWebhookTargetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
