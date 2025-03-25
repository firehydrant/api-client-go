// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostV1ScimV2GroupsReader is a Reader for the PostV1ScimV2Groups structure.
type PostV1ScimV2GroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ScimV2GroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1ScimV2GroupsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1ScimV2GroupsCreated creates a PostV1ScimV2GroupsCreated with default headers values
func NewPostV1ScimV2GroupsCreated() *PostV1ScimV2GroupsCreated {
	return &PostV1ScimV2GroupsCreated{}
}

/*
PostV1ScimV2GroupsCreated describes a response with status code 201, with default header values.

SCIM endpoint to create a new Team (Colloquial for Group in the SCIM protocol). Any members defined in the payload will be assigned to the team with no defined role.
*/
type PostV1ScimV2GroupsCreated struct {
}

// IsSuccess returns true when this post v1 scim v2 groups created response has a 2xx status code
func (o *PostV1ScimV2GroupsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 scim v2 groups created response has a 3xx status code
func (o *PostV1ScimV2GroupsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 scim v2 groups created response has a 4xx status code
func (o *PostV1ScimV2GroupsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 scim v2 groups created response has a 5xx status code
func (o *PostV1ScimV2GroupsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 scim v2 groups created response a status code equal to that given
func (o *PostV1ScimV2GroupsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1ScimV2GroupsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/scim/v2/Groups][%d] postV1ScimV2GroupsCreated ", 201)
}

func (o *PostV1ScimV2GroupsCreated) String() string {
	return fmt.Sprintf("[POST /v1/scim/v2/Groups][%d] postV1ScimV2GroupsCreated ", 201)
}

func (o *PostV1ScimV2GroupsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
