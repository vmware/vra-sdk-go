// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetPipelineTilesUsingGETReader is a Reader for the GetPipelineTilesUsingGET structure.
type GetPipelineTilesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelineTilesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelineTilesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPipelineTilesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPipelineTilesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPipelineTilesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPipelineTilesUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPipelineTilesUsingGETOK creates a GetPipelineTilesUsingGETOK with default headers values
func NewGetPipelineTilesUsingGETOK() *GetPipelineTilesUsingGETOK {
	return &GetPipelineTilesUsingGETOK{}
}

/*
GetPipelineTilesUsingGETOK describes a response with status code 200, with default header values.

'Success' with the requested Pipeline Tiles
*/
type GetPipelineTilesUsingGETOK struct {
	Payload models.Tiles
}

// IsSuccess returns true when this get pipeline tiles using g e t o k response has a 2xx status code
func (o *GetPipelineTilesUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get pipeline tiles using g e t o k response has a 3xx status code
func (o *GetPipelineTilesUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline tiles using g e t o k response has a 4xx status code
func (o *GetPipelineTilesUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pipeline tiles using g e t o k response has a 5xx status code
func (o *GetPipelineTilesUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline tiles using g e t o k response a status code equal to that given
func (o *GetPipelineTilesUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPipelineTilesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipeline-tiles][%d] getPipelineTilesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPipelineTilesUsingGETOK) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipeline-tiles][%d] getPipelineTilesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPipelineTilesUsingGETOK) GetPayload() models.Tiles {
	return o.Payload
}

func (o *GetPipelineTilesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalTiles(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetPipelineTilesUsingGETUnauthorized creates a GetPipelineTilesUsingGETUnauthorized with default headers values
func NewGetPipelineTilesUsingGETUnauthorized() *GetPipelineTilesUsingGETUnauthorized {
	return &GetPipelineTilesUsingGETUnauthorized{}
}

/*
GetPipelineTilesUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetPipelineTilesUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get pipeline tiles using g e t unauthorized response has a 2xx status code
func (o *GetPipelineTilesUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipeline tiles using g e t unauthorized response has a 3xx status code
func (o *GetPipelineTilesUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline tiles using g e t unauthorized response has a 4xx status code
func (o *GetPipelineTilesUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pipeline tiles using g e t unauthorized response has a 5xx status code
func (o *GetPipelineTilesUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline tiles using g e t unauthorized response a status code equal to that given
func (o *GetPipelineTilesUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPipelineTilesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipeline-tiles][%d] getPipelineTilesUsingGETUnauthorized ", 401)
}

func (o *GetPipelineTilesUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipeline-tiles][%d] getPipelineTilesUsingGETUnauthorized ", 401)
}

func (o *GetPipelineTilesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPipelineTilesUsingGETForbidden creates a GetPipelineTilesUsingGETForbidden with default headers values
func NewGetPipelineTilesUsingGETForbidden() *GetPipelineTilesUsingGETForbidden {
	return &GetPipelineTilesUsingGETForbidden{}
}

/*
GetPipelineTilesUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPipelineTilesUsingGETForbidden struct {
}

// IsSuccess returns true when this get pipeline tiles using g e t forbidden response has a 2xx status code
func (o *GetPipelineTilesUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipeline tiles using g e t forbidden response has a 3xx status code
func (o *GetPipelineTilesUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline tiles using g e t forbidden response has a 4xx status code
func (o *GetPipelineTilesUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pipeline tiles using g e t forbidden response has a 5xx status code
func (o *GetPipelineTilesUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline tiles using g e t forbidden response a status code equal to that given
func (o *GetPipelineTilesUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPipelineTilesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipeline-tiles][%d] getPipelineTilesUsingGETForbidden ", 403)
}

func (o *GetPipelineTilesUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipeline-tiles][%d] getPipelineTilesUsingGETForbidden ", 403)
}

func (o *GetPipelineTilesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPipelineTilesUsingGETNotFound creates a GetPipelineTilesUsingGETNotFound with default headers values
func NewGetPipelineTilesUsingGETNotFound() *GetPipelineTilesUsingGETNotFound {
	return &GetPipelineTilesUsingGETNotFound{}
}

/*
GetPipelineTilesUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetPipelineTilesUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get pipeline tiles using g e t not found response has a 2xx status code
func (o *GetPipelineTilesUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipeline tiles using g e t not found response has a 3xx status code
func (o *GetPipelineTilesUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline tiles using g e t not found response has a 4xx status code
func (o *GetPipelineTilesUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pipeline tiles using g e t not found response has a 5xx status code
func (o *GetPipelineTilesUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline tiles using g e t not found response a status code equal to that given
func (o *GetPipelineTilesUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPipelineTilesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipeline-tiles][%d] getPipelineTilesUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetPipelineTilesUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipeline-tiles][%d] getPipelineTilesUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetPipelineTilesUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPipelineTilesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelineTilesUsingGETInternalServerError creates a GetPipelineTilesUsingGETInternalServerError with default headers values
func NewGetPipelineTilesUsingGETInternalServerError() *GetPipelineTilesUsingGETInternalServerError {
	return &GetPipelineTilesUsingGETInternalServerError{}
}

/*
GetPipelineTilesUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetPipelineTilesUsingGETInternalServerError struct {
}

// IsSuccess returns true when this get pipeline tiles using g e t internal server error response has a 2xx status code
func (o *GetPipelineTilesUsingGETInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipeline tiles using g e t internal server error response has a 3xx status code
func (o *GetPipelineTilesUsingGETInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline tiles using g e t internal server error response has a 4xx status code
func (o *GetPipelineTilesUsingGETInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pipeline tiles using g e t internal server error response has a 5xx status code
func (o *GetPipelineTilesUsingGETInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get pipeline tiles using g e t internal server error response a status code equal to that given
func (o *GetPipelineTilesUsingGETInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPipelineTilesUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipeline-tiles][%d] getPipelineTilesUsingGETInternalServerError ", 500)
}

func (o *GetPipelineTilesUsingGETInternalServerError) String() string {
	return fmt.Sprintf("[GET /codestream/api/pipeline-tiles][%d] getPipelineTilesUsingGETInternalServerError ", 500)
}

func (o *GetPipelineTilesUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
