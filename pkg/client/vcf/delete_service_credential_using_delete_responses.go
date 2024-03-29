// Code generated by go-swagger; DO NOT EDIT.

package vcf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteServiceCredentialUsingDELETEReader is a Reader for the DeleteServiceCredentialUsingDELETE structure.
type DeleteServiceCredentialUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceCredentialUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteServiceCredentialUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteServiceCredentialUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteServiceCredentialUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteServiceCredentialUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteServiceCredentialUsingDELETENoContent creates a DeleteServiceCredentialUsingDELETENoContent with default headers values
func NewDeleteServiceCredentialUsingDELETENoContent() *DeleteServiceCredentialUsingDELETENoContent {
	return &DeleteServiceCredentialUsingDELETENoContent{}
}

/*
DeleteServiceCredentialUsingDELETENoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteServiceCredentialUsingDELETENoContent struct {
}

// IsSuccess returns true when this delete service credential using d e l e t e no content response has a 2xx status code
func (o *DeleteServiceCredentialUsingDELETENoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete service credential using d e l e t e no content response has a 3xx status code
func (o *DeleteServiceCredentialUsingDELETENoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service credential using d e l e t e no content response has a 4xx status code
func (o *DeleteServiceCredentialUsingDELETENoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete service credential using d e l e t e no content response has a 5xx status code
func (o *DeleteServiceCredentialUsingDELETENoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service credential using d e l e t e no content response a status code equal to that given
func (o *DeleteServiceCredentialUsingDELETENoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteServiceCredentialUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /content/api/vcf/{integrationId}/domain/{domainId}/service-accounts/{id}][%d] deleteServiceCredentialUsingDELETENoContent ", 204)
}

func (o *DeleteServiceCredentialUsingDELETENoContent) String() string {
	return fmt.Sprintf("[DELETE /content/api/vcf/{integrationId}/domain/{domainId}/service-accounts/{id}][%d] deleteServiceCredentialUsingDELETENoContent ", 204)
}

func (o *DeleteServiceCredentialUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceCredentialUsingDELETEUnauthorized creates a DeleteServiceCredentialUsingDELETEUnauthorized with default headers values
func NewDeleteServiceCredentialUsingDELETEUnauthorized() *DeleteServiceCredentialUsingDELETEUnauthorized {
	return &DeleteServiceCredentialUsingDELETEUnauthorized{}
}

/*
DeleteServiceCredentialUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteServiceCredentialUsingDELETEUnauthorized struct {
}

// IsSuccess returns true when this delete service credential using d e l e t e unauthorized response has a 2xx status code
func (o *DeleteServiceCredentialUsingDELETEUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete service credential using d e l e t e unauthorized response has a 3xx status code
func (o *DeleteServiceCredentialUsingDELETEUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service credential using d e l e t e unauthorized response has a 4xx status code
func (o *DeleteServiceCredentialUsingDELETEUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete service credential using d e l e t e unauthorized response has a 5xx status code
func (o *DeleteServiceCredentialUsingDELETEUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service credential using d e l e t e unauthorized response a status code equal to that given
func (o *DeleteServiceCredentialUsingDELETEUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteServiceCredentialUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /content/api/vcf/{integrationId}/domain/{domainId}/service-accounts/{id}][%d] deleteServiceCredentialUsingDELETEUnauthorized ", 401)
}

func (o *DeleteServiceCredentialUsingDELETEUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /content/api/vcf/{integrationId}/domain/{domainId}/service-accounts/{id}][%d] deleteServiceCredentialUsingDELETEUnauthorized ", 401)
}

func (o *DeleteServiceCredentialUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceCredentialUsingDELETEForbidden creates a DeleteServiceCredentialUsingDELETEForbidden with default headers values
func NewDeleteServiceCredentialUsingDELETEForbidden() *DeleteServiceCredentialUsingDELETEForbidden {
	return &DeleteServiceCredentialUsingDELETEForbidden{}
}

/*
DeleteServiceCredentialUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteServiceCredentialUsingDELETEForbidden struct {
}

// IsSuccess returns true when this delete service credential using d e l e t e forbidden response has a 2xx status code
func (o *DeleteServiceCredentialUsingDELETEForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete service credential using d e l e t e forbidden response has a 3xx status code
func (o *DeleteServiceCredentialUsingDELETEForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service credential using d e l e t e forbidden response has a 4xx status code
func (o *DeleteServiceCredentialUsingDELETEForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete service credential using d e l e t e forbidden response has a 5xx status code
func (o *DeleteServiceCredentialUsingDELETEForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service credential using d e l e t e forbidden response a status code equal to that given
func (o *DeleteServiceCredentialUsingDELETEForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteServiceCredentialUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /content/api/vcf/{integrationId}/domain/{domainId}/service-accounts/{id}][%d] deleteServiceCredentialUsingDELETEForbidden ", 403)
}

func (o *DeleteServiceCredentialUsingDELETEForbidden) String() string {
	return fmt.Sprintf("[DELETE /content/api/vcf/{integrationId}/domain/{domainId}/service-accounts/{id}][%d] deleteServiceCredentialUsingDELETEForbidden ", 403)
}

func (o *DeleteServiceCredentialUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceCredentialUsingDELETENotFound creates a DeleteServiceCredentialUsingDELETENotFound with default headers values
func NewDeleteServiceCredentialUsingDELETENotFound() *DeleteServiceCredentialUsingDELETENotFound {
	return &DeleteServiceCredentialUsingDELETENotFound{}
}

/*
DeleteServiceCredentialUsingDELETENotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteServiceCredentialUsingDELETENotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete service credential using d e l e t e not found response has a 2xx status code
func (o *DeleteServiceCredentialUsingDELETENotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete service credential using d e l e t e not found response has a 3xx status code
func (o *DeleteServiceCredentialUsingDELETENotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service credential using d e l e t e not found response has a 4xx status code
func (o *DeleteServiceCredentialUsingDELETENotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete service credential using d e l e t e not found response has a 5xx status code
func (o *DeleteServiceCredentialUsingDELETENotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service credential using d e l e t e not found response a status code equal to that given
func (o *DeleteServiceCredentialUsingDELETENotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteServiceCredentialUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /content/api/vcf/{integrationId}/domain/{domainId}/service-accounts/{id}][%d] deleteServiceCredentialUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteServiceCredentialUsingDELETENotFound) String() string {
	return fmt.Sprintf("[DELETE /content/api/vcf/{integrationId}/domain/{domainId}/service-accounts/{id}][%d] deleteServiceCredentialUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteServiceCredentialUsingDELETENotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServiceCredentialUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
