// Code generated by go-swagger; DO NOT EDIT.

package data_collector

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

// NewCreateDataCollectorParams creates a new CreateDataCollectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDataCollectorParams() *CreateDataCollectorParams {
	return &CreateDataCollectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDataCollectorParamsWithTimeout creates a new CreateDataCollectorParams object
// with the ability to set a timeout on a request.
func NewCreateDataCollectorParamsWithTimeout(timeout time.Duration) *CreateDataCollectorParams {
	return &CreateDataCollectorParams{
		timeout: timeout,
	}
}

// NewCreateDataCollectorParamsWithContext creates a new CreateDataCollectorParams object
// with the ability to set a context for a request.
func NewCreateDataCollectorParamsWithContext(ctx context.Context) *CreateDataCollectorParams {
	return &CreateDataCollectorParams{
		Context: ctx,
	}
}

// NewCreateDataCollectorParamsWithHTTPClient creates a new CreateDataCollectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDataCollectorParamsWithHTTPClient(client *http.Client) *CreateDataCollectorParams {
	return &CreateDataCollectorParams{
		HTTPClient: client,
	}
}

/*
CreateDataCollectorParams contains all the parameters to send to the API endpoint

	for the create data collector operation.

	Typically these are written to a http.Request.
*/
type CreateDataCollectorParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create data collector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDataCollectorParams) WithDefaults() *CreateDataCollectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create data collector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDataCollectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create data collector params
func (o *CreateDataCollectorParams) WithTimeout(timeout time.Duration) *CreateDataCollectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create data collector params
func (o *CreateDataCollectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create data collector params
func (o *CreateDataCollectorParams) WithContext(ctx context.Context) *CreateDataCollectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create data collector params
func (o *CreateDataCollectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create data collector params
func (o *CreateDataCollectorParams) WithHTTPClient(client *http.Client) *CreateDataCollectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create data collector params
func (o *CreateDataCollectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create data collector params
func (o *CreateDataCollectorParams) WithAPIVersion(aPIVersion *string) *CreateDataCollectorParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create data collector params
func (o *CreateDataCollectorParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDataCollectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
