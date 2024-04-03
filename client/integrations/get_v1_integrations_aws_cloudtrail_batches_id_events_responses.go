// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1IntegrationsAwsCloudtrailBatchesIDEventsReader is a Reader for the GetV1IntegrationsAwsCloudtrailBatchesIDEvents structure.
type GetV1IntegrationsAwsCloudtrailBatchesIDEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IntegrationsAwsCloudtrailBatchesIDEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/integrations/aws/cloudtrail_batches/{id}/events] getV1IntegrationsAwsCloudtrailBatchesIdEvents", response, response.Code())
	}
}

// NewGetV1IntegrationsAwsCloudtrailBatchesIDEventsOK creates a GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK with default headers values
func NewGetV1IntegrationsAwsCloudtrailBatchesIDEventsOK() *GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK {
	return &GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK{}
}

/*
GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK describes a response with status code 200, with default header values.

get Event(s)
*/
type GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK struct {
}

// IsSuccess returns true when this get v1 integrations aws cloudtrail batches Id events o k response has a 2xx status code
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 integrations aws cloudtrail batches Id events o k response has a 3xx status code
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 integrations aws cloudtrail batches Id events o k response has a 4xx status code
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 integrations aws cloudtrail batches Id events o k response has a 5xx status code
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 integrations aws cloudtrail batches Id events o k response a status code equal to that given
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 integrations aws cloudtrail batches Id events o k response
func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK) Code() int {
	return 200
}

func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK) Error() string {
	return fmt.Sprintf("[GET /v1/integrations/aws/cloudtrail_batches/{id}/events][%d] getV1IntegrationsAwsCloudtrailBatchesIdEventsOK ", 200)
}

func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK) String() string {
	return fmt.Sprintf("[GET /v1/integrations/aws/cloudtrail_batches/{id}/events][%d] getV1IntegrationsAwsCloudtrailBatchesIdEventsOK ", 200)
}

func (o *GetV1IntegrationsAwsCloudtrailBatchesIDEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
