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

// DeleteV1SeverityMatrixConditionsConditionIDReader is a Reader for the DeleteV1SeverityMatrixConditionsConditionID structure.
type DeleteV1SeverityMatrixConditionsConditionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1SeverityMatrixConditionsConditionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1SeverityMatrixConditionsConditionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1SeverityMatrixConditionsConditionIDOK creates a DeleteV1SeverityMatrixConditionsConditionIDOK with default headers values
func NewDeleteV1SeverityMatrixConditionsConditionIDOK() *DeleteV1SeverityMatrixConditionsConditionIDOK {
	return &DeleteV1SeverityMatrixConditionsConditionIDOK{}
}

/*
DeleteV1SeverityMatrixConditionsConditionIDOK describes a response with status code 200, with default header values.

Delete a specific condition
*/
type DeleteV1SeverityMatrixConditionsConditionIDOK struct {
	Payload *models.SeverityMatrixConditionEntity
}

// IsSuccess returns true when this delete v1 severity matrix conditions condition Id o k response has a 2xx status code
func (o *DeleteV1SeverityMatrixConditionsConditionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 severity matrix conditions condition Id o k response has a 3xx status code
func (o *DeleteV1SeverityMatrixConditionsConditionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 severity matrix conditions condition Id o k response has a 4xx status code
func (o *DeleteV1SeverityMatrixConditionsConditionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 severity matrix conditions condition Id o k response has a 5xx status code
func (o *DeleteV1SeverityMatrixConditionsConditionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 severity matrix conditions condition Id o k response a status code equal to that given
func (o *DeleteV1SeverityMatrixConditionsConditionIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteV1SeverityMatrixConditionsConditionIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/severity_matrix/conditions/{condition_id}][%d] deleteV1SeverityMatrixConditionsConditionIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1SeverityMatrixConditionsConditionIDOK) String() string {
	return fmt.Sprintf("[DELETE /v1/severity_matrix/conditions/{condition_id}][%d] deleteV1SeverityMatrixConditionsConditionIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1SeverityMatrixConditionsConditionIDOK) GetPayload() *models.SeverityMatrixConditionEntity {
	return o.Payload
}

func (o *DeleteV1SeverityMatrixConditionsConditionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SeverityMatrixConditionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
