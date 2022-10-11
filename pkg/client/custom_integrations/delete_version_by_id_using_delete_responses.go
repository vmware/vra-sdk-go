// Code generated by go-swagger; DO NOT EDIT.

package custom_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteVersionByIDUsingDELETEReader is a Reader for the DeleteVersionByIDUsingDELETE structure.
type DeleteVersionByIDUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVersionByIDUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVersionByIDUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteVersionByIDUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteVersionByIDUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVersionByIDUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteVersionByIDUsingDELETEInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVersionByIDUsingDELETEOK creates a DeleteVersionByIDUsingDELETEOK with default headers values
func NewDeleteVersionByIDUsingDELETEOK() *DeleteVersionByIDUsingDELETEOK {
	return &DeleteVersionByIDUsingDELETEOK{}
}

/*
DeleteVersionByIDUsingDELETEOK describes a response with status code 200, with default header values.

OK
*/
type DeleteVersionByIDUsingDELETEOK struct {
	Payload models.CustomIntegration
}

// IsSuccess returns true when this delete version by Id using d e l e t e o k response has a 2xx status code
func (o *DeleteVersionByIDUsingDELETEOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete version by Id using d e l e t e o k response has a 3xx status code
func (o *DeleteVersionByIDUsingDELETEOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete version by Id using d e l e t e o k response has a 4xx status code
func (o *DeleteVersionByIDUsingDELETEOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete version by Id using d e l e t e o k response has a 5xx status code
func (o *DeleteVersionByIDUsingDELETEOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete version by Id using d e l e t e o k response a status code equal to that given
func (o *DeleteVersionByIDUsingDELETEOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteVersionByIDUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}/versions/{version}][%d] deleteVersionByIdUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteVersionByIDUsingDELETEOK) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}/versions/{version}][%d] deleteVersionByIdUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteVersionByIDUsingDELETEOK) GetPayload() models.CustomIntegration {
	return o.Payload
}

func (o *DeleteVersionByIDUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalCustomIntegration(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeleteVersionByIDUsingDELETEUnauthorized creates a DeleteVersionByIDUsingDELETEUnauthorized with default headers values
func NewDeleteVersionByIDUsingDELETEUnauthorized() *DeleteVersionByIDUsingDELETEUnauthorized {
	return &DeleteVersionByIDUsingDELETEUnauthorized{}
}

/*
DeleteVersionByIDUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type DeleteVersionByIDUsingDELETEUnauthorized struct {
}

// IsSuccess returns true when this delete version by Id using d e l e t e unauthorized response has a 2xx status code
func (o *DeleteVersionByIDUsingDELETEUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete version by Id using d e l e t e unauthorized response has a 3xx status code
func (o *DeleteVersionByIDUsingDELETEUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete version by Id using d e l e t e unauthorized response has a 4xx status code
func (o *DeleteVersionByIDUsingDELETEUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete version by Id using d e l e t e unauthorized response has a 5xx status code
func (o *DeleteVersionByIDUsingDELETEUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete version by Id using d e l e t e unauthorized response a status code equal to that given
func (o *DeleteVersionByIDUsingDELETEUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteVersionByIDUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}/versions/{version}][%d] deleteVersionByIdUsingDELETEUnauthorized ", 401)
}

func (o *DeleteVersionByIDUsingDELETEUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}/versions/{version}][%d] deleteVersionByIdUsingDELETEUnauthorized ", 401)
}

func (o *DeleteVersionByIDUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVersionByIDUsingDELETEForbidden creates a DeleteVersionByIDUsingDELETEForbidden with default headers values
func NewDeleteVersionByIDUsingDELETEForbidden() *DeleteVersionByIDUsingDELETEForbidden {
	return &DeleteVersionByIDUsingDELETEForbidden{}
}

/*
DeleteVersionByIDUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteVersionByIDUsingDELETEForbidden struct {
}

// IsSuccess returns true when this delete version by Id using d e l e t e forbidden response has a 2xx status code
func (o *DeleteVersionByIDUsingDELETEForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete version by Id using d e l e t e forbidden response has a 3xx status code
func (o *DeleteVersionByIDUsingDELETEForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete version by Id using d e l e t e forbidden response has a 4xx status code
func (o *DeleteVersionByIDUsingDELETEForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete version by Id using d e l e t e forbidden response has a 5xx status code
func (o *DeleteVersionByIDUsingDELETEForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete version by Id using d e l e t e forbidden response a status code equal to that given
func (o *DeleteVersionByIDUsingDELETEForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteVersionByIDUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}/versions/{version}][%d] deleteVersionByIdUsingDELETEForbidden ", 403)
}

func (o *DeleteVersionByIDUsingDELETEForbidden) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}/versions/{version}][%d] deleteVersionByIdUsingDELETEForbidden ", 403)
}

func (o *DeleteVersionByIDUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVersionByIDUsingDELETENotFound creates a DeleteVersionByIDUsingDELETENotFound with default headers values
func NewDeleteVersionByIDUsingDELETENotFound() *DeleteVersionByIDUsingDELETENotFound {
	return &DeleteVersionByIDUsingDELETENotFound{}
}

/*
DeleteVersionByIDUsingDELETENotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteVersionByIDUsingDELETENotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete version by Id using d e l e t e not found response has a 2xx status code
func (o *DeleteVersionByIDUsingDELETENotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete version by Id using d e l e t e not found response has a 3xx status code
func (o *DeleteVersionByIDUsingDELETENotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete version by Id using d e l e t e not found response has a 4xx status code
func (o *DeleteVersionByIDUsingDELETENotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete version by Id using d e l e t e not found response has a 5xx status code
func (o *DeleteVersionByIDUsingDELETENotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete version by Id using d e l e t e not found response a status code equal to that given
func (o *DeleteVersionByIDUsingDELETENotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteVersionByIDUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}/versions/{version}][%d] deleteVersionByIdUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteVersionByIDUsingDELETENotFound) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}/versions/{version}][%d] deleteVersionByIdUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteVersionByIDUsingDELETENotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteVersionByIDUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVersionByIDUsingDELETEInternalServerError creates a DeleteVersionByIDUsingDELETEInternalServerError with default headers values
func NewDeleteVersionByIDUsingDELETEInternalServerError() *DeleteVersionByIDUsingDELETEInternalServerError {
	return &DeleteVersionByIDUsingDELETEInternalServerError{}
}

/*
DeleteVersionByIDUsingDELETEInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DeleteVersionByIDUsingDELETEInternalServerError struct {
}

// IsSuccess returns true when this delete version by Id using d e l e t e internal server error response has a 2xx status code
func (o *DeleteVersionByIDUsingDELETEInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete version by Id using d e l e t e internal server error response has a 3xx status code
func (o *DeleteVersionByIDUsingDELETEInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete version by Id using d e l e t e internal server error response has a 4xx status code
func (o *DeleteVersionByIDUsingDELETEInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete version by Id using d e l e t e internal server error response has a 5xx status code
func (o *DeleteVersionByIDUsingDELETEInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete version by Id using d e l e t e internal server error response a status code equal to that given
func (o *DeleteVersionByIDUsingDELETEInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteVersionByIDUsingDELETEInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}/versions/{version}][%d] deleteVersionByIdUsingDELETEInternalServerError ", 500)
}

func (o *DeleteVersionByIDUsingDELETEInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}/versions/{version}][%d] deleteVersionByIdUsingDELETEInternalServerError ", 500)
}

func (o *DeleteVersionByIDUsingDELETEInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}