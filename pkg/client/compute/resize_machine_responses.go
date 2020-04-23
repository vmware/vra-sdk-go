// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// ResizeMachineReader is a Reader for the ResizeMachine structure.
type ResizeMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResizeMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewResizeMachineAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewResizeMachineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResizeMachineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResizeMachineAccepted creates a ResizeMachineAccepted with default headers values
func NewResizeMachineAccepted() *ResizeMachineAccepted {
	return &ResizeMachineAccepted{}
}

/*ResizeMachineAccepted handles this case with default header values.

successful operation
*/
type ResizeMachineAccepted struct {
	Payload *models.RequestTracker
}

func (o *ResizeMachineAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/resize][%d] resizeMachineAccepted  %+v", 202, o.Payload)
}

func (o *ResizeMachineAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *ResizeMachineAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResizeMachineForbidden creates a ResizeMachineForbidden with default headers values
func NewResizeMachineForbidden() *ResizeMachineForbidden {
	return &ResizeMachineForbidden{}
}

/*ResizeMachineForbidden handles this case with default header values.

Forbidden
*/
type ResizeMachineForbidden struct {
}

func (o *ResizeMachineForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/resize][%d] resizeMachineForbidden ", 403)
}

func (o *ResizeMachineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResizeMachineNotFound creates a ResizeMachineNotFound with default headers values
func NewResizeMachineNotFound() *ResizeMachineNotFound {
	return &ResizeMachineNotFound{}
}

/*ResizeMachineNotFound handles this case with default header values.

Not Found
*/
type ResizeMachineNotFound struct {
	Payload *models.Error
}

func (o *ResizeMachineNotFound) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/resize][%d] resizeMachineNotFound  %+v", 404, o.Payload)
}

func (o *ResizeMachineNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ResizeMachineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
