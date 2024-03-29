// Code generated by go-swagger; DO NOT EDIT.

package load_balancer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetLoadBalancerNetworkInterfaceReader is a Reader for the GetLoadBalancerNetworkInterface structure.
type GetLoadBalancerNetworkInterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoadBalancerNetworkInterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLoadBalancerNetworkInterfaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetLoadBalancerNetworkInterfaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLoadBalancerNetworkInterfaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLoadBalancerNetworkInterfaceOK creates a GetLoadBalancerNetworkInterfaceOK with default headers values
func NewGetLoadBalancerNetworkInterfaceOK() *GetLoadBalancerNetworkInterfaceOK {
	return &GetLoadBalancerNetworkInterfaceOK{}
}

/*
GetLoadBalancerNetworkInterfaceOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLoadBalancerNetworkInterfaceOK struct {
	Payload *models.NetworkInterface
}

// IsSuccess returns true when this get load balancer network interface o k response has a 2xx status code
func (o *GetLoadBalancerNetworkInterfaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get load balancer network interface o k response has a 3xx status code
func (o *GetLoadBalancerNetworkInterfaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get load balancer network interface o k response has a 4xx status code
func (o *GetLoadBalancerNetworkInterfaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get load balancer network interface o k response has a 5xx status code
func (o *GetLoadBalancerNetworkInterfaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get load balancer network interface o k response a status code equal to that given
func (o *GetLoadBalancerNetworkInterfaceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLoadBalancerNetworkInterfaceOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/load-balancers/{id}/network-interfaces/{id1}][%d] getLoadBalancerNetworkInterfaceOK  %+v", 200, o.Payload)
}

func (o *GetLoadBalancerNetworkInterfaceOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/load-balancers/{id}/network-interfaces/{id1}][%d] getLoadBalancerNetworkInterfaceOK  %+v", 200, o.Payload)
}

func (o *GetLoadBalancerNetworkInterfaceOK) GetPayload() *models.NetworkInterface {
	return o.Payload
}

func (o *GetLoadBalancerNetworkInterfaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoadBalancerNetworkInterfaceForbidden creates a GetLoadBalancerNetworkInterfaceForbidden with default headers values
func NewGetLoadBalancerNetworkInterfaceForbidden() *GetLoadBalancerNetworkInterfaceForbidden {
	return &GetLoadBalancerNetworkInterfaceForbidden{}
}

/*
GetLoadBalancerNetworkInterfaceForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetLoadBalancerNetworkInterfaceForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get load balancer network interface forbidden response has a 2xx status code
func (o *GetLoadBalancerNetworkInterfaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get load balancer network interface forbidden response has a 3xx status code
func (o *GetLoadBalancerNetworkInterfaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get load balancer network interface forbidden response has a 4xx status code
func (o *GetLoadBalancerNetworkInterfaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get load balancer network interface forbidden response has a 5xx status code
func (o *GetLoadBalancerNetworkInterfaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get load balancer network interface forbidden response a status code equal to that given
func (o *GetLoadBalancerNetworkInterfaceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLoadBalancerNetworkInterfaceForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/load-balancers/{id}/network-interfaces/{id1}][%d] getLoadBalancerNetworkInterfaceForbidden  %+v", 403, o.Payload)
}

func (o *GetLoadBalancerNetworkInterfaceForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/load-balancers/{id}/network-interfaces/{id1}][%d] getLoadBalancerNetworkInterfaceForbidden  %+v", 403, o.Payload)
}

func (o *GetLoadBalancerNetworkInterfaceForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetLoadBalancerNetworkInterfaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoadBalancerNetworkInterfaceNotFound creates a GetLoadBalancerNetworkInterfaceNotFound with default headers values
func NewGetLoadBalancerNetworkInterfaceNotFound() *GetLoadBalancerNetworkInterfaceNotFound {
	return &GetLoadBalancerNetworkInterfaceNotFound{}
}

/*
GetLoadBalancerNetworkInterfaceNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetLoadBalancerNetworkInterfaceNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get load balancer network interface not found response has a 2xx status code
func (o *GetLoadBalancerNetworkInterfaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get load balancer network interface not found response has a 3xx status code
func (o *GetLoadBalancerNetworkInterfaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get load balancer network interface not found response has a 4xx status code
func (o *GetLoadBalancerNetworkInterfaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get load balancer network interface not found response has a 5xx status code
func (o *GetLoadBalancerNetworkInterfaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get load balancer network interface not found response a status code equal to that given
func (o *GetLoadBalancerNetworkInterfaceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLoadBalancerNetworkInterfaceNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/load-balancers/{id}/network-interfaces/{id1}][%d] getLoadBalancerNetworkInterfaceNotFound  %+v", 404, o.Payload)
}

func (o *GetLoadBalancerNetworkInterfaceNotFound) String() string {
	return fmt.Sprintf("[GET /iaas/api/load-balancers/{id}/network-interfaces/{id1}][%d] getLoadBalancerNetworkInterfaceNotFound  %+v", 404, o.Payload)
}

func (o *GetLoadBalancerNetworkInterfaceNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLoadBalancerNetworkInterfaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
