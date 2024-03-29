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

// GetAllGerritTriggersUsingGETReader is a Reader for the GetAllGerritTriggersUsingGET structure.
type GetAllGerritTriggersUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllGerritTriggersUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllGerritTriggersUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllGerritTriggersUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllGerritTriggersUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllGerritTriggersUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllGerritTriggersUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllGerritTriggersUsingGETOK creates a GetAllGerritTriggersUsingGETOK with default headers values
func NewGetAllGerritTriggersUsingGETOK() *GetAllGerritTriggersUsingGETOK {
	return &GetAllGerritTriggersUsingGETOK{}
}

/*
GetAllGerritTriggersUsingGETOK describes a response with status code 200, with default header values.

'Success' with get of gerrit triggers
*/
type GetAllGerritTriggersUsingGETOK struct {
	Payload models.GerritTriggers
}

// IsSuccess returns true when this get all gerrit triggers using g e t o k response has a 2xx status code
func (o *GetAllGerritTriggersUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all gerrit triggers using g e t o k response has a 3xx status code
func (o *GetAllGerritTriggersUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all gerrit triggers using g e t o k response has a 4xx status code
func (o *GetAllGerritTriggersUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all gerrit triggers using g e t o k response has a 5xx status code
func (o *GetAllGerritTriggersUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all gerrit triggers using g e t o k response a status code equal to that given
func (o *GetAllGerritTriggersUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllGerritTriggersUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers][%d] getAllGerritTriggersUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllGerritTriggersUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers][%d] getAllGerritTriggersUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllGerritTriggersUsingGETOK) GetPayload() models.GerritTriggers {
	return o.Payload
}

func (o *GetAllGerritTriggersUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritTriggers(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetAllGerritTriggersUsingGETUnauthorized creates a GetAllGerritTriggersUsingGETUnauthorized with default headers values
func NewGetAllGerritTriggersUsingGETUnauthorized() *GetAllGerritTriggersUsingGETUnauthorized {
	return &GetAllGerritTriggersUsingGETUnauthorized{}
}

/*
GetAllGerritTriggersUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetAllGerritTriggersUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get all gerrit triggers using g e t unauthorized response has a 2xx status code
func (o *GetAllGerritTriggersUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all gerrit triggers using g e t unauthorized response has a 3xx status code
func (o *GetAllGerritTriggersUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all gerrit triggers using g e t unauthorized response has a 4xx status code
func (o *GetAllGerritTriggersUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all gerrit triggers using g e t unauthorized response has a 5xx status code
func (o *GetAllGerritTriggersUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all gerrit triggers using g e t unauthorized response a status code equal to that given
func (o *GetAllGerritTriggersUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAllGerritTriggersUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers][%d] getAllGerritTriggersUsingGETUnauthorized ", 401)
}

func (o *GetAllGerritTriggersUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers][%d] getAllGerritTriggersUsingGETUnauthorized ", 401)
}

func (o *GetAllGerritTriggersUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllGerritTriggersUsingGETForbidden creates a GetAllGerritTriggersUsingGETForbidden with default headers values
func NewGetAllGerritTriggersUsingGETForbidden() *GetAllGerritTriggersUsingGETForbidden {
	return &GetAllGerritTriggersUsingGETForbidden{}
}

/*
GetAllGerritTriggersUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllGerritTriggersUsingGETForbidden struct {
}

// IsSuccess returns true when this get all gerrit triggers using g e t forbidden response has a 2xx status code
func (o *GetAllGerritTriggersUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all gerrit triggers using g e t forbidden response has a 3xx status code
func (o *GetAllGerritTriggersUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all gerrit triggers using g e t forbidden response has a 4xx status code
func (o *GetAllGerritTriggersUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all gerrit triggers using g e t forbidden response has a 5xx status code
func (o *GetAllGerritTriggersUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get all gerrit triggers using g e t forbidden response a status code equal to that given
func (o *GetAllGerritTriggersUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAllGerritTriggersUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers][%d] getAllGerritTriggersUsingGETForbidden ", 403)
}

func (o *GetAllGerritTriggersUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers][%d] getAllGerritTriggersUsingGETForbidden ", 403)
}

func (o *GetAllGerritTriggersUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllGerritTriggersUsingGETNotFound creates a GetAllGerritTriggersUsingGETNotFound with default headers values
func NewGetAllGerritTriggersUsingGETNotFound() *GetAllGerritTriggersUsingGETNotFound {
	return &GetAllGerritTriggersUsingGETNotFound{}
}

/*
GetAllGerritTriggersUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllGerritTriggersUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get all gerrit triggers using g e t not found response has a 2xx status code
func (o *GetAllGerritTriggersUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all gerrit triggers using g e t not found response has a 3xx status code
func (o *GetAllGerritTriggersUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all gerrit triggers using g e t not found response has a 4xx status code
func (o *GetAllGerritTriggersUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all gerrit triggers using g e t not found response has a 5xx status code
func (o *GetAllGerritTriggersUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all gerrit triggers using g e t not found response a status code equal to that given
func (o *GetAllGerritTriggersUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAllGerritTriggersUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers][%d] getAllGerritTriggersUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetAllGerritTriggersUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers][%d] getAllGerritTriggersUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetAllGerritTriggersUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllGerritTriggersUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllGerritTriggersUsingGETInternalServerError creates a GetAllGerritTriggersUsingGETInternalServerError with default headers values
func NewGetAllGerritTriggersUsingGETInternalServerError() *GetAllGerritTriggersUsingGETInternalServerError {
	return &GetAllGerritTriggersUsingGETInternalServerError{}
}

/*
GetAllGerritTriggersUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetAllGerritTriggersUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get all gerrit triggers using g e t internal server error response has a 2xx status code
func (o *GetAllGerritTriggersUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all gerrit triggers using g e t internal server error response has a 3xx status code
func (o *GetAllGerritTriggersUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all gerrit triggers using g e t internal server error response has a 4xx status code
func (o *GetAllGerritTriggersUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all gerrit triggers using g e t internal server error response has a 5xx status code
func (o *GetAllGerritTriggersUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all gerrit triggers using g e t internal server error response a status code equal to that given
func (o *GetAllGerritTriggersUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAllGerritTriggersUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers][%d] getAllGerritTriggersUsingGETInternalServerError ", 500)
}

func (o *GetAllGerritTriggersUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers][%d] getAllGerritTriggersUsingGETInternalServerError ", 500)
}

func (o *GetAllGerritTriggersUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
