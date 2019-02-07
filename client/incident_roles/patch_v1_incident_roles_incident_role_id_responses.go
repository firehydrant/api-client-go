// Code generated by go-swagger; DO NOT EDIT.

package incident_roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/firehydrant/api-client-go/models"
)

// PatchV1IncidentRolesIncidentRoleIDReader is a Reader for the PatchV1IncidentRolesIncidentRoleID structure.
type PatchV1IncidentRolesIncidentRoleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1IncidentRolesIncidentRoleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchV1IncidentRolesIncidentRoleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchV1IncidentRolesIncidentRoleIDOK creates a PatchV1IncidentRolesIncidentRoleIDOK with default headers values
func NewPatchV1IncidentRolesIncidentRoleIDOK() *PatchV1IncidentRolesIncidentRoleIDOK {
	return &PatchV1IncidentRolesIncidentRoleIDOK{}
}

/*PatchV1IncidentRolesIncidentRoleIDOK handles this case with default header values.

Update an incident role
*/
type PatchV1IncidentRolesIncidentRoleIDOK struct {
	Payload *models.IncidentRoleEntity
}

func (o *PatchV1IncidentRolesIncidentRoleIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/incident_roles/{incident_role_id}][%d] patchV1IncidentRolesIncidentRoleIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1IncidentRolesIncidentRoleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncidentRoleEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
