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

// UpdateVcfCloudAccountAsyncReader is a Reader for the UpdateVcfCloudAccountAsync structure.
type UpdateVcfCloudAccountAsyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVcfCloudAccountAsyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewUpdateVcfCloudAccountAsyncAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateVcfCloudAccountAsyncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVcfCloudAccountAsyncNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVcfCloudAccountAsyncAccepted creates a UpdateVcfCloudAccountAsyncAccepted with default headers values
func NewUpdateVcfCloudAccountAsyncAccepted() *UpdateVcfCloudAccountAsyncAccepted {
	return &UpdateVcfCloudAccountAsyncAccepted{}
}

/*
UpdateVcfCloudAccountAsyncAccepted describes a response with status code 202, with default header values.

successful operation
*/
type UpdateVcfCloudAccountAsyncAccepted struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this update vcf cloud account async accepted response has a 2xx status code
func (o *UpdateVcfCloudAccountAsyncAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update vcf cloud account async accepted response has a 3xx status code
func (o *UpdateVcfCloudAccountAsyncAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update vcf cloud account async accepted response has a 4xx status code
func (o *UpdateVcfCloudAccountAsyncAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this update vcf cloud account async accepted response has a 5xx status code
func (o *UpdateVcfCloudAccountAsyncAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this update vcf cloud account async accepted response a status code equal to that given
func (o *UpdateVcfCloudAccountAsyncAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *UpdateVcfCloudAccountAsyncAccepted) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vcf/{id}][%d] updateVcfCloudAccountAsyncAccepted  %+v", 202, o.Payload)
}

func (o *UpdateVcfCloudAccountAsyncAccepted) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vcf/{id}][%d] updateVcfCloudAccountAsyncAccepted  %+v", 202, o.Payload)
}

func (o *UpdateVcfCloudAccountAsyncAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *UpdateVcfCloudAccountAsyncAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVcfCloudAccountAsyncForbidden creates a UpdateVcfCloudAccountAsyncForbidden with default headers values
func NewUpdateVcfCloudAccountAsyncForbidden() *UpdateVcfCloudAccountAsyncForbidden {
	return &UpdateVcfCloudAccountAsyncForbidden{}
}

/*
UpdateVcfCloudAccountAsyncForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateVcfCloudAccountAsyncForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this update vcf cloud account async forbidden response has a 2xx status code
func (o *UpdateVcfCloudAccountAsyncForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update vcf cloud account async forbidden response has a 3xx status code
func (o *UpdateVcfCloudAccountAsyncForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update vcf cloud account async forbidden response has a 4xx status code
func (o *UpdateVcfCloudAccountAsyncForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update vcf cloud account async forbidden response has a 5xx status code
func (o *UpdateVcfCloudAccountAsyncForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update vcf cloud account async forbidden response a status code equal to that given
func (o *UpdateVcfCloudAccountAsyncForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateVcfCloudAccountAsyncForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vcf/{id}][%d] updateVcfCloudAccountAsyncForbidden  %+v", 403, o.Payload)
}

func (o *UpdateVcfCloudAccountAsyncForbidden) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vcf/{id}][%d] updateVcfCloudAccountAsyncForbidden  %+v", 403, o.Payload)
}

func (o *UpdateVcfCloudAccountAsyncForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *UpdateVcfCloudAccountAsyncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVcfCloudAccountAsyncNotFound creates a UpdateVcfCloudAccountAsyncNotFound with default headers values
func NewUpdateVcfCloudAccountAsyncNotFound() *UpdateVcfCloudAccountAsyncNotFound {
	return &UpdateVcfCloudAccountAsyncNotFound{}
}

/*
UpdateVcfCloudAccountAsyncNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateVcfCloudAccountAsyncNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update vcf cloud account async not found response has a 2xx status code
func (o *UpdateVcfCloudAccountAsyncNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update vcf cloud account async not found response has a 3xx status code
func (o *UpdateVcfCloudAccountAsyncNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update vcf cloud account async not found response has a 4xx status code
func (o *UpdateVcfCloudAccountAsyncNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update vcf cloud account async not found response has a 5xx status code
func (o *UpdateVcfCloudAccountAsyncNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update vcf cloud account async not found response a status code equal to that given
func (o *UpdateVcfCloudAccountAsyncNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateVcfCloudAccountAsyncNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vcf/{id}][%d] updateVcfCloudAccountAsyncNotFound  %+v", 404, o.Payload)
}

func (o *UpdateVcfCloudAccountAsyncNotFound) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vcf/{id}][%d] updateVcfCloudAccountAsyncNotFound  %+v", 404, o.Payload)
}

func (o *UpdateVcfCloudAccountAsyncNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateVcfCloudAccountAsyncNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
