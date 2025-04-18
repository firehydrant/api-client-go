// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1ScheduledMaintenancesScheduledMaintenanceIDReader is a Reader for the DeleteV1ScheduledMaintenancesScheduledMaintenanceID structure.
type DeleteV1ScheduledMaintenancesScheduledMaintenanceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent creates a DeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent with default headers values
func NewDeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent() *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent {
	return &DeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent{}
}

/*
DeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent describes a response with status code 204, with default header values.

Delete a scheduled maintenance event, preventing it from taking place.
*/
type DeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent struct {
}

// IsSuccess returns true when this delete v1 scheduled maintenances scheduled maintenance Id no content response has a 2xx status code
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 scheduled maintenances scheduled maintenance Id no content response has a 3xx status code
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 scheduled maintenances scheduled maintenance Id no content response has a 4xx status code
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 scheduled maintenances scheduled maintenance Id no content response has a 5xx status code
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 scheduled maintenances scheduled maintenance Id no content response a status code equal to that given
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/scheduled_maintenances/{scheduled_maintenance_id}][%d] deleteV1ScheduledMaintenancesScheduledMaintenanceIdNoContent ", 204)
}

func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/scheduled_maintenances/{scheduled_maintenance_id}][%d] deleteV1ScheduledMaintenancesScheduledMaintenanceIdNoContent ", 204)
}

func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
