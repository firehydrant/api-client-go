// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PatchV1TaskListsTaskListIDReader is a Reader for the PatchV1TaskListsTaskListID structure.
type PatchV1TaskListsTaskListIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1TaskListsTaskListIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1TaskListsTaskListIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1TaskListsTaskListIDOK creates a PatchV1TaskListsTaskListIDOK with default headers values
func NewPatchV1TaskListsTaskListIDOK() *PatchV1TaskListsTaskListIDOK {
	return &PatchV1TaskListsTaskListIDOK{}
}

/*
PatchV1TaskListsTaskListIDOK describes a response with status code 200, with default header values.

Updates a task list's attributes and task list items
*/
type PatchV1TaskListsTaskListIDOK struct {
	Payload *models.TaskListEntity
}

// IsSuccess returns true when this patch v1 task lists task list Id o k response has a 2xx status code
func (o *PatchV1TaskListsTaskListIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 task lists task list Id o k response has a 3xx status code
func (o *PatchV1TaskListsTaskListIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 task lists task list Id o k response has a 4xx status code
func (o *PatchV1TaskListsTaskListIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 task lists task list Id o k response has a 5xx status code
func (o *PatchV1TaskListsTaskListIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 task lists task list Id o k response a status code equal to that given
func (o *PatchV1TaskListsTaskListIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchV1TaskListsTaskListIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/task_lists/{task_list_id}][%d] patchV1TaskListsTaskListIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1TaskListsTaskListIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/task_lists/{task_list_id}][%d] patchV1TaskListsTaskListIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1TaskListsTaskListIDOK) GetPayload() *models.TaskListEntity {
	return o.Payload
}

func (o *PatchV1TaskListsTaskListIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskListEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
