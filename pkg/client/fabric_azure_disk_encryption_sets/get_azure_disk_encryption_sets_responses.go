// Code generated by go-swagger; DO NOT EDIT.

package fabric_azure_disk_encryption_sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetAzureDiskEncryptionSetsReader is a Reader for the GetAzureDiskEncryptionSets structure.
type GetAzureDiskEncryptionSetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAzureDiskEncryptionSetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAzureDiskEncryptionSetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAzureDiskEncryptionSetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAzureDiskEncryptionSetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAzureDiskEncryptionSetsOK creates a GetAzureDiskEncryptionSetsOK with default headers values
func NewGetAzureDiskEncryptionSetsOK() *GetAzureDiskEncryptionSetsOK {
	return &GetAzureDiskEncryptionSetsOK{}
}

/* GetAzureDiskEncryptionSetsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAzureDiskEncryptionSetsOK struct {
	Payload *models.DiskEncryptionSetList
}

func (o *GetAzureDiskEncryptionSetsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-disk-encryption-sets][%d] getAzureDiskEncryptionSetsOK  %+v", 200, o.Payload)
}
func (o *GetAzureDiskEncryptionSetsOK) GetPayload() *models.DiskEncryptionSetList {
	return o.Payload
}

func (o *GetAzureDiskEncryptionSetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DiskEncryptionSetList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAzureDiskEncryptionSetsForbidden creates a GetAzureDiskEncryptionSetsForbidden with default headers values
func NewGetAzureDiskEncryptionSetsForbidden() *GetAzureDiskEncryptionSetsForbidden {
	return &GetAzureDiskEncryptionSetsForbidden{}
}

/* GetAzureDiskEncryptionSetsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAzureDiskEncryptionSetsForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *GetAzureDiskEncryptionSetsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-disk-encryption-sets][%d] getAzureDiskEncryptionSetsForbidden  %+v", 403, o.Payload)
}
func (o *GetAzureDiskEncryptionSetsForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetAzureDiskEncryptionSetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAzureDiskEncryptionSetsNotFound creates a GetAzureDiskEncryptionSetsNotFound with default headers values
func NewGetAzureDiskEncryptionSetsNotFound() *GetAzureDiskEncryptionSetsNotFound {
	return &GetAzureDiskEncryptionSetsNotFound{}
}

/* GetAzureDiskEncryptionSetsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAzureDiskEncryptionSetsNotFound struct {
	Payload *models.Error
}

func (o *GetAzureDiskEncryptionSetsNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-disk-encryption-sets][%d] getAzureDiskEncryptionSetsNotFound  %+v", 404, o.Payload)
}
func (o *GetAzureDiskEncryptionSetsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAzureDiskEncryptionSetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
