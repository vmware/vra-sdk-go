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

// GetDeploymentFiltersUsingGET2Reader is a Reader for the GetDeploymentFiltersUsingGET2 structure.
type GetDeploymentFiltersUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentFiltersUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentFiltersUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeploymentFiltersUsingGET2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentFiltersUsingGET2OK creates a GetDeploymentFiltersUsingGET2OK with default headers values
func NewGetDeploymentFiltersUsingGET2OK() *GetDeploymentFiltersUsingGET2OK {
	return &GetDeploymentFiltersUsingGET2OK{}
}

/*
GetDeploymentFiltersUsingGET2OK describes a response with status code 200, with default header values.

OK
*/
type GetDeploymentFiltersUsingGET2OK struct {
	Payload *models.DeploymentFilterSchema
}

// IsSuccess returns true when this get deployment filters using g e t2 o k response has a 2xx status code
func (o *GetDeploymentFiltersUsingGET2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get deployment filters using g e t2 o k response has a 3xx status code
func (o *GetDeploymentFiltersUsingGET2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment filters using g e t2 o k response has a 4xx status code
func (o *GetDeploymentFiltersUsingGET2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployment filters using g e t2 o k response has a 5xx status code
func (o *GetDeploymentFiltersUsingGET2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment filters using g e t2 o k response a status code equal to that given
func (o *GetDeploymentFiltersUsingGET2OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDeploymentFiltersUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/filters][%d] getDeploymentFiltersUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetDeploymentFiltersUsingGET2OK) String() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/filters][%d] getDeploymentFiltersUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetDeploymentFiltersUsingGET2OK) GetPayload() *models.DeploymentFilterSchema {
	return o.Payload
}

func (o *GetDeploymentFiltersUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentFilterSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentFiltersUsingGET2Unauthorized creates a GetDeploymentFiltersUsingGET2Unauthorized with default headers values
func NewGetDeploymentFiltersUsingGET2Unauthorized() *GetDeploymentFiltersUsingGET2Unauthorized {
	return &GetDeploymentFiltersUsingGET2Unauthorized{}
}

/*
GetDeploymentFiltersUsingGET2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDeploymentFiltersUsingGET2Unauthorized struct {
}

// IsSuccess returns true when this get deployment filters using g e t2 unauthorized response has a 2xx status code
func (o *GetDeploymentFiltersUsingGET2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment filters using g e t2 unauthorized response has a 3xx status code
func (o *GetDeploymentFiltersUsingGET2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment filters using g e t2 unauthorized response has a 4xx status code
func (o *GetDeploymentFiltersUsingGET2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment filters using g e t2 unauthorized response has a 5xx status code
func (o *GetDeploymentFiltersUsingGET2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment filters using g e t2 unauthorized response a status code equal to that given
func (o *GetDeploymentFiltersUsingGET2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDeploymentFiltersUsingGET2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/filters][%d] getDeploymentFiltersUsingGET2Unauthorized ", 401)
}

func (o *GetDeploymentFiltersUsingGET2Unauthorized) String() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/filters][%d] getDeploymentFiltersUsingGET2Unauthorized ", 401)
}

func (o *GetDeploymentFiltersUsingGET2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
