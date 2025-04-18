// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1WebhooksWebhookIDDeliveriesReader is a Reader for the GetV1WebhooksWebhookIDDeliveries structure.
type GetV1WebhooksWebhookIDDeliveriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1WebhooksWebhookIDDeliveriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1WebhooksWebhookIDDeliveriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1WebhooksWebhookIDDeliveriesOK creates a GetV1WebhooksWebhookIDDeliveriesOK with default headers values
func NewGetV1WebhooksWebhookIDDeliveriesOK() *GetV1WebhooksWebhookIDDeliveriesOK {
	return &GetV1WebhooksWebhookIDDeliveriesOK{}
}

/*
GetV1WebhooksWebhookIDDeliveriesOK describes a response with status code 200, with default header values.

Get webhook deliveries
*/
type GetV1WebhooksWebhookIDDeliveriesOK struct {
}

// IsSuccess returns true when this get v1 webhooks webhook Id deliveries o k response has a 2xx status code
func (o *GetV1WebhooksWebhookIDDeliveriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 webhooks webhook Id deliveries o k response has a 3xx status code
func (o *GetV1WebhooksWebhookIDDeliveriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 webhooks webhook Id deliveries o k response has a 4xx status code
func (o *GetV1WebhooksWebhookIDDeliveriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 webhooks webhook Id deliveries o k response has a 5xx status code
func (o *GetV1WebhooksWebhookIDDeliveriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 webhooks webhook Id deliveries o k response a status code equal to that given
func (o *GetV1WebhooksWebhookIDDeliveriesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1WebhooksWebhookIDDeliveriesOK) Error() string {
	return fmt.Sprintf("[GET /v1/webhooks/{webhook_id}/deliveries][%d] getV1WebhooksWebhookIdDeliveriesOK ", 200)
}

func (o *GetV1WebhooksWebhookIDDeliveriesOK) String() string {
	return fmt.Sprintf("[GET /v1/webhooks/{webhook_id}/deliveries][%d] getV1WebhooksWebhookIdDeliveriesOK ", 200)
}

func (o *GetV1WebhooksWebhookIDDeliveriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
