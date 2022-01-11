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

// GetTerraformConfigurationSourceCommitListUsingGETReader is a Reader for the GetTerraformConfigurationSourceCommitListUsingGET structure.
type GetTerraformConfigurationSourceCommitListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTerraformConfigurationSourceCommitListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTerraformConfigurationSourceCommitListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTerraformConfigurationSourceCommitListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTerraformConfigurationSourceCommitListUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTerraformConfigurationSourceCommitListUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTerraformConfigurationSourceCommitListUsingGETOK creates a GetTerraformConfigurationSourceCommitListUsingGETOK with default headers values
func NewGetTerraformConfigurationSourceCommitListUsingGETOK() *GetTerraformConfigurationSourceCommitListUsingGETOK {
	return &GetTerraformConfigurationSourceCommitListUsingGETOK{}
}

/* GetTerraformConfigurationSourceCommitListUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetTerraformConfigurationSourceCommitListUsingGETOK struct {
	Payload *models.PageOfCommitDetails
}

func (o *GetTerraformConfigurationSourceCommitListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-source-commits][%d] getTerraformConfigurationSourceCommitListUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetTerraformConfigurationSourceCommitListUsingGETOK) GetPayload() *models.PageOfCommitDetails {
	return o.Payload
}

func (o *GetTerraformConfigurationSourceCommitListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfCommitDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTerraformConfigurationSourceCommitListUsingGETBadRequest creates a GetTerraformConfigurationSourceCommitListUsingGETBadRequest with default headers values
func NewGetTerraformConfigurationSourceCommitListUsingGETBadRequest() *GetTerraformConfigurationSourceCommitListUsingGETBadRequest {
	return &GetTerraformConfigurationSourceCommitListUsingGETBadRequest{}
}

/* GetTerraformConfigurationSourceCommitListUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetTerraformConfigurationSourceCommitListUsingGETBadRequest struct {
	Payload *models.Error
}

func (o *GetTerraformConfigurationSourceCommitListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-source-commits][%d] getTerraformConfigurationSourceCommitListUsingGETBadRequest  %+v", 400, o.Payload)
}
func (o *GetTerraformConfigurationSourceCommitListUsingGETBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTerraformConfigurationSourceCommitListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTerraformConfigurationSourceCommitListUsingGETUnauthorized creates a GetTerraformConfigurationSourceCommitListUsingGETUnauthorized with default headers values
func NewGetTerraformConfigurationSourceCommitListUsingGETUnauthorized() *GetTerraformConfigurationSourceCommitListUsingGETUnauthorized {
	return &GetTerraformConfigurationSourceCommitListUsingGETUnauthorized{}
}

/* GetTerraformConfigurationSourceCommitListUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTerraformConfigurationSourceCommitListUsingGETUnauthorized struct {
}

func (o *GetTerraformConfigurationSourceCommitListUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-source-commits][%d] getTerraformConfigurationSourceCommitListUsingGETUnauthorized ", 401)
}

func (o *GetTerraformConfigurationSourceCommitListUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTerraformConfigurationSourceCommitListUsingGETForbidden creates a GetTerraformConfigurationSourceCommitListUsingGETForbidden with default headers values
func NewGetTerraformConfigurationSourceCommitListUsingGETForbidden() *GetTerraformConfigurationSourceCommitListUsingGETForbidden {
	return &GetTerraformConfigurationSourceCommitListUsingGETForbidden{}
}

/* GetTerraformConfigurationSourceCommitListUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTerraformConfigurationSourceCommitListUsingGETForbidden struct {
}

func (o *GetTerraformConfigurationSourceCommitListUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-source-commits][%d] getTerraformConfigurationSourceCommitListUsingGETForbidden ", 403)
}

func (o *GetTerraformConfigurationSourceCommitListUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
