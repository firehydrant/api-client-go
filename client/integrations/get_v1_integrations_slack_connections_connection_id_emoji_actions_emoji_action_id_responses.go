// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDReader is a Reader for the GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionID structure.
type GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK creates a GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK with default headers values
func NewGetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK() *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK {
	return &GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK{}
}

/*
GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK describes a response with status code 200, with default header values.

Retrieves a slack emoji action
*/
type GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK struct {
}

// IsSuccess returns true when this get v1 integrations slack connections connection Id emoji actions emoji action Id o k response has a 2xx status code
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 integrations slack connections connection Id emoji actions emoji action Id o k response has a 3xx status code
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 integrations slack connections connection Id emoji actions emoji action Id o k response has a 4xx status code
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 integrations slack connections connection Id emoji actions emoji action Id o k response has a 5xx status code
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 integrations slack connections connection Id emoji actions emoji action Id o k response a status code equal to that given
func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/integrations/slack/connections/{connection_id}/emoji_actions/{emoji_action_id}][%d] getV1IntegrationsSlackConnectionsConnectionIdEmojiActionsEmojiActionIdOK ", 200)
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK) String() string {
	return fmt.Sprintf("[GET /v1/integrations/slack/connections/{connection_id}/emoji_actions/{emoji_action_id}][%d] getV1IntegrationsSlackConnectionsConnectionIdEmojiActionsEmojiActionIdOK ", 200)
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsEmojiActionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
