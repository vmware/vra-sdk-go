// Code generated by go-swagger; DO NOT EDIT.

package disk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDiskSnapshotsReader is a Reader for the GetDiskSnapshots structure.
type GetDiskSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDiskSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDiskSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetDiskSnapshotsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDiskSnapshotsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDiskSnapshotsOK creates a GetDiskSnapshotsOK with default headers values
func NewGetDiskSnapshotsOK() *GetDiskSnapshotsOK {
	return &GetDiskSnapshotsOK{}
}

/*GetDiskSnapshotsOK handles this case with default header values.

successful operation
*/
type GetDiskSnapshotsOK struct {
	Payload []*models.DiskSnapshot
}

func (o *GetDiskSnapshotsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/block-devices/{id}/snapshots][%d] getDiskSnapshotsOK  %+v", 200, o.Payload)
}

func (o *GetDiskSnapshotsOK) GetPayload() []*models.DiskSnapshot {
	return o.Payload
}

func (o *GetDiskSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiskSnapshotsForbidden creates a GetDiskSnapshotsForbidden with default headers values
func NewGetDiskSnapshotsForbidden() *GetDiskSnapshotsForbidden {
	return &GetDiskSnapshotsForbidden{}
}

/*GetDiskSnapshotsForbidden handles this case with default header values.

Forbidden
*/
type GetDiskSnapshotsForbidden struct {
}

func (o *GetDiskSnapshotsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/block-devices/{id}/snapshots][%d] getDiskSnapshotsForbidden ", 403)
}

func (o *GetDiskSnapshotsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDiskSnapshotsNotFound creates a GetDiskSnapshotsNotFound with default headers values
func NewGetDiskSnapshotsNotFound() *GetDiskSnapshotsNotFound {
	return &GetDiskSnapshotsNotFound{}
}

/*GetDiskSnapshotsNotFound handles this case with default header values.

Not Found
*/
type GetDiskSnapshotsNotFound struct {
	Payload *models.Error
}

func (o *GetDiskSnapshotsNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/block-devices/{id}/snapshots][%d] getDiskSnapshotsNotFound  %+v", 404, o.Payload)
}

func (o *GetDiskSnapshotsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDiskSnapshotsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
