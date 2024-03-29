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

// GetAllVariablesUsingGETReader is a Reader for the GetAllVariablesUsingGET structure.
type GetAllVariablesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllVariablesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllVariablesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllVariablesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllVariablesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllVariablesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllVariablesUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllVariablesUsingGETOK creates a GetAllVariablesUsingGETOK with default headers values
func NewGetAllVariablesUsingGETOK() *GetAllVariablesUsingGETOK {
	return &GetAllVariablesUsingGETOK{}
}

/*
GetAllVariablesUsingGETOK describes a response with status code 200, with default header values.

'Success' with Variables on pages
*/
type GetAllVariablesUsingGETOK struct {
	Payload models.Variables
}

// IsSuccess returns true when this get all variables using g e t o k response has a 2xx status code
func (o *GetAllVariablesUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all variables using g e t o k response has a 3xx status code
func (o *GetAllVariablesUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all variables using g e t o k response has a 4xx status code
func (o *GetAllVariablesUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all variables using g e t o k response has a 5xx status code
func (o *GetAllVariablesUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all variables using g e t o k response a status code equal to that given
func (o *GetAllVariablesUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllVariablesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables][%d] getAllVariablesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllVariablesUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/variables][%d] getAllVariablesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllVariablesUsingGETOK) GetPayload() models.Variables {
	return o.Payload
}

func (o *GetAllVariablesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalVariables(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetAllVariablesUsingGETUnauthorized creates a GetAllVariablesUsingGETUnauthorized with default headers values
func NewGetAllVariablesUsingGETUnauthorized() *GetAllVariablesUsingGETUnauthorized {
	return &GetAllVariablesUsingGETUnauthorized{}
}

/*
GetAllVariablesUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetAllVariablesUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get all variables using g e t unauthorized response has a 2xx status code
func (o *GetAllVariablesUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all variables using g e t unauthorized response has a 3xx status code
func (o *GetAllVariablesUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all variables using g e t unauthorized response has a 4xx status code
func (o *GetAllVariablesUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all variables using g e t unauthorized response has a 5xx status code
func (o *GetAllVariablesUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all variables using g e t unauthorized response a status code equal to that given
func (o *GetAllVariablesUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAllVariablesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables][%d] getAllVariablesUsingGETUnauthorized ", 401)
}

func (o *GetAllVariablesUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/variables][%d] getAllVariablesUsingGETUnauthorized ", 401)
}

func (o *GetAllVariablesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllVariablesUsingGETForbidden creates a GetAllVariablesUsingGETForbidden with default headers values
func NewGetAllVariablesUsingGETForbidden() *GetAllVariablesUsingGETForbidden {
	return &GetAllVariablesUsingGETForbidden{}
}

/*
GetAllVariablesUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllVariablesUsingGETForbidden struct {
}

// IsSuccess returns true when this get all variables using g e t forbidden response has a 2xx status code
func (o *GetAllVariablesUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all variables using g e t forbidden response has a 3xx status code
func (o *GetAllVariablesUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all variables using g e t forbidden response has a 4xx status code
func (o *GetAllVariablesUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all variables using g e t forbidden response has a 5xx status code
func (o *GetAllVariablesUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get all variables using g e t forbidden response a status code equal to that given
func (o *GetAllVariablesUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAllVariablesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables][%d] getAllVariablesUsingGETForbidden ", 403)
}

func (o *GetAllVariablesUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/variables][%d] getAllVariablesUsingGETForbidden ", 403)
}

func (o *GetAllVariablesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllVariablesUsingGETNotFound creates a GetAllVariablesUsingGETNotFound with default headers values
func NewGetAllVariablesUsingGETNotFound() *GetAllVariablesUsingGETNotFound {
	return &GetAllVariablesUsingGETNotFound{}
}

/*
GetAllVariablesUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllVariablesUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get all variables using g e t not found response has a 2xx status code
func (o *GetAllVariablesUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all variables using g e t not found response has a 3xx status code
func (o *GetAllVariablesUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all variables using g e t not found response has a 4xx status code
func (o *GetAllVariablesUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all variables using g e t not found response has a 5xx status code
func (o *GetAllVariablesUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all variables using g e t not found response a status code equal to that given
func (o *GetAllVariablesUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAllVariablesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables][%d] getAllVariablesUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetAllVariablesUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/variables][%d] getAllVariablesUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetAllVariablesUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllVariablesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllVariablesUsingGETInternalServerError creates a GetAllVariablesUsingGETInternalServerError with default headers values
func NewGetAllVariablesUsingGETInternalServerError() *GetAllVariablesUsingGETInternalServerError {
	return &GetAllVariablesUsingGETInternalServerError{}
}

/*
GetAllVariablesUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetAllVariablesUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get all variables using g e t internal server error response has a 2xx status code
func (o *GetAllVariablesUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all variables using g e t internal server error response has a 3xx status code
func (o *GetAllVariablesUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all variables using g e t internal server error response has a 4xx status code
func (o *GetAllVariablesUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all variables using g e t internal server error response has a 5xx status code
func (o *GetAllVariablesUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all variables using g e t internal server error response a status code equal to that given
func (o *GetAllVariablesUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAllVariablesUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables][%d] getAllVariablesUsingGETInternalServerError ", 500)
}

func (o *GetAllVariablesUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/variables][%d] getAllVariablesUsingGETInternalServerError ", 500)
}

func (o *GetAllVariablesUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
