// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// EnumeratePrivateImagesGcpReader is a Reader for the EnumeratePrivateImagesGcp structure.
type EnumeratePrivateImagesGcpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnumeratePrivateImagesGcpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewEnumeratePrivateImagesGcpAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnumeratePrivateImagesGcpBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnumeratePrivateImagesGcpForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnumeratePrivateImagesGcpAccepted creates a EnumeratePrivateImagesGcpAccepted with default headers values
func NewEnumeratePrivateImagesGcpAccepted() *EnumeratePrivateImagesGcpAccepted {
	return &EnumeratePrivateImagesGcpAccepted{}
}

/*
EnumeratePrivateImagesGcpAccepted describes a response with status code 202, with default header values.

successful operation
*/
type EnumeratePrivateImagesGcpAccepted struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this enumerate private images gcp accepted response has a 2xx status code
func (o *EnumeratePrivateImagesGcpAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enumerate private images gcp accepted response has a 3xx status code
func (o *EnumeratePrivateImagesGcpAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enumerate private images gcp accepted response has a 4xx status code
func (o *EnumeratePrivateImagesGcpAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this enumerate private images gcp accepted response has a 5xx status code
func (o *EnumeratePrivateImagesGcpAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this enumerate private images gcp accepted response a status code equal to that given
func (o *EnumeratePrivateImagesGcpAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *EnumeratePrivateImagesGcpAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-gcp/{id}/private-image-enumeration][%d] enumeratePrivateImagesGcpAccepted  %+v", 202, o.Payload)
}

func (o *EnumeratePrivateImagesGcpAccepted) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-gcp/{id}/private-image-enumeration][%d] enumeratePrivateImagesGcpAccepted  %+v", 202, o.Payload)
}

func (o *EnumeratePrivateImagesGcpAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *EnumeratePrivateImagesGcpAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumeratePrivateImagesGcpBadRequest creates a EnumeratePrivateImagesGcpBadRequest with default headers values
func NewEnumeratePrivateImagesGcpBadRequest() *EnumeratePrivateImagesGcpBadRequest {
	return &EnumeratePrivateImagesGcpBadRequest{}
}

/*
EnumeratePrivateImagesGcpBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type EnumeratePrivateImagesGcpBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this enumerate private images gcp bad request response has a 2xx status code
func (o *EnumeratePrivateImagesGcpBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enumerate private images gcp bad request response has a 3xx status code
func (o *EnumeratePrivateImagesGcpBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enumerate private images gcp bad request response has a 4xx status code
func (o *EnumeratePrivateImagesGcpBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this enumerate private images gcp bad request response has a 5xx status code
func (o *EnumeratePrivateImagesGcpBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this enumerate private images gcp bad request response a status code equal to that given
func (o *EnumeratePrivateImagesGcpBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EnumeratePrivateImagesGcpBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-gcp/{id}/private-image-enumeration][%d] enumeratePrivateImagesGcpBadRequest  %+v", 400, o.Payload)
}

func (o *EnumeratePrivateImagesGcpBadRequest) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-gcp/{id}/private-image-enumeration][%d] enumeratePrivateImagesGcpBadRequest  %+v", 400, o.Payload)
}

func (o *EnumeratePrivateImagesGcpBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *EnumeratePrivateImagesGcpBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumeratePrivateImagesGcpForbidden creates a EnumeratePrivateImagesGcpForbidden with default headers values
func NewEnumeratePrivateImagesGcpForbidden() *EnumeratePrivateImagesGcpForbidden {
	return &EnumeratePrivateImagesGcpForbidden{}
}

/*
EnumeratePrivateImagesGcpForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type EnumeratePrivateImagesGcpForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this enumerate private images gcp forbidden response has a 2xx status code
func (o *EnumeratePrivateImagesGcpForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enumerate private images gcp forbidden response has a 3xx status code
func (o *EnumeratePrivateImagesGcpForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enumerate private images gcp forbidden response has a 4xx status code
func (o *EnumeratePrivateImagesGcpForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this enumerate private images gcp forbidden response has a 5xx status code
func (o *EnumeratePrivateImagesGcpForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this enumerate private images gcp forbidden response a status code equal to that given
func (o *EnumeratePrivateImagesGcpForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EnumeratePrivateImagesGcpForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-gcp/{id}/private-image-enumeration][%d] enumeratePrivateImagesGcpForbidden  %+v", 403, o.Payload)
}

func (o *EnumeratePrivateImagesGcpForbidden) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-gcp/{id}/private-image-enumeration][%d] enumeratePrivateImagesGcpForbidden  %+v", 403, o.Payload)
}

func (o *EnumeratePrivateImagesGcpForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *EnumeratePrivateImagesGcpForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
