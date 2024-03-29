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

// DeleteGerritListenerByIDUsingDELETEReader is a Reader for the DeleteGerritListenerByIDUsingDELETE structure.
type DeleteGerritListenerByIDUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGerritListenerByIDUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteGerritListenerByIDUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteGerritListenerByIDUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteGerritListenerByIDUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteGerritListenerByIDUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteGerritListenerByIDUsingDELETEInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteGerritListenerByIDUsingDELETEOK creates a DeleteGerritListenerByIDUsingDELETEOK with default headers values
func NewDeleteGerritListenerByIDUsingDELETEOK() *DeleteGerritListenerByIDUsingDELETEOK {
	return &DeleteGerritListenerByIDUsingDELETEOK{}
}

/*
DeleteGerritListenerByIDUsingDELETEOK describes a response with status code 200, with default header values.

'Success' with Gerrit Listener Delete
*/
type DeleteGerritListenerByIDUsingDELETEOK struct {
	Payload models.GerritListener
}

// IsSuccess returns true when this delete gerrit listener by Id using d e l e t e o k response has a 2xx status code
func (o *DeleteGerritListenerByIDUsingDELETEOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete gerrit listener by Id using d e l e t e o k response has a 3xx status code
func (o *DeleteGerritListenerByIDUsingDELETEOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete gerrit listener by Id using d e l e t e o k response has a 4xx status code
func (o *DeleteGerritListenerByIDUsingDELETEOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete gerrit listener by Id using d e l e t e o k response has a 5xx status code
func (o *DeleteGerritListenerByIDUsingDELETEOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete gerrit listener by Id using d e l e t e o k response a status code equal to that given
func (o *DeleteGerritListenerByIDUsingDELETEOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteGerritListenerByIDUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteGerritListenerByIdUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteGerritListenerByIDUsingDELETEOK) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteGerritListenerByIdUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteGerritListenerByIDUsingDELETEOK) GetPayload() models.GerritListener {
	return o.Payload
}

func (o *DeleteGerritListenerByIDUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListener(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeleteGerritListenerByIDUsingDELETEUnauthorized creates a DeleteGerritListenerByIDUsingDELETEUnauthorized with default headers values
func NewDeleteGerritListenerByIDUsingDELETEUnauthorized() *DeleteGerritListenerByIDUsingDELETEUnauthorized {
	return &DeleteGerritListenerByIDUsingDELETEUnauthorized{}
}

/*
DeleteGerritListenerByIDUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type DeleteGerritListenerByIDUsingDELETEUnauthorized struct {
}

// IsSuccess returns true when this delete gerrit listener by Id using d e l e t e unauthorized response has a 2xx status code
func (o *DeleteGerritListenerByIDUsingDELETEUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete gerrit listener by Id using d e l e t e unauthorized response has a 3xx status code
func (o *DeleteGerritListenerByIDUsingDELETEUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete gerrit listener by Id using d e l e t e unauthorized response has a 4xx status code
func (o *DeleteGerritListenerByIDUsingDELETEUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete gerrit listener by Id using d e l e t e unauthorized response has a 5xx status code
func (o *DeleteGerritListenerByIDUsingDELETEUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete gerrit listener by Id using d e l e t e unauthorized response a status code equal to that given
func (o *DeleteGerritListenerByIDUsingDELETEUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteGerritListenerByIDUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteGerritListenerByIdUsingDELETEUnauthorized ", 401)
}

func (o *DeleteGerritListenerByIDUsingDELETEUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteGerritListenerByIdUsingDELETEUnauthorized ", 401)
}

func (o *DeleteGerritListenerByIDUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGerritListenerByIDUsingDELETEForbidden creates a DeleteGerritListenerByIDUsingDELETEForbidden with default headers values
func NewDeleteGerritListenerByIDUsingDELETEForbidden() *DeleteGerritListenerByIDUsingDELETEForbidden {
	return &DeleteGerritListenerByIDUsingDELETEForbidden{}
}

/*
DeleteGerritListenerByIDUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteGerritListenerByIDUsingDELETEForbidden struct {
}

// IsSuccess returns true when this delete gerrit listener by Id using d e l e t e forbidden response has a 2xx status code
func (o *DeleteGerritListenerByIDUsingDELETEForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete gerrit listener by Id using d e l e t e forbidden response has a 3xx status code
func (o *DeleteGerritListenerByIDUsingDELETEForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete gerrit listener by Id using d e l e t e forbidden response has a 4xx status code
func (o *DeleteGerritListenerByIDUsingDELETEForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete gerrit listener by Id using d e l e t e forbidden response has a 5xx status code
func (o *DeleteGerritListenerByIDUsingDELETEForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete gerrit listener by Id using d e l e t e forbidden response a status code equal to that given
func (o *DeleteGerritListenerByIDUsingDELETEForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteGerritListenerByIDUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteGerritListenerByIdUsingDELETEForbidden ", 403)
}

func (o *DeleteGerritListenerByIDUsingDELETEForbidden) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteGerritListenerByIdUsingDELETEForbidden ", 403)
}

func (o *DeleteGerritListenerByIDUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGerritListenerByIDUsingDELETENotFound creates a DeleteGerritListenerByIDUsingDELETENotFound with default headers values
func NewDeleteGerritListenerByIDUsingDELETENotFound() *DeleteGerritListenerByIDUsingDELETENotFound {
	return &DeleteGerritListenerByIDUsingDELETENotFound{}
}

/*
DeleteGerritListenerByIDUsingDELETENotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteGerritListenerByIDUsingDELETENotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete gerrit listener by Id using d e l e t e not found response has a 2xx status code
func (o *DeleteGerritListenerByIDUsingDELETENotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete gerrit listener by Id using d e l e t e not found response has a 3xx status code
func (o *DeleteGerritListenerByIDUsingDELETENotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete gerrit listener by Id using d e l e t e not found response has a 4xx status code
func (o *DeleteGerritListenerByIDUsingDELETENotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete gerrit listener by Id using d e l e t e not found response has a 5xx status code
func (o *DeleteGerritListenerByIDUsingDELETENotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete gerrit listener by Id using d e l e t e not found response a status code equal to that given
func (o *DeleteGerritListenerByIDUsingDELETENotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteGerritListenerByIDUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteGerritListenerByIdUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteGerritListenerByIDUsingDELETENotFound) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteGerritListenerByIdUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteGerritListenerByIDUsingDELETENotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteGerritListenerByIDUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGerritListenerByIDUsingDELETEInternalServerError creates a DeleteGerritListenerByIDUsingDELETEInternalServerError with default headers values
func NewDeleteGerritListenerByIDUsingDELETEInternalServerError() *DeleteGerritListenerByIDUsingDELETEInternalServerError {
	return &DeleteGerritListenerByIDUsingDELETEInternalServerError{}
}

/*
DeleteGerritListenerByIDUsingDELETEInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DeleteGerritListenerByIDUsingDELETEInternalServerError struct {
}

// IsSuccess returns true when this delete gerrit listener by Id using d e l e t e internal server error response has a 2xx status code
func (o *DeleteGerritListenerByIDUsingDELETEInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete gerrit listener by Id using d e l e t e internal server error response has a 3xx status code
func (o *DeleteGerritListenerByIDUsingDELETEInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete gerrit listener by Id using d e l e t e internal server error response has a 4xx status code
func (o *DeleteGerritListenerByIDUsingDELETEInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete gerrit listener by Id using d e l e t e internal server error response has a 5xx status code
func (o *DeleteGerritListenerByIDUsingDELETEInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete gerrit listener by Id using d e l e t e internal server error response a status code equal to that given
func (o *DeleteGerritListenerByIDUsingDELETEInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteGerritListenerByIDUsingDELETEInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteGerritListenerByIdUsingDELETEInternalServerError ", 500)
}

func (o *DeleteGerritListenerByIDUsingDELETEInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteGerritListenerByIdUsingDELETEInternalServerError ", 500)
}

func (o *DeleteGerritListenerByIDUsingDELETEInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
