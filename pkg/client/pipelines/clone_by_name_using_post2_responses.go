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

// CloneByNameUsingPOST2Reader is a Reader for the CloneByNameUsingPOST2 structure.
type CloneByNameUsingPOST2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneByNameUsingPOST2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloneByNameUsingPOST2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCloneByNameUsingPOST2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloneByNameUsingPOST2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloneByNameUsingPOST2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloneByNameUsingPOST2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloneByNameUsingPOST2OK creates a CloneByNameUsingPOST2OK with default headers values
func NewCloneByNameUsingPOST2OK() *CloneByNameUsingPOST2OK {
	return &CloneByNameUsingPOST2OK{}
}

/* CloneByNameUsingPOST2OK describes a response with status code 200, with default header values.

'Success' with the cloned Pipeline
*/
type CloneByNameUsingPOST2OK struct {
	Payload models.Pipeline
}

func (o *CloneByNameUsingPOST2OK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}][%d] cloneByNameUsingPOST2OK  %+v", 200, o.Payload)
}
func (o *CloneByNameUsingPOST2OK) GetPayload() models.Pipeline {
	return o.Payload
}

func (o *CloneByNameUsingPOST2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalPipeline(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewCloneByNameUsingPOST2Unauthorized creates a CloneByNameUsingPOST2Unauthorized with default headers values
func NewCloneByNameUsingPOST2Unauthorized() *CloneByNameUsingPOST2Unauthorized {
	return &CloneByNameUsingPOST2Unauthorized{}
}

/* CloneByNameUsingPOST2Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type CloneByNameUsingPOST2Unauthorized struct {
}

func (o *CloneByNameUsingPOST2Unauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}][%d] cloneByNameUsingPOST2Unauthorized ", 401)
}

func (o *CloneByNameUsingPOST2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloneByNameUsingPOST2Forbidden creates a CloneByNameUsingPOST2Forbidden with default headers values
func NewCloneByNameUsingPOST2Forbidden() *CloneByNameUsingPOST2Forbidden {
	return &CloneByNameUsingPOST2Forbidden{}
}

/* CloneByNameUsingPOST2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloneByNameUsingPOST2Forbidden struct {
}

func (o *CloneByNameUsingPOST2Forbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}][%d] cloneByNameUsingPOST2Forbidden ", 403)
}

func (o *CloneByNameUsingPOST2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloneByNameUsingPOST2NotFound creates a CloneByNameUsingPOST2NotFound with default headers values
func NewCloneByNameUsingPOST2NotFound() *CloneByNameUsingPOST2NotFound {
	return &CloneByNameUsingPOST2NotFound{}
}

/* CloneByNameUsingPOST2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type CloneByNameUsingPOST2NotFound struct {
	Payload *models.Error
}

func (o *CloneByNameUsingPOST2NotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}][%d] cloneByNameUsingPOST2NotFound  %+v", 404, o.Payload)
}
func (o *CloneByNameUsingPOST2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CloneByNameUsingPOST2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneByNameUsingPOST2InternalServerError creates a CloneByNameUsingPOST2InternalServerError with default headers values
func NewCloneByNameUsingPOST2InternalServerError() *CloneByNameUsingPOST2InternalServerError {
	return &CloneByNameUsingPOST2InternalServerError{}
}

/* CloneByNameUsingPOST2InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CloneByNameUsingPOST2InternalServerError struct {
}

func (o *CloneByNameUsingPOST2InternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{project}/{name}][%d] cloneByNameUsingPOST2InternalServerError ", 500)
}

func (o *CloneByNameUsingPOST2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
