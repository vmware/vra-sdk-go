// Code generated by go-swagger; DO NOT EDIT.

package resource_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// SubmitResourceActionRequestUsingPOST5Reader is a Reader for the SubmitResourceActionRequestUsingPOST5 structure.
type SubmitResourceActionRequestUsingPOST5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubmitResourceActionRequestUsingPOST5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubmitResourceActionRequestUsingPOST5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSubmitResourceActionRequestUsingPOST5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubmitResourceActionRequestUsingPOST5Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSubmitResourceActionRequestUsingPOST5NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSubmitResourceActionRequestUsingPOST5Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubmitResourceActionRequestUsingPOST5OK creates a SubmitResourceActionRequestUsingPOST5OK with default headers values
func NewSubmitResourceActionRequestUsingPOST5OK() *SubmitResourceActionRequestUsingPOST5OK {
	return &SubmitResourceActionRequestUsingPOST5OK{}
}

/*
SubmitResourceActionRequestUsingPOST5OK describes a response with status code 200, with default header values.

OK
*/
type SubmitResourceActionRequestUsingPOST5OK struct {
	Payload *models.Request
}

// IsSuccess returns true when this submit resource action request using p o s t5 o k response has a 2xx status code
func (o *SubmitResourceActionRequestUsingPOST5OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this submit resource action request using p o s t5 o k response has a 3xx status code
func (o *SubmitResourceActionRequestUsingPOST5OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit resource action request using p o s t5 o k response has a 4xx status code
func (o *SubmitResourceActionRequestUsingPOST5OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this submit resource action request using p o s t5 o k response has a 5xx status code
func (o *SubmitResourceActionRequestUsingPOST5OK) IsServerError() bool {
	return false
}

// IsCode returns true when this submit resource action request using p o s t5 o k response a status code equal to that given
func (o *SubmitResourceActionRequestUsingPOST5OK) IsCode(code int) bool {
	return code == 200
}

func (o *SubmitResourceActionRequestUsingPOST5OK) Error() string {
	return fmt.Sprintf("[POST /deployment/api/resources/{resourceId}/requests][%d] submitResourceActionRequestUsingPOST5OK  %+v", 200, o.Payload)
}

func (o *SubmitResourceActionRequestUsingPOST5OK) String() string {
	return fmt.Sprintf("[POST /deployment/api/resources/{resourceId}/requests][%d] submitResourceActionRequestUsingPOST5OK  %+v", 200, o.Payload)
}

func (o *SubmitResourceActionRequestUsingPOST5OK) GetPayload() *models.Request {
	return o.Payload
}

func (o *SubmitResourceActionRequestUsingPOST5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Request)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitResourceActionRequestUsingPOST5Unauthorized creates a SubmitResourceActionRequestUsingPOST5Unauthorized with default headers values
func NewSubmitResourceActionRequestUsingPOST5Unauthorized() *SubmitResourceActionRequestUsingPOST5Unauthorized {
	return &SubmitResourceActionRequestUsingPOST5Unauthorized{}
}

/*
SubmitResourceActionRequestUsingPOST5Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SubmitResourceActionRequestUsingPOST5Unauthorized struct {
}

// IsSuccess returns true when this submit resource action request using p o s t5 unauthorized response has a 2xx status code
func (o *SubmitResourceActionRequestUsingPOST5Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit resource action request using p o s t5 unauthorized response has a 3xx status code
func (o *SubmitResourceActionRequestUsingPOST5Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit resource action request using p o s t5 unauthorized response has a 4xx status code
func (o *SubmitResourceActionRequestUsingPOST5Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit resource action request using p o s t5 unauthorized response has a 5xx status code
func (o *SubmitResourceActionRequestUsingPOST5Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this submit resource action request using p o s t5 unauthorized response a status code equal to that given
func (o *SubmitResourceActionRequestUsingPOST5Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SubmitResourceActionRequestUsingPOST5Unauthorized) Error() string {
	return fmt.Sprintf("[POST /deployment/api/resources/{resourceId}/requests][%d] submitResourceActionRequestUsingPOST5Unauthorized ", 401)
}

func (o *SubmitResourceActionRequestUsingPOST5Unauthorized) String() string {
	return fmt.Sprintf("[POST /deployment/api/resources/{resourceId}/requests][%d] submitResourceActionRequestUsingPOST5Unauthorized ", 401)
}

func (o *SubmitResourceActionRequestUsingPOST5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitResourceActionRequestUsingPOST5Forbidden creates a SubmitResourceActionRequestUsingPOST5Forbidden with default headers values
func NewSubmitResourceActionRequestUsingPOST5Forbidden() *SubmitResourceActionRequestUsingPOST5Forbidden {
	return &SubmitResourceActionRequestUsingPOST5Forbidden{}
}

/*
SubmitResourceActionRequestUsingPOST5Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SubmitResourceActionRequestUsingPOST5Forbidden struct {
}

// IsSuccess returns true when this submit resource action request using p o s t5 forbidden response has a 2xx status code
func (o *SubmitResourceActionRequestUsingPOST5Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit resource action request using p o s t5 forbidden response has a 3xx status code
func (o *SubmitResourceActionRequestUsingPOST5Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit resource action request using p o s t5 forbidden response has a 4xx status code
func (o *SubmitResourceActionRequestUsingPOST5Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit resource action request using p o s t5 forbidden response has a 5xx status code
func (o *SubmitResourceActionRequestUsingPOST5Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this submit resource action request using p o s t5 forbidden response a status code equal to that given
func (o *SubmitResourceActionRequestUsingPOST5Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SubmitResourceActionRequestUsingPOST5Forbidden) Error() string {
	return fmt.Sprintf("[POST /deployment/api/resources/{resourceId}/requests][%d] submitResourceActionRequestUsingPOST5Forbidden ", 403)
}

func (o *SubmitResourceActionRequestUsingPOST5Forbidden) String() string {
	return fmt.Sprintf("[POST /deployment/api/resources/{resourceId}/requests][%d] submitResourceActionRequestUsingPOST5Forbidden ", 403)
}

func (o *SubmitResourceActionRequestUsingPOST5Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubmitResourceActionRequestUsingPOST5NotFound creates a SubmitResourceActionRequestUsingPOST5NotFound with default headers values
func NewSubmitResourceActionRequestUsingPOST5NotFound() *SubmitResourceActionRequestUsingPOST5NotFound {
	return &SubmitResourceActionRequestUsingPOST5NotFound{}
}

/*
SubmitResourceActionRequestUsingPOST5NotFound describes a response with status code 404, with default header values.

Not Found
*/
type SubmitResourceActionRequestUsingPOST5NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this submit resource action request using p o s t5 not found response has a 2xx status code
func (o *SubmitResourceActionRequestUsingPOST5NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit resource action request using p o s t5 not found response has a 3xx status code
func (o *SubmitResourceActionRequestUsingPOST5NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit resource action request using p o s t5 not found response has a 4xx status code
func (o *SubmitResourceActionRequestUsingPOST5NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit resource action request using p o s t5 not found response has a 5xx status code
func (o *SubmitResourceActionRequestUsingPOST5NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this submit resource action request using p o s t5 not found response a status code equal to that given
func (o *SubmitResourceActionRequestUsingPOST5NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SubmitResourceActionRequestUsingPOST5NotFound) Error() string {
	return fmt.Sprintf("[POST /deployment/api/resources/{resourceId}/requests][%d] submitResourceActionRequestUsingPOST5NotFound  %+v", 404, o.Payload)
}

func (o *SubmitResourceActionRequestUsingPOST5NotFound) String() string {
	return fmt.Sprintf("[POST /deployment/api/resources/{resourceId}/requests][%d] submitResourceActionRequestUsingPOST5NotFound  %+v", 404, o.Payload)
}

func (o *SubmitResourceActionRequestUsingPOST5NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SubmitResourceActionRequestUsingPOST5NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubmitResourceActionRequestUsingPOST5Conflict creates a SubmitResourceActionRequestUsingPOST5Conflict with default headers values
func NewSubmitResourceActionRequestUsingPOST5Conflict() *SubmitResourceActionRequestUsingPOST5Conflict {
	return &SubmitResourceActionRequestUsingPOST5Conflict{}
}

/*
SubmitResourceActionRequestUsingPOST5Conflict describes a response with status code 409, with default header values.

Conflict
*/
type SubmitResourceActionRequestUsingPOST5Conflict struct {
}

// IsSuccess returns true when this submit resource action request using p o s t5 conflict response has a 2xx status code
func (o *SubmitResourceActionRequestUsingPOST5Conflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this submit resource action request using p o s t5 conflict response has a 3xx status code
func (o *SubmitResourceActionRequestUsingPOST5Conflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this submit resource action request using p o s t5 conflict response has a 4xx status code
func (o *SubmitResourceActionRequestUsingPOST5Conflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this submit resource action request using p o s t5 conflict response has a 5xx status code
func (o *SubmitResourceActionRequestUsingPOST5Conflict) IsServerError() bool {
	return false
}

// IsCode returns true when this submit resource action request using p o s t5 conflict response a status code equal to that given
func (o *SubmitResourceActionRequestUsingPOST5Conflict) IsCode(code int) bool {
	return code == 409
}

func (o *SubmitResourceActionRequestUsingPOST5Conflict) Error() string {
	return fmt.Sprintf("[POST /deployment/api/resources/{resourceId}/requests][%d] submitResourceActionRequestUsingPOST5Conflict ", 409)
}

func (o *SubmitResourceActionRequestUsingPOST5Conflict) String() string {
	return fmt.Sprintf("[POST /deployment/api/resources/{resourceId}/requests][%d] submitResourceActionRequestUsingPOST5Conflict ", 409)
}

func (o *SubmitResourceActionRequestUsingPOST5Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
