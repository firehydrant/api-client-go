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

// PostV1SavedSearchesResourceTypeReader is a Reader for the PostV1SavedSearchesResourceType structure.
type PostV1SavedSearchesResourceTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1SavedSearchesResourceTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1SavedSearchesResourceTypeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1SavedSearchesResourceTypeCreated creates a PostV1SavedSearchesResourceTypeCreated with default headers values
func NewPostV1SavedSearchesResourceTypeCreated() *PostV1SavedSearchesResourceTypeCreated {
	return &PostV1SavedSearchesResourceTypeCreated{}
}

/*
PostV1SavedSearchesResourceTypeCreated describes a response with status code 201, with default header values.

Create a new saved search for a particular resource type
*/
type PostV1SavedSearchesResourceTypeCreated struct {
	Payload *models.SavedSearchEntity
}

// IsSuccess returns true when this post v1 saved searches resource type created response has a 2xx status code
func (o *PostV1SavedSearchesResourceTypeCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 saved searches resource type created response has a 3xx status code
func (o *PostV1SavedSearchesResourceTypeCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 saved searches resource type created response has a 4xx status code
func (o *PostV1SavedSearchesResourceTypeCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 saved searches resource type created response has a 5xx status code
func (o *PostV1SavedSearchesResourceTypeCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 saved searches resource type created response a status code equal to that given
func (o *PostV1SavedSearchesResourceTypeCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1SavedSearchesResourceTypeCreated) Error() string {
	return fmt.Sprintf("[POST /v1/saved_searches/{resource_type}][%d] postV1SavedSearchesResourceTypeCreated  %+v", 201, o.Payload)
}

func (o *PostV1SavedSearchesResourceTypeCreated) String() string {
	return fmt.Sprintf("[POST /v1/saved_searches/{resource_type}][%d] postV1SavedSearchesResourceTypeCreated  %+v", 201, o.Payload)
}

func (o *PostV1SavedSearchesResourceTypeCreated) GetPayload() *models.SavedSearchEntity {
	return o.Payload
}

func (o *PostV1SavedSearchesResourceTypeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SavedSearchEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
