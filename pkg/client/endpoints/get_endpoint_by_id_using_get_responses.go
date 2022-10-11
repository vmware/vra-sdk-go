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

/*
GetEndpointByIDUsingGETOK describes a response with status code 200, with default header values.

'Success' with the requested Endpoint
*/
type GetEndpointByIDUsingGETOK struct {
	Payload models.Endpoint
}

// IsSuccess returns true when this get endpoint by Id using g e t o k response has a 2xx status code
func (o *GetEndpointByIDUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get endpoint by Id using g e t o k response has a 3xx status code
func (o *GetEndpointByIDUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint by Id using g e t o k response has a 4xx status code
func (o *GetEndpointByIDUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get endpoint by Id using g e t o k response has a 5xx status code
func (o *GetEndpointByIDUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint by Id using g e t o k response a status code equal to that given
func (o *GetEndpointByIDUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetEndpointByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getEndpointByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetEndpointByIDUsingGETOK) String() string {
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

/*
GetEndpointByIDUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetEndpointByIDUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get endpoint by Id using g e t unauthorized response has a 2xx status code
func (o *GetEndpointByIDUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint by Id using g e t unauthorized response has a 3xx status code
func (o *GetEndpointByIDUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint by Id using g e t unauthorized response has a 4xx status code
func (o *GetEndpointByIDUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint by Id using g e t unauthorized response has a 5xx status code
func (o *GetEndpointByIDUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint by Id using g e t unauthorized response a status code equal to that given
func (o *GetEndpointByIDUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetEndpointByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getEndpointByIdUsingGETUnauthorized ", 401)
}

func (o *GetEndpointByIDUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getEndpointByIdUsingGETUnauthorized ", 401)
}

func (o *GetEndpointByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointByIDUsingGETForbidden creates a GetEndpointByIDUsingGETForbidden with default headers values
func NewGetEndpointByIDUsingGETForbidden() *GetEndpointByIDUsingGETForbidden {
	return &GetEndpointByIDUsingGETForbidden{}
}

/*
GetEndpointByIDUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetEndpointByIDUsingGETForbidden struct {
}

// IsSuccess returns true when this get endpoint by Id using g e t forbidden response has a 2xx status code
func (o *GetEndpointByIDUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint by Id using g e t forbidden response has a 3xx status code
func (o *GetEndpointByIDUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint by Id using g e t forbidden response has a 4xx status code
func (o *GetEndpointByIDUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint by Id using g e t forbidden response has a 5xx status code
func (o *GetEndpointByIDUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint by Id using g e t forbidden response a status code equal to that given
func (o *GetEndpointByIDUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetEndpointByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getEndpointByIdUsingGETForbidden ", 403)
}

func (o *GetEndpointByIDUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getEndpointByIdUsingGETForbidden ", 403)
}

func (o *GetEndpointByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointByIDUsingGETNotFound creates a GetEndpointByIDUsingGETNotFound with default headers values
func NewGetEndpointByIDUsingGETNotFound() *GetEndpointByIDUsingGETNotFound {
	return &GetEndpointByIDUsingGETNotFound{}
}

/*
GetEndpointByIDUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetEndpointByIDUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get endpoint by Id using g e t not found response has a 2xx status code
func (o *GetEndpointByIDUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint by Id using g e t not found response has a 3xx status code
func (o *GetEndpointByIDUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint by Id using g e t not found response has a 4xx status code
func (o *GetEndpointByIDUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint by Id using g e t not found response has a 5xx status code
func (o *GetEndpointByIDUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint by Id using g e t not found response a status code equal to that given
func (o *GetEndpointByIDUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetEndpointByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getEndpointByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetEndpointByIDUsingGETNotFound) String() string {
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

/*
GetEndpointByIDUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetEndpointByIDUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get endpoint by Id using g e t internal server error response has a 2xx status code
func (o *GetEndpointByIDUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint by Id using g e t internal server error response has a 3xx status code
func (o *GetEndpointByIDUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint by Id using g e t internal server error response has a 4xx status code
func (o *GetEndpointByIDUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get endpoint by Id using g e t internal server error response has a 5xx status code
func (o *GetEndpointByIDUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get endpoint by Id using g e t internal server error response a status code equal to that given
func (o *GetEndpointByIDUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetEndpointByIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getEndpointByIdUsingGETInternalServerError ", 500)
}

func (o *GetEndpointByIDUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getEndpointByIdUsingGETInternalServerError ", 500)
}

func (o *GetEndpointByIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
