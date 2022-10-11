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
	"github.com/go-openapi/strfmt"
)

// NewEnumeratePrivateImagesParams creates a new EnumeratePrivateImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnumeratePrivateImagesParams() *EnumeratePrivateImagesParams {
	return &EnumeratePrivateImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnumeratePrivateImagesParamsWithTimeout creates a new EnumeratePrivateImagesParams object
// with the ability to set a timeout on a request.
func NewEnumeratePrivateImagesParamsWithTimeout(timeout time.Duration) *EnumeratePrivateImagesParams {
	return &EnumeratePrivateImagesParams{
		timeout: timeout,
	}
}

// NewEnumeratePrivateImagesParamsWithContext creates a new EnumeratePrivateImagesParams object
// with the ability to set a context for a request.
func NewEnumeratePrivateImagesParamsWithContext(ctx context.Context) *EnumeratePrivateImagesParams {
	return &EnumeratePrivateImagesParams{
		Context: ctx,
	}
}

// NewEnumeratePrivateImagesParamsWithHTTPClient creates a new EnumeratePrivateImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnumeratePrivateImagesParamsWithHTTPClient(client *http.Client) *EnumeratePrivateImagesParams {
	return &EnumeratePrivateImagesParams{
		HTTPClient: client,
	}
}

/*
EnumeratePrivateImagesParams contains all the parameters to send to the API endpoint

	for the enumerate private images operation.

	Typically these are written to a http.Request.
*/
type EnumeratePrivateImagesParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   Id of cloud account to enumerate
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enumerate private images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnumeratePrivateImagesParams) WithDefaults() *EnumeratePrivateImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enumerate private images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnumeratePrivateImagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enumerate private images params
func (o *EnumeratePrivateImagesParams) WithTimeout(timeout time.Duration) *EnumeratePrivateImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enumerate private images params
func (o *EnumeratePrivateImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enumerate private images params
func (o *EnumeratePrivateImagesParams) WithContext(ctx context.Context) *EnumeratePrivateImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enumerate private images params
func (o *EnumeratePrivateImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enumerate private images params
func (o *EnumeratePrivateImagesParams) WithHTTPClient(client *http.Client) *EnumeratePrivateImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enumerate private images params
func (o *EnumeratePrivateImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the enumerate private images params
func (o *EnumeratePrivateImagesParams) WithAPIVersion(aPIVersion *string) *EnumeratePrivateImagesParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the enumerate private images params
func (o *EnumeratePrivateImagesParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the enumerate private images params
func (o *EnumeratePrivateImagesParams) WithID(id string) *EnumeratePrivateImagesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the enumerate private images params
func (o *EnumeratePrivateImagesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EnumeratePrivateImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
