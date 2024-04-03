// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostV1TeamsTeamIDEscalationPoliciesReader is a Reader for the PostV1TeamsTeamIDEscalationPolicies structure.
type PostV1TeamsTeamIDEscalationPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1TeamsTeamIDEscalationPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1TeamsTeamIDEscalationPoliciesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/teams/{team_id}/escalation_policies] postV1TeamsTeamIdEscalationPolicies", response, response.Code())
	}
}

// NewPostV1TeamsTeamIDEscalationPoliciesCreated creates a PostV1TeamsTeamIDEscalationPoliciesCreated with default headers values
func NewPostV1TeamsTeamIDEscalationPoliciesCreated() *PostV1TeamsTeamIDEscalationPoliciesCreated {
	return &PostV1TeamsTeamIDEscalationPoliciesCreated{}
}

/*
PostV1TeamsTeamIDEscalationPoliciesCreated describes a response with status code 201, with default header values.

Create a Signals escalation policy for a team.
*/
type PostV1TeamsTeamIDEscalationPoliciesCreated struct {
}

// IsSuccess returns true when this post v1 teams team Id escalation policies created response has a 2xx status code
func (o *PostV1TeamsTeamIDEscalationPoliciesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 teams team Id escalation policies created response has a 3xx status code
func (o *PostV1TeamsTeamIDEscalationPoliciesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 teams team Id escalation policies created response has a 4xx status code
func (o *PostV1TeamsTeamIDEscalationPoliciesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 teams team Id escalation policies created response has a 5xx status code
func (o *PostV1TeamsTeamIDEscalationPoliciesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 teams team Id escalation policies created response a status code equal to that given
func (o *PostV1TeamsTeamIDEscalationPoliciesCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post v1 teams team Id escalation policies created response
func (o *PostV1TeamsTeamIDEscalationPoliciesCreated) Code() int {
	return 201
}

func (o *PostV1TeamsTeamIDEscalationPoliciesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/teams/{team_id}/escalation_policies][%d] postV1TeamsTeamIdEscalationPoliciesCreated ", 201)
}

func (o *PostV1TeamsTeamIDEscalationPoliciesCreated) String() string {
	return fmt.Sprintf("[POST /v1/teams/{team_id}/escalation_policies][%d] postV1TeamsTeamIdEscalationPoliciesCreated ", 201)
}

func (o *PostV1TeamsTeamIDEscalationPoliciesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
