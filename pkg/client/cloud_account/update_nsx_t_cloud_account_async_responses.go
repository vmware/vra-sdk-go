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

// UpdateNsxTCloudAccountAsyncReader is a Reader for the UpdateNsxTCloudAccountAsync structure.
type UpdateNsxTCloudAccountAsyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNsxTCloudAccountAsyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewUpdateNsxTCloudAccountAsyncAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateNsxTCloudAccountAsyncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateNsxTCloudAccountAsyncNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNsxTCloudAccountAsyncAccepted creates a UpdateNsxTCloudAccountAsyncAccepted with default headers values
func NewUpdateNsxTCloudAccountAsyncAccepted() *UpdateNsxTCloudAccountAsyncAccepted {
	return &UpdateNsxTCloudAccountAsyncAccepted{}
}

/* UpdateNsxTCloudAccountAsyncAccepted describes a response with status code 202, with default header values.

successful operation
*/
type UpdateNsxTCloudAccountAsyncAccepted struct {
	Payload *models.RequestTracker
}

func (o *UpdateNsxTCloudAccountAsyncAccepted) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-nsx-t/{id}][%d] updateNsxTCloudAccountAsyncAccepted  %+v", 202, o.Payload)
}
func (o *UpdateNsxTCloudAccountAsyncAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *UpdateNsxTCloudAccountAsyncAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNsxTCloudAccountAsyncForbidden creates a UpdateNsxTCloudAccountAsyncForbidden with default headers values
func NewUpdateNsxTCloudAccountAsyncForbidden() *UpdateNsxTCloudAccountAsyncForbidden {
	return &UpdateNsxTCloudAccountAsyncForbidden{}
}

/* UpdateNsxTCloudAccountAsyncForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateNsxTCloudAccountAsyncForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *UpdateNsxTCloudAccountAsyncForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-nsx-t/{id}][%d] updateNsxTCloudAccountAsyncForbidden  %+v", 403, o.Payload)
}
func (o *UpdateNsxTCloudAccountAsyncForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *UpdateNsxTCloudAccountAsyncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNsxTCloudAccountAsyncNotFound creates a UpdateNsxTCloudAccountAsyncNotFound with default headers values
func NewUpdateNsxTCloudAccountAsyncNotFound() *UpdateNsxTCloudAccountAsyncNotFound {
	return &UpdateNsxTCloudAccountAsyncNotFound{}
}

/* UpdateNsxTCloudAccountAsyncNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateNsxTCloudAccountAsyncNotFound struct {
	Payload *models.Error
}

func (o *UpdateNsxTCloudAccountAsyncNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-nsx-t/{id}][%d] updateNsxTCloudAccountAsyncNotFound  %+v", 404, o.Payload)
}
func (o *UpdateNsxTCloudAccountAsyncNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateNsxTCloudAccountAsyncNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
