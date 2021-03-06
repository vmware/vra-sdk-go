// Code generated by go-swagger; DO NOT EDIT.

package content_source

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// ListContentSourcesUsingGETReader is a Reader for the ListContentSourcesUsingGET structure.
type ListContentSourcesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListContentSourcesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListContentSourcesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListContentSourcesUsingGETOK creates a ListContentSourcesUsingGETOK with default headers values
func NewListContentSourcesUsingGETOK() *ListContentSourcesUsingGETOK {
	return &ListContentSourcesUsingGETOK{}
}

/*ListContentSourcesUsingGETOK handles this case with default header values.

Content sources
*/
type ListContentSourcesUsingGETOK struct {
	Payload *models.ContentSources
}

func (o *ListContentSourcesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /content/api/sources][%d] listContentSourcesUsingGETOK  %+v", 200, o.Payload)
}

func (o *ListContentSourcesUsingGETOK) GetPayload() *models.ContentSources {
	return o.Payload
}

func (o *ListContentSourcesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentSources)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
