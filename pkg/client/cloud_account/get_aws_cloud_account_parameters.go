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

// NewGetAwsCloudAccountParams creates a new GetAwsCloudAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAwsCloudAccountParams() *GetAwsCloudAccountParams {
	return &GetAwsCloudAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAwsCloudAccountParamsWithTimeout creates a new GetAwsCloudAccountParams object
// with the ability to set a timeout on a request.
func NewGetAwsCloudAccountParamsWithTimeout(timeout time.Duration) *GetAwsCloudAccountParams {
	return &GetAwsCloudAccountParams{
		timeout: timeout,
	}
}

// NewGetAwsCloudAccountParamsWithContext creates a new GetAwsCloudAccountParams object
// with the ability to set a context for a request.
func NewGetAwsCloudAccountParamsWithContext(ctx context.Context) *GetAwsCloudAccountParams {
	return &GetAwsCloudAccountParams{
		Context: ctx,
	}
}

// NewGetAwsCloudAccountParamsWithHTTPClient creates a new GetAwsCloudAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAwsCloudAccountParamsWithHTTPClient(client *http.Client) *GetAwsCloudAccountParams {
	return &GetAwsCloudAccountParams{
		HTTPClient: client,
	}
}

/*
GetAwsCloudAccountParams contains all the parameters to send to the API endpoint

	for the get aws cloud account operation.

	Typically these are written to a http.Request.
*/
type GetAwsCloudAccountParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the Cloud Account
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get aws cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAwsCloudAccountParams) WithDefaults() *GetAwsCloudAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get aws cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAwsCloudAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get aws cloud account params
func (o *GetAwsCloudAccountParams) WithTimeout(timeout time.Duration) *GetAwsCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get aws cloud account params
func (o *GetAwsCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get aws cloud account params
func (o *GetAwsCloudAccountParams) WithContext(ctx context.Context) *GetAwsCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get aws cloud account params
func (o *GetAwsCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get aws cloud account params
func (o *GetAwsCloudAccountParams) WithHTTPClient(client *http.Client) *GetAwsCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get aws cloud account params
func (o *GetAwsCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get aws cloud account params
func (o *GetAwsCloudAccountParams) WithAPIVersion(aPIVersion *string) *GetAwsCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get aws cloud account params
func (o *GetAwsCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get aws cloud account params
func (o *GetAwsCloudAccountParams) WithID(id string) *GetAwsCloudAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get aws cloud account params
func (o *GetAwsCloudAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAwsCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
