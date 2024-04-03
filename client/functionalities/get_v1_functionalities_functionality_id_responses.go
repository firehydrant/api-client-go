// Code generated by go-swagger; DO NOT EDIT.

package functionalities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1FunctionalitiesFunctionalityIDReader is a Reader for the GetV1FunctionalitiesFunctionalityID structure.
type GetV1FunctionalitiesFunctionalityIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1FunctionalitiesFunctionalityIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1FunctionalitiesFunctionalityIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/functionalities/{functionality_id}] getV1FunctionalitiesFunctionalityId", response, response.Code())
	}
}

// NewGetV1FunctionalitiesFunctionalityIDOK creates a GetV1FunctionalitiesFunctionalityIDOK with default headers values
func NewGetV1FunctionalitiesFunctionalityIDOK() *GetV1FunctionalitiesFunctionalityIDOK {
	return &GetV1FunctionalitiesFunctionalityIDOK{}
}

/*
GetV1FunctionalitiesFunctionalityIDOK describes a response with status code 200, with default header values.

Retrieves a single functionality by ID
*/
type GetV1FunctionalitiesFunctionalityIDOK struct {
	Payload *models.FunctionalityEntity
}

// IsSuccess returns true when this get v1 functionalities functionality Id o k response has a 2xx status code
func (o *GetV1FunctionalitiesFunctionalityIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 functionalities functionality Id o k response has a 3xx status code
func (o *GetV1FunctionalitiesFunctionalityIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 functionalities functionality Id o k response has a 4xx status code
func (o *GetV1FunctionalitiesFunctionalityIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 functionalities functionality Id o k response has a 5xx status code
func (o *GetV1FunctionalitiesFunctionalityIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 functionalities functionality Id o k response a status code equal to that given
func (o *GetV1FunctionalitiesFunctionalityIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 functionalities functionality Id o k response
func (o *GetV1FunctionalitiesFunctionalityIDOK) Code() int {
	return 200
}

func (o *GetV1FunctionalitiesFunctionalityIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/functionalities/{functionality_id}][%d] getV1FunctionalitiesFunctionalityIdOK  %+v", 200, o.Payload)
}

func (o *GetV1FunctionalitiesFunctionalityIDOK) String() string {
	return fmt.Sprintf("[GET /v1/functionalities/{functionality_id}][%d] getV1FunctionalitiesFunctionalityIdOK  %+v", 200, o.Payload)
}

func (o *GetV1FunctionalitiesFunctionalityIDOK) GetPayload() *models.FunctionalityEntity {
	return o.Payload
}

func (o *GetV1FunctionalitiesFunctionalityIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FunctionalityEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
