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

// GetByNameUsingGET2Reader is a Reader for the GetByNameUsingGET2 structure.
type GetByNameUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetByNameUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetByNameUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetByNameUsingGET2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetByNameUsingGET2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetByNameUsingGET2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetByNameUsingGET2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetByNameUsingGET2OK creates a GetByNameUsingGET2OK with default headers values
func NewGetByNameUsingGET2OK() *GetByNameUsingGET2OK {
	return &GetByNameUsingGET2OK{}
}

/* GetByNameUsingGET2OK describes a response with status code 200, with default header values.

'Success' with gerrit trigger
*/
type GetByNameUsingGET2OK struct {
	Payload models.GerritTrigger
}

func (o *GetByNameUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getByNameUsingGET2OK  %+v", 200, o.Payload)
}
func (o *GetByNameUsingGET2OK) GetPayload() models.GerritTrigger {
	return o.Payload
}

func (o *GetByNameUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritTrigger(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetByNameUsingGET2Unauthorized creates a GetByNameUsingGET2Unauthorized with default headers values
func NewGetByNameUsingGET2Unauthorized() *GetByNameUsingGET2Unauthorized {
	return &GetByNameUsingGET2Unauthorized{}
}

/* GetByNameUsingGET2Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetByNameUsingGET2Unauthorized struct {
}

func (o *GetByNameUsingGET2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getByNameUsingGET2Unauthorized ", 401)
}

func (o *GetByNameUsingGET2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetByNameUsingGET2Forbidden creates a GetByNameUsingGET2Forbidden with default headers values
func NewGetByNameUsingGET2Forbidden() *GetByNameUsingGET2Forbidden {
	return &GetByNameUsingGET2Forbidden{}
}

/* GetByNameUsingGET2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetByNameUsingGET2Forbidden struct {
}

func (o *GetByNameUsingGET2Forbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getByNameUsingGET2Forbidden ", 403)
}

func (o *GetByNameUsingGET2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetByNameUsingGET2NotFound creates a GetByNameUsingGET2NotFound with default headers values
func NewGetByNameUsingGET2NotFound() *GetByNameUsingGET2NotFound {
	return &GetByNameUsingGET2NotFound{}
}

/* GetByNameUsingGET2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetByNameUsingGET2NotFound struct {
	Payload *models.Error
}

func (o *GetByNameUsingGET2NotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getByNameUsingGET2NotFound  %+v", 404, o.Payload)
}
func (o *GetByNameUsingGET2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetByNameUsingGET2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetByNameUsingGET2InternalServerError creates a GetByNameUsingGET2InternalServerError with default headers values
func NewGetByNameUsingGET2InternalServerError() *GetByNameUsingGET2InternalServerError {
	return &GetByNameUsingGET2InternalServerError{}
}

/* GetByNameUsingGET2InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetByNameUsingGET2InternalServerError struct {
}

func (o *GetByNameUsingGET2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getByNameUsingGET2InternalServerError ", 500)
}

func (o *GetByNameUsingGET2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
