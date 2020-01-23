// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetUsingGET3Reader is a Reader for the GetUsingGET3 structure.
type GetUsingGET3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsingGET3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsingGET3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsingGET3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsingGET3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsingGET3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsingGET3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsingGET3OK creates a GetUsingGET3OK with default headers values
func NewGetUsingGET3OK() *GetUsingGET3OK {
	return &GetUsingGET3OK{}
}

/*GetUsingGET3OK handles this case with default header values.

'Success' with the requested Pipeline
*/
type GetUsingGET3OK struct {
	Payload models.Pipeline
}

func (o *GetUsingGET3OK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}][%d] getUsingGET3OK  %+v", 200, o.Payload)
}

func (o *GetUsingGET3OK) GetPayload() models.Pipeline {
	return o.Payload
}

func (o *GetUsingGET3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalPipeline(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetUsingGET3Unauthorized creates a GetUsingGET3Unauthorized with default headers values
func NewGetUsingGET3Unauthorized() *GetUsingGET3Unauthorized {
	return &GetUsingGET3Unauthorized{}
}

/*GetUsingGET3Unauthorized handles this case with default header values.

Unauthorized Request
*/
type GetUsingGET3Unauthorized struct {
}

func (o *GetUsingGET3Unauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}][%d] getUsingGET3Unauthorized ", 401)
}

func (o *GetUsingGET3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGET3Forbidden creates a GetUsingGET3Forbidden with default headers values
func NewGetUsingGET3Forbidden() *GetUsingGET3Forbidden {
	return &GetUsingGET3Forbidden{}
}

/*GetUsingGET3Forbidden handles this case with default header values.

Forbidden
*/
type GetUsingGET3Forbidden struct {
}

func (o *GetUsingGET3Forbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}][%d] getUsingGET3Forbidden ", 403)
}

func (o *GetUsingGET3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGET3NotFound creates a GetUsingGET3NotFound with default headers values
func NewGetUsingGET3NotFound() *GetUsingGET3NotFound {
	return &GetUsingGET3NotFound{}
}

/*GetUsingGET3NotFound handles this case with default header values.

Not Found
*/
type GetUsingGET3NotFound struct {
}

func (o *GetUsingGET3NotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}][%d] getUsingGET3NotFound ", 404)
}

func (o *GetUsingGET3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGET3InternalServerError creates a GetUsingGET3InternalServerError with default headers values
func NewGetUsingGET3InternalServerError() *GetUsingGET3InternalServerError {
	return &GetUsingGET3InternalServerError{}
}

/*GetUsingGET3InternalServerError handles this case with default header values.

Server Error
*/
type GetUsingGET3InternalServerError struct {
}

func (o *GetUsingGET3InternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}][%d] getUsingGET3InternalServerError ", 500)
}

func (o *GetUsingGET3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
