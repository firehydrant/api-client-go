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

// GetV1FunctionalitiesFunctionalityIDServicesReader is a Reader for the GetV1FunctionalitiesFunctionalityIDServices structure.
type GetV1FunctionalitiesFunctionalityIDServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1FunctionalitiesFunctionalityIDServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1FunctionalitiesFunctionalityIDServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1FunctionalitiesFunctionalityIDServicesOK creates a GetV1FunctionalitiesFunctionalityIDServicesOK with default headers values
func NewGetV1FunctionalitiesFunctionalityIDServicesOK() *GetV1FunctionalitiesFunctionalityIDServicesOK {
	return &GetV1FunctionalitiesFunctionalityIDServicesOK{}
}

/*
GetV1FunctionalitiesFunctionalityIDServicesOK describes a response with status code 200, with default header values.

List services for a functionality
*/
type GetV1FunctionalitiesFunctionalityIDServicesOK struct {
	Payload *models.FunctionalityWithAllServicesEntity
}

// IsSuccess returns true when this get v1 functionalities functionality Id services o k response has a 2xx status code
func (o *GetV1FunctionalitiesFunctionalityIDServicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 functionalities functionality Id services o k response has a 3xx status code
func (o *GetV1FunctionalitiesFunctionalityIDServicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 functionalities functionality Id services o k response has a 4xx status code
func (o *GetV1FunctionalitiesFunctionalityIDServicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 functionalities functionality Id services o k response has a 5xx status code
func (o *GetV1FunctionalitiesFunctionalityIDServicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 functionalities functionality Id services o k response a status code equal to that given
func (o *GetV1FunctionalitiesFunctionalityIDServicesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1FunctionalitiesFunctionalityIDServicesOK) Error() string {
	return fmt.Sprintf("[GET /v1/functionalities/{functionality_id}/services][%d] getV1FunctionalitiesFunctionalityIdServicesOK  %+v", 200, o.Payload)
}

func (o *GetV1FunctionalitiesFunctionalityIDServicesOK) String() string {
	return fmt.Sprintf("[GET /v1/functionalities/{functionality_id}/services][%d] getV1FunctionalitiesFunctionalityIdServicesOK  %+v", 200, o.Payload)
}

func (o *GetV1FunctionalitiesFunctionalityIDServicesOK) GetPayload() *models.FunctionalityWithAllServicesEntity {
	return o.Payload
}

func (o *GetV1FunctionalitiesFunctionalityIDServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FunctionalityWithAllServicesEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
