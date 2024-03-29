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

// DeleteGerritListenerByNameUsingDELETEReader is a Reader for the DeleteGerritListenerByNameUsingDELETE structure.
type DeleteGerritListenerByNameUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGerritListenerByNameUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteGerritListenerByNameUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteGerritListenerByNameUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteGerritListenerByNameUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteGerritListenerByNameUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteGerritListenerByNameUsingDELETEInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteGerritListenerByNameUsingDELETEOK creates a DeleteGerritListenerByNameUsingDELETEOK with default headers values
func NewDeleteGerritListenerByNameUsingDELETEOK() *DeleteGerritListenerByNameUsingDELETEOK {
	return &DeleteGerritListenerByNameUsingDELETEOK{}
}

/*
DeleteGerritListenerByNameUsingDELETEOK describes a response with status code 200, with default header values.

'Success' with Gerrit Listener Delete
*/
type DeleteGerritListenerByNameUsingDELETEOK struct {
	Payload models.GerritListener
}

// IsSuccess returns true when this delete gerrit listener by name using d e l e t e o k response has a 2xx status code
func (o *DeleteGerritListenerByNameUsingDELETEOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete gerrit listener by name using d e l e t e o k response has a 3xx status code
func (o *DeleteGerritListenerByNameUsingDELETEOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete gerrit listener by name using d e l e t e o k response has a 4xx status code
func (o *DeleteGerritListenerByNameUsingDELETEOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete gerrit listener by name using d e l e t e o k response has a 5xx status code
func (o *DeleteGerritListenerByNameUsingDELETEOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete gerrit listener by name using d e l e t e o k response a status code equal to that given
func (o *DeleteGerritListenerByNameUsingDELETEOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteGerritListenerByNameUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteGerritListenerByNameUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteGerritListenerByNameUsingDELETEOK) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteGerritListenerByNameUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteGerritListenerByNameUsingDELETEOK) GetPayload() models.GerritListener {
	return o.Payload
}

func (o *DeleteGerritListenerByNameUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListener(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeleteGerritListenerByNameUsingDELETEUnauthorized creates a DeleteGerritListenerByNameUsingDELETEUnauthorized with default headers values
func NewDeleteGerritListenerByNameUsingDELETEUnauthorized() *DeleteGerritListenerByNameUsingDELETEUnauthorized {
	return &DeleteGerritListenerByNameUsingDELETEUnauthorized{}
}

/*
DeleteGerritListenerByNameUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type DeleteGerritListenerByNameUsingDELETEUnauthorized struct {
}

// IsSuccess returns true when this delete gerrit listener by name using d e l e t e unauthorized response has a 2xx status code
func (o *DeleteGerritListenerByNameUsingDELETEUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete gerrit listener by name using d e l e t e unauthorized response has a 3xx status code
func (o *DeleteGerritListenerByNameUsingDELETEUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete gerrit listener by name using d e l e t e unauthorized response has a 4xx status code
func (o *DeleteGerritListenerByNameUsingDELETEUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete gerrit listener by name using d e l e t e unauthorized response has a 5xx status code
func (o *DeleteGerritListenerByNameUsingDELETEUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete gerrit listener by name using d e l e t e unauthorized response a status code equal to that given
func (o *DeleteGerritListenerByNameUsingDELETEUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteGerritListenerByNameUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteGerritListenerByNameUsingDELETEUnauthorized ", 401)
}

func (o *DeleteGerritListenerByNameUsingDELETEUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteGerritListenerByNameUsingDELETEUnauthorized ", 401)
}

func (o *DeleteGerritListenerByNameUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGerritListenerByNameUsingDELETEForbidden creates a DeleteGerritListenerByNameUsingDELETEForbidden with default headers values
func NewDeleteGerritListenerByNameUsingDELETEForbidden() *DeleteGerritListenerByNameUsingDELETEForbidden {
	return &DeleteGerritListenerByNameUsingDELETEForbidden{}
}

/*
DeleteGerritListenerByNameUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteGerritListenerByNameUsingDELETEForbidden struct {
}

// IsSuccess returns true when this delete gerrit listener by name using d e l e t e forbidden response has a 2xx status code
func (o *DeleteGerritListenerByNameUsingDELETEForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete gerrit listener by name using d e l e t e forbidden response has a 3xx status code
func (o *DeleteGerritListenerByNameUsingDELETEForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete gerrit listener by name using d e l e t e forbidden response has a 4xx status code
func (o *DeleteGerritListenerByNameUsingDELETEForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete gerrit listener by name using d e l e t e forbidden response has a 5xx status code
func (o *DeleteGerritListenerByNameUsingDELETEForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete gerrit listener by name using d e l e t e forbidden response a status code equal to that given
func (o *DeleteGerritListenerByNameUsingDELETEForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteGerritListenerByNameUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteGerritListenerByNameUsingDELETEForbidden ", 403)
}

func (o *DeleteGerritListenerByNameUsingDELETEForbidden) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteGerritListenerByNameUsingDELETEForbidden ", 403)
}

func (o *DeleteGerritListenerByNameUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGerritListenerByNameUsingDELETENotFound creates a DeleteGerritListenerByNameUsingDELETENotFound with default headers values
func NewDeleteGerritListenerByNameUsingDELETENotFound() *DeleteGerritListenerByNameUsingDELETENotFound {
	return &DeleteGerritListenerByNameUsingDELETENotFound{}
}

/*
DeleteGerritListenerByNameUsingDELETENotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteGerritListenerByNameUsingDELETENotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete gerrit listener by name using d e l e t e not found response has a 2xx status code
func (o *DeleteGerritListenerByNameUsingDELETENotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete gerrit listener by name using d e l e t e not found response has a 3xx status code
func (o *DeleteGerritListenerByNameUsingDELETENotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete gerrit listener by name using d e l e t e not found response has a 4xx status code
func (o *DeleteGerritListenerByNameUsingDELETENotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete gerrit listener by name using d e l e t e not found response has a 5xx status code
func (o *DeleteGerritListenerByNameUsingDELETENotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete gerrit listener by name using d e l e t e not found response a status code equal to that given
func (o *DeleteGerritListenerByNameUsingDELETENotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteGerritListenerByNameUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteGerritListenerByNameUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteGerritListenerByNameUsingDELETENotFound) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteGerritListenerByNameUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteGerritListenerByNameUsingDELETENotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteGerritListenerByNameUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGerritListenerByNameUsingDELETEInternalServerError creates a DeleteGerritListenerByNameUsingDELETEInternalServerError with default headers values
func NewDeleteGerritListenerByNameUsingDELETEInternalServerError() *DeleteGerritListenerByNameUsingDELETEInternalServerError {
	return &DeleteGerritListenerByNameUsingDELETEInternalServerError{}
}

/*
DeleteGerritListenerByNameUsingDELETEInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DeleteGerritListenerByNameUsingDELETEInternalServerError struct {
}

// IsSuccess returns true when this delete gerrit listener by name using d e l e t e internal server error response has a 2xx status code
func (o *DeleteGerritListenerByNameUsingDELETEInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete gerrit listener by name using d e l e t e internal server error response has a 3xx status code
func (o *DeleteGerritListenerByNameUsingDELETEInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete gerrit listener by name using d e l e t e internal server error response has a 4xx status code
func (o *DeleteGerritListenerByNameUsingDELETEInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete gerrit listener by name using d e l e t e internal server error response has a 5xx status code
func (o *DeleteGerritListenerByNameUsingDELETEInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete gerrit listener by name using d e l e t e internal server error response a status code equal to that given
func (o *DeleteGerritListenerByNameUsingDELETEInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteGerritListenerByNameUsingDELETEInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteGerritListenerByNameUsingDELETEInternalServerError ", 500)
}

func (o *DeleteGerritListenerByNameUsingDELETEInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteGerritListenerByNameUsingDELETEInternalServerError ", 500)
}

func (o *DeleteGerritListenerByNameUsingDELETEInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
