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

// DeleteDraftByIDUsingDELETEReader is a Reader for the DeleteDraftByIDUsingDELETE structure.
type DeleteDraftByIDUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDraftByIDUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDraftByIDUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDraftByIDUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDraftByIDUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDraftByIDUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDraftByIDUsingDELETEInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDraftByIDUsingDELETEOK creates a DeleteDraftByIDUsingDELETEOK with default headers values
func NewDeleteDraftByIDUsingDELETEOK() *DeleteDraftByIDUsingDELETEOK {
	return &DeleteDraftByIDUsingDELETEOK{}
}

/* DeleteDraftByIDUsingDELETEOK describes a response with status code 200, with default header values.

OK
*/
type DeleteDraftByIDUsingDELETEOK struct {
	Payload models.CustomIntegration
}

func (o *DeleteDraftByIDUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}][%d] deleteDraftByIdUsingDELETEOK  %+v", 200, o.Payload)
}
func (o *DeleteDraftByIDUsingDELETEOK) GetPayload() models.CustomIntegration {
	return o.Payload
}

func (o *DeleteDraftByIDUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalCustomIntegration(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeleteDraftByIDUsingDELETEUnauthorized creates a DeleteDraftByIDUsingDELETEUnauthorized with default headers values
func NewDeleteDraftByIDUsingDELETEUnauthorized() *DeleteDraftByIDUsingDELETEUnauthorized {
	return &DeleteDraftByIDUsingDELETEUnauthorized{}
}

/* DeleteDraftByIDUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type DeleteDraftByIDUsingDELETEUnauthorized struct {
}

func (o *DeleteDraftByIDUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}][%d] deleteDraftByIdUsingDELETEUnauthorized ", 401)
}

func (o *DeleteDraftByIDUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDraftByIDUsingDELETEForbidden creates a DeleteDraftByIDUsingDELETEForbidden with default headers values
func NewDeleteDraftByIDUsingDELETEForbidden() *DeleteDraftByIDUsingDELETEForbidden {
	return &DeleteDraftByIDUsingDELETEForbidden{}
}

/* DeleteDraftByIDUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteDraftByIDUsingDELETEForbidden struct {
}

func (o *DeleteDraftByIDUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}][%d] deleteDraftByIdUsingDELETEForbidden ", 403)
}

func (o *DeleteDraftByIDUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDraftByIDUsingDELETENotFound creates a DeleteDraftByIDUsingDELETENotFound with default headers values
func NewDeleteDraftByIDUsingDELETENotFound() *DeleteDraftByIDUsingDELETENotFound {
	return &DeleteDraftByIDUsingDELETENotFound{}
}

/* DeleteDraftByIDUsingDELETENotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteDraftByIDUsingDELETENotFound struct {
	Payload *models.Error
}

func (o *DeleteDraftByIDUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}][%d] deleteDraftByIdUsingDELETENotFound  %+v", 404, o.Payload)
}
func (o *DeleteDraftByIDUsingDELETENotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteDraftByIDUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDraftByIDUsingDELETEInternalServerError creates a DeleteDraftByIDUsingDELETEInternalServerError with default headers values
func NewDeleteDraftByIDUsingDELETEInternalServerError() *DeleteDraftByIDUsingDELETEInternalServerError {
	return &DeleteDraftByIDUsingDELETEInternalServerError{}
}

/* DeleteDraftByIDUsingDELETEInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DeleteDraftByIDUsingDELETEInternalServerError struct {
}

func (o *DeleteDraftByIDUsingDELETEInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/custom-integrations/{id}][%d] deleteDraftByIdUsingDELETEInternalServerError ", 500)
}

func (o *DeleteDraftByIDUsingDELETEInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
