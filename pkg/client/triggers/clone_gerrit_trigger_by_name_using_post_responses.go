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

// CloneGerritTriggerByNameUsingPOSTReader is a Reader for the CloneGerritTriggerByNameUsingPOST structure.
type CloneGerritTriggerByNameUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneGerritTriggerByNameUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloneGerritTriggerByNameUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCloneGerritTriggerByNameUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloneGerritTriggerByNameUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloneGerritTriggerByNameUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloneGerritTriggerByNameUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloneGerritTriggerByNameUsingPOSTOK creates a CloneGerritTriggerByNameUsingPOSTOK with default headers values
func NewCloneGerritTriggerByNameUsingPOSTOK() *CloneGerritTriggerByNameUsingPOSTOK {
	return &CloneGerritTriggerByNameUsingPOSTOK{}
}

/*
CloneGerritTriggerByNameUsingPOSTOK describes a response with status code 200, with default header values.

'Success' with the cloned Gerrit Trigger
*/
type CloneGerritTriggerByNameUsingPOSTOK struct {
	Payload models.GerritTrigger
}

// IsSuccess returns true when this clone gerrit trigger by name using p o s t o k response has a 2xx status code
func (o *CloneGerritTriggerByNameUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this clone gerrit trigger by name using p o s t o k response has a 3xx status code
func (o *CloneGerritTriggerByNameUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clone gerrit trigger by name using p o s t o k response has a 4xx status code
func (o *CloneGerritTriggerByNameUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this clone gerrit trigger by name using p o s t o k response has a 5xx status code
func (o *CloneGerritTriggerByNameUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this clone gerrit trigger by name using p o s t o k response a status code equal to that given
func (o *CloneGerritTriggerByNameUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

func (o *CloneGerritTriggerByNameUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{project}/{name}][%d] cloneGerritTriggerByNameUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CloneGerritTriggerByNameUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{project}/{name}][%d] cloneGerritTriggerByNameUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CloneGerritTriggerByNameUsingPOSTOK) GetPayload() models.GerritTrigger {
	return o.Payload
}

func (o *CloneGerritTriggerByNameUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritTrigger(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewCloneGerritTriggerByNameUsingPOSTUnauthorized creates a CloneGerritTriggerByNameUsingPOSTUnauthorized with default headers values
func NewCloneGerritTriggerByNameUsingPOSTUnauthorized() *CloneGerritTriggerByNameUsingPOSTUnauthorized {
	return &CloneGerritTriggerByNameUsingPOSTUnauthorized{}
}

/*
CloneGerritTriggerByNameUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type CloneGerritTriggerByNameUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this clone gerrit trigger by name using p o s t unauthorized response has a 2xx status code
func (o *CloneGerritTriggerByNameUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this clone gerrit trigger by name using p o s t unauthorized response has a 3xx status code
func (o *CloneGerritTriggerByNameUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clone gerrit trigger by name using p o s t unauthorized response has a 4xx status code
func (o *CloneGerritTriggerByNameUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this clone gerrit trigger by name using p o s t unauthorized response has a 5xx status code
func (o *CloneGerritTriggerByNameUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this clone gerrit trigger by name using p o s t unauthorized response a status code equal to that given
func (o *CloneGerritTriggerByNameUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CloneGerritTriggerByNameUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{project}/{name}][%d] cloneGerritTriggerByNameUsingPOSTUnauthorized ", 401)
}

func (o *CloneGerritTriggerByNameUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{project}/{name}][%d] cloneGerritTriggerByNameUsingPOSTUnauthorized ", 401)
}

func (o *CloneGerritTriggerByNameUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloneGerritTriggerByNameUsingPOSTForbidden creates a CloneGerritTriggerByNameUsingPOSTForbidden with default headers values
func NewCloneGerritTriggerByNameUsingPOSTForbidden() *CloneGerritTriggerByNameUsingPOSTForbidden {
	return &CloneGerritTriggerByNameUsingPOSTForbidden{}
}

/*
CloneGerritTriggerByNameUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloneGerritTriggerByNameUsingPOSTForbidden struct {
}

// IsSuccess returns true when this clone gerrit trigger by name using p o s t forbidden response has a 2xx status code
func (o *CloneGerritTriggerByNameUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this clone gerrit trigger by name using p o s t forbidden response has a 3xx status code
func (o *CloneGerritTriggerByNameUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clone gerrit trigger by name using p o s t forbidden response has a 4xx status code
func (o *CloneGerritTriggerByNameUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this clone gerrit trigger by name using p o s t forbidden response has a 5xx status code
func (o *CloneGerritTriggerByNameUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this clone gerrit trigger by name using p o s t forbidden response a status code equal to that given
func (o *CloneGerritTriggerByNameUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CloneGerritTriggerByNameUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{project}/{name}][%d] cloneGerritTriggerByNameUsingPOSTForbidden ", 403)
}

func (o *CloneGerritTriggerByNameUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{project}/{name}][%d] cloneGerritTriggerByNameUsingPOSTForbidden ", 403)
}

func (o *CloneGerritTriggerByNameUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloneGerritTriggerByNameUsingPOSTNotFound creates a CloneGerritTriggerByNameUsingPOSTNotFound with default headers values
func NewCloneGerritTriggerByNameUsingPOSTNotFound() *CloneGerritTriggerByNameUsingPOSTNotFound {
	return &CloneGerritTriggerByNameUsingPOSTNotFound{}
}

/*
CloneGerritTriggerByNameUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CloneGerritTriggerByNameUsingPOSTNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this clone gerrit trigger by name using p o s t not found response has a 2xx status code
func (o *CloneGerritTriggerByNameUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this clone gerrit trigger by name using p o s t not found response has a 3xx status code
func (o *CloneGerritTriggerByNameUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clone gerrit trigger by name using p o s t not found response has a 4xx status code
func (o *CloneGerritTriggerByNameUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this clone gerrit trigger by name using p o s t not found response has a 5xx status code
func (o *CloneGerritTriggerByNameUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this clone gerrit trigger by name using p o s t not found response a status code equal to that given
func (o *CloneGerritTriggerByNameUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CloneGerritTriggerByNameUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{project}/{name}][%d] cloneGerritTriggerByNameUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *CloneGerritTriggerByNameUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{project}/{name}][%d] cloneGerritTriggerByNameUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *CloneGerritTriggerByNameUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CloneGerritTriggerByNameUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneGerritTriggerByNameUsingPOSTInternalServerError creates a CloneGerritTriggerByNameUsingPOSTInternalServerError with default headers values
func NewCloneGerritTriggerByNameUsingPOSTInternalServerError() *CloneGerritTriggerByNameUsingPOSTInternalServerError {
	return &CloneGerritTriggerByNameUsingPOSTInternalServerError{}
}

/*
CloneGerritTriggerByNameUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CloneGerritTriggerByNameUsingPOSTInternalServerError struct {
}

// IsSuccess returns true when this clone gerrit trigger by name using p o s t internal server error response has a 2xx status code
func (o *CloneGerritTriggerByNameUsingPOSTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this clone gerrit trigger by name using p o s t internal server error response has a 3xx status code
func (o *CloneGerritTriggerByNameUsingPOSTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clone gerrit trigger by name using p o s t internal server error response has a 4xx status code
func (o *CloneGerritTriggerByNameUsingPOSTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this clone gerrit trigger by name using p o s t internal server error response has a 5xx status code
func (o *CloneGerritTriggerByNameUsingPOSTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this clone gerrit trigger by name using p o s t internal server error response a status code equal to that given
func (o *CloneGerritTriggerByNameUsingPOSTInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CloneGerritTriggerByNameUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{project}/{name}][%d] cloneGerritTriggerByNameUsingPOSTInternalServerError ", 500)
}

func (o *CloneGerritTriggerByNameUsingPOSTInternalServerError) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{project}/{name}][%d] cloneGerritTriggerByNameUsingPOSTInternalServerError ", 500)
}

func (o *CloneGerritTriggerByNameUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
