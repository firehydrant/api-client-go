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

// PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDReader is a Reader for the PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigID structure.
type PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/ticketing/projects/{ticketing_project_id}/provider_project_configurations/{config_id}] patchV1TicketingProjectsTicketingProjectIdProviderProjectConfigurationsConfigId", response, response.Code())
	}
}

// NewPatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK creates a PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK with default headers values
func NewPatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK() *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK {
	return &PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK{}
}

/*
PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK describes a response with status code 200, with default header values.

Update configuration for a ticketing project
*/
type PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK struct {
	Payload *models.TicketingProjectConfigEntity
}

// IsSuccess returns true when this patch v1 ticketing projects ticketing project Id provider project configurations config Id o k response has a 2xx status code
func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 ticketing projects ticketing project Id provider project configurations config Id o k response has a 3xx status code
func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 ticketing projects ticketing project Id provider project configurations config Id o k response has a 4xx status code
func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 ticketing projects ticketing project Id provider project configurations config Id o k response has a 5xx status code
func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 ticketing projects ticketing project Id provider project configurations config Id o k response a status code equal to that given
func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch v1 ticketing projects ticketing project Id provider project configurations config Id o k response
func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) Code() int {
	return 200
}

func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/ticketing/projects/{ticketing_project_id}/provider_project_configurations/{config_id}][%d] patchV1TicketingProjectsTicketingProjectIdProviderProjectConfigurationsConfigIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/ticketing/projects/{ticketing_project_id}/provider_project_configurations/{config_id}][%d] patchV1TicketingProjectsTicketingProjectIdProviderProjectConfigurationsConfigIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) GetPayload() *models.TicketingProjectConfigEntity {
	return o.Payload
}

func (o *PatchV1TicketingProjectsTicketingProjectIDProviderProjectConfigurationsConfigIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TicketingProjectConfigEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
