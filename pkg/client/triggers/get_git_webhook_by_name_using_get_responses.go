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

// GetGitWebhookByNameUsingGETReader is a Reader for the GetGitWebhookByNameUsingGET structure.
type GetGitWebhookByNameUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGitWebhookByNameUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGitWebhookByNameUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGitWebhookByNameUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGitWebhookByNameUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGitWebhookByNameUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGitWebhookByNameUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGitWebhookByNameUsingGETOK creates a GetGitWebhookByNameUsingGETOK with default headers values
func NewGetGitWebhookByNameUsingGETOK() *GetGitWebhookByNameUsingGETOK {
	return &GetGitWebhookByNameUsingGETOK{}
}

/* GetGitWebhookByNameUsingGETOK describes a response with status code 200, with default header values.

'Success' with Git Webhook
*/
type GetGitWebhookByNameUsingGETOK struct {
	Payload models.GitWebhook
}

func (o *GetGitWebhookByNameUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-webhooks/{project}/{name}][%d] getGitWebhookByNameUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetGitWebhookByNameUsingGETOK) GetPayload() models.GitWebhook {
	return o.Payload
}

func (o *GetGitWebhookByNameUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGitWebhook(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetGitWebhookByNameUsingGETUnauthorized creates a GetGitWebhookByNameUsingGETUnauthorized with default headers values
func NewGetGitWebhookByNameUsingGETUnauthorized() *GetGitWebhookByNameUsingGETUnauthorized {
	return &GetGitWebhookByNameUsingGETUnauthorized{}
}

/* GetGitWebhookByNameUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetGitWebhookByNameUsingGETUnauthorized struct {
}

func (o *GetGitWebhookByNameUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-webhooks/{project}/{name}][%d] getGitWebhookByNameUsingGETUnauthorized ", 401)
}

func (o *GetGitWebhookByNameUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGitWebhookByNameUsingGETForbidden creates a GetGitWebhookByNameUsingGETForbidden with default headers values
func NewGetGitWebhookByNameUsingGETForbidden() *GetGitWebhookByNameUsingGETForbidden {
	return &GetGitWebhookByNameUsingGETForbidden{}
}

/* GetGitWebhookByNameUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetGitWebhookByNameUsingGETForbidden struct {
}

func (o *GetGitWebhookByNameUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-webhooks/{project}/{name}][%d] getGitWebhookByNameUsingGETForbidden ", 403)
}

func (o *GetGitWebhookByNameUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGitWebhookByNameUsingGETNotFound creates a GetGitWebhookByNameUsingGETNotFound with default headers values
func NewGetGitWebhookByNameUsingGETNotFound() *GetGitWebhookByNameUsingGETNotFound {
	return &GetGitWebhookByNameUsingGETNotFound{}
}

/* GetGitWebhookByNameUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetGitWebhookByNameUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetGitWebhookByNameUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-webhooks/{project}/{name}][%d] getGitWebhookByNameUsingGETNotFound  %+v", 404, o.Payload)
}
func (o *GetGitWebhookByNameUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGitWebhookByNameUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGitWebhookByNameUsingGETInternalServerError creates a GetGitWebhookByNameUsingGETInternalServerError with default headers values
func NewGetGitWebhookByNameUsingGETInternalServerError() *GetGitWebhookByNameUsingGETInternalServerError {
	return &GetGitWebhookByNameUsingGETInternalServerError{}
}

/* GetGitWebhookByNameUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetGitWebhookByNameUsingGETInternalServerError struct {
}

func (o *GetGitWebhookByNameUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-webhooks/{project}/{name}][%d] getGitWebhookByNameUsingGETInternalServerError ", 500)
}

func (o *GetGitWebhookByNameUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
