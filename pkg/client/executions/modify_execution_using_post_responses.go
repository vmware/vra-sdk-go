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

// ModifyExecutionUsingPOSTReader is a Reader for the ModifyExecutionUsingPOST structure.
type ModifyExecutionUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyExecutionUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModifyExecutionUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewModifyExecutionUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewModifyExecutionUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewModifyExecutionUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewModifyExecutionUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyExecutionUsingPOSTOK creates a ModifyExecutionUsingPOSTOK with default headers values
func NewModifyExecutionUsingPOSTOK() *ModifyExecutionUsingPOSTOK {
	return &ModifyExecutionUsingPOSTOK{}
}

/*ModifyExecutionUsingPOSTOK handles this case with default header values.

'Success' with Execution
*/
type ModifyExecutionUsingPOSTOK struct {
	Payload models.Execution
}

func (o *ModifyExecutionUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/executions/{id}][%d] modifyExecutionUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ModifyExecutionUsingPOSTOK) GetPayload() models.Execution {
	return o.Payload
}

func (o *ModifyExecutionUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecution(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewModifyExecutionUsingPOSTUnauthorized creates a ModifyExecutionUsingPOSTUnauthorized with default headers values
func NewModifyExecutionUsingPOSTUnauthorized() *ModifyExecutionUsingPOSTUnauthorized {
	return &ModifyExecutionUsingPOSTUnauthorized{}
}

/*ModifyExecutionUsingPOSTUnauthorized handles this case with default header values.

Unauthorized Request
*/
type ModifyExecutionUsingPOSTUnauthorized struct {
}

func (o *ModifyExecutionUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/executions/{id}][%d] modifyExecutionUsingPOSTUnauthorized ", 401)
}

func (o *ModifyExecutionUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyExecutionUsingPOSTForbidden creates a ModifyExecutionUsingPOSTForbidden with default headers values
func NewModifyExecutionUsingPOSTForbidden() *ModifyExecutionUsingPOSTForbidden {
	return &ModifyExecutionUsingPOSTForbidden{}
}

/*ModifyExecutionUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type ModifyExecutionUsingPOSTForbidden struct {
}

func (o *ModifyExecutionUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/executions/{id}][%d] modifyExecutionUsingPOSTForbidden ", 403)
}

func (o *ModifyExecutionUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyExecutionUsingPOSTNotFound creates a ModifyExecutionUsingPOSTNotFound with default headers values
func NewModifyExecutionUsingPOSTNotFound() *ModifyExecutionUsingPOSTNotFound {
	return &ModifyExecutionUsingPOSTNotFound{}
}

/*ModifyExecutionUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type ModifyExecutionUsingPOSTNotFound struct {
}

func (o *ModifyExecutionUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/executions/{id}][%d] modifyExecutionUsingPOSTNotFound ", 404)
}

func (o *ModifyExecutionUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyExecutionUsingPOSTInternalServerError creates a ModifyExecutionUsingPOSTInternalServerError with default headers values
func NewModifyExecutionUsingPOSTInternalServerError() *ModifyExecutionUsingPOSTInternalServerError {
	return &ModifyExecutionUsingPOSTInternalServerError{}
}

/*ModifyExecutionUsingPOSTInternalServerError handles this case with default header values.

Server Error
*/
type ModifyExecutionUsingPOSTInternalServerError struct {
}

func (o *ModifyExecutionUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/executions/{id}][%d] modifyExecutionUsingPOSTInternalServerError ", 500)
}

func (o *ModifyExecutionUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
