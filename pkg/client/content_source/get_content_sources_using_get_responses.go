// Code generated by go-swagger; DO NOT EDIT.

package content_source

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetContentSourcesUsingGETReader is a Reader for the GetContentSourcesUsingGET structure.
type GetContentSourcesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentSourcesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentSourcesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetContentSourcesUsingGETOK creates a GetContentSourcesUsingGETOK with default headers values
func NewGetContentSourcesUsingGETOK() *GetContentSourcesUsingGETOK {
	return &GetContentSourcesUsingGETOK{}
}

/*GetContentSourcesUsingGETOK handles this case with default header values.

Content sources
*/
type GetContentSourcesUsingGETOK struct {
	Payload *models.ContentSources
}

func (o *GetContentSourcesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /content/api/sources][%d] getContentSourcesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetContentSourcesUsingGETOK) GetPayload() *models.ContentSources {
	return o.Payload
}

func (o *GetContentSourcesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentSources)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}