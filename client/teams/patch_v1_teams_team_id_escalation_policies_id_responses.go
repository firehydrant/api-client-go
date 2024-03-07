// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PatchV1TeamsTeamIDEscalationPoliciesIDReader is a Reader for the PatchV1TeamsTeamIDEscalationPoliciesID structure.
type PatchV1TeamsTeamIDEscalationPoliciesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1TeamsTeamIDEscalationPoliciesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1TeamsTeamIDEscalationPoliciesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1TeamsTeamIDEscalationPoliciesIDOK creates a PatchV1TeamsTeamIDEscalationPoliciesIDOK with default headers values
func NewPatchV1TeamsTeamIDEscalationPoliciesIDOK() *PatchV1TeamsTeamIDEscalationPoliciesIDOK {
	return &PatchV1TeamsTeamIDEscalationPoliciesIDOK{}
}

/*
PatchV1TeamsTeamIDEscalationPoliciesIDOK describes a response with status code 200, with default header values.

Update a Signals escalation policy by ID
*/
type PatchV1TeamsTeamIDEscalationPoliciesIDOK struct {
}

// IsSuccess returns true when this patch v1 teams team Id escalation policies Id o k response has a 2xx status code
func (o *PatchV1TeamsTeamIDEscalationPoliciesIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 teams team Id escalation policies Id o k response has a 3xx status code
func (o *PatchV1TeamsTeamIDEscalationPoliciesIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 teams team Id escalation policies Id o k response has a 4xx status code
func (o *PatchV1TeamsTeamIDEscalationPoliciesIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 teams team Id escalation policies Id o k response has a 5xx status code
func (o *PatchV1TeamsTeamIDEscalationPoliciesIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 teams team Id escalation policies Id o k response a status code equal to that given
func (o *PatchV1TeamsTeamIDEscalationPoliciesIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchV1TeamsTeamIDEscalationPoliciesIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/teams/{team_id}/escalation_policies/{id}][%d] patchV1TeamsTeamIdEscalationPoliciesIdOK ", 200)
}

func (o *PatchV1TeamsTeamIDEscalationPoliciesIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/teams/{team_id}/escalation_policies/{id}][%d] patchV1TeamsTeamIdEscalationPoliciesIdOK ", 200)
}

func (o *PatchV1TeamsTeamIDEscalationPoliciesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
