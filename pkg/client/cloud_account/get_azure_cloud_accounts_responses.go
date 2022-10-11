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

// GetAzureCloudAccountsReader is a Reader for the GetAzureCloudAccounts structure.
type GetAzureCloudAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAzureCloudAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAzureCloudAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAzureCloudAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAzureCloudAccountsOK creates a GetAzureCloudAccountsOK with default headers values
func NewGetAzureCloudAccountsOK() *GetAzureCloudAccountsOK {
	return &GetAzureCloudAccountsOK{}
}

/*
GetAzureCloudAccountsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAzureCloudAccountsOK struct {
	Payload *models.CloudAccountAzureResult
}

// IsSuccess returns true when this get azure cloud accounts o k response has a 2xx status code
func (o *GetAzureCloudAccountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get azure cloud accounts o k response has a 3xx status code
func (o *GetAzureCloudAccountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure cloud accounts o k response has a 4xx status code
func (o *GetAzureCloudAccountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get azure cloud accounts o k response has a 5xx status code
func (o *GetAzureCloudAccountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get azure cloud accounts o k response a status code equal to that given
func (o *GetAzureCloudAccountsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAzureCloudAccountsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-azure][%d] getAzureCloudAccountsOK  %+v", 200, o.Payload)
}

func (o *GetAzureCloudAccountsOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-azure][%d] getAzureCloudAccountsOK  %+v", 200, o.Payload)
}

func (o *GetAzureCloudAccountsOK) GetPayload() *models.CloudAccountAzureResult {
	return o.Payload
}

func (o *GetAzureCloudAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountAzureResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAzureCloudAccountsForbidden creates a GetAzureCloudAccountsForbidden with default headers values
func NewGetAzureCloudAccountsForbidden() *GetAzureCloudAccountsForbidden {
	return &GetAzureCloudAccountsForbidden{}
}

/*
GetAzureCloudAccountsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAzureCloudAccountsForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get azure cloud accounts forbidden response has a 2xx status code
func (o *GetAzureCloudAccountsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get azure cloud accounts forbidden response has a 3xx status code
func (o *GetAzureCloudAccountsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get azure cloud accounts forbidden response has a 4xx status code
func (o *GetAzureCloudAccountsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get azure cloud accounts forbidden response has a 5xx status code
func (o *GetAzureCloudAccountsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get azure cloud accounts forbidden response a status code equal to that given
func (o *GetAzureCloudAccountsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAzureCloudAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-azure][%d] getAzureCloudAccountsForbidden  %+v", 403, o.Payload)
}

func (o *GetAzureCloudAccountsForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-azure][%d] getAzureCloudAccountsForbidden  %+v", 403, o.Payload)
}

func (o *GetAzureCloudAccountsForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetAzureCloudAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
