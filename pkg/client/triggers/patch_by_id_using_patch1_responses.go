// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// PatchByIDUsingPATCH1Reader is a Reader for the PatchByIDUsingPATCH1 structure.
type PatchByIDUsingPATCH1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchByIDUsingPATCH1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchByIDUsingPATCH1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchByIDUsingPATCH1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchByIDUsingPATCH1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchByIDUsingPATCH1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchByIDUsingPATCH1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchByIDUsingPATCH1OK creates a PatchByIDUsingPATCH1OK with default headers values
func NewPatchByIDUsingPATCH1OK() *PatchByIDUsingPATCH1OK {
	return &PatchByIDUsingPATCH1OK{}
}

/*PatchByIDUsingPATCH1OK handles this case with default header values.

'Success' with Gerrit Listener patch
*/
type PatchByIDUsingPATCH1OK struct {
	Payload models.GerritListener
}

func (o *PatchByIDUsingPATCH1OK) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchByIdUsingPATCH1OK  %+v", 200, o.Payload)
}

func (o *PatchByIDUsingPATCH1OK) GetPayload() models.GerritListener {
	return o.Payload
}

func (o *PatchByIDUsingPATCH1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListener(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewPatchByIDUsingPATCH1Unauthorized creates a PatchByIDUsingPATCH1Unauthorized with default headers values
func NewPatchByIDUsingPATCH1Unauthorized() *PatchByIDUsingPATCH1Unauthorized {
	return &PatchByIDUsingPATCH1Unauthorized{}
}

/*PatchByIDUsingPATCH1Unauthorized handles this case with default header values.

Unauthorized Request
*/
type PatchByIDUsingPATCH1Unauthorized struct {
}

func (o *PatchByIDUsingPATCH1Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchByIdUsingPATCH1Unauthorized ", 401)
}

func (o *PatchByIDUsingPATCH1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchByIDUsingPATCH1Forbidden creates a PatchByIDUsingPATCH1Forbidden with default headers values
func NewPatchByIDUsingPATCH1Forbidden() *PatchByIDUsingPATCH1Forbidden {
	return &PatchByIDUsingPATCH1Forbidden{}
}

/*PatchByIDUsingPATCH1Forbidden handles this case with default header values.

Forbidden
*/
type PatchByIDUsingPATCH1Forbidden struct {
}

func (o *PatchByIDUsingPATCH1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchByIdUsingPATCH1Forbidden ", 403)
}

func (o *PatchByIDUsingPATCH1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchByIDUsingPATCH1NotFound creates a PatchByIDUsingPATCH1NotFound with default headers values
func NewPatchByIDUsingPATCH1NotFound() *PatchByIDUsingPATCH1NotFound {
	return &PatchByIDUsingPATCH1NotFound{}
}

/*PatchByIDUsingPATCH1NotFound handles this case with default header values.

Not Found
*/
type PatchByIDUsingPATCH1NotFound struct {
}

func (o *PatchByIDUsingPATCH1NotFound) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchByIdUsingPATCH1NotFound ", 404)
}

func (o *PatchByIDUsingPATCH1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchByIDUsingPATCH1InternalServerError creates a PatchByIDUsingPATCH1InternalServerError with default headers values
func NewPatchByIDUsingPATCH1InternalServerError() *PatchByIDUsingPATCH1InternalServerError {
	return &PatchByIDUsingPATCH1InternalServerError{}
}

/*PatchByIDUsingPATCH1InternalServerError handles this case with default header values.

Server Error
*/
type PatchByIDUsingPATCH1InternalServerError struct {
}

func (o *PatchByIDUsingPATCH1InternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchByIdUsingPATCH1InternalServerError ", 500)
}

func (o *PatchByIDUsingPATCH1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
