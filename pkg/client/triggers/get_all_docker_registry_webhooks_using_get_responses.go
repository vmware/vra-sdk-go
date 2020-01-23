// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetAllDockerRegistryWebhooksUsingGETReader is a Reader for the GetAllDockerRegistryWebhooksUsingGET structure.
type GetAllDockerRegistryWebhooksUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllDockerRegistryWebhooksUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllDockerRegistryWebhooksUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllDockerRegistryWebhooksUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllDockerRegistryWebhooksUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllDockerRegistryWebhooksUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllDockerRegistryWebhooksUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllDockerRegistryWebhooksUsingGETOK creates a GetAllDockerRegistryWebhooksUsingGETOK with default headers values
func NewGetAllDockerRegistryWebhooksUsingGETOK() *GetAllDockerRegistryWebhooksUsingGETOK {
	return &GetAllDockerRegistryWebhooksUsingGETOK{}
}

/*GetAllDockerRegistryWebhooksUsingGETOK handles this case with default header values.

'Success' with get of Docker Registry Webhooks
*/
type GetAllDockerRegistryWebhooksUsingGETOK struct {
	Payload models.DockerRegistryWebhooks
}

func (o *GetAllDockerRegistryWebhooksUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks][%d] getAllDockerRegistryWebhooksUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllDockerRegistryWebhooksUsingGETOK) GetPayload() models.DockerRegistryWebhooks {
	return o.Payload
}

func (o *GetAllDockerRegistryWebhooksUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalDockerRegistryWebhooks(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetAllDockerRegistryWebhooksUsingGETUnauthorized creates a GetAllDockerRegistryWebhooksUsingGETUnauthorized with default headers values
func NewGetAllDockerRegistryWebhooksUsingGETUnauthorized() *GetAllDockerRegistryWebhooksUsingGETUnauthorized {
	return &GetAllDockerRegistryWebhooksUsingGETUnauthorized{}
}

/*GetAllDockerRegistryWebhooksUsingGETUnauthorized handles this case with default header values.

Unauthorized Request
*/
type GetAllDockerRegistryWebhooksUsingGETUnauthorized struct {
}

func (o *GetAllDockerRegistryWebhooksUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks][%d] getAllDockerRegistryWebhooksUsingGETUnauthorized ", 401)
}

func (o *GetAllDockerRegistryWebhooksUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllDockerRegistryWebhooksUsingGETForbidden creates a GetAllDockerRegistryWebhooksUsingGETForbidden with default headers values
func NewGetAllDockerRegistryWebhooksUsingGETForbidden() *GetAllDockerRegistryWebhooksUsingGETForbidden {
	return &GetAllDockerRegistryWebhooksUsingGETForbidden{}
}

/*GetAllDockerRegistryWebhooksUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAllDockerRegistryWebhooksUsingGETForbidden struct {
}

func (o *GetAllDockerRegistryWebhooksUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks][%d] getAllDockerRegistryWebhooksUsingGETForbidden ", 403)
}

func (o *GetAllDockerRegistryWebhooksUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllDockerRegistryWebhooksUsingGETNotFound creates a GetAllDockerRegistryWebhooksUsingGETNotFound with default headers values
func NewGetAllDockerRegistryWebhooksUsingGETNotFound() *GetAllDockerRegistryWebhooksUsingGETNotFound {
	return &GetAllDockerRegistryWebhooksUsingGETNotFound{}
}

/*GetAllDockerRegistryWebhooksUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetAllDockerRegistryWebhooksUsingGETNotFound struct {
}

func (o *GetAllDockerRegistryWebhooksUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks][%d] getAllDockerRegistryWebhooksUsingGETNotFound ", 404)
}

func (o *GetAllDockerRegistryWebhooksUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllDockerRegistryWebhooksUsingGETInternalServerError creates a GetAllDockerRegistryWebhooksUsingGETInternalServerError with default headers values
func NewGetAllDockerRegistryWebhooksUsingGETInternalServerError() *GetAllDockerRegistryWebhooksUsingGETInternalServerError {
	return &GetAllDockerRegistryWebhooksUsingGETInternalServerError{}
}

/*GetAllDockerRegistryWebhooksUsingGETInternalServerError handles this case with default header values.

Server Error
*/
type GetAllDockerRegistryWebhooksUsingGETInternalServerError struct {
}

func (o *GetAllDockerRegistryWebhooksUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-webhooks][%d] getAllDockerRegistryWebhooksUsingGETInternalServerError ", 500)
}

func (o *GetAllDockerRegistryWebhooksUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
