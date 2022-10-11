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

// NewGetDataCollectorParams creates a new GetDataCollectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDataCollectorParams() *GetDataCollectorParams {
	return &GetDataCollectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDataCollectorParamsWithTimeout creates a new GetDataCollectorParams object
// with the ability to set a timeout on a request.
func NewGetDataCollectorParamsWithTimeout(timeout time.Duration) *GetDataCollectorParams {
	return &GetDataCollectorParams{
		timeout: timeout,
	}
}

// NewGetDataCollectorParamsWithContext creates a new GetDataCollectorParams object
// with the ability to set a context for a request.
func NewGetDataCollectorParamsWithContext(ctx context.Context) *GetDataCollectorParams {
	return &GetDataCollectorParams{
		Context: ctx,
	}
}

// NewGetDataCollectorParamsWithHTTPClient creates a new GetDataCollectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDataCollectorParamsWithHTTPClient(client *http.Client) *GetDataCollectorParams {
	return &GetDataCollectorParams{
		HTTPClient: client,
	}
}

/*
GetDataCollectorParams contains all the parameters to send to the API endpoint

	for the get data collector operation.

	Typically these are written to a http.Request.
*/
type GetDataCollectorParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the Data Collector.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get data collector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataCollectorParams) WithDefaults() *GetDataCollectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get data collector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataCollectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get data collector params
func (o *GetDataCollectorParams) WithTimeout(timeout time.Duration) *GetDataCollectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get data collector params
func (o *GetDataCollectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get data collector params
func (o *GetDataCollectorParams) WithContext(ctx context.Context) *GetDataCollectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get data collector params
func (o *GetDataCollectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get data collector params
func (o *GetDataCollectorParams) WithHTTPClient(client *http.Client) *GetDataCollectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get data collector params
func (o *GetDataCollectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get data collector params
func (o *GetDataCollectorParams) WithAPIVersion(aPIVersion *string) *GetDataCollectorParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get data collector params
func (o *GetDataCollectorParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get data collector params
func (o *GetDataCollectorParams) WithID(id string) *GetDataCollectorParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get data collector params
func (o *GetDataCollectorParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDataCollectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
