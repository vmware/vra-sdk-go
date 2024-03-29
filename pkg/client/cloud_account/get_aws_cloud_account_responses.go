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

// GetAwsCloudAccountReader is a Reader for the GetAwsCloudAccount structure.
type GetAwsCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAwsCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAwsCloudAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAwsCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAwsCloudAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAwsCloudAccountOK creates a GetAwsCloudAccountOK with default headers values
func NewGetAwsCloudAccountOK() *GetAwsCloudAccountOK {
	return &GetAwsCloudAccountOK{}
}

/*
GetAwsCloudAccountOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAwsCloudAccountOK struct {
	Payload *models.CloudAccountAws
}

// IsSuccess returns true when this get aws cloud account o k response has a 2xx status code
func (o *GetAwsCloudAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get aws cloud account o k response has a 3xx status code
func (o *GetAwsCloudAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get aws cloud account o k response has a 4xx status code
func (o *GetAwsCloudAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get aws cloud account o k response has a 5xx status code
func (o *GetAwsCloudAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get aws cloud account o k response a status code equal to that given
func (o *GetAwsCloudAccountOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAwsCloudAccountOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-aws/{id}][%d] getAwsCloudAccountOK  %+v", 200, o.Payload)
}

func (o *GetAwsCloudAccountOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-aws/{id}][%d] getAwsCloudAccountOK  %+v", 200, o.Payload)
}

func (o *GetAwsCloudAccountOK) GetPayload() *models.CloudAccountAws {
	return o.Payload
}

func (o *GetAwsCloudAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountAws)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAwsCloudAccountForbidden creates a GetAwsCloudAccountForbidden with default headers values
func NewGetAwsCloudAccountForbidden() *GetAwsCloudAccountForbidden {
	return &GetAwsCloudAccountForbidden{}
}

/*
GetAwsCloudAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAwsCloudAccountForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get aws cloud account forbidden response has a 2xx status code
func (o *GetAwsCloudAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get aws cloud account forbidden response has a 3xx status code
func (o *GetAwsCloudAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get aws cloud account forbidden response has a 4xx status code
func (o *GetAwsCloudAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get aws cloud account forbidden response has a 5xx status code
func (o *GetAwsCloudAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get aws cloud account forbidden response a status code equal to that given
func (o *GetAwsCloudAccountForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAwsCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-aws/{id}][%d] getAwsCloudAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetAwsCloudAccountForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-aws/{id}][%d] getAwsCloudAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetAwsCloudAccountForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetAwsCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAwsCloudAccountNotFound creates a GetAwsCloudAccountNotFound with default headers values
func NewGetAwsCloudAccountNotFound() *GetAwsCloudAccountNotFound {
	return &GetAwsCloudAccountNotFound{}
}

/*
GetAwsCloudAccountNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAwsCloudAccountNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get aws cloud account not found response has a 2xx status code
func (o *GetAwsCloudAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get aws cloud account not found response has a 3xx status code
func (o *GetAwsCloudAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get aws cloud account not found response has a 4xx status code
func (o *GetAwsCloudAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get aws cloud account not found response has a 5xx status code
func (o *GetAwsCloudAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get aws cloud account not found response a status code equal to that given
func (o *GetAwsCloudAccountNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAwsCloudAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-aws/{id}][%d] getAwsCloudAccountNotFound  %+v", 404, o.Payload)
}

func (o *GetAwsCloudAccountNotFound) String() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-aws/{id}][%d] getAwsCloudAccountNotFound  %+v", 404, o.Payload)
}

func (o *GetAwsCloudAccountNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAwsCloudAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
