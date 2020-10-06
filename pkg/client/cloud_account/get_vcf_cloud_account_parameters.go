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

// NewGetVcfCloudAccountParams creates a new GetVcfCloudAccountParams object
// with the default values initialized.
func NewGetVcfCloudAccountParams() *GetVcfCloudAccountParams {
	var ()
	return &GetVcfCloudAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVcfCloudAccountParamsWithTimeout creates a new GetVcfCloudAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVcfCloudAccountParamsWithTimeout(timeout time.Duration) *GetVcfCloudAccountParams {
	var ()
	return &GetVcfCloudAccountParams{

		timeout: timeout,
	}
}

// NewGetVcfCloudAccountParamsWithContext creates a new GetVcfCloudAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVcfCloudAccountParamsWithContext(ctx context.Context) *GetVcfCloudAccountParams {
	var ()
	return &GetVcfCloudAccountParams{

		Context: ctx,
	}
}

// NewGetVcfCloudAccountParamsWithHTTPClient creates a new GetVcfCloudAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVcfCloudAccountParamsWithHTTPClient(client *http.Client) *GetVcfCloudAccountParams {
	var ()
	return &GetVcfCloudAccountParams{
		HTTPClient: client,
	}
}

/*GetVcfCloudAccountParams contains all the parameters to send to the API endpoint
for the get vcf cloud account operation typically these are written to a http.Request
*/
type GetVcfCloudAccountParams struct {

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

// WithTimeout adds the timeout to the get vcf cloud account params
func (o *GetVcfCloudAccountParams) WithTimeout(timeout time.Duration) *GetVcfCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vcf cloud account params
func (o *GetVcfCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vcf cloud account params
func (o *GetVcfCloudAccountParams) WithContext(ctx context.Context) *GetVcfCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vcf cloud account params
func (o *GetVcfCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vcf cloud account params
func (o *GetVcfCloudAccountParams) WithHTTPClient(client *http.Client) *GetVcfCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vcf cloud account params
func (o *GetVcfCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get vcf cloud account params
func (o *GetVcfCloudAccountParams) WithAPIVersion(aPIVersion *string) *GetVcfCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get vcf cloud account params
func (o *GetVcfCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get vcf cloud account params
func (o *GetVcfCloudAccountParams) WithID(id string) *GetVcfCloudAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vcf cloud account params
func (o *GetVcfCloudAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVcfCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
