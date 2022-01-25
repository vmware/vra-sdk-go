// Code generated by go-swagger; DO NOT EDIT.

package notification_scenario_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetScenarioConfigUsingGET2Reader is a Reader for the GetScenarioConfigUsingGET2 structure.
type GetScenarioConfigUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScenarioConfigUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScenarioConfigUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetScenarioConfigUsingGET2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScenarioConfigUsingGET2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScenarioConfigUsingGET2OK creates a GetScenarioConfigUsingGET2OK with default headers values
func NewGetScenarioConfigUsingGET2OK() *GetScenarioConfigUsingGET2OK {
	return &GetScenarioConfigUsingGET2OK{}
}

/* GetScenarioConfigUsingGET2OK describes a response with status code 200, with default header values.

OK
*/
type GetScenarioConfigUsingGET2OK struct {
	Payload *models.NotificationScenarioConfig
}

func (o *GetScenarioConfigUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /notification/api/scenario-configs/{scenarioId}][%d] getScenarioConfigUsingGET2OK  %+v", 200, o.Payload)
}
func (o *GetScenarioConfigUsingGET2OK) GetPayload() *models.NotificationScenarioConfig {
	return o.Payload
}

func (o *GetScenarioConfigUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotificationScenarioConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScenarioConfigUsingGET2Forbidden creates a GetScenarioConfigUsingGET2Forbidden with default headers values
func NewGetScenarioConfigUsingGET2Forbidden() *GetScenarioConfigUsingGET2Forbidden {
	return &GetScenarioConfigUsingGET2Forbidden{}
}

/* GetScenarioConfigUsingGET2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetScenarioConfigUsingGET2Forbidden struct {
}

func (o *GetScenarioConfigUsingGET2Forbidden) Error() string {
	return fmt.Sprintf("[GET /notification/api/scenario-configs/{scenarioId}][%d] getScenarioConfigUsingGET2Forbidden ", 403)
}

func (o *GetScenarioConfigUsingGET2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScenarioConfigUsingGET2NotFound creates a GetScenarioConfigUsingGET2NotFound with default headers values
func NewGetScenarioConfigUsingGET2NotFound() *GetScenarioConfigUsingGET2NotFound {
	return &GetScenarioConfigUsingGET2NotFound{}
}

/* GetScenarioConfigUsingGET2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetScenarioConfigUsingGET2NotFound struct {
	Payload *models.Error
}

func (o *GetScenarioConfigUsingGET2NotFound) Error() string {
	return fmt.Sprintf("[GET /notification/api/scenario-configs/{scenarioId}][%d] getScenarioConfigUsingGET2NotFound  %+v", 404, o.Payload)
}
func (o *GetScenarioConfigUsingGET2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetScenarioConfigUsingGET2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
