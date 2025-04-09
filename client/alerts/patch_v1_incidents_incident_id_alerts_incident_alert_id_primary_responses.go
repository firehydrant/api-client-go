// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryReader is a Reader for the PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary structure.
type PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK creates a PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK with default headers values
func NewPatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK() *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK {
	return &PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK{}
}

/*
PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK describes a response with status code 200, with default header values.

Setting an alert as primary will overwrite milestone times in the FireHydrant incident with times included in the primary alert. Services attached to the primary alert will also be automatically added to the incident.
*/
type PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK struct {
	Payload *models.IncidentsAlertEntity
}

// IsSuccess returns true when this patch v1 incidents incident Id alerts incident alert Id primary o k response has a 2xx status code
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 incidents incident Id alerts incident alert Id primary o k response has a 3xx status code
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 incidents incident Id alerts incident alert Id primary o k response has a 4xx status code
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 incidents incident Id alerts incident alert Id primary o k response has a 5xx status code
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 incidents incident Id alerts incident alert Id primary o k response a status code equal to that given
func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/incidents/{incident_id}/alerts/{incident_alert_id}/primary][%d] patchV1IncidentsIncidentIdAlertsIncidentAlertIdPrimaryOK  %+v", 200, o.Payload)
}

func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK) String() string {
	return fmt.Sprintf("[PATCH /v1/incidents/{incident_id}/alerts/{incident_alert_id}/primary][%d] patchV1IncidentsIncidentIdAlertsIncidentAlertIdPrimaryOK  %+v", 200, o.Payload)
}

func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK) GetPayload() *models.IncidentsAlertEntity {
	return o.Payload
}

func (o *PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncidentsAlertEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
