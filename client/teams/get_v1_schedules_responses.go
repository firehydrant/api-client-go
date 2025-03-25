// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1SchedulesReader is a Reader for the GetV1Schedules structure.
type GetV1SchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1SchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1SchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1SchedulesOK creates a GetV1SchedulesOK with default headers values
func NewGetV1SchedulesOK() *GetV1SchedulesOK {
	return &GetV1SchedulesOK{}
}

/*
GetV1SchedulesOK describes a response with status code 200, with default header values.

List all known schedules in FireHydrant as pulled from external sources
*/
type GetV1SchedulesOK struct {
	Payload *models.ScheduleEntityPaginated
}

// IsSuccess returns true when this get v1 schedules o k response has a 2xx status code
func (o *GetV1SchedulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 schedules o k response has a 3xx status code
func (o *GetV1SchedulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 schedules o k response has a 4xx status code
func (o *GetV1SchedulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 schedules o k response has a 5xx status code
func (o *GetV1SchedulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 schedules o k response a status code equal to that given
func (o *GetV1SchedulesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1SchedulesOK) Error() string {
	return fmt.Sprintf("[GET /v1/schedules][%d] getV1SchedulesOK  %+v", 200, o.Payload)
}

func (o *GetV1SchedulesOK) String() string {
	return fmt.Sprintf("[GET /v1/schedules][%d] getV1SchedulesOK  %+v", 200, o.Payload)
}

func (o *GetV1SchedulesOK) GetPayload() *models.ScheduleEntityPaginated {
	return o.Payload
}

func (o *GetV1SchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduleEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
