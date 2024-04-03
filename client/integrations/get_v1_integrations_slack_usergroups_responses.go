// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1IntegrationsSlackUsergroupsReader is a Reader for the GetV1IntegrationsSlackUsergroups structure.
type GetV1IntegrationsSlackUsergroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IntegrationsSlackUsergroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IntegrationsSlackUsergroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1IntegrationsSlackUsergroupsOK creates a GetV1IntegrationsSlackUsergroupsOK with default headers values
func NewGetV1IntegrationsSlackUsergroupsOK() *GetV1IntegrationsSlackUsergroupsOK {
	return &GetV1IntegrationsSlackUsergroupsOK{}
}

/*
GetV1IntegrationsSlackUsergroupsOK describes a response with status code 200, with default header values.

get Usergroup(s)
*/
type GetV1IntegrationsSlackUsergroupsOK struct {
}

// IsSuccess returns true when this get v1 integrations slack usergroups o k response has a 2xx status code
func (o *GetV1IntegrationsSlackUsergroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 integrations slack usergroups o k response has a 3xx status code
func (o *GetV1IntegrationsSlackUsergroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 integrations slack usergroups o k response has a 4xx status code
func (o *GetV1IntegrationsSlackUsergroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 integrations slack usergroups o k response has a 5xx status code
func (o *GetV1IntegrationsSlackUsergroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 integrations slack usergroups o k response a status code equal to that given
func (o *GetV1IntegrationsSlackUsergroupsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1IntegrationsSlackUsergroupsOK) Error() string {
	return fmt.Sprintf("[GET /v1/integrations/slack/usergroups][%d] getV1IntegrationsSlackUsergroupsOK ", 200)
}

func (o *GetV1IntegrationsSlackUsergroupsOK) String() string {
	return fmt.Sprintf("[GET /v1/integrations/slack/usergroups][%d] getV1IntegrationsSlackUsergroupsOK ", 200)
}

func (o *GetV1IntegrationsSlackUsergroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}