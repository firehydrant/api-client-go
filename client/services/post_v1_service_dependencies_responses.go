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

// PostV1ServiceDependenciesReader is a Reader for the PostV1ServiceDependencies structure.
type PostV1ServiceDependenciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ServiceDependenciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1ServiceDependenciesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1ServiceDependenciesCreated creates a PostV1ServiceDependenciesCreated with default headers values
func NewPostV1ServiceDependenciesCreated() *PostV1ServiceDependenciesCreated {
	return &PostV1ServiceDependenciesCreated{}
}

/*
PostV1ServiceDependenciesCreated describes a response with status code 201, with default header values.

Creates a service dependency relationship between two services
*/
type PostV1ServiceDependenciesCreated struct {
	Payload *models.ServiceDependencyEntity
}

// IsSuccess returns true when this post v1 service dependencies created response has a 2xx status code
func (o *PostV1ServiceDependenciesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 service dependencies created response has a 3xx status code
func (o *PostV1ServiceDependenciesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 service dependencies created response has a 4xx status code
func (o *PostV1ServiceDependenciesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 service dependencies created response has a 5xx status code
func (o *PostV1ServiceDependenciesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 service dependencies created response a status code equal to that given
func (o *PostV1ServiceDependenciesCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1ServiceDependenciesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/service_dependencies][%d] postV1ServiceDependenciesCreated  %+v", 201, o.Payload)
}

func (o *PostV1ServiceDependenciesCreated) String() string {
	return fmt.Sprintf("[POST /v1/service_dependencies][%d] postV1ServiceDependenciesCreated  %+v", 201, o.Payload)
}

func (o *PostV1ServiceDependenciesCreated) GetPayload() *models.ServiceDependencyEntity {
	return o.Payload
}

func (o *PostV1ServiceDependenciesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDependencyEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
