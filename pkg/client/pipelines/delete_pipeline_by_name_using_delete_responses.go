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

// DeletePipelineByNameUsingDELETEReader is a Reader for the DeletePipelineByNameUsingDELETE structure.
type DeletePipelineByNameUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePipelineByNameUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePipelineByNameUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeletePipelineByNameUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePipelineByNameUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePipelineByNameUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePipelineByNameUsingDELETEInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePipelineByNameUsingDELETEOK creates a DeletePipelineByNameUsingDELETEOK with default headers values
func NewDeletePipelineByNameUsingDELETEOK() *DeletePipelineByNameUsingDELETEOK {
	return &DeletePipelineByNameUsingDELETEOK{}
}

/* DeletePipelineByNameUsingDELETEOK describes a response with status code 200, with default header values.

'Success' with the deleted Pipeline
*/
type DeletePipelineByNameUsingDELETEOK struct {
	Payload models.Pipeline
}

func (o *DeletePipelineByNameUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/pipelines/{project}/{name}][%d] deletePipelineByNameUsingDELETEOK  %+v", 200, o.Payload)
}
func (o *DeletePipelineByNameUsingDELETEOK) GetPayload() models.Pipeline {
	return o.Payload
}

func (o *DeletePipelineByNameUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalPipeline(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeletePipelineByNameUsingDELETEUnauthorized creates a DeletePipelineByNameUsingDELETEUnauthorized with default headers values
func NewDeletePipelineByNameUsingDELETEUnauthorized() *DeletePipelineByNameUsingDELETEUnauthorized {
	return &DeletePipelineByNameUsingDELETEUnauthorized{}
}

/* DeletePipelineByNameUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type DeletePipelineByNameUsingDELETEUnauthorized struct {
}

func (o *DeletePipelineByNameUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/pipelines/{project}/{name}][%d] deletePipelineByNameUsingDELETEUnauthorized ", 401)
}

func (o *DeletePipelineByNameUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePipelineByNameUsingDELETEForbidden creates a DeletePipelineByNameUsingDELETEForbidden with default headers values
func NewDeletePipelineByNameUsingDELETEForbidden() *DeletePipelineByNameUsingDELETEForbidden {
	return &DeletePipelineByNameUsingDELETEForbidden{}
}

/* DeletePipelineByNameUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeletePipelineByNameUsingDELETEForbidden struct {
}

func (o *DeletePipelineByNameUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/pipelines/{project}/{name}][%d] deletePipelineByNameUsingDELETEForbidden ", 403)
}

func (o *DeletePipelineByNameUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePipelineByNameUsingDELETENotFound creates a DeletePipelineByNameUsingDELETENotFound with default headers values
func NewDeletePipelineByNameUsingDELETENotFound() *DeletePipelineByNameUsingDELETENotFound {
	return &DeletePipelineByNameUsingDELETENotFound{}
}

/* DeletePipelineByNameUsingDELETENotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeletePipelineByNameUsingDELETENotFound struct {
	Payload *models.Error
}

func (o *DeletePipelineByNameUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/pipelines/{project}/{name}][%d] deletePipelineByNameUsingDELETENotFound  %+v", 404, o.Payload)
}
func (o *DeletePipelineByNameUsingDELETENotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePipelineByNameUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePipelineByNameUsingDELETEInternalServerError creates a DeletePipelineByNameUsingDELETEInternalServerError with default headers values
func NewDeletePipelineByNameUsingDELETEInternalServerError() *DeletePipelineByNameUsingDELETEInternalServerError {
	return &DeletePipelineByNameUsingDELETEInternalServerError{}
}

/* DeletePipelineByNameUsingDELETEInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DeletePipelineByNameUsingDELETEInternalServerError struct {
}

func (o *DeletePipelineByNameUsingDELETEInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/pipelines/{project}/{name}][%d] deletePipelineByNameUsingDELETEInternalServerError ", 500)
}

func (o *DeletePipelineByNameUsingDELETEInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
