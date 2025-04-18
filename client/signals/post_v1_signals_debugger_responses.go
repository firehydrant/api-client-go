// Code generated by go-swagger; DO NOT EDIT.

package signals

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostV1SignalsDebuggerReader is a Reader for the PostV1SignalsDebugger structure.
type PostV1SignalsDebuggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1SignalsDebuggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1SignalsDebuggerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1SignalsDebuggerCreated creates a PostV1SignalsDebuggerCreated with default headers values
func NewPostV1SignalsDebuggerCreated() *PostV1SignalsDebuggerCreated {
	return &PostV1SignalsDebuggerCreated{}
}

/*
PostV1SignalsDebuggerCreated describes a response with status code 201, with default header values.

Debug Signals expressions
*/
type PostV1SignalsDebuggerCreated struct {
}

// IsSuccess returns true when this post v1 signals debugger created response has a 2xx status code
func (o *PostV1SignalsDebuggerCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 signals debugger created response has a 3xx status code
func (o *PostV1SignalsDebuggerCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 signals debugger created response has a 4xx status code
func (o *PostV1SignalsDebuggerCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 signals debugger created response has a 5xx status code
func (o *PostV1SignalsDebuggerCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 signals debugger created response a status code equal to that given
func (o *PostV1SignalsDebuggerCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1SignalsDebuggerCreated) Error() string {
	return fmt.Sprintf("[POST /v1/signals/debugger][%d] postV1SignalsDebuggerCreated ", 201)
}

func (o *PostV1SignalsDebuggerCreated) String() string {
	return fmt.Sprintf("[POST /v1/signals/debugger][%d] postV1SignalsDebuggerCreated ", 201)
}

func (o *PostV1SignalsDebuggerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
