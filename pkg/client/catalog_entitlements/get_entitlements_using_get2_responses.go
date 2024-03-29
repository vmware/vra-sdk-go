// Code generated by go-swagger; DO NOT EDIT.

package catalog_entitlements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetEntitlementsUsingGET2Reader is a Reader for the GetEntitlementsUsingGET2 structure.
type GetEntitlementsUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEntitlementsUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEntitlementsUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEntitlementsUsingGET2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEntitlementsUsingGET2OK creates a GetEntitlementsUsingGET2OK with default headers values
func NewGetEntitlementsUsingGET2OK() *GetEntitlementsUsingGET2OK {
	return &GetEntitlementsUsingGET2OK{}
}

/*
GetEntitlementsUsingGET2OK describes a response with status code 200, with default header values.

OK
*/
type GetEntitlementsUsingGET2OK struct {
	Payload []*models.Entitlement
}

// IsSuccess returns true when this get entitlements using g e t2 o k response has a 2xx status code
func (o *GetEntitlementsUsingGET2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get entitlements using g e t2 o k response has a 3xx status code
func (o *GetEntitlementsUsingGET2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get entitlements using g e t2 o k response has a 4xx status code
func (o *GetEntitlementsUsingGET2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get entitlements using g e t2 o k response has a 5xx status code
func (o *GetEntitlementsUsingGET2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get entitlements using g e t2 o k response a status code equal to that given
func (o *GetEntitlementsUsingGET2OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetEntitlementsUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /catalog/api/admin/entitlements][%d] getEntitlementsUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetEntitlementsUsingGET2OK) String() string {
	return fmt.Sprintf("[GET /catalog/api/admin/entitlements][%d] getEntitlementsUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetEntitlementsUsingGET2OK) GetPayload() []*models.Entitlement {
	return o.Payload
}

func (o *GetEntitlementsUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEntitlementsUsingGET2Unauthorized creates a GetEntitlementsUsingGET2Unauthorized with default headers values
func NewGetEntitlementsUsingGET2Unauthorized() *GetEntitlementsUsingGET2Unauthorized {
	return &GetEntitlementsUsingGET2Unauthorized{}
}

/*
GetEntitlementsUsingGET2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetEntitlementsUsingGET2Unauthorized struct {
}

// IsSuccess returns true when this get entitlements using g e t2 unauthorized response has a 2xx status code
func (o *GetEntitlementsUsingGET2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get entitlements using g e t2 unauthorized response has a 3xx status code
func (o *GetEntitlementsUsingGET2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get entitlements using g e t2 unauthorized response has a 4xx status code
func (o *GetEntitlementsUsingGET2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get entitlements using g e t2 unauthorized response has a 5xx status code
func (o *GetEntitlementsUsingGET2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get entitlements using g e t2 unauthorized response a status code equal to that given
func (o *GetEntitlementsUsingGET2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetEntitlementsUsingGET2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /catalog/api/admin/entitlements][%d] getEntitlementsUsingGET2Unauthorized ", 401)
}

func (o *GetEntitlementsUsingGET2Unauthorized) String() string {
	return fmt.Sprintf("[GET /catalog/api/admin/entitlements][%d] getEntitlementsUsingGET2Unauthorized ", 401)
}

func (o *GetEntitlementsUsingGET2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
