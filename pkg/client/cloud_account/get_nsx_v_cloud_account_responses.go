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

// GetNsxVCloudAccountReader is a Reader for the GetNsxVCloudAccount structure.
type GetNsxVCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNsxVCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNsxVCloudAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetNsxVCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNsxVCloudAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNsxVCloudAccountOK creates a GetNsxVCloudAccountOK with default headers values
func NewGetNsxVCloudAccountOK() *GetNsxVCloudAccountOK {
	return &GetNsxVCloudAccountOK{}
}

/*
GetNsxVCloudAccountOK describes a response with status code 200, with default header values.

successful operation
*/
type GetNsxVCloudAccountOK struct {
	Payload *models.CloudAccountNsxV
}

// IsSuccess returns true when this get nsx v cloud account o k response has a 2xx status code
func (o *GetNsxVCloudAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get nsx v cloud account o k response has a 3xx status code
func (o *GetNsxVCloudAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nsx v cloud account o k response has a 4xx status code
func (o *GetNsxVCloudAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get nsx v cloud account o k response has a 5xx status code
func (o *GetNsxVCloudAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get nsx v cloud account o k response a status code equal to that given
func (o *GetNsxVCloudAccountOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNsxVCloudAccountOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-nsx-v/{id}][%d] getNsxVCloudAccountOK  %+v", 200, o.Payload)
}

func (o *GetNsxVCloudAccountOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-nsx-v/{id}][%d] getNsxVCloudAccountOK  %+v", 200, o.Payload)
}

func (o *GetNsxVCloudAccountOK) GetPayload() *models.CloudAccountNsxV {
	return o.Payload
}

func (o *GetNsxVCloudAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountNsxV)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNsxVCloudAccountForbidden creates a GetNsxVCloudAccountForbidden with default headers values
func NewGetNsxVCloudAccountForbidden() *GetNsxVCloudAccountForbidden {
	return &GetNsxVCloudAccountForbidden{}
}

/*
GetNsxVCloudAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetNsxVCloudAccountForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get nsx v cloud account forbidden response has a 2xx status code
func (o *GetNsxVCloudAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get nsx v cloud account forbidden response has a 3xx status code
func (o *GetNsxVCloudAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nsx v cloud account forbidden response has a 4xx status code
func (o *GetNsxVCloudAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get nsx v cloud account forbidden response has a 5xx status code
func (o *GetNsxVCloudAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get nsx v cloud account forbidden response a status code equal to that given
func (o *GetNsxVCloudAccountForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetNsxVCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-nsx-v/{id}][%d] getNsxVCloudAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetNsxVCloudAccountForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-nsx-v/{id}][%d] getNsxVCloudAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetNsxVCloudAccountForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetNsxVCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNsxVCloudAccountNotFound creates a GetNsxVCloudAccountNotFound with default headers values
func NewGetNsxVCloudAccountNotFound() *GetNsxVCloudAccountNotFound {
	return &GetNsxVCloudAccountNotFound{}
}

/*
GetNsxVCloudAccountNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetNsxVCloudAccountNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get nsx v cloud account not found response has a 2xx status code
func (o *GetNsxVCloudAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get nsx v cloud account not found response has a 3xx status code
func (o *GetNsxVCloudAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nsx v cloud account not found response has a 4xx status code
func (o *GetNsxVCloudAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get nsx v cloud account not found response has a 5xx status code
func (o *GetNsxVCloudAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get nsx v cloud account not found response a status code equal to that given
func (o *GetNsxVCloudAccountNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetNsxVCloudAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-nsx-v/{id}][%d] getNsxVCloudAccountNotFound  %+v", 404, o.Payload)
}

func (o *GetNsxVCloudAccountNotFound) String() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-nsx-v/{id}][%d] getNsxVCloudAccountNotFound  %+v", 404, o.Payload)
}

func (o *GetNsxVCloudAccountNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNsxVCloudAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
