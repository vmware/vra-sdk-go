// Code generated by go-swagger; DO NOT EDIT.

package disk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteMachineDiskReader is a Reader for the DeleteMachineDisk structure.
type DeleteMachineDiskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMachineDiskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteMachineDiskAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteMachineDiskNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteMachineDiskForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteMachineDiskAccepted creates a DeleteMachineDiskAccepted with default headers values
func NewDeleteMachineDiskAccepted() *DeleteMachineDiskAccepted {
	return &DeleteMachineDiskAccepted{}
}

/* DeleteMachineDiskAccepted describes a response with status code 202, with default header values.

successful operation
*/
type DeleteMachineDiskAccepted struct {
	Payload *models.RequestTracker
}

func (o *DeleteMachineDiskAccepted) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/machines/{id}/disks/{diskId}][%d] deleteMachineDiskAccepted  %+v", 202, o.Payload)
}
func (o *DeleteMachineDiskAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *DeleteMachineDiskAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMachineDiskNoContent creates a DeleteMachineDiskNoContent with default headers values
func NewDeleteMachineDiskNoContent() *DeleteMachineDiskNoContent {
	return &DeleteMachineDiskNoContent{}
}

/* DeleteMachineDiskNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteMachineDiskNoContent struct {
}

func (o *DeleteMachineDiskNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/machines/{id}/disks/{diskId}][%d] deleteMachineDiskNoContent ", 204)
}

func (o *DeleteMachineDiskNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMachineDiskForbidden creates a DeleteMachineDiskForbidden with default headers values
func NewDeleteMachineDiskForbidden() *DeleteMachineDiskForbidden {
	return &DeleteMachineDiskForbidden{}
}

/* DeleteMachineDiskForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteMachineDiskForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *DeleteMachineDiskForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/machines/{id}/disks/{diskId}][%d] deleteMachineDiskForbidden  %+v", 403, o.Payload)
}
func (o *DeleteMachineDiskForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *DeleteMachineDiskForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
