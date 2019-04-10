// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// CreateAzureStorageProfileReader is a Reader for the CreateAzureStorageProfile structure.
type CreateAzureStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAzureStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateAzureStorageProfileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateAzureStorageProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateAzureStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAzureStorageProfileCreated creates a CreateAzureStorageProfileCreated with default headers values
func NewCreateAzureStorageProfileCreated() *CreateAzureStorageProfileCreated {
	return &CreateAzureStorageProfileCreated{}
}

/*CreateAzureStorageProfileCreated handles this case with default header values.

successful operation
*/
type CreateAzureStorageProfileCreated struct {
	Payload *models.AzureStorageProfile
}

func (o *CreateAzureStorageProfileCreated) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles-azure][%d] createAzureStorageProfileCreated  %+v", 201, o.Payload)
}

func (o *CreateAzureStorageProfileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureStorageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAzureStorageProfileBadRequest creates a CreateAzureStorageProfileBadRequest with default headers values
func NewCreateAzureStorageProfileBadRequest() *CreateAzureStorageProfileBadRequest {
	return &CreateAzureStorageProfileBadRequest{}
}

/*CreateAzureStorageProfileBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type CreateAzureStorageProfileBadRequest struct {
}

func (o *CreateAzureStorageProfileBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles-azure][%d] createAzureStorageProfileBadRequest ", 400)
}

func (o *CreateAzureStorageProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAzureStorageProfileForbidden creates a CreateAzureStorageProfileForbidden with default headers values
func NewCreateAzureStorageProfileForbidden() *CreateAzureStorageProfileForbidden {
	return &CreateAzureStorageProfileForbidden{}
}

/*CreateAzureStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type CreateAzureStorageProfileForbidden struct {
}

func (o *CreateAzureStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles-azure][%d] createAzureStorageProfileForbidden ", 403)
}

func (o *CreateAzureStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}