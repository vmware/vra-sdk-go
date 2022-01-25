// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetEndpointByIDUsingGETReader is a Reader for the GetEndpointByIDUsingGET structure.
type GetEndpointByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEndpointByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEndpointByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEndpointByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEndpointByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEndpointByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEndpointByIDUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEndpointByIDUsingGETOK creates a GetEndpointByIDUsingGETOK with default headers values
func NewGetEndpointByIDUsingGETOK() *GetEndpointByIDUsingGETOK {
	return &GetEndpointByIDUsingGETOK{}
}

/* GetEndpointByIDUsingGETOK describes a response with status code 200, with default header values.

'Success' with the requested Endpoint
*/
type GetEndpointByIDUsingGETOK struct {
	Payload models.Endpoint
}

func (o *GetEndpointByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getEndpointByIdUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetEndpointByIDUsingGETOK) GetPayload() models.Endpoint {
	return o.Payload
}

func (o *GetEndpointByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalEndpoint(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetEndpointByIDUsingGETUnauthorized creates a GetEndpointByIDUsingGETUnauthorized with default headers values
func NewGetEndpointByIDUsingGETUnauthorized() *GetEndpointByIDUsingGETUnauthorized {
	return &GetEndpointByIDUsingGETUnauthorized{}
}

/* GetEndpointByIDUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetEndpointByIDUsingGETUnauthorized struct {
}

func (o *GetEndpointByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getEndpointByIdUsingGETUnauthorized ", 401)
}

func (o *GetEndpointByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointByIDUsingGETForbidden creates a GetEndpointByIDUsingGETForbidden with default headers values
func NewGetEndpointByIDUsingGETForbidden() *GetEndpointByIDUsingGETForbidden {
	return &GetEndpointByIDUsingGETForbidden{}
}

/* GetEndpointByIDUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetEndpointByIDUsingGETForbidden struct {
}

func (o *GetEndpointByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getEndpointByIdUsingGETForbidden ", 403)
}

func (o *GetEndpointByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointByIDUsingGETNotFound creates a GetEndpointByIDUsingGETNotFound with default headers values
func NewGetEndpointByIDUsingGETNotFound() *GetEndpointByIDUsingGETNotFound {
	return &GetEndpointByIDUsingGETNotFound{}
}

/* GetEndpointByIDUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetEndpointByIDUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetEndpointByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getEndpointByIdUsingGETNotFound  %+v", 404, o.Payload)
}
func (o *GetEndpointByIDUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEndpointByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointByIDUsingGETInternalServerError creates a GetEndpointByIDUsingGETInternalServerError with default headers values
func NewGetEndpointByIDUsingGETInternalServerError() *GetEndpointByIDUsingGETInternalServerError {
	return &GetEndpointByIDUsingGETInternalServerError{}
}

/* GetEndpointByIDUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetEndpointByIDUsingGETInternalServerError struct {
}

func (o *GetEndpointByIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getEndpointByIdUsingGETInternalServerError ", 500)
}

func (o *GetEndpointByIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
