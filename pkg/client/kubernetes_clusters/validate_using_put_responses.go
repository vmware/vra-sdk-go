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

// ValidateUsingPUTReader is a Reader for the ValidateUsingPUT structure.
type ValidateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewValidateUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewValidateUsingPUTOK creates a ValidateUsingPUTOK with default headers values
func NewValidateUsingPUTOK() *ValidateUsingPUTOK {
	return &ValidateUsingPUTOK{}
}

/*
ValidateUsingPUTOK describes a response with status code 200, with default header values.

OK
*/
type ValidateUsingPUTOK struct {
	Payload *models.K8sClusterState
}

// IsSuccess returns true when this validate using p u t o k response has a 2xx status code
func (o *ValidateUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate using p u t o k response has a 3xx status code
func (o *ValidateUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate using p u t o k response has a 4xx status code
func (o *ValidateUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate using p u t o k response has a 5xx status code
func (o *ValidateUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate using p u t o k response a status code equal to that given
func (o *ValidateUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

func (o *ValidateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/k8s/clusters/validate][%d] validateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *ValidateUsingPUTOK) String() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/k8s/clusters/validate][%d] validateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *ValidateUsingPUTOK) GetPayload() *models.K8sClusterState {
	return o.Payload
}

func (o *ValidateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.K8sClusterState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateUsingPUTForbidden creates a ValidateUsingPUTForbidden with default headers values
func NewValidateUsingPUTForbidden() *ValidateUsingPUTForbidden {
	return &ValidateUsingPUTForbidden{}
}

/*
ValidateUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type ValidateUsingPUTForbidden struct {
}

// IsSuccess returns true when this validate using p u t forbidden response has a 2xx status code
func (o *ValidateUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this validate using p u t forbidden response has a 3xx status code
func (o *ValidateUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate using p u t forbidden response has a 4xx status code
func (o *ValidateUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this validate using p u t forbidden response has a 5xx status code
func (o *ValidateUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this validate using p u t forbidden response a status code equal to that given
func (o *ValidateUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ValidateUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/k8s/clusters/validate][%d] validateUsingPUTForbidden ", 403)
}

func (o *ValidateUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/k8s/clusters/validate][%d] validateUsingPUTForbidden ", 403)
}

func (o *ValidateUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
