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

// CreateGitWebhookUsingPOSTReader is a Reader for the CreateGitWebhookUsingPOST structure.
type CreateGitWebhookUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGitWebhookUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateGitWebhookUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateGitWebhookUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateGitWebhookUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateGitWebhookUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateGitWebhookUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateGitWebhookUsingPOSTOK creates a CreateGitWebhookUsingPOSTOK with default headers values
func NewCreateGitWebhookUsingPOSTOK() *CreateGitWebhookUsingPOSTOK {
	return &CreateGitWebhookUsingPOSTOK{}
}

/*
CreateGitWebhookUsingPOSTOK describes a response with status code 200, with default header values.

'Success' with Git Webhook Creation
*/
type CreateGitWebhookUsingPOSTOK struct {
	Payload models.GitWebhook
}

// IsSuccess returns true when this create git webhook using p o s t o k response has a 2xx status code
func (o *CreateGitWebhookUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create git webhook using p o s t o k response has a 3xx status code
func (o *CreateGitWebhookUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create git webhook using p o s t o k response has a 4xx status code
func (o *CreateGitWebhookUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create git webhook using p o s t o k response has a 5xx status code
func (o *CreateGitWebhookUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create git webhook using p o s t o k response a status code equal to that given
func (o *CreateGitWebhookUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateGitWebhookUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createGitWebhookUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateGitWebhookUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createGitWebhookUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateGitWebhookUsingPOSTOK) GetPayload() models.GitWebhook {
	return o.Payload
}

func (o *CreateGitWebhookUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGitWebhook(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewCreateGitWebhookUsingPOSTUnauthorized creates a CreateGitWebhookUsingPOSTUnauthorized with default headers values
func NewCreateGitWebhookUsingPOSTUnauthorized() *CreateGitWebhookUsingPOSTUnauthorized {
	return &CreateGitWebhookUsingPOSTUnauthorized{}
}

/*
CreateGitWebhookUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type CreateGitWebhookUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this create git webhook using p o s t unauthorized response has a 2xx status code
func (o *CreateGitWebhookUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create git webhook using p o s t unauthorized response has a 3xx status code
func (o *CreateGitWebhookUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create git webhook using p o s t unauthorized response has a 4xx status code
func (o *CreateGitWebhookUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create git webhook using p o s t unauthorized response has a 5xx status code
func (o *CreateGitWebhookUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create git webhook using p o s t unauthorized response a status code equal to that given
func (o *CreateGitWebhookUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateGitWebhookUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createGitWebhookUsingPOSTUnauthorized ", 401)
}

func (o *CreateGitWebhookUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createGitWebhookUsingPOSTUnauthorized ", 401)
}

func (o *CreateGitWebhookUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGitWebhookUsingPOSTForbidden creates a CreateGitWebhookUsingPOSTForbidden with default headers values
func NewCreateGitWebhookUsingPOSTForbidden() *CreateGitWebhookUsingPOSTForbidden {
	return &CreateGitWebhookUsingPOSTForbidden{}
}

/*
CreateGitWebhookUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateGitWebhookUsingPOSTForbidden struct {
}

// IsSuccess returns true when this create git webhook using p o s t forbidden response has a 2xx status code
func (o *CreateGitWebhookUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create git webhook using p o s t forbidden response has a 3xx status code
func (o *CreateGitWebhookUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create git webhook using p o s t forbidden response has a 4xx status code
func (o *CreateGitWebhookUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create git webhook using p o s t forbidden response has a 5xx status code
func (o *CreateGitWebhookUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create git webhook using p o s t forbidden response a status code equal to that given
func (o *CreateGitWebhookUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateGitWebhookUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createGitWebhookUsingPOSTForbidden ", 403)
}

func (o *CreateGitWebhookUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createGitWebhookUsingPOSTForbidden ", 403)
}

func (o *CreateGitWebhookUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGitWebhookUsingPOSTNotFound creates a CreateGitWebhookUsingPOSTNotFound with default headers values
func NewCreateGitWebhookUsingPOSTNotFound() *CreateGitWebhookUsingPOSTNotFound {
	return &CreateGitWebhookUsingPOSTNotFound{}
}

/*
CreateGitWebhookUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateGitWebhookUsingPOSTNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this create git webhook using p o s t not found response has a 2xx status code
func (o *CreateGitWebhookUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create git webhook using p o s t not found response has a 3xx status code
func (o *CreateGitWebhookUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create git webhook using p o s t not found response has a 4xx status code
func (o *CreateGitWebhookUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create git webhook using p o s t not found response has a 5xx status code
func (o *CreateGitWebhookUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create git webhook using p o s t not found response a status code equal to that given
func (o *CreateGitWebhookUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateGitWebhookUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createGitWebhookUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *CreateGitWebhookUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createGitWebhookUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *CreateGitWebhookUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGitWebhookUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGitWebhookUsingPOSTInternalServerError creates a CreateGitWebhookUsingPOSTInternalServerError with default headers values
func NewCreateGitWebhookUsingPOSTInternalServerError() *CreateGitWebhookUsingPOSTInternalServerError {
	return &CreateGitWebhookUsingPOSTInternalServerError{}
}

/*
CreateGitWebhookUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CreateGitWebhookUsingPOSTInternalServerError struct {
}

// IsSuccess returns true when this create git webhook using p o s t internal server error response has a 2xx status code
func (o *CreateGitWebhookUsingPOSTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create git webhook using p o s t internal server error response has a 3xx status code
func (o *CreateGitWebhookUsingPOSTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create git webhook using p o s t internal server error response has a 4xx status code
func (o *CreateGitWebhookUsingPOSTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create git webhook using p o s t internal server error response has a 5xx status code
func (o *CreateGitWebhookUsingPOSTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create git webhook using p o s t internal server error response a status code equal to that given
func (o *CreateGitWebhookUsingPOSTInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateGitWebhookUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createGitWebhookUsingPOSTInternalServerError ", 500)
}

func (o *CreateGitWebhookUsingPOSTInternalServerError) String() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createGitWebhookUsingPOSTInternalServerError ", 500)
}

func (o *CreateGitWebhookUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}