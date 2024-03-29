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

// PatchDockerRegistryWebhookByNameUsingPATCHReader is a Reader for the PatchDockerRegistryWebhookByNameUsingPATCH structure.
type PatchDockerRegistryWebhookByNameUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDockerRegistryWebhookByNameUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchDockerRegistryWebhookByNameUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchDockerRegistryWebhookByNameUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchDockerRegistryWebhookByNameUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchDockerRegistryWebhookByNameUsingPATCHNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchDockerRegistryWebhookByNameUsingPATCHInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchDockerRegistryWebhookByNameUsingPATCHOK creates a PatchDockerRegistryWebhookByNameUsingPATCHOK with default headers values
func NewPatchDockerRegistryWebhookByNameUsingPATCHOK() *PatchDockerRegistryWebhookByNameUsingPATCHOK {
	return &PatchDockerRegistryWebhookByNameUsingPATCHOK{}
}

/*
PatchDockerRegistryWebhookByNameUsingPATCHOK describes a response with status code 200, with default header values.

'Success' with Docker Registry Webhook patch
*/
type PatchDockerRegistryWebhookByNameUsingPATCHOK struct {
	Payload models.DockerRegistryWebHook
}

// IsSuccess returns true when this patch docker registry webhook by name using p a t c h o k response has a 2xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch docker registry webhook by name using p a t c h o k response has a 3xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch docker registry webhook by name using p a t c h o k response has a 4xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch docker registry webhook by name using p a t c h o k response has a 5xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch docker registry webhook by name using p a t c h o k response a status code equal to that given
func (o *PatchDockerRegistryWebhookByNameUsingPATCHOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchDockerRegistryWebhookByNameUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHOK) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchDockerRegistryWebhookByNameUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHOK) GetPayload() models.DockerRegistryWebHook {
	return o.Payload
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalDockerRegistryWebHook(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewPatchDockerRegistryWebhookByNameUsingPATCHUnauthorized creates a PatchDockerRegistryWebhookByNameUsingPATCHUnauthorized with default headers values
func NewPatchDockerRegistryWebhookByNameUsingPATCHUnauthorized() *PatchDockerRegistryWebhookByNameUsingPATCHUnauthorized {
	return &PatchDockerRegistryWebhookByNameUsingPATCHUnauthorized{}
}

/*
PatchDockerRegistryWebhookByNameUsingPATCHUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type PatchDockerRegistryWebhookByNameUsingPATCHUnauthorized struct {
}

// IsSuccess returns true when this patch docker registry webhook by name using p a t c h unauthorized response has a 2xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch docker registry webhook by name using p a t c h unauthorized response has a 3xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch docker registry webhook by name using p a t c h unauthorized response has a 4xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch docker registry webhook by name using p a t c h unauthorized response has a 5xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch docker registry webhook by name using p a t c h unauthorized response a status code equal to that given
func (o *PatchDockerRegistryWebhookByNameUsingPATCHUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchDockerRegistryWebhookByNameUsingPATCHUnauthorized ", 401)
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchDockerRegistryWebhookByNameUsingPATCHUnauthorized ", 401)
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchDockerRegistryWebhookByNameUsingPATCHForbidden creates a PatchDockerRegistryWebhookByNameUsingPATCHForbidden with default headers values
func NewPatchDockerRegistryWebhookByNameUsingPATCHForbidden() *PatchDockerRegistryWebhookByNameUsingPATCHForbidden {
	return &PatchDockerRegistryWebhookByNameUsingPATCHForbidden{}
}

/*
PatchDockerRegistryWebhookByNameUsingPATCHForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchDockerRegistryWebhookByNameUsingPATCHForbidden struct {
}

// IsSuccess returns true when this patch docker registry webhook by name using p a t c h forbidden response has a 2xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch docker registry webhook by name using p a t c h forbidden response has a 3xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch docker registry webhook by name using p a t c h forbidden response has a 4xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch docker registry webhook by name using p a t c h forbidden response has a 5xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch docker registry webhook by name using p a t c h forbidden response a status code equal to that given
func (o *PatchDockerRegistryWebhookByNameUsingPATCHForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchDockerRegistryWebhookByNameUsingPATCHForbidden ", 403)
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHForbidden) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchDockerRegistryWebhookByNameUsingPATCHForbidden ", 403)
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchDockerRegistryWebhookByNameUsingPATCHNotFound creates a PatchDockerRegistryWebhookByNameUsingPATCHNotFound with default headers values
func NewPatchDockerRegistryWebhookByNameUsingPATCHNotFound() *PatchDockerRegistryWebhookByNameUsingPATCHNotFound {
	return &PatchDockerRegistryWebhookByNameUsingPATCHNotFound{}
}

/*
PatchDockerRegistryWebhookByNameUsingPATCHNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchDockerRegistryWebhookByNameUsingPATCHNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this patch docker registry webhook by name using p a t c h not found response has a 2xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch docker registry webhook by name using p a t c h not found response has a 3xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch docker registry webhook by name using p a t c h not found response has a 4xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch docker registry webhook by name using p a t c h not found response has a 5xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch docker registry webhook by name using p a t c h not found response a status code equal to that given
func (o *PatchDockerRegistryWebhookByNameUsingPATCHNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHNotFound) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchDockerRegistryWebhookByNameUsingPATCHNotFound  %+v", 404, o.Payload)
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHNotFound) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchDockerRegistryWebhookByNameUsingPATCHNotFound  %+v", 404, o.Payload)
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDockerRegistryWebhookByNameUsingPATCHInternalServerError creates a PatchDockerRegistryWebhookByNameUsingPATCHInternalServerError with default headers values
func NewPatchDockerRegistryWebhookByNameUsingPATCHInternalServerError() *PatchDockerRegistryWebhookByNameUsingPATCHInternalServerError {
	return &PatchDockerRegistryWebhookByNameUsingPATCHInternalServerError{}
}

/*
PatchDockerRegistryWebhookByNameUsingPATCHInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PatchDockerRegistryWebhookByNameUsingPATCHInternalServerError struct {
}

// IsSuccess returns true when this patch docker registry webhook by name using p a t c h internal server error response has a 2xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch docker registry webhook by name using p a t c h internal server error response has a 3xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch docker registry webhook by name using p a t c h internal server error response has a 4xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch docker registry webhook by name using p a t c h internal server error response has a 5xx status code
func (o *PatchDockerRegistryWebhookByNameUsingPATCHInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch docker registry webhook by name using p a t c h internal server error response a status code equal to that given
func (o *PatchDockerRegistryWebhookByNameUsingPATCHInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchDockerRegistryWebhookByNameUsingPATCHInternalServerError ", 500)
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchDockerRegistryWebhookByNameUsingPATCHInternalServerError ", 500)
}

func (o *PatchDockerRegistryWebhookByNameUsingPATCHInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
