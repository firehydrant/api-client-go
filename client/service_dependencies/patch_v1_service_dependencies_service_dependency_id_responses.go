// Code generated by go-swagger; DO NOT EDIT.

package service_dependencies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PatchV1ServiceDependenciesServiceDependencyIDReader is a Reader for the PatchV1ServiceDependenciesServiceDependencyID structure.
type PatchV1ServiceDependenciesServiceDependencyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1ServiceDependenciesServiceDependencyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1ServiceDependenciesServiceDependencyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1ServiceDependenciesServiceDependencyIDOK creates a PatchV1ServiceDependenciesServiceDependencyIDOK with default headers values
func NewPatchV1ServiceDependenciesServiceDependencyIDOK() *PatchV1ServiceDependenciesServiceDependencyIDOK {
	return &PatchV1ServiceDependenciesServiceDependencyIDOK{}
}

/*
PatchV1ServiceDependenciesServiceDependencyIDOK describes a response with status code 200, with default header values.

Update the notes of the service dependency
*/
type PatchV1ServiceDependenciesServiceDependencyIDOK struct {
	Payload *models.ServiceDependencyEntity
}

// IsSuccess returns true when this patch v1 service dependencies service dependency Id o k response has a 2xx status code
func (o *PatchV1ServiceDependenciesServiceDependencyIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 service dependencies service dependency Id o k response has a 3xx status code
func (o *PatchV1ServiceDependenciesServiceDependencyIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 service dependencies service dependency Id o k response has a 4xx status code
func (o *PatchV1ServiceDependenciesServiceDependencyIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 service dependencies service dependency Id o k response has a 5xx status code
func (o *PatchV1ServiceDependenciesServiceDependencyIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 service dependencies service dependency Id o k response a status code equal to that given
func (o *PatchV1ServiceDependenciesServiceDependencyIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchV1ServiceDependenciesServiceDependencyIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/service_dependencies/{service_dependency_id}][%d] patchV1ServiceDependenciesServiceDependencyIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1ServiceDependenciesServiceDependencyIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/service_dependencies/{service_dependency_id}][%d] patchV1ServiceDependenciesServiceDependencyIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1ServiceDependenciesServiceDependencyIDOK) GetPayload() *models.ServiceDependencyEntity {
	return o.Payload
}

func (o *PatchV1ServiceDependenciesServiceDependencyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDependencyEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
