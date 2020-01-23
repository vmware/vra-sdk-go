// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetEndpointPropertiesUsingGETReader is a Reader for the GetEndpointPropertiesUsingGET structure.
type GetEndpointPropertiesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEndpointPropertiesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEndpointPropertiesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEndpointPropertiesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEndpointPropertiesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEndpointPropertiesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEndpointPropertiesUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEndpointPropertiesUsingGETOK creates a GetEndpointPropertiesUsingGETOK with default headers values
func NewGetEndpointPropertiesUsingGETOK() *GetEndpointPropertiesUsingGETOK {
	return &GetEndpointPropertiesUsingGETOK{}
}

/*GetEndpointPropertiesUsingGETOK handles this case with default header values.

'Success' with endpoint properties
*/
type GetEndpointPropertiesUsingGETOK struct {
	Payload interface{}
}

func (o *GetEndpointPropertiesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-tiles/{type}][%d] getEndpointPropertiesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetEndpointPropertiesUsingGETOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetEndpointPropertiesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointPropertiesUsingGETUnauthorized creates a GetEndpointPropertiesUsingGETUnauthorized with default headers values
func NewGetEndpointPropertiesUsingGETUnauthorized() *GetEndpointPropertiesUsingGETUnauthorized {
	return &GetEndpointPropertiesUsingGETUnauthorized{}
}

/*GetEndpointPropertiesUsingGETUnauthorized handles this case with default header values.

Unauthorized Request
*/
type GetEndpointPropertiesUsingGETUnauthorized struct {
}

func (o *GetEndpointPropertiesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-tiles/{type}][%d] getEndpointPropertiesUsingGETUnauthorized ", 401)
}

func (o *GetEndpointPropertiesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointPropertiesUsingGETForbidden creates a GetEndpointPropertiesUsingGETForbidden with default headers values
func NewGetEndpointPropertiesUsingGETForbidden() *GetEndpointPropertiesUsingGETForbidden {
	return &GetEndpointPropertiesUsingGETForbidden{}
}

/*GetEndpointPropertiesUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetEndpointPropertiesUsingGETForbidden struct {
}

func (o *GetEndpointPropertiesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-tiles/{type}][%d] getEndpointPropertiesUsingGETForbidden ", 403)
}

func (o *GetEndpointPropertiesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointPropertiesUsingGETNotFound creates a GetEndpointPropertiesUsingGETNotFound with default headers values
func NewGetEndpointPropertiesUsingGETNotFound() *GetEndpointPropertiesUsingGETNotFound {
	return &GetEndpointPropertiesUsingGETNotFound{}
}

/*GetEndpointPropertiesUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetEndpointPropertiesUsingGETNotFound struct {
}

func (o *GetEndpointPropertiesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-tiles/{type}][%d] getEndpointPropertiesUsingGETNotFound ", 404)
}

func (o *GetEndpointPropertiesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointPropertiesUsingGETInternalServerError creates a GetEndpointPropertiesUsingGETInternalServerError with default headers values
func NewGetEndpointPropertiesUsingGETInternalServerError() *GetEndpointPropertiesUsingGETInternalServerError {
	return &GetEndpointPropertiesUsingGETInternalServerError{}
}

/*GetEndpointPropertiesUsingGETInternalServerError handles this case with default header values.

Server Error
*/
type GetEndpointPropertiesUsingGETInternalServerError struct {
}

func (o *GetEndpointPropertiesUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-tiles/{type}][%d] getEndpointPropertiesUsingGETInternalServerError ", 500)
}

func (o *GetEndpointPropertiesUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
