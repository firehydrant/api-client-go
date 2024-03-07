// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1TeamsTeamIDOnCallSchedulesScheduleIDReader is a Reader for the GetV1TeamsTeamIDOnCallSchedulesScheduleID structure.
type GetV1TeamsTeamIDOnCallSchedulesScheduleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDOK creates a GetV1TeamsTeamIDOnCallSchedulesScheduleIDOK with default headers values
func NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDOK() *GetV1TeamsTeamIDOnCallSchedulesScheduleIDOK {
	return &GetV1TeamsTeamIDOnCallSchedulesScheduleIDOK{}
}

/*
GetV1TeamsTeamIDOnCallSchedulesScheduleIDOK describes a response with status code 200, with default header values.

Get a Signals on-call schedule by ID
*/
type GetV1TeamsTeamIDOnCallSchedulesScheduleIDOK struct {
}

// IsSuccess returns true when this get v1 teams team Id on call schedules schedule Id o k response has a 2xx status code
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 teams team Id on call schedules schedule Id o k response has a 3xx status code
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 teams team Id on call schedules schedule Id o k response has a 4xx status code
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 teams team Id on call schedules schedule Id o k response has a 5xx status code
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 teams team Id on call schedules schedule Id o k response a status code equal to that given
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/teams/{team_id}/on_call_schedules/{schedule_id}][%d] getV1TeamsTeamIdOnCallSchedulesScheduleIdOK ", 200)
}

func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDOK) String() string {
	return fmt.Sprintf("[GET /v1/teams/{team_id}/on_call_schedules/{schedule_id}][%d] getV1TeamsTeamIdOnCallSchedulesScheduleIdOK ", 200)
}

func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
