// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateVMCCloudAccountAsyncReader is a Reader for the UpdateVMCCloudAccountAsync structure.
type UpdateVMCCloudAccountAsyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVMCCloudAccountAsyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewUpdateVMCCloudAccountAsyncAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateVMCCloudAccountAsyncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVMCCloudAccountAsyncNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVMCCloudAccountAsyncAccepted creates a UpdateVMCCloudAccountAsyncAccepted with default headers values
func NewUpdateVMCCloudAccountAsyncAccepted() *UpdateVMCCloudAccountAsyncAccepted {
	return &UpdateVMCCloudAccountAsyncAccepted{}
}

/* UpdateVMCCloudAccountAsyncAccepted describes a response with status code 202, with default header values.

successful operation
*/
type UpdateVMCCloudAccountAsyncAccepted struct {
	Payload *models.RequestTracker
}

func (o *UpdateVMCCloudAccountAsyncAccepted) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vmc/{id}][%d] updateVmCCloudAccountAsyncAccepted  %+v", 202, o.Payload)
}
func (o *UpdateVMCCloudAccountAsyncAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *UpdateVMCCloudAccountAsyncAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMCCloudAccountAsyncForbidden creates a UpdateVMCCloudAccountAsyncForbidden with default headers values
func NewUpdateVMCCloudAccountAsyncForbidden() *UpdateVMCCloudAccountAsyncForbidden {
	return &UpdateVMCCloudAccountAsyncForbidden{}
}

/* UpdateVMCCloudAccountAsyncForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateVMCCloudAccountAsyncForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *UpdateVMCCloudAccountAsyncForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vmc/{id}][%d] updateVmCCloudAccountAsyncForbidden  %+v", 403, o.Payload)
}
func (o *UpdateVMCCloudAccountAsyncForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *UpdateVMCCloudAccountAsyncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMCCloudAccountAsyncNotFound creates a UpdateVMCCloudAccountAsyncNotFound with default headers values
func NewUpdateVMCCloudAccountAsyncNotFound() *UpdateVMCCloudAccountAsyncNotFound {
	return &UpdateVMCCloudAccountAsyncNotFound{}
}

/* UpdateVMCCloudAccountAsyncNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateVMCCloudAccountAsyncNotFound struct {
	Payload *models.Error
}

func (o *UpdateVMCCloudAccountAsyncNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vmc/{id}][%d] updateVmCCloudAccountAsyncNotFound  %+v", 404, o.Payload)
}
func (o *UpdateVMCCloudAccountAsyncNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateVMCCloudAccountAsyncNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
