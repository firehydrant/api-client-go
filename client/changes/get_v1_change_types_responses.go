// Code generated by go-swagger; DO NOT EDIT.

package changes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1ChangeTypesReader is a Reader for the GetV1ChangeTypes structure.
type GetV1ChangeTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ChangeTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ChangeTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1ChangeTypesOK creates a GetV1ChangeTypesOK with default headers values
func NewGetV1ChangeTypesOK() *GetV1ChangeTypesOK {
	return &GetV1ChangeTypesOK{}
}

/*
GetV1ChangeTypesOK describes a response with status code 200, with default header values.

List change types for the organization
*/
type GetV1ChangeTypesOK struct {
	Payload *models.ChangeTypeEntityPaginated
}

// IsSuccess returns true when this get v1 change types o k response has a 2xx status code
func (o *GetV1ChangeTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 change types o k response has a 3xx status code
func (o *GetV1ChangeTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 change types o k response has a 4xx status code
func (o *GetV1ChangeTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 change types o k response has a 5xx status code
func (o *GetV1ChangeTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 change types o k response a status code equal to that given
func (o *GetV1ChangeTypesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1ChangeTypesOK) Error() string {
	return fmt.Sprintf("[GET /v1/change_types][%d] getV1ChangeTypesOK  %+v", 200, o.Payload)
}

func (o *GetV1ChangeTypesOK) String() string {
	return fmt.Sprintf("[GET /v1/change_types][%d] getV1ChangeTypesOK  %+v", 200, o.Payload)
}

func (o *GetV1ChangeTypesOK) GetPayload() *models.ChangeTypeEntityPaginated {
	return o.Payload
}

func (o *GetV1ChangeTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChangeTypeEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
