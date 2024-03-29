// Code generated by go-swagger; DO NOT EDIT.

package user_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetAllUserOperationsUsingGETReader is a Reader for the GetAllUserOperationsUsingGET structure.
type GetAllUserOperationsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllUserOperationsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllUserOperationsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllUserOperationsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllUserOperationsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllUserOperationsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllUserOperationsUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllUserOperationsUsingGETOK creates a GetAllUserOperationsUsingGETOK with default headers values
func NewGetAllUserOperationsUsingGETOK() *GetAllUserOperationsUsingGETOK {
	return &GetAllUserOperationsUsingGETOK{}
}

/*
GetAllUserOperationsUsingGETOK describes a response with status code 200, with default header values.

'Success' with the requested User Operations
*/
type GetAllUserOperationsUsingGETOK struct {
	Payload models.UserOperations
}

// IsSuccess returns true when this get all user operations using g e t o k response has a 2xx status code
func (o *GetAllUserOperationsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all user operations using g e t o k response has a 3xx status code
func (o *GetAllUserOperationsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all user operations using g e t o k response has a 4xx status code
func (o *GetAllUserOperationsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all user operations using g e t o k response has a 5xx status code
func (o *GetAllUserOperationsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all user operations using g e t o k response a status code equal to that given
func (o *GetAllUserOperationsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllUserOperationsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations][%d] getAllUserOperationsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllUserOperationsUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations][%d] getAllUserOperationsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllUserOperationsUsingGETOK) GetPayload() models.UserOperations {
	return o.Payload
}

func (o *GetAllUserOperationsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalUserOperations(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetAllUserOperationsUsingGETUnauthorized creates a GetAllUserOperationsUsingGETUnauthorized with default headers values
func NewGetAllUserOperationsUsingGETUnauthorized() *GetAllUserOperationsUsingGETUnauthorized {
	return &GetAllUserOperationsUsingGETUnauthorized{}
}

/*
GetAllUserOperationsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetAllUserOperationsUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get all user operations using g e t unauthorized response has a 2xx status code
func (o *GetAllUserOperationsUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all user operations using g e t unauthorized response has a 3xx status code
func (o *GetAllUserOperationsUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all user operations using g e t unauthorized response has a 4xx status code
func (o *GetAllUserOperationsUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all user operations using g e t unauthorized response has a 5xx status code
func (o *GetAllUserOperationsUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all user operations using g e t unauthorized response a status code equal to that given
func (o *GetAllUserOperationsUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAllUserOperationsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations][%d] getAllUserOperationsUsingGETUnauthorized ", 401)
}

func (o *GetAllUserOperationsUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations][%d] getAllUserOperationsUsingGETUnauthorized ", 401)
}

func (o *GetAllUserOperationsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUserOperationsUsingGETForbidden creates a GetAllUserOperationsUsingGETForbidden with default headers values
func NewGetAllUserOperationsUsingGETForbidden() *GetAllUserOperationsUsingGETForbidden {
	return &GetAllUserOperationsUsingGETForbidden{}
}

/*
GetAllUserOperationsUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllUserOperationsUsingGETForbidden struct {
}

// IsSuccess returns true when this get all user operations using g e t forbidden response has a 2xx status code
func (o *GetAllUserOperationsUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all user operations using g e t forbidden response has a 3xx status code
func (o *GetAllUserOperationsUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all user operations using g e t forbidden response has a 4xx status code
func (o *GetAllUserOperationsUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all user operations using g e t forbidden response has a 5xx status code
func (o *GetAllUserOperationsUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get all user operations using g e t forbidden response a status code equal to that given
func (o *GetAllUserOperationsUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAllUserOperationsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations][%d] getAllUserOperationsUsingGETForbidden ", 403)
}

func (o *GetAllUserOperationsUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations][%d] getAllUserOperationsUsingGETForbidden ", 403)
}

func (o *GetAllUserOperationsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUserOperationsUsingGETNotFound creates a GetAllUserOperationsUsingGETNotFound with default headers values
func NewGetAllUserOperationsUsingGETNotFound() *GetAllUserOperationsUsingGETNotFound {
	return &GetAllUserOperationsUsingGETNotFound{}
}

/*
GetAllUserOperationsUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllUserOperationsUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get all user operations using g e t not found response has a 2xx status code
func (o *GetAllUserOperationsUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all user operations using g e t not found response has a 3xx status code
func (o *GetAllUserOperationsUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all user operations using g e t not found response has a 4xx status code
func (o *GetAllUserOperationsUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all user operations using g e t not found response has a 5xx status code
func (o *GetAllUserOperationsUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all user operations using g e t not found response a status code equal to that given
func (o *GetAllUserOperationsUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAllUserOperationsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations][%d] getAllUserOperationsUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetAllUserOperationsUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations][%d] getAllUserOperationsUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetAllUserOperationsUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllUserOperationsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllUserOperationsUsingGETInternalServerError creates a GetAllUserOperationsUsingGETInternalServerError with default headers values
func NewGetAllUserOperationsUsingGETInternalServerError() *GetAllUserOperationsUsingGETInternalServerError {
	return &GetAllUserOperationsUsingGETInternalServerError{}
}

/*
GetAllUserOperationsUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetAllUserOperationsUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get all user operations using g e t internal server error response has a 2xx status code
func (o *GetAllUserOperationsUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all user operations using g e t internal server error response has a 3xx status code
func (o *GetAllUserOperationsUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all user operations using g e t internal server error response has a 4xx status code
func (o *GetAllUserOperationsUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all user operations using g e t internal server error response has a 5xx status code
func (o *GetAllUserOperationsUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all user operations using g e t internal server error response a status code equal to that given
func (o *GetAllUserOperationsUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAllUserOperationsUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations][%d] getAllUserOperationsUsingGETInternalServerError ", 500)
}

func (o *GetAllUserOperationsUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations][%d] getAllUserOperationsUsingGETInternalServerError ", 500)
}

func (o *GetAllUserOperationsUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
