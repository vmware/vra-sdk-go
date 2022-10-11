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

// GetAllScenarioConfigsUsingGET2Reader is a Reader for the GetAllScenarioConfigsUsingGET2 structure.
type GetAllScenarioConfigsUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllScenarioConfigsUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllScenarioConfigsUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllScenarioConfigsUsingGET2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllScenarioConfigsUsingGET2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllScenarioConfigsUsingGET2OK creates a GetAllScenarioConfigsUsingGET2OK with default headers values
func NewGetAllScenarioConfigsUsingGET2OK() *GetAllScenarioConfigsUsingGET2OK {
	return &GetAllScenarioConfigsUsingGET2OK{}
}

/*
GetAllScenarioConfigsUsingGET2OK describes a response with status code 200, with default header values.

OK
*/
type GetAllScenarioConfigsUsingGET2OK struct {
	Payload *models.PageOfNotificationScenarioConfig
}

// IsSuccess returns true when this get all scenario configs using g e t2 o k response has a 2xx status code
func (o *GetAllScenarioConfigsUsingGET2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all scenario configs using g e t2 o k response has a 3xx status code
func (o *GetAllScenarioConfigsUsingGET2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all scenario configs using g e t2 o k response has a 4xx status code
func (o *GetAllScenarioConfigsUsingGET2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all scenario configs using g e t2 o k response has a 5xx status code
func (o *GetAllScenarioConfigsUsingGET2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all scenario configs using g e t2 o k response a status code equal to that given
func (o *GetAllScenarioConfigsUsingGET2OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllScenarioConfigsUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /notification/api/scenario-configs][%d] getAllScenarioConfigsUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetAllScenarioConfigsUsingGET2OK) String() string {
	return fmt.Sprintf("[GET /notification/api/scenario-configs][%d] getAllScenarioConfigsUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetAllScenarioConfigsUsingGET2OK) GetPayload() *models.PageOfNotificationScenarioConfig {
	return o.Payload
}

func (o *GetAllScenarioConfigsUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfNotificationScenarioConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllScenarioConfigsUsingGET2Unauthorized creates a GetAllScenarioConfigsUsingGET2Unauthorized with default headers values
func NewGetAllScenarioConfigsUsingGET2Unauthorized() *GetAllScenarioConfigsUsingGET2Unauthorized {
	return &GetAllScenarioConfigsUsingGET2Unauthorized{}
}

/*
GetAllScenarioConfigsUsingGET2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAllScenarioConfigsUsingGET2Unauthorized struct {
}

// IsSuccess returns true when this get all scenario configs using g e t2 unauthorized response has a 2xx status code
func (o *GetAllScenarioConfigsUsingGET2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all scenario configs using g e t2 unauthorized response has a 3xx status code
func (o *GetAllScenarioConfigsUsingGET2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all scenario configs using g e t2 unauthorized response has a 4xx status code
func (o *GetAllScenarioConfigsUsingGET2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all scenario configs using g e t2 unauthorized response has a 5xx status code
func (o *GetAllScenarioConfigsUsingGET2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all scenario configs using g e t2 unauthorized response a status code equal to that given
func (o *GetAllScenarioConfigsUsingGET2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAllScenarioConfigsUsingGET2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /notification/api/scenario-configs][%d] getAllScenarioConfigsUsingGET2Unauthorized ", 401)
}

func (o *GetAllScenarioConfigsUsingGET2Unauthorized) String() string {
	return fmt.Sprintf("[GET /notification/api/scenario-configs][%d] getAllScenarioConfigsUsingGET2Unauthorized ", 401)
}

func (o *GetAllScenarioConfigsUsingGET2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllScenarioConfigsUsingGET2Forbidden creates a GetAllScenarioConfigsUsingGET2Forbidden with default headers values
func NewGetAllScenarioConfigsUsingGET2Forbidden() *GetAllScenarioConfigsUsingGET2Forbidden {
	return &GetAllScenarioConfigsUsingGET2Forbidden{}
}

/*
GetAllScenarioConfigsUsingGET2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllScenarioConfigsUsingGET2Forbidden struct {
}

// IsSuccess returns true when this get all scenario configs using g e t2 forbidden response has a 2xx status code
func (o *GetAllScenarioConfigsUsingGET2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all scenario configs using g e t2 forbidden response has a 3xx status code
func (o *GetAllScenarioConfigsUsingGET2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all scenario configs using g e t2 forbidden response has a 4xx status code
func (o *GetAllScenarioConfigsUsingGET2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all scenario configs using g e t2 forbidden response has a 5xx status code
func (o *GetAllScenarioConfigsUsingGET2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get all scenario configs using g e t2 forbidden response a status code equal to that given
func (o *GetAllScenarioConfigsUsingGET2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAllScenarioConfigsUsingGET2Forbidden) Error() string {
	return fmt.Sprintf("[GET /notification/api/scenario-configs][%d] getAllScenarioConfigsUsingGET2Forbidden ", 403)
}

func (o *GetAllScenarioConfigsUsingGET2Forbidden) String() string {
	return fmt.Sprintf("[GET /notification/api/scenario-configs][%d] getAllScenarioConfigsUsingGET2Forbidden ", 403)
}

func (o *GetAllScenarioConfigsUsingGET2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
