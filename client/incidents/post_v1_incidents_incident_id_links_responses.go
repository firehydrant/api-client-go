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

// PostV1IncidentsIncidentIDLinksReader is a Reader for the PostV1IncidentsIncidentIDLinks structure.
type PostV1IncidentsIncidentIDLinksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1IncidentsIncidentIDLinksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1IncidentsIncidentIDLinksCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1IncidentsIncidentIDLinksCreated creates a PostV1IncidentsIncidentIDLinksCreated with default headers values
func NewPostV1IncidentsIncidentIDLinksCreated() *PostV1IncidentsIncidentIDLinksCreated {
	return &PostV1IncidentsIncidentIDLinksCreated{}
}

/*
PostV1IncidentsIncidentIDLinksCreated describes a response with status code 201, with default header values.

Allows adding adhoc links to an incident as an attachment
*/
type PostV1IncidentsIncidentIDLinksCreated struct {
	Payload *models.AttachmentsLinkEntity
}

// IsSuccess returns true when this post v1 incidents incident Id links created response has a 2xx status code
func (o *PostV1IncidentsIncidentIDLinksCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 incidents incident Id links created response has a 3xx status code
func (o *PostV1IncidentsIncidentIDLinksCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 incidents incident Id links created response has a 4xx status code
func (o *PostV1IncidentsIncidentIDLinksCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 incidents incident Id links created response has a 5xx status code
func (o *PostV1IncidentsIncidentIDLinksCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 incidents incident Id links created response a status code equal to that given
func (o *PostV1IncidentsIncidentIDLinksCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1IncidentsIncidentIDLinksCreated) Error() string {
	return fmt.Sprintf("[POST /v1/incidents/{incident_id}/links][%d] postV1IncidentsIncidentIdLinksCreated  %+v", 201, o.Payload)
}

func (o *PostV1IncidentsIncidentIDLinksCreated) String() string {
	return fmt.Sprintf("[POST /v1/incidents/{incident_id}/links][%d] postV1IncidentsIncidentIdLinksCreated  %+v", 201, o.Payload)
}

func (o *PostV1IncidentsIncidentIDLinksCreated) GetPayload() *models.AttachmentsLinkEntity {
	return o.Payload
}

func (o *PostV1IncidentsIncidentIDLinksCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AttachmentsLinkEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
