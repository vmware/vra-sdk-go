// Code generated by go-swagger; DO NOT EDIT.

package deployment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetSingleDeploymentReader is a Reader for the GetSingleDeployment structure.
type GetSingleDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSingleDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSingleDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSingleDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSingleDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSingleDeploymentOK creates a GetSingleDeploymentOK with default headers values
func NewGetSingleDeploymentOK() *GetSingleDeploymentOK {
	return &GetSingleDeploymentOK{}
}

/*GetSingleDeploymentOK handles this case with default header values.

successful operation
*/
type GetSingleDeploymentOK struct {
	Payload *models.IaaSDeployment
}

func (o *GetSingleDeploymentOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/deployments/{id}][%d] getSingleDeploymentOK  %+v", 200, o.Payload)
}

func (o *GetSingleDeploymentOK) GetPayload() *models.IaaSDeployment {
	return o.Payload
}

func (o *GetSingleDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IaaSDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSingleDeploymentForbidden creates a GetSingleDeploymentForbidden with default headers values
func NewGetSingleDeploymentForbidden() *GetSingleDeploymentForbidden {
	return &GetSingleDeploymentForbidden{}
}

/*GetSingleDeploymentForbidden handles this case with default header values.

Forbidden
*/
type GetSingleDeploymentForbidden struct {
}

func (o *GetSingleDeploymentForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/deployments/{id}][%d] getSingleDeploymentForbidden ", 403)
}

func (o *GetSingleDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSingleDeploymentNotFound creates a GetSingleDeploymentNotFound with default headers values
func NewGetSingleDeploymentNotFound() *GetSingleDeploymentNotFound {
	return &GetSingleDeploymentNotFound{}
}

/*GetSingleDeploymentNotFound handles this case with default header values.

Not Found
*/
type GetSingleDeploymentNotFound struct {
}

func (o *GetSingleDeploymentNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/deployments/{id}][%d] getSingleDeploymentNotFound ", 404)
}

func (o *GetSingleDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}