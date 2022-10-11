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

// CreateTerraformVersionUsingPOSTReader is a Reader for the CreateTerraformVersionUsingPOST structure.
type CreateTerraformVersionUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTerraformVersionUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTerraformVersionUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTerraformVersionUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateTerraformVersionUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateTerraformVersionUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTerraformVersionUsingPOSTCreated creates a CreateTerraformVersionUsingPOSTCreated with default headers values
func NewCreateTerraformVersionUsingPOSTCreated() *CreateTerraformVersionUsingPOSTCreated {
	return &CreateTerraformVersionUsingPOSTCreated{}
}

/*
CreateTerraformVersionUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type CreateTerraformVersionUsingPOSTCreated struct {
	Payload *models.TerraformVersion
}

// IsSuccess returns true when this create terraform version using p o s t created response has a 2xx status code
func (o *CreateTerraformVersionUsingPOSTCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create terraform version using p o s t created response has a 3xx status code
func (o *CreateTerraformVersionUsingPOSTCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create terraform version using p o s t created response has a 4xx status code
func (o *CreateTerraformVersionUsingPOSTCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create terraform version using p o s t created response has a 5xx status code
func (o *CreateTerraformVersionUsingPOSTCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create terraform version using p o s t created response a status code equal to that given
func (o *CreateTerraformVersionUsingPOSTCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateTerraformVersionUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprint-integrations/terraform/versions][%d] createTerraformVersionUsingPOSTCreated  %+v", 201, o.Payload)
}

func (o *CreateTerraformVersionUsingPOSTCreated) String() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprint-integrations/terraform/versions][%d] createTerraformVersionUsingPOSTCreated  %+v", 201, o.Payload)
}

func (o *CreateTerraformVersionUsingPOSTCreated) GetPayload() *models.TerraformVersion {
	return o.Payload
}

func (o *CreateTerraformVersionUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TerraformVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTerraformVersionUsingPOSTBadRequest creates a CreateTerraformVersionUsingPOSTBadRequest with default headers values
func NewCreateTerraformVersionUsingPOSTBadRequest() *CreateTerraformVersionUsingPOSTBadRequest {
	return &CreateTerraformVersionUsingPOSTBadRequest{}
}

/*
CreateTerraformVersionUsingPOSTBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateTerraformVersionUsingPOSTBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create terraform version using p o s t bad request response has a 2xx status code
func (o *CreateTerraformVersionUsingPOSTBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create terraform version using p o s t bad request response has a 3xx status code
func (o *CreateTerraformVersionUsingPOSTBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create terraform version using p o s t bad request response has a 4xx status code
func (o *CreateTerraformVersionUsingPOSTBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create terraform version using p o s t bad request response has a 5xx status code
func (o *CreateTerraformVersionUsingPOSTBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create terraform version using p o s t bad request response a status code equal to that given
func (o *CreateTerraformVersionUsingPOSTBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateTerraformVersionUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprint-integrations/terraform/versions][%d] createTerraformVersionUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTerraformVersionUsingPOSTBadRequest) String() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprint-integrations/terraform/versions][%d] createTerraformVersionUsingPOSTBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTerraformVersionUsingPOSTBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTerraformVersionUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTerraformVersionUsingPOSTUnauthorized creates a CreateTerraformVersionUsingPOSTUnauthorized with default headers values
func NewCreateTerraformVersionUsingPOSTUnauthorized() *CreateTerraformVersionUsingPOSTUnauthorized {
	return &CreateTerraformVersionUsingPOSTUnauthorized{}
}

/*
CreateTerraformVersionUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateTerraformVersionUsingPOSTUnauthorized struct {
}

// IsSuccess returns true when this create terraform version using p o s t unauthorized response has a 2xx status code
func (o *CreateTerraformVersionUsingPOSTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create terraform version using p o s t unauthorized response has a 3xx status code
func (o *CreateTerraformVersionUsingPOSTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create terraform version using p o s t unauthorized response has a 4xx status code
func (o *CreateTerraformVersionUsingPOSTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create terraform version using p o s t unauthorized response has a 5xx status code
func (o *CreateTerraformVersionUsingPOSTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create terraform version using p o s t unauthorized response a status code equal to that given
func (o *CreateTerraformVersionUsingPOSTUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateTerraformVersionUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprint-integrations/terraform/versions][%d] createTerraformVersionUsingPOSTUnauthorized ", 401)
}

func (o *CreateTerraformVersionUsingPOSTUnauthorized) String() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprint-integrations/terraform/versions][%d] createTerraformVersionUsingPOSTUnauthorized ", 401)
}

func (o *CreateTerraformVersionUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTerraformVersionUsingPOSTForbidden creates a CreateTerraformVersionUsingPOSTForbidden with default headers values
func NewCreateTerraformVersionUsingPOSTForbidden() *CreateTerraformVersionUsingPOSTForbidden {
	return &CreateTerraformVersionUsingPOSTForbidden{}
}

/*
CreateTerraformVersionUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateTerraformVersionUsingPOSTForbidden struct {
}

// IsSuccess returns true when this create terraform version using p o s t forbidden response has a 2xx status code
func (o *CreateTerraformVersionUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create terraform version using p o s t forbidden response has a 3xx status code
func (o *CreateTerraformVersionUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create terraform version using p o s t forbidden response has a 4xx status code
func (o *CreateTerraformVersionUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create terraform version using p o s t forbidden response has a 5xx status code
func (o *CreateTerraformVersionUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create terraform version using p o s t forbidden response a status code equal to that given
func (o *CreateTerraformVersionUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateTerraformVersionUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprint-integrations/terraform/versions][%d] createTerraformVersionUsingPOSTForbidden ", 403)
}

func (o *CreateTerraformVersionUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprint-integrations/terraform/versions][%d] createTerraformVersionUsingPOSTForbidden ", 403)
}

func (o *CreateTerraformVersionUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}