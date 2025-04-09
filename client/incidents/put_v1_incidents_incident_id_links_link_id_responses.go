// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutV1IncidentsIncidentIDLinksLinkIDReader is a Reader for the PutV1IncidentsIncidentIDLinksLinkID structure.
type PutV1IncidentsIncidentIDLinksLinkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1IncidentsIncidentIDLinksLinkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutV1IncidentsIncidentIDLinksLinkIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutV1IncidentsIncidentIDLinksLinkIDOK creates a PutV1IncidentsIncidentIDLinksLinkIDOK with default headers values
func NewPutV1IncidentsIncidentIDLinksLinkIDOK() *PutV1IncidentsIncidentIDLinksLinkIDOK {
	return &PutV1IncidentsIncidentIDLinksLinkIDOK{}
}

/*
PutV1IncidentsIncidentIDLinksLinkIDOK describes a response with status code 200, with default header values.

Update the external incident link attributes
*/
type PutV1IncidentsIncidentIDLinksLinkIDOK struct {
}

// IsSuccess returns true when this put v1 incidents incident Id links link Id o k response has a 2xx status code
func (o *PutV1IncidentsIncidentIDLinksLinkIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put v1 incidents incident Id links link Id o k response has a 3xx status code
func (o *PutV1IncidentsIncidentIDLinksLinkIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 incidents incident Id links link Id o k response has a 4xx status code
func (o *PutV1IncidentsIncidentIDLinksLinkIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put v1 incidents incident Id links link Id o k response has a 5xx status code
func (o *PutV1IncidentsIncidentIDLinksLinkIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 incidents incident Id links link Id o k response a status code equal to that given
func (o *PutV1IncidentsIncidentIDLinksLinkIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutV1IncidentsIncidentIDLinksLinkIDOK) Error() string {
	return fmt.Sprintf("[PUT /v1/incidents/{incident_id}/links/{link_id}][%d] putV1IncidentsIncidentIdLinksLinkIdOK ", 200)
}

func (o *PutV1IncidentsIncidentIDLinksLinkIDOK) String() string {
	return fmt.Sprintf("[PUT /v1/incidents/{incident_id}/links/{link_id}][%d] putV1IncidentsIncidentIdLinksLinkIdOK ", 200)
}

func (o *PutV1IncidentsIncidentIDLinksLinkIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
