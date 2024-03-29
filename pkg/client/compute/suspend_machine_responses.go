// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// SuspendMachineReader is a Reader for the SuspendMachine structure.
type SuspendMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuspendMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSuspendMachineAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSuspendMachineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSuspendMachineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSuspendMachineAccepted creates a SuspendMachineAccepted with default headers values
func NewSuspendMachineAccepted() *SuspendMachineAccepted {
	return &SuspendMachineAccepted{}
}

/*
SuspendMachineAccepted describes a response with status code 202, with default header values.

successful operation
*/
type SuspendMachineAccepted struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this suspend machine accepted response has a 2xx status code
func (o *SuspendMachineAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this suspend machine accepted response has a 3xx status code
func (o *SuspendMachineAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this suspend machine accepted response has a 4xx status code
func (o *SuspendMachineAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this suspend machine accepted response has a 5xx status code
func (o *SuspendMachineAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this suspend machine accepted response a status code equal to that given
func (o *SuspendMachineAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *SuspendMachineAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/suspend][%d] suspendMachineAccepted  %+v", 202, o.Payload)
}

func (o *SuspendMachineAccepted) String() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/suspend][%d] suspendMachineAccepted  %+v", 202, o.Payload)
}

func (o *SuspendMachineAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *SuspendMachineAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuspendMachineForbidden creates a SuspendMachineForbidden with default headers values
func NewSuspendMachineForbidden() *SuspendMachineForbidden {
	return &SuspendMachineForbidden{}
}

/*
SuspendMachineForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SuspendMachineForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this suspend machine forbidden response has a 2xx status code
func (o *SuspendMachineForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this suspend machine forbidden response has a 3xx status code
func (o *SuspendMachineForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this suspend machine forbidden response has a 4xx status code
func (o *SuspendMachineForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this suspend machine forbidden response has a 5xx status code
func (o *SuspendMachineForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this suspend machine forbidden response a status code equal to that given
func (o *SuspendMachineForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SuspendMachineForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/suspend][%d] suspendMachineForbidden  %+v", 403, o.Payload)
}

func (o *SuspendMachineForbidden) String() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/suspend][%d] suspendMachineForbidden  %+v", 403, o.Payload)
}

func (o *SuspendMachineForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *SuspendMachineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuspendMachineNotFound creates a SuspendMachineNotFound with default headers values
func NewSuspendMachineNotFound() *SuspendMachineNotFound {
	return &SuspendMachineNotFound{}
}

/*
SuspendMachineNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SuspendMachineNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this suspend machine not found response has a 2xx status code
func (o *SuspendMachineNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this suspend machine not found response has a 3xx status code
func (o *SuspendMachineNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this suspend machine not found response has a 4xx status code
func (o *SuspendMachineNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this suspend machine not found response has a 5xx status code
func (o *SuspendMachineNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this suspend machine not found response a status code equal to that given
func (o *SuspendMachineNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SuspendMachineNotFound) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/suspend][%d] suspendMachineNotFound  %+v", 404, o.Payload)
}

func (o *SuspendMachineNotFound) String() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/suspend][%d] suspendMachineNotFound  %+v", 404, o.Payload)
}

func (o *SuspendMachineNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SuspendMachineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
