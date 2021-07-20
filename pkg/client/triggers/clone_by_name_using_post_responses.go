// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CloneByNameUsingPOSTReader is a Reader for the CloneByNameUsingPOST structure.
type CloneByNameUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneByNameUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloneByNameUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCloneByNameUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloneByNameUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloneByNameUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloneByNameUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloneByNameUsingPOSTOK creates a CloneByNameUsingPOSTOK with default headers values
func NewCloneByNameUsingPOSTOK() *CloneByNameUsingPOSTOK {
	return &CloneByNameUsingPOSTOK{}
}

/* CloneByNameUsingPOSTOK describes a response with status code 200, with default header values.

'Success' with the cloned Pipeline
*/
type CloneByNameUsingPOSTOK struct {
	Payload models.GerritListener
}

func (o *CloneByNameUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneByNameUsingPOSTOK  %+v", 200, o.Payload)
}
func (o *CloneByNameUsingPOSTOK) GetPayload() models.GerritListener {
	return o.Payload
}

func (o *CloneByNameUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListener(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewCloneByNameUsingPOSTUnauthorized creates a CloneByNameUsingPOSTUnauthorized with default headers values
func NewCloneByNameUsingPOSTUnauthorized() *CloneByNameUsingPOSTUnauthorized {
	return &CloneByNameUsingPOSTUnauthorized{}
}

/* CloneByNameUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type CloneByNameUsingPOSTUnauthorized struct {
}

func (o *CloneByNameUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneByNameUsingPOSTUnauthorized ", 401)
}

func (o *CloneByNameUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloneByNameUsingPOSTForbidden creates a CloneByNameUsingPOSTForbidden with default headers values
func NewCloneByNameUsingPOSTForbidden() *CloneByNameUsingPOSTForbidden {
	return &CloneByNameUsingPOSTForbidden{}
}

/* CloneByNameUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloneByNameUsingPOSTForbidden struct {
}

func (o *CloneByNameUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneByNameUsingPOSTForbidden ", 403)
}

func (o *CloneByNameUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloneByNameUsingPOSTNotFound creates a CloneByNameUsingPOSTNotFound with default headers values
func NewCloneByNameUsingPOSTNotFound() *CloneByNameUsingPOSTNotFound {
	return &CloneByNameUsingPOSTNotFound{}
}

/* CloneByNameUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CloneByNameUsingPOSTNotFound struct {
	Payload *models.Error
}

func (o *CloneByNameUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneByNameUsingPOSTNotFound  %+v", 404, o.Payload)
}
func (o *CloneByNameUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CloneByNameUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneByNameUsingPOSTInternalServerError creates a CloneByNameUsingPOSTInternalServerError with default headers values
func NewCloneByNameUsingPOSTInternalServerError() *CloneByNameUsingPOSTInternalServerError {
	return &CloneByNameUsingPOSTInternalServerError{}
}

/* CloneByNameUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CloneByNameUsingPOSTInternalServerError struct {
}

func (o *CloneByNameUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners/{project}/{name}][%d] cloneByNameUsingPOSTInternalServerError ", 500)
}

func (o *CloneByNameUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
