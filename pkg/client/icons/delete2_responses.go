// Code generated by go-swagger; DO NOT EDIT.

package icons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Delete2Reader is a Reader for the Delete2 structure.
type Delete2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Delete2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDelete2NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDelete2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDelete2NoContent creates a Delete2NoContent with default headers values
func NewDelete2NoContent() *Delete2NoContent {
	return &Delete2NoContent{}
}

/* Delete2NoContent describes a response with status code 204, with default header values.

Success - delete the icon
*/
type Delete2NoContent struct {
}

func (o *Delete2NoContent) Error() string {
	return fmt.Sprintf("[DELETE /icon/api/icons/{id}][%d] delete2NoContent ", 204)
}

func (o *Delete2NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDelete2Forbidden creates a Delete2Forbidden with default headers values
func NewDelete2Forbidden() *Delete2Forbidden {
	return &Delete2Forbidden{}
}

/* Delete2Forbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type Delete2Forbidden struct {
}

func (o *Delete2Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /icon/api/icons/{id}][%d] delete2Forbidden ", 403)
}

func (o *Delete2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
