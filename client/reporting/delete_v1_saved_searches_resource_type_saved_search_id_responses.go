// Code generated by go-swagger; DO NOT EDIT.

package reporting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// DeleteV1SavedSearchesResourceTypeSavedSearchIDReader is a Reader for the DeleteV1SavedSearchesResourceTypeSavedSearchID structure.
type DeleteV1SavedSearchesResourceTypeSavedSearchIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1SavedSearchesResourceTypeSavedSearchIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1SavedSearchesResourceTypeSavedSearchIDOK creates a DeleteV1SavedSearchesResourceTypeSavedSearchIDOK with default headers values
func NewDeleteV1SavedSearchesResourceTypeSavedSearchIDOK() *DeleteV1SavedSearchesResourceTypeSavedSearchIDOK {
	return &DeleteV1SavedSearchesResourceTypeSavedSearchIDOK{}
}

/*
DeleteV1SavedSearchesResourceTypeSavedSearchIDOK describes a response with status code 200, with default header values.

Delete a specific saved search
*/
type DeleteV1SavedSearchesResourceTypeSavedSearchIDOK struct {
	Payload *models.SavedSearchEntity
}

// IsSuccess returns true when this delete v1 saved searches resource type saved search Id o k response has a 2xx status code
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 saved searches resource type saved search Id o k response has a 3xx status code
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 saved searches resource type saved search Id o k response has a 4xx status code
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 saved searches resource type saved search Id o k response has a 5xx status code
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 saved searches resource type saved search Id o k response a status code equal to that given
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/saved_searches/{resource_type}/{saved_search_id}][%d] deleteV1SavedSearchesResourceTypeSavedSearchIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDOK) String() string {
	return fmt.Sprintf("[DELETE /v1/saved_searches/{resource_type}/{saved_search_id}][%d] deleteV1SavedSearchesResourceTypeSavedSearchIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDOK) GetPayload() *models.SavedSearchEntity {
	return o.Payload
}

func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SavedSearchEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
