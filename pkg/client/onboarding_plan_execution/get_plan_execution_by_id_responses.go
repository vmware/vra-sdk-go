// Code generated by go-swagger; DO NOT EDIT.

package onboarding_plan_execution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetPlanExecutionByIDReader is a Reader for the GetPlanExecutionByID structure.
type GetPlanExecutionByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlanExecutionByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPlanExecutionByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPlanExecutionByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPlanExecutionByIDOK creates a GetPlanExecutionByIDOK with default headers values
func NewGetPlanExecutionByIDOK() *GetPlanExecutionByIDOK {
	return &GetPlanExecutionByIDOK{}
}

/*
GetPlanExecutionByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetPlanExecutionByIDOK struct {
	Payload *models.PlanExecutionResponse
}

// IsSuccess returns true when this get plan execution by Id o k response has a 2xx status code
func (o *GetPlanExecutionByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get plan execution by Id o k response has a 3xx status code
func (o *GetPlanExecutionByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get plan execution by Id o k response has a 4xx status code
func (o *GetPlanExecutionByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get plan execution by Id o k response has a 5xx status code
func (o *GetPlanExecutionByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get plan execution by Id o k response a status code equal to that given
func (o *GetPlanExecutionByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPlanExecutionByIDOK) Error() string {
	return fmt.Sprintf("[GET /relocation/api/wo/execute-plan/{id}][%d] getPlanExecutionByIdOK  %+v", 200, o.Payload)
}

func (o *GetPlanExecutionByIDOK) String() string {
	return fmt.Sprintf("[GET /relocation/api/wo/execute-plan/{id}][%d] getPlanExecutionByIdOK  %+v", 200, o.Payload)
}

func (o *GetPlanExecutionByIDOK) GetPayload() *models.PlanExecutionResponse {
	return o.Payload
}

func (o *GetPlanExecutionByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlanExecutionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlanExecutionByIDUnauthorized creates a GetPlanExecutionByIDUnauthorized with default headers values
func NewGetPlanExecutionByIDUnauthorized() *GetPlanExecutionByIDUnauthorized {
	return &GetPlanExecutionByIDUnauthorized{}
}

/*
GetPlanExecutionByIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPlanExecutionByIDUnauthorized struct {
}

// IsSuccess returns true when this get plan execution by Id unauthorized response has a 2xx status code
func (o *GetPlanExecutionByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get plan execution by Id unauthorized response has a 3xx status code
func (o *GetPlanExecutionByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get plan execution by Id unauthorized response has a 4xx status code
func (o *GetPlanExecutionByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get plan execution by Id unauthorized response has a 5xx status code
func (o *GetPlanExecutionByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get plan execution by Id unauthorized response a status code equal to that given
func (o *GetPlanExecutionByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPlanExecutionByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /relocation/api/wo/execute-plan/{id}][%d] getPlanExecutionByIdUnauthorized ", 401)
}

func (o *GetPlanExecutionByIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /relocation/api/wo/execute-plan/{id}][%d] getPlanExecutionByIdUnauthorized ", 401)
}

func (o *GetPlanExecutionByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
