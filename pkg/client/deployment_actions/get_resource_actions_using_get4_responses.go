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

// GetResourceActionsUsingGET4Reader is a Reader for the GetResourceActionsUsingGET4 structure.
type GetResourceActionsUsingGET4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceActionsUsingGET4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceActionsUsingGET4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourceActionsUsingGET4Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResourceActionsUsingGET4NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourceActionsUsingGET4OK creates a GetResourceActionsUsingGET4OK with default headers values
func NewGetResourceActionsUsingGET4OK() *GetResourceActionsUsingGET4OK {
	return &GetResourceActionsUsingGET4OK{}
}

/*
GetResourceActionsUsingGET4OK describes a response with status code 200, with default header values.

OK
*/
type GetResourceActionsUsingGET4OK struct {
	Payload []*models.ResourceAction
}

// IsSuccess returns true when this get resource actions using g e t4 o k response has a 2xx status code
func (o *GetResourceActionsUsingGET4OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resource actions using g e t4 o k response has a 3xx status code
func (o *GetResourceActionsUsingGET4OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource actions using g e t4 o k response has a 4xx status code
func (o *GetResourceActionsUsingGET4OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resource actions using g e t4 o k response has a 5xx status code
func (o *GetResourceActionsUsingGET4OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource actions using g e t4 o k response a status code equal to that given
func (o *GetResourceActionsUsingGET4OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetResourceActionsUsingGET4OK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources/{resourceId}/actions][%d] getResourceActionsUsingGET4OK  %+v", 200, o.Payload)
}

func (o *GetResourceActionsUsingGET4OK) String() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources/{resourceId}/actions][%d] getResourceActionsUsingGET4OK  %+v", 200, o.Payload)
}

func (o *GetResourceActionsUsingGET4OK) GetPayload() []*models.ResourceAction {
	return o.Payload
}

func (o *GetResourceActionsUsingGET4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceActionsUsingGET4Unauthorized creates a GetResourceActionsUsingGET4Unauthorized with default headers values
func NewGetResourceActionsUsingGET4Unauthorized() *GetResourceActionsUsingGET4Unauthorized {
	return &GetResourceActionsUsingGET4Unauthorized{}
}

/*
GetResourceActionsUsingGET4Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetResourceActionsUsingGET4Unauthorized struct {
}

// IsSuccess returns true when this get resource actions using g e t4 unauthorized response has a 2xx status code
func (o *GetResourceActionsUsingGET4Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource actions using g e t4 unauthorized response has a 3xx status code
func (o *GetResourceActionsUsingGET4Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource actions using g e t4 unauthorized response has a 4xx status code
func (o *GetResourceActionsUsingGET4Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource actions using g e t4 unauthorized response has a 5xx status code
func (o *GetResourceActionsUsingGET4Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource actions using g e t4 unauthorized response a status code equal to that given
func (o *GetResourceActionsUsingGET4Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetResourceActionsUsingGET4Unauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources/{resourceId}/actions][%d] getResourceActionsUsingGET4Unauthorized ", 401)
}

func (o *GetResourceActionsUsingGET4Unauthorized) String() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources/{resourceId}/actions][%d] getResourceActionsUsingGET4Unauthorized ", 401)
}

func (o *GetResourceActionsUsingGET4Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourceActionsUsingGET4NotFound creates a GetResourceActionsUsingGET4NotFound with default headers values
func NewGetResourceActionsUsingGET4NotFound() *GetResourceActionsUsingGET4NotFound {
	return &GetResourceActionsUsingGET4NotFound{}
}

/*
GetResourceActionsUsingGET4NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetResourceActionsUsingGET4NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get resource actions using g e t4 not found response has a 2xx status code
func (o *GetResourceActionsUsingGET4NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource actions using g e t4 not found response has a 3xx status code
func (o *GetResourceActionsUsingGET4NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource actions using g e t4 not found response has a 4xx status code
func (o *GetResourceActionsUsingGET4NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource actions using g e t4 not found response has a 5xx status code
func (o *GetResourceActionsUsingGET4NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource actions using g e t4 not found response a status code equal to that given
func (o *GetResourceActionsUsingGET4NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetResourceActionsUsingGET4NotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources/{resourceId}/actions][%d] getResourceActionsUsingGET4NotFound  %+v", 404, o.Payload)
}

func (o *GetResourceActionsUsingGET4NotFound) String() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources/{resourceId}/actions][%d] getResourceActionsUsingGET4NotFound  %+v", 404, o.Payload)
}

func (o *GetResourceActionsUsingGET4NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetResourceActionsUsingGET4NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}