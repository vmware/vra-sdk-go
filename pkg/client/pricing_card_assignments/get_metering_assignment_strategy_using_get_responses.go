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

// GetMeteringAssignmentStrategyUsingGETReader is a Reader for the GetMeteringAssignmentStrategyUsingGET structure.
type GetMeteringAssignmentStrategyUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeteringAssignmentStrategyUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMeteringAssignmentStrategyUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetMeteringAssignmentStrategyUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMeteringAssignmentStrategyUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMeteringAssignmentStrategyUsingGETOK creates a GetMeteringAssignmentStrategyUsingGETOK with default headers values
func NewGetMeteringAssignmentStrategyUsingGETOK() *GetMeteringAssignmentStrategyUsingGETOK {
	return &GetMeteringAssignmentStrategyUsingGETOK{}
}

/*GetMeteringAssignmentStrategyUsingGETOK handles this case with default header values.

OK
*/
type GetMeteringAssignmentStrategyUsingGETOK struct {
	Payload *models.MeteringAssignmentStrategy
}

func (o *GetMeteringAssignmentStrategyUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /price/api/private/pricing-card-assignments/strategy][%d] getMeteringAssignmentStrategyUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetMeteringAssignmentStrategyUsingGETOK) GetPayload() *models.MeteringAssignmentStrategy {
	return o.Payload
}

func (o *GetMeteringAssignmentStrategyUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MeteringAssignmentStrategy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeteringAssignmentStrategyUsingGETUnauthorized creates a GetMeteringAssignmentStrategyUsingGETUnauthorized with default headers values
func NewGetMeteringAssignmentStrategyUsingGETUnauthorized() *GetMeteringAssignmentStrategyUsingGETUnauthorized {
	return &GetMeteringAssignmentStrategyUsingGETUnauthorized{}
}

/*GetMeteringAssignmentStrategyUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetMeteringAssignmentStrategyUsingGETUnauthorized struct {
}

func (o *GetMeteringAssignmentStrategyUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /price/api/private/pricing-card-assignments/strategy][%d] getMeteringAssignmentStrategyUsingGETUnauthorized ", 401)
}

func (o *GetMeteringAssignmentStrategyUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMeteringAssignmentStrategyUsingGETNotFound creates a GetMeteringAssignmentStrategyUsingGETNotFound with default headers values
func NewGetMeteringAssignmentStrategyUsingGETNotFound() *GetMeteringAssignmentStrategyUsingGETNotFound {
	return &GetMeteringAssignmentStrategyUsingGETNotFound{}
}

/*GetMeteringAssignmentStrategyUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetMeteringAssignmentStrategyUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetMeteringAssignmentStrategyUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /price/api/private/pricing-card-assignments/strategy][%d] getMeteringAssignmentStrategyUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetMeteringAssignmentStrategyUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMeteringAssignmentStrategyUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
