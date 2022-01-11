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

// GetAllGitWebhooksUsingGETReader is a Reader for the GetAllGitWebhooksUsingGET structure.
type GetAllGitWebhooksUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllGitWebhooksUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllGitWebhooksUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllGitWebhooksUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllGitWebhooksUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllGitWebhooksUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllGitWebhooksUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllGitWebhooksUsingGETOK creates a GetAllGitWebhooksUsingGETOK with default headers values
func NewGetAllGitWebhooksUsingGETOK() *GetAllGitWebhooksUsingGETOK {
	return &GetAllGitWebhooksUsingGETOK{}
}

/* GetAllGitWebhooksUsingGETOK describes a response with status code 200, with default header values.

'Success' with get of Git Webhooks
*/
type GetAllGitWebhooksUsingGETOK struct {
	Payload models.GitWebhooks
}

func (o *GetAllGitWebhooksUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-webhooks][%d] getAllGitWebhooksUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetAllGitWebhooksUsingGETOK) GetPayload() models.GitWebhooks {
	return o.Payload
}

func (o *GetAllGitWebhooksUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGitWebhooks(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetAllGitWebhooksUsingGETUnauthorized creates a GetAllGitWebhooksUsingGETUnauthorized with default headers values
func NewGetAllGitWebhooksUsingGETUnauthorized() *GetAllGitWebhooksUsingGETUnauthorized {
	return &GetAllGitWebhooksUsingGETUnauthorized{}
}

/* GetAllGitWebhooksUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetAllGitWebhooksUsingGETUnauthorized struct {
}

func (o *GetAllGitWebhooksUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-webhooks][%d] getAllGitWebhooksUsingGETUnauthorized ", 401)
}

func (o *GetAllGitWebhooksUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllGitWebhooksUsingGETForbidden creates a GetAllGitWebhooksUsingGETForbidden with default headers values
func NewGetAllGitWebhooksUsingGETForbidden() *GetAllGitWebhooksUsingGETForbidden {
	return &GetAllGitWebhooksUsingGETForbidden{}
}

/* GetAllGitWebhooksUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllGitWebhooksUsingGETForbidden struct {
}

func (o *GetAllGitWebhooksUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-webhooks][%d] getAllGitWebhooksUsingGETForbidden ", 403)
}

func (o *GetAllGitWebhooksUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllGitWebhooksUsingGETNotFound creates a GetAllGitWebhooksUsingGETNotFound with default headers values
func NewGetAllGitWebhooksUsingGETNotFound() *GetAllGitWebhooksUsingGETNotFound {
	return &GetAllGitWebhooksUsingGETNotFound{}
}

/* GetAllGitWebhooksUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllGitWebhooksUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetAllGitWebhooksUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-webhooks][%d] getAllGitWebhooksUsingGETNotFound  %+v", 404, o.Payload)
}
func (o *GetAllGitWebhooksUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllGitWebhooksUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllGitWebhooksUsingGETInternalServerError creates a GetAllGitWebhooksUsingGETInternalServerError with default headers values
func NewGetAllGitWebhooksUsingGETInternalServerError() *GetAllGitWebhooksUsingGETInternalServerError {
	return &GetAllGitWebhooksUsingGETInternalServerError{}
}

/* GetAllGitWebhooksUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetAllGitWebhooksUsingGETInternalServerError struct {
}

func (o *GetAllGitWebhooksUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-webhooks][%d] getAllGitWebhooksUsingGETInternalServerError ", 500)
}

func (o *GetAllGitWebhooksUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
