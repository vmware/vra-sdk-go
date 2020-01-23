// Code generated by go-swagger; DO NOT EDIT.

package custom_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDraftByIDUsingGETReader is a Reader for the GetDraftByIDUsingGET structure.
type GetDraftByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDraftByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDraftByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDraftByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDraftByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDraftByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDraftByIDUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDraftByIDUsingGETOK creates a GetDraftByIDUsingGETOK with default headers values
func NewGetDraftByIDUsingGETOK() *GetDraftByIDUsingGETOK {
	return &GetDraftByIDUsingGETOK{}
}

/*GetDraftByIDUsingGETOK handles this case with default header values.

OK
*/
type GetDraftByIDUsingGETOK struct {
	Payload models.CustomIntegration
}

func (o *GetDraftByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetDraftByIDUsingGETOK) GetPayload() models.CustomIntegration {
	return o.Payload
}

func (o *GetDraftByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalCustomIntegration(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetDraftByIDUsingGETUnauthorized creates a GetDraftByIDUsingGETUnauthorized with default headers values
func NewGetDraftByIDUsingGETUnauthorized() *GetDraftByIDUsingGETUnauthorized {
	return &GetDraftByIDUsingGETUnauthorized{}
}

/*GetDraftByIDUsingGETUnauthorized handles this case with default header values.

Unauthorized Request
*/
type GetDraftByIDUsingGETUnauthorized struct {
}

func (o *GetDraftByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETUnauthorized ", 401)
}

func (o *GetDraftByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDraftByIDUsingGETForbidden creates a GetDraftByIDUsingGETForbidden with default headers values
func NewGetDraftByIDUsingGETForbidden() *GetDraftByIDUsingGETForbidden {
	return &GetDraftByIDUsingGETForbidden{}
}

/*GetDraftByIDUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetDraftByIDUsingGETForbidden struct {
}

func (o *GetDraftByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETForbidden ", 403)
}

func (o *GetDraftByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDraftByIDUsingGETNotFound creates a GetDraftByIDUsingGETNotFound with default headers values
func NewGetDraftByIDUsingGETNotFound() *GetDraftByIDUsingGETNotFound {
	return &GetDraftByIDUsingGETNotFound{}
}

/*GetDraftByIDUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetDraftByIDUsingGETNotFound struct {
}

func (o *GetDraftByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETNotFound ", 404)
}

func (o *GetDraftByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDraftByIDUsingGETInternalServerError creates a GetDraftByIDUsingGETInternalServerError with default headers values
func NewGetDraftByIDUsingGETInternalServerError() *GetDraftByIDUsingGETInternalServerError {
	return &GetDraftByIDUsingGETInternalServerError{}
}

/*GetDraftByIDUsingGETInternalServerError handles this case with default header values.

Server Error
*/
type GetDraftByIDUsingGETInternalServerError struct {
}

func (o *GetDraftByIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/custom-integrations/{id}][%d] getDraftByIdUsingGETInternalServerError ", 500)
}

func (o *GetDraftByIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
