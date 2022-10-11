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

// GetEndpointByNameUsingGETReader is a Reader for the GetEndpointByNameUsingGET structure.
type GetEndpointByNameUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEndpointByNameUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEndpointByNameUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEndpointByNameUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEndpointByNameUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEndpointByNameUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEndpointByNameUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEndpointByNameUsingGETOK creates a GetEndpointByNameUsingGETOK with default headers values
func NewGetEndpointByNameUsingGETOK() *GetEndpointByNameUsingGETOK {
	return &GetEndpointByNameUsingGETOK{}
}

/*
GetEndpointByNameUsingGETOK describes a response with status code 200, with default header values.

'Success' with the requested Endpoint
*/
type GetEndpointByNameUsingGETOK struct {
	Payload models.Endpoint
}

// IsSuccess returns true when this get endpoint by name using g e t o k response has a 2xx status code
func (o *GetEndpointByNameUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get endpoint by name using g e t o k response has a 3xx status code
func (o *GetEndpointByNameUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint by name using g e t o k response has a 4xx status code
func (o *GetEndpointByNameUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get endpoint by name using g e t o k response has a 5xx status code
func (o *GetEndpointByNameUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint by name using g e t o k response a status code equal to that given
func (o *GetEndpointByNameUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetEndpointByNameUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{project}/{name}][%d] getEndpointByNameUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetEndpointByNameUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{project}/{name}][%d] getEndpointByNameUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetEndpointByNameUsingGETOK) GetPayload() models.Endpoint {
	return o.Payload
}

func (o *GetEndpointByNameUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalEndpoint(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetEndpointByNameUsingGETUnauthorized creates a GetEndpointByNameUsingGETUnauthorized with default headers values
func NewGetEndpointByNameUsingGETUnauthorized() *GetEndpointByNameUsingGETUnauthorized {
	return &GetEndpointByNameUsingGETUnauthorized{}
}

/*
GetEndpointByNameUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetEndpointByNameUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get endpoint by name using g e t unauthorized response has a 2xx status code
func (o *GetEndpointByNameUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint by name using g e t unauthorized response has a 3xx status code
func (o *GetEndpointByNameUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint by name using g e t unauthorized response has a 4xx status code
func (o *GetEndpointByNameUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint by name using g e t unauthorized response has a 5xx status code
func (o *GetEndpointByNameUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint by name using g e t unauthorized response a status code equal to that given
func (o *GetEndpointByNameUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetEndpointByNameUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{project}/{name}][%d] getEndpointByNameUsingGETUnauthorized ", 401)
}

func (o *GetEndpointByNameUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{project}/{name}][%d] getEndpointByNameUsingGETUnauthorized ", 401)
}

func (o *GetEndpointByNameUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointByNameUsingGETForbidden creates a GetEndpointByNameUsingGETForbidden with default headers values
func NewGetEndpointByNameUsingGETForbidden() *GetEndpointByNameUsingGETForbidden {
	return &GetEndpointByNameUsingGETForbidden{}
}

/*
GetEndpointByNameUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetEndpointByNameUsingGETForbidden struct {
}

// IsSuccess returns true when this get endpoint by name using g e t forbidden response has a 2xx status code
func (o *GetEndpointByNameUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint by name using g e t forbidden response has a 3xx status code
func (o *GetEndpointByNameUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint by name using g e t forbidden response has a 4xx status code
func (o *GetEndpointByNameUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint by name using g e t forbidden response has a 5xx status code
func (o *GetEndpointByNameUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint by name using g e t forbidden response a status code equal to that given
func (o *GetEndpointByNameUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetEndpointByNameUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{project}/{name}][%d] getEndpointByNameUsingGETForbidden ", 403)
}

func (o *GetEndpointByNameUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{project}/{name}][%d] getEndpointByNameUsingGETForbidden ", 403)
}

func (o *GetEndpointByNameUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointByNameUsingGETNotFound creates a GetEndpointByNameUsingGETNotFound with default headers values
func NewGetEndpointByNameUsingGETNotFound() *GetEndpointByNameUsingGETNotFound {
	return &GetEndpointByNameUsingGETNotFound{}
}

/*
GetEndpointByNameUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetEndpointByNameUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get endpoint by name using g e t not found response has a 2xx status code
func (o *GetEndpointByNameUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint by name using g e t not found response has a 3xx status code
func (o *GetEndpointByNameUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint by name using g e t not found response has a 4xx status code
func (o *GetEndpointByNameUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint by name using g e t not found response has a 5xx status code
func (o *GetEndpointByNameUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint by name using g e t not found response a status code equal to that given
func (o *GetEndpointByNameUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetEndpointByNameUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{project}/{name}][%d] getEndpointByNameUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetEndpointByNameUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{project}/{name}][%d] getEndpointByNameUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetEndpointByNameUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEndpointByNameUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointByNameUsingGETInternalServerError creates a GetEndpointByNameUsingGETInternalServerError with default headers values
func NewGetEndpointByNameUsingGETInternalServerError() *GetEndpointByNameUsingGETInternalServerError {
	return &GetEndpointByNameUsingGETInternalServerError{}
}

/*
GetEndpointByNameUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetEndpointByNameUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get endpoint by name using g e t internal server error response has a 2xx status code
func (o *GetEndpointByNameUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint by name using g e t internal server error response has a 3xx status code
func (o *GetEndpointByNameUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint by name using g e t internal server error response has a 4xx status code
func (o *GetEndpointByNameUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get endpoint by name using g e t internal server error response has a 5xx status code
func (o *GetEndpointByNameUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get endpoint by name using g e t internal server error response a status code equal to that given
func (o *GetEndpointByNameUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetEndpointByNameUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{project}/{name}][%d] getEndpointByNameUsingGETInternalServerError ", 500)
}

func (o *GetEndpointByNameUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{project}/{name}][%d] getEndpointByNameUsingGETInternalServerError ", 500)
}

func (o *GetEndpointByNameUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}