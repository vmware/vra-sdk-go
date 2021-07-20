// Code generated by go-swagger; DO NOT EDIT.

package executions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteUsingDELETE3Reader is a Reader for the DeleteUsingDELETE3 structure.
type DeleteUsingDELETE3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsingDELETE3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUsingDELETE3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUsingDELETE3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUsingDELETE3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUsingDELETE3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUsingDELETE3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUsingDELETE3OK creates a DeleteUsingDELETE3OK with default headers values
func NewDeleteUsingDELETE3OK() *DeleteUsingDELETE3OK {
	return &DeleteUsingDELETE3OK{}
}

/* DeleteUsingDELETE3OK describes a response with status code 200, with default header values.

'Success' with execution delete
*/
type DeleteUsingDELETE3OK struct {
	Payload models.Execution
}

func (o *DeleteUsingDELETE3OK) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/executions/{id}][%d] deleteUsingDELETE3OK  %+v", 200, o.Payload)
}
func (o *DeleteUsingDELETE3OK) GetPayload() models.Execution {
	return o.Payload
}

func (o *DeleteUsingDELETE3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecution(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeleteUsingDELETE3Unauthorized creates a DeleteUsingDELETE3Unauthorized with default headers values
func NewDeleteUsingDELETE3Unauthorized() *DeleteUsingDELETE3Unauthorized {
	return &DeleteUsingDELETE3Unauthorized{}
}

/* DeleteUsingDELETE3Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type DeleteUsingDELETE3Unauthorized struct {
}

func (o *DeleteUsingDELETE3Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/executions/{id}][%d] deleteUsingDELETE3Unauthorized ", 401)
}

func (o *DeleteUsingDELETE3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETE3Forbidden creates a DeleteUsingDELETE3Forbidden with default headers values
func NewDeleteUsingDELETE3Forbidden() *DeleteUsingDELETE3Forbidden {
	return &DeleteUsingDELETE3Forbidden{}
}

/* DeleteUsingDELETE3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteUsingDELETE3Forbidden struct {
}

func (o *DeleteUsingDELETE3Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/executions/{id}][%d] deleteUsingDELETE3Forbidden ", 403)
}

func (o *DeleteUsingDELETE3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETE3NotFound creates a DeleteUsingDELETE3NotFound with default headers values
func NewDeleteUsingDELETE3NotFound() *DeleteUsingDELETE3NotFound {
	return &DeleteUsingDELETE3NotFound{}
}

/* DeleteUsingDELETE3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteUsingDELETE3NotFound struct {
	Payload *models.Error
}

func (o *DeleteUsingDELETE3NotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/executions/{id}][%d] deleteUsingDELETE3NotFound  %+v", 404, o.Payload)
}
func (o *DeleteUsingDELETE3NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteUsingDELETE3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUsingDELETE3InternalServerError creates a DeleteUsingDELETE3InternalServerError with default headers values
func NewDeleteUsingDELETE3InternalServerError() *DeleteUsingDELETE3InternalServerError {
	return &DeleteUsingDELETE3InternalServerError{}
}

/* DeleteUsingDELETE3InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DeleteUsingDELETE3InternalServerError struct {
}

func (o *DeleteUsingDELETE3InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/executions/{id}][%d] deleteUsingDELETE3InternalServerError ", 500)
}

func (o *DeleteUsingDELETE3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
