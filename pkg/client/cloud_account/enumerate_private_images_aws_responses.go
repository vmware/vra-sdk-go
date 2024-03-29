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

// EnumeratePrivateImagesAWSReader is a Reader for the EnumeratePrivateImagesAWS structure.
type EnumeratePrivateImagesAWSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnumeratePrivateImagesAWSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewEnumeratePrivateImagesAWSAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnumeratePrivateImagesAWSBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnumeratePrivateImagesAWSForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnumeratePrivateImagesAWSAccepted creates a EnumeratePrivateImagesAWSAccepted with default headers values
func NewEnumeratePrivateImagesAWSAccepted() *EnumeratePrivateImagesAWSAccepted {
	return &EnumeratePrivateImagesAWSAccepted{}
}

/*
EnumeratePrivateImagesAWSAccepted describes a response with status code 202, with default header values.

successful operation
*/
type EnumeratePrivateImagesAWSAccepted struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this enumerate private images a w s accepted response has a 2xx status code
func (o *EnumeratePrivateImagesAWSAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enumerate private images a w s accepted response has a 3xx status code
func (o *EnumeratePrivateImagesAWSAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enumerate private images a w s accepted response has a 4xx status code
func (o *EnumeratePrivateImagesAWSAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this enumerate private images a w s accepted response has a 5xx status code
func (o *EnumeratePrivateImagesAWSAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this enumerate private images a w s accepted response a status code equal to that given
func (o *EnumeratePrivateImagesAWSAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *EnumeratePrivateImagesAWSAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-aws/{id}/private-image-enumeration][%d] enumeratePrivateImagesAWSAccepted  %+v", 202, o.Payload)
}

func (o *EnumeratePrivateImagesAWSAccepted) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-aws/{id}/private-image-enumeration][%d] enumeratePrivateImagesAWSAccepted  %+v", 202, o.Payload)
}

func (o *EnumeratePrivateImagesAWSAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *EnumeratePrivateImagesAWSAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumeratePrivateImagesAWSBadRequest creates a EnumeratePrivateImagesAWSBadRequest with default headers values
func NewEnumeratePrivateImagesAWSBadRequest() *EnumeratePrivateImagesAWSBadRequest {
	return &EnumeratePrivateImagesAWSBadRequest{}
}

/*
EnumeratePrivateImagesAWSBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type EnumeratePrivateImagesAWSBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this enumerate private images a w s bad request response has a 2xx status code
func (o *EnumeratePrivateImagesAWSBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enumerate private images a w s bad request response has a 3xx status code
func (o *EnumeratePrivateImagesAWSBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enumerate private images a w s bad request response has a 4xx status code
func (o *EnumeratePrivateImagesAWSBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this enumerate private images a w s bad request response has a 5xx status code
func (o *EnumeratePrivateImagesAWSBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this enumerate private images a w s bad request response a status code equal to that given
func (o *EnumeratePrivateImagesAWSBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EnumeratePrivateImagesAWSBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-aws/{id}/private-image-enumeration][%d] enumeratePrivateImagesAWSBadRequest  %+v", 400, o.Payload)
}

func (o *EnumeratePrivateImagesAWSBadRequest) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-aws/{id}/private-image-enumeration][%d] enumeratePrivateImagesAWSBadRequest  %+v", 400, o.Payload)
}

func (o *EnumeratePrivateImagesAWSBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *EnumeratePrivateImagesAWSBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumeratePrivateImagesAWSForbidden creates a EnumeratePrivateImagesAWSForbidden with default headers values
func NewEnumeratePrivateImagesAWSForbidden() *EnumeratePrivateImagesAWSForbidden {
	return &EnumeratePrivateImagesAWSForbidden{}
}

/*
EnumeratePrivateImagesAWSForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type EnumeratePrivateImagesAWSForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this enumerate private images a w s forbidden response has a 2xx status code
func (o *EnumeratePrivateImagesAWSForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enumerate private images a w s forbidden response has a 3xx status code
func (o *EnumeratePrivateImagesAWSForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enumerate private images a w s forbidden response has a 4xx status code
func (o *EnumeratePrivateImagesAWSForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this enumerate private images a w s forbidden response has a 5xx status code
func (o *EnumeratePrivateImagesAWSForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this enumerate private images a w s forbidden response a status code equal to that given
func (o *EnumeratePrivateImagesAWSForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EnumeratePrivateImagesAWSForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-aws/{id}/private-image-enumeration][%d] enumeratePrivateImagesAWSForbidden  %+v", 403, o.Payload)
}

func (o *EnumeratePrivateImagesAWSForbidden) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-aws/{id}/private-image-enumeration][%d] enumeratePrivateImagesAWSForbidden  %+v", 403, o.Payload)
}

func (o *EnumeratePrivateImagesAWSForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *EnumeratePrivateImagesAWSForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
