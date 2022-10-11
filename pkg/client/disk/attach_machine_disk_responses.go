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

// AttachMachineDiskReader is a Reader for the AttachMachineDisk structure.
type AttachMachineDiskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachMachineDiskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttachMachineDiskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAttachMachineDiskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAttachMachineDiskForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAttachMachineDiskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAttachMachineDiskOK creates a AttachMachineDiskOK with default headers values
func NewAttachMachineDiskOK() *AttachMachineDiskOK {
	return &AttachMachineDiskOK{}
}

/*
AttachMachineDiskOK describes a response with status code 200, with default header values.

successful operation
*/
type AttachMachineDiskOK struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this attach machine disk o k response has a 2xx status code
func (o *AttachMachineDiskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this attach machine disk o k response has a 3xx status code
func (o *AttachMachineDiskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach machine disk o k response has a 4xx status code
func (o *AttachMachineDiskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this attach machine disk o k response has a 5xx status code
func (o *AttachMachineDiskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this attach machine disk o k response a status code equal to that given
func (o *AttachMachineDiskOK) IsCode(code int) bool {
	return code == 200
}

func (o *AttachMachineDiskOK) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/disks][%d] attachMachineDiskOK  %+v", 200, o.Payload)
}

func (o *AttachMachineDiskOK) String() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/disks][%d] attachMachineDiskOK  %+v", 200, o.Payload)
}

func (o *AttachMachineDiskOK) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *AttachMachineDiskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachMachineDiskBadRequest creates a AttachMachineDiskBadRequest with default headers values
func NewAttachMachineDiskBadRequest() *AttachMachineDiskBadRequest {
	return &AttachMachineDiskBadRequest{}
}

/*
AttachMachineDiskBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type AttachMachineDiskBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this attach machine disk bad request response has a 2xx status code
func (o *AttachMachineDiskBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach machine disk bad request response has a 3xx status code
func (o *AttachMachineDiskBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach machine disk bad request response has a 4xx status code
func (o *AttachMachineDiskBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach machine disk bad request response has a 5xx status code
func (o *AttachMachineDiskBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this attach machine disk bad request response a status code equal to that given
func (o *AttachMachineDiskBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AttachMachineDiskBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/disks][%d] attachMachineDiskBadRequest  %+v", 400, o.Payload)
}

func (o *AttachMachineDiskBadRequest) String() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/disks][%d] attachMachineDiskBadRequest  %+v", 400, o.Payload)
}

func (o *AttachMachineDiskBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AttachMachineDiskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachMachineDiskForbidden creates a AttachMachineDiskForbidden with default headers values
func NewAttachMachineDiskForbidden() *AttachMachineDiskForbidden {
	return &AttachMachineDiskForbidden{}
}

/*
AttachMachineDiskForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AttachMachineDiskForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this attach machine disk forbidden response has a 2xx status code
func (o *AttachMachineDiskForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach machine disk forbidden response has a 3xx status code
func (o *AttachMachineDiskForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach machine disk forbidden response has a 4xx status code
func (o *AttachMachineDiskForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach machine disk forbidden response has a 5xx status code
func (o *AttachMachineDiskForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this attach machine disk forbidden response a status code equal to that given
func (o *AttachMachineDiskForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AttachMachineDiskForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/disks][%d] attachMachineDiskForbidden  %+v", 403, o.Payload)
}

func (o *AttachMachineDiskForbidden) String() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/disks][%d] attachMachineDiskForbidden  %+v", 403, o.Payload)
}

func (o *AttachMachineDiskForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *AttachMachineDiskForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachMachineDiskNotFound creates a AttachMachineDiskNotFound with default headers values
func NewAttachMachineDiskNotFound() *AttachMachineDiskNotFound {
	return &AttachMachineDiskNotFound{}
}

/*
AttachMachineDiskNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AttachMachineDiskNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this attach machine disk not found response has a 2xx status code
func (o *AttachMachineDiskNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach machine disk not found response has a 3xx status code
func (o *AttachMachineDiskNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach machine disk not found response has a 4xx status code
func (o *AttachMachineDiskNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach machine disk not found response has a 5xx status code
func (o *AttachMachineDiskNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this attach machine disk not found response a status code equal to that given
func (o *AttachMachineDiskNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AttachMachineDiskNotFound) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/disks][%d] attachMachineDiskNotFound  %+v", 404, o.Payload)
}

func (o *AttachMachineDiskNotFound) String() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/disks][%d] attachMachineDiskNotFound  %+v", 404, o.Payload)
}

func (o *AttachMachineDiskNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *AttachMachineDiskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
