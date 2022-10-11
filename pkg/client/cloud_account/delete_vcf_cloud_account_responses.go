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

// DeleteVcfCloudAccountReader is a Reader for the DeleteVcfCloudAccount structure.
type DeleteVcfCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVcfCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteVcfCloudAccountAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteVcfCloudAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteVcfCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVcfCloudAccountAccepted creates a DeleteVcfCloudAccountAccepted with default headers values
func NewDeleteVcfCloudAccountAccepted() *DeleteVcfCloudAccountAccepted {
	return &DeleteVcfCloudAccountAccepted{}
}

/*
DeleteVcfCloudAccountAccepted describes a response with status code 202, with default header values.

successful operation
*/
type DeleteVcfCloudAccountAccepted struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this delete vcf cloud account accepted response has a 2xx status code
func (o *DeleteVcfCloudAccountAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete vcf cloud account accepted response has a 3xx status code
func (o *DeleteVcfCloudAccountAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete vcf cloud account accepted response has a 4xx status code
func (o *DeleteVcfCloudAccountAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete vcf cloud account accepted response has a 5xx status code
func (o *DeleteVcfCloudAccountAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete vcf cloud account accepted response a status code equal to that given
func (o *DeleteVcfCloudAccountAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DeleteVcfCloudAccountAccepted) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts-vcf/{id}][%d] deleteVcfCloudAccountAccepted  %+v", 202, o.Payload)
}

func (o *DeleteVcfCloudAccountAccepted) String() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts-vcf/{id}][%d] deleteVcfCloudAccountAccepted  %+v", 202, o.Payload)
}

func (o *DeleteVcfCloudAccountAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *DeleteVcfCloudAccountAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVcfCloudAccountNoContent creates a DeleteVcfCloudAccountNoContent with default headers values
func NewDeleteVcfCloudAccountNoContent() *DeleteVcfCloudAccountNoContent {
	return &DeleteVcfCloudAccountNoContent{}
}

/*
DeleteVcfCloudAccountNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteVcfCloudAccountNoContent struct {
}

// IsSuccess returns true when this delete vcf cloud account no content response has a 2xx status code
func (o *DeleteVcfCloudAccountNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete vcf cloud account no content response has a 3xx status code
func (o *DeleteVcfCloudAccountNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete vcf cloud account no content response has a 4xx status code
func (o *DeleteVcfCloudAccountNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete vcf cloud account no content response has a 5xx status code
func (o *DeleteVcfCloudAccountNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete vcf cloud account no content response a status code equal to that given
func (o *DeleteVcfCloudAccountNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteVcfCloudAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts-vcf/{id}][%d] deleteVcfCloudAccountNoContent ", 204)
}

func (o *DeleteVcfCloudAccountNoContent) String() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts-vcf/{id}][%d] deleteVcfCloudAccountNoContent ", 204)
}

func (o *DeleteVcfCloudAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVcfCloudAccountForbidden creates a DeleteVcfCloudAccountForbidden with default headers values
func NewDeleteVcfCloudAccountForbidden() *DeleteVcfCloudAccountForbidden {
	return &DeleteVcfCloudAccountForbidden{}
}

/*
DeleteVcfCloudAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteVcfCloudAccountForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this delete vcf cloud account forbidden response has a 2xx status code
func (o *DeleteVcfCloudAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete vcf cloud account forbidden response has a 3xx status code
func (o *DeleteVcfCloudAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete vcf cloud account forbidden response has a 4xx status code
func (o *DeleteVcfCloudAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete vcf cloud account forbidden response has a 5xx status code
func (o *DeleteVcfCloudAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete vcf cloud account forbidden response a status code equal to that given
func (o *DeleteVcfCloudAccountForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteVcfCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts-vcf/{id}][%d] deleteVcfCloudAccountForbidden  %+v", 403, o.Payload)
}

func (o *DeleteVcfCloudAccountForbidden) String() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts-vcf/{id}][%d] deleteVcfCloudAccountForbidden  %+v", 403, o.Payload)
}

func (o *DeleteVcfCloudAccountForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *DeleteVcfCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
