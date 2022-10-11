// Code generated by go-swagger; DO NOT EDIT.

package user_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// ModifyPatchUserOperationUsingPATCHReader is a Reader for the ModifyPatchUserOperationUsingPATCH structure.
type ModifyPatchUserOperationUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyPatchUserOperationUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModifyPatchUserOperationUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewModifyPatchUserOperationUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewModifyPatchUserOperationUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewModifyPatchUserOperationUsingPATCHNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewModifyPatchUserOperationUsingPATCHInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewModifyPatchUserOperationUsingPATCHOK creates a ModifyPatchUserOperationUsingPATCHOK with default headers values
func NewModifyPatchUserOperationUsingPATCHOK() *ModifyPatchUserOperationUsingPATCHOK {
	return &ModifyPatchUserOperationUsingPATCHOK{}
}

/*
ModifyPatchUserOperationUsingPATCHOK describes a response with status code 200, with default header values.

'Success' with the modified User Operation
*/
type ModifyPatchUserOperationUsingPATCHOK struct {
	Payload *models.UserOpResource
}

// IsSuccess returns true when this modify patch user operation using p a t c h o k response has a 2xx status code
func (o *ModifyPatchUserOperationUsingPATCHOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this modify patch user operation using p a t c h o k response has a 3xx status code
func (o *ModifyPatchUserOperationUsingPATCHOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify patch user operation using p a t c h o k response has a 4xx status code
func (o *ModifyPatchUserOperationUsingPATCHOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this modify patch user operation using p a t c h o k response has a 5xx status code
func (o *ModifyPatchUserOperationUsingPATCHOK) IsServerError() bool {
	return false
}

// IsCode returns true when this modify patch user operation using p a t c h o k response a status code equal to that given
func (o *ModifyPatchUserOperationUsingPATCHOK) IsCode(code int) bool {
	return code == 200
}

func (o *ModifyPatchUserOperationUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/user-operations/{id}][%d] modifyPatchUserOperationUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *ModifyPatchUserOperationUsingPATCHOK) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/user-operations/{id}][%d] modifyPatchUserOperationUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *ModifyPatchUserOperationUsingPATCHOK) GetPayload() *models.UserOpResource {
	return o.Payload
}

func (o *ModifyPatchUserOperationUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserOpResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyPatchUserOperationUsingPATCHUnauthorized creates a ModifyPatchUserOperationUsingPATCHUnauthorized with default headers values
func NewModifyPatchUserOperationUsingPATCHUnauthorized() *ModifyPatchUserOperationUsingPATCHUnauthorized {
	return &ModifyPatchUserOperationUsingPATCHUnauthorized{}
}

/*
ModifyPatchUserOperationUsingPATCHUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type ModifyPatchUserOperationUsingPATCHUnauthorized struct {
}

// IsSuccess returns true when this modify patch user operation using p a t c h unauthorized response has a 2xx status code
func (o *ModifyPatchUserOperationUsingPATCHUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify patch user operation using p a t c h unauthorized response has a 3xx status code
func (o *ModifyPatchUserOperationUsingPATCHUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify patch user operation using p a t c h unauthorized response has a 4xx status code
func (o *ModifyPatchUserOperationUsingPATCHUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify patch user operation using p a t c h unauthorized response has a 5xx status code
func (o *ModifyPatchUserOperationUsingPATCHUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this modify patch user operation using p a t c h unauthorized response a status code equal to that given
func (o *ModifyPatchUserOperationUsingPATCHUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ModifyPatchUserOperationUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/user-operations/{id}][%d] modifyPatchUserOperationUsingPATCHUnauthorized ", 401)
}

func (o *ModifyPatchUserOperationUsingPATCHUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/user-operations/{id}][%d] modifyPatchUserOperationUsingPATCHUnauthorized ", 401)
}

func (o *ModifyPatchUserOperationUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyPatchUserOperationUsingPATCHForbidden creates a ModifyPatchUserOperationUsingPATCHForbidden with default headers values
func NewModifyPatchUserOperationUsingPATCHForbidden() *ModifyPatchUserOperationUsingPATCHForbidden {
	return &ModifyPatchUserOperationUsingPATCHForbidden{}
}

/*
ModifyPatchUserOperationUsingPATCHForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ModifyPatchUserOperationUsingPATCHForbidden struct {
}

// IsSuccess returns true when this modify patch user operation using p a t c h forbidden response has a 2xx status code
func (o *ModifyPatchUserOperationUsingPATCHForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify patch user operation using p a t c h forbidden response has a 3xx status code
func (o *ModifyPatchUserOperationUsingPATCHForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify patch user operation using p a t c h forbidden response has a 4xx status code
func (o *ModifyPatchUserOperationUsingPATCHForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify patch user operation using p a t c h forbidden response has a 5xx status code
func (o *ModifyPatchUserOperationUsingPATCHForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this modify patch user operation using p a t c h forbidden response a status code equal to that given
func (o *ModifyPatchUserOperationUsingPATCHForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ModifyPatchUserOperationUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/user-operations/{id}][%d] modifyPatchUserOperationUsingPATCHForbidden ", 403)
}

func (o *ModifyPatchUserOperationUsingPATCHForbidden) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/user-operations/{id}][%d] modifyPatchUserOperationUsingPATCHForbidden ", 403)
}

func (o *ModifyPatchUserOperationUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyPatchUserOperationUsingPATCHNotFound creates a ModifyPatchUserOperationUsingPATCHNotFound with default headers values
func NewModifyPatchUserOperationUsingPATCHNotFound() *ModifyPatchUserOperationUsingPATCHNotFound {
	return &ModifyPatchUserOperationUsingPATCHNotFound{}
}

/*
ModifyPatchUserOperationUsingPATCHNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ModifyPatchUserOperationUsingPATCHNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this modify patch user operation using p a t c h not found response has a 2xx status code
func (o *ModifyPatchUserOperationUsingPATCHNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify patch user operation using p a t c h not found response has a 3xx status code
func (o *ModifyPatchUserOperationUsingPATCHNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify patch user operation using p a t c h not found response has a 4xx status code
func (o *ModifyPatchUserOperationUsingPATCHNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify patch user operation using p a t c h not found response has a 5xx status code
func (o *ModifyPatchUserOperationUsingPATCHNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this modify patch user operation using p a t c h not found response a status code equal to that given
func (o *ModifyPatchUserOperationUsingPATCHNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ModifyPatchUserOperationUsingPATCHNotFound) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/user-operations/{id}][%d] modifyPatchUserOperationUsingPATCHNotFound  %+v", 404, o.Payload)
}

func (o *ModifyPatchUserOperationUsingPATCHNotFound) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/user-operations/{id}][%d] modifyPatchUserOperationUsingPATCHNotFound  %+v", 404, o.Payload)
}

func (o *ModifyPatchUserOperationUsingPATCHNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ModifyPatchUserOperationUsingPATCHNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyPatchUserOperationUsingPATCHInternalServerError creates a ModifyPatchUserOperationUsingPATCHInternalServerError with default headers values
func NewModifyPatchUserOperationUsingPATCHInternalServerError() *ModifyPatchUserOperationUsingPATCHInternalServerError {
	return &ModifyPatchUserOperationUsingPATCHInternalServerError{}
}

/*
ModifyPatchUserOperationUsingPATCHInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ModifyPatchUserOperationUsingPATCHInternalServerError struct {
}

// IsSuccess returns true when this modify patch user operation using p a t c h internal server error response has a 2xx status code
func (o *ModifyPatchUserOperationUsingPATCHInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify patch user operation using p a t c h internal server error response has a 3xx status code
func (o *ModifyPatchUserOperationUsingPATCHInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify patch user operation using p a t c h internal server error response has a 4xx status code
func (o *ModifyPatchUserOperationUsingPATCHInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this modify patch user operation using p a t c h internal server error response has a 5xx status code
func (o *ModifyPatchUserOperationUsingPATCHInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this modify patch user operation using p a t c h internal server error response a status code equal to that given
func (o *ModifyPatchUserOperationUsingPATCHInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ModifyPatchUserOperationUsingPATCHInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/user-operations/{id}][%d] modifyPatchUserOperationUsingPATCHInternalServerError ", 500)
}

func (o *ModifyPatchUserOperationUsingPATCHInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/user-operations/{id}][%d] modifyPatchUserOperationUsingPATCHInternalServerError ", 500)
}

func (o *ModifyPatchUserOperationUsingPATCHInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
