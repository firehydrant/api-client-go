// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchV1IntegrationsConnectionsOpsgenieIDRefreshReader is a Reader for the PatchV1IntegrationsConnectionsOpsgenieIDRefresh structure.
type PatchV1IntegrationsConnectionsOpsgenieIDRefreshReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1IntegrationsConnectionsOpsgenieIDRefreshReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1IntegrationsConnectionsOpsgenieIDRefreshOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1IntegrationsConnectionsOpsgenieIDRefreshOK creates a PatchV1IntegrationsConnectionsOpsgenieIDRefreshOK with default headers values
func NewPatchV1IntegrationsConnectionsOpsgenieIDRefreshOK() *PatchV1IntegrationsConnectionsOpsgenieIDRefreshOK {
	return &PatchV1IntegrationsConnectionsOpsgenieIDRefreshOK{}
}

/* PatchV1IntegrationsConnectionsOpsgenieIDRefreshOK describes a response with status code 200, with default header values.

patched Refresh
*/
type PatchV1IntegrationsConnectionsOpsgenieIDRefreshOK struct {
}

func (o *PatchV1IntegrationsConnectionsOpsgenieIDRefreshOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/integrations/connections/opsgenie/{id}/refresh][%d] patchV1IntegrationsConnectionsOpsgenieIdRefreshOK ", 200)
}

func (o *PatchV1IntegrationsConnectionsOpsgenieIDRefreshOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}