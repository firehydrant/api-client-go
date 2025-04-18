// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1ProcessingLogEntriesReader is a Reader for the GetV1ProcessingLogEntries structure.
type GetV1ProcessingLogEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1ProcessingLogEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1ProcessingLogEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1ProcessingLogEntriesOK creates a GetV1ProcessingLogEntriesOK with default headers values
func NewGetV1ProcessingLogEntriesOK() *GetV1ProcessingLogEntriesOK {
	return &GetV1ProcessingLogEntriesOK{}
}

/*
GetV1ProcessingLogEntriesOK describes a response with status code 200, with default header values.

Processing Log Entries for a specific alert
*/
type GetV1ProcessingLogEntriesOK struct {
	Payload *models.AlertsProcessingLogEntryEntityPaginated
}

// IsSuccess returns true when this get v1 processing log entries o k response has a 2xx status code
func (o *GetV1ProcessingLogEntriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 processing log entries o k response has a 3xx status code
func (o *GetV1ProcessingLogEntriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 processing log entries o k response has a 4xx status code
func (o *GetV1ProcessingLogEntriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 processing log entries o k response has a 5xx status code
func (o *GetV1ProcessingLogEntriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 processing log entries o k response a status code equal to that given
func (o *GetV1ProcessingLogEntriesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1ProcessingLogEntriesOK) Error() string {
	return fmt.Sprintf("[GET /v1/processing_log_entries][%d] getV1ProcessingLogEntriesOK  %+v", 200, o.Payload)
}

func (o *GetV1ProcessingLogEntriesOK) String() string {
	return fmt.Sprintf("[GET /v1/processing_log_entries][%d] getV1ProcessingLogEntriesOK  %+v", 200, o.Payload)
}

func (o *GetV1ProcessingLogEntriesOK) GetPayload() *models.AlertsProcessingLogEntryEntityPaginated {
	return o.Payload
}

func (o *GetV1ProcessingLogEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertsProcessingLogEntryEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
