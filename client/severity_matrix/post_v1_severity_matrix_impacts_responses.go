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

// PostV1SeverityMatrixImpactsReader is a Reader for the PostV1SeverityMatrixImpacts structure.
type PostV1SeverityMatrixImpactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1SeverityMatrixImpactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1SeverityMatrixImpactsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /v1/severity_matrix/impacts] postV1SeverityMatrixImpacts", response, response.Code())
	}
}

// NewPostV1SeverityMatrixImpactsCreated creates a PostV1SeverityMatrixImpactsCreated with default headers values
func NewPostV1SeverityMatrixImpactsCreated() *PostV1SeverityMatrixImpactsCreated {
	return &PostV1SeverityMatrixImpactsCreated{}
}

/*
PostV1SeverityMatrixImpactsCreated describes a response with status code 201, with default header values.

Create a new impact
*/
type PostV1SeverityMatrixImpactsCreated struct {
	Payload *models.SeverityMatrixImpactEntity
}

// IsSuccess returns true when this post v1 severity matrix impacts created response has a 2xx status code
func (o *PostV1SeverityMatrixImpactsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 severity matrix impacts created response has a 3xx status code
func (o *PostV1SeverityMatrixImpactsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 severity matrix impacts created response has a 4xx status code
func (o *PostV1SeverityMatrixImpactsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 severity matrix impacts created response has a 5xx status code
func (o *PostV1SeverityMatrixImpactsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 severity matrix impacts created response a status code equal to that given
func (o *PostV1SeverityMatrixImpactsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post v1 severity matrix impacts created response
func (o *PostV1SeverityMatrixImpactsCreated) Code() int {
	return 201
}

func (o *PostV1SeverityMatrixImpactsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/severity_matrix/impacts][%d] postV1SeverityMatrixImpactsCreated  %+v", 201, o.Payload)
}

func (o *PostV1SeverityMatrixImpactsCreated) String() string {
	return fmt.Sprintf("[POST /v1/severity_matrix/impacts][%d] postV1SeverityMatrixImpactsCreated  %+v", 201, o.Payload)
}

func (o *PostV1SeverityMatrixImpactsCreated) GetPayload() *models.SeverityMatrixImpactEntity {
	return o.Payload
}

func (o *PostV1SeverityMatrixImpactsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SeverityMatrixImpactEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
