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

// GetPipelineByIDUsingGETReader is a Reader for the GetPipelineByIDUsingGET structure.
type GetPipelineByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelineByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelineByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPipelineByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPipelineByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPipelineByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPipelineByIDUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPipelineByIDUsingGETOK creates a GetPipelineByIDUsingGETOK with default headers values
func NewGetPipelineByIDUsingGETOK() *GetPipelineByIDUsingGETOK {
	return &GetPipelineByIDUsingGETOK{}
}

/*
GetPipelineByIDUsingGETOK describes a response with status code 200, with default header values.

'Success' with the requested Pipeline
*/
type GetPipelineByIDUsingGETOK struct {
	Payload models.Pipeline
}

// IsSuccess returns true when this get pipeline by Id using g e t o k response has a 2xx status code
func (o *GetPipelineByIDUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get pipeline by Id using g e t o k response has a 3xx status code
func (o *GetPipelineByIDUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline by Id using g e t o k response has a 4xx status code
func (o *GetPipelineByIDUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pipeline by Id using g e t o k response has a 5xx status code
func (o *GetPipelineByIDUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline by Id using g e t o k response a status code equal to that given
func (o *GetPipelineByIDUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPipelineByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getPipelineByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPipelineByIDUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getPipelineByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPipelineByIDUsingGETOK) GetPayload() models.Pipeline {
	return o.Payload
}

func (o *GetPipelineByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalPipeline(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetPipelineByIDUsingGETUnauthorized creates a GetPipelineByIDUsingGETUnauthorized with default headers values
func NewGetPipelineByIDUsingGETUnauthorized() *GetPipelineByIDUsingGETUnauthorized {
	return &GetPipelineByIDUsingGETUnauthorized{}
}

/*
GetPipelineByIDUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetPipelineByIDUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get pipeline by Id using g e t unauthorized response has a 2xx status code
func (o *GetPipelineByIDUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipeline by Id using g e t unauthorized response has a 3xx status code
func (o *GetPipelineByIDUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline by Id using g e t unauthorized response has a 4xx status code
func (o *GetPipelineByIDUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pipeline by Id using g e t unauthorized response has a 5xx status code
func (o *GetPipelineByIDUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline by Id using g e t unauthorized response a status code equal to that given
func (o *GetPipelineByIDUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPipelineByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getPipelineByIdUsingGETUnauthorized ", 401)
}

func (o *GetPipelineByIDUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getPipelineByIdUsingGETUnauthorized ", 401)
}

func (o *GetPipelineByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPipelineByIDUsingGETForbidden creates a GetPipelineByIDUsingGETForbidden with default headers values
func NewGetPipelineByIDUsingGETForbidden() *GetPipelineByIDUsingGETForbidden {
	return &GetPipelineByIDUsingGETForbidden{}
}

/*
GetPipelineByIDUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPipelineByIDUsingGETForbidden struct {
}

// IsSuccess returns true when this get pipeline by Id using g e t forbidden response has a 2xx status code
func (o *GetPipelineByIDUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipeline by Id using g e t forbidden response has a 3xx status code
func (o *GetPipelineByIDUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline by Id using g e t forbidden response has a 4xx status code
func (o *GetPipelineByIDUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pipeline by Id using g e t forbidden response has a 5xx status code
func (o *GetPipelineByIDUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline by Id using g e t forbidden response a status code equal to that given
func (o *GetPipelineByIDUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPipelineByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getPipelineByIdUsingGETForbidden ", 403)
}

func (o *GetPipelineByIDUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getPipelineByIdUsingGETForbidden ", 403)
}

func (o *GetPipelineByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPipelineByIDUsingGETNotFound creates a GetPipelineByIDUsingGETNotFound with default headers values
func NewGetPipelineByIDUsingGETNotFound() *GetPipelineByIDUsingGETNotFound {
	return &GetPipelineByIDUsingGETNotFound{}
}

/*
GetPipelineByIDUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetPipelineByIDUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get pipeline by Id using g e t not found response has a 2xx status code
func (o *GetPipelineByIDUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipeline by Id using g e t not found response has a 3xx status code
func (o *GetPipelineByIDUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline by Id using g e t not found response has a 4xx status code
func (o *GetPipelineByIDUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pipeline by Id using g e t not found response has a 5xx status code
func (o *GetPipelineByIDUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline by Id using g e t not found response a status code equal to that given
func (o *GetPipelineByIDUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPipelineByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getPipelineByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetPipelineByIDUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getPipelineByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetPipelineByIDUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPipelineByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelineByIDUsingGETInternalServerError creates a GetPipelineByIDUsingGETInternalServerError with default headers values
func NewGetPipelineByIDUsingGETInternalServerError() *GetPipelineByIDUsingGETInternalServerError {
	return &GetPipelineByIDUsingGETInternalServerError{}
}

/*
GetPipelineByIDUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetPipelineByIDUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get pipeline by Id using g e t internal server error response has a 2xx status code
func (o *GetPipelineByIDUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipeline by Id using g e t internal server error response has a 3xx status code
func (o *GetPipelineByIDUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline by Id using g e t internal server error response has a 4xx status code
func (o *GetPipelineByIDUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pipeline by Id using g e t internal server error response has a 5xx status code
func (o *GetPipelineByIDUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get pipeline by Id using g e t internal server error response a status code equal to that given
func (o *GetPipelineByIDUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPipelineByIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getPipelineByIdUsingGETInternalServerError ", 500)
}

func (o *GetPipelineByIDUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getPipelineByIdUsingGETInternalServerError ", 500)
}

func (o *GetPipelineByIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
