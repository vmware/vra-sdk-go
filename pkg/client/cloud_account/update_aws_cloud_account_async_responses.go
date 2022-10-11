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

// UpdateAWSCloudAccountAsyncReader is a Reader for the UpdateAWSCloudAccountAsync structure.
type UpdateAWSCloudAccountAsyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAWSCloudAccountAsyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewUpdateAWSCloudAccountAsyncAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateAWSCloudAccountAsyncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAWSCloudAccountAsyncNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAWSCloudAccountAsyncAccepted creates a UpdateAWSCloudAccountAsyncAccepted with default headers values
func NewUpdateAWSCloudAccountAsyncAccepted() *UpdateAWSCloudAccountAsyncAccepted {
	return &UpdateAWSCloudAccountAsyncAccepted{}
}

/*
UpdateAWSCloudAccountAsyncAccepted describes a response with status code 202, with default header values.

successful operation
*/
type UpdateAWSCloudAccountAsyncAccepted struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this update a w s cloud account async accepted response has a 2xx status code
func (o *UpdateAWSCloudAccountAsyncAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update a w s cloud account async accepted response has a 3xx status code
func (o *UpdateAWSCloudAccountAsyncAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update a w s cloud account async accepted response has a 4xx status code
func (o *UpdateAWSCloudAccountAsyncAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this update a w s cloud account async accepted response has a 5xx status code
func (o *UpdateAWSCloudAccountAsyncAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this update a w s cloud account async accepted response a status code equal to that given
func (o *UpdateAWSCloudAccountAsyncAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *UpdateAWSCloudAccountAsyncAccepted) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-aws/{id}][%d] updateAWSCloudAccountAsyncAccepted  %+v", 202, o.Payload)
}

func (o *UpdateAWSCloudAccountAsyncAccepted) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-aws/{id}][%d] updateAWSCloudAccountAsyncAccepted  %+v", 202, o.Payload)
}

func (o *UpdateAWSCloudAccountAsyncAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *UpdateAWSCloudAccountAsyncAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAWSCloudAccountAsyncForbidden creates a UpdateAWSCloudAccountAsyncForbidden with default headers values
func NewUpdateAWSCloudAccountAsyncForbidden() *UpdateAWSCloudAccountAsyncForbidden {
	return &UpdateAWSCloudAccountAsyncForbidden{}
}

/*
UpdateAWSCloudAccountAsyncForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateAWSCloudAccountAsyncForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this update a w s cloud account async forbidden response has a 2xx status code
func (o *UpdateAWSCloudAccountAsyncForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update a w s cloud account async forbidden response has a 3xx status code
func (o *UpdateAWSCloudAccountAsyncForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update a w s cloud account async forbidden response has a 4xx status code
func (o *UpdateAWSCloudAccountAsyncForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update a w s cloud account async forbidden response has a 5xx status code
func (o *UpdateAWSCloudAccountAsyncForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update a w s cloud account async forbidden response a status code equal to that given
func (o *UpdateAWSCloudAccountAsyncForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateAWSCloudAccountAsyncForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-aws/{id}][%d] updateAWSCloudAccountAsyncForbidden  %+v", 403, o.Payload)
}

func (o *UpdateAWSCloudAccountAsyncForbidden) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-aws/{id}][%d] updateAWSCloudAccountAsyncForbidden  %+v", 403, o.Payload)
}

func (o *UpdateAWSCloudAccountAsyncForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *UpdateAWSCloudAccountAsyncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAWSCloudAccountAsyncNotFound creates a UpdateAWSCloudAccountAsyncNotFound with default headers values
func NewUpdateAWSCloudAccountAsyncNotFound() *UpdateAWSCloudAccountAsyncNotFound {
	return &UpdateAWSCloudAccountAsyncNotFound{}
}

/*
UpdateAWSCloudAccountAsyncNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateAWSCloudAccountAsyncNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update a w s cloud account async not found response has a 2xx status code
func (o *UpdateAWSCloudAccountAsyncNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update a w s cloud account async not found response has a 3xx status code
func (o *UpdateAWSCloudAccountAsyncNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update a w s cloud account async not found response has a 4xx status code
func (o *UpdateAWSCloudAccountAsyncNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update a w s cloud account async not found response has a 5xx status code
func (o *UpdateAWSCloudAccountAsyncNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update a w s cloud account async not found response a status code equal to that given
func (o *UpdateAWSCloudAccountAsyncNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateAWSCloudAccountAsyncNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-aws/{id}][%d] updateAWSCloudAccountAsyncNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAWSCloudAccountAsyncNotFound) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-aws/{id}][%d] updateAWSCloudAccountAsyncNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAWSCloudAccountAsyncNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAWSCloudAccountAsyncNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}