// Code generated by go-swagger; DO NOT EDIT.

package custom_fields

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1CustomFieldsDefinitionsReader is a Reader for the GetV1CustomFieldsDefinitions structure.
type GetV1CustomFieldsDefinitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1CustomFieldsDefinitionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1CustomFieldsDefinitionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1CustomFieldsDefinitionsOK creates a GetV1CustomFieldsDefinitionsOK with default headers values
func NewGetV1CustomFieldsDefinitionsOK() *GetV1CustomFieldsDefinitionsOK {
	return &GetV1CustomFieldsDefinitionsOK{}
}

/*
GetV1CustomFieldsDefinitionsOK describes a response with status code 200, with default header values.

List all custom field definitions
*/
type GetV1CustomFieldsDefinitionsOK struct {
	Payload *models.OrganizationsCustomFieldDefinitionEntity
}

// IsSuccess returns true when this get v1 custom fields definitions o k response has a 2xx status code
func (o *GetV1CustomFieldsDefinitionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 custom fields definitions o k response has a 3xx status code
func (o *GetV1CustomFieldsDefinitionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 custom fields definitions o k response has a 4xx status code
func (o *GetV1CustomFieldsDefinitionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 custom fields definitions o k response has a 5xx status code
func (o *GetV1CustomFieldsDefinitionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 custom fields definitions o k response a status code equal to that given
func (o *GetV1CustomFieldsDefinitionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1CustomFieldsDefinitionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/custom_fields/definitions][%d] getV1CustomFieldsDefinitionsOK  %+v", 200, o.Payload)
}

func (o *GetV1CustomFieldsDefinitionsOK) String() string {
	return fmt.Sprintf("[GET /v1/custom_fields/definitions][%d] getV1CustomFieldsDefinitionsOK  %+v", 200, o.Payload)
}

func (o *GetV1CustomFieldsDefinitionsOK) GetPayload() *models.OrganizationsCustomFieldDefinitionEntity {
	return o.Payload
}

func (o *GetV1CustomFieldsDefinitionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationsCustomFieldDefinitionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}