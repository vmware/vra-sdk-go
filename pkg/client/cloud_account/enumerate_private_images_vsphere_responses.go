// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// EnumeratePrivateImagesVSphereReader is a Reader for the EnumeratePrivateImagesVSphere structure.
type EnumeratePrivateImagesVSphereReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnumeratePrivateImagesVSphereReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewEnumeratePrivateImagesVSphereDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewEnumeratePrivateImagesVSphereDefault creates a EnumeratePrivateImagesVSphereDefault with default headers values
func NewEnumeratePrivateImagesVSphereDefault(code int) *EnumeratePrivateImagesVSphereDefault {
	return &EnumeratePrivateImagesVSphereDefault{
		_statusCode: code,
	}
}

/*EnumeratePrivateImagesVSphereDefault handles this case with default header values.

successful operation
*/
type EnumeratePrivateImagesVSphereDefault struct {
	_statusCode int
}

// Code gets the status code for the enumerate private images v sphere default response
func (o *EnumeratePrivateImagesVSphereDefault) Code() int {
	return o._statusCode
}

func (o *EnumeratePrivateImagesVSphereDefault) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vsphere/{id}/private-image-enumeration][%d] enumeratePrivateImagesVSphere default ", o._statusCode)
}

func (o *EnumeratePrivateImagesVSphereDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
