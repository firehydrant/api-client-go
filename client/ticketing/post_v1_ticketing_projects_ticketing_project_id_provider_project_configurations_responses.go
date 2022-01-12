// Code generated by go-swagger; DO NOT EDIT.

package ticketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsReader is a Reader for the PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurations structure.
type PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsCreated creates a PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsCreated with default headers values
func NewPostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsCreated() *PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsCreated {
	return &PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsCreated{}
}

/* PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsCreated describes a response with status code 201, with default header values.

Creates configuration for a ticketing project
*/
type PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsCreated struct {
	Payload *models.ProjectConfigEntity
}

func (o *PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/ticketing/projects/{ticketing_project_id}/provider_project_configurations][%d] postV1TicketingProjectsTicketingProjectIdProviderProjectConfigurationsCreated  %+v", 201, o.Payload)
}
func (o *PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsCreated) GetPayload() *models.ProjectConfigEntity {
	return o.Payload
}

func (o *PostV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectConfigEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}