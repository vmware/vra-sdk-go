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

// GetDeploymentExpenseHistoryByIDUsingGET2Reader is a Reader for the GetDeploymentExpenseHistoryByIDUsingGET2 structure.
type GetDeploymentExpenseHistoryByIDUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentExpenseHistoryByIDUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentExpenseHistoryByIDUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeploymentExpenseHistoryByIDUsingGET2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeploymentExpenseHistoryByIDUsingGET2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentExpenseHistoryByIDUsingGET2OK creates a GetDeploymentExpenseHistoryByIDUsingGET2OK with default headers values
func NewGetDeploymentExpenseHistoryByIDUsingGET2OK() *GetDeploymentExpenseHistoryByIDUsingGET2OK {
	return &GetDeploymentExpenseHistoryByIDUsingGET2OK{}
}

/*
GetDeploymentExpenseHistoryByIDUsingGET2OK describes a response with status code 200, with default header values.

OK
*/
type GetDeploymentExpenseHistoryByIDUsingGET2OK struct {
	Payload *models.DeploymentExpenseHistory
}

// IsSuccess returns true when this get deployment expense history by Id using g e t2 o k response has a 2xx status code
func (o *GetDeploymentExpenseHistoryByIDUsingGET2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get deployment expense history by Id using g e t2 o k response has a 3xx status code
func (o *GetDeploymentExpenseHistoryByIDUsingGET2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment expense history by Id using g e t2 o k response has a 4xx status code
func (o *GetDeploymentExpenseHistoryByIDUsingGET2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployment expense history by Id using g e t2 o k response has a 5xx status code
func (o *GetDeploymentExpenseHistoryByIDUsingGET2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment expense history by Id using g e t2 o k response a status code equal to that given
func (o *GetDeploymentExpenseHistoryByIDUsingGET2OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDeploymentExpenseHistoryByIDUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/expense-history][%d] getDeploymentExpenseHistoryByIdUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetDeploymentExpenseHistoryByIDUsingGET2OK) String() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/expense-history][%d] getDeploymentExpenseHistoryByIdUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetDeploymentExpenseHistoryByIDUsingGET2OK) GetPayload() *models.DeploymentExpenseHistory {
	return o.Payload
}

func (o *GetDeploymentExpenseHistoryByIDUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentExpenseHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentExpenseHistoryByIDUsingGET2Unauthorized creates a GetDeploymentExpenseHistoryByIDUsingGET2Unauthorized with default headers values
func NewGetDeploymentExpenseHistoryByIDUsingGET2Unauthorized() *GetDeploymentExpenseHistoryByIDUsingGET2Unauthorized {
	return &GetDeploymentExpenseHistoryByIDUsingGET2Unauthorized{}
}

/*
GetDeploymentExpenseHistoryByIDUsingGET2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDeploymentExpenseHistoryByIDUsingGET2Unauthorized struct {
}

// IsSuccess returns true when this get deployment expense history by Id using g e t2 unauthorized response has a 2xx status code
func (o *GetDeploymentExpenseHistoryByIDUsingGET2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment expense history by Id using g e t2 unauthorized response has a 3xx status code
func (o *GetDeploymentExpenseHistoryByIDUsingGET2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment expense history by Id using g e t2 unauthorized response has a 4xx status code
func (o *GetDeploymentExpenseHistoryByIDUsingGET2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment expense history by Id using g e t2 unauthorized response has a 5xx status code
func (o *GetDeploymentExpenseHistoryByIDUsingGET2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment expense history by Id using g e t2 unauthorized response a status code equal to that given
func (o *GetDeploymentExpenseHistoryByIDUsingGET2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDeploymentExpenseHistoryByIDUsingGET2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/expense-history][%d] getDeploymentExpenseHistoryByIdUsingGET2Unauthorized ", 401)
}

func (o *GetDeploymentExpenseHistoryByIDUsingGET2Unauthorized) String() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/expense-history][%d] getDeploymentExpenseHistoryByIdUsingGET2Unauthorized ", 401)
}

func (o *GetDeploymentExpenseHistoryByIDUsingGET2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentExpenseHistoryByIDUsingGET2NotFound creates a GetDeploymentExpenseHistoryByIDUsingGET2NotFound with default headers values
func NewGetDeploymentExpenseHistoryByIDUsingGET2NotFound() *GetDeploymentExpenseHistoryByIDUsingGET2NotFound {
	return &GetDeploymentExpenseHistoryByIDUsingGET2NotFound{}
}

/*
GetDeploymentExpenseHistoryByIDUsingGET2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDeploymentExpenseHistoryByIDUsingGET2NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get deployment expense history by Id using g e t2 not found response has a 2xx status code
func (o *GetDeploymentExpenseHistoryByIDUsingGET2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment expense history by Id using g e t2 not found response has a 3xx status code
func (o *GetDeploymentExpenseHistoryByIDUsingGET2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment expense history by Id using g e t2 not found response has a 4xx status code
func (o *GetDeploymentExpenseHistoryByIDUsingGET2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment expense history by Id using g e t2 not found response has a 5xx status code
func (o *GetDeploymentExpenseHistoryByIDUsingGET2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment expense history by Id using g e t2 not found response a status code equal to that given
func (o *GetDeploymentExpenseHistoryByIDUsingGET2NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDeploymentExpenseHistoryByIDUsingGET2NotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/expense-history][%d] getDeploymentExpenseHistoryByIdUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentExpenseHistoryByIDUsingGET2NotFound) String() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/expense-history][%d] getDeploymentExpenseHistoryByIdUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentExpenseHistoryByIDUsingGET2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDeploymentExpenseHistoryByIDUsingGET2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
