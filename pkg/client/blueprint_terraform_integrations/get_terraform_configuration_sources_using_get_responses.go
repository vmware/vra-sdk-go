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

// GetTerraformConfigurationSourcesUsingGETReader is a Reader for the GetTerraformConfigurationSourcesUsingGET structure.
type GetTerraformConfigurationSourcesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTerraformConfigurationSourcesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTerraformConfigurationSourcesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTerraformConfigurationSourcesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTerraformConfigurationSourcesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTerraformConfigurationSourcesUsingGETOK creates a GetTerraformConfigurationSourcesUsingGETOK with default headers values
func NewGetTerraformConfigurationSourcesUsingGETOK() *GetTerraformConfigurationSourcesUsingGETOK {
	return &GetTerraformConfigurationSourcesUsingGETOK{}
}

/*
GetTerraformConfigurationSourcesUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetTerraformConfigurationSourcesUsingGETOK struct {
	Payload *models.PageOfBlueprintContentSource
}

// IsSuccess returns true when this get terraform configuration sources using g e t o k response has a 2xx status code
func (o *GetTerraformConfigurationSourcesUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get terraform configuration sources using g e t o k response has a 3xx status code
func (o *GetTerraformConfigurationSourcesUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get terraform configuration sources using g e t o k response has a 4xx status code
func (o *GetTerraformConfigurationSourcesUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get terraform configuration sources using g e t o k response has a 5xx status code
func (o *GetTerraformConfigurationSourcesUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get terraform configuration sources using g e t o k response a status code equal to that given
func (o *GetTerraformConfigurationSourcesUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTerraformConfigurationSourcesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-sources][%d] getTerraformConfigurationSourcesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetTerraformConfigurationSourcesUsingGETOK) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-sources][%d] getTerraformConfigurationSourcesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetTerraformConfigurationSourcesUsingGETOK) GetPayload() *models.PageOfBlueprintContentSource {
	return o.Payload
}

func (o *GetTerraformConfigurationSourcesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfBlueprintContentSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTerraformConfigurationSourcesUsingGETUnauthorized creates a GetTerraformConfigurationSourcesUsingGETUnauthorized with default headers values
func NewGetTerraformConfigurationSourcesUsingGETUnauthorized() *GetTerraformConfigurationSourcesUsingGETUnauthorized {
	return &GetTerraformConfigurationSourcesUsingGETUnauthorized{}
}

/*
GetTerraformConfigurationSourcesUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTerraformConfigurationSourcesUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get terraform configuration sources using g e t unauthorized response has a 2xx status code
func (o *GetTerraformConfigurationSourcesUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get terraform configuration sources using g e t unauthorized response has a 3xx status code
func (o *GetTerraformConfigurationSourcesUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get terraform configuration sources using g e t unauthorized response has a 4xx status code
func (o *GetTerraformConfigurationSourcesUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get terraform configuration sources using g e t unauthorized response has a 5xx status code
func (o *GetTerraformConfigurationSourcesUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get terraform configuration sources using g e t unauthorized response a status code equal to that given
func (o *GetTerraformConfigurationSourcesUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTerraformConfigurationSourcesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-sources][%d] getTerraformConfigurationSourcesUsingGETUnauthorized ", 401)
}

func (o *GetTerraformConfigurationSourcesUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-sources][%d] getTerraformConfigurationSourcesUsingGETUnauthorized ", 401)
}

func (o *GetTerraformConfigurationSourcesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTerraformConfigurationSourcesUsingGETForbidden creates a GetTerraformConfigurationSourcesUsingGETForbidden with default headers values
func NewGetTerraformConfigurationSourcesUsingGETForbidden() *GetTerraformConfigurationSourcesUsingGETForbidden {
	return &GetTerraformConfigurationSourcesUsingGETForbidden{}
}

/*
GetTerraformConfigurationSourcesUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTerraformConfigurationSourcesUsingGETForbidden struct {
}

// IsSuccess returns true when this get terraform configuration sources using g e t forbidden response has a 2xx status code
func (o *GetTerraformConfigurationSourcesUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get terraform configuration sources using g e t forbidden response has a 3xx status code
func (o *GetTerraformConfigurationSourcesUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get terraform configuration sources using g e t forbidden response has a 4xx status code
func (o *GetTerraformConfigurationSourcesUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get terraform configuration sources using g e t forbidden response has a 5xx status code
func (o *GetTerraformConfigurationSourcesUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get terraform configuration sources using g e t forbidden response a status code equal to that given
func (o *GetTerraformConfigurationSourcesUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTerraformConfigurationSourcesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-sources][%d] getTerraformConfigurationSourcesUsingGETForbidden ", 403)
}

func (o *GetTerraformConfigurationSourcesUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-sources][%d] getTerraformConfigurationSourcesUsingGETForbidden ", 403)
}

func (o *GetTerraformConfigurationSourcesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
