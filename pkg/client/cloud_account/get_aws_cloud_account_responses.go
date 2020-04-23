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
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAwsCloudAccountOK creates a GetAwsCloudAccountOK with default headers values
func NewGetAwsCloudAccountOK() *GetAwsCloudAccountOK {
	return &GetAwsCloudAccountOK{}
}

/*GetAwsCloudAccountOK handles this case with default header values.

successful operation
*/
type GetAwsCloudAccountOK struct {
	Payload *models.CloudAccountAws
}

func (o *GetAwsCloudAccountOK) Error() string {
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

/*GetAwsCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type GetAwsCloudAccountForbidden struct {
}

func (o *GetAwsCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-aws/{id}][%d] getAwsCloudAccountForbidden ", 403)
}

func (o *GetAwsCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAwsCloudAccountNotFound creates a GetAwsCloudAccountNotFound with default headers values
func NewGetAwsCloudAccountNotFound() *GetAwsCloudAccountNotFound {
	return &GetAwsCloudAccountNotFound{}
}

/*GetAwsCloudAccountNotFound handles this case with default header values.

Not Found
*/
type GetAwsCloudAccountNotFound struct {
	Payload *models.Error
}

func (o *GetAwsCloudAccountNotFound) Error() string {
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
