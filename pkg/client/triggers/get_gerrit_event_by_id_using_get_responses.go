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

// GetGerritEventByIDUsingGETReader is a Reader for the GetGerritEventByIDUsingGET structure.
type GetGerritEventByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGerritEventByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGerritEventByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGerritEventByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGerritEventByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGerritEventByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGerritEventByIDUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGerritEventByIDUsingGETOK creates a GetGerritEventByIDUsingGETOK with default headers values
func NewGetGerritEventByIDUsingGETOK() *GetGerritEventByIDUsingGETOK {
	return &GetGerritEventByIDUsingGETOK{}
}

/*
GetGerritEventByIDUsingGETOK describes a response with status code 200, with default header values.

'Success' with gerrit Event
*/
type GetGerritEventByIDUsingGETOK struct {
	Payload models.GerritEvent
}

// IsSuccess returns true when this get gerrit event by Id using g e t o k response has a 2xx status code
func (o *GetGerritEventByIDUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gerrit event by Id using g e t o k response has a 3xx status code
func (o *GetGerritEventByIDUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit event by Id using g e t o k response has a 4xx status code
func (o *GetGerritEventByIDUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gerrit event by Id using g e t o k response has a 5xx status code
func (o *GetGerritEventByIDUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit event by Id using g e t o k response a status code equal to that given
func (o *GetGerritEventByIDUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGerritEventByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-events/{id}][%d] getGerritEventByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGerritEventByIDUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-events/{id}][%d] getGerritEventByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGerritEventByIDUsingGETOK) GetPayload() models.GerritEvent {
	return o.Payload
}

func (o *GetGerritEventByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritEvent(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetGerritEventByIDUsingGETUnauthorized creates a GetGerritEventByIDUsingGETUnauthorized with default headers values
func NewGetGerritEventByIDUsingGETUnauthorized() *GetGerritEventByIDUsingGETUnauthorized {
	return &GetGerritEventByIDUsingGETUnauthorized{}
}

/*
GetGerritEventByIDUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetGerritEventByIDUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get gerrit event by Id using g e t unauthorized response has a 2xx status code
func (o *GetGerritEventByIDUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit event by Id using g e t unauthorized response has a 3xx status code
func (o *GetGerritEventByIDUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit event by Id using g e t unauthorized response has a 4xx status code
func (o *GetGerritEventByIDUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gerrit event by Id using g e t unauthorized response has a 5xx status code
func (o *GetGerritEventByIDUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit event by Id using g e t unauthorized response a status code equal to that given
func (o *GetGerritEventByIDUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetGerritEventByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-events/{id}][%d] getGerritEventByIdUsingGETUnauthorized ", 401)
}

func (o *GetGerritEventByIDUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-events/{id}][%d] getGerritEventByIdUsingGETUnauthorized ", 401)
}

func (o *GetGerritEventByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGerritEventByIDUsingGETForbidden creates a GetGerritEventByIDUsingGETForbidden with default headers values
func NewGetGerritEventByIDUsingGETForbidden() *GetGerritEventByIDUsingGETForbidden {
	return &GetGerritEventByIDUsingGETForbidden{}
}

/*
GetGerritEventByIDUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetGerritEventByIDUsingGETForbidden struct {
}

// IsSuccess returns true when this get gerrit event by Id using g e t forbidden response has a 2xx status code
func (o *GetGerritEventByIDUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit event by Id using g e t forbidden response has a 3xx status code
func (o *GetGerritEventByIDUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit event by Id using g e t forbidden response has a 4xx status code
func (o *GetGerritEventByIDUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gerrit event by Id using g e t forbidden response has a 5xx status code
func (o *GetGerritEventByIDUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit event by Id using g e t forbidden response a status code equal to that given
func (o *GetGerritEventByIDUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetGerritEventByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-events/{id}][%d] getGerritEventByIdUsingGETForbidden ", 403)
}

func (o *GetGerritEventByIDUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-events/{id}][%d] getGerritEventByIdUsingGETForbidden ", 403)
}

func (o *GetGerritEventByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGerritEventByIDUsingGETNotFound creates a GetGerritEventByIDUsingGETNotFound with default headers values
func NewGetGerritEventByIDUsingGETNotFound() *GetGerritEventByIDUsingGETNotFound {
	return &GetGerritEventByIDUsingGETNotFound{}
}

/*
GetGerritEventByIDUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetGerritEventByIDUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get gerrit event by Id using g e t not found response has a 2xx status code
func (o *GetGerritEventByIDUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit event by Id using g e t not found response has a 3xx status code
func (o *GetGerritEventByIDUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit event by Id using g e t not found response has a 4xx status code
func (o *GetGerritEventByIDUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gerrit event by Id using g e t not found response has a 5xx status code
func (o *GetGerritEventByIDUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit event by Id using g e t not found response a status code equal to that given
func (o *GetGerritEventByIDUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetGerritEventByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-events/{id}][%d] getGerritEventByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGerritEventByIDUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-events/{id}][%d] getGerritEventByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGerritEventByIDUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGerritEventByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGerritEventByIDUsingGETInternalServerError creates a GetGerritEventByIDUsingGETInternalServerError with default headers values
func NewGetGerritEventByIDUsingGETInternalServerError() *GetGerritEventByIDUsingGETInternalServerError {
	return &GetGerritEventByIDUsingGETInternalServerError{}
}

/*
GetGerritEventByIDUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetGerritEventByIDUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get gerrit event by Id using g e t internal server error response has a 2xx status code
func (o *GetGerritEventByIDUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit event by Id using g e t internal server error response has a 3xx status code
func (o *GetGerritEventByIDUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit event by Id using g e t internal server error response has a 4xx status code
func (o *GetGerritEventByIDUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gerrit event by Id using g e t internal server error response has a 5xx status code
func (o *GetGerritEventByIDUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get gerrit event by Id using g e t internal server error response a status code equal to that given
func (o *GetGerritEventByIDUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetGerritEventByIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-events/{id}][%d] getGerritEventByIdUsingGETInternalServerError ", 500)
}

func (o *GetGerritEventByIDUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-events/{id}][%d] getGerritEventByIdUsingGETInternalServerError ", 500)
}

func (o *GetGerritEventByIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}