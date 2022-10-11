// Code generated by go-swagger; DO NOT EDIT.

package onboarding_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetPlanByIDReader is a Reader for the GetPlanByID structure.
type GetPlanByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlanByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPlanByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPlanByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPlanByIDOK creates a GetPlanByIDOK with default headers values
func NewGetPlanByIDOK() *GetPlanByIDOK {
	return &GetPlanByIDOK{}
}

/*
GetPlanByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetPlanByIDOK struct {
	Payload *models.OnboardingPlanResponse
}

// IsSuccess returns true when this get plan by Id o k response has a 2xx status code
func (o *GetPlanByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get plan by Id o k response has a 3xx status code
func (o *GetPlanByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get plan by Id o k response has a 4xx status code
func (o *GetPlanByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get plan by Id o k response has a 5xx status code
func (o *GetPlanByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get plan by Id o k response a status code equal to that given
func (o *GetPlanByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPlanByIDOK) Error() string {
	return fmt.Sprintf("[GET /relocation/onboarding/plan/{id}][%d] getPlanByIdOK  %+v", 200, o.Payload)
}

func (o *GetPlanByIDOK) String() string {
	return fmt.Sprintf("[GET /relocation/onboarding/plan/{id}][%d] getPlanByIdOK  %+v", 200, o.Payload)
}

func (o *GetPlanByIDOK) GetPayload() *models.OnboardingPlanResponse {
	return o.Payload
}

func (o *GetPlanByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OnboardingPlanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlanByIDUnauthorized creates a GetPlanByIDUnauthorized with default headers values
func NewGetPlanByIDUnauthorized() *GetPlanByIDUnauthorized {
	return &GetPlanByIDUnauthorized{}
}

/*
GetPlanByIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPlanByIDUnauthorized struct {
}

// IsSuccess returns true when this get plan by Id unauthorized response has a 2xx status code
func (o *GetPlanByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get plan by Id unauthorized response has a 3xx status code
func (o *GetPlanByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get plan by Id unauthorized response has a 4xx status code
func (o *GetPlanByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get plan by Id unauthorized response has a 5xx status code
func (o *GetPlanByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get plan by Id unauthorized response a status code equal to that given
func (o *GetPlanByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPlanByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /relocation/onboarding/plan/{id}][%d] getPlanByIdUnauthorized ", 401)
}

func (o *GetPlanByIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /relocation/onboarding/plan/{id}][%d] getPlanByIdUnauthorized ", 401)
}

func (o *GetPlanByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
