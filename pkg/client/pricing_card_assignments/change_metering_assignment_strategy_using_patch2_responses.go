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

// ChangeMeteringAssignmentStrategyUsingPATCH2Reader is a Reader for the ChangeMeteringAssignmentStrategyUsingPATCH2 structure.
type ChangeMeteringAssignmentStrategyUsingPATCH2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeMeteringAssignmentStrategyUsingPATCH2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeMeteringAssignmentStrategyUsingPATCH2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewChangeMeteringAssignmentStrategyUsingPATCH2OK creates a ChangeMeteringAssignmentStrategyUsingPATCH2OK with default headers values
func NewChangeMeteringAssignmentStrategyUsingPATCH2OK() *ChangeMeteringAssignmentStrategyUsingPATCH2OK {
	return &ChangeMeteringAssignmentStrategyUsingPATCH2OK{}
}

/*
ChangeMeteringAssignmentStrategyUsingPATCH2OK describes a response with status code 200, with default header values.

OK
*/
type ChangeMeteringAssignmentStrategyUsingPATCH2OK struct {
	Payload *models.MeteringAssignmentStrategy
}

// IsSuccess returns true when this change metering assignment strategy using p a t c h2 o k response has a 2xx status code
func (o *ChangeMeteringAssignmentStrategyUsingPATCH2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this change metering assignment strategy using p a t c h2 o k response has a 3xx status code
func (o *ChangeMeteringAssignmentStrategyUsingPATCH2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change metering assignment strategy using p a t c h2 o k response has a 4xx status code
func (o *ChangeMeteringAssignmentStrategyUsingPATCH2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this change metering assignment strategy using p a t c h2 o k response has a 5xx status code
func (o *ChangeMeteringAssignmentStrategyUsingPATCH2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this change metering assignment strategy using p a t c h2 o k response a status code equal to that given
func (o *ChangeMeteringAssignmentStrategyUsingPATCH2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ChangeMeteringAssignmentStrategyUsingPATCH2OK) Error() string {
	return fmt.Sprintf("[PATCH /price/api/private/pricing-card-assignments/strategy][%d] changeMeteringAssignmentStrategyUsingPATCH2OK  %+v", 200, o.Payload)
}

func (o *ChangeMeteringAssignmentStrategyUsingPATCH2OK) String() string {
	return fmt.Sprintf("[PATCH /price/api/private/pricing-card-assignments/strategy][%d] changeMeteringAssignmentStrategyUsingPATCH2OK  %+v", 200, o.Payload)
}

func (o *ChangeMeteringAssignmentStrategyUsingPATCH2OK) GetPayload() *models.MeteringAssignmentStrategy {
	return o.Payload
}

func (o *ChangeMeteringAssignmentStrategyUsingPATCH2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MeteringAssignmentStrategy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized creates a ChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized with default headers values
func NewChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized() *ChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized {
	return &ChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized{}
}

/*
ChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized struct {
}

// IsSuccess returns true when this change metering assignment strategy using p a t c h2 unauthorized response has a 2xx status code
func (o *ChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this change metering assignment strategy using p a t c h2 unauthorized response has a 3xx status code
func (o *ChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change metering assignment strategy using p a t c h2 unauthorized response has a 4xx status code
func (o *ChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this change metering assignment strategy using p a t c h2 unauthorized response has a 5xx status code
func (o *ChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this change metering assignment strategy using p a t c h2 unauthorized response a status code equal to that given
func (o *ChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /price/api/private/pricing-card-assignments/strategy][%d] changeMeteringAssignmentStrategyUsingPATCH2Unauthorized ", 401)
}

func (o *ChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized) String() string {
	return fmt.Sprintf("[PATCH /price/api/private/pricing-card-assignments/strategy][%d] changeMeteringAssignmentStrategyUsingPATCH2Unauthorized ", 401)
}

func (o *ChangeMeteringAssignmentStrategyUsingPATCH2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}