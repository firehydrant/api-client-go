// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1ServiceDependenciesServiceDependencyIDReader is a Reader for the GetV1ServiceDependenciesServiceDependencyID structure.
type GetV1ServiceDependenciesServiceDependencyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ServiceDependenciesServiceDependencyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ServiceDependenciesServiceDependencyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1ServiceDependenciesServiceDependencyIDOK creates a GetV1ServiceDependenciesServiceDependencyIDOK with default headers values
func NewGetV1ServiceDependenciesServiceDependencyIDOK() *GetV1ServiceDependenciesServiceDependencyIDOK {
	return &GetV1ServiceDependenciesServiceDependencyIDOK{}
}

/*
GetV1ServiceDependenciesServiceDependencyIDOK describes a response with status code 200, with default header values.

Retrieves a single service dependency by ID
*/
type GetV1ServiceDependenciesServiceDependencyIDOK struct {
	Payload *models.ServiceDependencyEntity
}

// IsSuccess returns true when this get v1 service dependencies service dependency Id o k response has a 2xx status code
func (o *GetV1ServiceDependenciesServiceDependencyIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 service dependencies service dependency Id o k response has a 3xx status code
func (o *GetV1ServiceDependenciesServiceDependencyIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 service dependencies service dependency Id o k response has a 4xx status code
func (o *GetV1ServiceDependenciesServiceDependencyIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 service dependencies service dependency Id o k response has a 5xx status code
func (o *GetV1ServiceDependenciesServiceDependencyIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 service dependencies service dependency Id o k response a status code equal to that given
func (o *GetV1ServiceDependenciesServiceDependencyIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1ServiceDependenciesServiceDependencyIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/service_dependencies/{service_dependency_id}][%d] getV1ServiceDependenciesServiceDependencyIdOK  %+v", 200, o.Payload)
}

func (o *GetV1ServiceDependenciesServiceDependencyIDOK) String() string {
	return fmt.Sprintf("[GET /v1/service_dependencies/{service_dependency_id}][%d] getV1ServiceDependenciesServiceDependencyIdOK  %+v", 200, o.Payload)
}

func (o *GetV1ServiceDependenciesServiceDependencyIDOK) GetPayload() *models.ServiceDependencyEntity {
	return o.Payload
}

func (o *GetV1ServiceDependenciesServiceDependencyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDependencyEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
