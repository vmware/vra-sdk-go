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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewEnumerateAzureRegionsAsyncParams creates a new EnumerateAzureRegionsAsyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnumerateAzureRegionsAsyncParams() *EnumerateAzureRegionsAsyncParams {
	return &EnumerateAzureRegionsAsyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnumerateAzureRegionsAsyncParamsWithTimeout creates a new EnumerateAzureRegionsAsyncParams object
// with the ability to set a timeout on a request.
func NewEnumerateAzureRegionsAsyncParamsWithTimeout(timeout time.Duration) *EnumerateAzureRegionsAsyncParams {
	return &EnumerateAzureRegionsAsyncParams{
		timeout: timeout,
	}
}

// NewEnumerateAzureRegionsAsyncParamsWithContext creates a new EnumerateAzureRegionsAsyncParams object
// with the ability to set a context for a request.
func NewEnumerateAzureRegionsAsyncParamsWithContext(ctx context.Context) *EnumerateAzureRegionsAsyncParams {
	return &EnumerateAzureRegionsAsyncParams{
		Context: ctx,
	}
}

// NewEnumerateAzureRegionsAsyncParamsWithHTTPClient creates a new EnumerateAzureRegionsAsyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnumerateAzureRegionsAsyncParamsWithHTTPClient(client *http.Client) *EnumerateAzureRegionsAsyncParams {
	return &EnumerateAzureRegionsAsyncParams{
		HTTPClient: client,
	}
}

/*
EnumerateAzureRegionsAsyncParams contains all the parameters to send to the API endpoint

	for the enumerate azure regions async operation.

	Typically these are written to a http.Request.
*/
type EnumerateAzureRegionsAsyncParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Azure region enumeration specification
	*/
	Body *models.CloudAccountAzureRegionEnumerationSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enumerate azure regions async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnumerateAzureRegionsAsyncParams) WithDefaults() *EnumerateAzureRegionsAsyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enumerate azure regions async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnumerateAzureRegionsAsyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enumerate azure regions async params
func (o *EnumerateAzureRegionsAsyncParams) WithTimeout(timeout time.Duration) *EnumerateAzureRegionsAsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enumerate azure regions async params
func (o *EnumerateAzureRegionsAsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enumerate azure regions async params
func (o *EnumerateAzureRegionsAsyncParams) WithContext(ctx context.Context) *EnumerateAzureRegionsAsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enumerate azure regions async params
func (o *EnumerateAzureRegionsAsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enumerate azure regions async params
func (o *EnumerateAzureRegionsAsyncParams) WithHTTPClient(client *http.Client) *EnumerateAzureRegionsAsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enumerate azure regions async params
func (o *EnumerateAzureRegionsAsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the enumerate azure regions async params
func (o *EnumerateAzureRegionsAsyncParams) WithAPIVersion(aPIVersion *string) *EnumerateAzureRegionsAsyncParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the enumerate azure regions async params
func (o *EnumerateAzureRegionsAsyncParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the enumerate azure regions async params
func (o *EnumerateAzureRegionsAsyncParams) WithBody(body *models.CloudAccountAzureRegionEnumerationSpecification) *EnumerateAzureRegionsAsyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the enumerate azure regions async params
func (o *EnumerateAzureRegionsAsyncParams) SetBody(body *models.CloudAccountAzureRegionEnumerationSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EnumerateAzureRegionsAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
