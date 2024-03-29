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

// ExecutePipelineByNameUsingPOSTReader is a Reader for the ExecutePipelineByNameUsingPOST structure.
type ExecutePipelineByNameUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecutePipelineByNameUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecutePipelineByNameUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewExecutePipelineByNameUsingPOSTAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewExecutePipelineByNameUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExecutePipelineByNameUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExecutePipelineByNameUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExecutePipelineByNameUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExecutePipelineByNameUsingPOSTOK creates a ExecutePipelineByNameUsingPOSTOK with default headers values
func NewExecutePipelineByNameUsingPOSTOK() *ExecutePipelineByNameUsingPOSTOK {
	return &ExecutePipelineByNameUsingPOSTOK{}
}

/*
ExecutePipelineByNameUsingPOSTOK describes a response with status code 200, with default header values.

'Success' with the created execution response
*/
type ExecutePipelineByNameUsingPOSTOK struct {
	Payload models.ExecutionResponse
}

// IsSuccess returns true when this execute pipeline by name using p o s t o k response has a 2xx status code
func (o *ExecutePipelineByNameUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this execute pipeline by name using p o s t o k response has a 3xx status code
func (o *ExecutePipelineByNameUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execute pipeline by name using p o s t o k response has a 4xx status code
func (o *ExecutePipelineByNameUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this execute pipeline by name using p o s t o k response has a 5xx status code
func (o *ExecutePipelineByNameUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this execute pipeline by name using p o s t o k response a status code equal to that given
func (o *ExecutePipelineByNameUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExecutePipelineByNameUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}/executions][%d] executePipelineByNameUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ExecutePipelineByNameUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}/executions][%d] executePipelineByNameUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ExecutePipelineByNameUsingPOSTOK) GetPayload() models.ExecutionResponse {
	return o.Payload
}

func (o *ExecutePipelineByNameUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecutionResponse(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewExecutePipelineByNameUsingPOSTAccepted creates a ExecutePipelineByNameUsingPOSTAccepted with default headers values
func NewExecutePipelineByNameUsingPOSTAccepted() *ExecutePipelineByNameUsingPOSTAccepted {
	return &ExecutePipelineByNameUsingPOSTAccepted{}
}

/*
ExecutePipelineByNameUsingPOSTAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ExecutePipelineByNameUsingPOSTAccepted struct {
	Payload models.ExecutionResponse
}

// IsSuccess returns true when this execute pipeline by name using p o s t accepted response has a 2xx status code
func (o *ExecutePipelineByNameUsingPOSTAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this execute pipeline by name using p o s t accepted response has a 3xx status code
func (o *ExecutePipelineByNameUsingPOSTAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execute pipeline by name using p o s t accepted response has a 4xx status code
func (o *ExecutePipelineByNameUsingPOSTAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this execute pipeline by name using p o s t accepted response has a 5xx status code
func (o *ExecutePipelineByNameUsingPOSTAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this execute pipeline by name using p o s t accepted response a status code equal to that given
func (o *ExecutePipelineByNameUsingPOSTAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *ExecutePipelineByNameUsingPOSTAccepted) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}/executions][%d] executePipelineByNameUsingPOSTAccepted  %+v", 202, o.Payload)
}

func (o *ExecutePipelineByNameUsingPOSTAccepted) String() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}/executions][%d] executePipelineByNameUsingPOSTAccepted  %+v", 202, o.Payload)
}

func (o *ExecutePipelineByNameUsingPOSTAccepted) GetPayload() models.ExecutionResponse {
	return o.Payload
}

func (o *ExecutePipelineByNameUsingPOSTAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecutionResponse(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewExecutePipelineByNameUsingPOSTUnauthorized creates a ExecutePipelineByNameUsingPOSTUnauthorized with default headers values
func NewExecutePipelineByNameUsingPOSTUnauthorized() *ExecutePipelineByNameUsingPOSTUnauthorized {
	return &ExecutePipelineByNameUsingPOSTUnauthorized{}
}

/*
ExecutePipelineByNameUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type ExecutePipelineByNameUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this execute pipeline by name using p o s t unauthorized response has a 2xx status code
func (o *ExecutePipelineByNameUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this execute pipeline by name using p o s t unauthorized response has a 3xx status code
func (o *ExecutePipelineByNameUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execute pipeline by name using p o s t unauthorized response has a 4xx status code
func (o *ExecutePipelineByNameUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this execute pipeline by name using p o s t unauthorized response has a 5xx status code
func (o *ExecutePipelineByNameUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this execute pipeline by name using p o s t unauthorized response a status code equal to that given
func (o *ExecutePipelineByNameUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ExecutePipelineByNameUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}/executions][%d] executePipelineByNameUsingPOSTUnauthorized ", 401)
}

func (o *ExecutePipelineByNameUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}/executions][%d] executePipelineByNameUsingPOSTUnauthorized ", 401)
}

func (o *ExecutePipelineByNameUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecutePipelineByNameUsingPOSTForbidden creates a ExecutePipelineByNameUsingPOSTForbidden with default headers values
func NewExecutePipelineByNameUsingPOSTForbidden() *ExecutePipelineByNameUsingPOSTForbidden {
	return &ExecutePipelineByNameUsingPOSTForbidden{}
}

/*
ExecutePipelineByNameUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ExecutePipelineByNameUsingPOSTForbidden struct {
}

// IsSuccess returns true when this execute pipeline by name using p o s t forbidden response has a 2xx status code
func (o *ExecutePipelineByNameUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this execute pipeline by name using p o s t forbidden response has a 3xx status code
func (o *ExecutePipelineByNameUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execute pipeline by name using p o s t forbidden response has a 4xx status code
func (o *ExecutePipelineByNameUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this execute pipeline by name using p o s t forbidden response has a 5xx status code
func (o *ExecutePipelineByNameUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this execute pipeline by name using p o s t forbidden response a status code equal to that given
func (o *ExecutePipelineByNameUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ExecutePipelineByNameUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}/executions][%d] executePipelineByNameUsingPOSTForbidden ", 403)
}

func (o *ExecutePipelineByNameUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}/executions][%d] executePipelineByNameUsingPOSTForbidden ", 403)
}

func (o *ExecutePipelineByNameUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecutePipelineByNameUsingPOSTNotFound creates a ExecutePipelineByNameUsingPOSTNotFound with default headers values
func NewExecutePipelineByNameUsingPOSTNotFound() *ExecutePipelineByNameUsingPOSTNotFound {
	return &ExecutePipelineByNameUsingPOSTNotFound{}
}

/*
ExecutePipelineByNameUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ExecutePipelineByNameUsingPOSTNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this execute pipeline by name using p o s t not found response has a 2xx status code
func (o *ExecutePipelineByNameUsingPOSTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this execute pipeline by name using p o s t not found response has a 3xx status code
func (o *ExecutePipelineByNameUsingPOSTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execute pipeline by name using p o s t not found response has a 4xx status code
func (o *ExecutePipelineByNameUsingPOSTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this execute pipeline by name using p o s t not found response has a 5xx status code
func (o *ExecutePipelineByNameUsingPOSTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this execute pipeline by name using p o s t not found response a status code equal to that given
func (o *ExecutePipelineByNameUsingPOSTNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ExecutePipelineByNameUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}/executions][%d] executePipelineByNameUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *ExecutePipelineByNameUsingPOSTNotFound) String() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}/executions][%d] executePipelineByNameUsingPOSTNotFound  %+v", 404, o.Payload)
}

func (o *ExecutePipelineByNameUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ExecutePipelineByNameUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecutePipelineByNameUsingPOSTInternalServerError creates a ExecutePipelineByNameUsingPOSTInternalServerError with default headers values
func NewExecutePipelineByNameUsingPOSTInternalServerError() *ExecutePipelineByNameUsingPOSTInternalServerError {
	return &ExecutePipelineByNameUsingPOSTInternalServerError{}
}

/*
ExecutePipelineByNameUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ExecutePipelineByNameUsingPOSTInternalServerError struct {
}

// IsSuccess returns true when this execute pipeline by name using p o s t internal server error response has a 2xx status code
func (o *ExecutePipelineByNameUsingPOSTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this execute pipeline by name using p o s t internal server error response has a 3xx status code
func (o *ExecutePipelineByNameUsingPOSTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this execute pipeline by name using p o s t internal server error response has a 4xx status code
func (o *ExecutePipelineByNameUsingPOSTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this execute pipeline by name using p o s t internal server error response has a 5xx status code
func (o *ExecutePipelineByNameUsingPOSTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this execute pipeline by name using p o s t internal server error response a status code equal to that given
func (o *ExecutePipelineByNameUsingPOSTInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ExecutePipelineByNameUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}/executions][%d] executePipelineByNameUsingPOSTInternalServerError ", 500)
}

func (o *ExecutePipelineByNameUsingPOSTInternalServerError) String() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}/executions][%d] executePipelineByNameUsingPOSTInternalServerError ", 500)
}

func (o *ExecutePipelineByNameUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
