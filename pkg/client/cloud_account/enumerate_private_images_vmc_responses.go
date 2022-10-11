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

// EnumeratePrivateImagesVMCReader is a Reader for the EnumeratePrivateImagesVMC structure.
type EnumeratePrivateImagesVMCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnumeratePrivateImagesVMCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewEnumeratePrivateImagesVMCAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnumeratePrivateImagesVMCBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnumeratePrivateImagesVMCForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnumeratePrivateImagesVMCAccepted creates a EnumeratePrivateImagesVMCAccepted with default headers values
func NewEnumeratePrivateImagesVMCAccepted() *EnumeratePrivateImagesVMCAccepted {
	return &EnumeratePrivateImagesVMCAccepted{}
}

/*
EnumeratePrivateImagesVMCAccepted describes a response with status code 202, with default header values.

successful operation
*/
type EnumeratePrivateImagesVMCAccepted struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this enumerate private images Vm c accepted response has a 2xx status code
func (o *EnumeratePrivateImagesVMCAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enumerate private images Vm c accepted response has a 3xx status code
func (o *EnumeratePrivateImagesVMCAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enumerate private images Vm c accepted response has a 4xx status code
func (o *EnumeratePrivateImagesVMCAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this enumerate private images Vm c accepted response has a 5xx status code
func (o *EnumeratePrivateImagesVMCAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this enumerate private images Vm c accepted response a status code equal to that given
func (o *EnumeratePrivateImagesVMCAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *EnumeratePrivateImagesVMCAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc/{id}/private-image-enumeration][%d] enumeratePrivateImagesVmCAccepted  %+v", 202, o.Payload)
}

func (o *EnumeratePrivateImagesVMCAccepted) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc/{id}/private-image-enumeration][%d] enumeratePrivateImagesVmCAccepted  %+v", 202, o.Payload)
}

func (o *EnumeratePrivateImagesVMCAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *EnumeratePrivateImagesVMCAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumeratePrivateImagesVMCBadRequest creates a EnumeratePrivateImagesVMCBadRequest with default headers values
func NewEnumeratePrivateImagesVMCBadRequest() *EnumeratePrivateImagesVMCBadRequest {
	return &EnumeratePrivateImagesVMCBadRequest{}
}

/*
EnumeratePrivateImagesVMCBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type EnumeratePrivateImagesVMCBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this enumerate private images Vm c bad request response has a 2xx status code
func (o *EnumeratePrivateImagesVMCBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enumerate private images Vm c bad request response has a 3xx status code
func (o *EnumeratePrivateImagesVMCBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enumerate private images Vm c bad request response has a 4xx status code
func (o *EnumeratePrivateImagesVMCBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this enumerate private images Vm c bad request response has a 5xx status code
func (o *EnumeratePrivateImagesVMCBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this enumerate private images Vm c bad request response a status code equal to that given
func (o *EnumeratePrivateImagesVMCBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EnumeratePrivateImagesVMCBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc/{id}/private-image-enumeration][%d] enumeratePrivateImagesVmCBadRequest  %+v", 400, o.Payload)
}

func (o *EnumeratePrivateImagesVMCBadRequest) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc/{id}/private-image-enumeration][%d] enumeratePrivateImagesVmCBadRequest  %+v", 400, o.Payload)
}

func (o *EnumeratePrivateImagesVMCBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *EnumeratePrivateImagesVMCBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumeratePrivateImagesVMCForbidden creates a EnumeratePrivateImagesVMCForbidden with default headers values
func NewEnumeratePrivateImagesVMCForbidden() *EnumeratePrivateImagesVMCForbidden {
	return &EnumeratePrivateImagesVMCForbidden{}
}

/*
EnumeratePrivateImagesVMCForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type EnumeratePrivateImagesVMCForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this enumerate private images Vm c forbidden response has a 2xx status code
func (o *EnumeratePrivateImagesVMCForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enumerate private images Vm c forbidden response has a 3xx status code
func (o *EnumeratePrivateImagesVMCForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enumerate private images Vm c forbidden response has a 4xx status code
func (o *EnumeratePrivateImagesVMCForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this enumerate private images Vm c forbidden response has a 5xx status code
func (o *EnumeratePrivateImagesVMCForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this enumerate private images Vm c forbidden response a status code equal to that given
func (o *EnumeratePrivateImagesVMCForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EnumeratePrivateImagesVMCForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc/{id}/private-image-enumeration][%d] enumeratePrivateImagesVmCForbidden  %+v", 403, o.Payload)
}

func (o *EnumeratePrivateImagesVMCForbidden) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc/{id}/private-image-enumeration][%d] enumeratePrivateImagesVmCForbidden  %+v", 403, o.Payload)
}

func (o *EnumeratePrivateImagesVMCForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *EnumeratePrivateImagesVMCForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
