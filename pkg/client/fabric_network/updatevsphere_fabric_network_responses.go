// Code generated by go-swagger; DO NOT EDIT.

package fabric_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdatevSphereFabricNetworkReader is a Reader for the UpdatevSphereFabricNetwork structure.
type UpdatevSphereFabricNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatevSphereFabricNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatevSphereFabricNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdatevSphereFabricNetworkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatevSphereFabricNetworkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdatevSphereFabricNetworkOK creates a UpdatevSphereFabricNetworkOK with default headers values
func NewUpdatevSphereFabricNetworkOK() *UpdatevSphereFabricNetworkOK {
	return &UpdatevSphereFabricNetworkOK{}
}

/*
UpdatevSphereFabricNetworkOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdatevSphereFabricNetworkOK struct {
	Payload *models.FabricNetworkVsphere
}

// IsSuccess returns true when this updatev sphere fabric network o k response has a 2xx status code
func (o *UpdatevSphereFabricNetworkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this updatev sphere fabric network o k response has a 3xx status code
func (o *UpdatevSphereFabricNetworkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this updatev sphere fabric network o k response has a 4xx status code
func (o *UpdatevSphereFabricNetworkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this updatev sphere fabric network o k response has a 5xx status code
func (o *UpdatevSphereFabricNetworkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this updatev sphere fabric network o k response a status code equal to that given
func (o *UpdatevSphereFabricNetworkOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdatevSphereFabricNetworkOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/fabric-networks-vsphere/{id}][%d] updatevSphereFabricNetworkOK  %+v", 200, o.Payload)
}

func (o *UpdatevSphereFabricNetworkOK) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/fabric-networks-vsphere/{id}][%d] updatevSphereFabricNetworkOK  %+v", 200, o.Payload)
}

func (o *UpdatevSphereFabricNetworkOK) GetPayload() *models.FabricNetworkVsphere {
	return o.Payload
}

func (o *UpdatevSphereFabricNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricNetworkVsphere)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatevSphereFabricNetworkForbidden creates a UpdatevSphereFabricNetworkForbidden with default headers values
func NewUpdatevSphereFabricNetworkForbidden() *UpdatevSphereFabricNetworkForbidden {
	return &UpdatevSphereFabricNetworkForbidden{}
}

/*
UpdatevSphereFabricNetworkForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdatevSphereFabricNetworkForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this updatev sphere fabric network forbidden response has a 2xx status code
func (o *UpdatevSphereFabricNetworkForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this updatev sphere fabric network forbidden response has a 3xx status code
func (o *UpdatevSphereFabricNetworkForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this updatev sphere fabric network forbidden response has a 4xx status code
func (o *UpdatevSphereFabricNetworkForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this updatev sphere fabric network forbidden response has a 5xx status code
func (o *UpdatevSphereFabricNetworkForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this updatev sphere fabric network forbidden response a status code equal to that given
func (o *UpdatevSphereFabricNetworkForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdatevSphereFabricNetworkForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/fabric-networks-vsphere/{id}][%d] updatevSphereFabricNetworkForbidden  %+v", 403, o.Payload)
}

func (o *UpdatevSphereFabricNetworkForbidden) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/fabric-networks-vsphere/{id}][%d] updatevSphereFabricNetworkForbidden  %+v", 403, o.Payload)
}

func (o *UpdatevSphereFabricNetworkForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *UpdatevSphereFabricNetworkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatevSphereFabricNetworkNotFound creates a UpdatevSphereFabricNetworkNotFound with default headers values
func NewUpdatevSphereFabricNetworkNotFound() *UpdatevSphereFabricNetworkNotFound {
	return &UpdatevSphereFabricNetworkNotFound{}
}

/*
UpdatevSphereFabricNetworkNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdatevSphereFabricNetworkNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this updatev sphere fabric network not found response has a 2xx status code
func (o *UpdatevSphereFabricNetworkNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this updatev sphere fabric network not found response has a 3xx status code
func (o *UpdatevSphereFabricNetworkNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this updatev sphere fabric network not found response has a 4xx status code
func (o *UpdatevSphereFabricNetworkNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this updatev sphere fabric network not found response has a 5xx status code
func (o *UpdatevSphereFabricNetworkNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this updatev sphere fabric network not found response a status code equal to that given
func (o *UpdatevSphereFabricNetworkNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdatevSphereFabricNetworkNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/fabric-networks-vsphere/{id}][%d] updatevSphereFabricNetworkNotFound  %+v", 404, o.Payload)
}

func (o *UpdatevSphereFabricNetworkNotFound) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/fabric-networks-vsphere/{id}][%d] updatevSphereFabricNetworkNotFound  %+v", 404, o.Payload)
}

func (o *UpdatevSphereFabricNetworkNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdatevSphereFabricNetworkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
