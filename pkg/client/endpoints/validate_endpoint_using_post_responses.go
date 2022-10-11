// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// ValidateEndpointUsingPOSTReader is a Reader for the ValidateEndpointUsingPOST structure.
type ValidateEndpointUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateEndpointUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateEndpointUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewValidateEndpointUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewValidateEndpointUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewValidateEndpointUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewValidateEndpointUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewValidateEndpointUsingPOSTOK creates a ValidateEndpointUsingPOSTOK with default headers values
func NewValidateEndpointUsingPOSTOK() *ValidateEndpointUsingPOSTOK {
	return &ValidateEndpointUsingPOSTOK{}
}

/*
ValidateEndpointUsingPOSTOK describes a response with status code 200, with default header values.

'Success' with endpoint validations
*/
type ValidateEndpointUsingPOSTOK struct {
	Payload *models.TileExecutorResponse
}

// IsSuccess returns true when this validate endpoint using p o s t o k response has a 2xx status code
func (o *ValidateEndpointUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate endpoint using p o s t o k response has a 3xx status code
func (o *ValidateEndpointUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate endpoint using p o s t o k response has a 4xx status code
func (o *ValidateEndpointUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate endpoint using p o s t o k response has a 5xx status code
func (o *ValidateEndpointUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate endpoint using p o s t o k response a status code equal to that given
func (o *ValidateEndpointUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

func (o *ValidateEndpointUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/endpoint-validation][%d] validateEndpointUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ValidateEndpointUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /codestream/api/endpoint-validation][%d] validateEndpointUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ValidateEndpointUsingPOSTOK) GetPayload() *models.TileExecutorResponse {
	return o.Payload
}

func (o *ValidateEndpointUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TileExecutorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateEndpointUsingPOSTUnauthorized creates a ValidateEndpointUsingPOSTUnauthorized with default headers values
func NewValidateEndpointUsingPOSTUnauthorized() *ValidateEndpointUsingPOSTUnauthorized {
	return &ValidateEndpointUsingPOSTUnauthorized{}
}

/*
ValidateEndpointUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type ValidateEndpointUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this validate endpoint using p o s t unauthorized response has a 2xx status code
func (o *ValidateEndpointUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate endpoint using p o s t unauthorized response has a 3xx status code
func (o *ValidateEndpointUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate endpoint using p o s t unauthorized response has a 4xx status code
func (o *ValidateEndpointUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate endpoint using p o s t unauthorized response has a 5xx status code
func (o *ValidateEndpointUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this validate endpoint using p o s t unauthorized response a status code equal to that given
func (o *ValidateEndpointUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ValidateEndpointUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/endpoint-validation][%d] validateEndpointUsingPOSTUnauthorized ", 401)
}

func (o *ValidateEndpointUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /codestream/api/endpoint-validation][%d] validateEndpointUsingPOSTUnauthorized ", 401)
}

func (o *ValidateEndpointUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateEndpointUsingPOSTForbidden creates a ValidateEndpointUsingPOSTForbidden with default headers values
func NewValidateEndpointUsingPOSTForbidden() *ValidateEndpointUsingPOSTForbidden {
	return &ValidateEndpointUsingPOSTForbidden{}
}

/*
ValidateEndpointUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ValidateEndpointUsingPOSTForbidden struct {
}

// IsSuccess returns true when this validate endpoint using p o s t forbidden response has a 2xx status code
func (o *ValidateEndpointUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate endpoint using p o s t forbidden response has a 3xx status code
func (o *ValidateEndpointUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate endpoint using p o s t forbidden response has a 4xx status code
func (o *ValidateEndpointUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate endpoint using p o s t forbidden response has a 5xx status code
func (o *ValidateEndpointUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this validate endpoint using p o s t forbidden response a status code equal to that given
func (o *ValidateEndpointUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ValidateEndpointUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/endpoint-validation][%d] validateEndpointUsingPOSTForbidden ", 403)
}

func (o *ValidateEndpointUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /codestream/api/endpoint-validation][%d] validateEndpointUsingPOSTForbidden ", 403)
}

func (o *ValidateEndpointUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateEndpointUsingPOSTNotFound creates a ValidateEndpointUsingPOSTNotFound with default headers values
func NewValidateEndpointUsingPOSTNotFound() *ValidateEndpointUsingPOSTNotFound {
	return &ValidateEndpointUsingPOSTNotFound{}
}

/*
ValidateEndpointUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ValidateEndpointUsingPOSTNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this validate endpoint using p o s t not found response has a 2xx status code
func (o *ValidateEndpointUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate endpoint using p o s t not found response has a 3xx status code
func (o *ValidateEndpointUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate endpoint using p o s t not found response has a 4xx status code
func (o *ValidateEndpointUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate endpoint using p o s t not found response has a 5xx status code
func (o *ValidateEndpointUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this validate endpoint using p o s t not found response a status code equal to that given
func (o *ValidateEndpointUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ValidateEndpointUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/endpoint-validation][%d] validateEndpointUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *ValidateEndpointUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /codestream/api/endpoint-validation][%d] validateEndpointUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *ValidateEndpointUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateEndpointUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateEndpointUsingPOSTInternalServerError creates a ValidateEndpointUsingPOSTInternalServerError with default headers values
func NewValidateEndpointUsingPOSTInternalServerError() *ValidateEndpointUsingPOSTInternalServerError {
	return &ValidateEndpointUsingPOSTInternalServerError{}
}

/*
ValidateEndpointUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ValidateEndpointUsingPOSTInternalServerError struct {
}

// IsSuccess returns true when this validate endpoint using p o s t internal server error response has a 2xx status code
func (o *ValidateEndpointUsingPOSTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate endpoint using p o s t internal server error response has a 3xx status code
func (o *ValidateEndpointUsingPOSTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate endpoint using p o s t internal server error response has a 4xx status code
func (o *ValidateEndpointUsingPOSTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate endpoint using p o s t internal server error response has a 5xx status code
func (o *ValidateEndpointUsingPOSTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this validate endpoint using p o s t internal server error response a status code equal to that given
func (o *ValidateEndpointUsingPOSTInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ValidateEndpointUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/endpoint-validation][%d] validateEndpointUsingPOSTInternalServerError ", 500)
}

func (o *ValidateEndpointUsingPOSTInternalServerError) String() string {
	return fmt.Sprintf("[POST /codestream/api/endpoint-validation][%d] validateEndpointUsingPOSTInternalServerError ", 500)
}

func (o *ValidateEndpointUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
