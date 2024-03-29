// Code generated by go-swagger; DO NOT EDIT.

package blueprint_terraform_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteTerraformVersionUsingDELETEReader is a Reader for the DeleteTerraformVersionUsingDELETE structure.
type DeleteTerraformVersionUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTerraformVersionUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteTerraformVersionUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteTerraformVersionUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTerraformVersionUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTerraformVersionUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTerraformVersionUsingDELETENoContent creates a DeleteTerraformVersionUsingDELETENoContent with default headers values
func NewDeleteTerraformVersionUsingDELETENoContent() *DeleteTerraformVersionUsingDELETENoContent {
	return &DeleteTerraformVersionUsingDELETENoContent{}
}

/*
DeleteTerraformVersionUsingDELETENoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteTerraformVersionUsingDELETENoContent struct {
}

// IsSuccess returns true when this delete terraform version using d e l e t e no content response has a 2xx status code
func (o *DeleteTerraformVersionUsingDELETENoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete terraform version using d e l e t e no content response has a 3xx status code
func (o *DeleteTerraformVersionUsingDELETENoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete terraform version using d e l e t e no content response has a 4xx status code
func (o *DeleteTerraformVersionUsingDELETENoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete terraform version using d e l e t e no content response has a 5xx status code
func (o *DeleteTerraformVersionUsingDELETENoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete terraform version using d e l e t e no content response a status code equal to that given
func (o *DeleteTerraformVersionUsingDELETENoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteTerraformVersionUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] deleteTerraformVersionUsingDELETENoContent ", 204)
}

func (o *DeleteTerraformVersionUsingDELETENoContent) String() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] deleteTerraformVersionUsingDELETENoContent ", 204)
}

func (o *DeleteTerraformVersionUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTerraformVersionUsingDELETEUnauthorized creates a DeleteTerraformVersionUsingDELETEUnauthorized with default headers values
func NewDeleteTerraformVersionUsingDELETEUnauthorized() *DeleteTerraformVersionUsingDELETEUnauthorized {
	return &DeleteTerraformVersionUsingDELETEUnauthorized{}
}

/*
DeleteTerraformVersionUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteTerraformVersionUsingDELETEUnauthorized struct {
}

// IsSuccess returns true when this delete terraform version using d e l e t e unauthorized response has a 2xx status code
func (o *DeleteTerraformVersionUsingDELETEUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete terraform version using d e l e t e unauthorized response has a 3xx status code
func (o *DeleteTerraformVersionUsingDELETEUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete terraform version using d e l e t e unauthorized response has a 4xx status code
func (o *DeleteTerraformVersionUsingDELETEUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete terraform version using d e l e t e unauthorized response has a 5xx status code
func (o *DeleteTerraformVersionUsingDELETEUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete terraform version using d e l e t e unauthorized response a status code equal to that given
func (o *DeleteTerraformVersionUsingDELETEUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteTerraformVersionUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] deleteTerraformVersionUsingDELETEUnauthorized ", 401)
}

func (o *DeleteTerraformVersionUsingDELETEUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] deleteTerraformVersionUsingDELETEUnauthorized ", 401)
}

func (o *DeleteTerraformVersionUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTerraformVersionUsingDELETEForbidden creates a DeleteTerraformVersionUsingDELETEForbidden with default headers values
func NewDeleteTerraformVersionUsingDELETEForbidden() *DeleteTerraformVersionUsingDELETEForbidden {
	return &DeleteTerraformVersionUsingDELETEForbidden{}
}

/*
DeleteTerraformVersionUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteTerraformVersionUsingDELETEForbidden struct {
}

// IsSuccess returns true when this delete terraform version using d e l e t e forbidden response has a 2xx status code
func (o *DeleteTerraformVersionUsingDELETEForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete terraform version using d e l e t e forbidden response has a 3xx status code
func (o *DeleteTerraformVersionUsingDELETEForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete terraform version using d e l e t e forbidden response has a 4xx status code
func (o *DeleteTerraformVersionUsingDELETEForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete terraform version using d e l e t e forbidden response has a 5xx status code
func (o *DeleteTerraformVersionUsingDELETEForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete terraform version using d e l e t e forbidden response a status code equal to that given
func (o *DeleteTerraformVersionUsingDELETEForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteTerraformVersionUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] deleteTerraformVersionUsingDELETEForbidden ", 403)
}

func (o *DeleteTerraformVersionUsingDELETEForbidden) String() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] deleteTerraformVersionUsingDELETEForbidden ", 403)
}

func (o *DeleteTerraformVersionUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTerraformVersionUsingDELETENotFound creates a DeleteTerraformVersionUsingDELETENotFound with default headers values
func NewDeleteTerraformVersionUsingDELETENotFound() *DeleteTerraformVersionUsingDELETENotFound {
	return &DeleteTerraformVersionUsingDELETENotFound{}
}

/*
DeleteTerraformVersionUsingDELETENotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteTerraformVersionUsingDELETENotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete terraform version using d e l e t e not found response has a 2xx status code
func (o *DeleteTerraformVersionUsingDELETENotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete terraform version using d e l e t e not found response has a 3xx status code
func (o *DeleteTerraformVersionUsingDELETENotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete terraform version using d e l e t e not found response has a 4xx status code
func (o *DeleteTerraformVersionUsingDELETENotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete terraform version using d e l e t e not found response has a 5xx status code
func (o *DeleteTerraformVersionUsingDELETENotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete terraform version using d e l e t e not found response a status code equal to that given
func (o *DeleteTerraformVersionUsingDELETENotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteTerraformVersionUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] deleteTerraformVersionUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteTerraformVersionUsingDELETENotFound) String() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] deleteTerraformVersionUsingDELETENotFound  %+v", 404, o.Payload)
}

func (o *DeleteTerraformVersionUsingDELETENotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTerraformVersionUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
