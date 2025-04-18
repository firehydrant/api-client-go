// Code generated by go-swagger; DO NOT EDIT.

package status_pages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1NuncConnectionsNuncConnectionIDLinksReader is a Reader for the PostV1NuncConnectionsNuncConnectionIDLinks structure.
type PostV1NuncConnectionsNuncConnectionIDLinksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1NuncConnectionsNuncConnectionIDLinksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1NuncConnectionsNuncConnectionIDLinksCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1NuncConnectionsNuncConnectionIDLinksCreated creates a PostV1NuncConnectionsNuncConnectionIDLinksCreated with default headers values
func NewPostV1NuncConnectionsNuncConnectionIDLinksCreated() *PostV1NuncConnectionsNuncConnectionIDLinksCreated {
	return &PostV1NuncConnectionsNuncConnectionIDLinksCreated{}
}

/*
PostV1NuncConnectionsNuncConnectionIDLinksCreated describes a response with status code 201, with default header values.

Add a link to be displayed on a FireHydrant status page
*/
type PostV1NuncConnectionsNuncConnectionIDLinksCreated struct {
	Payload *models.NuncConnectionEntity
}

// IsSuccess returns true when this post v1 nunc connections nunc connection Id links created response has a 2xx status code
func (o *PostV1NuncConnectionsNuncConnectionIDLinksCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 nunc connections nunc connection Id links created response has a 3xx status code
func (o *PostV1NuncConnectionsNuncConnectionIDLinksCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 nunc connections nunc connection Id links created response has a 4xx status code
func (o *PostV1NuncConnectionsNuncConnectionIDLinksCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 nunc connections nunc connection Id links created response has a 5xx status code
func (o *PostV1NuncConnectionsNuncConnectionIDLinksCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 nunc connections nunc connection Id links created response a status code equal to that given
func (o *PostV1NuncConnectionsNuncConnectionIDLinksCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1NuncConnectionsNuncConnectionIDLinksCreated) Error() string {
	return fmt.Sprintf("[POST /v1/nunc_connections/{nunc_connection_id}/links][%d] postV1NuncConnectionsNuncConnectionIdLinksCreated  %+v", 201, o.Payload)
}

func (o *PostV1NuncConnectionsNuncConnectionIDLinksCreated) String() string {
	return fmt.Sprintf("[POST /v1/nunc_connections/{nunc_connection_id}/links][%d] postV1NuncConnectionsNuncConnectionIdLinksCreated  %+v", 201, o.Payload)
}

func (o *PostV1NuncConnectionsNuncConnectionIDLinksCreated) GetPayload() *models.NuncConnectionEntity {
	return o.Payload
}

func (o *PostV1NuncConnectionsNuncConnectionIDLinksCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NuncConnectionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
