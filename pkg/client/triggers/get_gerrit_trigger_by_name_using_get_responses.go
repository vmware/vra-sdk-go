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

// GetGerritTriggerByNameUsingGETReader is a Reader for the GetGerritTriggerByNameUsingGET structure.
type GetGerritTriggerByNameUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGerritTriggerByNameUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGerritTriggerByNameUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGerritTriggerByNameUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGerritTriggerByNameUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGerritTriggerByNameUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGerritTriggerByNameUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGerritTriggerByNameUsingGETOK creates a GetGerritTriggerByNameUsingGETOK with default headers values
func NewGetGerritTriggerByNameUsingGETOK() *GetGerritTriggerByNameUsingGETOK {
	return &GetGerritTriggerByNameUsingGETOK{}
}

/*
GetGerritTriggerByNameUsingGETOK describes a response with status code 200, with default header values.

'Success' with gerrit trigger
*/
type GetGerritTriggerByNameUsingGETOK struct {
	Payload models.GerritTrigger
}

// IsSuccess returns true when this get gerrit trigger by name using g e t o k response has a 2xx status code
func (o *GetGerritTriggerByNameUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gerrit trigger by name using g e t o k response has a 3xx status code
func (o *GetGerritTriggerByNameUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit trigger by name using g e t o k response has a 4xx status code
func (o *GetGerritTriggerByNameUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gerrit trigger by name using g e t o k response has a 5xx status code
func (o *GetGerritTriggerByNameUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit trigger by name using g e t o k response a status code equal to that given
func (o *GetGerritTriggerByNameUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGerritTriggerByNameUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getGerritTriggerByNameUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGerritTriggerByNameUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getGerritTriggerByNameUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetGerritTriggerByNameUsingGETOK) GetPayload() models.GerritTrigger {
	return o.Payload
}

func (o *GetGerritTriggerByNameUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritTrigger(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetGerritTriggerByNameUsingGETUnauthorized creates a GetGerritTriggerByNameUsingGETUnauthorized with default headers values
func NewGetGerritTriggerByNameUsingGETUnauthorized() *GetGerritTriggerByNameUsingGETUnauthorized {
	return &GetGerritTriggerByNameUsingGETUnauthorized{}
}

/*
GetGerritTriggerByNameUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetGerritTriggerByNameUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get gerrit trigger by name using g e t unauthorized response has a 2xx status code
func (o *GetGerritTriggerByNameUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit trigger by name using g e t unauthorized response has a 3xx status code
func (o *GetGerritTriggerByNameUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit trigger by name using g e t unauthorized response has a 4xx status code
func (o *GetGerritTriggerByNameUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gerrit trigger by name using g e t unauthorized response has a 5xx status code
func (o *GetGerritTriggerByNameUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit trigger by name using g e t unauthorized response a status code equal to that given
func (o *GetGerritTriggerByNameUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetGerritTriggerByNameUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getGerritTriggerByNameUsingGETUnauthorized ", 401)
}

func (o *GetGerritTriggerByNameUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getGerritTriggerByNameUsingGETUnauthorized ", 401)
}

func (o *GetGerritTriggerByNameUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGerritTriggerByNameUsingGETForbidden creates a GetGerritTriggerByNameUsingGETForbidden with default headers values
func NewGetGerritTriggerByNameUsingGETForbidden() *GetGerritTriggerByNameUsingGETForbidden {
	return &GetGerritTriggerByNameUsingGETForbidden{}
}

/*
GetGerritTriggerByNameUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetGerritTriggerByNameUsingGETForbidden struct {
}

// IsSuccess returns true when this get gerrit trigger by name using g e t forbidden response has a 2xx status code
func (o *GetGerritTriggerByNameUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit trigger by name using g e t forbidden response has a 3xx status code
func (o *GetGerritTriggerByNameUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit trigger by name using g e t forbidden response has a 4xx status code
func (o *GetGerritTriggerByNameUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gerrit trigger by name using g e t forbidden response has a 5xx status code
func (o *GetGerritTriggerByNameUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit trigger by name using g e t forbidden response a status code equal to that given
func (o *GetGerritTriggerByNameUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetGerritTriggerByNameUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getGerritTriggerByNameUsingGETForbidden ", 403)
}

func (o *GetGerritTriggerByNameUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getGerritTriggerByNameUsingGETForbidden ", 403)
}

func (o *GetGerritTriggerByNameUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGerritTriggerByNameUsingGETNotFound creates a GetGerritTriggerByNameUsingGETNotFound with default headers values
func NewGetGerritTriggerByNameUsingGETNotFound() *GetGerritTriggerByNameUsingGETNotFound {
	return &GetGerritTriggerByNameUsingGETNotFound{}
}

/*
GetGerritTriggerByNameUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetGerritTriggerByNameUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get gerrit trigger by name using g e t not found response has a 2xx status code
func (o *GetGerritTriggerByNameUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit trigger by name using g e t not found response has a 3xx status code
func (o *GetGerritTriggerByNameUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit trigger by name using g e t not found response has a 4xx status code
func (o *GetGerritTriggerByNameUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gerrit trigger by name using g e t not found response has a 5xx status code
func (o *GetGerritTriggerByNameUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get gerrit trigger by name using g e t not found response a status code equal to that given
func (o *GetGerritTriggerByNameUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetGerritTriggerByNameUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getGerritTriggerByNameUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGerritTriggerByNameUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getGerritTriggerByNameUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetGerritTriggerByNameUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGerritTriggerByNameUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGerritTriggerByNameUsingGETInternalServerError creates a GetGerritTriggerByNameUsingGETInternalServerError with default headers values
func NewGetGerritTriggerByNameUsingGETInternalServerError() *GetGerritTriggerByNameUsingGETInternalServerError {
	return &GetGerritTriggerByNameUsingGETInternalServerError{}
}

/*
GetGerritTriggerByNameUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetGerritTriggerByNameUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get gerrit trigger by name using g e t internal server error response has a 2xx status code
func (o *GetGerritTriggerByNameUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gerrit trigger by name using g e t internal server error response has a 3xx status code
func (o *GetGerritTriggerByNameUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gerrit trigger by name using g e t internal server error response has a 4xx status code
func (o *GetGerritTriggerByNameUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gerrit trigger by name using g e t internal server error response has a 5xx status code
func (o *GetGerritTriggerByNameUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get gerrit trigger by name using g e t internal server error response a status code equal to that given
func (o *GetGerritTriggerByNameUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetGerritTriggerByNameUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getGerritTriggerByNameUsingGETInternalServerError ", 500)
}

func (o *GetGerritTriggerByNameUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{project}/{name}][%d] getGerritTriggerByNameUsingGETInternalServerError ", 500)
}

func (o *GetGerritTriggerByNameUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
