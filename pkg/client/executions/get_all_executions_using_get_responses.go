// Code generated by go-swagger; DO NOT EDIT.

package executions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetAllExecutionsUsingGETReader is a Reader for the GetAllExecutionsUsingGET structure.
type GetAllExecutionsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllExecutionsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllExecutionsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllExecutionsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllExecutionsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllExecutionsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllExecutionsUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllExecutionsUsingGETOK creates a GetAllExecutionsUsingGETOK with default headers values
func NewGetAllExecutionsUsingGETOK() *GetAllExecutionsUsingGETOK {
	return &GetAllExecutionsUsingGETOK{}
}

/*
GetAllExecutionsUsingGETOK describes a response with status code 200, with default header values.

'Success' with the requested Executions
*/
type GetAllExecutionsUsingGETOK struct {
	Payload models.Executions
}

// IsSuccess returns true when this get all executions using g e t o k response has a 2xx status code
func (o *GetAllExecutionsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all executions using g e t o k response has a 3xx status code
func (o *GetAllExecutionsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all executions using g e t o k response has a 4xx status code
func (o *GetAllExecutionsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all executions using g e t o k response has a 5xx status code
func (o *GetAllExecutionsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all executions using g e t o k response a status code equal to that given
func (o *GetAllExecutionsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllExecutionsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllExecutionsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllExecutionsUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllExecutionsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllExecutionsUsingGETOK) GetPayload() models.Executions {
	return o.Payload
}

func (o *GetAllExecutionsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecutions(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetAllExecutionsUsingGETUnauthorized creates a GetAllExecutionsUsingGETUnauthorized with default headers values
func NewGetAllExecutionsUsingGETUnauthorized() *GetAllExecutionsUsingGETUnauthorized {
	return &GetAllExecutionsUsingGETUnauthorized{}
}

/*
GetAllExecutionsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetAllExecutionsUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get all executions using g e t unauthorized response has a 2xx status code
func (o *GetAllExecutionsUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all executions using g e t unauthorized response has a 3xx status code
func (o *GetAllExecutionsUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all executions using g e t unauthorized response has a 4xx status code
func (o *GetAllExecutionsUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all executions using g e t unauthorized response has a 5xx status code
func (o *GetAllExecutionsUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all executions using g e t unauthorized response a status code equal to that given
func (o *GetAllExecutionsUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAllExecutionsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllExecutionsUsingGETUnauthorized ", 401)
}

func (o *GetAllExecutionsUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllExecutionsUsingGETUnauthorized ", 401)
}

func (o *GetAllExecutionsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllExecutionsUsingGETForbidden creates a GetAllExecutionsUsingGETForbidden with default headers values
func NewGetAllExecutionsUsingGETForbidden() *GetAllExecutionsUsingGETForbidden {
	return &GetAllExecutionsUsingGETForbidden{}
}

/*
GetAllExecutionsUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllExecutionsUsingGETForbidden struct {
}

// IsSuccess returns true when this get all executions using g e t forbidden response has a 2xx status code
func (o *GetAllExecutionsUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all executions using g e t forbidden response has a 3xx status code
func (o *GetAllExecutionsUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all executions using g e t forbidden response has a 4xx status code
func (o *GetAllExecutionsUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all executions using g e t forbidden response has a 5xx status code
func (o *GetAllExecutionsUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get all executions using g e t forbidden response a status code equal to that given
func (o *GetAllExecutionsUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAllExecutionsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllExecutionsUsingGETForbidden ", 403)
}

func (o *GetAllExecutionsUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllExecutionsUsingGETForbidden ", 403)
}

func (o *GetAllExecutionsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllExecutionsUsingGETNotFound creates a GetAllExecutionsUsingGETNotFound with default headers values
func NewGetAllExecutionsUsingGETNotFound() *GetAllExecutionsUsingGETNotFound {
	return &GetAllExecutionsUsingGETNotFound{}
}

/*
GetAllExecutionsUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllExecutionsUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get all executions using g e t not found response has a 2xx status code
func (o *GetAllExecutionsUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all executions using g e t not found response has a 3xx status code
func (o *GetAllExecutionsUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all executions using g e t not found response has a 4xx status code
func (o *GetAllExecutionsUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all executions using g e t not found response has a 5xx status code
func (o *GetAllExecutionsUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all executions using g e t not found response a status code equal to that given
func (o *GetAllExecutionsUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAllExecutionsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllExecutionsUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetAllExecutionsUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllExecutionsUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetAllExecutionsUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllExecutionsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllExecutionsUsingGETInternalServerError creates a GetAllExecutionsUsingGETInternalServerError with default headers values
func NewGetAllExecutionsUsingGETInternalServerError() *GetAllExecutionsUsingGETInternalServerError {
	return &GetAllExecutionsUsingGETInternalServerError{}
}

/*
GetAllExecutionsUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetAllExecutionsUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get all executions using g e t internal server error response has a 2xx status code
func (o *GetAllExecutionsUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all executions using g e t internal server error response has a 3xx status code
func (o *GetAllExecutionsUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all executions using g e t internal server error response has a 4xx status code
func (o *GetAllExecutionsUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all executions using g e t internal server error response has a 5xx status code
func (o *GetAllExecutionsUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all executions using g e t internal server error response a status code equal to that given
func (o *GetAllExecutionsUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAllExecutionsUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllExecutionsUsingGETInternalServerError ", 500)
}

func (o *GetAllExecutionsUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllExecutionsUsingGETInternalServerError ", 500)
}

func (o *GetAllExecutionsUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}