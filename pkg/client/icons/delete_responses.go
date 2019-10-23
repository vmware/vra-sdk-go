// Code generated by go-swagger; DO NOT EDIT.

package icons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteReader is a Reader for the Delete structure.
type DeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteNoContent creates a DeleteNoContent with default headers values
func NewDeleteNoContent() *DeleteNoContent {
	return &DeleteNoContent{}
}

/*DeleteNoContent handles this case with default header values.

Success - delete the icon
*/
type DeleteNoContent struct {
}

func (o *DeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /icon/api/icons/{id}][%d] deleteNoContent ", 204)
}

func (o *DeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteForbidden creates a DeleteForbidden with default headers values
func NewDeleteForbidden() *DeleteForbidden {
	return &DeleteForbidden{}
}

/*DeleteForbidden handles this case with default header values.

Forbidden.
*/
type DeleteForbidden struct {
}

func (o *DeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /icon/api/icons/{id}][%d] deleteForbidden ", 403)
}

func (o *DeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
