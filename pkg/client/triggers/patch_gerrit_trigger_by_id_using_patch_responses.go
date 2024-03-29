// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// PatchGerritTriggerByIDUsingPATCHReader is a Reader for the PatchGerritTriggerByIDUsingPATCH structure.
type PatchGerritTriggerByIDUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchGerritTriggerByIDUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchGerritTriggerByIDUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchGerritTriggerByIDUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchGerritTriggerByIDUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchGerritTriggerByIDUsingPATCHNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchGerritTriggerByIDUsingPATCHInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchGerritTriggerByIDUsingPATCHOK creates a PatchGerritTriggerByIDUsingPATCHOK with default headers values
func NewPatchGerritTriggerByIDUsingPATCHOK() *PatchGerritTriggerByIDUsingPATCHOK {
	return &PatchGerritTriggerByIDUsingPATCHOK{}
}

/*
PatchGerritTriggerByIDUsingPATCHOK describes a response with status code 200, with default header values.

'Success' with Gerrit Trigger patch
*/
type PatchGerritTriggerByIDUsingPATCHOK struct {
	Payload models.GerritTrigger
}

// IsSuccess returns true when this patch gerrit trigger by Id using p a t c h o k response has a 2xx status code
func (o *PatchGerritTriggerByIDUsingPATCHOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch gerrit trigger by Id using p a t c h o k response has a 3xx status code
func (o *PatchGerritTriggerByIDUsingPATCHOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch gerrit trigger by Id using p a t c h o k response has a 4xx status code
func (o *PatchGerritTriggerByIDUsingPATCHOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch gerrit trigger by Id using p a t c h o k response has a 5xx status code
func (o *PatchGerritTriggerByIDUsingPATCHOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch gerrit trigger by Id using p a t c h o k response a status code equal to that given
func (o *PatchGerritTriggerByIDUsingPATCHOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchGerritTriggerByIDUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-triggers/{id}][%d] patchGerritTriggerByIdUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchGerritTriggerByIDUsingPATCHOK) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-triggers/{id}][%d] patchGerritTriggerByIdUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchGerritTriggerByIDUsingPATCHOK) GetPayload() models.GerritTrigger {
	return o.Payload
}

func (o *PatchGerritTriggerByIDUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritTrigger(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewPatchGerritTriggerByIDUsingPATCHUnauthorized creates a PatchGerritTriggerByIDUsingPATCHUnauthorized with default headers values
func NewPatchGerritTriggerByIDUsingPATCHUnauthorized() *PatchGerritTriggerByIDUsingPATCHUnauthorized {
	return &PatchGerritTriggerByIDUsingPATCHUnauthorized{}
}

/*
PatchGerritTriggerByIDUsingPATCHUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type PatchGerritTriggerByIDUsingPATCHUnauthorized struct {
}

// IsSuccess returns true when this patch gerrit trigger by Id using p a t c h unauthorized response has a 2xx status code
func (o *PatchGerritTriggerByIDUsingPATCHUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch gerrit trigger by Id using p a t c h unauthorized response has a 3xx status code
func (o *PatchGerritTriggerByIDUsingPATCHUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch gerrit trigger by Id using p a t c h unauthorized response has a 4xx status code
func (o *PatchGerritTriggerByIDUsingPATCHUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch gerrit trigger by Id using p a t c h unauthorized response has a 5xx status code
func (o *PatchGerritTriggerByIDUsingPATCHUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch gerrit trigger by Id using p a t c h unauthorized response a status code equal to that given
func (o *PatchGerritTriggerByIDUsingPATCHUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchGerritTriggerByIDUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-triggers/{id}][%d] patchGerritTriggerByIdUsingPATCHUnauthorized ", 401)
}

func (o *PatchGerritTriggerByIDUsingPATCHUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-triggers/{id}][%d] patchGerritTriggerByIdUsingPATCHUnauthorized ", 401)
}

func (o *PatchGerritTriggerByIDUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchGerritTriggerByIDUsingPATCHForbidden creates a PatchGerritTriggerByIDUsingPATCHForbidden with default headers values
func NewPatchGerritTriggerByIDUsingPATCHForbidden() *PatchGerritTriggerByIDUsingPATCHForbidden {
	return &PatchGerritTriggerByIDUsingPATCHForbidden{}
}

/*
PatchGerritTriggerByIDUsingPATCHForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchGerritTriggerByIDUsingPATCHForbidden struct {
}

// IsSuccess returns true when this patch gerrit trigger by Id using p a t c h forbidden response has a 2xx status code
func (o *PatchGerritTriggerByIDUsingPATCHForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch gerrit trigger by Id using p a t c h forbidden response has a 3xx status code
func (o *PatchGerritTriggerByIDUsingPATCHForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch gerrit trigger by Id using p a t c h forbidden response has a 4xx status code
func (o *PatchGerritTriggerByIDUsingPATCHForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch gerrit trigger by Id using p a t c h forbidden response has a 5xx status code
func (o *PatchGerritTriggerByIDUsingPATCHForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch gerrit trigger by Id using p a t c h forbidden response a status code equal to that given
func (o *PatchGerritTriggerByIDUsingPATCHForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchGerritTriggerByIDUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-triggers/{id}][%d] patchGerritTriggerByIdUsingPATCHForbidden ", 403)
}

func (o *PatchGerritTriggerByIDUsingPATCHForbidden) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-triggers/{id}][%d] patchGerritTriggerByIdUsingPATCHForbidden ", 403)
}

func (o *PatchGerritTriggerByIDUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchGerritTriggerByIDUsingPATCHNotFound creates a PatchGerritTriggerByIDUsingPATCHNotFound with default headers values
func NewPatchGerritTriggerByIDUsingPATCHNotFound() *PatchGerritTriggerByIDUsingPATCHNotFound {
	return &PatchGerritTriggerByIDUsingPATCHNotFound{}
}

/*
PatchGerritTriggerByIDUsingPATCHNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchGerritTriggerByIDUsingPATCHNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this patch gerrit trigger by Id using p a t c h not found response has a 2xx status code
func (o *PatchGerritTriggerByIDUsingPATCHNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch gerrit trigger by Id using p a t c h not found response has a 3xx status code
func (o *PatchGerritTriggerByIDUsingPATCHNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch gerrit trigger by Id using p a t c h not found response has a 4xx status code
func (o *PatchGerritTriggerByIDUsingPATCHNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch gerrit trigger by Id using p a t c h not found response has a 5xx status code
func (o *PatchGerritTriggerByIDUsingPATCHNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch gerrit trigger by Id using p a t c h not found response a status code equal to that given
func (o *PatchGerritTriggerByIDUsingPATCHNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchGerritTriggerByIDUsingPATCHNotFound) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-triggers/{id}][%d] patchGerritTriggerByIdUsingPATCHNotFound  %+v", 404, o.Payload)
}

func (o *PatchGerritTriggerByIDUsingPATCHNotFound) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-triggers/{id}][%d] patchGerritTriggerByIdUsingPATCHNotFound  %+v", 404, o.Payload)
}

func (o *PatchGerritTriggerByIDUsingPATCHNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchGerritTriggerByIDUsingPATCHNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGerritTriggerByIDUsingPATCHInternalServerError creates a PatchGerritTriggerByIDUsingPATCHInternalServerError with default headers values
func NewPatchGerritTriggerByIDUsingPATCHInternalServerError() *PatchGerritTriggerByIDUsingPATCHInternalServerError {
	return &PatchGerritTriggerByIDUsingPATCHInternalServerError{}
}

/*
PatchGerritTriggerByIDUsingPATCHInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PatchGerritTriggerByIDUsingPATCHInternalServerError struct {
}

// IsSuccess returns true when this patch gerrit trigger by Id using p a t c h internal server error response has a 2xx status code
func (o *PatchGerritTriggerByIDUsingPATCHInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch gerrit trigger by Id using p a t c h internal server error response has a 3xx status code
func (o *PatchGerritTriggerByIDUsingPATCHInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch gerrit trigger by Id using p a t c h internal server error response has a 4xx status code
func (o *PatchGerritTriggerByIDUsingPATCHInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch gerrit trigger by Id using p a t c h internal server error response has a 5xx status code
func (o *PatchGerritTriggerByIDUsingPATCHInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch gerrit trigger by Id using p a t c h internal server error response a status code equal to that given
func (o *PatchGerritTriggerByIDUsingPATCHInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchGerritTriggerByIDUsingPATCHInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-triggers/{id}][%d] patchGerritTriggerByIdUsingPATCHInternalServerError ", 500)
}

func (o *PatchGerritTriggerByIDUsingPATCHInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-triggers/{id}][%d] patchGerritTriggerByIdUsingPATCHInternalServerError ", 500)
}

func (o *PatchGerritTriggerByIDUsingPATCHInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
