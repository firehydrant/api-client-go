// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// DeleteV1IntegrationsStatuspageConnectionsConnectionIDReader is a Reader for the DeleteV1IntegrationsStatuspageConnectionsConnectionID structure.
type DeleteV1IntegrationsStatuspageConnectionsConnectionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1IntegrationsStatuspageConnectionsConnectionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteV1IntegrationsStatuspageConnectionsConnectionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/integrations/statuspage/connections/{connection_id}] deleteV1IntegrationsStatuspageConnectionsConnectionId", response, response.Code())
	}
}

// NewDeleteV1IntegrationsStatuspageConnectionsConnectionIDOK creates a DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK with default headers values
func NewDeleteV1IntegrationsStatuspageConnectionsConnectionIDOK() *DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK {
	return &DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK{}
}

/*
DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK describes a response with status code 200, with default header values.

Deletes the given Statuspage integration connection.
*/
type DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK struct {
	Payload *models.IntegrationsStatuspageConnectionEntity
}

// IsSuccess returns true when this delete v1 integrations statuspage connections connection Id o k response has a 2xx status code
func (o *DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 integrations statuspage connections connection Id o k response has a 3xx status code
func (o *DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 integrations statuspage connections connection Id o k response has a 4xx status code
func (o *DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 integrations statuspage connections connection Id o k response has a 5xx status code
func (o *DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 integrations statuspage connections connection Id o k response a status code equal to that given
func (o *DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete v1 integrations statuspage connections connection Id o k response
func (o *DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK) Code() int {
	return 200
}

func (o *DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/integrations/statuspage/connections/{connection_id}][%d] deleteV1IntegrationsStatuspageConnectionsConnectionIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK) String() string {
	return fmt.Sprintf("[DELETE /v1/integrations/statuspage/connections/{connection_id}][%d] deleteV1IntegrationsStatuspageConnectionsConnectionIdOK  %+v", 200, o.Payload)
}

func (o *DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK) GetPayload() *models.IntegrationsStatuspageConnectionEntity {
	return o.Payload
}

func (o *DeleteV1IntegrationsStatuspageConnectionsConnectionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntegrationsStatuspageConnectionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
