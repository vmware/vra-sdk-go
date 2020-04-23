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

// UpdateNsxTCloudAccountReader is a Reader for the UpdateNsxTCloudAccount structure.
type UpdateNsxTCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNsxTCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNsxTCloudAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateNsxTCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateNsxTCloudAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNsxTCloudAccountOK creates a UpdateNsxTCloudAccountOK with default headers values
func NewUpdateNsxTCloudAccountOK() *UpdateNsxTCloudAccountOK {
	return &UpdateNsxTCloudAccountOK{}
}

/*UpdateNsxTCloudAccountOK handles this case with default header values.

successful operation
*/
type UpdateNsxTCloudAccountOK struct {
	Payload *models.CloudAccountNsxT
}

func (o *UpdateNsxTCloudAccountOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-nsx-t/{id}][%d] updateNsxTCloudAccountOK  %+v", 200, o.Payload)
}

func (o *UpdateNsxTCloudAccountOK) GetPayload() *models.CloudAccountNsxT {
	return o.Payload
}

func (o *UpdateNsxTCloudAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountNsxT)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNsxTCloudAccountForbidden creates a UpdateNsxTCloudAccountForbidden with default headers values
func NewUpdateNsxTCloudAccountForbidden() *UpdateNsxTCloudAccountForbidden {
	return &UpdateNsxTCloudAccountForbidden{}
}

/*UpdateNsxTCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type UpdateNsxTCloudAccountForbidden struct {
}

func (o *UpdateNsxTCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-nsx-t/{id}][%d] updateNsxTCloudAccountForbidden ", 403)
}

func (o *UpdateNsxTCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNsxTCloudAccountNotFound creates a UpdateNsxTCloudAccountNotFound with default headers values
func NewUpdateNsxTCloudAccountNotFound() *UpdateNsxTCloudAccountNotFound {
	return &UpdateNsxTCloudAccountNotFound{}
}

/*UpdateNsxTCloudAccountNotFound handles this case with default header values.

Not Found
*/
type UpdateNsxTCloudAccountNotFound struct {
	Payload *models.Error
}

func (o *UpdateNsxTCloudAccountNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-nsx-t/{id}][%d] updateNsxTCloudAccountNotFound  %+v", 404, o.Payload)
}

func (o *UpdateNsxTCloudAccountNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateNsxTCloudAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
