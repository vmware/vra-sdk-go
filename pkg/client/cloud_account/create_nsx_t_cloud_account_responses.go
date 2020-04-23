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

// CreateNsxTCloudAccountReader is a Reader for the CreateNsxTCloudAccount structure.
type CreateNsxTCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNsxTCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNsxTCloudAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateNsxTCloudAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateNsxTCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateNsxTCloudAccountCreated creates a CreateNsxTCloudAccountCreated with default headers values
func NewCreateNsxTCloudAccountCreated() *CreateNsxTCloudAccountCreated {
	return &CreateNsxTCloudAccountCreated{}
}

/*CreateNsxTCloudAccountCreated handles this case with default header values.

successful operation
*/
type CreateNsxTCloudAccountCreated struct {
	Payload *models.CloudAccountNsxT
}

func (o *CreateNsxTCloudAccountCreated) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-nsx-t][%d] createNsxTCloudAccountCreated  %+v", 201, o.Payload)
}

func (o *CreateNsxTCloudAccountCreated) GetPayload() *models.CloudAccountNsxT {
	return o.Payload
}

func (o *CreateNsxTCloudAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountNsxT)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNsxTCloudAccountBadRequest creates a CreateNsxTCloudAccountBadRequest with default headers values
func NewCreateNsxTCloudAccountBadRequest() *CreateNsxTCloudAccountBadRequest {
	return &CreateNsxTCloudAccountBadRequest{}
}

/*CreateNsxTCloudAccountBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type CreateNsxTCloudAccountBadRequest struct {
	Payload *models.Error
}

func (o *CreateNsxTCloudAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-nsx-t][%d] createNsxTCloudAccountBadRequest  %+v", 400, o.Payload)
}

func (o *CreateNsxTCloudAccountBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateNsxTCloudAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNsxTCloudAccountForbidden creates a CreateNsxTCloudAccountForbidden with default headers values
func NewCreateNsxTCloudAccountForbidden() *CreateNsxTCloudAccountForbidden {
	return &CreateNsxTCloudAccountForbidden{}
}

/*CreateNsxTCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type CreateNsxTCloudAccountForbidden struct {
}

func (o *CreateNsxTCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-nsx-t][%d] createNsxTCloudAccountForbidden ", 403)
}

func (o *CreateNsxTCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
