// Code generated by go-swagger; DO NOT EDIT.

package installers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetKubernetesPropertiesUsingGETReader is a Reader for the GetKubernetesPropertiesUsingGET structure.
type GetKubernetesPropertiesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKubernetesPropertiesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKubernetesPropertiesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetKubernetesPropertiesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKubernetesPropertiesUsingGETOK creates a GetKubernetesPropertiesUsingGETOK with default headers values
func NewGetKubernetesPropertiesUsingGETOK() *GetKubernetesPropertiesUsingGETOK {
	return &GetKubernetesPropertiesUsingGETOK{}
}

/*
GetKubernetesPropertiesUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetKubernetesPropertiesUsingGETOK struct {
	Payload map[string]string
}

// IsSuccess returns true when this get kubernetes properties using g e t o k response has a 2xx status code
func (o *GetKubernetesPropertiesUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kubernetes properties using g e t o k response has a 3xx status code
func (o *GetKubernetesPropertiesUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes properties using g e t o k response has a 4xx status code
func (o *GetKubernetesPropertiesUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kubernetes properties using g e t o k response has a 5xx status code
func (o *GetKubernetesPropertiesUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes properties using g e t o k response a status code equal to that given
func (o *GetKubernetesPropertiesUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKubernetesPropertiesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/installers/{id}/properties][%d] getKubernetesPropertiesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesPropertiesUsingGETOK) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/installers/{id}/properties][%d] getKubernetesPropertiesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetKubernetesPropertiesUsingGETOK) GetPayload() map[string]string {
	return o.Payload
}

func (o *GetKubernetesPropertiesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKubernetesPropertiesUsingGETForbidden creates a GetKubernetesPropertiesUsingGETForbidden with default headers values
func NewGetKubernetesPropertiesUsingGETForbidden() *GetKubernetesPropertiesUsingGETForbidden {
	return &GetKubernetesPropertiesUsingGETForbidden{}
}

/*
GetKubernetesPropertiesUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type GetKubernetesPropertiesUsingGETForbidden struct {
}

// IsSuccess returns true when this get kubernetes properties using g e t forbidden response has a 2xx status code
func (o *GetKubernetesPropertiesUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get kubernetes properties using g e t forbidden response has a 3xx status code
func (o *GetKubernetesPropertiesUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kubernetes properties using g e t forbidden response has a 4xx status code
func (o *GetKubernetesPropertiesUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get kubernetes properties using g e t forbidden response has a 5xx status code
func (o *GetKubernetesPropertiesUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get kubernetes properties using g e t forbidden response a status code equal to that given
func (o *GetKubernetesPropertiesUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetKubernetesPropertiesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/installers/{id}/properties][%d] getKubernetesPropertiesUsingGETForbidden ", 403)
}

func (o *GetKubernetesPropertiesUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/installers/{id}/properties][%d] getKubernetesPropertiesUsingGETForbidden ", 403)
}

func (o *GetKubernetesPropertiesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}