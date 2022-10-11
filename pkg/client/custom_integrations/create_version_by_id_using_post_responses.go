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

// CreateVersionByIDUsingPOSTReader is a Reader for the CreateVersionByIDUsingPOST structure.
type CreateVersionByIDUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVersionByIDUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVersionByIDUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateVersionByIDUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateVersionByIDUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateVersionByIDUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateVersionByIDUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVersionByIDUsingPOSTOK creates a CreateVersionByIDUsingPOSTOK with default headers values
func NewCreateVersionByIDUsingPOSTOK() *CreateVersionByIDUsingPOSTOK {
	return &CreateVersionByIDUsingPOSTOK{}
}

/*
CreateVersionByIDUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type CreateVersionByIDUsingPOSTOK struct {
	Payload models.CustomIntegration
}

// IsSuccess returns true when this create version by Id using p o s t o k response has a 2xx status code
func (o *CreateVersionByIDUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create version by Id using p o s t o k response has a 3xx status code
func (o *CreateVersionByIDUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create version by Id using p o s t o k response has a 4xx status code
func (o *CreateVersionByIDUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create version by Id using p o s t o k response has a 5xx status code
func (o *CreateVersionByIDUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create version by Id using p o s t o k response a status code equal to that given
func (o *CreateVersionByIDUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateVersionByIDUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions][%d] createVersionByIdUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateVersionByIDUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions][%d] createVersionByIdUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateVersionByIDUsingPOSTOK) GetPayload() models.CustomIntegration {
	return o.Payload
}

func (o *CreateVersionByIDUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalCustomIntegration(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewCreateVersionByIDUsingPOSTUnauthorized creates a CreateVersionByIDUsingPOSTUnauthorized with default headers values
func NewCreateVersionByIDUsingPOSTUnauthorized() *CreateVersionByIDUsingPOSTUnauthorized {
	return &CreateVersionByIDUsingPOSTUnauthorized{}
}

/*
CreateVersionByIDUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type CreateVersionByIDUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this create version by Id using p o s t unauthorized response has a 2xx status code
func (o *CreateVersionByIDUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create version by Id using p o s t unauthorized response has a 3xx status code
func (o *CreateVersionByIDUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create version by Id using p o s t unauthorized response has a 4xx status code
func (o *CreateVersionByIDUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create version by Id using p o s t unauthorized response has a 5xx status code
func (o *CreateVersionByIDUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create version by Id using p o s t unauthorized response a status code equal to that given
func (o *CreateVersionByIDUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateVersionByIDUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions][%d] createVersionByIdUsingPOSTUnauthorized ", 401)
}

func (o *CreateVersionByIDUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions][%d] createVersionByIdUsingPOSTUnauthorized ", 401)
}

func (o *CreateVersionByIDUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateVersionByIDUsingPOSTForbidden creates a CreateVersionByIDUsingPOSTForbidden with default headers values
func NewCreateVersionByIDUsingPOSTForbidden() *CreateVersionByIDUsingPOSTForbidden {
	return &CreateVersionByIDUsingPOSTForbidden{}
}

/*
CreateVersionByIDUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateVersionByIDUsingPOSTForbidden struct {
}

// IsSuccess returns true when this create version by Id using p o s t forbidden response has a 2xx status code
func (o *CreateVersionByIDUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create version by Id using p o s t forbidden response has a 3xx status code
func (o *CreateVersionByIDUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create version by Id using p o s t forbidden response has a 4xx status code
func (o *CreateVersionByIDUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create version by Id using p o s t forbidden response has a 5xx status code
func (o *CreateVersionByIDUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create version by Id using p o s t forbidden response a status code equal to that given
func (o *CreateVersionByIDUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateVersionByIDUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions][%d] createVersionByIdUsingPOSTForbidden ", 403)
}

func (o *CreateVersionByIDUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions][%d] createVersionByIdUsingPOSTForbidden ", 403)
}

func (o *CreateVersionByIDUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateVersionByIDUsingPOSTNotFound creates a CreateVersionByIDUsingPOSTNotFound with default headers values
func NewCreateVersionByIDUsingPOSTNotFound() *CreateVersionByIDUsingPOSTNotFound {
	return &CreateVersionByIDUsingPOSTNotFound{}
}

/*
CreateVersionByIDUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateVersionByIDUsingPOSTNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this create version by Id using p o s t not found response has a 2xx status code
func (o *CreateVersionByIDUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create version by Id using p o s t not found response has a 3xx status code
func (o *CreateVersionByIDUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create version by Id using p o s t not found response has a 4xx status code
func (o *CreateVersionByIDUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create version by Id using p o s t not found response has a 5xx status code
func (o *CreateVersionByIDUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create version by Id using p o s t not found response a status code equal to that given
func (o *CreateVersionByIDUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateVersionByIDUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions][%d] createVersionByIdUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *CreateVersionByIDUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions][%d] createVersionByIdUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *CreateVersionByIDUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateVersionByIDUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVersionByIDUsingPOSTInternalServerError creates a CreateVersionByIDUsingPOSTInternalServerError with default headers values
func NewCreateVersionByIDUsingPOSTInternalServerError() *CreateVersionByIDUsingPOSTInternalServerError {
	return &CreateVersionByIDUsingPOSTInternalServerError{}
}

/*
CreateVersionByIDUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CreateVersionByIDUsingPOSTInternalServerError struct {
}

// IsSuccess returns true when this create version by Id using p o s t internal server error response has a 2xx status code
func (o *CreateVersionByIDUsingPOSTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create version by Id using p o s t internal server error response has a 3xx status code
func (o *CreateVersionByIDUsingPOSTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create version by Id using p o s t internal server error response has a 4xx status code
func (o *CreateVersionByIDUsingPOSTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create version by Id using p o s t internal server error response has a 5xx status code
func (o *CreateVersionByIDUsingPOSTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create version by Id using p o s t internal server error response a status code equal to that given
func (o *CreateVersionByIDUsingPOSTInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateVersionByIDUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions][%d] createVersionByIdUsingPOSTInternalServerError ", 500)
}

func (o *CreateVersionByIDUsingPOSTInternalServerError) String() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions][%d] createVersionByIdUsingPOSTInternalServerError ", 500)
}

func (o *CreateVersionByIDUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}