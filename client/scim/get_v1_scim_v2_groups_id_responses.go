// Code generated by go-swagger; DO NOT EDIT.

package scim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1ScimV2GroupsIDReader is a Reader for the GetV1ScimV2GroupsID structure.
type GetV1ScimV2GroupsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ScimV2GroupsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ScimV2GroupsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1ScimV2GroupsIDOK creates a GetV1ScimV2GroupsIDOK with default headers values
func NewGetV1ScimV2GroupsIDOK() *GetV1ScimV2GroupsIDOK {
	return &GetV1ScimV2GroupsIDOK{}
}

/*
GetV1ScimV2GroupsIDOK describes a response with status code 200, with default header values.

SCIM endpoint that lists a Team (Colloquial for Group in the SCIM protocol)
*/
type GetV1ScimV2GroupsIDOK struct {
}

// IsSuccess returns true when this get v1 scim v2 groups Id o k response has a 2xx status code
func (o *GetV1ScimV2GroupsIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 scim v2 groups Id o k response has a 3xx status code
func (o *GetV1ScimV2GroupsIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 scim v2 groups Id o k response has a 4xx status code
func (o *GetV1ScimV2GroupsIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 scim v2 groups Id o k response has a 5xx status code
func (o *GetV1ScimV2GroupsIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 scim v2 groups Id o k response a status code equal to that given
func (o *GetV1ScimV2GroupsIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1ScimV2GroupsIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/scim/v2/Groups/{id}][%d] getV1ScimV2GroupsIdOK ", 200)
}

func (o *GetV1ScimV2GroupsIDOK) String() string {
	return fmt.Sprintf("[GET /v1/scim/v2/Groups/{id}][%d] getV1ScimV2GroupsIdOK ", 200)
}

func (o *GetV1ScimV2GroupsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
