// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDReader is a Reader for the GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID structure.
type GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/teams/{team_id}/on_call_schedules/{schedule_id}/shifts/{id}] getV1TeamsTeamIdOnCallSchedulesScheduleIdShiftsId", response, response.Code())
	}
}

// NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK creates a GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK with default headers values
func NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK() *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK {
	return &GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK{}
}

/*
GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK describes a response with status code 200, with default header values.

Get a Signals on-call shift by ID
*/
type GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK struct {
}

// IsSuccess returns true when this get v1 teams team Id on call schedules schedule Id shifts Id o k response has a 2xx status code
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 teams team Id on call schedules schedule Id shifts Id o k response has a 3xx status code
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 teams team Id on call schedules schedule Id shifts Id o k response has a 4xx status code
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 teams team Id on call schedules schedule Id shifts Id o k response has a 5xx status code
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 teams team Id on call schedules schedule Id shifts Id o k response a status code equal to that given
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 teams team Id on call schedules schedule Id shifts Id o k response
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK) Code() int {
	return 200
}

func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/teams/{team_id}/on_call_schedules/{schedule_id}/shifts/{id}][%d] getV1TeamsTeamIdOnCallSchedulesScheduleIdShiftsIdOK ", 200)
}

func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK) String() string {
	return fmt.Sprintf("[GET /v1/teams/{team_id}/on_call_schedules/{schedule_id}/shifts/{id}][%d] getV1TeamsTeamIdOnCallSchedulesScheduleIdShiftsIdOK ", 200)
}

func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
