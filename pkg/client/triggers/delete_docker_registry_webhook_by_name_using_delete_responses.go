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

// DeleteDockerRegistryWebhookByNameUsingDELETEReader is a Reader for the DeleteDockerRegistryWebhookByNameUsingDELETE structure.
type DeleteDockerRegistryWebhookByNameUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDockerRegistryWebhookByNameUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDockerRegistryWebhookByNameUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDockerRegistryWebhookByNameUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDockerRegistryWebhookByNameUsingDELETEOK creates a DeleteDockerRegistryWebhookByNameUsingDELETEOK with default headers values
func NewDeleteDockerRegistryWebhookByNameUsingDELETEOK() *DeleteDockerRegistryWebhookByNameUsingDELETEOK {
	return &DeleteDockerRegistryWebhookByNameUsingDELETEOK{}
}

/*
DeleteDockerRegistryWebhookByNameUsingDELETEOK describes a response with status code 200, with default header values.

'Success' with Docker Registry Webhook Delete
*/
type DeleteDockerRegistryWebhookByNameUsingDELETEOK struct {
	Payload models.DockerRegistryWebHook
}

// IsSuccess returns true when this delete docker registry webhook by name using d e l e t e o k response has a 2xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete docker registry webhook by name using d e l e t e o k response has a 3xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete docker registry webhook by name using d e l e t e o k response has a 4xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete docker registry webhook by name using d e l e t e o k response has a 5xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete docker registry webhook by name using d e l e t e o k response a status code equal to that given
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-webhooks/{project}/{name}][%d] deleteDockerRegistryWebhookByNameUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETEOK) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-webhooks/{project}/{name}][%d] deleteDockerRegistryWebhookByNameUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETEOK) GetPayload() models.DockerRegistryWebHook {
	return o.Payload
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalDockerRegistryWebHook(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized creates a DeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized with default headers values
func NewDeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized() *DeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized {
	return &DeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized{}
}

/*
DeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type DeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized struct {
}

// IsSuccess returns true when this delete docker registry webhook by name using d e l e t e unauthorized response has a 2xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete docker registry webhook by name using d e l e t e unauthorized response has a 3xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete docker registry webhook by name using d e l e t e unauthorized response has a 4xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete docker registry webhook by name using d e l e t e unauthorized response has a 5xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete docker registry webhook by name using d e l e t e unauthorized response a status code equal to that given
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-webhooks/{project}/{name}][%d] deleteDockerRegistryWebhookByNameUsingDELETEUnauthorized ", 401)
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-webhooks/{project}/{name}][%d] deleteDockerRegistryWebhookByNameUsingDELETEUnauthorized ", 401)
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDockerRegistryWebhookByNameUsingDELETEForbidden creates a DeleteDockerRegistryWebhookByNameUsingDELETEForbidden with default headers values
func NewDeleteDockerRegistryWebhookByNameUsingDELETEForbidden() *DeleteDockerRegistryWebhookByNameUsingDELETEForbidden {
	return &DeleteDockerRegistryWebhookByNameUsingDELETEForbidden{}
}

/*
DeleteDockerRegistryWebhookByNameUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteDockerRegistryWebhookByNameUsingDELETEForbidden struct {
}

// IsSuccess returns true when this delete docker registry webhook by name using d e l e t e forbidden response has a 2xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete docker registry webhook by name using d e l e t e forbidden response has a 3xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete docker registry webhook by name using d e l e t e forbidden response has a 4xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete docker registry webhook by name using d e l e t e forbidden response has a 5xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete docker registry webhook by name using d e l e t e forbidden response a status code equal to that given
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-webhooks/{project}/{name}][%d] deleteDockerRegistryWebhookByNameUsingDELETEForbidden ", 403)
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETEForbidden) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-webhooks/{project}/{name}][%d] deleteDockerRegistryWebhookByNameUsingDELETEForbidden ", 403)
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDockerRegistryWebhookByNameUsingDELETENotFound creates a DeleteDockerRegistryWebhookByNameUsingDELETENotFound with default headers values
func NewDeleteDockerRegistryWebhookByNameUsingDELETENotFound() *DeleteDockerRegistryWebhookByNameUsingDELETENotFound {
	return &DeleteDockerRegistryWebhookByNameUsingDELETENotFound{}
}

/*
DeleteDockerRegistryWebhookByNameUsingDELETENotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteDockerRegistryWebhookByNameUsingDELETENotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete docker registry webhook by name using d e l e t e not found response has a 2xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETENotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete docker registry webhook by name using d e l e t e not found response has a 3xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETENotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete docker registry webhook by name using d e l e t e not found response has a 4xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETENotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete docker registry webhook by name using d e l e t e not found response has a 5xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETENotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete docker registry webhook by name using d e l e t e not found response a status code equal to that given
func (o *DeleteDockerRegistryWebhookByNameUsingDELETENotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-webhooks/{project}/{name}][%d] deleteDockerRegistryWebhookByNameUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETENotFound) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-webhooks/{project}/{name}][%d] deleteDockerRegistryWebhookByNameUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETENotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError creates a DeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError with default headers values
func NewDeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError() *DeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError {
	return &DeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError{}
}

/*
DeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError struct {
}

// IsSuccess returns true when this delete docker registry webhook by name using d e l e t e internal server error response has a 2xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete docker registry webhook by name using d e l e t e internal server error response has a 3xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete docker registry webhook by name using d e l e t e internal server error response has a 4xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete docker registry webhook by name using d e l e t e internal server error response has a 5xx status code
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete docker registry webhook by name using d e l e t e internal server error response a status code equal to that given
func (o *DeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-webhooks/{project}/{name}][%d] deleteDockerRegistryWebhookByNameUsingDELETEInternalServerError ", 500)
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-webhooks/{project}/{name}][%d] deleteDockerRegistryWebhookByNameUsingDELETEInternalServerError ", 500)
}

func (o *DeleteDockerRegistryWebhookByNameUsingDELETEInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
