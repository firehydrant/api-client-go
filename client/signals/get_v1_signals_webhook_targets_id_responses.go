// Code generated by go-swagger; DO NOT EDIT.

package signals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1SignalsWebhookTargetsIDReader is a Reader for the GetV1SignalsWebhookTargetsID structure.
type GetV1SignalsWebhookTargetsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1SignalsWebhookTargetsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1SignalsWebhookTargetsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1SignalsWebhookTargetsIDOK creates a GetV1SignalsWebhookTargetsIDOK with default headers values
func NewGetV1SignalsWebhookTargetsIDOK() *GetV1SignalsWebhookTargetsIDOK {
	return &GetV1SignalsWebhookTargetsIDOK{}
}

/*
GetV1SignalsWebhookTargetsIDOK describes a response with status code 200, with default header values.

Get a Signals webhook target by ID
*/
type GetV1SignalsWebhookTargetsIDOK struct {
}

// IsSuccess returns true when this get v1 signals webhook targets Id o k response has a 2xx status code
func (o *GetV1SignalsWebhookTargetsIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 signals webhook targets Id o k response has a 3xx status code
func (o *GetV1SignalsWebhookTargetsIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 signals webhook targets Id o k response has a 4xx status code
func (o *GetV1SignalsWebhookTargetsIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 signals webhook targets Id o k response has a 5xx status code
func (o *GetV1SignalsWebhookTargetsIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 signals webhook targets Id o k response a status code equal to that given
func (o *GetV1SignalsWebhookTargetsIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1SignalsWebhookTargetsIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/signals/webhook_targets/{id}][%d] getV1SignalsWebhookTargetsIdOK ", 200)
}

func (o *GetV1SignalsWebhookTargetsIDOK) String() string {
	return fmt.Sprintf("[GET /v1/signals/webhook_targets/{id}][%d] getV1SignalsWebhookTargetsIdOK ", 200)
}

func (o *GetV1SignalsWebhookTargetsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
