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

// CreateLoadBalancerReader is a Reader for the CreateLoadBalancer structure.
type CreateLoadBalancerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLoadBalancerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCreateLoadBalancerAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateLoadBalancerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateLoadBalancerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateLoadBalancerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateLoadBalancerAccepted creates a CreateLoadBalancerAccepted with default headers values
func NewCreateLoadBalancerAccepted() *CreateLoadBalancerAccepted {
	return &CreateLoadBalancerAccepted{}
}

/*
CreateLoadBalancerAccepted describes a response with status code 202, with default header values.

successful operation
*/
type CreateLoadBalancerAccepted struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this create load balancer accepted response has a 2xx status code
func (o *CreateLoadBalancerAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create load balancer accepted response has a 3xx status code
func (o *CreateLoadBalancerAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create load balancer accepted response has a 4xx status code
func (o *CreateLoadBalancerAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create load balancer accepted response has a 5xx status code
func (o *CreateLoadBalancerAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create load balancer accepted response a status code equal to that given
func (o *CreateLoadBalancerAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateLoadBalancerAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/load-balancers][%d] createLoadBalancerAccepted  %+v", 202, o.Payload)
}

func (o *CreateLoadBalancerAccepted) String() string {
	return fmt.Sprintf("[POST /iaas/api/load-balancers][%d] createLoadBalancerAccepted  %+v", 202, o.Payload)
}

func (o *CreateLoadBalancerAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *CreateLoadBalancerAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLoadBalancerBadRequest creates a CreateLoadBalancerBadRequest with default headers values
func NewCreateLoadBalancerBadRequest() *CreateLoadBalancerBadRequest {
	return &CreateLoadBalancerBadRequest{}
}

/*
CreateLoadBalancerBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type CreateLoadBalancerBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create load balancer bad request response has a 2xx status code
func (o *CreateLoadBalancerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create load balancer bad request response has a 3xx status code
func (o *CreateLoadBalancerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create load balancer bad request response has a 4xx status code
func (o *CreateLoadBalancerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create load balancer bad request response has a 5xx status code
func (o *CreateLoadBalancerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create load balancer bad request response a status code equal to that given
func (o *CreateLoadBalancerBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateLoadBalancerBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/load-balancers][%d] createLoadBalancerBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLoadBalancerBadRequest) String() string {
	return fmt.Sprintf("[POST /iaas/api/load-balancers][%d] createLoadBalancerBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLoadBalancerBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateLoadBalancerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLoadBalancerForbidden creates a CreateLoadBalancerForbidden with default headers values
func NewCreateLoadBalancerForbidden() *CreateLoadBalancerForbidden {
	return &CreateLoadBalancerForbidden{}
}

/*
CreateLoadBalancerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateLoadBalancerForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this create load balancer forbidden response has a 2xx status code
func (o *CreateLoadBalancerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create load balancer forbidden response has a 3xx status code
func (o *CreateLoadBalancerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create load balancer forbidden response has a 4xx status code
func (o *CreateLoadBalancerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create load balancer forbidden response has a 5xx status code
func (o *CreateLoadBalancerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create load balancer forbidden response a status code equal to that given
func (o *CreateLoadBalancerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateLoadBalancerForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/load-balancers][%d] createLoadBalancerForbidden  %+v", 403, o.Payload)
}

func (o *CreateLoadBalancerForbidden) String() string {
	return fmt.Sprintf("[POST /iaas/api/load-balancers][%d] createLoadBalancerForbidden  %+v", 403, o.Payload)
}

func (o *CreateLoadBalancerForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *CreateLoadBalancerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLoadBalancerNotFound creates a CreateLoadBalancerNotFound with default headers values
func NewCreateLoadBalancerNotFound() *CreateLoadBalancerNotFound {
	return &CreateLoadBalancerNotFound{}
}

/*
CreateLoadBalancerNotFound describes a response with status code 404, with default header values.

Input(s) not Found
*/
type CreateLoadBalancerNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this create load balancer not found response has a 2xx status code
func (o *CreateLoadBalancerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create load balancer not found response has a 3xx status code
func (o *CreateLoadBalancerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create load balancer not found response has a 4xx status code
func (o *CreateLoadBalancerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create load balancer not found response has a 5xx status code
func (o *CreateLoadBalancerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create load balancer not found response a status code equal to that given
func (o *CreateLoadBalancerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateLoadBalancerNotFound) Error() string {
	return fmt.Sprintf("[POST /iaas/api/load-balancers][%d] createLoadBalancerNotFound  %+v", 404, o.Payload)
}

func (o *CreateLoadBalancerNotFound) String() string {
	return fmt.Sprintf("[POST /iaas/api/load-balancers][%d] createLoadBalancerNotFound  %+v", 404, o.Payload)
}

func (o *CreateLoadBalancerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateLoadBalancerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
