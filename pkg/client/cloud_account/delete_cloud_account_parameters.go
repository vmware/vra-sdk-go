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

// NewDeleteCloudAccountParams creates a new DeleteCloudAccountParams object
// with the default values initialized.
func NewDeleteCloudAccountParams() *DeleteCloudAccountParams {
	var ()
	return &DeleteCloudAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCloudAccountParamsWithTimeout creates a new DeleteCloudAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCloudAccountParamsWithTimeout(timeout time.Duration) *DeleteCloudAccountParams {
	var ()
	return &DeleteCloudAccountParams{

		timeout: timeout,
	}
}

// NewDeleteCloudAccountParamsWithContext creates a new DeleteCloudAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCloudAccountParamsWithContext(ctx context.Context) *DeleteCloudAccountParams {
	var ()
	return &DeleteCloudAccountParams{

		Context: ctx,
	}
}

// NewDeleteCloudAccountParamsWithHTTPClient creates a new DeleteCloudAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCloudAccountParamsWithHTTPClient(client *http.Client) *DeleteCloudAccountParams {
	var ()
	return &DeleteCloudAccountParams{
		HTTPClient: client,
	}
}

/*DeleteCloudAccountParams contains all the parameters to send to the API endpoint
for the delete cloud account operation typically these are written to a http.Request
*/
type DeleteCloudAccountParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the Cloud Account

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cloud account params
func (o *DeleteCloudAccountParams) WithTimeout(timeout time.Duration) *DeleteCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cloud account params
func (o *DeleteCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cloud account params
func (o *DeleteCloudAccountParams) WithContext(ctx context.Context) *DeleteCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cloud account params
func (o *DeleteCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cloud account params
func (o *DeleteCloudAccountParams) WithHTTPClient(client *http.Client) *DeleteCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cloud account params
func (o *DeleteCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete cloud account params
func (o *DeleteCloudAccountParams) WithAPIVersion(aPIVersion *string) *DeleteCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete cloud account params
func (o *DeleteCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete cloud account params
func (o *DeleteCloudAccountParams) WithID(id string) *DeleteCloudAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete cloud account params
func (o *DeleteCloudAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
