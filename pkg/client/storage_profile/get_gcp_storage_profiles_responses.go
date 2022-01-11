// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetGcpStorageProfilesReader is a Reader for the GetGcpStorageProfiles structure.
type GetGcpStorageProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGcpStorageProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGcpStorageProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetGcpStorageProfilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGcpStorageProfilesOK creates a GetGcpStorageProfilesOK with default headers values
func NewGetGcpStorageProfilesOK() *GetGcpStorageProfilesOK {
	return &GetGcpStorageProfilesOK{}
}

/* GetGcpStorageProfilesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetGcpStorageProfilesOK struct {
	Payload *models.StorageProfileGcpResult
}

func (o *GetGcpStorageProfilesOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-gcp][%d] getGcpStorageProfilesOK  %+v", 200, o.Payload)
}
func (o *GetGcpStorageProfilesOK) GetPayload() *models.StorageProfileGcpResult {
	return o.Payload
}

func (o *GetGcpStorageProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageProfileGcpResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGcpStorageProfilesForbidden creates a GetGcpStorageProfilesForbidden with default headers values
func NewGetGcpStorageProfilesForbidden() *GetGcpStorageProfilesForbidden {
	return &GetGcpStorageProfilesForbidden{}
}

/* GetGcpStorageProfilesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetGcpStorageProfilesForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *GetGcpStorageProfilesForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-gcp][%d] getGcpStorageProfilesForbidden  %+v", 403, o.Payload)
}
func (o *GetGcpStorageProfilesForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetGcpStorageProfilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
