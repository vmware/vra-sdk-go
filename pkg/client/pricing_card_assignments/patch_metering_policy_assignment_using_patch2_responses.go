// Code generated by go-swagger; DO NOT EDIT.

package pricing_card_assignments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// PatchMeteringPolicyAssignmentUsingPATCH2Reader is a Reader for the PatchMeteringPolicyAssignmentUsingPATCH2 structure.
type PatchMeteringPolicyAssignmentUsingPATCH2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchMeteringPolicyAssignmentUsingPATCH2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchMeteringPolicyAssignmentUsingPATCH2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchMeteringPolicyAssignmentUsingPATCH2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchMeteringPolicyAssignmentUsingPATCH2OK creates a PatchMeteringPolicyAssignmentUsingPATCH2OK with default headers values
func NewPatchMeteringPolicyAssignmentUsingPATCH2OK() *PatchMeteringPolicyAssignmentUsingPATCH2OK {
	return &PatchMeteringPolicyAssignmentUsingPATCH2OK{}
}

/*
PatchMeteringPolicyAssignmentUsingPATCH2OK describes a response with status code 200, with default header values.

OK
*/
type PatchMeteringPolicyAssignmentUsingPATCH2OK struct {
	Payload *models.MeteringPolicyAssignment
}

// IsSuccess returns true when this patch metering policy assignment using p a t c h2 o k response has a 2xx status code
func (o *PatchMeteringPolicyAssignmentUsingPATCH2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch metering policy assignment using p a t c h2 o k response has a 3xx status code
func (o *PatchMeteringPolicyAssignmentUsingPATCH2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch metering policy assignment using p a t c h2 o k response has a 4xx status code
func (o *PatchMeteringPolicyAssignmentUsingPATCH2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch metering policy assignment using p a t c h2 o k response has a 5xx status code
func (o *PatchMeteringPolicyAssignmentUsingPATCH2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch metering policy assignment using p a t c h2 o k response a status code equal to that given
func (o *PatchMeteringPolicyAssignmentUsingPATCH2OK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchMeteringPolicyAssignmentUsingPATCH2OK) Error() string {
	return fmt.Sprintf("[PATCH /price/api/private/pricing-card-assignments/{id}][%d] patchMeteringPolicyAssignmentUsingPATCH2OK  %+v", 200, o.Payload)
}

func (o *PatchMeteringPolicyAssignmentUsingPATCH2OK) String() string {
	return fmt.Sprintf("[PATCH /price/api/private/pricing-card-assignments/{id}][%d] patchMeteringPolicyAssignmentUsingPATCH2OK  %+v", 200, o.Payload)
}

func (o *PatchMeteringPolicyAssignmentUsingPATCH2OK) GetPayload() *models.MeteringPolicyAssignment {
	return o.Payload
}

func (o *PatchMeteringPolicyAssignmentUsingPATCH2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MeteringPolicyAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchMeteringPolicyAssignmentUsingPATCH2Unauthorized creates a PatchMeteringPolicyAssignmentUsingPATCH2Unauthorized with default headers values
func NewPatchMeteringPolicyAssignmentUsingPATCH2Unauthorized() *PatchMeteringPolicyAssignmentUsingPATCH2Unauthorized {
	return &PatchMeteringPolicyAssignmentUsingPATCH2Unauthorized{}
}

/*
PatchMeteringPolicyAssignmentUsingPATCH2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PatchMeteringPolicyAssignmentUsingPATCH2Unauthorized struct {
}

// IsSuccess returns true when this patch metering policy assignment using p a t c h2 unauthorized response has a 2xx status code
func (o *PatchMeteringPolicyAssignmentUsingPATCH2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch metering policy assignment using p a t c h2 unauthorized response has a 3xx status code
func (o *PatchMeteringPolicyAssignmentUsingPATCH2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch metering policy assignment using p a t c h2 unauthorized response has a 4xx status code
func (o *PatchMeteringPolicyAssignmentUsingPATCH2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch metering policy assignment using p a t c h2 unauthorized response has a 5xx status code
func (o *PatchMeteringPolicyAssignmentUsingPATCH2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch metering policy assignment using p a t c h2 unauthorized response a status code equal to that given
func (o *PatchMeteringPolicyAssignmentUsingPATCH2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchMeteringPolicyAssignmentUsingPATCH2Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /price/api/private/pricing-card-assignments/{id}][%d] patchMeteringPolicyAssignmentUsingPATCH2Unauthorized ", 401)
}

func (o *PatchMeteringPolicyAssignmentUsingPATCH2Unauthorized) String() string {
	return fmt.Sprintf("[PATCH /price/api/private/pricing-card-assignments/{id}][%d] patchMeteringPolicyAssignmentUsingPATCH2Unauthorized ", 401)
}

func (o *PatchMeteringPolicyAssignmentUsingPATCH2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
