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

// GetConfigurationPropertyReader is a Reader for the GetConfigurationProperty structure.
type GetConfigurationPropertyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationPropertyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigurationPropertyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetConfigurationPropertyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConfigurationPropertyOK creates a GetConfigurationPropertyOK with default headers values
func NewGetConfigurationPropertyOK() *GetConfigurationPropertyOK {
	return &GetConfigurationPropertyOK{}
}

/*
GetConfigurationPropertyOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConfigurationPropertyOK struct {
	Payload *models.ConfigurationProperty
}

// IsSuccess returns true when this get configuration property o k response has a 2xx status code
func (o *GetConfigurationPropertyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get configuration property o k response has a 3xx status code
func (o *GetConfigurationPropertyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration property o k response has a 4xx status code
func (o *GetConfigurationPropertyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get configuration property o k response has a 5xx status code
func (o *GetConfigurationPropertyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration property o k response a status code equal to that given
func (o *GetConfigurationPropertyOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConfigurationPropertyOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/configuration-properties/{id}][%d] getConfigurationPropertyOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationPropertyOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/configuration-properties/{id}][%d] getConfigurationPropertyOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationPropertyOK) GetPayload() *models.ConfigurationProperty {
	return o.Payload
}

func (o *GetConfigurationPropertyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigurationProperty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationPropertyForbidden creates a GetConfigurationPropertyForbidden with default headers values
func NewGetConfigurationPropertyForbidden() *GetConfigurationPropertyForbidden {
	return &GetConfigurationPropertyForbidden{}
}

/*
GetConfigurationPropertyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetConfigurationPropertyForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get configuration property forbidden response has a 2xx status code
func (o *GetConfigurationPropertyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration property forbidden response has a 3xx status code
func (o *GetConfigurationPropertyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration property forbidden response has a 4xx status code
func (o *GetConfigurationPropertyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration property forbidden response has a 5xx status code
func (o *GetConfigurationPropertyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration property forbidden response a status code equal to that given
func (o *GetConfigurationPropertyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConfigurationPropertyForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/configuration-properties/{id}][%d] getConfigurationPropertyForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigurationPropertyForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/configuration-properties/{id}][%d] getConfigurationPropertyForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigurationPropertyForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetConfigurationPropertyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
