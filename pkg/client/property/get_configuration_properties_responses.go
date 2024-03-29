// Code generated by go-swagger; DO NOT EDIT.

package property

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetConfigurationPropertiesReader is a Reader for the GetConfigurationProperties structure.
type GetConfigurationPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigurationPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetConfigurationPropertiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConfigurationPropertiesOK creates a GetConfigurationPropertiesOK with default headers values
func NewGetConfigurationPropertiesOK() *GetConfigurationPropertiesOK {
	return &GetConfigurationPropertiesOK{}
}

/*
GetConfigurationPropertiesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConfigurationPropertiesOK struct {
	Payload *models.ConfigurationPropertyResult
}

// IsSuccess returns true when this get configuration properties o k response has a 2xx status code
func (o *GetConfigurationPropertiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get configuration properties o k response has a 3xx status code
func (o *GetConfigurationPropertiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration properties o k response has a 4xx status code
func (o *GetConfigurationPropertiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get configuration properties o k response has a 5xx status code
func (o *GetConfigurationPropertiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration properties o k response a status code equal to that given
func (o *GetConfigurationPropertiesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConfigurationPropertiesOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/configuration-properties][%d] getConfigurationPropertiesOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationPropertiesOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/configuration-properties][%d] getConfigurationPropertiesOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationPropertiesOK) GetPayload() *models.ConfigurationPropertyResult {
	return o.Payload
}

func (o *GetConfigurationPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigurationPropertyResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationPropertiesForbidden creates a GetConfigurationPropertiesForbidden with default headers values
func NewGetConfigurationPropertiesForbidden() *GetConfigurationPropertiesForbidden {
	return &GetConfigurationPropertiesForbidden{}
}

/*
GetConfigurationPropertiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetConfigurationPropertiesForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get configuration properties forbidden response has a 2xx status code
func (o *GetConfigurationPropertiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration properties forbidden response has a 3xx status code
func (o *GetConfigurationPropertiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration properties forbidden response has a 4xx status code
func (o *GetConfigurationPropertiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration properties forbidden response has a 5xx status code
func (o *GetConfigurationPropertiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration properties forbidden response a status code equal to that given
func (o *GetConfigurationPropertiesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConfigurationPropertiesForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/configuration-properties][%d] getConfigurationPropertiesForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigurationPropertiesForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/configuration-properties][%d] getConfigurationPropertiesForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigurationPropertiesForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetConfigurationPropertiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
