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

// DeleteTerraformVersionUsingDELETE1Reader is a Reader for the DeleteTerraformVersionUsingDELETE1 structure.
type DeleteTerraformVersionUsingDELETE1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTerraformVersionUsingDELETE1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteTerraformVersionUsingDELETE1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteTerraformVersionUsingDELETE1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTerraformVersionUsingDELETE1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTerraformVersionUsingDELETE1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTerraformVersionUsingDELETE1NoContent creates a DeleteTerraformVersionUsingDELETE1NoContent with default headers values
func NewDeleteTerraformVersionUsingDELETE1NoContent() *DeleteTerraformVersionUsingDELETE1NoContent {
	return &DeleteTerraformVersionUsingDELETE1NoContent{}
}

/*DeleteTerraformVersionUsingDELETE1NoContent handles this case with default header values.

No Content
*/
type DeleteTerraformVersionUsingDELETE1NoContent struct {
}

func (o *DeleteTerraformVersionUsingDELETE1NoContent) Error() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] deleteTerraformVersionUsingDELETE1NoContent ", 204)
}

func (o *DeleteTerraformVersionUsingDELETE1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTerraformVersionUsingDELETE1Unauthorized creates a DeleteTerraformVersionUsingDELETE1Unauthorized with default headers values
func NewDeleteTerraformVersionUsingDELETE1Unauthorized() *DeleteTerraformVersionUsingDELETE1Unauthorized {
	return &DeleteTerraformVersionUsingDELETE1Unauthorized{}
}

/*DeleteTerraformVersionUsingDELETE1Unauthorized handles this case with default header values.

Unauthorized
*/
type DeleteTerraformVersionUsingDELETE1Unauthorized struct {
}

func (o *DeleteTerraformVersionUsingDELETE1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] deleteTerraformVersionUsingDELETE1Unauthorized ", 401)
}

func (o *DeleteTerraformVersionUsingDELETE1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTerraformVersionUsingDELETE1Forbidden creates a DeleteTerraformVersionUsingDELETE1Forbidden with default headers values
func NewDeleteTerraformVersionUsingDELETE1Forbidden() *DeleteTerraformVersionUsingDELETE1Forbidden {
	return &DeleteTerraformVersionUsingDELETE1Forbidden{}
}

/*DeleteTerraformVersionUsingDELETE1Forbidden handles this case with default header values.

Forbidden
*/
type DeleteTerraformVersionUsingDELETE1Forbidden struct {
}

func (o *DeleteTerraformVersionUsingDELETE1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] deleteTerraformVersionUsingDELETE1Forbidden ", 403)
}

func (o *DeleteTerraformVersionUsingDELETE1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTerraformVersionUsingDELETE1NotFound creates a DeleteTerraformVersionUsingDELETE1NotFound with default headers values
func NewDeleteTerraformVersionUsingDELETE1NotFound() *DeleteTerraformVersionUsingDELETE1NotFound {
	return &DeleteTerraformVersionUsingDELETE1NotFound{}
}

/*DeleteTerraformVersionUsingDELETE1NotFound handles this case with default header values.

Not Found
*/
type DeleteTerraformVersionUsingDELETE1NotFound struct {
	Payload *models.Error
}

func (o *DeleteTerraformVersionUsingDELETE1NotFound) Error() string {
	return fmt.Sprintf("[DELETE /blueprint/api/blueprint-integrations/terraform/versions/{versionId}][%d] deleteTerraformVersionUsingDELETE1NotFound  %+v", 404, o.Payload)
}

func (o *DeleteTerraformVersionUsingDELETE1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTerraformVersionUsingDELETE1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
