// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PatchV1IntegrationsFieldMapsFieldMapIDReader is a Reader for the PatchV1IntegrationsFieldMapsFieldMapID structure.
type PatchV1IntegrationsFieldMapsFieldMapIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1IntegrationsFieldMapsFieldMapIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1IntegrationsFieldMapsFieldMapIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1IntegrationsFieldMapsFieldMapIDOK creates a PatchV1IntegrationsFieldMapsFieldMapIDOK with default headers values
func NewPatchV1IntegrationsFieldMapsFieldMapIDOK() *PatchV1IntegrationsFieldMapsFieldMapIDOK {
	return &PatchV1IntegrationsFieldMapsFieldMapIDOK{}
}

/*
PatchV1IntegrationsFieldMapsFieldMapIDOK describes a response with status code 200, with default header values.

Update field mapping
*/
type PatchV1IntegrationsFieldMapsFieldMapIDOK struct {
	Payload *models.FieldMappingFieldMapEntity
}

// IsSuccess returns true when this patch v1 integrations field maps field map Id o k response has a 2xx status code
func (o *PatchV1IntegrationsFieldMapsFieldMapIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 integrations field maps field map Id o k response has a 3xx status code
func (o *PatchV1IntegrationsFieldMapsFieldMapIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 integrations field maps field map Id o k response has a 4xx status code
func (o *PatchV1IntegrationsFieldMapsFieldMapIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 integrations field maps field map Id o k response has a 5xx status code
func (o *PatchV1IntegrationsFieldMapsFieldMapIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 integrations field maps field map Id o k response a status code equal to that given
func (o *PatchV1IntegrationsFieldMapsFieldMapIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchV1IntegrationsFieldMapsFieldMapIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/integrations/field_maps/{field_map_id}][%d] patchV1IntegrationsFieldMapsFieldMapIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1IntegrationsFieldMapsFieldMapIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/integrations/field_maps/{field_map_id}][%d] patchV1IntegrationsFieldMapsFieldMapIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1IntegrationsFieldMapsFieldMapIDOK) GetPayload() *models.FieldMappingFieldMapEntity {
	return o.Payload
}

func (o *PatchV1IntegrationsFieldMapsFieldMapIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FieldMappingFieldMapEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
