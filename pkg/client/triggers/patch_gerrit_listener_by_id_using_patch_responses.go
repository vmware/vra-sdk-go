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

// PatchGerritListenerByIDUsingPATCHReader is a Reader for the PatchGerritListenerByIDUsingPATCH structure.
type PatchGerritListenerByIDUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchGerritListenerByIDUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchGerritListenerByIDUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchGerritListenerByIDUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchGerritListenerByIDUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchGerritListenerByIDUsingPATCHNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchGerritListenerByIDUsingPATCHInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchGerritListenerByIDUsingPATCHOK creates a PatchGerritListenerByIDUsingPATCHOK with default headers values
func NewPatchGerritListenerByIDUsingPATCHOK() *PatchGerritListenerByIDUsingPATCHOK {
	return &PatchGerritListenerByIDUsingPATCHOK{}
}

/*
PatchGerritListenerByIDUsingPATCHOK describes a response with status code 200, with default header values.

'Success' with Gerrit Listener patch
*/
type PatchGerritListenerByIDUsingPATCHOK struct {
	Payload models.GerritListener
}

// IsSuccess returns true when this patch gerrit listener by Id using p a t c h o k response has a 2xx status code
func (o *PatchGerritListenerByIDUsingPATCHOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch gerrit listener by Id using p a t c h o k response has a 3xx status code
func (o *PatchGerritListenerByIDUsingPATCHOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch gerrit listener by Id using p a t c h o k response has a 4xx status code
func (o *PatchGerritListenerByIDUsingPATCHOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch gerrit listener by Id using p a t c h o k response has a 5xx status code
func (o *PatchGerritListenerByIDUsingPATCHOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch gerrit listener by Id using p a t c h o k response a status code equal to that given
func (o *PatchGerritListenerByIDUsingPATCHOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchGerritListenerByIDUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchGerritListenerByIdUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchGerritListenerByIDUsingPATCHOK) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchGerritListenerByIdUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchGerritListenerByIDUsingPATCHOK) GetPayload() models.GerritListener {
	return o.Payload
}

func (o *PatchGerritListenerByIDUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListener(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewPatchGerritListenerByIDUsingPATCHUnauthorized creates a PatchGerritListenerByIDUsingPATCHUnauthorized with default headers values
func NewPatchGerritListenerByIDUsingPATCHUnauthorized() *PatchGerritListenerByIDUsingPATCHUnauthorized {
	return &PatchGerritListenerByIDUsingPATCHUnauthorized{}
}

/*
PatchGerritListenerByIDUsingPATCHUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type PatchGerritListenerByIDUsingPATCHUnauthorized struct {
}

// IsSuccess returns true when this patch gerrit listener by Id using p a t c h unauthorized response has a 2xx status code
func (o *PatchGerritListenerByIDUsingPATCHUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch gerrit listener by Id using p a t c h unauthorized response has a 3xx status code
func (o *PatchGerritListenerByIDUsingPATCHUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch gerrit listener by Id using p a t c h unauthorized response has a 4xx status code
func (o *PatchGerritListenerByIDUsingPATCHUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch gerrit listener by Id using p a t c h unauthorized response has a 5xx status code
func (o *PatchGerritListenerByIDUsingPATCHUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch gerrit listener by Id using p a t c h unauthorized response a status code equal to that given
func (o *PatchGerritListenerByIDUsingPATCHUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchGerritListenerByIDUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchGerritListenerByIdUsingPATCHUnauthorized ", 401)
}

func (o *PatchGerritListenerByIDUsingPATCHUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchGerritListenerByIdUsingPATCHUnauthorized ", 401)
}

func (o *PatchGerritListenerByIDUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchGerritListenerByIDUsingPATCHForbidden creates a PatchGerritListenerByIDUsingPATCHForbidden with default headers values
func NewPatchGerritListenerByIDUsingPATCHForbidden() *PatchGerritListenerByIDUsingPATCHForbidden {
	return &PatchGerritListenerByIDUsingPATCHForbidden{}
}

/*
PatchGerritListenerByIDUsingPATCHForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchGerritListenerByIDUsingPATCHForbidden struct {
}

// IsSuccess returns true when this patch gerrit listener by Id using p a t c h forbidden response has a 2xx status code
func (o *PatchGerritListenerByIDUsingPATCHForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch gerrit listener by Id using p a t c h forbidden response has a 3xx status code
func (o *PatchGerritListenerByIDUsingPATCHForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch gerrit listener by Id using p a t c h forbidden response has a 4xx status code
func (o *PatchGerritListenerByIDUsingPATCHForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch gerrit listener by Id using p a t c h forbidden response has a 5xx status code
func (o *PatchGerritListenerByIDUsingPATCHForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch gerrit listener by Id using p a t c h forbidden response a status code equal to that given
func (o *PatchGerritListenerByIDUsingPATCHForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchGerritListenerByIDUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchGerritListenerByIdUsingPATCHForbidden ", 403)
}

func (o *PatchGerritListenerByIDUsingPATCHForbidden) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchGerritListenerByIdUsingPATCHForbidden ", 403)
}

func (o *PatchGerritListenerByIDUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchGerritListenerByIDUsingPATCHNotFound creates a PatchGerritListenerByIDUsingPATCHNotFound with default headers values
func NewPatchGerritListenerByIDUsingPATCHNotFound() *PatchGerritListenerByIDUsingPATCHNotFound {
	return &PatchGerritListenerByIDUsingPATCHNotFound{}
}

/*
PatchGerritListenerByIDUsingPATCHNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchGerritListenerByIDUsingPATCHNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this patch gerrit listener by Id using p a t c h not found response has a 2xx status code
func (o *PatchGerritListenerByIDUsingPATCHNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch gerrit listener by Id using p a t c h not found response has a 3xx status code
func (o *PatchGerritListenerByIDUsingPATCHNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch gerrit listener by Id using p a t c h not found response has a 4xx status code
func (o *PatchGerritListenerByIDUsingPATCHNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch gerrit listener by Id using p a t c h not found response has a 5xx status code
func (o *PatchGerritListenerByIDUsingPATCHNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch gerrit listener by Id using p a t c h not found response a status code equal to that given
func (o *PatchGerritListenerByIDUsingPATCHNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchGerritListenerByIDUsingPATCHNotFound) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchGerritListenerByIdUsingPATCHNotFound  %+v", 404, o.Payload)
}

func (o *PatchGerritListenerByIDUsingPATCHNotFound) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchGerritListenerByIdUsingPATCHNotFound  %+v", 404, o.Payload)
}

func (o *PatchGerritListenerByIDUsingPATCHNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchGerritListenerByIDUsingPATCHNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGerritListenerByIDUsingPATCHInternalServerError creates a PatchGerritListenerByIDUsingPATCHInternalServerError with default headers values
func NewPatchGerritListenerByIDUsingPATCHInternalServerError() *PatchGerritListenerByIDUsingPATCHInternalServerError {
	return &PatchGerritListenerByIDUsingPATCHInternalServerError{}
}

/*
PatchGerritListenerByIDUsingPATCHInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PatchGerritListenerByIDUsingPATCHInternalServerError struct {
}

// IsSuccess returns true when this patch gerrit listener by Id using p a t c h internal server error response has a 2xx status code
func (o *PatchGerritListenerByIDUsingPATCHInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch gerrit listener by Id using p a t c h internal server error response has a 3xx status code
func (o *PatchGerritListenerByIDUsingPATCHInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch gerrit listener by Id using p a t c h internal server error response has a 4xx status code
func (o *PatchGerritListenerByIDUsingPATCHInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch gerrit listener by Id using p a t c h internal server error response has a 5xx status code
func (o *PatchGerritListenerByIDUsingPATCHInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch gerrit listener by Id using p a t c h internal server error response a status code equal to that given
func (o *PatchGerritListenerByIDUsingPATCHInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchGerritListenerByIDUsingPATCHInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchGerritListenerByIdUsingPATCHInternalServerError ", 500)
}

func (o *PatchGerritListenerByIDUsingPATCHInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/gerrit-listeners/{id}][%d] patchGerritListenerByIdUsingPATCHInternalServerError ", 500)
}

func (o *PatchGerritListenerByIDUsingPATCHInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
