// Code generated by go-swagger; DO NOT EDIT.

package executions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteUsingDELETE10Reader is a Reader for the DeleteUsingDELETE10 structure.
type DeleteUsingDELETE10Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsingDELETE10Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUsingDELETE10OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUsingDELETE10Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUsingDELETE10Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUsingDELETE10NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUsingDELETE10InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUsingDELETE10OK creates a DeleteUsingDELETE10OK with default headers values
func NewDeleteUsingDELETE10OK() *DeleteUsingDELETE10OK {
	return &DeleteUsingDELETE10OK{}
}

/*DeleteUsingDELETE10OK handles this case with default header values.

'Success' with the deleted User Operation
*/
type DeleteUsingDELETE10OK struct {
	Payload models.UserOperation
}

func (o *DeleteUsingDELETE10OK) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/user-operations/{id}][%d] deleteUsingDELETE10OK  %+v", 200, o.Payload)
}

func (o *DeleteUsingDELETE10OK) GetPayload() models.UserOperation {
	return o.Payload
}

func (o *DeleteUsingDELETE10OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalUserOperation(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeleteUsingDELETE10Unauthorized creates a DeleteUsingDELETE10Unauthorized with default headers values
func NewDeleteUsingDELETE10Unauthorized() *DeleteUsingDELETE10Unauthorized {
	return &DeleteUsingDELETE10Unauthorized{}
}

/*DeleteUsingDELETE10Unauthorized handles this case with default header values.

Unauthorized Request
*/
type DeleteUsingDELETE10Unauthorized struct {
}

func (o *DeleteUsingDELETE10Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/user-operations/{id}][%d] deleteUsingDELETE10Unauthorized ", 401)
}

func (o *DeleteUsingDELETE10Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETE10Forbidden creates a DeleteUsingDELETE10Forbidden with default headers values
func NewDeleteUsingDELETE10Forbidden() *DeleteUsingDELETE10Forbidden {
	return &DeleteUsingDELETE10Forbidden{}
}

/*DeleteUsingDELETE10Forbidden handles this case with default header values.

Forbidden
*/
type DeleteUsingDELETE10Forbidden struct {
}

func (o *DeleteUsingDELETE10Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/user-operations/{id}][%d] deleteUsingDELETE10Forbidden ", 403)
}

func (o *DeleteUsingDELETE10Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETE10NotFound creates a DeleteUsingDELETE10NotFound with default headers values
func NewDeleteUsingDELETE10NotFound() *DeleteUsingDELETE10NotFound {
	return &DeleteUsingDELETE10NotFound{}
}

/*DeleteUsingDELETE10NotFound handles this case with default header values.

Not Found
*/
type DeleteUsingDELETE10NotFound struct {
}

func (o *DeleteUsingDELETE10NotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/user-operations/{id}][%d] deleteUsingDELETE10NotFound ", 404)
}

func (o *DeleteUsingDELETE10NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETE10InternalServerError creates a DeleteUsingDELETE10InternalServerError with default headers values
func NewDeleteUsingDELETE10InternalServerError() *DeleteUsingDELETE10InternalServerError {
	return &DeleteUsingDELETE10InternalServerError{}
}

/*DeleteUsingDELETE10InternalServerError handles this case with default header values.

Server Error
*/
type DeleteUsingDELETE10InternalServerError struct {
}

func (o *DeleteUsingDELETE10InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/user-operations/{id}][%d] deleteUsingDELETE10InternalServerError ", 500)
}

func (o *DeleteUsingDELETE10InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
