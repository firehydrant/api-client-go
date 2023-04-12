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

// GetV1IntegrationsDataExportExportRequestsExportRequestIDReader is a Reader for the GetV1IntegrationsDataExportExportRequestsExportRequestID structure.
type GetV1IntegrationsDataExportExportRequestsExportRequestIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IntegrationsDataExportExportRequestsExportRequestIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IntegrationsDataExportExportRequestsExportRequestIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1IntegrationsDataExportExportRequestsExportRequestIDOK creates a GetV1IntegrationsDataExportExportRequestsExportRequestIDOK with default headers values
func NewGetV1IntegrationsDataExportExportRequestsExportRequestIDOK() *GetV1IntegrationsDataExportExportRequestsExportRequestIDOK {
	return &GetV1IntegrationsDataExportExportRequestsExportRequestIDOK{}
}

/*
GetV1IntegrationsDataExportExportRequestsExportRequestIDOK describes a response with status code 200, with default header values.

Retrieves a single export request for data exporting
*/
type GetV1IntegrationsDataExportExportRequestsExportRequestIDOK struct {
	Payload *models.IntegrationsDataExportExportRequestEntity
}

// IsSuccess returns true when this get v1 integrations data export export requests export request Id o k response has a 2xx status code
func (o *GetV1IntegrationsDataExportExportRequestsExportRequestIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 integrations data export export requests export request Id o k response has a 3xx status code
func (o *GetV1IntegrationsDataExportExportRequestsExportRequestIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 integrations data export export requests export request Id o k response has a 4xx status code
func (o *GetV1IntegrationsDataExportExportRequestsExportRequestIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 integrations data export export requests export request Id o k response has a 5xx status code
func (o *GetV1IntegrationsDataExportExportRequestsExportRequestIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 integrations data export export requests export request Id o k response a status code equal to that given
func (o *GetV1IntegrationsDataExportExportRequestsExportRequestIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1IntegrationsDataExportExportRequestsExportRequestIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/integrations/data_export/export_requests/{export_request_id}][%d] getV1IntegrationsDataExportExportRequestsExportRequestIdOK  %+v", 200, o.Payload)
}

func (o *GetV1IntegrationsDataExportExportRequestsExportRequestIDOK) String() string {
	return fmt.Sprintf("[GET /v1/integrations/data_export/export_requests/{export_request_id}][%d] getV1IntegrationsDataExportExportRequestsExportRequestIdOK  %+v", 200, o.Payload)
}

func (o *GetV1IntegrationsDataExportExportRequestsExportRequestIDOK) GetPayload() *models.IntegrationsDataExportExportRequestEntity {
	return o.Payload
}

func (o *GetV1IntegrationsDataExportExportRequestsExportRequestIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntegrationsDataExportExportRequestEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
