// Code generated by go-swagger; DO NOT EDIT.

package fabric_vsphere_datastore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetFabricVSphereDatastoreReader is a Reader for the GetFabricVSphereDatastore structure.
type GetFabricVSphereDatastoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFabricVSphereDatastoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFabricVSphereDatastoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFabricVSphereDatastoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFabricVSphereDatastoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFabricVSphereDatastoreOK creates a GetFabricVSphereDatastoreOK with default headers values
func NewGetFabricVSphereDatastoreOK() *GetFabricVSphereDatastoreOK {
	return &GetFabricVSphereDatastoreOK{}
}

/*
GetFabricVSphereDatastoreOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFabricVSphereDatastoreOK struct {
	Payload *models.FabricVsphereDatastore
}

// IsSuccess returns true when this get fabric v sphere datastore o k response has a 2xx status code
func (o *GetFabricVSphereDatastoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get fabric v sphere datastore o k response has a 3xx status code
func (o *GetFabricVSphereDatastoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fabric v sphere datastore o k response has a 4xx status code
func (o *GetFabricVSphereDatastoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fabric v sphere datastore o k response has a 5xx status code
func (o *GetFabricVSphereDatastoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get fabric v sphere datastore o k response a status code equal to that given
func (o *GetFabricVSphereDatastoreOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFabricVSphereDatastoreOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-vsphere-datastores/{id}][%d] getFabricVSphereDatastoreOK  %+v", 200, o.Payload)
}

func (o *GetFabricVSphereDatastoreOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-vsphere-datastores/{id}][%d] getFabricVSphereDatastoreOK  %+v", 200, o.Payload)
}

func (o *GetFabricVSphereDatastoreOK) GetPayload() *models.FabricVsphereDatastore {
	return o.Payload
}

func (o *GetFabricVSphereDatastoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricVsphereDatastore)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricVSphereDatastoreForbidden creates a GetFabricVSphereDatastoreForbidden with default headers values
func NewGetFabricVSphereDatastoreForbidden() *GetFabricVSphereDatastoreForbidden {
	return &GetFabricVSphereDatastoreForbidden{}
}

/*
GetFabricVSphereDatastoreForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFabricVSphereDatastoreForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get fabric v sphere datastore forbidden response has a 2xx status code
func (o *GetFabricVSphereDatastoreForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fabric v sphere datastore forbidden response has a 3xx status code
func (o *GetFabricVSphereDatastoreForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fabric v sphere datastore forbidden response has a 4xx status code
func (o *GetFabricVSphereDatastoreForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fabric v sphere datastore forbidden response has a 5xx status code
func (o *GetFabricVSphereDatastoreForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get fabric v sphere datastore forbidden response a status code equal to that given
func (o *GetFabricVSphereDatastoreForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFabricVSphereDatastoreForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-vsphere-datastores/{id}][%d] getFabricVSphereDatastoreForbidden  %+v", 403, o.Payload)
}

func (o *GetFabricVSphereDatastoreForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-vsphere-datastores/{id}][%d] getFabricVSphereDatastoreForbidden  %+v", 403, o.Payload)
}

func (o *GetFabricVSphereDatastoreForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetFabricVSphereDatastoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricVSphereDatastoreNotFound creates a GetFabricVSphereDatastoreNotFound with default headers values
func NewGetFabricVSphereDatastoreNotFound() *GetFabricVSphereDatastoreNotFound {
	return &GetFabricVSphereDatastoreNotFound{}
}

/*
GetFabricVSphereDatastoreNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetFabricVSphereDatastoreNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get fabric v sphere datastore not found response has a 2xx status code
func (o *GetFabricVSphereDatastoreNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fabric v sphere datastore not found response has a 3xx status code
func (o *GetFabricVSphereDatastoreNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fabric v sphere datastore not found response has a 4xx status code
func (o *GetFabricVSphereDatastoreNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fabric v sphere datastore not found response has a 5xx status code
func (o *GetFabricVSphereDatastoreNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get fabric v sphere datastore not found response a status code equal to that given
func (o *GetFabricVSphereDatastoreNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFabricVSphereDatastoreNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-vsphere-datastores/{id}][%d] getFabricVSphereDatastoreNotFound  %+v", 404, o.Payload)
}

func (o *GetFabricVSphereDatastoreNotFound) String() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-vsphere-datastores/{id}][%d] getFabricVSphereDatastoreNotFound  %+v", 404, o.Payload)
}

func (o *GetFabricVSphereDatastoreNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFabricVSphereDatastoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
