// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PutV1PostMortemsQuestionsReader is a Reader for the PutV1PostMortemsQuestions structure.
type PutV1PostMortemsQuestionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1PostMortemsQuestionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutV1PostMortemsQuestionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /v1/post_mortems/questions] putV1PostMortemsQuestions", response, response.Code())
	}
}

// NewPutV1PostMortemsQuestionsOK creates a PutV1PostMortemsQuestionsOK with default headers values
func NewPutV1PostMortemsQuestionsOK() *PutV1PostMortemsQuestionsOK {
	return &PutV1PostMortemsQuestionsOK{}
}

/*
PutV1PostMortemsQuestionsOK describes a response with status code 200, with default header values.

Update the questions configured to be provided and filled out on future retrospective reports.
*/
type PutV1PostMortemsQuestionsOK struct {
	Payload *models.PostMortemsQuestionTypeEntity
}

// IsSuccess returns true when this put v1 post mortems questions o k response has a 2xx status code
func (o *PutV1PostMortemsQuestionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put v1 post mortems questions o k response has a 3xx status code
func (o *PutV1PostMortemsQuestionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 post mortems questions o k response has a 4xx status code
func (o *PutV1PostMortemsQuestionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put v1 post mortems questions o k response has a 5xx status code
func (o *PutV1PostMortemsQuestionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 post mortems questions o k response a status code equal to that given
func (o *PutV1PostMortemsQuestionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put v1 post mortems questions o k response
func (o *PutV1PostMortemsQuestionsOK) Code() int {
	return 200
}

func (o *PutV1PostMortemsQuestionsOK) Error() string {
	return fmt.Sprintf("[PUT /v1/post_mortems/questions][%d] putV1PostMortemsQuestionsOK  %+v", 200, o.Payload)
}

func (o *PutV1PostMortemsQuestionsOK) String() string {
	return fmt.Sprintf("[PUT /v1/post_mortems/questions][%d] putV1PostMortemsQuestionsOK  %+v", 200, o.Payload)
}

func (o *PutV1PostMortemsQuestionsOK) GetPayload() *models.PostMortemsQuestionTypeEntity {
	return o.Payload
}

func (o *PutV1PostMortemsQuestionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostMortemsQuestionTypeEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
