// Code generated by go-swagger; DO NOT EDIT.

package custom_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDraftByIDUsingGETReader is a Reader for the GetDraftByIDUsingGET structure.
type GetDraftByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDraftByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDraftByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDraftByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDraftByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDraftByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDraftByIDUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDraftByIDUsingGETOK creates a GetDraftByIDUsingGETOK with default headers values
func NewGetDraftByIDUsingGETOK() *GetDraftByIDUsingGETOK {
	return &GetDraftByIDUsingGETOK{}
}

/*
GetDraftByIDUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetDraftByIDUsingGETOK struct {
	Payload models.CustomIntegration
}

// IsSuccess returns true when this get draft by Id using g e t o k response has a 2xx status code
func (o *GetDraftByIDUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get draft by Id using g e t o k response has a 3xx status code
func (o *GetDraftByIDUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get draft by Id using g e t o k response has a 4xx status code
func (o *GetDraftByIDUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get draft by Id using g e t o k response has a 5xx status code
func (o *GetDraftByIDUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get draft by Id using g e t o k response a status code equal to that given
func (o *GetDraftByIDUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDraftByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetDraftByIDUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetDraftByIDUsingGETOK) GetPayload() models.CustomIntegration {
	return o.Payload
}

func (o *GetDraftByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalCustomIntegration(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetDraftByIDUsingGETUnauthorized creates a GetDraftByIDUsingGETUnauthorized with default headers values
func NewGetDraftByIDUsingGETUnauthorized() *GetDraftByIDUsingGETUnauthorized {
	return &GetDraftByIDUsingGETUnauthorized{}
}

/*
GetDraftByIDUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetDraftByIDUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get draft by Id using g e t unauthorized response has a 2xx status code
func (o *GetDraftByIDUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get draft by Id using g e t unauthorized response has a 3xx status code
func (o *GetDraftByIDUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get draft by Id using g e t unauthorized response has a 4xx status code
func (o *GetDraftByIDUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get draft by Id using g e t unauthorized response has a 5xx status code
func (o *GetDraftByIDUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get draft by Id using g e t unauthorized response a status code equal to that given
func (o *GetDraftByIDUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDraftByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETUnauthorized ", 401)
}

func (o *GetDraftByIDUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETUnauthorized ", 401)
}

func (o *GetDraftByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDraftByIDUsingGETForbidden creates a GetDraftByIDUsingGETForbidden with default headers values
func NewGetDraftByIDUsingGETForbidden() *GetDraftByIDUsingGETForbidden {
	return &GetDraftByIDUsingGETForbidden{}
}

/*
GetDraftByIDUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDraftByIDUsingGETForbidden struct {
}

// IsSuccess returns true when this get draft by Id using g e t forbidden response has a 2xx status code
func (o *GetDraftByIDUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get draft by Id using g e t forbidden response has a 3xx status code
func (o *GetDraftByIDUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get draft by Id using g e t forbidden response has a 4xx status code
func (o *GetDraftByIDUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get draft by Id using g e t forbidden response has a 5xx status code
func (o *GetDraftByIDUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get draft by Id using g e t forbidden response a status code equal to that given
func (o *GetDraftByIDUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDraftByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETForbidden ", 403)
}

func (o *GetDraftByIDUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETForbidden ", 403)
}

func (o *GetDraftByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDraftByIDUsingGETNotFound creates a GetDraftByIDUsingGETNotFound with default headers values
func NewGetDraftByIDUsingGETNotFound() *GetDraftByIDUsingGETNotFound {
	return &GetDraftByIDUsingGETNotFound{}
}

/*
GetDraftByIDUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDraftByIDUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get draft by Id using g e t not found response has a 2xx status code
func (o *GetDraftByIDUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get draft by Id using g e t not found response has a 3xx status code
func (o *GetDraftByIDUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get draft by Id using g e t not found response has a 4xx status code
func (o *GetDraftByIDUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get draft by Id using g e t not found response has a 5xx status code
func (o *GetDraftByIDUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get draft by Id using g e t not found response a status code equal to that given
func (o *GetDraftByIDUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDraftByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetDraftByIDUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetDraftByIDUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDraftByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDraftByIDUsingGETInternalServerError creates a GetDraftByIDUsingGETInternalServerError with default headers values
func NewGetDraftByIDUsingGETInternalServerError() *GetDraftByIDUsingGETInternalServerError {
	return &GetDraftByIDUsingGETInternalServerError{}
}

/*
GetDraftByIDUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetDraftByIDUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get draft by Id using g e t internal server error response has a 2xx status code
func (o *GetDraftByIDUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get draft by Id using g e t internal server error response has a 3xx status code
func (o *GetDraftByIDUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get draft by Id using g e t internal server error response has a 4xx status code
func (o *GetDraftByIDUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get draft by Id using g e t internal server error response has a 5xx status code
func (o *GetDraftByIDUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get draft by Id using g e t internal server error response a status code equal to that given
func (o *GetDraftByIDUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDraftByIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETInternalServerError ", 500)
}

func (o *GetDraftByIDUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETInternalServerError ", 500)
}

func (o *GetDraftByIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
