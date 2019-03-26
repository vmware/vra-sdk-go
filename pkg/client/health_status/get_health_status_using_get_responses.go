// Code generated by go-swagger; DO NOT EDIT.

package health_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetHealthStatusUsingGETReader is a Reader for the GetHealthStatusUsingGET structure.
type GetHealthStatusUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHealthStatusUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetHealthStatusUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetHealthStatusUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetHealthStatusUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetHealthStatusUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetHealthStatusUsingGETOK creates a GetHealthStatusUsingGETOK with default headers values
func NewGetHealthStatusUsingGETOK() *GetHealthStatusUsingGETOK {
	return &GetHealthStatusUsingGETOK{}
}

/*GetHealthStatusUsingGETOK handles this case with default header values.

OK
*/
type GetHealthStatusUsingGETOK struct {
}

func (o *GetHealthStatusUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /catalog/api/health][%d] getHealthStatusUsingGETOK ", 200)
}

func (o *GetHealthStatusUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHealthStatusUsingGETUnauthorized creates a GetHealthStatusUsingGETUnauthorized with default headers values
func NewGetHealthStatusUsingGETUnauthorized() *GetHealthStatusUsingGETUnauthorized {
	return &GetHealthStatusUsingGETUnauthorized{}
}

/*GetHealthStatusUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetHealthStatusUsingGETUnauthorized struct {
}

func (o *GetHealthStatusUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /catalog/api/health][%d] getHealthStatusUsingGETUnauthorized ", 401)
}

func (o *GetHealthStatusUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHealthStatusUsingGETForbidden creates a GetHealthStatusUsingGETForbidden with default headers values
func NewGetHealthStatusUsingGETForbidden() *GetHealthStatusUsingGETForbidden {
	return &GetHealthStatusUsingGETForbidden{}
}

/*GetHealthStatusUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetHealthStatusUsingGETForbidden struct {
}

func (o *GetHealthStatusUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /catalog/api/health][%d] getHealthStatusUsingGETForbidden ", 403)
}

func (o *GetHealthStatusUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHealthStatusUsingGETNotFound creates a GetHealthStatusUsingGETNotFound with default headers values
func NewGetHealthStatusUsingGETNotFound() *GetHealthStatusUsingGETNotFound {
	return &GetHealthStatusUsingGETNotFound{}
}

/*GetHealthStatusUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetHealthStatusUsingGETNotFound struct {
}

func (o *GetHealthStatusUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /catalog/api/health][%d] getHealthStatusUsingGETNotFound ", 404)
}

func (o *GetHealthStatusUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}