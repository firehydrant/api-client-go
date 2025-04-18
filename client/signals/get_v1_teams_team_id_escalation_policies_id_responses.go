// Code generated by go-swagger; DO NOT EDIT.

package signals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1TeamsTeamIDEscalationPoliciesIDReader is a Reader for the GetV1TeamsTeamIDEscalationPoliciesID structure.
type GetV1TeamsTeamIDEscalationPoliciesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1TeamsTeamIDEscalationPoliciesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1TeamsTeamIDEscalationPoliciesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1TeamsTeamIDEscalationPoliciesIDOK creates a GetV1TeamsTeamIDEscalationPoliciesIDOK with default headers values
func NewGetV1TeamsTeamIDEscalationPoliciesIDOK() *GetV1TeamsTeamIDEscalationPoliciesIDOK {
	return &GetV1TeamsTeamIDEscalationPoliciesIDOK{}
}

/*
GetV1TeamsTeamIDEscalationPoliciesIDOK describes a response with status code 200, with default header values.

Get a Signals escalation policy by ID
*/
type GetV1TeamsTeamIDEscalationPoliciesIDOK struct {
}

// IsSuccess returns true when this get v1 teams team Id escalation policies Id o k response has a 2xx status code
func (o *GetV1TeamsTeamIDEscalationPoliciesIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 teams team Id escalation policies Id o k response has a 3xx status code
func (o *GetV1TeamsTeamIDEscalationPoliciesIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 teams team Id escalation policies Id o k response has a 4xx status code
func (o *GetV1TeamsTeamIDEscalationPoliciesIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 teams team Id escalation policies Id o k response has a 5xx status code
func (o *GetV1TeamsTeamIDEscalationPoliciesIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 teams team Id escalation policies Id o k response a status code equal to that given
func (o *GetV1TeamsTeamIDEscalationPoliciesIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1TeamsTeamIDEscalationPoliciesIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/teams/{team_id}/escalation_policies/{id}][%d] getV1TeamsTeamIdEscalationPoliciesIdOK ", 200)
}

func (o *GetV1TeamsTeamIDEscalationPoliciesIDOK) String() string {
	return fmt.Sprintf("[GET /v1/teams/{team_id}/escalation_policies/{id}][%d] getV1TeamsTeamIdEscalationPoliciesIdOK ", 200)
}

func (o *GetV1TeamsTeamIDEscalationPoliciesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
