// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetExecutionsByNameUsingGETReader is a Reader for the GetExecutionsByNameUsingGET structure.
type GetExecutionsByNameUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExecutionsByNameUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExecutionsByNameUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetExecutionsByNameUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExecutionsByNameUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExecutionsByNameUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExecutionsByNameUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExecutionsByNameUsingGETOK creates a GetExecutionsByNameUsingGETOK with default headers values
func NewGetExecutionsByNameUsingGETOK() *GetExecutionsByNameUsingGETOK {
	return &GetExecutionsByNameUsingGETOK{}
}

/*
GetExecutionsByNameUsingGETOK describes a response with status code 200, with default header values.

'Success' with Executions on pages
*/
type GetExecutionsByNameUsingGETOK struct {
	Payload models.Executions
}

// IsSuccess returns true when this get executions by name using g e t o k response has a 2xx status code
func (o *GetExecutionsByNameUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get executions by name using g e t o k response has a 3xx status code
func (o *GetExecutionsByNameUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get executions by name using g e t o k response has a 4xx status code
func (o *GetExecutionsByNameUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get executions by name using g e t o k response has a 5xx status code
func (o *GetExecutionsByNameUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get executions by name using g e t o k response a status code equal to that given
func (o *GetExecutionsByNameUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetExecutionsByNameUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetExecutionsByNameUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetExecutionsByNameUsingGETOK) GetPayload() models.Executions {
	return o.Payload
}

func (o *GetExecutionsByNameUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecutions(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetExecutionsByNameUsingGETUnauthorized creates a GetExecutionsByNameUsingGETUnauthorized with default headers values
func NewGetExecutionsByNameUsingGETUnauthorized() *GetExecutionsByNameUsingGETUnauthorized {
	return &GetExecutionsByNameUsingGETUnauthorized{}
}

/*
GetExecutionsByNameUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetExecutionsByNameUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get executions by name using g e t unauthorized response has a 2xx status code
func (o *GetExecutionsByNameUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get executions by name using g e t unauthorized response has a 3xx status code
func (o *GetExecutionsByNameUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get executions by name using g e t unauthorized response has a 4xx status code
func (o *GetExecutionsByNameUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get executions by name using g e t unauthorized response has a 5xx status code
func (o *GetExecutionsByNameUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get executions by name using g e t unauthorized response a status code equal to that given
func (o *GetExecutionsByNameUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetExecutionsByNameUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETUnauthorized ", 401)
}

func (o *GetExecutionsByNameUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETUnauthorized ", 401)
}

func (o *GetExecutionsByNameUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExecutionsByNameUsingGETForbidden creates a GetExecutionsByNameUsingGETForbidden with default headers values
func NewGetExecutionsByNameUsingGETForbidden() *GetExecutionsByNameUsingGETForbidden {
	return &GetExecutionsByNameUsingGETForbidden{}
}

/*
GetExecutionsByNameUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetExecutionsByNameUsingGETForbidden struct {
}

// IsSuccess returns true when this get executions by name using g e t forbidden response has a 2xx status code
func (o *GetExecutionsByNameUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get executions by name using g e t forbidden response has a 3xx status code
func (o *GetExecutionsByNameUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get executions by name using g e t forbidden response has a 4xx status code
func (o *GetExecutionsByNameUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get executions by name using g e t forbidden response has a 5xx status code
func (o *GetExecutionsByNameUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get executions by name using g e t forbidden response a status code equal to that given
func (o *GetExecutionsByNameUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetExecutionsByNameUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETForbidden ", 403)
}

func (o *GetExecutionsByNameUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETForbidden ", 403)
}

func (o *GetExecutionsByNameUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExecutionsByNameUsingGETNotFound creates a GetExecutionsByNameUsingGETNotFound with default headers values
func NewGetExecutionsByNameUsingGETNotFound() *GetExecutionsByNameUsingGETNotFound {
	return &GetExecutionsByNameUsingGETNotFound{}
}

/*
GetExecutionsByNameUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetExecutionsByNameUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get executions by name using g e t not found response has a 2xx status code
func (o *GetExecutionsByNameUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get executions by name using g e t not found response has a 3xx status code
func (o *GetExecutionsByNameUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get executions by name using g e t not found response has a 4xx status code
func (o *GetExecutionsByNameUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get executions by name using g e t not found response has a 5xx status code
func (o *GetExecutionsByNameUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get executions by name using g e t not found response a status code equal to that given
func (o *GetExecutionsByNameUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetExecutionsByNameUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetExecutionsByNameUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetExecutionsByNameUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetExecutionsByNameUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExecutionsByNameUsingGETInternalServerError creates a GetExecutionsByNameUsingGETInternalServerError with default headers values
func NewGetExecutionsByNameUsingGETInternalServerError() *GetExecutionsByNameUsingGETInternalServerError {
	return &GetExecutionsByNameUsingGETInternalServerError{}
}

/*
GetExecutionsByNameUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetExecutionsByNameUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get executions by name using g e t internal server error response has a 2xx status code
func (o *GetExecutionsByNameUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get executions by name using g e t internal server error response has a 3xx status code
func (o *GetExecutionsByNameUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get executions by name using g e t internal server error response has a 4xx status code
func (o *GetExecutionsByNameUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get executions by name using g e t internal server error response has a 5xx status code
func (o *GetExecutionsByNameUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get executions by name using g e t internal server error response a status code equal to that given
func (o *GetExecutionsByNameUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetExecutionsByNameUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETInternalServerError ", 500)
}

func (o *GetExecutionsByNameUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETInternalServerError ", 500)
}

func (o *GetExecutionsByNameUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
