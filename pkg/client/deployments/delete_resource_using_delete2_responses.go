// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteResourceUsingDELETE2Reader is a Reader for the DeleteResourceUsingDELETE2 structure.
type DeleteResourceUsingDELETE2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteResourceUsingDELETE2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteResourceUsingDELETE2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteResourceUsingDELETE2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteResourceUsingDELETE2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteResourceUsingDELETE2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteResourceUsingDELETE2OK creates a DeleteResourceUsingDELETE2OK with default headers values
func NewDeleteResourceUsingDELETE2OK() *DeleteResourceUsingDELETE2OK {
	return &DeleteResourceUsingDELETE2OK{}
}

/*
DeleteResourceUsingDELETE2OK describes a response with status code 200, with default header values.

OK
*/
type DeleteResourceUsingDELETE2OK struct {
	Payload *models.Request
}

// IsSuccess returns true when this delete resource using d e l e t e2 o k response has a 2xx status code
func (o *DeleteResourceUsingDELETE2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete resource using d e l e t e2 o k response has a 3xx status code
func (o *DeleteResourceUsingDELETE2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resource using d e l e t e2 o k response has a 4xx status code
func (o *DeleteResourceUsingDELETE2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete resource using d e l e t e2 o k response has a 5xx status code
func (o *DeleteResourceUsingDELETE2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete resource using d e l e t e2 o k response a status code equal to that given
func (o *DeleteResourceUsingDELETE2OK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteResourceUsingDELETE2OK) Error() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] deleteResourceUsingDELETE2OK  %+v", 200, o.Payload)
}

func (o *DeleteResourceUsingDELETE2OK) String() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] deleteResourceUsingDELETE2OK  %+v", 200, o.Payload)
}

func (o *DeleteResourceUsingDELETE2OK) GetPayload() *models.Request {
	return o.Payload
}

func (o *DeleteResourceUsingDELETE2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Request)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteResourceUsingDELETE2Unauthorized creates a DeleteResourceUsingDELETE2Unauthorized with default headers values
func NewDeleteResourceUsingDELETE2Unauthorized() *DeleteResourceUsingDELETE2Unauthorized {
	return &DeleteResourceUsingDELETE2Unauthorized{}
}

/*
DeleteResourceUsingDELETE2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteResourceUsingDELETE2Unauthorized struct {
}

// IsSuccess returns true when this delete resource using d e l e t e2 unauthorized response has a 2xx status code
func (o *DeleteResourceUsingDELETE2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete resource using d e l e t e2 unauthorized response has a 3xx status code
func (o *DeleteResourceUsingDELETE2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resource using d e l e t e2 unauthorized response has a 4xx status code
func (o *DeleteResourceUsingDELETE2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete resource using d e l e t e2 unauthorized response has a 5xx status code
func (o *DeleteResourceUsingDELETE2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete resource using d e l e t e2 unauthorized response a status code equal to that given
func (o *DeleteResourceUsingDELETE2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteResourceUsingDELETE2Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] deleteResourceUsingDELETE2Unauthorized ", 401)
}

func (o *DeleteResourceUsingDELETE2Unauthorized) String() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] deleteResourceUsingDELETE2Unauthorized ", 401)
}

func (o *DeleteResourceUsingDELETE2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResourceUsingDELETE2Forbidden creates a DeleteResourceUsingDELETE2Forbidden with default headers values
func NewDeleteResourceUsingDELETE2Forbidden() *DeleteResourceUsingDELETE2Forbidden {
	return &DeleteResourceUsingDELETE2Forbidden{}
}

/*
DeleteResourceUsingDELETE2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteResourceUsingDELETE2Forbidden struct {
}

// IsSuccess returns true when this delete resource using d e l e t e2 forbidden response has a 2xx status code
func (o *DeleteResourceUsingDELETE2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete resource using d e l e t e2 forbidden response has a 3xx status code
func (o *DeleteResourceUsingDELETE2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resource using d e l e t e2 forbidden response has a 4xx status code
func (o *DeleteResourceUsingDELETE2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete resource using d e l e t e2 forbidden response has a 5xx status code
func (o *DeleteResourceUsingDELETE2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete resource using d e l e t e2 forbidden response a status code equal to that given
func (o *DeleteResourceUsingDELETE2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteResourceUsingDELETE2Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] deleteResourceUsingDELETE2Forbidden ", 403)
}

func (o *DeleteResourceUsingDELETE2Forbidden) String() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] deleteResourceUsingDELETE2Forbidden ", 403)
}

func (o *DeleteResourceUsingDELETE2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResourceUsingDELETE2NotFound creates a DeleteResourceUsingDELETE2NotFound with default headers values
func NewDeleteResourceUsingDELETE2NotFound() *DeleteResourceUsingDELETE2NotFound {
	return &DeleteResourceUsingDELETE2NotFound{}
}

/*
DeleteResourceUsingDELETE2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteResourceUsingDELETE2NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete resource using d e l e t e2 not found response has a 2xx status code
func (o *DeleteResourceUsingDELETE2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete resource using d e l e t e2 not found response has a 3xx status code
func (o *DeleteResourceUsingDELETE2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete resource using d e l e t e2 not found response has a 4xx status code
func (o *DeleteResourceUsingDELETE2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete resource using d e l e t e2 not found response has a 5xx status code
func (o *DeleteResourceUsingDELETE2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete resource using d e l e t e2 not found response a status code equal to that given
func (o *DeleteResourceUsingDELETE2NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteResourceUsingDELETE2NotFound) Error() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] deleteResourceUsingDELETE2NotFound  %+v", 404, o.Payload)
}

func (o *DeleteResourceUsingDELETE2NotFound) String() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] deleteResourceUsingDELETE2NotFound  %+v", 404, o.Payload)
}

func (o *DeleteResourceUsingDELETE2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteResourceUsingDELETE2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
