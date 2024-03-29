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

// CreateGerritListenerUsingPOSTReader is a Reader for the CreateGerritListenerUsingPOST structure.
type CreateGerritListenerUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGerritListenerUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateGerritListenerUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateGerritListenerUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateGerritListenerUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateGerritListenerUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateGerritListenerUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateGerritListenerUsingPOSTOK creates a CreateGerritListenerUsingPOSTOK with default headers values
func NewCreateGerritListenerUsingPOSTOK() *CreateGerritListenerUsingPOSTOK {
	return &CreateGerritListenerUsingPOSTOK{}
}

/*
CreateGerritListenerUsingPOSTOK describes a response with status code 200, with default header values.

'Success' with Gerrit Listener Creation
*/
type CreateGerritListenerUsingPOSTOK struct {
	Payload models.GerritListener
}

// IsSuccess returns true when this create gerrit listener using p o s t o k response has a 2xx status code
func (o *CreateGerritListenerUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create gerrit listener using p o s t o k response has a 3xx status code
func (o *CreateGerritListenerUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create gerrit listener using p o s t o k response has a 4xx status code
func (o *CreateGerritListenerUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create gerrit listener using p o s t o k response has a 5xx status code
func (o *CreateGerritListenerUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create gerrit listener using p o s t o k response a status code equal to that given
func (o *CreateGerritListenerUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateGerritListenerUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateGerritListenerUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateGerritListenerUsingPOSTOK) GetPayload() models.GerritListener {
	return o.Payload
}

func (o *CreateGerritListenerUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListener(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewCreateGerritListenerUsingPOSTUnauthorized creates a CreateGerritListenerUsingPOSTUnauthorized with default headers values
func NewCreateGerritListenerUsingPOSTUnauthorized() *CreateGerritListenerUsingPOSTUnauthorized {
	return &CreateGerritListenerUsingPOSTUnauthorized{}
}

/*
CreateGerritListenerUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type CreateGerritListenerUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this create gerrit listener using p o s t unauthorized response has a 2xx status code
func (o *CreateGerritListenerUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create gerrit listener using p o s t unauthorized response has a 3xx status code
func (o *CreateGerritListenerUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create gerrit listener using p o s t unauthorized response has a 4xx status code
func (o *CreateGerritListenerUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create gerrit listener using p o s t unauthorized response has a 5xx status code
func (o *CreateGerritListenerUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create gerrit listener using p o s t unauthorized response a status code equal to that given
func (o *CreateGerritListenerUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateGerritListenerUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTUnauthorized ", 401)
}

func (o *CreateGerritListenerUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTUnauthorized ", 401)
}

func (o *CreateGerritListenerUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGerritListenerUsingPOSTForbidden creates a CreateGerritListenerUsingPOSTForbidden with default headers values
func NewCreateGerritListenerUsingPOSTForbidden() *CreateGerritListenerUsingPOSTForbidden {
	return &CreateGerritListenerUsingPOSTForbidden{}
}

/*
CreateGerritListenerUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateGerritListenerUsingPOSTForbidden struct {
}

// IsSuccess returns true when this create gerrit listener using p o s t forbidden response has a 2xx status code
func (o *CreateGerritListenerUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create gerrit listener using p o s t forbidden response has a 3xx status code
func (o *CreateGerritListenerUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create gerrit listener using p o s t forbidden response has a 4xx status code
func (o *CreateGerritListenerUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create gerrit listener using p o s t forbidden response has a 5xx status code
func (o *CreateGerritListenerUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create gerrit listener using p o s t forbidden response a status code equal to that given
func (o *CreateGerritListenerUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateGerritListenerUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTForbidden ", 403)
}

func (o *CreateGerritListenerUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTForbidden ", 403)
}

func (o *CreateGerritListenerUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGerritListenerUsingPOSTNotFound creates a CreateGerritListenerUsingPOSTNotFound with default headers values
func NewCreateGerritListenerUsingPOSTNotFound() *CreateGerritListenerUsingPOSTNotFound {
	return &CreateGerritListenerUsingPOSTNotFound{}
}

/*
CreateGerritListenerUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateGerritListenerUsingPOSTNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this create gerrit listener using p o s t not found response has a 2xx status code
func (o *CreateGerritListenerUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create gerrit listener using p o s t not found response has a 3xx status code
func (o *CreateGerritListenerUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create gerrit listener using p o s t not found response has a 4xx status code
func (o *CreateGerritListenerUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create gerrit listener using p o s t not found response has a 5xx status code
func (o *CreateGerritListenerUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create gerrit listener using p o s t not found response a status code equal to that given
func (o *CreateGerritListenerUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateGerritListenerUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *CreateGerritListenerUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *CreateGerritListenerUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGerritListenerUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGerritListenerUsingPOSTInternalServerError creates a CreateGerritListenerUsingPOSTInternalServerError with default headers values
func NewCreateGerritListenerUsingPOSTInternalServerError() *CreateGerritListenerUsingPOSTInternalServerError {
	return &CreateGerritListenerUsingPOSTInternalServerError{}
}

/*
CreateGerritListenerUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CreateGerritListenerUsingPOSTInternalServerError struct {
}

// IsSuccess returns true when this create gerrit listener using p o s t internal server error response has a 2xx status code
func (o *CreateGerritListenerUsingPOSTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create gerrit listener using p o s t internal server error response has a 3xx status code
func (o *CreateGerritListenerUsingPOSTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create gerrit listener using p o s t internal server error response has a 4xx status code
func (o *CreateGerritListenerUsingPOSTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create gerrit listener using p o s t internal server error response has a 5xx status code
func (o *CreateGerritListenerUsingPOSTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create gerrit listener using p o s t internal server error response a status code equal to that given
func (o *CreateGerritListenerUsingPOSTInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateGerritListenerUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTInternalServerError ", 500)
}

func (o *CreateGerritListenerUsingPOSTInternalServerError) String() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTInternalServerError ", 500)
}

func (o *CreateGerritListenerUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
