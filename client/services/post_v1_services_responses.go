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

// PostV1ServicesReader is a Reader for the PostV1Services structure.
type PostV1ServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1ServicesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1ServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1ServicesCreated creates a PostV1ServicesCreated with default headers values
func NewPostV1ServicesCreated() *PostV1ServicesCreated {
	return &PostV1ServicesCreated{}
}

/*
PostV1ServicesCreated describes a response with status code 201, with default header values.

Creates a service for the organization, you may also create or attach functionalities to the service on create.
*/
type PostV1ServicesCreated struct {
	Payload *models.ServiceEntity
}

// IsSuccess returns true when this post v1 services created response has a 2xx status code
func (o *PostV1ServicesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 services created response has a 3xx status code
func (o *PostV1ServicesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 services created response has a 4xx status code
func (o *PostV1ServicesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 services created response has a 5xx status code
func (o *PostV1ServicesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 services created response a status code equal to that given
func (o *PostV1ServicesCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1ServicesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/services][%d] postV1ServicesCreated  %+v", 201, o.Payload)
}

func (o *PostV1ServicesCreated) String() string {
	return fmt.Sprintf("[POST /v1/services][%d] postV1ServicesCreated  %+v", 201, o.Payload)
}

func (o *PostV1ServicesCreated) GetPayload() *models.ServiceEntity {
	return o.Payload
}

func (o *PostV1ServicesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ServicesBadRequest creates a PostV1ServicesBadRequest with default headers values
func NewPostV1ServicesBadRequest() *PostV1ServicesBadRequest {
	return &PostV1ServicesBadRequest{}
}

/*
PostV1ServicesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostV1ServicesBadRequest struct {
	Payload *models.ErrorEntity
}

// IsSuccess returns true when this post v1 services bad request response has a 2xx status code
func (o *PostV1ServicesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 services bad request response has a 3xx status code
func (o *PostV1ServicesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 services bad request response has a 4xx status code
func (o *PostV1ServicesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 services bad request response has a 5xx status code
func (o *PostV1ServicesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 services bad request response a status code equal to that given
func (o *PostV1ServicesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostV1ServicesBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/services][%d] postV1ServicesBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1ServicesBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/services][%d] postV1ServicesBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1ServicesBadRequest) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *PostV1ServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
