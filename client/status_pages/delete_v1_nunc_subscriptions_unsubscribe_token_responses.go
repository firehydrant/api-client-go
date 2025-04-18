// Code generated by go-swagger; DO NOT EDIT.

package status_pages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// DeleteV1NuncSubscriptionsUnsubscribeTokenReader is a Reader for the DeleteV1NuncSubscriptionsUnsubscribeToken structure.
type DeleteV1NuncSubscriptionsUnsubscribeTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1NuncSubscriptionsUnsubscribeTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1NuncSubscriptionsUnsubscribeTokenOK creates a DeleteV1NuncSubscriptionsUnsubscribeTokenOK with default headers values
func NewDeleteV1NuncSubscriptionsUnsubscribeTokenOK() *DeleteV1NuncSubscriptionsUnsubscribeTokenOK {
	return &DeleteV1NuncSubscriptionsUnsubscribeTokenOK{}
}

/*
DeleteV1NuncSubscriptionsUnsubscribeTokenOK describes a response with status code 200, with default header values.

Unsubscribe from status page updates
*/
type DeleteV1NuncSubscriptionsUnsubscribeTokenOK struct {
	Payload *models.NuncNuncSubscription
}

// IsSuccess returns true when this delete v1 nunc subscriptions unsubscribe token o k response has a 2xx status code
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 nunc subscriptions unsubscribe token o k response has a 3xx status code
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 nunc subscriptions unsubscribe token o k response has a 4xx status code
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 nunc subscriptions unsubscribe token o k response has a 5xx status code
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 nunc subscriptions unsubscribe token o k response a status code equal to that given
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/nunc/subscriptions/{unsubscribe_token}][%d] deleteV1NuncSubscriptionsUnsubscribeTokenOK  %+v", 200, o.Payload)
}

func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenOK) String() string {
	return fmt.Sprintf("[DELETE /v1/nunc/subscriptions/{unsubscribe_token}][%d] deleteV1NuncSubscriptionsUnsubscribeTokenOK  %+v", 200, o.Payload)
}

func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenOK) GetPayload() *models.NuncNuncSubscription {
	return o.Payload
}

func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NuncNuncSubscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
