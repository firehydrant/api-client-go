// Code generated by go-swagger; DO NOT EDIT.

package severity_matrix

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1SeverityMatrixImpactsReader is a Reader for the GetV1SeverityMatrixImpacts structure.
type GetV1SeverityMatrixImpactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1SeverityMatrixImpactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1SeverityMatrixImpactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/severity_matrix/impacts] getV1SeverityMatrixImpacts", response, response.Code())
	}
}

// NewGetV1SeverityMatrixImpactsOK creates a GetV1SeverityMatrixImpactsOK with default headers values
func NewGetV1SeverityMatrixImpactsOK() *GetV1SeverityMatrixImpactsOK {
	return &GetV1SeverityMatrixImpactsOK{}
}

/*
GetV1SeverityMatrixImpactsOK describes a response with status code 200, with default header values.

Lists impacts
*/
type GetV1SeverityMatrixImpactsOK struct {
	Payload *models.SeverityMatrixImpactEntity
}

// IsSuccess returns true when this get v1 severity matrix impacts o k response has a 2xx status code
func (o *GetV1SeverityMatrixImpactsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 severity matrix impacts o k response has a 3xx status code
func (o *GetV1SeverityMatrixImpactsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 severity matrix impacts o k response has a 4xx status code
func (o *GetV1SeverityMatrixImpactsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 severity matrix impacts o k response has a 5xx status code
func (o *GetV1SeverityMatrixImpactsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 severity matrix impacts o k response a status code equal to that given
func (o *GetV1SeverityMatrixImpactsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 severity matrix impacts o k response
func (o *GetV1SeverityMatrixImpactsOK) Code() int {
	return 200
}

func (o *GetV1SeverityMatrixImpactsOK) Error() string {
	return fmt.Sprintf("[GET /v1/severity_matrix/impacts][%d] getV1SeverityMatrixImpactsOK  %+v", 200, o.Payload)
}

func (o *GetV1SeverityMatrixImpactsOK) String() string {
	return fmt.Sprintf("[GET /v1/severity_matrix/impacts][%d] getV1SeverityMatrixImpactsOK  %+v", 200, o.Payload)
}

func (o *GetV1SeverityMatrixImpactsOK) GetPayload() *models.SeverityMatrixImpactEntity {
	return o.Payload
}

func (o *GetV1SeverityMatrixImpactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SeverityMatrixImpactEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
