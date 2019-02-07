// Code generated by go-swagger; DO NOT EDIT.

package incidents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firehydrant/api-client-go/models"
)

// PatchV1IncidentsIncidentIDNotesNoteIDReader is a Reader for the PatchV1IncidentsIncidentIDNotesNoteID structure.
type PatchV1IncidentsIncidentIDNotesNoteIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1IncidentsIncidentIDNotesNoteIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchV1IncidentsIncidentIDNotesNoteIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchV1IncidentsIncidentIDNotesNoteIDOK creates a PatchV1IncidentsIncidentIDNotesNoteIDOK with default headers values
func NewPatchV1IncidentsIncidentIDNotesNoteIDOK() *PatchV1IncidentsIncidentIDNotesNoteIDOK {
	return &PatchV1IncidentsIncidentIDNotesNoteIDOK{}
}

/*PatchV1IncidentsIncidentIDNotesNoteIDOK handles this case with default header values.

Update a note
*/
type PatchV1IncidentsIncidentIDNotesNoteIDOK struct {
	Payload *models.NoteEntity
}

func (o *PatchV1IncidentsIncidentIDNotesNoteIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/incidents/{incident_id}/notes/{note_id}][%d] patchV1IncidentsIncidentIdNotesNoteIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1IncidentsIncidentIDNotesNoteIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NoteEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
