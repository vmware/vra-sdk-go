// Code generated by go-swagger; DO NOT EDIT.

package network_ip_range

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteNetworkIPRangeReader is a Reader for the DeleteNetworkIPRange structure.
type DeleteNetworkIPRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkIPRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkIPRangeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteNetworkIPRangeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteNetworkIPRangeNoContent creates a DeleteNetworkIPRangeNoContent with default headers values
func NewDeleteNetworkIPRangeNoContent() *DeleteNetworkIPRangeNoContent {
	return &DeleteNetworkIPRangeNoContent{}
}

/*DeleteNetworkIPRangeNoContent handles this case with default header values.

No Content
*/
type DeleteNetworkIPRangeNoContent struct {
}

func (o *DeleteNetworkIPRangeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/network-ip-ranges/{id}][%d] deleteNetworkIpRangeNoContent ", 204)
}

func (o *DeleteNetworkIPRangeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworkIPRangeForbidden creates a DeleteNetworkIPRangeForbidden with default headers values
func NewDeleteNetworkIPRangeForbidden() *DeleteNetworkIPRangeForbidden {
	return &DeleteNetworkIPRangeForbidden{}
}

/*DeleteNetworkIPRangeForbidden handles this case with default header values.

Forbidden
*/
type DeleteNetworkIPRangeForbidden struct {
}

func (o *DeleteNetworkIPRangeForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/network-ip-ranges/{id}][%d] deleteNetworkIpRangeForbidden ", 403)
}

func (o *DeleteNetworkIPRangeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
