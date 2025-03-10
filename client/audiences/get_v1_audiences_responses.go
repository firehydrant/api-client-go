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

// GetV1AudiencesReader is a Reader for the GetV1Audiences structure.
type GetV1AudiencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1AudiencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1AudiencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1AudiencesOK creates a GetV1AudiencesOK with default headers values
func NewGetV1AudiencesOK() *GetV1AudiencesOK {
	return &GetV1AudiencesOK{}
}

/*
GetV1AudiencesOK describes a response with status code 200, with default header values.

List all audiences
*/
type GetV1AudiencesOK struct {
	Payload *models.AudiencesEntitiesAudienceEntity
}

// IsSuccess returns true when this get v1 audiences o k response has a 2xx status code
func (o *GetV1AudiencesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 audiences o k response has a 3xx status code
func (o *GetV1AudiencesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 audiences o k response has a 4xx status code
func (o *GetV1AudiencesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 audiences o k response has a 5xx status code
func (o *GetV1AudiencesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 audiences o k response a status code equal to that given
func (o *GetV1AudiencesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1AudiencesOK) Error() string {
	return fmt.Sprintf("[GET /v1/audiences][%d] getV1AudiencesOK  %+v", 200, o.Payload)
}

func (o *GetV1AudiencesOK) String() string {
	return fmt.Sprintf("[GET /v1/audiences][%d] getV1AudiencesOK  %+v", 200, o.Payload)
}

func (o *GetV1AudiencesOK) GetPayload() *models.AudiencesEntitiesAudienceEntity {
	return o.Payload
}

func (o *GetV1AudiencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AudiencesEntitiesAudienceEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
