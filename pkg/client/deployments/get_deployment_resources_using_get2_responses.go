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

// GetDeploymentResourcesUsingGET2Reader is a Reader for the GetDeploymentResourcesUsingGET2 structure.
type GetDeploymentResourcesUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentResourcesUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentResourcesUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeploymentResourcesUsingGET2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeploymentResourcesUsingGET2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentResourcesUsingGET2OK creates a GetDeploymentResourcesUsingGET2OK with default headers values
func NewGetDeploymentResourcesUsingGET2OK() *GetDeploymentResourcesUsingGET2OK {
	return &GetDeploymentResourcesUsingGET2OK{}
}

/*
GetDeploymentResourcesUsingGET2OK describes a response with status code 200, with default header values.

OK
*/
type GetDeploymentResourcesUsingGET2OK struct {
	Payload *models.PageOfDeploymentResource
}

// IsSuccess returns true when this get deployment resources using g e t2 o k response has a 2xx status code
func (o *GetDeploymentResourcesUsingGET2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get deployment resources using g e t2 o k response has a 3xx status code
func (o *GetDeploymentResourcesUsingGET2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment resources using g e t2 o k response has a 4xx status code
func (o *GetDeploymentResourcesUsingGET2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployment resources using g e t2 o k response has a 5xx status code
func (o *GetDeploymentResourcesUsingGET2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment resources using g e t2 o k response a status code equal to that given
func (o *GetDeploymentResourcesUsingGET2OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDeploymentResourcesUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources][%d] getDeploymentResourcesUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetDeploymentResourcesUsingGET2OK) String() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources][%d] getDeploymentResourcesUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetDeploymentResourcesUsingGET2OK) GetPayload() *models.PageOfDeploymentResource {
	return o.Payload
}

func (o *GetDeploymentResourcesUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfDeploymentResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentResourcesUsingGET2Unauthorized creates a GetDeploymentResourcesUsingGET2Unauthorized with default headers values
func NewGetDeploymentResourcesUsingGET2Unauthorized() *GetDeploymentResourcesUsingGET2Unauthorized {
	return &GetDeploymentResourcesUsingGET2Unauthorized{}
}

/*
GetDeploymentResourcesUsingGET2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDeploymentResourcesUsingGET2Unauthorized struct {
}

// IsSuccess returns true when this get deployment resources using g e t2 unauthorized response has a 2xx status code
func (o *GetDeploymentResourcesUsingGET2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment resources using g e t2 unauthorized response has a 3xx status code
func (o *GetDeploymentResourcesUsingGET2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment resources using g e t2 unauthorized response has a 4xx status code
func (o *GetDeploymentResourcesUsingGET2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment resources using g e t2 unauthorized response has a 5xx status code
func (o *GetDeploymentResourcesUsingGET2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment resources using g e t2 unauthorized response a status code equal to that given
func (o *GetDeploymentResourcesUsingGET2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDeploymentResourcesUsingGET2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources][%d] getDeploymentResourcesUsingGET2Unauthorized ", 401)
}

func (o *GetDeploymentResourcesUsingGET2Unauthorized) String() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources][%d] getDeploymentResourcesUsingGET2Unauthorized ", 401)
}

func (o *GetDeploymentResourcesUsingGET2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentResourcesUsingGET2NotFound creates a GetDeploymentResourcesUsingGET2NotFound with default headers values
func NewGetDeploymentResourcesUsingGET2NotFound() *GetDeploymentResourcesUsingGET2NotFound {
	return &GetDeploymentResourcesUsingGET2NotFound{}
}

/*
GetDeploymentResourcesUsingGET2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDeploymentResourcesUsingGET2NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get deployment resources using g e t2 not found response has a 2xx status code
func (o *GetDeploymentResourcesUsingGET2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment resources using g e t2 not found response has a 3xx status code
func (o *GetDeploymentResourcesUsingGET2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment resources using g e t2 not found response has a 4xx status code
func (o *GetDeploymentResourcesUsingGET2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment resources using g e t2 not found response has a 5xx status code
func (o *GetDeploymentResourcesUsingGET2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment resources using g e t2 not found response a status code equal to that given
func (o *GetDeploymentResourcesUsingGET2NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDeploymentResourcesUsingGET2NotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources][%d] getDeploymentResourcesUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentResourcesUsingGET2NotFound) String() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources][%d] getDeploymentResourcesUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentResourcesUsingGET2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDeploymentResourcesUsingGET2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
