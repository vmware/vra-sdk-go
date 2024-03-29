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

// GetGerritListenerByIDUsingGETReader is a Reader for the GetGerritListenerByIDUsingGET structure.
type GetGerritListenerByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGerritListenerByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGerritListenerByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGerritListenerByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGerritListenerByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGerritListenerByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGerritListenerByIDUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGerritListenerByIDUsingGETOK creates a GetGerritListenerByIDUsingGETOK with default headers values
func NewGetGerritListenerByIDUsingGETOK() *GetGerritListenerByIDUsingGETOK {
	return &GetGerritListenerByIDUsingGETOK{}
}

/*
GetGerritListenerByIDUsingGETOK describes a response with status code 200, with default header values.

'Success' with gerrit listener
*/
type GetGerritListenerByIDUsingGETOK struct {
	Payload models.GerritListener
}

// IsSuccess returns true when this get gerrit listener by Id using g e t o k response has a 2xx status code
func (o *GetGerritListenerByIDUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gerrit listener by Id using g e t o k response has a 3xx status code
func (o *GetGerritListenerByIDUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit listener by Id using g e t o k response has a 4xx status code
func (o *GetGerritListenerByIDUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gerrit listener by Id using g e t o k response has a 5xx status code
func (o *GetGerritListenerByIDUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit listener by Id using g e t o k response a status code equal to that given
func (o *GetGerritListenerByIDUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGerritListenerByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}][%d] getGerritListenerByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGerritListenerByIDUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}][%d] getGerritListenerByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGerritListenerByIDUsingGETOK) GetPayload() models.GerritListener {
	return o.Payload
}

func (o *GetGerritListenerByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListener(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetGerritListenerByIDUsingGETUnauthorized creates a GetGerritListenerByIDUsingGETUnauthorized with default headers values
func NewGetGerritListenerByIDUsingGETUnauthorized() *GetGerritListenerByIDUsingGETUnauthorized {
	return &GetGerritListenerByIDUsingGETUnauthorized{}
}

/*
GetGerritListenerByIDUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetGerritListenerByIDUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get gerrit listener by Id using g e t unauthorized response has a 2xx status code
func (o *GetGerritListenerByIDUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit listener by Id using g e t unauthorized response has a 3xx status code
func (o *GetGerritListenerByIDUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit listener by Id using g e t unauthorized response has a 4xx status code
func (o *GetGerritListenerByIDUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gerrit listener by Id using g e t unauthorized response has a 5xx status code
func (o *GetGerritListenerByIDUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit listener by Id using g e t unauthorized response a status code equal to that given
func (o *GetGerritListenerByIDUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetGerritListenerByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}][%d] getGerritListenerByIdUsingGETUnauthorized ", 401)
}

func (o *GetGerritListenerByIDUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}][%d] getGerritListenerByIdUsingGETUnauthorized ", 401)
}

func (o *GetGerritListenerByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGerritListenerByIDUsingGETForbidden creates a GetGerritListenerByIDUsingGETForbidden with default headers values
func NewGetGerritListenerByIDUsingGETForbidden() *GetGerritListenerByIDUsingGETForbidden {
	return &GetGerritListenerByIDUsingGETForbidden{}
}

/*
GetGerritListenerByIDUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetGerritListenerByIDUsingGETForbidden struct {
}

// IsSuccess returns true when this get gerrit listener by Id using g e t forbidden response has a 2xx status code
func (o *GetGerritListenerByIDUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit listener by Id using g e t forbidden response has a 3xx status code
func (o *GetGerritListenerByIDUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit listener by Id using g e t forbidden response has a 4xx status code
func (o *GetGerritListenerByIDUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gerrit listener by Id using g e t forbidden response has a 5xx status code
func (o *GetGerritListenerByIDUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit listener by Id using g e t forbidden response a status code equal to that given
func (o *GetGerritListenerByIDUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetGerritListenerByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}][%d] getGerritListenerByIdUsingGETForbidden ", 403)
}

func (o *GetGerritListenerByIDUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}][%d] getGerritListenerByIdUsingGETForbidden ", 403)
}

func (o *GetGerritListenerByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGerritListenerByIDUsingGETNotFound creates a GetGerritListenerByIDUsingGETNotFound with default headers values
func NewGetGerritListenerByIDUsingGETNotFound() *GetGerritListenerByIDUsingGETNotFound {
	return &GetGerritListenerByIDUsingGETNotFound{}
}

/*
GetGerritListenerByIDUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetGerritListenerByIDUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get gerrit listener by Id using g e t not found response has a 2xx status code
func (o *GetGerritListenerByIDUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit listener by Id using g e t not found response has a 3xx status code
func (o *GetGerritListenerByIDUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit listener by Id using g e t not found response has a 4xx status code
func (o *GetGerritListenerByIDUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gerrit listener by Id using g e t not found response has a 5xx status code
func (o *GetGerritListenerByIDUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit listener by Id using g e t not found response a status code equal to that given
func (o *GetGerritListenerByIDUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetGerritListenerByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}][%d] getGerritListenerByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGerritListenerByIDUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}][%d] getGerritListenerByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGerritListenerByIDUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGerritListenerByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGerritListenerByIDUsingGETInternalServerError creates a GetGerritListenerByIDUsingGETInternalServerError with default headers values
func NewGetGerritListenerByIDUsingGETInternalServerError() *GetGerritListenerByIDUsingGETInternalServerError {
	return &GetGerritListenerByIDUsingGETInternalServerError{}
}

/*
GetGerritListenerByIDUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetGerritListenerByIDUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get gerrit listener by Id using g e t internal server error response has a 2xx status code
func (o *GetGerritListenerByIDUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit listener by Id using g e t internal server error response has a 3xx status code
func (o *GetGerritListenerByIDUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit listener by Id using g e t internal server error response has a 4xx status code
func (o *GetGerritListenerByIDUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gerrit listener by Id using g e t internal server error response has a 5xx status code
func (o *GetGerritListenerByIDUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get gerrit listener by Id using g e t internal server error response a status code equal to that given
func (o *GetGerritListenerByIDUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetGerritListenerByIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}][%d] getGerritListenerByIdUsingGETInternalServerError ", 500)
}

func (o *GetGerritListenerByIDUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners/{id}][%d] getGerritListenerByIdUsingGETInternalServerError ", 500)
}

func (o *GetGerritListenerByIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
