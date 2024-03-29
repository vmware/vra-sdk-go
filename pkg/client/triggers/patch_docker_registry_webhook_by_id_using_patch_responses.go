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

// PatchDockerRegistryWebhookByIDUsingPATCHReader is a Reader for the PatchDockerRegistryWebhookByIDUsingPATCH structure.
type PatchDockerRegistryWebhookByIDUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDockerRegistryWebhookByIDUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchDockerRegistryWebhookByIDUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchDockerRegistryWebhookByIDUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchDockerRegistryWebhookByIDUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchDockerRegistryWebhookByIDUsingPATCHNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchDockerRegistryWebhookByIDUsingPATCHInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchDockerRegistryWebhookByIDUsingPATCHOK creates a PatchDockerRegistryWebhookByIDUsingPATCHOK with default headers values
func NewPatchDockerRegistryWebhookByIDUsingPATCHOK() *PatchDockerRegistryWebhookByIDUsingPATCHOK {
	return &PatchDockerRegistryWebhookByIDUsingPATCHOK{}
}

/*
PatchDockerRegistryWebhookByIDUsingPATCHOK describes a response with status code 200, with default header values.

'Success' with Docker Registry Webhook patch
*/
type PatchDockerRegistryWebhookByIDUsingPATCHOK struct {
	Payload models.DockerRegistryWebHook
}

// IsSuccess returns true when this patch docker registry webhook by Id using p a t c h o k response has a 2xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch docker registry webhook by Id using p a t c h o k response has a 3xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch docker registry webhook by Id using p a t c h o k response has a 4xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch docker registry webhook by Id using p a t c h o k response has a 5xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch docker registry webhook by Id using p a t c h o k response a status code equal to that given
func (o *PatchDockerRegistryWebhookByIDUsingPATCHOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{id}][%d] patchDockerRegistryWebhookByIdUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHOK) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{id}][%d] patchDockerRegistryWebhookByIdUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHOK) GetPayload() models.DockerRegistryWebHook {
	return o.Payload
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalDockerRegistryWebHook(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewPatchDockerRegistryWebhookByIDUsingPATCHUnauthorized creates a PatchDockerRegistryWebhookByIDUsingPATCHUnauthorized with default headers values
func NewPatchDockerRegistryWebhookByIDUsingPATCHUnauthorized() *PatchDockerRegistryWebhookByIDUsingPATCHUnauthorized {
	return &PatchDockerRegistryWebhookByIDUsingPATCHUnauthorized{}
}

/*
PatchDockerRegistryWebhookByIDUsingPATCHUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type PatchDockerRegistryWebhookByIDUsingPATCHUnauthorized struct {
}

// IsSuccess returns true when this patch docker registry webhook by Id using p a t c h unauthorized response has a 2xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch docker registry webhook by Id using p a t c h unauthorized response has a 3xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch docker registry webhook by Id using p a t c h unauthorized response has a 4xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch docker registry webhook by Id using p a t c h unauthorized response has a 5xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch docker registry webhook by Id using p a t c h unauthorized response a status code equal to that given
func (o *PatchDockerRegistryWebhookByIDUsingPATCHUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{id}][%d] patchDockerRegistryWebhookByIdUsingPATCHUnauthorized ", 401)
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{id}][%d] patchDockerRegistryWebhookByIdUsingPATCHUnauthorized ", 401)
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchDockerRegistryWebhookByIDUsingPATCHForbidden creates a PatchDockerRegistryWebhookByIDUsingPATCHForbidden with default headers values
func NewPatchDockerRegistryWebhookByIDUsingPATCHForbidden() *PatchDockerRegistryWebhookByIDUsingPATCHForbidden {
	return &PatchDockerRegistryWebhookByIDUsingPATCHForbidden{}
}

/*
PatchDockerRegistryWebhookByIDUsingPATCHForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchDockerRegistryWebhookByIDUsingPATCHForbidden struct {
}

// IsSuccess returns true when this patch docker registry webhook by Id using p a t c h forbidden response has a 2xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch docker registry webhook by Id using p a t c h forbidden response has a 3xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch docker registry webhook by Id using p a t c h forbidden response has a 4xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch docker registry webhook by Id using p a t c h forbidden response has a 5xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch docker registry webhook by Id using p a t c h forbidden response a status code equal to that given
func (o *PatchDockerRegistryWebhookByIDUsingPATCHForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{id}][%d] patchDockerRegistryWebhookByIdUsingPATCHForbidden ", 403)
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHForbidden) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{id}][%d] patchDockerRegistryWebhookByIdUsingPATCHForbidden ", 403)
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchDockerRegistryWebhookByIDUsingPATCHNotFound creates a PatchDockerRegistryWebhookByIDUsingPATCHNotFound with default headers values
func NewPatchDockerRegistryWebhookByIDUsingPATCHNotFound() *PatchDockerRegistryWebhookByIDUsingPATCHNotFound {
	return &PatchDockerRegistryWebhookByIDUsingPATCHNotFound{}
}

/*
PatchDockerRegistryWebhookByIDUsingPATCHNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchDockerRegistryWebhookByIDUsingPATCHNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this patch docker registry webhook by Id using p a t c h not found response has a 2xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch docker registry webhook by Id using p a t c h not found response has a 3xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch docker registry webhook by Id using p a t c h not found response has a 4xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch docker registry webhook by Id using p a t c h not found response has a 5xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch docker registry webhook by Id using p a t c h not found response a status code equal to that given
func (o *PatchDockerRegistryWebhookByIDUsingPATCHNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHNotFound) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{id}][%d] patchDockerRegistryWebhookByIdUsingPATCHNotFound  %+v", 404, o.Payload)
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHNotFound) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{id}][%d] patchDockerRegistryWebhookByIdUsingPATCHNotFound  %+v", 404, o.Payload)
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDockerRegistryWebhookByIDUsingPATCHInternalServerError creates a PatchDockerRegistryWebhookByIDUsingPATCHInternalServerError with default headers values
func NewPatchDockerRegistryWebhookByIDUsingPATCHInternalServerError() *PatchDockerRegistryWebhookByIDUsingPATCHInternalServerError {
	return &PatchDockerRegistryWebhookByIDUsingPATCHInternalServerError{}
}

/*
PatchDockerRegistryWebhookByIDUsingPATCHInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PatchDockerRegistryWebhookByIDUsingPATCHInternalServerError struct {
}

// IsSuccess returns true when this patch docker registry webhook by Id using p a t c h internal server error response has a 2xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch docker registry webhook by Id using p a t c h internal server error response has a 3xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch docker registry webhook by Id using p a t c h internal server error response has a 4xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch docker registry webhook by Id using p a t c h internal server error response has a 5xx status code
func (o *PatchDockerRegistryWebhookByIDUsingPATCHInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch docker registry webhook by Id using p a t c h internal server error response a status code equal to that given
func (o *PatchDockerRegistryWebhookByIDUsingPATCHInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{id}][%d] patchDockerRegistryWebhookByIdUsingPATCHInternalServerError ", 500)
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{id}][%d] patchDockerRegistryWebhookByIdUsingPATCHInternalServerError ", 500)
}

func (o *PatchDockerRegistryWebhookByIDUsingPATCHInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
