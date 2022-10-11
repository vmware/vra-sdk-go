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

// CreateVmcCloudAccountAsyncReader is a Reader for the CreateVmcCloudAccountAsync structure.
type CreateVmcCloudAccountAsyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVmcCloudAccountAsyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCreateVmcCloudAccountAsyncAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVmcCloudAccountAsyncBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateVmcCloudAccountAsyncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVmcCloudAccountAsyncAccepted creates a CreateVmcCloudAccountAsyncAccepted with default headers values
func NewCreateVmcCloudAccountAsyncAccepted() *CreateVmcCloudAccountAsyncAccepted {
	return &CreateVmcCloudAccountAsyncAccepted{}
}

/*
CreateVmcCloudAccountAsyncAccepted describes a response with status code 202, with default header values.

successful operation
*/
type CreateVmcCloudAccountAsyncAccepted struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this create vmc cloud account async accepted response has a 2xx status code
func (o *CreateVmcCloudAccountAsyncAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create vmc cloud account async accepted response has a 3xx status code
func (o *CreateVmcCloudAccountAsyncAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create vmc cloud account async accepted response has a 4xx status code
func (o *CreateVmcCloudAccountAsyncAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create vmc cloud account async accepted response has a 5xx status code
func (o *CreateVmcCloudAccountAsyncAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create vmc cloud account async accepted response a status code equal to that given
func (o *CreateVmcCloudAccountAsyncAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateVmcCloudAccountAsyncAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc][%d] createVmcCloudAccountAsyncAccepted  %+v", 202, o.Payload)
}

func (o *CreateVmcCloudAccountAsyncAccepted) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc][%d] createVmcCloudAccountAsyncAccepted  %+v", 202, o.Payload)
}

func (o *CreateVmcCloudAccountAsyncAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *CreateVmcCloudAccountAsyncAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVmcCloudAccountAsyncBadRequest creates a CreateVmcCloudAccountAsyncBadRequest with default headers values
func NewCreateVmcCloudAccountAsyncBadRequest() *CreateVmcCloudAccountAsyncBadRequest {
	return &CreateVmcCloudAccountAsyncBadRequest{}
}

/*
CreateVmcCloudAccountAsyncBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type CreateVmcCloudAccountAsyncBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create vmc cloud account async bad request response has a 2xx status code
func (o *CreateVmcCloudAccountAsyncBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create vmc cloud account async bad request response has a 3xx status code
func (o *CreateVmcCloudAccountAsyncBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create vmc cloud account async bad request response has a 4xx status code
func (o *CreateVmcCloudAccountAsyncBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create vmc cloud account async bad request response has a 5xx status code
func (o *CreateVmcCloudAccountAsyncBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create vmc cloud account async bad request response a status code equal to that given
func (o *CreateVmcCloudAccountAsyncBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateVmcCloudAccountAsyncBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc][%d] createVmcCloudAccountAsyncBadRequest  %+v", 400, o.Payload)
}

func (o *CreateVmcCloudAccountAsyncBadRequest) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc][%d] createVmcCloudAccountAsyncBadRequest  %+v", 400, o.Payload)
}

func (o *CreateVmcCloudAccountAsyncBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateVmcCloudAccountAsyncBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVmcCloudAccountAsyncForbidden creates a CreateVmcCloudAccountAsyncForbidden with default headers values
func NewCreateVmcCloudAccountAsyncForbidden() *CreateVmcCloudAccountAsyncForbidden {
	return &CreateVmcCloudAccountAsyncForbidden{}
}

/*
CreateVmcCloudAccountAsyncForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateVmcCloudAccountAsyncForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this create vmc cloud account async forbidden response has a 2xx status code
func (o *CreateVmcCloudAccountAsyncForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create vmc cloud account async forbidden response has a 3xx status code
func (o *CreateVmcCloudAccountAsyncForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create vmc cloud account async forbidden response has a 4xx status code
func (o *CreateVmcCloudAccountAsyncForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create vmc cloud account async forbidden response has a 5xx status code
func (o *CreateVmcCloudAccountAsyncForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create vmc cloud account async forbidden response a status code equal to that given
func (o *CreateVmcCloudAccountAsyncForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateVmcCloudAccountAsyncForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc][%d] createVmcCloudAccountAsyncForbidden  %+v", 403, o.Payload)
}

func (o *CreateVmcCloudAccountAsyncForbidden) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc][%d] createVmcCloudAccountAsyncForbidden  %+v", 403, o.Payload)
}

func (o *CreateVmcCloudAccountAsyncForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *CreateVmcCloudAccountAsyncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
