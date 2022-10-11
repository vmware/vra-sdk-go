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

// CloneGerritListenerByNameUsingPOSTReader is a Reader for the CloneGerritListenerByNameUsingPOST structure.
type CloneGerritListenerByNameUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneGerritListenerByNameUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloneGerritListenerByNameUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCloneGerritListenerByNameUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloneGerritListenerByNameUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloneGerritListenerByNameUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloneGerritListenerByNameUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloneGerritListenerByNameUsingPOSTOK creates a CloneGerritListenerByNameUsingPOSTOK with default headers values
func NewCloneGerritListenerByNameUsingPOSTOK() *CloneGerritListenerByNameUsingPOSTOK {
	return &CloneGerritListenerByNameUsingPOSTOK{}
}

/*
CloneGerritListenerByNameUsingPOSTOK describes a response with status code 200, with default header values.

'Success' with the cloned Pipeline
*/
type CloneGerritListenerByNameUsingPOSTOK struct {
	Payload models.GerritListener
}

// IsSuccess returns true when this clone gerrit listener by name using p o s t o k response has a 2xx status code
func (o *CloneGerritListenerByNameUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this clone gerrit listener by name using p o s t o k response has a 3xx status code
func (o *CloneGerritListenerByNameUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clone gerrit listener by name using p o s t o k response has a 4xx status code
func (o *CloneGerritListenerByNameUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this clone gerrit listener by name using p o s t o k response has a 5xx status code
func (o *CloneGerritListenerByNameUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this clone gerrit listener by name using p o s t o k response a status code equal to that given
func (o *CloneGerritListenerByNameUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

func (o *CloneGerritListenerByNameUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneGerritListenerByNameUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CloneGerritListenerByNameUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneGerritListenerByNameUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CloneGerritListenerByNameUsingPOSTOK) GetPayload() models.GerritListener {
	return o.Payload
}

func (o *CloneGerritListenerByNameUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListener(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewCloneGerritListenerByNameUsingPOSTUnauthorized creates a CloneGerritListenerByNameUsingPOSTUnauthorized with default headers values
func NewCloneGerritListenerByNameUsingPOSTUnauthorized() *CloneGerritListenerByNameUsingPOSTUnauthorized {
	return &CloneGerritListenerByNameUsingPOSTUnauthorized{}
}

/*
CloneGerritListenerByNameUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type CloneGerritListenerByNameUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this clone gerrit listener by name using p o s t unauthorized response has a 2xx status code
func (o *CloneGerritListenerByNameUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this clone gerrit listener by name using p o s t unauthorized response has a 3xx status code
func (o *CloneGerritListenerByNameUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clone gerrit listener by name using p o s t unauthorized response has a 4xx status code
func (o *CloneGerritListenerByNameUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this clone gerrit listener by name using p o s t unauthorized response has a 5xx status code
func (o *CloneGerritListenerByNameUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this clone gerrit listener by name using p o s t unauthorized response a status code equal to that given
func (o *CloneGerritListenerByNameUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CloneGerritListenerByNameUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneGerritListenerByNameUsingPOSTUnauthorized ", 401)
}

func (o *CloneGerritListenerByNameUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneGerritListenerByNameUsingPOSTUnauthorized ", 401)
}

func (o *CloneGerritListenerByNameUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloneGerritListenerByNameUsingPOSTForbidden creates a CloneGerritListenerByNameUsingPOSTForbidden with default headers values
func NewCloneGerritListenerByNameUsingPOSTForbidden() *CloneGerritListenerByNameUsingPOSTForbidden {
	return &CloneGerritListenerByNameUsingPOSTForbidden{}
}

/*
CloneGerritListenerByNameUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloneGerritListenerByNameUsingPOSTForbidden struct {
}

// IsSuccess returns true when this clone gerrit listener by name using p o s t forbidden response has a 2xx status code
func (o *CloneGerritListenerByNameUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this clone gerrit listener by name using p o s t forbidden response has a 3xx status code
func (o *CloneGerritListenerByNameUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clone gerrit listener by name using p o s t forbidden response has a 4xx status code
func (o *CloneGerritListenerByNameUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this clone gerrit listener by name using p o s t forbidden response has a 5xx status code
func (o *CloneGerritListenerByNameUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this clone gerrit listener by name using p o s t forbidden response a status code equal to that given
func (o *CloneGerritListenerByNameUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CloneGerritListenerByNameUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneGerritListenerByNameUsingPOSTForbidden ", 403)
}

func (o *CloneGerritListenerByNameUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneGerritListenerByNameUsingPOSTForbidden ", 403)
}

func (o *CloneGerritListenerByNameUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloneGerritListenerByNameUsingPOSTNotFound creates a CloneGerritListenerByNameUsingPOSTNotFound with default headers values
func NewCloneGerritListenerByNameUsingPOSTNotFound() *CloneGerritListenerByNameUsingPOSTNotFound {
	return &CloneGerritListenerByNameUsingPOSTNotFound{}
}

/*
CloneGerritListenerByNameUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CloneGerritListenerByNameUsingPOSTNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this clone gerrit listener by name using p o s t not found response has a 2xx status code
func (o *CloneGerritListenerByNameUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this clone gerrit listener by name using p o s t not found response has a 3xx status code
func (o *CloneGerritListenerByNameUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clone gerrit listener by name using p o s t not found response has a 4xx status code
func (o *CloneGerritListenerByNameUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this clone gerrit listener by name using p o s t not found response has a 5xx status code
func (o *CloneGerritListenerByNameUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this clone gerrit listener by name using p o s t not found response a status code equal to that given
func (o *CloneGerritListenerByNameUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CloneGerritListenerByNameUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneGerritListenerByNameUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *CloneGerritListenerByNameUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneGerritListenerByNameUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *CloneGerritListenerByNameUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CloneGerritListenerByNameUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneGerritListenerByNameUsingPOSTInternalServerError creates a CloneGerritListenerByNameUsingPOSTInternalServerError with default headers values
func NewCloneGerritListenerByNameUsingPOSTInternalServerError() *CloneGerritListenerByNameUsingPOSTInternalServerError {
	return &CloneGerritListenerByNameUsingPOSTInternalServerError{}
}

/*
CloneGerritListenerByNameUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CloneGerritListenerByNameUsingPOSTInternalServerError struct {
}

// IsSuccess returns true when this clone gerrit listener by name using p o s t internal server error response has a 2xx status code
func (o *CloneGerritListenerByNameUsingPOSTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this clone gerrit listener by name using p o s t internal server error response has a 3xx status code
func (o *CloneGerritListenerByNameUsingPOSTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clone gerrit listener by name using p o s t internal server error response has a 4xx status code
func (o *CloneGerritListenerByNameUsingPOSTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this clone gerrit listener by name using p o s t internal server error response has a 5xx status code
func (o *CloneGerritListenerByNameUsingPOSTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this clone gerrit listener by name using p o s t internal server error response a status code equal to that given
func (o *CloneGerritListenerByNameUsingPOSTInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CloneGerritListenerByNameUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneGerritListenerByNameUsingPOSTInternalServerError ", 500)
}

func (o *CloneGerritListenerByNameUsingPOSTInternalServerError) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneGerritListenerByNameUsingPOSTInternalServerError ", 500)
}

func (o *CloneGerritListenerByNameUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
