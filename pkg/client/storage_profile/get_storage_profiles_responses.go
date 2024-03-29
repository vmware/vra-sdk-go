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

// GetStorageProfilesReader is a Reader for the GetStorageProfiles structure.
type GetStorageProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetStorageProfilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStorageProfilesOK creates a GetStorageProfilesOK with default headers values
func NewGetStorageProfilesOK() *GetStorageProfilesOK {
	return &GetStorageProfilesOK{}
}

/*
GetStorageProfilesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetStorageProfilesOK struct {
	Payload *models.StorageProfileResult
}

// IsSuccess returns true when this get storage profiles o k response has a 2xx status code
func (o *GetStorageProfilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get storage profiles o k response has a 3xx status code
func (o *GetStorageProfilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get storage profiles o k response has a 4xx status code
func (o *GetStorageProfilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get storage profiles o k response has a 5xx status code
func (o *GetStorageProfilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get storage profiles o k response a status code equal to that given
func (o *GetStorageProfilesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetStorageProfilesOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles][%d] getStorageProfilesOK  %+v", 200, o.Payload)
}

func (o *GetStorageProfilesOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles][%d] getStorageProfilesOK  %+v", 200, o.Payload)
}

func (o *GetStorageProfilesOK) GetPayload() *models.StorageProfileResult {
	return o.Payload
}

func (o *GetStorageProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageProfileResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageProfilesForbidden creates a GetStorageProfilesForbidden with default headers values
func NewGetStorageProfilesForbidden() *GetStorageProfilesForbidden {
	return &GetStorageProfilesForbidden{}
}

/*
GetStorageProfilesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetStorageProfilesForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get storage profiles forbidden response has a 2xx status code
func (o *GetStorageProfilesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get storage profiles forbidden response has a 3xx status code
func (o *GetStorageProfilesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get storage profiles forbidden response has a 4xx status code
func (o *GetStorageProfilesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get storage profiles forbidden response has a 5xx status code
func (o *GetStorageProfilesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get storage profiles forbidden response a status code equal to that given
func (o *GetStorageProfilesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetStorageProfilesForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles][%d] getStorageProfilesForbidden  %+v", 403, o.Payload)
}

func (o *GetStorageProfilesForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles][%d] getStorageProfilesForbidden  %+v", 403, o.Payload)
}

func (o *GetStorageProfilesForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetStorageProfilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
