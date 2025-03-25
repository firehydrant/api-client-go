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

// GetV1SeveritiesReader is a Reader for the GetV1Severities structure.
type GetV1SeveritiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1SeveritiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1SeveritiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1SeveritiesOK creates a GetV1SeveritiesOK with default headers values
func NewGetV1SeveritiesOK() *GetV1SeveritiesOK {
	return &GetV1SeveritiesOK{}
}

/*
GetV1SeveritiesOK describes a response with status code 200, with default header values.

Lists severities
*/
type GetV1SeveritiesOK struct {
	Payload *models.SeverityEntityPaginated
}

// IsSuccess returns true when this get v1 severities o k response has a 2xx status code
func (o *GetV1SeveritiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 severities o k response has a 3xx status code
func (o *GetV1SeveritiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 severities o k response has a 4xx status code
func (o *GetV1SeveritiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 severities o k response has a 5xx status code
func (o *GetV1SeveritiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 severities o k response a status code equal to that given
func (o *GetV1SeveritiesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1SeveritiesOK) Error() string {
	return fmt.Sprintf("[GET /v1/severities][%d] getV1SeveritiesOK  %+v", 200, o.Payload)
}

func (o *GetV1SeveritiesOK) String() string {
	return fmt.Sprintf("[GET /v1/severities][%d] getV1SeveritiesOK  %+v", 200, o.Payload)
}

func (o *GetV1SeveritiesOK) GetPayload() *models.SeverityEntityPaginated {
	return o.Payload
}

func (o *GetV1SeveritiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SeverityEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
