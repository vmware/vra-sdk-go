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

// GetCustomIntegrationVersionByIDUsingGETReader is a Reader for the GetCustomIntegrationVersionByIDUsingGET structure.
type GetCustomIntegrationVersionByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomIntegrationVersionByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomIntegrationVersionByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCustomIntegrationVersionByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCustomIntegrationVersionByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCustomIntegrationVersionByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCustomIntegrationVersionByIDUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCustomIntegrationVersionByIDUsingGETOK creates a GetCustomIntegrationVersionByIDUsingGETOK with default headers values
func NewGetCustomIntegrationVersionByIDUsingGETOK() *GetCustomIntegrationVersionByIDUsingGETOK {
	return &GetCustomIntegrationVersionByIDUsingGETOK{}
}

/*
GetCustomIntegrationVersionByIDUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetCustomIntegrationVersionByIDUsingGETOK struct {
	Payload models.CustomIntegration
}

// IsSuccess returns true when this get custom integration version by Id using g e t o k response has a 2xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get custom integration version by Id using g e t o k response has a 3xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom integration version by Id using g e t o k response has a 4xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get custom integration version by Id using g e t o k response has a 5xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom integration version by Id using g e t o k response a status code equal to that given
func (o *GetCustomIntegrationVersionByIDUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCustomIntegrationVersionByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}/versions/{version}][%d] getCustomIntegrationVersionByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetCustomIntegrationVersionByIDUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}/versions/{version}][%d] getCustomIntegrationVersionByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetCustomIntegrationVersionByIDUsingGETOK) GetPayload() models.CustomIntegration {
	return o.Payload
}

func (o *GetCustomIntegrationVersionByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalCustomIntegration(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetCustomIntegrationVersionByIDUsingGETUnauthorized creates a GetCustomIntegrationVersionByIDUsingGETUnauthorized with default headers values
func NewGetCustomIntegrationVersionByIDUsingGETUnauthorized() *GetCustomIntegrationVersionByIDUsingGETUnauthorized {
	return &GetCustomIntegrationVersionByIDUsingGETUnauthorized{}
}

/*
GetCustomIntegrationVersionByIDUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetCustomIntegrationVersionByIDUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get custom integration version by Id using g e t unauthorized response has a 2xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get custom integration version by Id using g e t unauthorized response has a 3xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom integration version by Id using g e t unauthorized response has a 4xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get custom integration version by Id using g e t unauthorized response has a 5xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom integration version by Id using g e t unauthorized response a status code equal to that given
func (o *GetCustomIntegrationVersionByIDUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetCustomIntegrationVersionByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}/versions/{version}][%d] getCustomIntegrationVersionByIdUsingGETUnauthorized ", 401)
}

func (o *GetCustomIntegrationVersionByIDUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}/versions/{version}][%d] getCustomIntegrationVersionByIdUsingGETUnauthorized ", 401)
}

func (o *GetCustomIntegrationVersionByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomIntegrationVersionByIDUsingGETForbidden creates a GetCustomIntegrationVersionByIDUsingGETForbidden with default headers values
func NewGetCustomIntegrationVersionByIDUsingGETForbidden() *GetCustomIntegrationVersionByIDUsingGETForbidden {
	return &GetCustomIntegrationVersionByIDUsingGETForbidden{}
}

/*
GetCustomIntegrationVersionByIDUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCustomIntegrationVersionByIDUsingGETForbidden struct {
}

// IsSuccess returns true when this get custom integration version by Id using g e t forbidden response has a 2xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get custom integration version by Id using g e t forbidden response has a 3xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom integration version by Id using g e t forbidden response has a 4xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get custom integration version by Id using g e t forbidden response has a 5xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom integration version by Id using g e t forbidden response a status code equal to that given
func (o *GetCustomIntegrationVersionByIDUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCustomIntegrationVersionByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}/versions/{version}][%d] getCustomIntegrationVersionByIdUsingGETForbidden ", 403)
}

func (o *GetCustomIntegrationVersionByIDUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}/versions/{version}][%d] getCustomIntegrationVersionByIdUsingGETForbidden ", 403)
}

func (o *GetCustomIntegrationVersionByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCustomIntegrationVersionByIDUsingGETNotFound creates a GetCustomIntegrationVersionByIDUsingGETNotFound with default headers values
func NewGetCustomIntegrationVersionByIDUsingGETNotFound() *GetCustomIntegrationVersionByIDUsingGETNotFound {
	return &GetCustomIntegrationVersionByIDUsingGETNotFound{}
}

/*
GetCustomIntegrationVersionByIDUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetCustomIntegrationVersionByIDUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get custom integration version by Id using g e t not found response has a 2xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get custom integration version by Id using g e t not found response has a 3xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom integration version by Id using g e t not found response has a 4xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get custom integration version by Id using g e t not found response has a 5xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom integration version by Id using g e t not found response a status code equal to that given
func (o *GetCustomIntegrationVersionByIDUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCustomIntegrationVersionByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}/versions/{version}][%d] getCustomIntegrationVersionByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetCustomIntegrationVersionByIDUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}/versions/{version}][%d] getCustomIntegrationVersionByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetCustomIntegrationVersionByIDUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCustomIntegrationVersionByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomIntegrationVersionByIDUsingGETInternalServerError creates a GetCustomIntegrationVersionByIDUsingGETInternalServerError with default headers values
func NewGetCustomIntegrationVersionByIDUsingGETInternalServerError() *GetCustomIntegrationVersionByIDUsingGETInternalServerError {
	return &GetCustomIntegrationVersionByIDUsingGETInternalServerError{}
}

/*
GetCustomIntegrationVersionByIDUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetCustomIntegrationVersionByIDUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get custom integration version by Id using g e t internal server error response has a 2xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get custom integration version by Id using g e t internal server error response has a 3xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom integration version by Id using g e t internal server error response has a 4xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get custom integration version by Id using g e t internal server error response has a 5xx status code
func (o *GetCustomIntegrationVersionByIDUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get custom integration version by Id using g e t internal server error response a status code equal to that given
func (o *GetCustomIntegrationVersionByIDUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCustomIntegrationVersionByIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}/versions/{version}][%d] getCustomIntegrationVersionByIdUsingGETInternalServerError ", 500)
}

func (o *GetCustomIntegrationVersionByIDUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}/versions/{version}][%d] getCustomIntegrationVersionByIdUsingGETInternalServerError ", 500)
}

func (o *GetCustomIntegrationVersionByIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
