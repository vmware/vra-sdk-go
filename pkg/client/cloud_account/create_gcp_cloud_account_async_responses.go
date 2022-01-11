// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateGcpCloudAccountAsyncReader is a Reader for the CreateGcpCloudAccountAsync structure.
type CreateGcpCloudAccountAsyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGcpCloudAccountAsyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCreateGcpCloudAccountAsyncAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateGcpCloudAccountAsyncBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateGcpCloudAccountAsyncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateGcpCloudAccountAsyncAccepted creates a CreateGcpCloudAccountAsyncAccepted with default headers values
func NewCreateGcpCloudAccountAsyncAccepted() *CreateGcpCloudAccountAsyncAccepted {
	return &CreateGcpCloudAccountAsyncAccepted{}
}

/* CreateGcpCloudAccountAsyncAccepted describes a response with status code 202, with default header values.

successful operation
*/
type CreateGcpCloudAccountAsyncAccepted struct {
	Payload *models.RequestTracker
}

func (o *CreateGcpCloudAccountAsyncAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-gcp][%d] createGcpCloudAccountAsyncAccepted  %+v", 202, o.Payload)
}
func (o *CreateGcpCloudAccountAsyncAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *CreateGcpCloudAccountAsyncAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGcpCloudAccountAsyncBadRequest creates a CreateGcpCloudAccountAsyncBadRequest with default headers values
func NewCreateGcpCloudAccountAsyncBadRequest() *CreateGcpCloudAccountAsyncBadRequest {
	return &CreateGcpCloudAccountAsyncBadRequest{}
}

/* CreateGcpCloudAccountAsyncBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type CreateGcpCloudAccountAsyncBadRequest struct {
	Payload *models.Error
}

func (o *CreateGcpCloudAccountAsyncBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-gcp][%d] createGcpCloudAccountAsyncBadRequest  %+v", 400, o.Payload)
}
func (o *CreateGcpCloudAccountAsyncBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGcpCloudAccountAsyncBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGcpCloudAccountAsyncForbidden creates a CreateGcpCloudAccountAsyncForbidden with default headers values
func NewCreateGcpCloudAccountAsyncForbidden() *CreateGcpCloudAccountAsyncForbidden {
	return &CreateGcpCloudAccountAsyncForbidden{}
}

/* CreateGcpCloudAccountAsyncForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateGcpCloudAccountAsyncForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *CreateGcpCloudAccountAsyncForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-gcp][%d] createGcpCloudAccountAsyncForbidden  %+v", 403, o.Payload)
}
func (o *CreateGcpCloudAccountAsyncForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *CreateGcpCloudAccountAsyncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
