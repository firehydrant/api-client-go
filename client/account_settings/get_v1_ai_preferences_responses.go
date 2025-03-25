// Code generated by go-swagger; DO NOT EDIT.

package account_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1AiPreferencesReader is a Reader for the GetV1AiPreferences structure.
type GetV1AiPreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1AiPreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1AiPreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1AiPreferencesOK creates a GetV1AiPreferencesOK with default headers values
func NewGetV1AiPreferencesOK() *GetV1AiPreferencesOK {
	return &GetV1AiPreferencesOK{}
}

/*
GetV1AiPreferencesOK describes a response with status code 200, with default header values.

Retrieves the current AI preferences
*/
type GetV1AiPreferencesOK struct {
	Payload *models.AIEntitiesPreferencesEntity
}

// IsSuccess returns true when this get v1 ai preferences o k response has a 2xx status code
func (o *GetV1AiPreferencesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 ai preferences o k response has a 3xx status code
func (o *GetV1AiPreferencesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 ai preferences o k response has a 4xx status code
func (o *GetV1AiPreferencesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 ai preferences o k response has a 5xx status code
func (o *GetV1AiPreferencesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 ai preferences o k response a status code equal to that given
func (o *GetV1AiPreferencesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1AiPreferencesOK) Error() string {
	return fmt.Sprintf("[GET /v1/ai/preferences][%d] getV1AiPreferencesOK  %+v", 200, o.Payload)
}

func (o *GetV1AiPreferencesOK) String() string {
	return fmt.Sprintf("[GET /v1/ai/preferences][%d] getV1AiPreferencesOK  %+v", 200, o.Payload)
}

func (o *GetV1AiPreferencesOK) GetPayload() *models.AIEntitiesPreferencesEntity {
	return o.Payload
}

func (o *GetV1AiPreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AIEntitiesPreferencesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
