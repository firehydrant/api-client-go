// Code generated by go-swagger; DO NOT EDIT.

package audiences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1AudiencesAudienceIDDetailsReader is a Reader for the PostV1AudiencesAudienceIDDetails structure.
type PostV1AudiencesAudienceIDDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1AudiencesAudienceIDDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1AudiencesAudienceIDDetailsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1AudiencesAudienceIDDetailsCreated creates a PostV1AudiencesAudienceIDDetailsCreated with default headers values
func NewPostV1AudiencesAudienceIDDetailsCreated() *PostV1AudiencesAudienceIDDetailsCreated {
	return &PostV1AudiencesAudienceIDDetailsCreated{}
}

/*
PostV1AudiencesAudienceIDDetailsCreated describes a response with status code 201, with default header values.

Create a new incident detail for an audience
*/
type PostV1AudiencesAudienceIDDetailsCreated struct {
	Payload *models.AudiencesEntitiesDetailEntity
}

// IsSuccess returns true when this post v1 audiences audience Id details created response has a 2xx status code
func (o *PostV1AudiencesAudienceIDDetailsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 audiences audience Id details created response has a 3xx status code
func (o *PostV1AudiencesAudienceIDDetailsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 audiences audience Id details created response has a 4xx status code
func (o *PostV1AudiencesAudienceIDDetailsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 audiences audience Id details created response has a 5xx status code
func (o *PostV1AudiencesAudienceIDDetailsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 audiences audience Id details created response a status code equal to that given
func (o *PostV1AudiencesAudienceIDDetailsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1AudiencesAudienceIDDetailsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/audiences/{audience_id}/details][%d] postV1AudiencesAudienceIdDetailsCreated  %+v", 201, o.Payload)
}

func (o *PostV1AudiencesAudienceIDDetailsCreated) String() string {
	return fmt.Sprintf("[POST /v1/audiences/{audience_id}/details][%d] postV1AudiencesAudienceIdDetailsCreated  %+v", 201, o.Payload)
}

func (o *PostV1AudiencesAudienceIDDetailsCreated) GetPayload() *models.AudiencesEntitiesDetailEntity {
	return o.Payload
}

func (o *PostV1AudiencesAudienceIDDetailsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AudiencesEntitiesDetailEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
