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

// GetUsingGETMixin4Reader is a Reader for the GetUsingGETMixin4 structure.
type GetUsingGETMixin4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsingGETMixin4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsingGETMixin4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsingGETMixin4Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsingGETMixin4Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsingGETMixin4NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsingGETMixin4InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsingGETMixin4OK creates a GetUsingGETMixin4OK with default headers values
func NewGetUsingGETMixin4OK() *GetUsingGETMixin4OK {
	return &GetUsingGETMixin4OK{}
}

/*
GetUsingGETMixin4OK describes a response with status code 200, with default header values.

'Success' with endpoint certificate details
*/
type GetUsingGETMixin4OK struct {
	Payload models.EndpointCertificateChain
}

// IsSuccess returns true when this get using g e t mixin4 o k response has a 2xx status code
func (o *GetUsingGETMixin4OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get using g e t mixin4 o k response has a 3xx status code
func (o *GetUsingGETMixin4OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get using g e t mixin4 o k response has a 4xx status code
func (o *GetUsingGETMixin4OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get using g e t mixin4 o k response has a 5xx status code
func (o *GetUsingGETMixin4OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get using g e t mixin4 o k response a status code equal to that given
func (o *GetUsingGETMixin4OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUsingGETMixin4OK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-certificate][%d] getUsingGETMixin4OK  %+v", 200, o.Payload)
}

func (o *GetUsingGETMixin4OK) String() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-certificate][%d] getUsingGETMixin4OK  %+v", 200, o.Payload)
}

func (o *GetUsingGETMixin4OK) GetPayload() models.EndpointCertificateChain {
	return o.Payload
}

func (o *GetUsingGETMixin4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalEndpointCertificateChain(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetUsingGETMixin4Unauthorized creates a GetUsingGETMixin4Unauthorized with default headers values
func NewGetUsingGETMixin4Unauthorized() *GetUsingGETMixin4Unauthorized {
	return &GetUsingGETMixin4Unauthorized{}
}

/*
GetUsingGETMixin4Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetUsingGETMixin4Unauthorized struct {
}

// IsSuccess returns true when this get using g e t mixin4 unauthorized response has a 2xx status code
func (o *GetUsingGETMixin4Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get using g e t mixin4 unauthorized response has a 3xx status code
func (o *GetUsingGETMixin4Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get using g e t mixin4 unauthorized response has a 4xx status code
func (o *GetUsingGETMixin4Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get using g e t mixin4 unauthorized response has a 5xx status code
func (o *GetUsingGETMixin4Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get using g e t mixin4 unauthorized response a status code equal to that given
func (o *GetUsingGETMixin4Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUsingGETMixin4Unauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-certificate][%d] getUsingGETMixin4Unauthorized ", 401)
}

func (o *GetUsingGETMixin4Unauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-certificate][%d] getUsingGETMixin4Unauthorized ", 401)
}

func (o *GetUsingGETMixin4Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGETMixin4Forbidden creates a GetUsingGETMixin4Forbidden with default headers values
func NewGetUsingGETMixin4Forbidden() *GetUsingGETMixin4Forbidden {
	return &GetUsingGETMixin4Forbidden{}
}

/*
GetUsingGETMixin4Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUsingGETMixin4Forbidden struct {
}

// IsSuccess returns true when this get using g e t mixin4 forbidden response has a 2xx status code
func (o *GetUsingGETMixin4Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get using g e t mixin4 forbidden response has a 3xx status code
func (o *GetUsingGETMixin4Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get using g e t mixin4 forbidden response has a 4xx status code
func (o *GetUsingGETMixin4Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get using g e t mixin4 forbidden response has a 5xx status code
func (o *GetUsingGETMixin4Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get using g e t mixin4 forbidden response a status code equal to that given
func (o *GetUsingGETMixin4Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUsingGETMixin4Forbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-certificate][%d] getUsingGETMixin4Forbidden ", 403)
}

func (o *GetUsingGETMixin4Forbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-certificate][%d] getUsingGETMixin4Forbidden ", 403)
}

func (o *GetUsingGETMixin4Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGETMixin4NotFound creates a GetUsingGETMixin4NotFound with default headers values
func NewGetUsingGETMixin4NotFound() *GetUsingGETMixin4NotFound {
	return &GetUsingGETMixin4NotFound{}
}

/*
GetUsingGETMixin4NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetUsingGETMixin4NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get using g e t mixin4 not found response has a 2xx status code
func (o *GetUsingGETMixin4NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get using g e t mixin4 not found response has a 3xx status code
func (o *GetUsingGETMixin4NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get using g e t mixin4 not found response has a 4xx status code
func (o *GetUsingGETMixin4NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get using g e t mixin4 not found response has a 5xx status code
func (o *GetUsingGETMixin4NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get using g e t mixin4 not found response a status code equal to that given
func (o *GetUsingGETMixin4NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUsingGETMixin4NotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-certificate][%d] getUsingGETMixin4NotFound  %+v", 404, o.Payload)
}

func (o *GetUsingGETMixin4NotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-certificate][%d] getUsingGETMixin4NotFound  %+v", 404, o.Payload)
}

func (o *GetUsingGETMixin4NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUsingGETMixin4NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsingGETMixin4InternalServerError creates a GetUsingGETMixin4InternalServerError with default headers values
func NewGetUsingGETMixin4InternalServerError() *GetUsingGETMixin4InternalServerError {
	return &GetUsingGETMixin4InternalServerError{}
}

/*
GetUsingGETMixin4InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetUsingGETMixin4InternalServerError struct {
}

// IsSuccess returns true when this get using g e t mixin4 internal server error response has a 2xx status code
func (o *GetUsingGETMixin4InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get using g e t mixin4 internal server error response has a 3xx status code
func (o *GetUsingGETMixin4InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get using g e t mixin4 internal server error response has a 4xx status code
func (o *GetUsingGETMixin4InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get using g e t mixin4 internal server error response has a 5xx status code
func (o *GetUsingGETMixin4InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get using g e t mixin4 internal server error response a status code equal to that given
func (o *GetUsingGETMixin4InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUsingGETMixin4InternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-certificate][%d] getUsingGETMixin4InternalServerError ", 500)
}

func (o *GetUsingGETMixin4InternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/endpoint-certificate][%d] getUsingGETMixin4InternalServerError ", 500)
}

func (o *GetUsingGETMixin4InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
