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

// ValidateGerritListenerUsingPOSTReader is a Reader for the ValidateGerritListenerUsingPOST structure.
type ValidateGerritListenerUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateGerritListenerUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateGerritListenerUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewValidateGerritListenerUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewValidateGerritListenerUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewValidateGerritListenerUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewValidateGerritListenerUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewValidateGerritListenerUsingPOSTOK creates a ValidateGerritListenerUsingPOSTOK with default headers values
func NewValidateGerritListenerUsingPOSTOK() *ValidateGerritListenerUsingPOSTOK {
	return &ValidateGerritListenerUsingPOSTOK{}
}

/*
ValidateGerritListenerUsingPOSTOK describes a response with status code 200, with default header values.

'Success' with Gerrit Listener Validation
*/
type ValidateGerritListenerUsingPOSTOK struct {
	Payload models.GerritListener
}

// IsSuccess returns true when this validate gerrit listener using p o s t o k response has a 2xx status code
func (o *ValidateGerritListenerUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate gerrit listener using p o s t o k response has a 3xx status code
func (o *ValidateGerritListenerUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate gerrit listener using p o s t o k response has a 4xx status code
func (o *ValidateGerritListenerUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate gerrit listener using p o s t o k response has a 5xx status code
func (o *ValidateGerritListenerUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate gerrit listener using p o s t o k response a status code equal to that given
func (o *ValidateGerritListenerUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

func (o *ValidateGerritListenerUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/validate][%d] validateGerritListenerUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ValidateGerritListenerUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/validate][%d] validateGerritListenerUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ValidateGerritListenerUsingPOSTOK) GetPayload() models.GerritListener {
	return o.Payload
}

func (o *ValidateGerritListenerUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListener(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewValidateGerritListenerUsingPOSTUnauthorized creates a ValidateGerritListenerUsingPOSTUnauthorized with default headers values
func NewValidateGerritListenerUsingPOSTUnauthorized() *ValidateGerritListenerUsingPOSTUnauthorized {
	return &ValidateGerritListenerUsingPOSTUnauthorized{}
}

/*
ValidateGerritListenerUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type ValidateGerritListenerUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this validate gerrit listener using p o s t unauthorized response has a 2xx status code
func (o *ValidateGerritListenerUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate gerrit listener using p o s t unauthorized response has a 3xx status code
func (o *ValidateGerritListenerUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate gerrit listener using p o s t unauthorized response has a 4xx status code
func (o *ValidateGerritListenerUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate gerrit listener using p o s t unauthorized response has a 5xx status code
func (o *ValidateGerritListenerUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this validate gerrit listener using p o s t unauthorized response a status code equal to that given
func (o *ValidateGerritListenerUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ValidateGerritListenerUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/validate][%d] validateGerritListenerUsingPOSTUnauthorized ", 401)
}

func (o *ValidateGerritListenerUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/validate][%d] validateGerritListenerUsingPOSTUnauthorized ", 401)
}

func (o *ValidateGerritListenerUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateGerritListenerUsingPOSTForbidden creates a ValidateGerritListenerUsingPOSTForbidden with default headers values
func NewValidateGerritListenerUsingPOSTForbidden() *ValidateGerritListenerUsingPOSTForbidden {
	return &ValidateGerritListenerUsingPOSTForbidden{}
}

/*
ValidateGerritListenerUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ValidateGerritListenerUsingPOSTForbidden struct {
}

// IsSuccess returns true when this validate gerrit listener using p o s t forbidden response has a 2xx status code
func (o *ValidateGerritListenerUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate gerrit listener using p o s t forbidden response has a 3xx status code
func (o *ValidateGerritListenerUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate gerrit listener using p o s t forbidden response has a 4xx status code
func (o *ValidateGerritListenerUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate gerrit listener using p o s t forbidden response has a 5xx status code
func (o *ValidateGerritListenerUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this validate gerrit listener using p o s t forbidden response a status code equal to that given
func (o *ValidateGerritListenerUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ValidateGerritListenerUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/validate][%d] validateGerritListenerUsingPOSTForbidden ", 403)
}

func (o *ValidateGerritListenerUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/validate][%d] validateGerritListenerUsingPOSTForbidden ", 403)
}

func (o *ValidateGerritListenerUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateGerritListenerUsingPOSTNotFound creates a ValidateGerritListenerUsingPOSTNotFound with default headers values
func NewValidateGerritListenerUsingPOSTNotFound() *ValidateGerritListenerUsingPOSTNotFound {
	return &ValidateGerritListenerUsingPOSTNotFound{}
}

/*
ValidateGerritListenerUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ValidateGerritListenerUsingPOSTNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this validate gerrit listener using p o s t not found response has a 2xx status code
func (o *ValidateGerritListenerUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate gerrit listener using p o s t not found response has a 3xx status code
func (o *ValidateGerritListenerUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate gerrit listener using p o s t not found response has a 4xx status code
func (o *ValidateGerritListenerUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate gerrit listener using p o s t not found response has a 5xx status code
func (o *ValidateGerritListenerUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this validate gerrit listener using p o s t not found response a status code equal to that given
func (o *ValidateGerritListenerUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ValidateGerritListenerUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/validate][%d] validateGerritListenerUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *ValidateGerritListenerUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/validate][%d] validateGerritListenerUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *ValidateGerritListenerUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateGerritListenerUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateGerritListenerUsingPOSTInternalServerError creates a ValidateGerritListenerUsingPOSTInternalServerError with default headers values
func NewValidateGerritListenerUsingPOSTInternalServerError() *ValidateGerritListenerUsingPOSTInternalServerError {
	return &ValidateGerritListenerUsingPOSTInternalServerError{}
}

/*
ValidateGerritListenerUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ValidateGerritListenerUsingPOSTInternalServerError struct {
}

// IsSuccess returns true when this validate gerrit listener using p o s t internal server error response has a 2xx status code
func (o *ValidateGerritListenerUsingPOSTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate gerrit listener using p o s t internal server error response has a 3xx status code
func (o *ValidateGerritListenerUsingPOSTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate gerrit listener using p o s t internal server error response has a 4xx status code
func (o *ValidateGerritListenerUsingPOSTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate gerrit listener using p o s t internal server error response has a 5xx status code
func (o *ValidateGerritListenerUsingPOSTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this validate gerrit listener using p o s t internal server error response a status code equal to that given
func (o *ValidateGerritListenerUsingPOSTInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ValidateGerritListenerUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/validate][%d] validateGerritListenerUsingPOSTInternalServerError ", 500)
}

func (o *ValidateGerritListenerUsingPOSTInternalServerError) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/validate][%d] validateGerritListenerUsingPOSTInternalServerError ", 500)
}

func (o *ValidateGerritListenerUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
