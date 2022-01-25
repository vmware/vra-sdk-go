// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDeploymentByIDV3UsingGETReader is a Reader for the GetDeploymentByIDV3UsingGET structure.
type GetDeploymentByIDV3UsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentByIDV3UsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentByIDV3UsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeploymentByIDV3UsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeploymentByIDV3UsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentByIDV3UsingGETOK creates a GetDeploymentByIDV3UsingGETOK with default headers values
func NewGetDeploymentByIDV3UsingGETOK() *GetDeploymentByIDV3UsingGETOK {
	return &GetDeploymentByIDV3UsingGETOK{}
}

/* GetDeploymentByIDV3UsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetDeploymentByIDV3UsingGETOK struct {
	Payload *models.Deployment
}

func (o *GetDeploymentByIDV3UsingGETOK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}][%d] getDeploymentByIdV3UsingGETOK  %+v", 200, o.Payload)
}
func (o *GetDeploymentByIDV3UsingGETOK) GetPayload() *models.Deployment {
	return o.Payload
}

func (o *GetDeploymentByIDV3UsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentByIDV3UsingGETUnauthorized creates a GetDeploymentByIDV3UsingGETUnauthorized with default headers values
func NewGetDeploymentByIDV3UsingGETUnauthorized() *GetDeploymentByIDV3UsingGETUnauthorized {
	return &GetDeploymentByIDV3UsingGETUnauthorized{}
}

/* GetDeploymentByIDV3UsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDeploymentByIDV3UsingGETUnauthorized struct {
}

func (o *GetDeploymentByIDV3UsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}][%d] getDeploymentByIdV3UsingGETUnauthorized ", 401)
}

func (o *GetDeploymentByIDV3UsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentByIDV3UsingGETNotFound creates a GetDeploymentByIDV3UsingGETNotFound with default headers values
func NewGetDeploymentByIDV3UsingGETNotFound() *GetDeploymentByIDV3UsingGETNotFound {
	return &GetDeploymentByIDV3UsingGETNotFound{}
}

/* GetDeploymentByIDV3UsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDeploymentByIDV3UsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetDeploymentByIDV3UsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}][%d] getDeploymentByIdV3UsingGETNotFound  %+v", 404, o.Payload)
}
func (o *GetDeploymentByIDV3UsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDeploymentByIDV3UsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
