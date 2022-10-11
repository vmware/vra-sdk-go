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

// DeleteConfigurationPropertyReader is a Reader for the DeleteConfigurationProperty structure.
type DeleteConfigurationPropertyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConfigurationPropertyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteConfigurationPropertyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteConfigurationPropertyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteConfigurationPropertyOK creates a DeleteConfigurationPropertyOK with default headers values
func NewDeleteConfigurationPropertyOK() *DeleteConfigurationPropertyOK {
	return &DeleteConfigurationPropertyOK{}
}

/*
DeleteConfigurationPropertyOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteConfigurationPropertyOK struct {
	Payload *models.ConfigurationProperty
}

// IsSuccess returns true when this delete configuration property o k response has a 2xx status code
func (o *DeleteConfigurationPropertyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete configuration property o k response has a 3xx status code
func (o *DeleteConfigurationPropertyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete configuration property o k response has a 4xx status code
func (o *DeleteConfigurationPropertyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete configuration property o k response has a 5xx status code
func (o *DeleteConfigurationPropertyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete configuration property o k response a status code equal to that given
func (o *DeleteConfigurationPropertyOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteConfigurationPropertyOK) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/configuration-properties/{id}][%d] deleteConfigurationPropertyOK  %+v", 200, o.Payload)
}

func (o *DeleteConfigurationPropertyOK) String() string {
	return fmt.Sprintf("[DELETE /iaas/api/configuration-properties/{id}][%d] deleteConfigurationPropertyOK  %+v", 200, o.Payload)
}

func (o *DeleteConfigurationPropertyOK) GetPayload() *models.ConfigurationProperty {
	return o.Payload
}

func (o *DeleteConfigurationPropertyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigurationProperty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConfigurationPropertyForbidden creates a DeleteConfigurationPropertyForbidden with default headers values
func NewDeleteConfigurationPropertyForbidden() *DeleteConfigurationPropertyForbidden {
	return &DeleteConfigurationPropertyForbidden{}
}

/*
DeleteConfigurationPropertyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteConfigurationPropertyForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this delete configuration property forbidden response has a 2xx status code
func (o *DeleteConfigurationPropertyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete configuration property forbidden response has a 3xx status code
func (o *DeleteConfigurationPropertyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete configuration property forbidden response has a 4xx status code
func (o *DeleteConfigurationPropertyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete configuration property forbidden response has a 5xx status code
func (o *DeleteConfigurationPropertyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete configuration property forbidden response a status code equal to that given
func (o *DeleteConfigurationPropertyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteConfigurationPropertyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/configuration-properties/{id}][%d] deleteConfigurationPropertyForbidden  %+v", 403, o.Payload)
}

func (o *DeleteConfigurationPropertyForbidden) String() string {
	return fmt.Sprintf("[DELETE /iaas/api/configuration-properties/{id}][%d] deleteConfigurationPropertyForbidden  %+v", 403, o.Payload)
}

func (o *DeleteConfigurationPropertyForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *DeleteConfigurationPropertyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
