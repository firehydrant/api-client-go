// Code generated by go-swagger; DO NOT EDIT.

package reporting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1MetricsInfraTypeReader is a Reader for the GetV1MetricsInfraType structure.
type GetV1MetricsInfraTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1MetricsInfraTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1MetricsInfraTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1MetricsInfraTypeOK creates a GetV1MetricsInfraTypeOK with default headers values
func NewGetV1MetricsInfraTypeOK() *GetV1MetricsInfraTypeOK {
	return &GetV1MetricsInfraTypeOK{}
}

/*
GetV1MetricsInfraTypeOK describes a response with status code 200, with default header values.

Returns metrics for all components of a given type
*/
type GetV1MetricsInfraTypeOK struct {
	Payload *models.MetricsInfrastructureListEntity
}

// IsSuccess returns true when this get v1 metrics infra type o k response has a 2xx status code
func (o *GetV1MetricsInfraTypeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 metrics infra type o k response has a 3xx status code
func (o *GetV1MetricsInfraTypeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 metrics infra type o k response has a 4xx status code
func (o *GetV1MetricsInfraTypeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 metrics infra type o k response has a 5xx status code
func (o *GetV1MetricsInfraTypeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 metrics infra type o k response a status code equal to that given
func (o *GetV1MetricsInfraTypeOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1MetricsInfraTypeOK) Error() string {
	return fmt.Sprintf("[GET /v1/metrics/{infra_type}][%d] getV1MetricsInfraTypeOK  %+v", 200, o.Payload)
}

func (o *GetV1MetricsInfraTypeOK) String() string {
	return fmt.Sprintf("[GET /v1/metrics/{infra_type}][%d] getV1MetricsInfraTypeOK  %+v", 200, o.Payload)
}

func (o *GetV1MetricsInfraTypeOK) GetPayload() *models.MetricsInfrastructureListEntity {
	return o.Payload
}

func (o *GetV1MetricsInfraTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetricsInfrastructureListEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
