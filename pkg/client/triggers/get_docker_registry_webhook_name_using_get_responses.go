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

// GetDockerRegistryWebhookNameUsingGETReader is a Reader for the GetDockerRegistryWebhookNameUsingGET structure.
type GetDockerRegistryWebhookNameUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDockerRegistryWebhookNameUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDockerRegistryWebhookNameUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDockerRegistryWebhookNameUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDockerRegistryWebhookNameUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDockerRegistryWebhookNameUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDockerRegistryWebhookNameUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDockerRegistryWebhookNameUsingGETOK creates a GetDockerRegistryWebhookNameUsingGETOK with default headers values
func NewGetDockerRegistryWebhookNameUsingGETOK() *GetDockerRegistryWebhookNameUsingGETOK {
	return &GetDockerRegistryWebhookNameUsingGETOK{}
}

/*
GetDockerRegistryWebhookNameUsingGETOK describes a response with status code 200, with default header values.

'Success' with Docker Registry Webhook
*/
type GetDockerRegistryWebhookNameUsingGETOK struct {
	Payload models.DockerRegistryWebHook
}

// IsSuccess returns true when this get docker registry webhook name using g e t o k response has a 2xx status code
func (o *GetDockerRegistryWebhookNameUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get docker registry webhook name using g e t o k response has a 3xx status code
func (o *GetDockerRegistryWebhookNameUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get docker registry webhook name using g e t o k response has a 4xx status code
func (o *GetDockerRegistryWebhookNameUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get docker registry webhook name using g e t o k response has a 5xx status code
func (o *GetDockerRegistryWebhookNameUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get docker registry webhook name using g e t o k response a status code equal to that given
func (o *GetDockerRegistryWebhookNameUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDockerRegistryWebhookNameUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks/{project}/{name}][%d] getDockerRegistryWebhookNameUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetDockerRegistryWebhookNameUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks/{project}/{name}][%d] getDockerRegistryWebhookNameUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetDockerRegistryWebhookNameUsingGETOK) GetPayload() models.DockerRegistryWebHook {
	return o.Payload
}

func (o *GetDockerRegistryWebhookNameUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalDockerRegistryWebHook(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetDockerRegistryWebhookNameUsingGETUnauthorized creates a GetDockerRegistryWebhookNameUsingGETUnauthorized with default headers values
func NewGetDockerRegistryWebhookNameUsingGETUnauthorized() *GetDockerRegistryWebhookNameUsingGETUnauthorized {
	return &GetDockerRegistryWebhookNameUsingGETUnauthorized{}
}

/*
GetDockerRegistryWebhookNameUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetDockerRegistryWebhookNameUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get docker registry webhook name using g e t unauthorized response has a 2xx status code
func (o *GetDockerRegistryWebhookNameUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get docker registry webhook name using g e t unauthorized response has a 3xx status code
func (o *GetDockerRegistryWebhookNameUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get docker registry webhook name using g e t unauthorized response has a 4xx status code
func (o *GetDockerRegistryWebhookNameUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get docker registry webhook name using g e t unauthorized response has a 5xx status code
func (o *GetDockerRegistryWebhookNameUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get docker registry webhook name using g e t unauthorized response a status code equal to that given
func (o *GetDockerRegistryWebhookNameUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDockerRegistryWebhookNameUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks/{project}/{name}][%d] getDockerRegistryWebhookNameUsingGETUnauthorized ", 401)
}

func (o *GetDockerRegistryWebhookNameUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks/{project}/{name}][%d] getDockerRegistryWebhookNameUsingGETUnauthorized ", 401)
}

func (o *GetDockerRegistryWebhookNameUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDockerRegistryWebhookNameUsingGETForbidden creates a GetDockerRegistryWebhookNameUsingGETForbidden with default headers values
func NewGetDockerRegistryWebhookNameUsingGETForbidden() *GetDockerRegistryWebhookNameUsingGETForbidden {
	return &GetDockerRegistryWebhookNameUsingGETForbidden{}
}

/*
GetDockerRegistryWebhookNameUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDockerRegistryWebhookNameUsingGETForbidden struct {
}

// IsSuccess returns true when this get docker registry webhook name using g e t forbidden response has a 2xx status code
func (o *GetDockerRegistryWebhookNameUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get docker registry webhook name using g e t forbidden response has a 3xx status code
func (o *GetDockerRegistryWebhookNameUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get docker registry webhook name using g e t forbidden response has a 4xx status code
func (o *GetDockerRegistryWebhookNameUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get docker registry webhook name using g e t forbidden response has a 5xx status code
func (o *GetDockerRegistryWebhookNameUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get docker registry webhook name using g e t forbidden response a status code equal to that given
func (o *GetDockerRegistryWebhookNameUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDockerRegistryWebhookNameUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks/{project}/{name}][%d] getDockerRegistryWebhookNameUsingGETForbidden ", 403)
}

func (o *GetDockerRegistryWebhookNameUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks/{project}/{name}][%d] getDockerRegistryWebhookNameUsingGETForbidden ", 403)
}

func (o *GetDockerRegistryWebhookNameUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDockerRegistryWebhookNameUsingGETNotFound creates a GetDockerRegistryWebhookNameUsingGETNotFound with default headers values
func NewGetDockerRegistryWebhookNameUsingGETNotFound() *GetDockerRegistryWebhookNameUsingGETNotFound {
	return &GetDockerRegistryWebhookNameUsingGETNotFound{}
}

/*
GetDockerRegistryWebhookNameUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDockerRegistryWebhookNameUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get docker registry webhook name using g e t not found response has a 2xx status code
func (o *GetDockerRegistryWebhookNameUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get docker registry webhook name using g e t not found response has a 3xx status code
func (o *GetDockerRegistryWebhookNameUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get docker registry webhook name using g e t not found response has a 4xx status code
func (o *GetDockerRegistryWebhookNameUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get docker registry webhook name using g e t not found response has a 5xx status code
func (o *GetDockerRegistryWebhookNameUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get docker registry webhook name using g e t not found response a status code equal to that given
func (o *GetDockerRegistryWebhookNameUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDockerRegistryWebhookNameUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks/{project}/{name}][%d] getDockerRegistryWebhookNameUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetDockerRegistryWebhookNameUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks/{project}/{name}][%d] getDockerRegistryWebhookNameUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetDockerRegistryWebhookNameUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDockerRegistryWebhookNameUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDockerRegistryWebhookNameUsingGETInternalServerError creates a GetDockerRegistryWebhookNameUsingGETInternalServerError with default headers values
func NewGetDockerRegistryWebhookNameUsingGETInternalServerError() *GetDockerRegistryWebhookNameUsingGETInternalServerError {
	return &GetDockerRegistryWebhookNameUsingGETInternalServerError{}
}

/*
GetDockerRegistryWebhookNameUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetDockerRegistryWebhookNameUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get docker registry webhook name using g e t internal server error response has a 2xx status code
func (o *GetDockerRegistryWebhookNameUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get docker registry webhook name using g e t internal server error response has a 3xx status code
func (o *GetDockerRegistryWebhookNameUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get docker registry webhook name using g e t internal server error response has a 4xx status code
func (o *GetDockerRegistryWebhookNameUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get docker registry webhook name using g e t internal server error response has a 5xx status code
func (o *GetDockerRegistryWebhookNameUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get docker registry webhook name using g e t internal server error response a status code equal to that given
func (o *GetDockerRegistryWebhookNameUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDockerRegistryWebhookNameUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks/{project}/{name}][%d] getDockerRegistryWebhookNameUsingGETInternalServerError ", 500)
}

func (o *GetDockerRegistryWebhookNameUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks/{project}/{name}][%d] getDockerRegistryWebhookNameUsingGETInternalServerError ", 500)
}

func (o *GetDockerRegistryWebhookNameUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
