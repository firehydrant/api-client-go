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

// PostV1ServicesServiceLinksReader is a Reader for the PostV1ServicesServiceLinks structure.
type PostV1ServicesServiceLinksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ServicesServiceLinksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1ServicesServiceLinksCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1ServicesServiceLinksCreated creates a PostV1ServicesServiceLinksCreated with default headers values
func NewPostV1ServicesServiceLinksCreated() *PostV1ServicesServiceLinksCreated {
	return &PostV1ServicesServiceLinksCreated{}
}

/*
PostV1ServicesServiceLinksCreated describes a response with status code 201, with default header values.

Creates a service with the appropriate integration for each external service ID passed
*/
type PostV1ServicesServiceLinksCreated struct {
	Payload []*models.ServiceLinkEntity
}

// IsSuccess returns true when this post v1 services service links created response has a 2xx status code
func (o *PostV1ServicesServiceLinksCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 services service links created response has a 3xx status code
func (o *PostV1ServicesServiceLinksCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 services service links created response has a 4xx status code
func (o *PostV1ServicesServiceLinksCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 services service links created response has a 5xx status code
func (o *PostV1ServicesServiceLinksCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 services service links created response a status code equal to that given
func (o *PostV1ServicesServiceLinksCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1ServicesServiceLinksCreated) Error() string {
	return fmt.Sprintf("[POST /v1/services/service_links][%d] postV1ServicesServiceLinksCreated  %+v", 201, o.Payload)
}

func (o *PostV1ServicesServiceLinksCreated) String() string {
	return fmt.Sprintf("[POST /v1/services/service_links][%d] postV1ServicesServiceLinksCreated  %+v", 201, o.Payload)
}

func (o *PostV1ServicesServiceLinksCreated) GetPayload() []*models.ServiceLinkEntity {
	return o.Payload
}

func (o *PostV1ServicesServiceLinksCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
