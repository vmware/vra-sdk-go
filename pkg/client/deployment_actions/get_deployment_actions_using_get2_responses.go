// Code generated by go-swagger; DO NOT EDIT.

package deployment_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDeploymentActionsUsingGET2Reader is a Reader for the GetDeploymentActionsUsingGET2 structure.
type GetDeploymentActionsUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentActionsUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentActionsUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeploymentActionsUsingGET2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeploymentActionsUsingGET2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentActionsUsingGET2OK creates a GetDeploymentActionsUsingGET2OK with default headers values
func NewGetDeploymentActionsUsingGET2OK() *GetDeploymentActionsUsingGET2OK {
	return &GetDeploymentActionsUsingGET2OK{}
}

/*
GetDeploymentActionsUsingGET2OK describes a response with status code 200, with default header values.

OK
*/
type GetDeploymentActionsUsingGET2OK struct {
	Payload []*models.ResourceAction
}

// IsSuccess returns true when this get deployment actions using g e t2 o k response has a 2xx status code
func (o *GetDeploymentActionsUsingGET2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get deployment actions using g e t2 o k response has a 3xx status code
func (o *GetDeploymentActionsUsingGET2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment actions using g e t2 o k response has a 4xx status code
func (o *GetDeploymentActionsUsingGET2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployment actions using g e t2 o k response has a 5xx status code
func (o *GetDeploymentActionsUsingGET2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment actions using g e t2 o k response a status code equal to that given
func (o *GetDeploymentActionsUsingGET2OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDeploymentActionsUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/actions][%d] getDeploymentActionsUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetDeploymentActionsUsingGET2OK) String() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/actions][%d] getDeploymentActionsUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetDeploymentActionsUsingGET2OK) GetPayload() []*models.ResourceAction {
	return o.Payload
}

func (o *GetDeploymentActionsUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentActionsUsingGET2Unauthorized creates a GetDeploymentActionsUsingGET2Unauthorized with default headers values
func NewGetDeploymentActionsUsingGET2Unauthorized() *GetDeploymentActionsUsingGET2Unauthorized {
	return &GetDeploymentActionsUsingGET2Unauthorized{}
}

/*
GetDeploymentActionsUsingGET2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDeploymentActionsUsingGET2Unauthorized struct {
}

// IsSuccess returns true when this get deployment actions using g e t2 unauthorized response has a 2xx status code
func (o *GetDeploymentActionsUsingGET2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment actions using g e t2 unauthorized response has a 3xx status code
func (o *GetDeploymentActionsUsingGET2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment actions using g e t2 unauthorized response has a 4xx status code
func (o *GetDeploymentActionsUsingGET2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment actions using g e t2 unauthorized response has a 5xx status code
func (o *GetDeploymentActionsUsingGET2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment actions using g e t2 unauthorized response a status code equal to that given
func (o *GetDeploymentActionsUsingGET2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDeploymentActionsUsingGET2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/actions][%d] getDeploymentActionsUsingGET2Unauthorized ", 401)
}

func (o *GetDeploymentActionsUsingGET2Unauthorized) String() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/actions][%d] getDeploymentActionsUsingGET2Unauthorized ", 401)
}

func (o *GetDeploymentActionsUsingGET2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentActionsUsingGET2NotFound creates a GetDeploymentActionsUsingGET2NotFound with default headers values
func NewGetDeploymentActionsUsingGET2NotFound() *GetDeploymentActionsUsingGET2NotFound {
	return &GetDeploymentActionsUsingGET2NotFound{}
}

/*
GetDeploymentActionsUsingGET2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDeploymentActionsUsingGET2NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get deployment actions using g e t2 not found response has a 2xx status code
func (o *GetDeploymentActionsUsingGET2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment actions using g e t2 not found response has a 3xx status code
func (o *GetDeploymentActionsUsingGET2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment actions using g e t2 not found response has a 4xx status code
func (o *GetDeploymentActionsUsingGET2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment actions using g e t2 not found response has a 5xx status code
func (o *GetDeploymentActionsUsingGET2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment actions using g e t2 not found response a status code equal to that given
func (o *GetDeploymentActionsUsingGET2NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDeploymentActionsUsingGET2NotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/actions][%d] getDeploymentActionsUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentActionsUsingGET2NotFound) String() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/actions][%d] getDeploymentActionsUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentActionsUsingGET2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDeploymentActionsUsingGET2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}