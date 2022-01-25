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

// GetGitEventByIDUsingGETReader is a Reader for the GetGitEventByIDUsingGET structure.
type GetGitEventByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGitEventByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGitEventByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGitEventByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGitEventByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGitEventByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGitEventByIDUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGitEventByIDUsingGETOK creates a GetGitEventByIDUsingGETOK with default headers values
func NewGetGitEventByIDUsingGETOK() *GetGitEventByIDUsingGETOK {
	return &GetGitEventByIDUsingGETOK{}
}

/* GetGitEventByIDUsingGETOK describes a response with status code 200, with default header values.

'Success' with Git Event
*/
type GetGitEventByIDUsingGETOK struct {
	Payload models.GitEvent
}

func (o *GetGitEventByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-events/{id}][%d] getGitEventByIdUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetGitEventByIDUsingGETOK) GetPayload() models.GitEvent {
	return o.Payload
}

func (o *GetGitEventByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGitEvent(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetGitEventByIDUsingGETUnauthorized creates a GetGitEventByIDUsingGETUnauthorized with default headers values
func NewGetGitEventByIDUsingGETUnauthorized() *GetGitEventByIDUsingGETUnauthorized {
	return &GetGitEventByIDUsingGETUnauthorized{}
}

/* GetGitEventByIDUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetGitEventByIDUsingGETUnauthorized struct {
}

func (o *GetGitEventByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-events/{id}][%d] getGitEventByIdUsingGETUnauthorized ", 401)
}

func (o *GetGitEventByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGitEventByIDUsingGETForbidden creates a GetGitEventByIDUsingGETForbidden with default headers values
func NewGetGitEventByIDUsingGETForbidden() *GetGitEventByIDUsingGETForbidden {
	return &GetGitEventByIDUsingGETForbidden{}
}

/* GetGitEventByIDUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetGitEventByIDUsingGETForbidden struct {
}

func (o *GetGitEventByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-events/{id}][%d] getGitEventByIdUsingGETForbidden ", 403)
}

func (o *GetGitEventByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGitEventByIDUsingGETNotFound creates a GetGitEventByIDUsingGETNotFound with default headers values
func NewGetGitEventByIDUsingGETNotFound() *GetGitEventByIDUsingGETNotFound {
	return &GetGitEventByIDUsingGETNotFound{}
}

/* GetGitEventByIDUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetGitEventByIDUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetGitEventByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-events/{id}][%d] getGitEventByIdUsingGETNotFound  %+v", 404, o.Payload)
}
func (o *GetGitEventByIDUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGitEventByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGitEventByIDUsingGETInternalServerError creates a GetGitEventByIDUsingGETInternalServerError with default headers values
func NewGetGitEventByIDUsingGETInternalServerError() *GetGitEventByIDUsingGETInternalServerError {
	return &GetGitEventByIDUsingGETInternalServerError{}
}

/* GetGitEventByIDUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetGitEventByIDUsingGETInternalServerError struct {
}

func (o *GetGitEventByIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-events/{id}][%d] getGitEventByIdUsingGETInternalServerError ", 500)
}

func (o *GetGitEventByIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
