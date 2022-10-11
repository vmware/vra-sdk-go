// Code generated by go-swagger; DO NOT EDIT.

package custom_naming

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteCustomnameReader is a Reader for the DeleteCustomname structure.
type DeleteCustomnameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomnameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCustomnameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteCustomnameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCustomnameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCustomnameNoContent creates a DeleteCustomnameNoContent with default headers values
func NewDeleteCustomnameNoContent() *DeleteCustomnameNoContent {
	return &DeleteCustomnameNoContent{}
}

/*
DeleteCustomnameNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteCustomnameNoContent struct {
}

// IsSuccess returns true when this delete customname no content response has a 2xx status code
func (o *DeleteCustomnameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete customname no content response has a 3xx status code
func (o *DeleteCustomnameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete customname no content response has a 4xx status code
func (o *DeleteCustomnameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete customname no content response has a 5xx status code
func (o *DeleteCustomnameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete customname no content response a status code equal to that given
func (o *DeleteCustomnameNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteCustomnameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/naming/{id}][%d] deleteCustomnameNoContent ", 204)
}

func (o *DeleteCustomnameNoContent) String() string {
	return fmt.Sprintf("[DELETE /iaas/api/naming/{id}][%d] deleteCustomnameNoContent ", 204)
}

func (o *DeleteCustomnameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCustomnameForbidden creates a DeleteCustomnameForbidden with default headers values
func NewDeleteCustomnameForbidden() *DeleteCustomnameForbidden {
	return &DeleteCustomnameForbidden{}
}

/*
DeleteCustomnameForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteCustomnameForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this delete customname forbidden response has a 2xx status code
func (o *DeleteCustomnameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete customname forbidden response has a 3xx status code
func (o *DeleteCustomnameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete customname forbidden response has a 4xx status code
func (o *DeleteCustomnameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete customname forbidden response has a 5xx status code
func (o *DeleteCustomnameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete customname forbidden response a status code equal to that given
func (o *DeleteCustomnameForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteCustomnameForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/naming/{id}][%d] deleteCustomnameForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCustomnameForbidden) String() string {
	return fmt.Sprintf("[DELETE /iaas/api/naming/{id}][%d] deleteCustomnameForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCustomnameForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *DeleteCustomnameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomnameNotFound creates a DeleteCustomnameNotFound with default headers values
func NewDeleteCustomnameNotFound() *DeleteCustomnameNotFound {
	return &DeleteCustomnameNotFound{}
}

/*
DeleteCustomnameNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteCustomnameNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete customname not found response has a 2xx status code
func (o *DeleteCustomnameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete customname not found response has a 3xx status code
func (o *DeleteCustomnameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete customname not found response has a 4xx status code
func (o *DeleteCustomnameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete customname not found response has a 5xx status code
func (o *DeleteCustomnameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete customname not found response a status code equal to that given
func (o *DeleteCustomnameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteCustomnameNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/naming/{id}][%d] deleteCustomnameNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCustomnameNotFound) String() string {
	return fmt.Sprintf("[DELETE /iaas/api/naming/{id}][%d] deleteCustomnameNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCustomnameNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCustomnameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}