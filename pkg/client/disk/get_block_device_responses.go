// Code generated by go-swagger; DO NOT EDIT.

package disk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetBlockDeviceReader is a Reader for the GetBlockDevice structure.
type GetBlockDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlockDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBlockDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetBlockDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBlockDeviceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlockDeviceOK creates a GetBlockDeviceOK with default headers values
func NewGetBlockDeviceOK() *GetBlockDeviceOK {
	return &GetBlockDeviceOK{}
}

/*GetBlockDeviceOK handles this case with default header values.

successful operation
*/
type GetBlockDeviceOK struct {
	Payload *models.BlockDevice
}

func (o *GetBlockDeviceOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/block-devices/{id}][%d] getBlockDeviceOK  %+v", 200, o.Payload)
}

func (o *GetBlockDeviceOK) GetPayload() *models.BlockDevice {
	return o.Payload
}

func (o *GetBlockDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BlockDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlockDeviceForbidden creates a GetBlockDeviceForbidden with default headers values
func NewGetBlockDeviceForbidden() *GetBlockDeviceForbidden {
	return &GetBlockDeviceForbidden{}
}

/*GetBlockDeviceForbidden handles this case with default header values.

Forbidden
*/
type GetBlockDeviceForbidden struct {
}

func (o *GetBlockDeviceForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/block-devices/{id}][%d] getBlockDeviceForbidden ", 403)
}

func (o *GetBlockDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlockDeviceNotFound creates a GetBlockDeviceNotFound with default headers values
func NewGetBlockDeviceNotFound() *GetBlockDeviceNotFound {
	return &GetBlockDeviceNotFound{}
}

/*GetBlockDeviceNotFound handles this case with default header values.

Not Found
*/
type GetBlockDeviceNotFound struct {
	Payload *models.Error
}

func (o *GetBlockDeviceNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/block-devices/{id}][%d] getBlockDeviceNotFound  %+v", 404, o.Payload)
}

func (o *GetBlockDeviceNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBlockDeviceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
