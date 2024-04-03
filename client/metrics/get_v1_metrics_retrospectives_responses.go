// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1MetricsRetrospectivesReader is a Reader for the GetV1MetricsRetrospectives structure.
type GetV1MetricsRetrospectivesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1MetricsRetrospectivesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1MetricsRetrospectivesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/metrics/retrospectives] getV1MetricsRetrospectives", response, response.Code())
	}
}

// NewGetV1MetricsRetrospectivesOK creates a GetV1MetricsRetrospectivesOK with default headers values
func NewGetV1MetricsRetrospectivesOK() *GetV1MetricsRetrospectivesOK {
	return &GetV1MetricsRetrospectivesOK{}
}

/*
GetV1MetricsRetrospectivesOK describes a response with status code 200, with default header values.

Returns a report with retrospective analytics data
*/
type GetV1MetricsRetrospectivesOK struct {
	Payload *models.MetricsRetrospectiveEntity
}

// IsSuccess returns true when this get v1 metrics retrospectives o k response has a 2xx status code
func (o *GetV1MetricsRetrospectivesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 metrics retrospectives o k response has a 3xx status code
func (o *GetV1MetricsRetrospectivesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 metrics retrospectives o k response has a 4xx status code
func (o *GetV1MetricsRetrospectivesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 metrics retrospectives o k response has a 5xx status code
func (o *GetV1MetricsRetrospectivesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 metrics retrospectives o k response a status code equal to that given
func (o *GetV1MetricsRetrospectivesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 metrics retrospectives o k response
func (o *GetV1MetricsRetrospectivesOK) Code() int {
	return 200
}

func (o *GetV1MetricsRetrospectivesOK) Error() string {
	return fmt.Sprintf("[GET /v1/metrics/retrospectives][%d] getV1MetricsRetrospectivesOK  %+v", 200, o.Payload)
}

func (o *GetV1MetricsRetrospectivesOK) String() string {
	return fmt.Sprintf("[GET /v1/metrics/retrospectives][%d] getV1MetricsRetrospectivesOK  %+v", 200, o.Payload)
}

func (o *GetV1MetricsRetrospectivesOK) GetPayload() *models.MetricsRetrospectiveEntity {
	return o.Payload
}

func (o *GetV1MetricsRetrospectivesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetricsRetrospectiveEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
