// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateVariableByNameUsingPUTReader is a Reader for the UpdateVariableByNameUsingPUT structure.
type UpdateVariableByNameUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVariableByNameUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVariableByNameUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateVariableByNameUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateVariableByNameUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVariableByNameUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateVariableByNameUsingPUTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVariableByNameUsingPUTOK creates a UpdateVariableByNameUsingPUTOK with default headers values
func NewUpdateVariableByNameUsingPUTOK() *UpdateVariableByNameUsingPUTOK {
	return &UpdateVariableByNameUsingPUTOK{}
}

/*
UpdateVariableByNameUsingPUTOK describes a response with status code 200, with default header values.

'Success' with updated Variable
*/
type UpdateVariableByNameUsingPUTOK struct {
	Payload *models.Variable
}

// IsSuccess returns true when this update variable by name using p u t o k response has a 2xx status code
func (o *UpdateVariableByNameUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update variable by name using p u t o k response has a 3xx status code
func (o *UpdateVariableByNameUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update variable by name using p u t o k response has a 4xx status code
func (o *UpdateVariableByNameUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update variable by name using p u t o k response has a 5xx status code
func (o *UpdateVariableByNameUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update variable by name using p u t o k response a status code equal to that given
func (o *UpdateVariableByNameUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateVariableByNameUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateVariableByNameUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateVariableByNameUsingPUTOK) String() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateVariableByNameUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateVariableByNameUsingPUTOK) GetPayload() *models.Variable {
	return o.Payload
}

func (o *UpdateVariableByNameUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Variable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVariableByNameUsingPUTUnauthorized creates a UpdateVariableByNameUsingPUTUnauthorized with default headers values
func NewUpdateVariableByNameUsingPUTUnauthorized() *UpdateVariableByNameUsingPUTUnauthorized {
	return &UpdateVariableByNameUsingPUTUnauthorized{}
}

/*
UpdateVariableByNameUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type UpdateVariableByNameUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this update variable by name using p u t unauthorized response has a 2xx status code
func (o *UpdateVariableByNameUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update variable by name using p u t unauthorized response has a 3xx status code
func (o *UpdateVariableByNameUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update variable by name using p u t unauthorized response has a 4xx status code
func (o *UpdateVariableByNameUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update variable by name using p u t unauthorized response has a 5xx status code
func (o *UpdateVariableByNameUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update variable by name using p u t unauthorized response a status code equal to that given
func (o *UpdateVariableByNameUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateVariableByNameUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateVariableByNameUsingPUTUnauthorized ", 401)
}

func (o *UpdateVariableByNameUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateVariableByNameUsingPUTUnauthorized ", 401)
}

func (o *UpdateVariableByNameUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVariableByNameUsingPUTForbidden creates a UpdateVariableByNameUsingPUTForbidden with default headers values
func NewUpdateVariableByNameUsingPUTForbidden() *UpdateVariableByNameUsingPUTForbidden {
	return &UpdateVariableByNameUsingPUTForbidden{}
}

/*
UpdateVariableByNameUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateVariableByNameUsingPUTForbidden struct {
}

// IsSuccess returns true when this update variable by name using p u t forbidden response has a 2xx status code
func (o *UpdateVariableByNameUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update variable by name using p u t forbidden response has a 3xx status code
func (o *UpdateVariableByNameUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update variable by name using p u t forbidden response has a 4xx status code
func (o *UpdateVariableByNameUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update variable by name using p u t forbidden response has a 5xx status code
func (o *UpdateVariableByNameUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update variable by name using p u t forbidden response a status code equal to that given
func (o *UpdateVariableByNameUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateVariableByNameUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateVariableByNameUsingPUTForbidden ", 403)
}

func (o *UpdateVariableByNameUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateVariableByNameUsingPUTForbidden ", 403)
}

func (o *UpdateVariableByNameUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVariableByNameUsingPUTNotFound creates a UpdateVariableByNameUsingPUTNotFound with default headers values
func NewUpdateVariableByNameUsingPUTNotFound() *UpdateVariableByNameUsingPUTNotFound {
	return &UpdateVariableByNameUsingPUTNotFound{}
}

/*
UpdateVariableByNameUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateVariableByNameUsingPUTNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update variable by name using p u t not found response has a 2xx status code
func (o *UpdateVariableByNameUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update variable by name using p u t not found response has a 3xx status code
func (o *UpdateVariableByNameUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update variable by name using p u t not found response has a 4xx status code
func (o *UpdateVariableByNameUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update variable by name using p u t not found response has a 5xx status code
func (o *UpdateVariableByNameUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update variable by name using p u t not found response a status code equal to that given
func (o *UpdateVariableByNameUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateVariableByNameUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateVariableByNameUsingPUTNotFound  %+v", 404, o.Payload)
}

func (o *UpdateVariableByNameUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateVariableByNameUsingPUTNotFound  %+v", 404, o.Payload)
}

func (o *UpdateVariableByNameUsingPUTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateVariableByNameUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVariableByNameUsingPUTInternalServerError creates a UpdateVariableByNameUsingPUTInternalServerError with default headers values
func NewUpdateVariableByNameUsingPUTInternalServerError() *UpdateVariableByNameUsingPUTInternalServerError {
	return &UpdateVariableByNameUsingPUTInternalServerError{}
}

/*
UpdateVariableByNameUsingPUTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdateVariableByNameUsingPUTInternalServerError struct {
}

// IsSuccess returns true when this update variable by name using p u t internal server error response has a 2xx status code
func (o *UpdateVariableByNameUsingPUTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update variable by name using p u t internal server error response has a 3xx status code
func (o *UpdateVariableByNameUsingPUTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update variable by name using p u t internal server error response has a 4xx status code
func (o *UpdateVariableByNameUsingPUTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update variable by name using p u t internal server error response has a 5xx status code
func (o *UpdateVariableByNameUsingPUTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update variable by name using p u t internal server error response a status code equal to that given
func (o *UpdateVariableByNameUsingPUTInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateVariableByNameUsingPUTInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateVariableByNameUsingPUTInternalServerError ", 500)
}

func (o *UpdateVariableByNameUsingPUTInternalServerError) String() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateVariableByNameUsingPUTInternalServerError ", 500)
}

func (o *UpdateVariableByNameUsingPUTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
