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

// GetGerritProjectsByIDUsingGETReader is a Reader for the GetGerritProjectsByIDUsingGET structure.
type GetGerritProjectsByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGerritProjectsByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGerritProjectsByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGerritProjectsByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGerritProjectsByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGerritProjectsByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGerritProjectsByIDUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGerritProjectsByIDUsingGETOK creates a GetGerritProjectsByIDUsingGETOK with default headers values
func NewGetGerritProjectsByIDUsingGETOK() *GetGerritProjectsByIDUsingGETOK {
	return &GetGerritProjectsByIDUsingGETOK{}
}

/*
GetGerritProjectsByIDUsingGETOK describes a response with status code 200, with default header values.

'Success' with gerrit projects
*/
type GetGerritProjectsByIDUsingGETOK struct {
	Payload string
}

// IsSuccess returns true when this get gerrit projects by Id using g e t o k response has a 2xx status code
func (o *GetGerritProjectsByIDUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gerrit projects by Id using g e t o k response has a 3xx status code
func (o *GetGerritProjectsByIDUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit projects by Id using g e t o k response has a 4xx status code
func (o *GetGerritProjectsByIDUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gerrit projects by Id using g e t o k response has a 5xx status code
func (o *GetGerritProjectsByIDUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit projects by Id using g e t o k response a status code equal to that given
func (o *GetGerritProjectsByIDUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGerritProjectsByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}/projects][%d] getGerritProjectsByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGerritProjectsByIDUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}/projects][%d] getGerritProjectsByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGerritProjectsByIDUsingGETOK) GetPayload() string {
	return o.Payload
}

func (o *GetGerritProjectsByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGerritProjectsByIDUsingGETUnauthorized creates a GetGerritProjectsByIDUsingGETUnauthorized with default headers values
func NewGetGerritProjectsByIDUsingGETUnauthorized() *GetGerritProjectsByIDUsingGETUnauthorized {
	return &GetGerritProjectsByIDUsingGETUnauthorized{}
}

/*
GetGerritProjectsByIDUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetGerritProjectsByIDUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get gerrit projects by Id using g e t unauthorized response has a 2xx status code
func (o *GetGerritProjectsByIDUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit projects by Id using g e t unauthorized response has a 3xx status code
func (o *GetGerritProjectsByIDUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit projects by Id using g e t unauthorized response has a 4xx status code
func (o *GetGerritProjectsByIDUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gerrit projects by Id using g e t unauthorized response has a 5xx status code
func (o *GetGerritProjectsByIDUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit projects by Id using g e t unauthorized response a status code equal to that given
func (o *GetGerritProjectsByIDUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetGerritProjectsByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}/projects][%d] getGerritProjectsByIdUsingGETUnauthorized ", 401)
}

func (o *GetGerritProjectsByIDUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}/projects][%d] getGerritProjectsByIdUsingGETUnauthorized ", 401)
}

func (o *GetGerritProjectsByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGerritProjectsByIDUsingGETForbidden creates a GetGerritProjectsByIDUsingGETForbidden with default headers values
func NewGetGerritProjectsByIDUsingGETForbidden() *GetGerritProjectsByIDUsingGETForbidden {
	return &GetGerritProjectsByIDUsingGETForbidden{}
}

/*
GetGerritProjectsByIDUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetGerritProjectsByIDUsingGETForbidden struct {
}

// IsSuccess returns true when this get gerrit projects by Id using g e t forbidden response has a 2xx status code
func (o *GetGerritProjectsByIDUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit projects by Id using g e t forbidden response has a 3xx status code
func (o *GetGerritProjectsByIDUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit projects by Id using g e t forbidden response has a 4xx status code
func (o *GetGerritProjectsByIDUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gerrit projects by Id using g e t forbidden response has a 5xx status code
func (o *GetGerritProjectsByIDUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit projects by Id using g e t forbidden response a status code equal to that given
func (o *GetGerritProjectsByIDUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetGerritProjectsByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}/projects][%d] getGerritProjectsByIdUsingGETForbidden ", 403)
}

func (o *GetGerritProjectsByIDUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}/projects][%d] getGerritProjectsByIdUsingGETForbidden ", 403)
}

func (o *GetGerritProjectsByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGerritProjectsByIDUsingGETNotFound creates a GetGerritProjectsByIDUsingGETNotFound with default headers values
func NewGetGerritProjectsByIDUsingGETNotFound() *GetGerritProjectsByIDUsingGETNotFound {
	return &GetGerritProjectsByIDUsingGETNotFound{}
}

/*
GetGerritProjectsByIDUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetGerritProjectsByIDUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get gerrit projects by Id using g e t not found response has a 2xx status code
func (o *GetGerritProjectsByIDUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit projects by Id using g e t not found response has a 3xx status code
func (o *GetGerritProjectsByIDUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit projects by Id using g e t not found response has a 4xx status code
func (o *GetGerritProjectsByIDUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gerrit projects by Id using g e t not found response has a 5xx status code
func (o *GetGerritProjectsByIDUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit projects by Id using g e t not found response a status code equal to that given
func (o *GetGerritProjectsByIDUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetGerritProjectsByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}/projects][%d] getGerritProjectsByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGerritProjectsByIDUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}/projects][%d] getGerritProjectsByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGerritProjectsByIDUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGerritProjectsByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGerritProjectsByIDUsingGETInternalServerError creates a GetGerritProjectsByIDUsingGETInternalServerError with default headers values
func NewGetGerritProjectsByIDUsingGETInternalServerError() *GetGerritProjectsByIDUsingGETInternalServerError {
	return &GetGerritProjectsByIDUsingGETInternalServerError{}
}

/*
GetGerritProjectsByIDUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetGerritProjectsByIDUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get gerrit projects by Id using g e t internal server error response has a 2xx status code
func (o *GetGerritProjectsByIDUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit projects by Id using g e t internal server error response has a 3xx status code
func (o *GetGerritProjectsByIDUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit projects by Id using g e t internal server error response has a 4xx status code
func (o *GetGerritProjectsByIDUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gerrit projects by Id using g e t internal server error response has a 5xx status code
func (o *GetGerritProjectsByIDUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get gerrit projects by Id using g e t internal server error response a status code equal to that given
func (o *GetGerritProjectsByIDUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetGerritProjectsByIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}/projects][%d] getGerritProjectsByIdUsingGETInternalServerError ", 500)
}

func (o *GetGerritProjectsByIDUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}/projects][%d] getGerritProjectsByIdUsingGETInternalServerError ", 500)
}

func (o *GetGerritProjectsByIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}