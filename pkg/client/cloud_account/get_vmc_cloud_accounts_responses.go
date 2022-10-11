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

// GetVmcCloudAccountsReader is a Reader for the GetVmcCloudAccounts structure.
type GetVmcCloudAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVmcCloudAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVmcCloudAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetVmcCloudAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVmcCloudAccountsOK creates a GetVmcCloudAccountsOK with default headers values
func NewGetVmcCloudAccountsOK() *GetVmcCloudAccountsOK {
	return &GetVmcCloudAccountsOK{}
}

/*
GetVmcCloudAccountsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetVmcCloudAccountsOK struct {
	Payload *models.CloudAccountVmcResult
}

// IsSuccess returns true when this get vmc cloud accounts o k response has a 2xx status code
func (o *GetVmcCloudAccountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vmc cloud accounts o k response has a 3xx status code
func (o *GetVmcCloudAccountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vmc cloud accounts o k response has a 4xx status code
func (o *GetVmcCloudAccountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vmc cloud accounts o k response has a 5xx status code
func (o *GetVmcCloudAccountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vmc cloud accounts o k response a status code equal to that given
func (o *GetVmcCloudAccountsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetVmcCloudAccountsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-vmc][%d] getVmcCloudAccountsOK  %+v", 200, o.Payload)
}

func (o *GetVmcCloudAccountsOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-vmc][%d] getVmcCloudAccountsOK  %+v", 200, o.Payload)
}

func (o *GetVmcCloudAccountsOK) GetPayload() *models.CloudAccountVmcResult {
	return o.Payload
}

func (o *GetVmcCloudAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountVmcResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVmcCloudAccountsForbidden creates a GetVmcCloudAccountsForbidden with default headers values
func NewGetVmcCloudAccountsForbidden() *GetVmcCloudAccountsForbidden {
	return &GetVmcCloudAccountsForbidden{}
}

/*
GetVmcCloudAccountsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetVmcCloudAccountsForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get vmc cloud accounts forbidden response has a 2xx status code
func (o *GetVmcCloudAccountsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get vmc cloud accounts forbidden response has a 3xx status code
func (o *GetVmcCloudAccountsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vmc cloud accounts forbidden response has a 4xx status code
func (o *GetVmcCloudAccountsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get vmc cloud accounts forbidden response has a 5xx status code
func (o *GetVmcCloudAccountsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get vmc cloud accounts forbidden response a status code equal to that given
func (o *GetVmcCloudAccountsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetVmcCloudAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-vmc][%d] getVmcCloudAccountsForbidden  %+v", 403, o.Payload)
}

func (o *GetVmcCloudAccountsForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-vmc][%d] getVmcCloudAccountsForbidden  %+v", 403, o.Payload)
}

func (o *GetVmcCloudAccountsForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetVmcCloudAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
