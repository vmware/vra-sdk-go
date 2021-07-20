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

// DeleteUsingDELETE4Reader is a Reader for the DeleteUsingDELETE4 structure.
type DeleteUsingDELETE4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsingDELETE4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUsingDELETE4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUsingDELETE4Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUsingDELETE4Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUsingDELETE4NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUsingDELETE4InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUsingDELETE4OK creates a DeleteUsingDELETE4OK with default headers values
func NewDeleteUsingDELETE4OK() *DeleteUsingDELETE4OK {
	return &DeleteUsingDELETE4OK{}
}

/* DeleteUsingDELETE4OK describes a response with status code 200, with default header values.

'Success' with Delete a Gerrit Event
*/
type DeleteUsingDELETE4OK struct {
	Payload models.GerritEvent
}

func (o *DeleteUsingDELETE4OK) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-events/{id}][%d] deleteUsingDELETE4OK  %+v", 200, o.Payload)
}
func (o *DeleteUsingDELETE4OK) GetPayload() models.GerritEvent {
	return o.Payload
}

func (o *DeleteUsingDELETE4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritEvent(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeleteUsingDELETE4Unauthorized creates a DeleteUsingDELETE4Unauthorized with default headers values
func NewDeleteUsingDELETE4Unauthorized() *DeleteUsingDELETE4Unauthorized {
	return &DeleteUsingDELETE4Unauthorized{}
}

/* DeleteUsingDELETE4Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type DeleteUsingDELETE4Unauthorized struct {
}

func (o *DeleteUsingDELETE4Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-events/{id}][%d] deleteUsingDELETE4Unauthorized ", 401)
}

func (o *DeleteUsingDELETE4Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETE4Forbidden creates a DeleteUsingDELETE4Forbidden with default headers values
func NewDeleteUsingDELETE4Forbidden() *DeleteUsingDELETE4Forbidden {
	return &DeleteUsingDELETE4Forbidden{}
}

/* DeleteUsingDELETE4Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteUsingDELETE4Forbidden struct {
}

func (o *DeleteUsingDELETE4Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-events/{id}][%d] deleteUsingDELETE4Forbidden ", 403)
}

func (o *DeleteUsingDELETE4Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETE4NotFound creates a DeleteUsingDELETE4NotFound with default headers values
func NewDeleteUsingDELETE4NotFound() *DeleteUsingDELETE4NotFound {
	return &DeleteUsingDELETE4NotFound{}
}

/* DeleteUsingDELETE4NotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteUsingDELETE4NotFound struct {
	Payload *models.Error
}

func (o *DeleteUsingDELETE4NotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-events/{id}][%d] deleteUsingDELETE4NotFound  %+v", 404, o.Payload)
}
func (o *DeleteUsingDELETE4NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUsingDELETE4NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUsingDELETE4InternalServerError creates a DeleteUsingDELETE4InternalServerError with default headers values
func NewDeleteUsingDELETE4InternalServerError() *DeleteUsingDELETE4InternalServerError {
	return &DeleteUsingDELETE4InternalServerError{}
}

/* DeleteUsingDELETE4InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DeleteUsingDELETE4InternalServerError struct {
}

func (o *DeleteUsingDELETE4InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-events/{id}][%d] deleteUsingDELETE4InternalServerError ", 500)
}

func (o *DeleteUsingDELETE4InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
