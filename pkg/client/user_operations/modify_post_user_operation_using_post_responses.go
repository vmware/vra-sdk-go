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

// ModifyPostUserOperationUsingPOSTReader is a Reader for the ModifyPostUserOperationUsingPOST structure.
type ModifyPostUserOperationUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyPostUserOperationUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModifyPostUserOperationUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewModifyPostUserOperationUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewModifyPostUserOperationUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewModifyPostUserOperationUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewModifyPostUserOperationUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewModifyPostUserOperationUsingPOSTOK creates a ModifyPostUserOperationUsingPOSTOK with default headers values
func NewModifyPostUserOperationUsingPOSTOK() *ModifyPostUserOperationUsingPOSTOK {
	return &ModifyPostUserOperationUsingPOSTOK{}
}

/*
ModifyPostUserOperationUsingPOSTOK describes a response with status code 200, with default header values.

'Success' with the modified User Operation
*/
type ModifyPostUserOperationUsingPOSTOK struct {
	Payload *models.UserOpResource
}

// IsSuccess returns true when this modify post user operation using p o s t o k response has a 2xx status code
func (o *ModifyPostUserOperationUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this modify post user operation using p o s t o k response has a 3xx status code
func (o *ModifyPostUserOperationUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify post user operation using p o s t o k response has a 4xx status code
func (o *ModifyPostUserOperationUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this modify post user operation using p o s t o k response has a 5xx status code
func (o *ModifyPostUserOperationUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this modify post user operation using p o s t o k response a status code equal to that given
func (o *ModifyPostUserOperationUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

func (o *ModifyPostUserOperationUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/user-operations/{id}][%d] modifyPostUserOperationUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ModifyPostUserOperationUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /codestream/api/user-operations/{id}][%d] modifyPostUserOperationUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ModifyPostUserOperationUsingPOSTOK) GetPayload() *models.UserOpResource {
	return o.Payload
}

func (o *ModifyPostUserOperationUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserOpResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyPostUserOperationUsingPOSTUnauthorized creates a ModifyPostUserOperationUsingPOSTUnauthorized with default headers values
func NewModifyPostUserOperationUsingPOSTUnauthorized() *ModifyPostUserOperationUsingPOSTUnauthorized {
	return &ModifyPostUserOperationUsingPOSTUnauthorized{}
}

/*
ModifyPostUserOperationUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type ModifyPostUserOperationUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this modify post user operation using p o s t unauthorized response has a 2xx status code
func (o *ModifyPostUserOperationUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify post user operation using p o s t unauthorized response has a 3xx status code
func (o *ModifyPostUserOperationUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify post user operation using p o s t unauthorized response has a 4xx status code
func (o *ModifyPostUserOperationUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify post user operation using p o s t unauthorized response has a 5xx status code
func (o *ModifyPostUserOperationUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this modify post user operation using p o s t unauthorized response a status code equal to that given
func (o *ModifyPostUserOperationUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ModifyPostUserOperationUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/user-operations/{id}][%d] modifyPostUserOperationUsingPOSTUnauthorized ", 401)
}

func (o *ModifyPostUserOperationUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /codestream/api/user-operations/{id}][%d] modifyPostUserOperationUsingPOSTUnauthorized ", 401)
}

func (o *ModifyPostUserOperationUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyPostUserOperationUsingPOSTForbidden creates a ModifyPostUserOperationUsingPOSTForbidden with default headers values
func NewModifyPostUserOperationUsingPOSTForbidden() *ModifyPostUserOperationUsingPOSTForbidden {
	return &ModifyPostUserOperationUsingPOSTForbidden{}
}

/*
ModifyPostUserOperationUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ModifyPostUserOperationUsingPOSTForbidden struct {
}

// IsSuccess returns true when this modify post user operation using p o s t forbidden response has a 2xx status code
func (o *ModifyPostUserOperationUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify post user operation using p o s t forbidden response has a 3xx status code
func (o *ModifyPostUserOperationUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify post user operation using p o s t forbidden response has a 4xx status code
func (o *ModifyPostUserOperationUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify post user operation using p o s t forbidden response has a 5xx status code
func (o *ModifyPostUserOperationUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this modify post user operation using p o s t forbidden response a status code equal to that given
func (o *ModifyPostUserOperationUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ModifyPostUserOperationUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/user-operations/{id}][%d] modifyPostUserOperationUsingPOSTForbidden ", 403)
}

func (o *ModifyPostUserOperationUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /codestream/api/user-operations/{id}][%d] modifyPostUserOperationUsingPOSTForbidden ", 403)
}

func (o *ModifyPostUserOperationUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyPostUserOperationUsingPOSTNotFound creates a ModifyPostUserOperationUsingPOSTNotFound with default headers values
func NewModifyPostUserOperationUsingPOSTNotFound() *ModifyPostUserOperationUsingPOSTNotFound {
	return &ModifyPostUserOperationUsingPOSTNotFound{}
}

/*
ModifyPostUserOperationUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ModifyPostUserOperationUsingPOSTNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this modify post user operation using p o s t not found response has a 2xx status code
func (o *ModifyPostUserOperationUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify post user operation using p o s t not found response has a 3xx status code
func (o *ModifyPostUserOperationUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify post user operation using p o s t not found response has a 4xx status code
func (o *ModifyPostUserOperationUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify post user operation using p o s t not found response has a 5xx status code
func (o *ModifyPostUserOperationUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this modify post user operation using p o s t not found response a status code equal to that given
func (o *ModifyPostUserOperationUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ModifyPostUserOperationUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/user-operations/{id}][%d] modifyPostUserOperationUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *ModifyPostUserOperationUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /codestream/api/user-operations/{id}][%d] modifyPostUserOperationUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *ModifyPostUserOperationUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ModifyPostUserOperationUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyPostUserOperationUsingPOSTInternalServerError creates a ModifyPostUserOperationUsingPOSTInternalServerError with default headers values
func NewModifyPostUserOperationUsingPOSTInternalServerError() *ModifyPostUserOperationUsingPOSTInternalServerError {
	return &ModifyPostUserOperationUsingPOSTInternalServerError{}
}

/*
ModifyPostUserOperationUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ModifyPostUserOperationUsingPOSTInternalServerError struct {
}

// IsSuccess returns true when this modify post user operation using p o s t internal server error response has a 2xx status code
func (o *ModifyPostUserOperationUsingPOSTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify post user operation using p o s t internal server error response has a 3xx status code
func (o *ModifyPostUserOperationUsingPOSTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify post user operation using p o s t internal server error response has a 4xx status code
func (o *ModifyPostUserOperationUsingPOSTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this modify post user operation using p o s t internal server error response has a 5xx status code
func (o *ModifyPostUserOperationUsingPOSTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this modify post user operation using p o s t internal server error response a status code equal to that given
func (o *ModifyPostUserOperationUsingPOSTInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ModifyPostUserOperationUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/user-operations/{id}][%d] modifyPostUserOperationUsingPOSTInternalServerError ", 500)
}

func (o *ModifyPostUserOperationUsingPOSTInternalServerError) String() string {
	return fmt.Sprintf("[POST /codestream/api/user-operations/{id}][%d] modifyPostUserOperationUsingPOSTInternalServerError ", 500)
}

func (o *ModifyPostUserOperationUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
