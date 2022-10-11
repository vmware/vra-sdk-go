// Code generated by go-swagger; DO NOT EDIT.

package onboarding_blueprints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateBlueprintReader is a Reader for the CreateBlueprint structure.
type CreateBlueprintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBlueprintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBlueprintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateBlueprintUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateBlueprintOK creates a CreateBlueprintOK with default headers values
func NewCreateBlueprintOK() *CreateBlueprintOK {
	return &CreateBlueprintOK{}
}

/*
CreateBlueprintOK describes a response with status code 200, with default header values.

Success
*/
type CreateBlueprintOK struct {
	Payload *models.OnboardingBlueprintResponse
}

// IsSuccess returns true when this create blueprint o k response has a 2xx status code
func (o *CreateBlueprintOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create blueprint o k response has a 3xx status code
func (o *CreateBlueprintOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create blueprint o k response has a 4xx status code
func (o *CreateBlueprintOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create blueprint o k response has a 5xx status code
func (o *CreateBlueprintOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create blueprint o k response a status code equal to that given
func (o *CreateBlueprintOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateBlueprintOK) Error() string {
	return fmt.Sprintf("[POST /relocation/onboarding/blueprint][%d] createBlueprintOK  %+v", 200, o.Payload)
}

func (o *CreateBlueprintOK) String() string {
	return fmt.Sprintf("[POST /relocation/onboarding/blueprint][%d] createBlueprintOK  %+v", 200, o.Payload)
}

func (o *CreateBlueprintOK) GetPayload() *models.OnboardingBlueprintResponse {
	return o.Payload
}

func (o *CreateBlueprintOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OnboardingBlueprintResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBlueprintUnauthorized creates a CreateBlueprintUnauthorized with default headers values
func NewCreateBlueprintUnauthorized() *CreateBlueprintUnauthorized {
	return &CreateBlueprintUnauthorized{}
}

/*
CreateBlueprintUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateBlueprintUnauthorized struct {
}

// IsSuccess returns true when this create blueprint unauthorized response has a 2xx status code
func (o *CreateBlueprintUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create blueprint unauthorized response has a 3xx status code
func (o *CreateBlueprintUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create blueprint unauthorized response has a 4xx status code
func (o *CreateBlueprintUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create blueprint unauthorized response has a 5xx status code
func (o *CreateBlueprintUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create blueprint unauthorized response a status code equal to that given
func (o *CreateBlueprintUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateBlueprintUnauthorized) Error() string {
	return fmt.Sprintf("[POST /relocation/onboarding/blueprint][%d] createBlueprintUnauthorized ", 401)
}

func (o *CreateBlueprintUnauthorized) String() string {
	return fmt.Sprintf("[POST /relocation/onboarding/blueprint][%d] createBlueprintUnauthorized ", 401)
}

func (o *CreateBlueprintUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
