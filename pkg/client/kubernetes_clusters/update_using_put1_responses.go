// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateUsingPUT1Reader is a Reader for the UpdateUsingPUT1 structure.
type UpdateUsingPUT1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUsingPUT1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUsingPUT1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateUsingPUT1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUsingPUT1OK creates a UpdateUsingPUT1OK with default headers values
func NewUpdateUsingPUT1OK() *UpdateUsingPUT1OK {
	return &UpdateUsingPUT1OK{}
}

/*
UpdateUsingPUT1OK describes a response with status code 200, with default header values.

OK
*/
type UpdateUsingPUT1OK struct {
	Payload *models.K8SCluster
}

// IsSuccess returns true when this update using p u t1 o k response has a 2xx status code
func (o *UpdateUsingPUT1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update using p u t1 o k response has a 3xx status code
func (o *UpdateUsingPUT1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update using p u t1 o k response has a 4xx status code
func (o *UpdateUsingPUT1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update using p u t1 o k response has a 5xx status code
func (o *UpdateUsingPUT1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this update using p u t1 o k response a status code equal to that given
func (o *UpdateUsingPUT1OK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateUsingPUT1OK) Error() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/k8s/clusters/{id}][%d] updateUsingPUT1OK  %+v", 200, o.Payload)
}

func (o *UpdateUsingPUT1OK) String() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/k8s/clusters/{id}][%d] updateUsingPUT1OK  %+v", 200, o.Payload)
}

func (o *UpdateUsingPUT1OK) GetPayload() *models.K8SCluster {
	return o.Payload
}

func (o *UpdateUsingPUT1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.K8SCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUsingPUT1Forbidden creates a UpdateUsingPUT1Forbidden with default headers values
func NewUpdateUsingPUT1Forbidden() *UpdateUsingPUT1Forbidden {
	return &UpdateUsingPUT1Forbidden{}
}

/*
UpdateUsingPUT1Forbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type UpdateUsingPUT1Forbidden struct {
}

// IsSuccess returns true when this update using p u t1 forbidden response has a 2xx status code
func (o *UpdateUsingPUT1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update using p u t1 forbidden response has a 3xx status code
func (o *UpdateUsingPUT1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update using p u t1 forbidden response has a 4xx status code
func (o *UpdateUsingPUT1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update using p u t1 forbidden response has a 5xx status code
func (o *UpdateUsingPUT1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update using p u t1 forbidden response a status code equal to that given
func (o *UpdateUsingPUT1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateUsingPUT1Forbidden) Error() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/k8s/clusters/{id}][%d] updateUsingPUT1Forbidden ", 403)
}

func (o *UpdateUsingPUT1Forbidden) String() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/k8s/clusters/{id}][%d] updateUsingPUT1Forbidden ", 403)
}

func (o *UpdateUsingPUT1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}