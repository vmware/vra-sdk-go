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

// ExecutePipelineByIDUsingPOSTReader is a Reader for the ExecutePipelineByIDUsingPOST structure.
type ExecutePipelineByIDUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecutePipelineByIDUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecutePipelineByIDUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewExecutePipelineByIDUsingPOSTAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewExecutePipelineByIDUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExecutePipelineByIDUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExecutePipelineByIDUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExecutePipelineByIDUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExecutePipelineByIDUsingPOSTOK creates a ExecutePipelineByIDUsingPOSTOK with default headers values
func NewExecutePipelineByIDUsingPOSTOK() *ExecutePipelineByIDUsingPOSTOK {
	return &ExecutePipelineByIDUsingPOSTOK{}
}

/* ExecutePipelineByIDUsingPOSTOK describes a response with status code 200, with default header values.

'Success' with the created execution response
*/
type ExecutePipelineByIDUsingPOSTOK struct {
	Payload models.ExecutionResponse
}

func (o *ExecutePipelineByIDUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}/executions][%d] executePipelineByIdUsingPOSTOK  %+v", 200, o.Payload)
}
func (o *ExecutePipelineByIDUsingPOSTOK) GetPayload() models.ExecutionResponse {
	return o.Payload
}

func (o *ExecutePipelineByIDUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecutionResponse(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewExecutePipelineByIDUsingPOSTAccepted creates a ExecutePipelineByIDUsingPOSTAccepted with default headers values
func NewExecutePipelineByIDUsingPOSTAccepted() *ExecutePipelineByIDUsingPOSTAccepted {
	return &ExecutePipelineByIDUsingPOSTAccepted{}
}

/* ExecutePipelineByIDUsingPOSTAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ExecutePipelineByIDUsingPOSTAccepted struct {
	Payload models.ExecutionResponse
}

func (o *ExecutePipelineByIDUsingPOSTAccepted) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}/executions][%d] executePipelineByIdUsingPOSTAccepted  %+v", 202, o.Payload)
}
func (o *ExecutePipelineByIDUsingPOSTAccepted) GetPayload() models.ExecutionResponse {
	return o.Payload
}

func (o *ExecutePipelineByIDUsingPOSTAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecutionResponse(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewExecutePipelineByIDUsingPOSTUnauthorized creates a ExecutePipelineByIDUsingPOSTUnauthorized with default headers values
func NewExecutePipelineByIDUsingPOSTUnauthorized() *ExecutePipelineByIDUsingPOSTUnauthorized {
	return &ExecutePipelineByIDUsingPOSTUnauthorized{}
}

/* ExecutePipelineByIDUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type ExecutePipelineByIDUsingPOSTUnauthorized struct {
}

func (o *ExecutePipelineByIDUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}/executions][%d] executePipelineByIdUsingPOSTUnauthorized ", 401)
}

func (o *ExecutePipelineByIDUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecutePipelineByIDUsingPOSTForbidden creates a ExecutePipelineByIDUsingPOSTForbidden with default headers values
func NewExecutePipelineByIDUsingPOSTForbidden() *ExecutePipelineByIDUsingPOSTForbidden {
	return &ExecutePipelineByIDUsingPOSTForbidden{}
}

/* ExecutePipelineByIDUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ExecutePipelineByIDUsingPOSTForbidden struct {
}

func (o *ExecutePipelineByIDUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}/executions][%d] executePipelineByIdUsingPOSTForbidden ", 403)
}

func (o *ExecutePipelineByIDUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecutePipelineByIDUsingPOSTNotFound creates a ExecutePipelineByIDUsingPOSTNotFound with default headers values
func NewExecutePipelineByIDUsingPOSTNotFound() *ExecutePipelineByIDUsingPOSTNotFound {
	return &ExecutePipelineByIDUsingPOSTNotFound{}
}

/* ExecutePipelineByIDUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ExecutePipelineByIDUsingPOSTNotFound struct {
	Payload *models.Error
}

func (o *ExecutePipelineByIDUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}/executions][%d] executePipelineByIdUsingPOSTNotFound  %+v", 404, o.Payload)
}
func (o *ExecutePipelineByIDUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ExecutePipelineByIDUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecutePipelineByIDUsingPOSTInternalServerError creates a ExecutePipelineByIDUsingPOSTInternalServerError with default headers values
func NewExecutePipelineByIDUsingPOSTInternalServerError() *ExecutePipelineByIDUsingPOSTInternalServerError {
	return &ExecutePipelineByIDUsingPOSTInternalServerError{}
}

/* ExecutePipelineByIDUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ExecutePipelineByIDUsingPOSTInternalServerError struct {
}

func (o *ExecutePipelineByIDUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}/executions][%d] executePipelineByIdUsingPOSTInternalServerError ", 500)
}

func (o *ExecutePipelineByIDUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
