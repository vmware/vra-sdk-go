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

// GetMachineDiskReader is a Reader for the GetMachineDisk structure.
type GetMachineDiskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachineDiskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMachineDiskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetMachineDiskForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMachineDiskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMachineDiskOK creates a GetMachineDiskOK with default headers values
func NewGetMachineDiskOK() *GetMachineDiskOK {
	return &GetMachineDiskOK{}
}

/* GetMachineDiskOK describes a response with status code 200, with default header values.

successful operation
*/
type GetMachineDiskOK struct {
	Payload *models.BlockDevice
}

func (o *GetMachineDiskOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines/{id}/disks/{diskId}][%d] getMachineDiskOK  %+v", 200, o.Payload)
}
func (o *GetMachineDiskOK) GetPayload() *models.BlockDevice {
	return o.Payload
}

func (o *GetMachineDiskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BlockDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachineDiskForbidden creates a GetMachineDiskForbidden with default headers values
func NewGetMachineDiskForbidden() *GetMachineDiskForbidden {
	return &GetMachineDiskForbidden{}
}

/* GetMachineDiskForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetMachineDiskForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *GetMachineDiskForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines/{id}/disks/{diskId}][%d] getMachineDiskForbidden  %+v", 403, o.Payload)
}
func (o *GetMachineDiskForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetMachineDiskForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachineDiskNotFound creates a GetMachineDiskNotFound with default headers values
func NewGetMachineDiskNotFound() *GetMachineDiskNotFound {
	return &GetMachineDiskNotFound{}
}

/* GetMachineDiskNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetMachineDiskNotFound struct {
	Payload *models.Error
}

func (o *GetMachineDiskNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines/{id}/disks/{diskId}][%d] getMachineDiskNotFound  %+v", 404, o.Payload)
}
func (o *GetMachineDiskNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMachineDiskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
