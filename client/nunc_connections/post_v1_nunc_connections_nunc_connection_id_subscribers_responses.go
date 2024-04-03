// Code generated by go-swagger; DO NOT EDIT.

package nunc_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1NuncConnectionsNuncConnectionIDSubscribersReader is a Reader for the PostV1NuncConnectionsNuncConnectionIDSubscribers structure.
type PostV1NuncConnectionsNuncConnectionIDSubscribersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1NuncConnectionsNuncConnectionIDSubscribersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1NuncConnectionsNuncConnectionIDSubscribersCreated creates a PostV1NuncConnectionsNuncConnectionIDSubscribersCreated with default headers values
func NewPostV1NuncConnectionsNuncConnectionIDSubscribersCreated() *PostV1NuncConnectionsNuncConnectionIDSubscribersCreated {
	return &PostV1NuncConnectionsNuncConnectionIDSubscribersCreated{}
}

/*
PostV1NuncConnectionsNuncConnectionIDSubscribersCreated describes a response with status code 201, with default header values.

Subscribes a comma-separated string of emails to status page updates
*/
type PostV1NuncConnectionsNuncConnectionIDSubscribersCreated struct {
	Payload *models.NuncEmailSubscribersEntity
}

// IsSuccess returns true when this post v1 nunc connections nunc connection Id subscribers created response has a 2xx status code
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 nunc connections nunc connection Id subscribers created response has a 3xx status code
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 nunc connections nunc connection Id subscribers created response has a 4xx status code
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 nunc connections nunc connection Id subscribers created response has a 5xx status code
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 nunc connections nunc connection Id subscribers created response a status code equal to that given
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersCreated) Error() string {
	return fmt.Sprintf("[POST /v1/nunc_connections/{nunc_connection_id}/subscribers][%d] postV1NuncConnectionsNuncConnectionIdSubscribersCreated  %+v", 201, o.Payload)
}

func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersCreated) String() string {
	return fmt.Sprintf("[POST /v1/nunc_connections/{nunc_connection_id}/subscribers][%d] postV1NuncConnectionsNuncConnectionIdSubscribersCreated  %+v", 201, o.Payload)
}

func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersCreated) GetPayload() *models.NuncEmailSubscribersEntity {
	return o.Payload
}

func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NuncEmailSubscribersEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}