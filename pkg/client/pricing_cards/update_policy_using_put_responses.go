// Code generated by go-swagger; DO NOT EDIT.

package pricing_cards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdatePolicyUsingPUTReader is a Reader for the UpdatePolicyUsingPUT structure.
type UpdatePolicyUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePolicyUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePolicyUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdatePolicyUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePolicyUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePolicyUsingPUTOK creates a UpdatePolicyUsingPUTOK with default headers values
func NewUpdatePolicyUsingPUTOK() *UpdatePolicyUsingPUTOK {
	return &UpdatePolicyUsingPUTOK{}
}

/*UpdatePolicyUsingPUTOK handles this case with default header values.

OK
*/
type UpdatePolicyUsingPUTOK struct {
	Payload *models.MeteringPolicy
}

func (o *UpdatePolicyUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /price/api/private/pricing-cards/{id}][%d] updatePolicyUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdatePolicyUsingPUTOK) GetPayload() *models.MeteringPolicy {
	return o.Payload
}

func (o *UpdatePolicyUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MeteringPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePolicyUsingPUTUnauthorized creates a UpdatePolicyUsingPUTUnauthorized with default headers values
func NewUpdatePolicyUsingPUTUnauthorized() *UpdatePolicyUsingPUTUnauthorized {
	return &UpdatePolicyUsingPUTUnauthorized{}
}

/*UpdatePolicyUsingPUTUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdatePolicyUsingPUTUnauthorized struct {
}

func (o *UpdatePolicyUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /price/api/private/pricing-cards/{id}][%d] updatePolicyUsingPUTUnauthorized ", 401)
}

func (o *UpdatePolicyUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePolicyUsingPUTNotFound creates a UpdatePolicyUsingPUTNotFound with default headers values
func NewUpdatePolicyUsingPUTNotFound() *UpdatePolicyUsingPUTNotFound {
	return &UpdatePolicyUsingPUTNotFound{}
}

/*UpdatePolicyUsingPUTNotFound handles this case with default header values.

Not Found
*/
type UpdatePolicyUsingPUTNotFound struct {
	Payload *models.Error
}

func (o *UpdatePolicyUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /price/api/private/pricing-cards/{id}][%d] updatePolicyUsingPUTNotFound  %+v", 404, o.Payload)
}

func (o *UpdatePolicyUsingPUTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdatePolicyUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
