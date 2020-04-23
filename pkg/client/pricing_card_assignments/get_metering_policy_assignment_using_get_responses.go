// Code generated by go-swagger; DO NOT EDIT.

package pricing_card_assignments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetMeteringPolicyAssignmentUsingGETReader is a Reader for the GetMeteringPolicyAssignmentUsingGET structure.
type GetMeteringPolicyAssignmentUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeteringPolicyAssignmentUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMeteringPolicyAssignmentUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetMeteringPolicyAssignmentUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMeteringPolicyAssignmentUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMeteringPolicyAssignmentUsingGETOK creates a GetMeteringPolicyAssignmentUsingGETOK with default headers values
func NewGetMeteringPolicyAssignmentUsingGETOK() *GetMeteringPolicyAssignmentUsingGETOK {
	return &GetMeteringPolicyAssignmentUsingGETOK{}
}

/*GetMeteringPolicyAssignmentUsingGETOK handles this case with default header values.

OK
*/
type GetMeteringPolicyAssignmentUsingGETOK struct {
	Payload *models.MeteringPolicyAssignment
}

func (o *GetMeteringPolicyAssignmentUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /price/api/private/pricing-card-assignments/{id}][%d] getMeteringPolicyAssignmentUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetMeteringPolicyAssignmentUsingGETOK) GetPayload() *models.MeteringPolicyAssignment {
	return o.Payload
}

func (o *GetMeteringPolicyAssignmentUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MeteringPolicyAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeteringPolicyAssignmentUsingGETUnauthorized creates a GetMeteringPolicyAssignmentUsingGETUnauthorized with default headers values
func NewGetMeteringPolicyAssignmentUsingGETUnauthorized() *GetMeteringPolicyAssignmentUsingGETUnauthorized {
	return &GetMeteringPolicyAssignmentUsingGETUnauthorized{}
}

/*GetMeteringPolicyAssignmentUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetMeteringPolicyAssignmentUsingGETUnauthorized struct {
}

func (o *GetMeteringPolicyAssignmentUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /price/api/private/pricing-card-assignments/{id}][%d] getMeteringPolicyAssignmentUsingGETUnauthorized ", 401)
}

func (o *GetMeteringPolicyAssignmentUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMeteringPolicyAssignmentUsingGETNotFound creates a GetMeteringPolicyAssignmentUsingGETNotFound with default headers values
func NewGetMeteringPolicyAssignmentUsingGETNotFound() *GetMeteringPolicyAssignmentUsingGETNotFound {
	return &GetMeteringPolicyAssignmentUsingGETNotFound{}
}

/*GetMeteringPolicyAssignmentUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetMeteringPolicyAssignmentUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetMeteringPolicyAssignmentUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /price/api/private/pricing-card-assignments/{id}][%d] getMeteringPolicyAssignmentUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetMeteringPolicyAssignmentUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMeteringPolicyAssignmentUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
