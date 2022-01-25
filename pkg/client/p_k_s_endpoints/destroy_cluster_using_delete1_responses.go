// Code generated by go-swagger; DO NOT EDIT.

package p_k_s_endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DestroyClusterUsingDELETE1Reader is a Reader for the DestroyClusterUsingDELETE1 structure.
type DestroyClusterUsingDELETE1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DestroyClusterUsingDELETE1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDestroyClusterUsingDELETE1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDestroyClusterUsingDELETE1OK creates a DestroyClusterUsingDELETE1OK with default headers values
func NewDestroyClusterUsingDELETE1OK() *DestroyClusterUsingDELETE1OK {
	return &DestroyClusterUsingDELETE1OK{}
}

/* DestroyClusterUsingDELETE1OK describes a response with status code 200, with default header values.

OK
*/
type DestroyClusterUsingDELETE1OK struct {
}

func (o *DestroyClusterUsingDELETE1OK) Error() string {
	return fmt.Sprintf("[DELETE /cmx/api/resources/pks/endpoints/{id}/clusters/{clusterId}][%d] destroyClusterUsingDELETE1OK ", 200)
}

func (o *DestroyClusterUsingDELETE1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
