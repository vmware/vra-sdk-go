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

// UpdateVSphereCloudAccountAsyncReader is a Reader for the UpdateVSphereCloudAccountAsync structure.
type UpdateVSphereCloudAccountAsyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVSphereCloudAccountAsyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewUpdateVSphereCloudAccountAsyncAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateVSphereCloudAccountAsyncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVSphereCloudAccountAsyncNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVSphereCloudAccountAsyncAccepted creates a UpdateVSphereCloudAccountAsyncAccepted with default headers values
func NewUpdateVSphereCloudAccountAsyncAccepted() *UpdateVSphereCloudAccountAsyncAccepted {
	return &UpdateVSphereCloudAccountAsyncAccepted{}
}

/*
UpdateVSphereCloudAccountAsyncAccepted describes a response with status code 202, with default header values.

successful operation
*/
type UpdateVSphereCloudAccountAsyncAccepted struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this update v sphere cloud account async accepted response has a 2xx status code
func (o *UpdateVSphereCloudAccountAsyncAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update v sphere cloud account async accepted response has a 3xx status code
func (o *UpdateVSphereCloudAccountAsyncAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update v sphere cloud account async accepted response has a 4xx status code
func (o *UpdateVSphereCloudAccountAsyncAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this update v sphere cloud account async accepted response has a 5xx status code
func (o *UpdateVSphereCloudAccountAsyncAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this update v sphere cloud account async accepted response a status code equal to that given
func (o *UpdateVSphereCloudAccountAsyncAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *UpdateVSphereCloudAccountAsyncAccepted) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vsphere/{id}][%d] updateVSphereCloudAccountAsyncAccepted  %+v", 202, o.Payload)
}

func (o *UpdateVSphereCloudAccountAsyncAccepted) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vsphere/{id}][%d] updateVSphereCloudAccountAsyncAccepted  %+v", 202, o.Payload)
}

func (o *UpdateVSphereCloudAccountAsyncAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *UpdateVSphereCloudAccountAsyncAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVSphereCloudAccountAsyncForbidden creates a UpdateVSphereCloudAccountAsyncForbidden with default headers values
func NewUpdateVSphereCloudAccountAsyncForbidden() *UpdateVSphereCloudAccountAsyncForbidden {
	return &UpdateVSphereCloudAccountAsyncForbidden{}
}

/*
UpdateVSphereCloudAccountAsyncForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateVSphereCloudAccountAsyncForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this update v sphere cloud account async forbidden response has a 2xx status code
func (o *UpdateVSphereCloudAccountAsyncForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update v sphere cloud account async forbidden response has a 3xx status code
func (o *UpdateVSphereCloudAccountAsyncForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update v sphere cloud account async forbidden response has a 4xx status code
func (o *UpdateVSphereCloudAccountAsyncForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update v sphere cloud account async forbidden response has a 5xx status code
func (o *UpdateVSphereCloudAccountAsyncForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update v sphere cloud account async forbidden response a status code equal to that given
func (o *UpdateVSphereCloudAccountAsyncForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateVSphereCloudAccountAsyncForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vsphere/{id}][%d] updateVSphereCloudAccountAsyncForbidden  %+v", 403, o.Payload)
}

func (o *UpdateVSphereCloudAccountAsyncForbidden) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vsphere/{id}][%d] updateVSphereCloudAccountAsyncForbidden  %+v", 403, o.Payload)
}

func (o *UpdateVSphereCloudAccountAsyncForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *UpdateVSphereCloudAccountAsyncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVSphereCloudAccountAsyncNotFound creates a UpdateVSphereCloudAccountAsyncNotFound with default headers values
func NewUpdateVSphereCloudAccountAsyncNotFound() *UpdateVSphereCloudAccountAsyncNotFound {
	return &UpdateVSphereCloudAccountAsyncNotFound{}
}

/*
UpdateVSphereCloudAccountAsyncNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateVSphereCloudAccountAsyncNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update v sphere cloud account async not found response has a 2xx status code
func (o *UpdateVSphereCloudAccountAsyncNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update v sphere cloud account async not found response has a 3xx status code
func (o *UpdateVSphereCloudAccountAsyncNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update v sphere cloud account async not found response has a 4xx status code
func (o *UpdateVSphereCloudAccountAsyncNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update v sphere cloud account async not found response has a 5xx status code
func (o *UpdateVSphereCloudAccountAsyncNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update v sphere cloud account async not found response a status code equal to that given
func (o *UpdateVSphereCloudAccountAsyncNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateVSphereCloudAccountAsyncNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vsphere/{id}][%d] updateVSphereCloudAccountAsyncNotFound  %+v", 404, o.Payload)
}

func (o *UpdateVSphereCloudAccountAsyncNotFound) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vsphere/{id}][%d] updateVSphereCloudAccountAsyncNotFound  %+v", 404, o.Payload)
}

func (o *UpdateVSphereCloudAccountAsyncNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateVSphereCloudAccountAsyncNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
