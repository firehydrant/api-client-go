// Code generated by go-swagger; DO NOT EDIT.

package severities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PatchV1SeveritiesSeveritySlugReader is a Reader for the PatchV1SeveritiesSeveritySlug structure.
type PatchV1SeveritiesSeveritySlugReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1SeveritiesSeveritySlugReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1SeveritiesSeveritySlugOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/severities/{severity_slug}] patchV1SeveritiesSeveritySlug", response, response.Code())
	}
}

// NewPatchV1SeveritiesSeveritySlugOK creates a PatchV1SeveritiesSeveritySlugOK with default headers values
func NewPatchV1SeveritiesSeveritySlugOK() *PatchV1SeveritiesSeveritySlugOK {
	return &PatchV1SeveritiesSeveritySlugOK{}
}

/*
PatchV1SeveritiesSeveritySlugOK describes a response with status code 200, with default header values.

Update a specific severity
*/
type PatchV1SeveritiesSeveritySlugOK struct {
	Payload *models.SeverityEntity
}

// IsSuccess returns true when this patch v1 severities severity slug o k response has a 2xx status code
func (o *PatchV1SeveritiesSeveritySlugOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 severities severity slug o k response has a 3xx status code
func (o *PatchV1SeveritiesSeveritySlugOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 severities severity slug o k response has a 4xx status code
func (o *PatchV1SeveritiesSeveritySlugOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 severities severity slug o k response has a 5xx status code
func (o *PatchV1SeveritiesSeveritySlugOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 severities severity slug o k response a status code equal to that given
func (o *PatchV1SeveritiesSeveritySlugOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch v1 severities severity slug o k response
func (o *PatchV1SeveritiesSeveritySlugOK) Code() int {
	return 200
}

func (o *PatchV1SeveritiesSeveritySlugOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/severities/{severity_slug}][%d] patchV1SeveritiesSeveritySlugOK  %+v", 200, o.Payload)
}

func (o *PatchV1SeveritiesSeveritySlugOK) String() string {
	return fmt.Sprintf("[PATCH /v1/severities/{severity_slug}][%d] patchV1SeveritiesSeveritySlugOK  %+v", 200, o.Payload)
}

func (o *PatchV1SeveritiesSeveritySlugOK) GetPayload() *models.SeverityEntity {
	return o.Payload
}

func (o *PatchV1SeveritiesSeveritySlugOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SeverityEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
