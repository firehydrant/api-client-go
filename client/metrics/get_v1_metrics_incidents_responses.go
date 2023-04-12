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

// GetV1MetricsIncidentsReader is a Reader for the GetV1MetricsIncidents structure.
type GetV1MetricsIncidentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1MetricsIncidentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1MetricsIncidentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1MetricsIncidentsOK creates a GetV1MetricsIncidentsOK with default headers values
func NewGetV1MetricsIncidentsOK() *GetV1MetricsIncidentsOK {
	return &GetV1MetricsIncidentsOK{}
}

/*
GetV1MetricsIncidentsOK describes a response with status code 200, with default header values.

Returns a report with time bucketed analytics data
*/
type GetV1MetricsIncidentsOK struct {
	Payload *models.MetricsMetricsEntity
}

// IsSuccess returns true when this get v1 metrics incidents o k response has a 2xx status code
func (o *GetV1MetricsIncidentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 metrics incidents o k response has a 3xx status code
func (o *GetV1MetricsIncidentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 metrics incidents o k response has a 4xx status code
func (o *GetV1MetricsIncidentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 metrics incidents o k response has a 5xx status code
func (o *GetV1MetricsIncidentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 metrics incidents o k response a status code equal to that given
func (o *GetV1MetricsIncidentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1MetricsIncidentsOK) Error() string {
	return fmt.Sprintf("[GET /v1/metrics/incidents][%d] getV1MetricsIncidentsOK  %+v", 200, o.Payload)
}

func (o *GetV1MetricsIncidentsOK) String() string {
	return fmt.Sprintf("[GET /v1/metrics/incidents][%d] getV1MetricsIncidentsOK  %+v", 200, o.Payload)
}

func (o *GetV1MetricsIncidentsOK) GetPayload() *models.MetricsMetricsEntity {
	return o.Payload
}

func (o *GetV1MetricsIncidentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetricsMetricsEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
