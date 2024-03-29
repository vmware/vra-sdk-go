// Code generated by go-swagger; DO NOT EDIT.

package fabric_azure_storage_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetFabricAzureStorageAccountReader is a Reader for the GetFabricAzureStorageAccount structure.
type GetFabricAzureStorageAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFabricAzureStorageAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFabricAzureStorageAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFabricAzureStorageAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFabricAzureStorageAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFabricAzureStorageAccountOK creates a GetFabricAzureStorageAccountOK with default headers values
func NewGetFabricAzureStorageAccountOK() *GetFabricAzureStorageAccountOK {
	return &GetFabricAzureStorageAccountOK{}
}

/*
GetFabricAzureStorageAccountOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFabricAzureStorageAccountOK struct {
	Payload *models.FabricAzureStorageAccount
}

// IsSuccess returns true when this get fabric azure storage account o k response has a 2xx status code
func (o *GetFabricAzureStorageAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get fabric azure storage account o k response has a 3xx status code
func (o *GetFabricAzureStorageAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fabric azure storage account o k response has a 4xx status code
func (o *GetFabricAzureStorageAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fabric azure storage account o k response has a 5xx status code
func (o *GetFabricAzureStorageAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get fabric azure storage account o k response a status code equal to that given
func (o *GetFabricAzureStorageAccountOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFabricAzureStorageAccountOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-storage-accounts/{id}][%d] getFabricAzureStorageAccountOK  %+v", 200, o.Payload)
}

func (o *GetFabricAzureStorageAccountOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-storage-accounts/{id}][%d] getFabricAzureStorageAccountOK  %+v", 200, o.Payload)
}

func (o *GetFabricAzureStorageAccountOK) GetPayload() *models.FabricAzureStorageAccount {
	return o.Payload
}

func (o *GetFabricAzureStorageAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricAzureStorageAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricAzureStorageAccountForbidden creates a GetFabricAzureStorageAccountForbidden with default headers values
func NewGetFabricAzureStorageAccountForbidden() *GetFabricAzureStorageAccountForbidden {
	return &GetFabricAzureStorageAccountForbidden{}
}

/*
GetFabricAzureStorageAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFabricAzureStorageAccountForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get fabric azure storage account forbidden response has a 2xx status code
func (o *GetFabricAzureStorageAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fabric azure storage account forbidden response has a 3xx status code
func (o *GetFabricAzureStorageAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fabric azure storage account forbidden response has a 4xx status code
func (o *GetFabricAzureStorageAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fabric azure storage account forbidden response has a 5xx status code
func (o *GetFabricAzureStorageAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get fabric azure storage account forbidden response a status code equal to that given
func (o *GetFabricAzureStorageAccountForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFabricAzureStorageAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-storage-accounts/{id}][%d] getFabricAzureStorageAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetFabricAzureStorageAccountForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-storage-accounts/{id}][%d] getFabricAzureStorageAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetFabricAzureStorageAccountForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetFabricAzureStorageAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricAzureStorageAccountNotFound creates a GetFabricAzureStorageAccountNotFound with default headers values
func NewGetFabricAzureStorageAccountNotFound() *GetFabricAzureStorageAccountNotFound {
	return &GetFabricAzureStorageAccountNotFound{}
}

/*
GetFabricAzureStorageAccountNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetFabricAzureStorageAccountNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get fabric azure storage account not found response has a 2xx status code
func (o *GetFabricAzureStorageAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fabric azure storage account not found response has a 3xx status code
func (o *GetFabricAzureStorageAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fabric azure storage account not found response has a 4xx status code
func (o *GetFabricAzureStorageAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fabric azure storage account not found response has a 5xx status code
func (o *GetFabricAzureStorageAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get fabric azure storage account not found response a status code equal to that given
func (o *GetFabricAzureStorageAccountNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFabricAzureStorageAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-storage-accounts/{id}][%d] getFabricAzureStorageAccountNotFound  %+v", 404, o.Payload)
}

func (o *GetFabricAzureStorageAccountNotFound) String() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-storage-accounts/{id}][%d] getFabricAzureStorageAccountNotFound  %+v", 404, o.Payload)
}

func (o *GetFabricAzureStorageAccountNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFabricAzureStorageAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
