// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1CatalogsCatalogIDIngestReader is a Reader for the PostV1CatalogsCatalogIDIngest structure.
type PostV1CatalogsCatalogIDIngestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1CatalogsCatalogIDIngestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1CatalogsCatalogIDIngestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1CatalogsCatalogIDIngestCreated creates a PostV1CatalogsCatalogIDIngestCreated with default headers values
func NewPostV1CatalogsCatalogIDIngestCreated() *PostV1CatalogsCatalogIDIngestCreated {
	return &PostV1CatalogsCatalogIDIngestCreated{}
}

/*
PostV1CatalogsCatalogIDIngestCreated describes a response with status code 201, with default header values.

Accepts catalog data in the configured format and asyncronously processes the data to incorporate changes into service catalog.
*/
type PostV1CatalogsCatalogIDIngestCreated struct {
	Payload *models.ImportsImportEntity
}

// IsSuccess returns true when this post v1 catalogs catalog Id ingest created response has a 2xx status code
func (o *PostV1CatalogsCatalogIDIngestCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 catalogs catalog Id ingest created response has a 3xx status code
func (o *PostV1CatalogsCatalogIDIngestCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 catalogs catalog Id ingest created response has a 4xx status code
func (o *PostV1CatalogsCatalogIDIngestCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 catalogs catalog Id ingest created response has a 5xx status code
func (o *PostV1CatalogsCatalogIDIngestCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 catalogs catalog Id ingest created response a status code equal to that given
func (o *PostV1CatalogsCatalogIDIngestCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1CatalogsCatalogIDIngestCreated) Error() string {
	return fmt.Sprintf("[POST /v1/catalogs/{catalog_id}/ingest][%d] postV1CatalogsCatalogIdIngestCreated  %+v", 201, o.Payload)
}

func (o *PostV1CatalogsCatalogIDIngestCreated) String() string {
	return fmt.Sprintf("[POST /v1/catalogs/{catalog_id}/ingest][%d] postV1CatalogsCatalogIdIngestCreated  %+v", 201, o.Payload)
}

func (o *PostV1CatalogsCatalogIDIngestCreated) GetPayload() *models.ImportsImportEntity {
	return o.Payload
}

func (o *PostV1CatalogsCatalogIDIngestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImportsImportEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
