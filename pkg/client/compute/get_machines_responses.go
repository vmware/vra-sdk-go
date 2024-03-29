// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetMachinesReader is a Reader for the GetMachines structure.
type GetMachinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMachinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetMachinesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMachinesOK creates a GetMachinesOK with default headers values
func NewGetMachinesOK() *GetMachinesOK {
	return &GetMachinesOK{}
}

/*
GetMachinesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetMachinesOK struct {
	Payload *models.MachineResult
}

// IsSuccess returns true when this get machines o k response has a 2xx status code
func (o *GetMachinesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get machines o k response has a 3xx status code
func (o *GetMachinesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get machines o k response has a 4xx status code
func (o *GetMachinesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get machines o k response has a 5xx status code
func (o *GetMachinesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get machines o k response a status code equal to that given
func (o *GetMachinesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetMachinesOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines][%d] getMachinesOK  %+v", 200, o.Payload)
}

func (o *GetMachinesOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/machines][%d] getMachinesOK  %+v", 200, o.Payload)
}

func (o *GetMachinesOK) GetPayload() *models.MachineResult {
	return o.Payload
}

func (o *GetMachinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachinesForbidden creates a GetMachinesForbidden with default headers values
func NewGetMachinesForbidden() *GetMachinesForbidden {
	return &GetMachinesForbidden{}
}

/*
GetMachinesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetMachinesForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get machines forbidden response has a 2xx status code
func (o *GetMachinesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get machines forbidden response has a 3xx status code
func (o *GetMachinesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get machines forbidden response has a 4xx status code
func (o *GetMachinesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get machines forbidden response has a 5xx status code
func (o *GetMachinesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get machines forbidden response a status code equal to that given
func (o *GetMachinesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetMachinesForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines][%d] getMachinesForbidden  %+v", 403, o.Payload)
}

func (o *GetMachinesForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/machines][%d] getMachinesForbidden  %+v", 403, o.Payload)
}

func (o *GetMachinesForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetMachinesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
