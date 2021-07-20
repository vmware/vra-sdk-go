// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetUsingGET5Reader is a Reader for the GetUsingGET5 structure.
type GetUsingGET5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsingGET5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsingGET5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsingGET5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsingGET5Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsingGET5NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsingGET5InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsingGET5OK creates a GetUsingGET5OK with default headers values
func NewGetUsingGET5OK() *GetUsingGET5OK {
	return &GetUsingGET5OK{}
}

/* GetUsingGET5OK describes a response with status code 200, with default header values.

'Success' with the requested Variable
*/
type GetUsingGET5OK struct {
	Payload *models.Variable
}

func (o *GetUsingGET5OK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables/{id}][%d] getUsingGET5OK  %+v", 200, o.Payload)
}
func (o *GetUsingGET5OK) GetPayload() *models.Variable {
	return o.Payload
}

func (o *GetUsingGET5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Variable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsingGET5Unauthorized creates a GetUsingGET5Unauthorized with default headers values
func NewGetUsingGET5Unauthorized() *GetUsingGET5Unauthorized {
	return &GetUsingGET5Unauthorized{}
}

/* GetUsingGET5Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetUsingGET5Unauthorized struct {
}

func (o *GetUsingGET5Unauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables/{id}][%d] getUsingGET5Unauthorized ", 401)
}

func (o *GetUsingGET5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGET5Forbidden creates a GetUsingGET5Forbidden with default headers values
func NewGetUsingGET5Forbidden() *GetUsingGET5Forbidden {
	return &GetUsingGET5Forbidden{}
}

/* GetUsingGET5Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUsingGET5Forbidden struct {
}

func (o *GetUsingGET5Forbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables/{id}][%d] getUsingGET5Forbidden ", 403)
}

func (o *GetUsingGET5Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGET5NotFound creates a GetUsingGET5NotFound with default headers values
func NewGetUsingGET5NotFound() *GetUsingGET5NotFound {
	return &GetUsingGET5NotFound{}
}

/* GetUsingGET5NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetUsingGET5NotFound struct {
	Payload *models.Error
}

func (o *GetUsingGET5NotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables/{id}][%d] getUsingGET5NotFound  %+v", 404, o.Payload)
}
func (o *GetUsingGET5NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUsingGET5NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsingGET5InternalServerError creates a GetUsingGET5InternalServerError with default headers values
func NewGetUsingGET5InternalServerError() *GetUsingGET5InternalServerError {
	return &GetUsingGET5InternalServerError{}
}

/* GetUsingGET5InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetUsingGET5InternalServerError struct {
}

func (o *GetUsingGET5InternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables/{id}][%d] getUsingGET5InternalServerError ", 500)
}

func (o *GetUsingGET5InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
