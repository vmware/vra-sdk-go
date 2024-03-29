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

// RestartMachineReader is a Reader for the RestartMachine structure.
type RestartMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRestartMachineAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRestartMachineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRestartMachineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRestartMachineAccepted creates a RestartMachineAccepted with default headers values
func NewRestartMachineAccepted() *RestartMachineAccepted {
	return &RestartMachineAccepted{}
}

/*
RestartMachineAccepted describes a response with status code 202, with default header values.

successful operation
*/
type RestartMachineAccepted struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this restart machine accepted response has a 2xx status code
func (o *RestartMachineAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restart machine accepted response has a 3xx status code
func (o *RestartMachineAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restart machine accepted response has a 4xx status code
func (o *RestartMachineAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this restart machine accepted response has a 5xx status code
func (o *RestartMachineAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this restart machine accepted response a status code equal to that given
func (o *RestartMachineAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *RestartMachineAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/restart][%d] restartMachineAccepted  %+v", 202, o.Payload)
}

func (o *RestartMachineAccepted) String() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/restart][%d] restartMachineAccepted  %+v", 202, o.Payload)
}

func (o *RestartMachineAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *RestartMachineAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartMachineForbidden creates a RestartMachineForbidden with default headers values
func NewRestartMachineForbidden() *RestartMachineForbidden {
	return &RestartMachineForbidden{}
}

/*
RestartMachineForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RestartMachineForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this restart machine forbidden response has a 2xx status code
func (o *RestartMachineForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restart machine forbidden response has a 3xx status code
func (o *RestartMachineForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restart machine forbidden response has a 4xx status code
func (o *RestartMachineForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this restart machine forbidden response has a 5xx status code
func (o *RestartMachineForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this restart machine forbidden response a status code equal to that given
func (o *RestartMachineForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *RestartMachineForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/restart][%d] restartMachineForbidden  %+v", 403, o.Payload)
}

func (o *RestartMachineForbidden) String() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/restart][%d] restartMachineForbidden  %+v", 403, o.Payload)
}

func (o *RestartMachineForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *RestartMachineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartMachineNotFound creates a RestartMachineNotFound with default headers values
func NewRestartMachineNotFound() *RestartMachineNotFound {
	return &RestartMachineNotFound{}
}

/*
RestartMachineNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RestartMachineNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this restart machine not found response has a 2xx status code
func (o *RestartMachineNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restart machine not found response has a 3xx status code
func (o *RestartMachineNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restart machine not found response has a 4xx status code
func (o *RestartMachineNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this restart machine not found response has a 5xx status code
func (o *RestartMachineNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this restart machine not found response a status code equal to that given
func (o *RestartMachineNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *RestartMachineNotFound) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/restart][%d] restartMachineNotFound  %+v", 404, o.Payload)
}

func (o *RestartMachineNotFound) String() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/restart][%d] restartMachineNotFound  %+v", 404, o.Payload)
}

func (o *RestartMachineNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RestartMachineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
