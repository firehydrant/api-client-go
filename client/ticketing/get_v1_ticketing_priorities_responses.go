// Code generated by go-swagger; DO NOT EDIT.

package ticketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1TicketingPrioritiesReader is a Reader for the GetV1TicketingPriorities structure.
type GetV1TicketingPrioritiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1TicketingPrioritiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1TicketingPrioritiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/ticketing/priorities] getV1TicketingPriorities", response, response.Code())
	}
}

// NewGetV1TicketingPrioritiesOK creates a GetV1TicketingPrioritiesOK with default headers values
func NewGetV1TicketingPrioritiesOK() *GetV1TicketingPrioritiesOK {
	return &GetV1TicketingPrioritiesOK{}
}

/*
GetV1TicketingPrioritiesOK describes a response with status code 200, with default header values.

List all ticketing priorities available to the organization
*/
type GetV1TicketingPrioritiesOK struct {
	Payload *models.TicketingPriorityEntity
}

// IsSuccess returns true when this get v1 ticketing priorities o k response has a 2xx status code
func (o *GetV1TicketingPrioritiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 ticketing priorities o k response has a 3xx status code
func (o *GetV1TicketingPrioritiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 ticketing priorities o k response has a 4xx status code
func (o *GetV1TicketingPrioritiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 ticketing priorities o k response has a 5xx status code
func (o *GetV1TicketingPrioritiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 ticketing priorities o k response a status code equal to that given
func (o *GetV1TicketingPrioritiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 ticketing priorities o k response
func (o *GetV1TicketingPrioritiesOK) Code() int {
	return 200
}

func (o *GetV1TicketingPrioritiesOK) Error() string {
	return fmt.Sprintf("[GET /v1/ticketing/priorities][%d] getV1TicketingPrioritiesOK  %+v", 200, o.Payload)
}

func (o *GetV1TicketingPrioritiesOK) String() string {
	return fmt.Sprintf("[GET /v1/ticketing/priorities][%d] getV1TicketingPrioritiesOK  %+v", 200, o.Payload)
}

func (o *GetV1TicketingPrioritiesOK) GetPayload() *models.TicketingPriorityEntity {
	return o.Payload
}

func (o *GetV1TicketingPrioritiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TicketingPriorityEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
