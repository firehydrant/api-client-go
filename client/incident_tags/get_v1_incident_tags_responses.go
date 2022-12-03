// Code generated by go-swagger; DO NOT EDIT.

package incident_tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1IncidentTagsReader is a Reader for the GetV1IncidentTags structure.
type GetV1IncidentTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IncidentTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IncidentTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1IncidentTagsOK creates a GetV1IncidentTagsOK with default headers values
func NewGetV1IncidentTagsOK() *GetV1IncidentTagsOK {
	return &GetV1IncidentTagsOK{}
}

/*
GetV1IncidentTagsOK describes a response with status code 200, with default header values.

List all of the incident tags in the organization
*/
type GetV1IncidentTagsOK struct {
	Payload *models.TagEntityPaginated
}

// IsSuccess returns true when this get v1 incident tags o k response has a 2xx status code
func (o *GetV1IncidentTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 incident tags o k response has a 3xx status code
func (o *GetV1IncidentTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 incident tags o k response has a 4xx status code
func (o *GetV1IncidentTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 incident tags o k response has a 5xx status code
func (o *GetV1IncidentTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 incident tags o k response a status code equal to that given
func (o *GetV1IncidentTagsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1IncidentTagsOK) Error() string {
	return fmt.Sprintf("[GET /v1/incident_tags][%d] getV1IncidentTagsOK  %+v", 200, o.Payload)
}

func (o *GetV1IncidentTagsOK) String() string {
	return fmt.Sprintf("[GET /v1/incident_tags][%d] getV1IncidentTagsOK  %+v", 200, o.Payload)
}

func (o *GetV1IncidentTagsOK) GetPayload() *models.TagEntityPaginated {
	return o.Payload
}

func (o *GetV1IncidentTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
