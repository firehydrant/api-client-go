// Code generated by go-swagger; DO NOT EDIT.

package lifecycles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDReader is a Reader for the DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionID structure.
type DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent creates a DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent with default headers values
func NewDeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent() *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent {
	return &DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent{}
}

/*
DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent describes a response with status code 204, with default header values.

Archives a measurement definition which will hide it from lists and metrics
*/
type DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent struct {
}

// IsSuccess returns true when this delete v1 lifecycles measurement definitions measurement definition Id no content response has a 2xx status code
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 lifecycles measurement definitions measurement definition Id no content response has a 3xx status code
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 lifecycles measurement definitions measurement definition Id no content response has a 4xx status code
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 lifecycles measurement definitions measurement definition Id no content response has a 5xx status code
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 lifecycles measurement definitions measurement definition Id no content response a status code equal to that given
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/lifecycles/measurement_definitions/{measurement_definition_id}][%d] deleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIdNoContent ", 204)
}

func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/lifecycles/measurement_definitions/{measurement_definition_id}][%d] deleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIdNoContent ", 204)
}

func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
