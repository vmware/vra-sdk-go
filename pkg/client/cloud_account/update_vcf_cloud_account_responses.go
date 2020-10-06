// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateVcfCloudAccountReader is a Reader for the UpdateVcfCloudAccount structure.
type UpdateVcfCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVcfCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVcfCloudAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateVcfCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVcfCloudAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateVcfCloudAccountOK creates a UpdateVcfCloudAccountOK with default headers values
func NewUpdateVcfCloudAccountOK() *UpdateVcfCloudAccountOK {
	return &UpdateVcfCloudAccountOK{}
}

/*UpdateVcfCloudAccountOK handles this case with default header values.

successful operation
*/
type UpdateVcfCloudAccountOK struct {
	Payload *models.CloudAccountVcf
}

func (o *UpdateVcfCloudAccountOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vcf/{id}][%d] updateVcfCloudAccountOK  %+v", 200, o.Payload)
}

func (o *UpdateVcfCloudAccountOK) GetPayload() *models.CloudAccountVcf {
	return o.Payload
}

func (o *UpdateVcfCloudAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountVcf)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVcfCloudAccountForbidden creates a UpdateVcfCloudAccountForbidden with default headers values
func NewUpdateVcfCloudAccountForbidden() *UpdateVcfCloudAccountForbidden {
	return &UpdateVcfCloudAccountForbidden{}
}

/*UpdateVcfCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type UpdateVcfCloudAccountForbidden struct {
}

func (o *UpdateVcfCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vcf/{id}][%d] updateVcfCloudAccountForbidden ", 403)
}

func (o *UpdateVcfCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVcfCloudAccountNotFound creates a UpdateVcfCloudAccountNotFound with default headers values
func NewUpdateVcfCloudAccountNotFound() *UpdateVcfCloudAccountNotFound {
	return &UpdateVcfCloudAccountNotFound{}
}

/*UpdateVcfCloudAccountNotFound handles this case with default header values.

Not Found
*/
type UpdateVcfCloudAccountNotFound struct {
	Payload *models.Error
}

func (o *UpdateVcfCloudAccountNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vcf/{id}][%d] updateVcfCloudAccountNotFound  %+v", 404, o.Payload)
}

func (o *UpdateVcfCloudAccountNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateVcfCloudAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
