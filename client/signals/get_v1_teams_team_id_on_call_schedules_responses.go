// Code generated by go-swagger; DO NOT EDIT.

package signals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1TeamsTeamIDOnCallSchedulesReader is a Reader for the GetV1TeamsTeamIDOnCallSchedules structure.
type GetV1TeamsTeamIDOnCallSchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1TeamsTeamIDOnCallSchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1TeamsTeamIDOnCallSchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1TeamsTeamIDOnCallSchedulesOK creates a GetV1TeamsTeamIDOnCallSchedulesOK with default headers values
func NewGetV1TeamsTeamIDOnCallSchedulesOK() *GetV1TeamsTeamIDOnCallSchedulesOK {
	return &GetV1TeamsTeamIDOnCallSchedulesOK{}
}

/*
GetV1TeamsTeamIDOnCallSchedulesOK describes a response with status code 200, with default header values.

List all Signals on-call schedules for a team.
*/
type GetV1TeamsTeamIDOnCallSchedulesOK struct {
}

// IsSuccess returns true when this get v1 teams team Id on call schedules o k response has a 2xx status code
func (o *GetV1TeamsTeamIDOnCallSchedulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 teams team Id on call schedules o k response has a 3xx status code
func (o *GetV1TeamsTeamIDOnCallSchedulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 teams team Id on call schedules o k response has a 4xx status code
func (o *GetV1TeamsTeamIDOnCallSchedulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 teams team Id on call schedules o k response has a 5xx status code
func (o *GetV1TeamsTeamIDOnCallSchedulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 teams team Id on call schedules o k response a status code equal to that given
func (o *GetV1TeamsTeamIDOnCallSchedulesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1TeamsTeamIDOnCallSchedulesOK) Error() string {
	return fmt.Sprintf("[GET /v1/teams/{team_id}/on_call_schedules][%d] getV1TeamsTeamIdOnCallSchedulesOK ", 200)
}

func (o *GetV1TeamsTeamIDOnCallSchedulesOK) String() string {
	return fmt.Sprintf("[GET /v1/teams/{team_id}/on_call_schedules][%d] getV1TeamsTeamIdOnCallSchedulesOK ", 200)
}

func (o *GetV1TeamsTeamIDOnCallSchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
