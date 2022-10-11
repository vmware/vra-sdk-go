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

// GetExecutionByIndexAndPipelineIDUsingGETReader is a Reader for the GetExecutionByIndexAndPipelineIDUsingGET structure.
type GetExecutionByIndexAndPipelineIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExecutionByIndexAndPipelineIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExecutionByIndexAndPipelineIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetExecutionByIndexAndPipelineIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExecutionByIndexAndPipelineIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExecutionByIndexAndPipelineIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExecutionByIndexAndPipelineIDUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExecutionByIndexAndPipelineIDUsingGETOK creates a GetExecutionByIndexAndPipelineIDUsingGETOK with default headers values
func NewGetExecutionByIndexAndPipelineIDUsingGETOK() *GetExecutionByIndexAndPipelineIDUsingGETOK {
	return &GetExecutionByIndexAndPipelineIDUsingGETOK{}
}

/*
GetExecutionByIndexAndPipelineIDUsingGETOK describes a response with status code 200, with default header values.

'Success' with the requested Execution
*/
type GetExecutionByIndexAndPipelineIDUsingGETOK struct {
	Payload models.Execution
}

// IsSuccess returns true when this get execution by index and pipeline Id using g e t o k response has a 2xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get execution by index and pipeline Id using g e t o k response has a 3xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get execution by index and pipeline Id using g e t o k response has a 4xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get execution by index and pipeline Id using g e t o k response has a 5xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get execution by index and pipeline Id using g e t o k response a status code equal to that given
func (o *GetExecutionByIndexAndPipelineIDUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions/{index}][%d] getExecutionByIndexAndPipelineIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions/{index}][%d] getExecutionByIndexAndPipelineIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETOK) GetPayload() models.Execution {
	return o.Payload
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecution(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetExecutionByIndexAndPipelineIDUsingGETUnauthorized creates a GetExecutionByIndexAndPipelineIDUsingGETUnauthorized with default headers values
func NewGetExecutionByIndexAndPipelineIDUsingGETUnauthorized() *GetExecutionByIndexAndPipelineIDUsingGETUnauthorized {
	return &GetExecutionByIndexAndPipelineIDUsingGETUnauthorized{}
}

/*
GetExecutionByIndexAndPipelineIDUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetExecutionByIndexAndPipelineIDUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get execution by index and pipeline Id using g e t unauthorized response has a 2xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get execution by index and pipeline Id using g e t unauthorized response has a 3xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get execution by index and pipeline Id using g e t unauthorized response has a 4xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get execution by index and pipeline Id using g e t unauthorized response has a 5xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get execution by index and pipeline Id using g e t unauthorized response a status code equal to that given
func (o *GetExecutionByIndexAndPipelineIDUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions/{index}][%d] getExecutionByIndexAndPipelineIdUsingGETUnauthorized ", 401)
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions/{index}][%d] getExecutionByIndexAndPipelineIdUsingGETUnauthorized ", 401)
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExecutionByIndexAndPipelineIDUsingGETForbidden creates a GetExecutionByIndexAndPipelineIDUsingGETForbidden with default headers values
func NewGetExecutionByIndexAndPipelineIDUsingGETForbidden() *GetExecutionByIndexAndPipelineIDUsingGETForbidden {
	return &GetExecutionByIndexAndPipelineIDUsingGETForbidden{}
}

/*
GetExecutionByIndexAndPipelineIDUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetExecutionByIndexAndPipelineIDUsingGETForbidden struct {
}

// IsSuccess returns true when this get execution by index and pipeline Id using g e t forbidden response has a 2xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get execution by index and pipeline Id using g e t forbidden response has a 3xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get execution by index and pipeline Id using g e t forbidden response has a 4xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get execution by index and pipeline Id using g e t forbidden response has a 5xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get execution by index and pipeline Id using g e t forbidden response a status code equal to that given
func (o *GetExecutionByIndexAndPipelineIDUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions/{index}][%d] getExecutionByIndexAndPipelineIdUsingGETForbidden ", 403)
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions/{index}][%d] getExecutionByIndexAndPipelineIdUsingGETForbidden ", 403)
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExecutionByIndexAndPipelineIDUsingGETNotFound creates a GetExecutionByIndexAndPipelineIDUsingGETNotFound with default headers values
func NewGetExecutionByIndexAndPipelineIDUsingGETNotFound() *GetExecutionByIndexAndPipelineIDUsingGETNotFound {
	return &GetExecutionByIndexAndPipelineIDUsingGETNotFound{}
}

/*
GetExecutionByIndexAndPipelineIDUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetExecutionByIndexAndPipelineIDUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get execution by index and pipeline Id using g e t not found response has a 2xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get execution by index and pipeline Id using g e t not found response has a 3xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get execution by index and pipeline Id using g e t not found response has a 4xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get execution by index and pipeline Id using g e t not found response has a 5xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get execution by index and pipeline Id using g e t not found response a status code equal to that given
func (o *GetExecutionByIndexAndPipelineIDUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions/{index}][%d] getExecutionByIndexAndPipelineIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions/{index}][%d] getExecutionByIndexAndPipelineIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExecutionByIndexAndPipelineIDUsingGETInternalServerError creates a GetExecutionByIndexAndPipelineIDUsingGETInternalServerError with default headers values
func NewGetExecutionByIndexAndPipelineIDUsingGETInternalServerError() *GetExecutionByIndexAndPipelineIDUsingGETInternalServerError {
	return &GetExecutionByIndexAndPipelineIDUsingGETInternalServerError{}
}

/*
GetExecutionByIndexAndPipelineIDUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetExecutionByIndexAndPipelineIDUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get execution by index and pipeline Id using g e t internal server error response has a 2xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get execution by index and pipeline Id using g e t internal server error response has a 3xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get execution by index and pipeline Id using g e t internal server error response has a 4xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get execution by index and pipeline Id using g e t internal server error response has a 5xx status code
func (o *GetExecutionByIndexAndPipelineIDUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get execution by index and pipeline Id using g e t internal server error response a status code equal to that given
func (o *GetExecutionByIndexAndPipelineIDUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions/{index}][%d] getExecutionByIndexAndPipelineIdUsingGETInternalServerError ", 500)
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions/{index}][%d] getExecutionByIndexAndPipelineIdUsingGETInternalServerError ", 500)
}

func (o *GetExecutionByIndexAndPipelineIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
