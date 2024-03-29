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

// PatchV1IntegrationsStatuspageConnectionsConnectionIDReader is a Reader for the PatchV1IntegrationsStatuspageConnectionsConnectionID structure.
type PatchV1IntegrationsStatuspageConnectionsConnectionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1IntegrationsStatuspageConnectionsConnectionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1IntegrationsStatuspageConnectionsConnectionIDOK creates a PatchV1IntegrationsStatuspageConnectionsConnectionIDOK with default headers values
func NewPatchV1IntegrationsStatuspageConnectionsConnectionIDOK() *PatchV1IntegrationsStatuspageConnectionsConnectionIDOK {
	return &PatchV1IntegrationsStatuspageConnectionsConnectionIDOK{}
}

/*
PatchV1IntegrationsStatuspageConnectionsConnectionIDOK describes a response with status code 200, with default header values.

Update the given Statuspage integration connection.
*/
type PatchV1IntegrationsStatuspageConnectionsConnectionIDOK struct {
	Payload *models.IntegrationsStatuspageConnectionEntity
}

// IsSuccess returns true when this patch v1 integrations statuspage connections connection Id o k response has a 2xx status code
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 integrations statuspage connections connection Id o k response has a 3xx status code
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 integrations statuspage connections connection Id o k response has a 4xx status code
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 integrations statuspage connections connection Id o k response has a 5xx status code
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 integrations statuspage connections connection Id o k response a status code equal to that given
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/integrations/statuspage/connections/{connection_id}][%d] patchV1IntegrationsStatuspageConnectionsConnectionIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/integrations/statuspage/connections/{connection_id}][%d] patchV1IntegrationsStatuspageConnectionsConnectionIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDOK) GetPayload() *models.IntegrationsStatuspageConnectionEntity {
	return o.Payload
}

func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntegrationsStatuspageConnectionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
