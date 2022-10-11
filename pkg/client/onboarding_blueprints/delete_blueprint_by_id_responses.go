// Code generated by go-swagger; DO NOT EDIT.

package onboarding_blueprints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteBlueprintByIDReader is a Reader for the DeleteBlueprintByID structure.
type DeleteBlueprintByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBlueprintByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteBlueprintByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteBlueprintByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteBlueprintByIDOK creates a DeleteBlueprintByIDOK with default headers values
func NewDeleteBlueprintByIDOK() *DeleteBlueprintByIDOK {
	return &DeleteBlueprintByIDOK{}
}

/*
DeleteBlueprintByIDOK describes a response with status code 200, with default header values.

Success
*/
type DeleteBlueprintByIDOK struct {
}

// IsSuccess returns true when this delete blueprint by Id o k response has a 2xx status code
func (o *DeleteBlueprintByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete blueprint by Id o k response has a 3xx status code
func (o *DeleteBlueprintByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete blueprint by Id o k response has a 4xx status code
func (o *DeleteBlueprintByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete blueprint by Id o k response has a 5xx status code
func (o *DeleteBlueprintByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete blueprint by Id o k response a status code equal to that given
func (o *DeleteBlueprintByIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteBlueprintByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /relocation/onboarding/blueprint/{id}][%d] deleteBlueprintByIdOK ", 200)
}

func (o *DeleteBlueprintByIDOK) String() string {
	return fmt.Sprintf("[DELETE /relocation/onboarding/blueprint/{id}][%d] deleteBlueprintByIdOK ", 200)
}

func (o *DeleteBlueprintByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBlueprintByIDUnauthorized creates a DeleteBlueprintByIDUnauthorized with default headers values
func NewDeleteBlueprintByIDUnauthorized() *DeleteBlueprintByIDUnauthorized {
	return &DeleteBlueprintByIDUnauthorized{}
}

/*
DeleteBlueprintByIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteBlueprintByIDUnauthorized struct {
}

// IsSuccess returns true when this delete blueprint by Id unauthorized response has a 2xx status code
func (o *DeleteBlueprintByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete blueprint by Id unauthorized response has a 3xx status code
func (o *DeleteBlueprintByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete blueprint by Id unauthorized response has a 4xx status code
func (o *DeleteBlueprintByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete blueprint by Id unauthorized response has a 5xx status code
func (o *DeleteBlueprintByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete blueprint by Id unauthorized response a status code equal to that given
func (o *DeleteBlueprintByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteBlueprintByIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /relocation/onboarding/blueprint/{id}][%d] deleteBlueprintByIdUnauthorized ", 401)
}

func (o *DeleteBlueprintByIDUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /relocation/onboarding/blueprint/{id}][%d] deleteBlueprintByIdUnauthorized ", 401)
}

func (o *DeleteBlueprintByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
