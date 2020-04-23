// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewEnumeratePrivateImagesVSphereParams creates a new EnumeratePrivateImagesVSphereParams object
// with the default values initialized.
func NewEnumeratePrivateImagesVSphereParams() *EnumeratePrivateImagesVSphereParams {
	var ()
	return &EnumeratePrivateImagesVSphereParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnumeratePrivateImagesVSphereParamsWithTimeout creates a new EnumeratePrivateImagesVSphereParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnumeratePrivateImagesVSphereParamsWithTimeout(timeout time.Duration) *EnumeratePrivateImagesVSphereParams {
	var ()
	return &EnumeratePrivateImagesVSphereParams{

		timeout: timeout,
	}
}

// NewEnumeratePrivateImagesVSphereParamsWithContext creates a new EnumeratePrivateImagesVSphereParams object
// with the default values initialized, and the ability to set a context for a request
func NewEnumeratePrivateImagesVSphereParamsWithContext(ctx context.Context) *EnumeratePrivateImagesVSphereParams {
	var ()
	return &EnumeratePrivateImagesVSphereParams{

		Context: ctx,
	}
}

// NewEnumeratePrivateImagesVSphereParamsWithHTTPClient creates a new EnumeratePrivateImagesVSphereParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnumeratePrivateImagesVSphereParamsWithHTTPClient(client *http.Client) *EnumeratePrivateImagesVSphereParams {
	var ()
	return &EnumeratePrivateImagesVSphereParams{
		HTTPClient: client,
	}
}

/*EnumeratePrivateImagesVSphereParams contains all the parameters to send to the API endpoint
for the enumerate private images v sphere operation typically these are written to a http.Request
*/
type EnumeratePrivateImagesVSphereParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  Id of vSphere cloud account to enumerate

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enumerate private images v sphere params
func (o *EnumeratePrivateImagesVSphereParams) WithTimeout(timeout time.Duration) *EnumeratePrivateImagesVSphereParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enumerate private images v sphere params
func (o *EnumeratePrivateImagesVSphereParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enumerate private images v sphere params
func (o *EnumeratePrivateImagesVSphereParams) WithContext(ctx context.Context) *EnumeratePrivateImagesVSphereParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enumerate private images v sphere params
func (o *EnumeratePrivateImagesVSphereParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enumerate private images v sphere params
func (o *EnumeratePrivateImagesVSphereParams) WithHTTPClient(client *http.Client) *EnumeratePrivateImagesVSphereParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enumerate private images v sphere params
func (o *EnumeratePrivateImagesVSphereParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the enumerate private images v sphere params
func (o *EnumeratePrivateImagesVSphereParams) WithAPIVersion(aPIVersion *string) *EnumeratePrivateImagesVSphereParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the enumerate private images v sphere params
func (o *EnumeratePrivateImagesVSphereParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the enumerate private images v sphere params
func (o *EnumeratePrivateImagesVSphereParams) WithID(id string) *EnumeratePrivateImagesVSphereParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the enumerate private images v sphere params
func (o *EnumeratePrivateImagesVSphereParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EnumeratePrivateImagesVSphereParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// query param apiVersion
		var qrAPIVersion string
		if o.APIVersion != nil {
			qrAPIVersion = *o.APIVersion
		}
		qAPIVersion := qrAPIVersion
		if qAPIVersion != "" {
			if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}