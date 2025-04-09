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

// GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDReader is a Reader for the GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionID structure.
type GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK creates a GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK with default headers values
func NewGetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK() *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK {
	return &GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK{}
}

/*
GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK describes a response with status code 200, with default header values.

Lists the available and configured integrations
*/
type GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK struct {
	Payload *models.IntegrationsAuthedProviderEntityPaginated
}

// IsSuccess returns true when this get v1 integrations authed providers integration slug connection Id o k response has a 2xx status code
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 integrations authed providers integration slug connection Id o k response has a 3xx status code
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 integrations authed providers integration slug connection Id o k response has a 4xx status code
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 integrations authed providers integration slug connection Id o k response has a 5xx status code
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 integrations authed providers integration slug connection Id o k response a status code equal to that given
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/integrations/authed_providers/{integration_slug}/{connection_id}][%d] getV1IntegrationsAuthedProvidersIntegrationSlugConnectionIdOK  %+v", 200, o.Payload)
}

func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK) String() string {
	return fmt.Sprintf("[GET /v1/integrations/authed_providers/{integration_slug}/{connection_id}][%d] getV1IntegrationsAuthedProvidersIntegrationSlugConnectionIdOK  %+v", 200, o.Payload)
}

func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK) GetPayload() *models.IntegrationsAuthedProviderEntityPaginated {
	return o.Payload
}

func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntegrationsAuthedProviderEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
