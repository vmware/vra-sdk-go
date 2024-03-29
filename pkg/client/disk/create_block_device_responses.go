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

// CreateBlockDeviceReader is a Reader for the CreateBlockDevice structure.
type CreateBlockDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBlockDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCreateBlockDeviceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateBlockDeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateBlockDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateBlockDeviceAccepted creates a CreateBlockDeviceAccepted with default headers values
func NewCreateBlockDeviceAccepted() *CreateBlockDeviceAccepted {
	return &CreateBlockDeviceAccepted{}
}

/*
CreateBlockDeviceAccepted describes a response with status code 202, with default header values.

successful operation
*/
type CreateBlockDeviceAccepted struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this create block device accepted response has a 2xx status code
func (o *CreateBlockDeviceAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create block device accepted response has a 3xx status code
func (o *CreateBlockDeviceAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create block device accepted response has a 4xx status code
func (o *CreateBlockDeviceAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create block device accepted response has a 5xx status code
func (o *CreateBlockDeviceAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create block device accepted response a status code equal to that given
func (o *CreateBlockDeviceAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateBlockDeviceAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices][%d] createBlockDeviceAccepted  %+v", 202, o.Payload)
}

func (o *CreateBlockDeviceAccepted) String() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices][%d] createBlockDeviceAccepted  %+v", 202, o.Payload)
}

func (o *CreateBlockDeviceAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *CreateBlockDeviceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBlockDeviceBadRequest creates a CreateBlockDeviceBadRequest with default headers values
func NewCreateBlockDeviceBadRequest() *CreateBlockDeviceBadRequest {
	return &CreateBlockDeviceBadRequest{}
}

/*
CreateBlockDeviceBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type CreateBlockDeviceBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create block device bad request response has a 2xx status code
func (o *CreateBlockDeviceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create block device bad request response has a 3xx status code
func (o *CreateBlockDeviceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create block device bad request response has a 4xx status code
func (o *CreateBlockDeviceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create block device bad request response has a 5xx status code
func (o *CreateBlockDeviceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create block device bad request response a status code equal to that given
func (o *CreateBlockDeviceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateBlockDeviceBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices][%d] createBlockDeviceBadRequest  %+v", 400, o.Payload)
}

func (o *CreateBlockDeviceBadRequest) String() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices][%d] createBlockDeviceBadRequest  %+v", 400, o.Payload)
}

func (o *CreateBlockDeviceBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateBlockDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBlockDeviceForbidden creates a CreateBlockDeviceForbidden with default headers values
func NewCreateBlockDeviceForbidden() *CreateBlockDeviceForbidden {
	return &CreateBlockDeviceForbidden{}
}

/*
CreateBlockDeviceForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateBlockDeviceForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this create block device forbidden response has a 2xx status code
func (o *CreateBlockDeviceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create block device forbidden response has a 3xx status code
func (o *CreateBlockDeviceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create block device forbidden response has a 4xx status code
func (o *CreateBlockDeviceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create block device forbidden response has a 5xx status code
func (o *CreateBlockDeviceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create block device forbidden response a status code equal to that given
func (o *CreateBlockDeviceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateBlockDeviceForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices][%d] createBlockDeviceForbidden  %+v", 403, o.Payload)
}

func (o *CreateBlockDeviceForbidden) String() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices][%d] createBlockDeviceForbidden  %+v", 403, o.Payload)
}

func (o *CreateBlockDeviceForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *CreateBlockDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
