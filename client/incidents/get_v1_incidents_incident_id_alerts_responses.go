// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1IncidentsIncidentIDAlertsReader is a Reader for the GetV1IncidentsIncidentIDAlerts structure.
type GetV1IncidentsIncidentIDAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IncidentsIncidentIDAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IncidentsIncidentIDAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1IncidentsIncidentIDAlertsOK creates a GetV1IncidentsIncidentIDAlertsOK with default headers values
func NewGetV1IncidentsIncidentIDAlertsOK() *GetV1IncidentsIncidentIDAlertsOK {
	return &GetV1IncidentsIncidentIDAlertsOK{}
}

/*
GetV1IncidentsIncidentIDAlertsOK describes a response with status code 200, with default header values.

List alerts that have been attached to an incident
*/
type GetV1IncidentsIncidentIDAlertsOK struct {
	Payload *models.AlertEntityPaginated
}

// IsSuccess returns true when this get v1 incidents incident Id alerts o k response has a 2xx status code
func (o *GetV1IncidentsIncidentIDAlertsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 incidents incident Id alerts o k response has a 3xx status code
func (o *GetV1IncidentsIncidentIDAlertsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 incidents incident Id alerts o k response has a 4xx status code
func (o *GetV1IncidentsIncidentIDAlertsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 incidents incident Id alerts o k response has a 5xx status code
func (o *GetV1IncidentsIncidentIDAlertsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 incidents incident Id alerts o k response a status code equal to that given
func (o *GetV1IncidentsIncidentIDAlertsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1IncidentsIncidentIDAlertsOK) Error() string {
	return fmt.Sprintf("[GET /v1/incidents/{incident_id}/alerts][%d] getV1IncidentsIncidentIdAlertsOK  %+v", 200, o.Payload)
}

func (o *GetV1IncidentsIncidentIDAlertsOK) String() string {
	return fmt.Sprintf("[GET /v1/incidents/{incident_id}/alerts][%d] getV1IncidentsIncidentIdAlertsOK  %+v", 200, o.Payload)
}

func (o *GetV1IncidentsIncidentIDAlertsOK) GetPayload() *models.AlertEntityPaginated {
	return o.Payload
}

func (o *GetV1IncidentsIncidentIDAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
