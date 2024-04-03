// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostV1IntegrationsConnectionsSlugReader is a Reader for the PostV1IntegrationsConnectionsSlug structure.
type PostV1IntegrationsConnectionsSlugReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1IntegrationsConnectionsSlugReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1IntegrationsConnectionsSlugCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/integrations/connections/{slug}] postV1IntegrationsConnectionsSlug", response, response.Code())
	}
}

// NewPostV1IntegrationsConnectionsSlugCreated creates a PostV1IntegrationsConnectionsSlugCreated with default headers values
func NewPostV1IntegrationsConnectionsSlugCreated() *PostV1IntegrationsConnectionsSlugCreated {
	return &PostV1IntegrationsConnectionsSlugCreated{}
}

/*
PostV1IntegrationsConnectionsSlugCreated describes a response with status code 201, with default header values.

created Connection
*/
type PostV1IntegrationsConnectionsSlugCreated struct {
}

// IsSuccess returns true when this post v1 integrations connections slug created response has a 2xx status code
func (o *PostV1IntegrationsConnectionsSlugCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 integrations connections slug created response has a 3xx status code
func (o *PostV1IntegrationsConnectionsSlugCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 integrations connections slug created response has a 4xx status code
func (o *PostV1IntegrationsConnectionsSlugCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 integrations connections slug created response has a 5xx status code
func (o *PostV1IntegrationsConnectionsSlugCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 integrations connections slug created response a status code equal to that given
func (o *PostV1IntegrationsConnectionsSlugCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post v1 integrations connections slug created response
func (o *PostV1IntegrationsConnectionsSlugCreated) Code() int {
	return 201
}

func (o *PostV1IntegrationsConnectionsSlugCreated) Error() string {
	return fmt.Sprintf("[POST /v1/integrations/connections/{slug}][%d] postV1IntegrationsConnectionsSlugCreated ", 201)
}

func (o *PostV1IntegrationsConnectionsSlugCreated) String() string {
	return fmt.Sprintf("[POST /v1/integrations/connections/{slug}][%d] postV1IntegrationsConnectionsSlugCreated ", 201)
}

func (o *PostV1IntegrationsConnectionsSlugCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
