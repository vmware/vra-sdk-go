// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// ActOnPipelineUsingPOSTReader is a Reader for the ActOnPipelineUsingPOST structure.
type ActOnPipelineUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActOnPipelineUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActOnPipelineUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewActOnPipelineUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewActOnPipelineUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewActOnPipelineUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewActOnPipelineUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActOnPipelineUsingPOSTOK creates a ActOnPipelineUsingPOSTOK with default headers values
func NewActOnPipelineUsingPOSTOK() *ActOnPipelineUsingPOSTOK {
	return &ActOnPipelineUsingPOSTOK{}
}

/*ActOnPipelineUsingPOSTOK handles this case with default header values.

'Success' with the cloned Pipeline
*/
type ActOnPipelineUsingPOSTOK struct {
	Payload models.Pipeline
}

func (o *ActOnPipelineUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}][%d] actOnPipelineUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ActOnPipelineUsingPOSTOK) GetPayload() models.Pipeline {
	return o.Payload
}

func (o *ActOnPipelineUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalPipeline(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewActOnPipelineUsingPOSTUnauthorized creates a ActOnPipelineUsingPOSTUnauthorized with default headers values
func NewActOnPipelineUsingPOSTUnauthorized() *ActOnPipelineUsingPOSTUnauthorized {
	return &ActOnPipelineUsingPOSTUnauthorized{}
}

/*ActOnPipelineUsingPOSTUnauthorized handles this case with default header values.

Unauthorized Request
*/
type ActOnPipelineUsingPOSTUnauthorized struct {
}

func (o *ActOnPipelineUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}][%d] actOnPipelineUsingPOSTUnauthorized ", 401)
}

func (o *ActOnPipelineUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActOnPipelineUsingPOSTForbidden creates a ActOnPipelineUsingPOSTForbidden with default headers values
func NewActOnPipelineUsingPOSTForbidden() *ActOnPipelineUsingPOSTForbidden {
	return &ActOnPipelineUsingPOSTForbidden{}
}

/*ActOnPipelineUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type ActOnPipelineUsingPOSTForbidden struct {
}

func (o *ActOnPipelineUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}][%d] actOnPipelineUsingPOSTForbidden ", 403)
}

func (o *ActOnPipelineUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActOnPipelineUsingPOSTNotFound creates a ActOnPipelineUsingPOSTNotFound with default headers values
func NewActOnPipelineUsingPOSTNotFound() *ActOnPipelineUsingPOSTNotFound {
	return &ActOnPipelineUsingPOSTNotFound{}
}

/*ActOnPipelineUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type ActOnPipelineUsingPOSTNotFound struct {
}

func (o *ActOnPipelineUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}][%d] actOnPipelineUsingPOSTNotFound ", 404)
}

func (o *ActOnPipelineUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActOnPipelineUsingPOSTInternalServerError creates a ActOnPipelineUsingPOSTInternalServerError with default headers values
func NewActOnPipelineUsingPOSTInternalServerError() *ActOnPipelineUsingPOSTInternalServerError {
	return &ActOnPipelineUsingPOSTInternalServerError{}
}

/*ActOnPipelineUsingPOSTInternalServerError handles this case with default header values.

Server Error
*/
type ActOnPipelineUsingPOSTInternalServerError struct {
}

func (o *ActOnPipelineUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}][%d] actOnPipelineUsingPOSTInternalServerError ", 500)
}

func (o *ActOnPipelineUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
