// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDReader is a Reader for the DeleteV1TeamsTeamIDOnCallSchedulesScheduleID structure.
type DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/teams/{team_id}/on_call_schedules/{schedule_id}] deleteV1TeamsTeamIdOnCallSchedulesScheduleId", response, response.Code())
	}
}

// NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent creates a DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent with default headers values
func NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent() *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent {
	return &DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent{}
}

/*
DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent describes a response with status code 204, with default header values.

Delete a Signals on-call schedule by ID
*/
type DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent struct {
}

// IsSuccess returns true when this delete v1 teams team Id on call schedules schedule Id no content response has a 2xx status code
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 teams team Id on call schedules schedule Id no content response has a 3xx status code
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 teams team Id on call schedules schedule Id no content response has a 4xx status code
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 teams team Id on call schedules schedule Id no content response has a 5xx status code
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 teams team Id on call schedules schedule Id no content response a status code equal to that given
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete v1 teams team Id on call schedules schedule Id no content response
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent) Code() int {
	return 204
}

func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/teams/{team_id}/on_call_schedules/{schedule_id}][%d] deleteV1TeamsTeamIdOnCallSchedulesScheduleIdNoContent ", 204)
}

func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/teams/{team_id}/on_call_schedules/{schedule_id}][%d] deleteV1TeamsTeamIdOnCallSchedulesScheduleIdNoContent ", 204)
}

func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
