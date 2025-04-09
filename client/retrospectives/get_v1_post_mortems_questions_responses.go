// Code generated by go-swagger; DO NOT EDIT.

package retrospectives

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1PostMortemsQuestionsReader is a Reader for the GetV1PostMortemsQuestions structure.
type GetV1PostMortemsQuestionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1PostMortemsQuestionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1PostMortemsQuestionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1PostMortemsQuestionsOK creates a GetV1PostMortemsQuestionsOK with default headers values
func NewGetV1PostMortemsQuestionsOK() *GetV1PostMortemsQuestionsOK {
	return &GetV1PostMortemsQuestionsOK{}
}

/*
GetV1PostMortemsQuestionsOK describes a response with status code 200, with default header values.

List the questions configured to be provided and filled out on each retrospective report.
*/
type GetV1PostMortemsQuestionsOK struct {
	Payload *models.PostMortemsQuestionTypeEntityPaginated
}

// IsSuccess returns true when this get v1 post mortems questions o k response has a 2xx status code
func (o *GetV1PostMortemsQuestionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 post mortems questions o k response has a 3xx status code
func (o *GetV1PostMortemsQuestionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 post mortems questions o k response has a 4xx status code
func (o *GetV1PostMortemsQuestionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 post mortems questions o k response has a 5xx status code
func (o *GetV1PostMortemsQuestionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 post mortems questions o k response a status code equal to that given
func (o *GetV1PostMortemsQuestionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1PostMortemsQuestionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/post_mortems/questions][%d] getV1PostMortemsQuestionsOK  %+v", 200, o.Payload)
}

func (o *GetV1PostMortemsQuestionsOK) String() string {
	return fmt.Sprintf("[GET /v1/post_mortems/questions][%d] getV1PostMortemsQuestionsOK  %+v", 200, o.Payload)
}

func (o *GetV1PostMortemsQuestionsOK) GetPayload() *models.PostMortemsQuestionTypeEntityPaginated {
	return o.Payload
}

func (o *GetV1PostMortemsQuestionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMortemsQuestionTypeEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
