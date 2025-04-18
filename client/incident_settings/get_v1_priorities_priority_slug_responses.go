// Code generated by go-swagger; DO NOT EDIT.

package incident_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1PrioritiesPrioritySlugReader is a Reader for the GetV1PrioritiesPrioritySlug structure.
type GetV1PrioritiesPrioritySlugReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1PrioritiesPrioritySlugReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1PrioritiesPrioritySlugOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1PrioritiesPrioritySlugOK creates a GetV1PrioritiesPrioritySlugOK with default headers values
func NewGetV1PrioritiesPrioritySlugOK() *GetV1PrioritiesPrioritySlugOK {
	return &GetV1PrioritiesPrioritySlugOK{}
}

/*
GetV1PrioritiesPrioritySlugOK describes a response with status code 200, with default header values.

Retrieve a specific priority
*/
type GetV1PrioritiesPrioritySlugOK struct {
	Payload *models.PriorityEntity
}

// IsSuccess returns true when this get v1 priorities priority slug o k response has a 2xx status code
func (o *GetV1PrioritiesPrioritySlugOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 priorities priority slug o k response has a 3xx status code
func (o *GetV1PrioritiesPrioritySlugOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 priorities priority slug o k response has a 4xx status code
func (o *GetV1PrioritiesPrioritySlugOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 priorities priority slug o k response has a 5xx status code
func (o *GetV1PrioritiesPrioritySlugOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 priorities priority slug o k response a status code equal to that given
func (o *GetV1PrioritiesPrioritySlugOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1PrioritiesPrioritySlugOK) Error() string {
	return fmt.Sprintf("[GET /v1/priorities/{priority_slug}][%d] getV1PrioritiesPrioritySlugOK  %+v", 200, o.Payload)
}

func (o *GetV1PrioritiesPrioritySlugOK) String() string {
	return fmt.Sprintf("[GET /v1/priorities/{priority_slug}][%d] getV1PrioritiesPrioritySlugOK  %+v", 200, o.Payload)
}

func (o *GetV1PrioritiesPrioritySlugOK) GetPayload() *models.PriorityEntity {
	return o.Payload
}

func (o *GetV1PrioritiesPrioritySlugOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PriorityEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
