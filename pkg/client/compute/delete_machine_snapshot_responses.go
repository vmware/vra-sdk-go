// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteMachineSnapshotReader is a Reader for the DeleteMachineSnapshot structure.
type DeleteMachineSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMachineSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMachineSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteMachineSnapshotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteMachineSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteMachineSnapshotOK creates a DeleteMachineSnapshotOK with default headers values
func NewDeleteMachineSnapshotOK() *DeleteMachineSnapshotOK {
	return &DeleteMachineSnapshotOK{}
}

/*DeleteMachineSnapshotOK handles this case with default header values.

successful operation
*/
type DeleteMachineSnapshotOK struct {
}

func (o *DeleteMachineSnapshotOK) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/machines/{id}/snapshots/{id1}][%d] deleteMachineSnapshotOK ", 200)
}

func (o *DeleteMachineSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMachineSnapshotForbidden creates a DeleteMachineSnapshotForbidden with default headers values
func NewDeleteMachineSnapshotForbidden() *DeleteMachineSnapshotForbidden {
	return &DeleteMachineSnapshotForbidden{}
}

/*DeleteMachineSnapshotForbidden handles this case with default header values.

Forbidden
*/
type DeleteMachineSnapshotForbidden struct {
}

func (o *DeleteMachineSnapshotForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/machines/{id}/snapshots/{id1}][%d] deleteMachineSnapshotForbidden ", 403)
}

func (o *DeleteMachineSnapshotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMachineSnapshotNotFound creates a DeleteMachineSnapshotNotFound with default headers values
func NewDeleteMachineSnapshotNotFound() *DeleteMachineSnapshotNotFound {
	return &DeleteMachineSnapshotNotFound{}
}

/*DeleteMachineSnapshotNotFound handles this case with default header values.

Not Found
*/
type DeleteMachineSnapshotNotFound struct {
}

func (o *DeleteMachineSnapshotNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/machines/{id}/snapshots/{id1}][%d] deleteMachineSnapshotNotFound ", 404)
}

func (o *DeleteMachineSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}