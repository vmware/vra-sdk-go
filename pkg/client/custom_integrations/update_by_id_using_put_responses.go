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

// UpdateByIDUsingPUTReader is a Reader for the UpdateByIDUsingPUT structure.
type UpdateByIDUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateByIDUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateByIDUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateByIDUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateByIDUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateByIDUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateByIDUsingPUTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateByIDUsingPUTOK creates a UpdateByIDUsingPUTOK with default headers values
func NewUpdateByIDUsingPUTOK() *UpdateByIDUsingPUTOK {
	return &UpdateByIDUsingPUTOK{}
}

/*
UpdateByIDUsingPUTOK describes a response with status code 200, with default header values.

OK
*/
type UpdateByIDUsingPUTOK struct {
	Payload models.CustomIntegration
}

// IsSuccess returns true when this update by Id using p u t o k response has a 2xx status code
func (o *UpdateByIDUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update by Id using p u t o k response has a 3xx status code
func (o *UpdateByIDUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update by Id using p u t o k response has a 4xx status code
func (o *UpdateByIDUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update by Id using p u t o k response has a 5xx status code
func (o *UpdateByIDUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update by Id using p u t o k response a status code equal to that given
func (o *UpdateByIDUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateByIDUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/custom-integrations/{id}][%d] updateByIdUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateByIDUsingPUTOK) String() string {
	return fmt.Sprintf("[PUT /codestream/api/custom-integrations/{id}][%d] updateByIdUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateByIDUsingPUTOK) GetPayload() models.CustomIntegration {
	return o.Payload
}

func (o *UpdateByIDUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalCustomIntegration(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateByIDUsingPUTUnauthorized creates a UpdateByIDUsingPUTUnauthorized with default headers values
func NewUpdateByIDUsingPUTUnauthorized() *UpdateByIDUsingPUTUnauthorized {
	return &UpdateByIDUsingPUTUnauthorized{}
}

/*
UpdateByIDUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type UpdateByIDUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this update by Id using p u t unauthorized response has a 2xx status code
func (o *UpdateByIDUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update by Id using p u t unauthorized response has a 3xx status code
func (o *UpdateByIDUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update by Id using p u t unauthorized response has a 4xx status code
func (o *UpdateByIDUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update by Id using p u t unauthorized response has a 5xx status code
func (o *UpdateByIDUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update by Id using p u t unauthorized response a status code equal to that given
func (o *UpdateByIDUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateByIDUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/custom-integrations/{id}][%d] updateByIdUsingPUTUnauthorized ", 401)
}

func (o *UpdateByIDUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /codestream/api/custom-integrations/{id}][%d] updateByIdUsingPUTUnauthorized ", 401)
}

func (o *UpdateByIDUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateByIDUsingPUTForbidden creates a UpdateByIDUsingPUTForbidden with default headers values
func NewUpdateByIDUsingPUTForbidden() *UpdateByIDUsingPUTForbidden {
	return &UpdateByIDUsingPUTForbidden{}
}

/*
UpdateByIDUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateByIDUsingPUTForbidden struct {
}

// IsSuccess returns true when this update by Id using p u t forbidden response has a 2xx status code
func (o *UpdateByIDUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update by Id using p u t forbidden response has a 3xx status code
func (o *UpdateByIDUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update by Id using p u t forbidden response has a 4xx status code
func (o *UpdateByIDUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update by Id using p u t forbidden response has a 5xx status code
func (o *UpdateByIDUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update by Id using p u t forbidden response a status code equal to that given
func (o *UpdateByIDUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateByIDUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/custom-integrations/{id}][%d] updateByIdUsingPUTForbidden ", 403)
}

func (o *UpdateByIDUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /codestream/api/custom-integrations/{id}][%d] updateByIdUsingPUTForbidden ", 403)
}

func (o *UpdateByIDUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateByIDUsingPUTNotFound creates a UpdateByIDUsingPUTNotFound with default headers values
func NewUpdateByIDUsingPUTNotFound() *UpdateByIDUsingPUTNotFound {
	return &UpdateByIDUsingPUTNotFound{}
}

/*
UpdateByIDUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateByIDUsingPUTNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update by Id using p u t not found response has a 2xx status code
func (o *UpdateByIDUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update by Id using p u t not found response has a 3xx status code
func (o *UpdateByIDUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update by Id using p u t not found response has a 4xx status code
func (o *UpdateByIDUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update by Id using p u t not found response has a 5xx status code
func (o *UpdateByIDUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update by Id using p u t not found response a status code equal to that given
func (o *UpdateByIDUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateByIDUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/custom-integrations/{id}][%d] updateByIdUsingPUTNotFound  %+v", 404, o.Payload)
}

func (o *UpdateByIDUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /codestream/api/custom-integrations/{id}][%d] updateByIdUsingPUTNotFound  %+v", 404, o.Payload)
}

func (o *UpdateByIDUsingPUTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateByIDUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateByIDUsingPUTInternalServerError creates a UpdateByIDUsingPUTInternalServerError with default headers values
func NewUpdateByIDUsingPUTInternalServerError() *UpdateByIDUsingPUTInternalServerError {
	return &UpdateByIDUsingPUTInternalServerError{}
}

/*
UpdateByIDUsingPUTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdateByIDUsingPUTInternalServerError struct {
}

// IsSuccess returns true when this update by Id using p u t internal server error response has a 2xx status code
func (o *UpdateByIDUsingPUTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update by Id using p u t internal server error response has a 3xx status code
func (o *UpdateByIDUsingPUTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update by Id using p u t internal server error response has a 4xx status code
func (o *UpdateByIDUsingPUTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update by Id using p u t internal server error response has a 5xx status code
func (o *UpdateByIDUsingPUTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update by Id using p u t internal server error response a status code equal to that given
func (o *UpdateByIDUsingPUTInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateByIDUsingPUTInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/custom-integrations/{id}][%d] updateByIdUsingPUTInternalServerError ", 500)
}

func (o *UpdateByIDUsingPUTInternalServerError) String() string {
	return fmt.Sprintf("[PUT /codestream/api/custom-integrations/{id}][%d] updateByIdUsingPUTInternalServerError ", 500)
}

func (o *UpdateByIDUsingPUTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
