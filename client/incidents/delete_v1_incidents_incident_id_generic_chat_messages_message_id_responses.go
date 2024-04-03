// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDReader is a Reader for the DeleteV1IncidentsIncidentIDGenericChatMessagesMessageID structure.
type DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/incidents/{incident_id}/generic_chat_messages/{message_id}] deleteV1IncidentsIncidentIdGenericChatMessagesMessageId", response, response.Code())
	}
}

// NewDeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK creates a DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK with default headers values
func NewDeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK() *DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK {
	return &DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK{}
}

/*
DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK describes a response with status code 200, with default header values.

Delete an existing generic chat message on an incident.
*/
type DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK struct {
	Payload *models.EventGenericChatMessageEntity
}

// IsSuccess returns true when this delete v1 incidents incident Id generic chat messages message Id o k response has a 2xx status code
func (o *DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 incidents incident Id generic chat messages message Id o k response has a 3xx status code
func (o *DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 incidents incident Id generic chat messages message Id o k response has a 4xx status code
func (o *DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 incidents incident Id generic chat messages message Id o k response has a 5xx status code
func (o *DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 incidents incident Id generic chat messages message Id o k response a status code equal to that given
func (o *DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete v1 incidents incident Id generic chat messages message Id o k response
func (o *DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK) Code() int {
	return 200
}

func (o *DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/incidents/{incident_id}/generic_chat_messages/{message_id}][%d] deleteV1IncidentsIncidentIdGenericChatMessagesMessageIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK) String() string {
	return fmt.Sprintf("[DELETE /v1/incidents/{incident_id}/generic_chat_messages/{message_id}][%d] deleteV1IncidentsIncidentIdGenericChatMessagesMessageIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK) GetPayload() *models.EventGenericChatMessageEntity {
	return o.Payload
}

func (o *DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventGenericChatMessageEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
